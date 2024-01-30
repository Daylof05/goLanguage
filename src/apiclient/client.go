package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const server = "http://localhost:32244"

type APIFolderManager struct{}

type APIFileManager struct{}

func (afm APIFolderManager) CreateFolder(title string) error {
	folderData := map[string]string{"title": title}
	jsonData, err := json.Marshal(folderData)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du JSON : %v", err)
	}

	url := fmt.Sprintf("%s/createfolder", server)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête POST : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la création du dossier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afm APIFolderManager) RenameFolder(oldTitle, newTitle string) error {
	url := fmt.Sprintf("%s/renamefolder/%s/%s", server, oldTitle, newTitle)
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec du renommage du dossier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afm APIFolderManager) OpenFolder(title string) error {
	url := fmt.Sprintf("%s/openfolder/%s", server, title)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de l'ouverture du dossier, code d'état : %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	fmt.Println("Contenu du dossier:", string(body))

	return nil
}

func (afm APIFolderManager) DeleteFolder(title string) error {
	url := fmt.Sprintf("%s/deletefolder/%s", server, title)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la suppression du dossier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afim APIFileManager) Creationfile(title, content string) error {
	data := map[string]string{"title": title, "content": content}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du JSON : %v", err)
	}

	url := fmt.Sprintf("%s/createfile", server)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la création du fichier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afim APIFileManager) Readfile(title string) error {
	url := fmt.Sprintf("%s/readfile/%s", server, title)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la lecture du fichier, code d'état : %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	fmt.Println("Contenu du fichier:", string(body))
	return nil
}

func (afim APIFileManager) Renamefile(oldTitle, newTitle string) error {
	url := fmt.Sprintf("%s/renamefile/%s/%s", server, oldTitle, newTitle)
	req, err := http.NewRequest(http.MethodPatch, url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec du renommage du fichier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afim APIFileManager) Changecontenufile(title, content string) error {
	data := map[string]string{"title": title, "content": content}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du JSON : %v", err)
	}

	url := fmt.Sprintf("%s/changecontentfile", server)
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la modification du contenu du fichier, code d'état : %d", resp.StatusCode)
	}

	return nil
}

func (afim APIFileManager) DeleteFile(title string) error {
	url := fmt.Sprintf("%s/deletefile/%s", server, title)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erreur lors de l'envoi de la requête : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("échec de la suppression du fichier, code d'état : %d", resp.StatusCode)
	}

	return nil
}
