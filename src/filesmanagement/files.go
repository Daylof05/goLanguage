package filesmanagement

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Creationfile(title string, content string) {
	// Création
	fichier, _ := os.Create(title)
	io.WriteString(fichier, content)
	fichier.Close()
}

func Readfile(title string) {
	// Lecture
	read, _ := ioutil.ReadFile(title)
	fmt.Println(string(read))
}

func Renamefile(title string, newtitle string) {
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
