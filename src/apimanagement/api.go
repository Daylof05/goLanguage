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
	router.PATCH("/renamefile/:oldTitle/:newTitle", renameFileAPI)
	router.PATCH("/changecontentfile/:title", changeContentFileAPI)
	router.DELETE("deletefile/:title", deleteFileAPI)
	router.Run("localhost:32244")
}

func createFolderAPI(c *gin.Context) {
	var newfolder Folder

	if err := c.BindJSON(&newfolder); err != nil {
		return
	}
	folderManager := foldersmanagement.CmdFolderManager{}
	err := folderManager.CreateFolder(newfolder.Title)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func deleteFolderAPI(c *gin.Context) {
	title := c.Param("title")

	folderManager := foldersmanagement.CmdFolderManager{}
	err := folderManager.DeleteFolder(title)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

	folderManager := foldersmanagement.CmdFolderManager{}
	if oldTitle == "" || newTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ancien ou nouveau nom du dossier manquant"})
		return
	}

	err := folderManager.RenameFolder(oldTitle, newTitle)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dossier renommé avec succès"})
}

func createFileAPI(c *gin.Context) {
	var newfile Files

	fileManager := filesmanagement.CmdFileManager{}
	if err := c.BindJSON(&newfile); err != nil {
		return
	}
	err := fileManager.Creationfile(newfile.Title, newfile.Content)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func deleteFileAPI(c *gin.Context) {
	title := c.Param("title")

	fileManager := filesmanagement.CmdFileManager{}
	err := fileManager.DeleteFile(title)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func renameFileAPI(c *gin.Context) {
	newTitle := c.Param("newTitle")
	oldTitle := c.Param("oldTitle")

	fileManager := filesmanagement.CmdFileManager{}
	if oldTitle == "" || newTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ancien ou nouveau nom du dossier manquant"})
		return
	}

	err := fileManager.Renamefile(oldTitle, newTitle)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dossier renommé avec succès"})
}

type ChangeContentRequest struct {
	NewContent string `json:"newContent"`
}

func changeContentFileAPI(c *gin.Context) {
	title := c.Param("title")
	var req ChangeContentRequest

	fileManager := filesmanagement.CmdFileManager{}
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := fileManager.Changecontenufile(title, req.NewContent)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Contenu du fichier mis à jour"})
}

func readContentFileAPI(c *gin.Context) {
	title := c.Param("title")

	fileManager := filesmanagement.CmdFileManager{}

	sql.Connection()
	sql.WriteUpdate("filesmanagement -> Readfile", title, "Success")
	err := fileManager.Readfile(title)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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
