package main

import (
	"os"
	"project/filesmanagement"
	"project/foldersmanagement"
)

func main() {

	switch os.Args[1] {
	case "createfolder":
		foldersmanagement.CreateFolder(os.Args[2])
	case "openfolder":
		foldersmanagement.OpenFolder(os.Args[2])
	case "renamefolder":
		foldersmanagement.RenameFolder(os.Args[2], os.Args[3])
	case "deletefolder":
		foldersmanagement.DeleteFolder(os.Args[2])
	case "createfile":
		filesmanagement.Creationfile(os.Args[2], os.Args[3])
	case "readfile":
		filesmanagement.Readfile(os.Args[2])
	case "renamefile":
		filesmanagement.Renamefile(os.Args[2], os.Args[3])
	case "changecontenufile":
		filesmanagement.Changecontenufile(os.Args[2], os.Args[3])
	case "deletefile":
		filesmanagement.DeleteFile(os.Args[2])
	}
}
