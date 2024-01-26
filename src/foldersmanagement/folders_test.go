package foldersmanagement

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestCreation(t *testing.T) {
	os.Mkdir("TestCreate", 0755)
	if _, err := os.Stat("TestCreate"); err != nil {
		fmt.Println("Folder does not exist")
		CreateFolder("Created")
	}
	os.RemoveAll("TestCreate")
	os.RemoveAll("Created")
}

func TestCreationInvalid(t *testing.T) {
	os.Mkdir("TestCreateInvalid", 0755)
	if _, err := os.Stat("TestCreateInvalid"); err == nil {
		t.Error("Folder already exists")
	}
	os.RemoveAll("TestCreateInvalid")
}

func TestCreationEmpty(t *testing.T) {
	os.Mkdir("TestCreateEmpty", 0755)
	if _, err := os.Stat(""); err == nil {
		t.Error("Folder name can't be empty")
	}
	os.RemoveAll("TestCreateEmpty")
}

func TestCreationName(t *testing.T) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if invalidCharsRegex.MatchString("TestCreationName") {
		t.Error("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	}
}

func TestOpen(t *testing.T) {
	os.Mkdir("TestOpen", 0755)
	if _, err := os.Stat("TestOpen"); err == nil {
		fmt.Println("Folder open")
		OpenFolder("TestOpen")
	}
	os.RemoveAll("TestOpen")
}

func TestOpenInvalid(t *testing.T) {
	os.Mkdir("TestOpenInvalid", 0755)
	if _, err := os.Stat("Testinvalid"); err != nil {
		t.Error("Folder does not exist")
	}
	os.Remove("TestOpenInvalid")
}

func TestRename(t *testing.T) {
	os.Mkdir("TestRename", 0755)
	if _, err := os.Stat("TestRename"); err == nil {
		fmt.Println("Folder renamed")
		RenameFolder("TestRename", "NewTest")
	}
	os.RemoveAll("NewTest")
}

func TestRenameInvalid(t *testing.T) {
	os.Mkdir("TestRenameInvalid", 0755)
	if _, err := os.Stat("Testinvalid"); err != nil {
		t.Error("Folder does not exist")
	}
	os.RemoveAll("TestRenameInvalid")
}

func TestDelete(t *testing.T) {
	os.Mkdir("TestDelete", 0755)
	if _, err := os.Stat("TestDelete"); err == nil {
		fmt.Println("Folder deleted")
		DeleteFolder("TestDelete")
	}
}

func TestDeleteInvalid(t *testing.T) {
	os.Mkdir("TestDeleteInvalid", 0755)
	if _, err := os.Stat("Testinvalid"); err != nil {
		t.Error("Folder does not exist")
	}
	os.RemoveAll("TestDeleteInvalid")
}
