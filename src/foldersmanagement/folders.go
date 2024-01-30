package foldersmanagement

import (
	"fmt"
	"os"
	"project/sql"
	"regexp"
)

type FolderManager interface {
	CreateFolder(title string) error
	OpenFolder(title string) error
	RenameFolder(oldTitle string, newtitle string) error
	DeleteFolder(title string) error
}

type CmdFolderManager struct{}

func (fm CmdFolderManager) CreateFolder(title string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err == nil {
		fmt.Println("Folder already exists")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Fail")
	} else {
		os.Mkdir(title, 0755)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> CreateFolder", title, "Success")
	}
	return nil
}

func (fm CmdFolderManager) OpenFolder(title string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Fail")
	} else {
		fmt.Println(os.ReadDir(title))
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> OpenFolder", title, "Success")
	}
	return nil
}

func (fm CmdFolderManager) RenameFolder(oldTitle string, newTitle string) error {
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
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else if invalidCharsRegex.MatchString(newTitle) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Fail")
	} else {
		os.Rename(oldTitle, newTitle)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> RenameFolder", oldTitle+"|"+newTitle, "Success")
	}
	return nil
}

func (fm CmdFolderManager) DeleteFolder(title string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Fail")
	} else {
		os.RemoveAll(title)
		sql.Connection()
		sql.WriteUpdate("foldersmanagement -> DeleteFolder", title, "Success")
	}
	return nil
}
