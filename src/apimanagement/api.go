package apimanagement

import (
	"fmt"
	"net/http"
	"os"
	"project/filesmanagement"
	"project/foldersmanagement"
	"project/sql"

	"github.com/gin-gonic/gin"
)

type Folder struct {
	ID    string `JSON:"id"`
	Title string `JSON:"title"`
}

type Files struct {
	ID      string `JSON:"id"`
	Title   string `JSON:"title"`
	Content string `JSON:"content"`
}

func ConnectionAPI() {
	router := gin.Default()
	router.POST("/createfolder", createFolderAPI)
	router.GET("/openfolder/:title", openFolderAPI)
	router.PATCH("/renamefolder/:oldTitle/:newTitle", renameFolderAPI)
	router.DELETE("/deletefolder/:title", deleteFolderAPI)
	router.POST("/createfile", createFileAPI)
	router.GET("/readfile/:title", readContentFileAPI)
	router.GET("/history", historyAPI)
	router.PATCH("/renamefile/:oldTitle", renameFileAPI)
	router.PATCH("/changecontentfile/:title", changeContentFileAPI)
	router.DELETE("deletefile/:title", deleteFileAPI)
	router.Run("localhost:32244")
}

func createFolderAPI(c *gin.Context) {
	var newfolder Folder

	if err := c.BindJSON(&newfolder); err != nil {
		return
	}
	foldersmanagement.CreateFolder(newfolder.Title)
}

func deleteFolderAPI(c *gin.Context) {
	title := c.Param("title")

	foldersmanagement.DeleteFolder(title)
}

func openFolderAPI(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nom du dossier manquant"})
		return
	}

	files, err := os.ReadDir(title)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du dossier :", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture du dossier"})
		return
	}

	fileNames := make([]string, len(files))
	for i, file := range files {
		fileNames[i] = file.Name()
	}

	c.JSON(http.StatusOK, gin.H{"files": fileNames})
}

func renameFolderAPI(c *gin.Context) {
	oldTitle := c.Param("oldTitle")
	newTitle := c.Param("newTitle")
	if oldTitle == "" || newTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ancien ou nouveau nom du dossier manquant"})
		return
	}
	foldersmanagement.RenameFolder(oldTitle, newTitle)
	c.JSON(http.StatusOK, gin.H{"message": "Dossier renommé avec succès"})
}

func createFileAPI(c *gin.Context) {
	var newfile Files

	if err := c.BindJSON(&newfile); err != nil {
		return
	}
	filesmanagement.Creationfile(newfile.Title, newfile.Content)
}

func deleteFileAPI(c *gin.Context) {
	title := c.Param("title")

	filesmanagement.DeleteFile(title)
}

type RenameFileRequest struct {
	NewTitle string `json:"newTitle"`
}

func renameFileAPI(c *gin.Context) {
	oldTitle := c.Param("oldTitle")

	var req RenameFileRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filesmanagement.Renamefile(oldTitle, req.NewTitle)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Fichier renommé"})
}

type ChangeContentRequest struct {
	NewContent string `json:"newContent"`
}

func changeContentFileAPI(c *gin.Context) {
	title := c.Param("title")
	var req ChangeContentRequest

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filesmanagement.Changecontenufile(title, req.NewContent)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Contenu du fichier mis à jour"})
}

func readContentFileAPI(c *gin.Context) {
	title := c.Param("title")

	content, err := os.ReadFile(title)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sql.Connection()
	sql.WriteUpdate("filesmanagement -> Readfile", title, "Success")

	c.String(http.StatusOK, string(content))
}

func historyAPI(c *gin.Context) {
	sql.Connection()
	sql.WriteUpdate("sql -> History", "none", "Success")
	updates, err := sql.PrintUpdates()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updates)
}
