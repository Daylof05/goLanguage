package filesmanagement

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"project/sql"
	"regexp"
)

func Creationfile(title string, content string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Création
	if _, err := os.Stat(title); err == nil {
		fmt.Println("Le fichier existe déjà avant le test.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Fail")
	} else {
		fichier, _ := os.Create(title)
		io.WriteString(fichier, content)
		fichier.Close()
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Creationfile", title+"|"+content, "Success")
	}
}

func Readfile(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Lecture
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Fail")
	} else {
		read, _ := ioutil.ReadFile(title)
		fmt.Println(string(read))
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Readfile", title, "Success")
	}
}

func Renamefile(title string, newtitle string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Renommer
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", title+"|"+newtitle, "Fail")
	} else if _, err := os.Stat(newtitle); err == nil {
		fmt.Println("Le fichier existe déjà avant le test.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", title+"|"+newtitle, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", title+"|"+newtitle, "Fail")
	} else {
		os.Rename(title, newtitle)
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Renamefile", title+"|"+newtitle, "Success")
	}
}

func Changecontenufile(title string, content string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Changer contenu
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Fail")
	} else {
		fichier, _ := os.OpenFile(title, os.O_WRONLY|os.O_TRUNC, 0666)
		io.WriteString(fichier, content)
		fichier.Close()
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> Changecontenufile", title+"|"+content, "Success")
	}
}

func DeleteFile(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Fail")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Fail")
	} else {
		os.Remove(title)
		sql.Connection()
		sql.WriteUpdate("filesmanagement -> DeleteFile", title, "Success")
	}
}
