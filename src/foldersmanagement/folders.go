package foldersmanagement

import (
	"fmt"
	"os"
)

func CreateFolder(title string) {
	os.Mkdir(title, 0755)
}

func OpenFolder(title string) {
	fmt.Println(os.ReadDir(title))
}

func RenameFolder(oldTitle string, newTitle string) {
	os.Rename(oldTitle, newTitle)
}

func DeleteFolder(title string) {
	os.RemoveAll(title)
}
