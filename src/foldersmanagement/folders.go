package foldersmanagement

import (
	"fmt"
	"os"
	"regexp"
)

func CreateFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err == nil {
		fmt.Println("Folder already exists")
	} else if title == "" {
		fmt.Println("Folder name can't be empty")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		os.Mkdir(title, 0755)
	}
}

func OpenFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		fmt.Println(os.ReadDir(title))
	}
}

func RenameFolder(oldTitle string, newTitle string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(oldTitle); err != nil {
		fmt.Println("Folder doesn't exist")
	} else if _, err := os.Stat(newTitle); err == nil {
		fmt.Println("This folder already exist")
	} else if invalidCharsRegex.MatchString(oldTitle) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
	} else if invalidCharsRegex.MatchString(newTitle) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		os.Rename(oldTitle, newTitle)
	}
}

func DeleteFolder(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Folder doesn't exist")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de dossiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		os.RemoveAll(title)
	}
}
