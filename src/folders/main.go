package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	case "createfile":
		creationfile(os.Args[2], os.Args[3])
	case "readfile":
		readfile(os.Args[2])
	case "renamefile":
		renamefile(os.Args[2], os.Args[3])
	case "changecontenufile":
		Changecontenufile(os.Args[2], os.Args[3])
	case "deletefile":
		DeleteFile(os.Args[2])
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

func creationfile(title string, content string) {
	// Création
	fichier, _ := os.Create(title)
	io.WriteString(fichier, content)
	fichier.Close()
}

func readfile(title string) {
	// Lecture
	read, _ := ioutil.ReadFile(title)
	fmt.Println(string(read))
}

func renamefile(title string, newtitle string) {
	// Renommer

	os.Rename(title, newtitle)
}

func Changecontenufile(title string, content string) {
	// Changer contenu
	fichier, _ := os.OpenFile(title, os.O_WRONLY|os.O_TRUNC, 0666)
	io.WriteString(fichier, content)
	fichier.Close()
}

func DeleteFile(title string) {
	os.Remove(title)
}
