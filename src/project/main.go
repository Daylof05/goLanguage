package main

import (
	"fmt"
	"log"
	"os"
	"project/apiclient"
	"project/apimanagement"
	"project/filesmanagement"
	"project/foldersmanagement"
	"project/sql"
)

var onlineMode bool = true

func main() {
	var folderManager foldersmanagement.FolderManager
	var fileManager filesmanagement.FilesManager

	if onlineMode {
		folderManager = apiclient.APIFolderManager{}
		fileManager = apiclient.APIFileManager{}
	} else {
		folderManager = foldersmanagement.CmdFolderManager{}
		fileManager = filesmanagement.CmdFileManager{}
	}
	switch os.Args[1] {
	case "createfolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> CreateFolder", "error", "Fail")
		} else {
			err := folderManager.CreateFolder(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "openfolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> OpenFolder", "error", "Fail")
		} else {
			err := folderManager.OpenFolder(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "renamefolder":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> RenameFolder", "error", "Fail")
		} else {
			err := folderManager.RenameFolder(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "deletefolder":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("foldersmanagement -> DeleteFolder", "error", "Fail")
		} else {
			err := folderManager.DeleteFolder(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "createfile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> Creationfile", "error", "Fail")
		} else {
			err := fileManager.Creationfile(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "readfile":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> Readfile", "error", "Fail")
		} else {
			err := fileManager.Readfile(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "renamefile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> Renamefile", "error", "Fail")
		} else {
			err := fileManager.Renamefile(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "changecontenufile":
		if len(os.Args) != 4 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> Changecontenufile", "error", "Fail")
		} else {
			err := fileManager.Changecontenufile(os.Args[2], os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "deletefile":
		if len(os.Args) != 3 {
			fmt.Println("Erreur lors de la saisie")
			sql.Connection()
			sql.WriteUpdate("filesmanagement -> DeleteFile", "error", "Fail")
		} else {
			err := fileManager.DeleteFile(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
		}
	case "history":
		sql.Connection()
		sql.WriteUpdate("sql -> History", "none", "Success")
		updates, err := sql.PrintUpdates()
		if err != nil {
			log.Fatal(err)
		}
		for _, update := range updates {
			fmt.Printf("Update found: %+v\n", update)
		}
	case "API":
		apimanagement.ConnectionAPI()
	}
}
