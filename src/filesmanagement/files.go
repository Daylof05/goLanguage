package filesmanagement

import (
	"fmt"
	"io"
	"os"
	"project/sql"
	"regexp"
)

type FilesManager interface {
	Creationfile(title string, content string) error
	Readfile(title string) error
	Renamefile(oldTitle string, newTitle string) error
	Changecontenufile(title string, content string) error
	DeleteFile(title string) error
}

type CmdFileManager struct{}

func (fm CmdFileManager) Creationfile(title string, content string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err == nil {
		fmt.Println("File already exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Files can't have special characters")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Fail")
	} else {
		fichier, _ := os.Create(title)
		io.WriteString(fichier, content)
		fichier.Close()
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Success")
	}
	return nil
}

func (fm CmdFileManager) Readfile(title string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("File doesn't exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Fail")
	} else {
		read, _ := os.ReadFile(title)
		fmt.Println(string(read))
		// fmt.Println(os.ReadFile(title))
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Success")
	}
	return nil
}

func (fm CmdFileManager) Renamefile(oldTitle string, newTitle string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(oldTitle); err != nil {
		fmt.Println("File doesn't exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", oldTitle+"|"+newTitle, "Fail")
	} else if _, err := os.Stat(newTitle); err == nil {
		fmt.Println("File already exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", oldTitle+"|"+newTitle, "Fail")
	} else if invalidCharsRegex.MatchString(oldTitle) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", oldTitle+"|"+newTitle, "Fail")
	} else {
		os.Rename(oldTitle, newTitle)
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", oldTitle+"|"+newTitle, "Success")
	}
	return nil
}

func (fm CmdFileManager) Changecontenufile(title string, content string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("File doesn't exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Fail")
	} else {
		fichier, _ := os.OpenFile(title, os.O_WRONLY|os.O_TRUNC, 0666)
		io.WriteString(fichier, content)
		fichier.Close()
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Success")
	}
	return nil
}

func (fm CmdFileManager) DeleteFile(title string) error {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("File doesn't exist")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Folders can't have special characters")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Fail")
	} else {
		os.Remove(title)
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Success")
	}
	return nil
}
