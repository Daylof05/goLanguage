// main_test.go
package filesmanagement

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreationfileExistant(t *testing.T) {

	// Test 1 le fichier existe
	if _, err := os.Stat("Aurevoir.txt"); err == nil {
		t.Error("Le fichier existe déjà avant le test.")
	}

}
func TestCreationfile(t *testing.T) {

	// Test 2  le fichier existe pas
	if _, err := os.Stat("TestCreationfile.txt"); err != nil {
		fmt.Println("Le fichier n'existait pas, il a donc été créer.")

		ioutil.WriteFile("TestCreationfile.txt", []byte("Texte du fichier créer pour le Creationfile"), 0644)
	}
	DeleteFile("TestCreationfile.txt")
}
func TestCreationfileEmptyTitle(t *testing.T) {

	title := ""

	// Test 3 Vérifier si le fichier à un nom vide
	if title == "" {
		t.Error("Le nom du fichier ne peut pas être vide.")
	}
}

func TestReadfileexist(t *testing.T) {
	ioutil.WriteFile("TestReadFileexist.txt", []byte("Texte du fichier créer pour le Readfileexist"), 0644)
	//Test 1 fichier existe

	if _, err := os.Stat("TestReadFileexist.txt"); err == nil {
		fmt.Println("Le fichier existe.")
		Readfile("TestReadFileexist.txt")

	}
	DeleteFile("TestReadFileexist.txt")

}

func TestReadfilenotexist(t *testing.T) {

	//Test 2 fichier existe pas

	if _, err := os.Stat("Aurevoir.txt"); err != nil {
		t.Error("Fichier Inexistant.")

	}

}

func TestRenamefileexist(t *testing.T) {
	ioutil.WriteFile("TestRenameFileexist.txt", []byte("Texte du fichier créer pour le Renamefileexist"), 0644)
	//Test 1 fichier existe

	if _, err := os.Stat("TestRenameFileexist.txt"); err == nil {
		fmt.Println("Le fichier existe.")
		Renamefile("TestRenameFileexist.txt", "NouveauTestRenameFileexist.txt")

	}
	DeleteFile("NouveauTestRenameFileexist.txt")
}

func TestRenamefilenotexist(t *testing.T) {

	//Test 2 fichier existe pas

	if _, err := os.Stat("Aurevoir.txt"); err != nil {
		t.Error("Fichier Inexistant.")

	}

}

func TestChangecontenufilefileexist(t *testing.T) {
	ioutil.WriteFile("TestChangecontenufilefileexist.txt", []byte("Texte du fichier créer pour le Changecontenufilefileexist"), 0644)
	//Test 1 fichier existe

	if _, err := os.Stat("TestChangecontenufilefileexist.txt"); err == nil {
		fmt.Println("Le fichier existe.")
		Changecontenufile("TestChangecontenufilefileexist.txt", "Nouveau contenu du changecontenu")

	}
	DeleteFile("TestChangecontenufilefileexist.txt")
}

func TestChangecontenufilenotexist(t *testing.T) {

	//Test 2 fichier existe pas

	if _, err := os.Stat("Aurevoir.txt"); err != nil {
		t.Error("Fichier Inexistant.")

	}

}

func TestDeletefileexist(t *testing.T) {
	ioutil.WriteFile("TestDeleteFileexist.txt", []byte("Texte du fichier créer pour le Deletefileexist"), 0644)
	//Test 1 fichier existe

	if _, err := os.Stat("TestDeleteFileexist.txt"); err == nil {
		fmt.Println("Le fichier a été supprimé.")

	}
	DeleteFile("TestDeletefileexist.txt")
}

func TestDeletefilenotexist(t *testing.T) {

	//Test 2 fichier existe pas

	if _, err := os.Stat("Aurevoir.txt"); err != nil {
		t.Error("Fichier Inexistant.")

	}
}
