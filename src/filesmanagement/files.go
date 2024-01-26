package filesmanagement

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func Creationfile(title string, content string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Création
	if _, err := os.Stat(title); err == nil {
		fmt.Println("Le fichier existe déjà avant le test.")
	} else if title == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		fichier, _ := os.Create(title)
		io.WriteString(fichier, content)
		fichier.Close()
	}
}

func Readfile(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Lecture
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
	} else if title == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		read, _ := ioutil.ReadFile(title)
		fmt.Println(string(read))
	}
}

func Renamefile(title string, newtitle string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Renommer
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
	} else if _, err := os.Stat(newtitle); err == nil {
		fmt.Println("Le fichier existe déjà avant le test.")
	} else if title == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if newtitle == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		os.Rename(title, newtitle)
	}
}

func Changecontenufile(title string, content string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	// Changer contenu
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
	} else if title == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		fichier, _ := os.OpenFile(title, os.O_WRONLY|os.O_TRUNC, 0666)
		io.WriteString(fichier, content)
		fichier.Close()
	}
}

func DeleteFile(title string) {
	invalidCharsRegex := regexp.MustCompile(`[<>:"/\|?*]`)
	if _, err := os.Stat(title); err != nil {
		fmt.Println("Fichier Inexistant.")
	} else if title == "" {
		fmt.Println("Le nom du fichier ne peut pas être vide.")
	} else if invalidCharsRegex.MatchString(title) {
		fmt.Println("Les noms de fichiers ne peuvent pas contenir certains characteres spéciaux")
	} else {
		os.Remove(title)
	}
}
