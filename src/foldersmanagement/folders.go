package foldersmanagement

import (
	"fmt"
	"os"
	"project/sql"
	"regexp"
)

func CreateFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err == nil {
		fmt.Println("Folder already exists")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Fail")
	} else {
		os.Mkdir(title, 0755)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Success")
	}
}

func OpenFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Fail")
	} else {
		fmt.Println(os.ReadDir(title))
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Success")
	}
}

func RenameFolder(oldTitle string, newTitle string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(oldTitle); err != nil {
		fmt.Println("Folder doesn't exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else if _, err := os.Stat(newTitle); err == nil {
		fmt.Println("This folder already exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else if invalidCharsRegex.MatchString(oldTitle) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else if invalidCharsRegex.MatchString(newTitle) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else {
		os.Rename(oldTitle, newTitle)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Success")
	}
}

func DeleteFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Fail")
	} else {
		os.RemoveAll(title)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Success")
	}
}
