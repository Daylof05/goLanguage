package main

import (
	"fmt"
	"log"
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

var folders = []Folder{
	{ID: "1", Title: "AncienDossier"},
}

func main() {

	router := gin.Default()
	router.GET("/folders", getFolders)
	router.POST("/createfolder", createFolderAPI)
	router.Run("localhost:32244")
	switch os.Args[1] {
	case "createfolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> CreateFolder", "error", "Fail")
		} else {
			foldersmanagement.CreateFolder(os.Args[2])
		}
	case "openfolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> OpenFolder", "error", "Fail")
		} else {
			foldersmanagement.OpenFolder(os.Args[2])
		}
	case "renamefolder":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> RenameFolder", "error", "Fail")
		} else {
			foldersmanagement.RenameFolder(os.Args[2], os.Args[3])
		}
	case "deletefolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> DeleteFolder", "error", "Fail")
		} else {
			foldersmanagement.DeleteFolder(os.Args[2])
		}
	case "createfile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> Creationfile", "error", "Fail")
		} else {
			filesmanagement.Creationfile(os.Args[2], os.Args[3])
		}
	case "readfile":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.WriteUpdate("filesmanagement -> Readfile", "error", "Fail")
		} else {
			filesmanagement.Readfile(os.Args[2])
		}
	case "renamefile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.WriteUpdate("filesmanagement -> Renamefile", "error", "Fail")
		} else {
			filesmanagement.Renamefile(os.Args[2], os.Args[3])
		}
	case "changecontenufile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.WriteUpdate("filesmanagement -> Changecontenufile", "error", "Fail")
		} else {
			filesmanagement.Changecontenufile(os.Args[2], os.Args[3])
		}
	case "deletefile":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.WriteUpdate("filesmanagement -> DeleteFile", "error", "Fail")
		} else {
			filesmanagement.DeleteFile(os.Args[2])
		}
	case "history":
		sql.Connection()
		sql.WriteUpdate("sql -> WriteUpdate", "none", "Success")
		updates, err := sql.PrintUpdates()
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			fmt.Printf("Update found: %+v\n", update)
		}
	}
}

func getFolders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, folders)
}

func createFolderAPI(c *gin.Context) {
	// var newfolder Folder

	// if err := c.BindJSON(&newfolder); err != nil {
	// 	return
	// }
	// folders = append(folders, newfolder)
	// c.IndentedJSON(http.StatusCreated, newfolder)
	// foldersmanagement.CreateFolder("NouveauDossierTest")
	var newfolder Folder

	if err := c.BindJSON(&newfolder); err != nil {
		return
	}
	foldersmanagement.CreateFolder(newfolder.Title)
}
