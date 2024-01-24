package main

import (
	"fmt"
	"os"
)

func main() {

	switch os.Args[1] {
	case "createfolder":
		createFolder(os.Args[2])
	case "openfolder":
		openFolder(os.Args[2])
	case "renamefolder":
		renameFolder(os.Args[2], os.Args[3])
	case "deletefolder":
		deleteFolder(os.Args[2])
	}
}

func createFolder(title string) {
	os.Mkdir(title, 0755)
}

func openFolder(title string) {
	fmt.Println(os.ReadDir(title))
}

func renameFolder(oldTitle string, newTitle string) {
	os.Rename(oldTitle, newTitle)
}

func deleteFolder(title string) {
	os.RemoveAll(title)
}
