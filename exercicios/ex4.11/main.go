package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobook/ch4/github"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var token = os.Getenv("GITHUB_TOKEN")

func main() {
	GithubIssuesCrud()
}

// exercicio 4.11
// crie uma ferramenta de linha de comando que permita
// aos usuarios criar ler atualizar e fechar issues do
// Github a partir da linha de comando, chamando seu editor
// de texto preferido quando houver a necessidade de de fornecer
// uma quantidade subtancial de texto de entrada

func GithubIssuesCrud() {
	args := os.Args[1:]
	cmd, owner, repo := args[0], args[1], args[2]
	switch cmd {
	case "create":
		fmt.Println("create issue")
		body, err := GetTextEditor("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro call editor: %s", err)
			os.Exit(1)
		}
		issue, err := CreateIssue(owner, repo, body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro creating issue: %s", err)
			os.Exit(1)
		}
		fmt.Println(issue)
	case "read":
		fmt.Println("Read issue")
		issue, err := ReadIssue(owner, repo, args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro reading issue: %s", err)
			os.Exit(1)
		}
		fmt.Println(issue)
	case "update":
		issue, err := ReadIssue(owner, repo, args[3])
		fmt.Println("update issue")
		body, err := GetTextEditor(issue.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro call editor: %s", err)
			os.Exit(1)
		}
		issue, err = UpdateIssue(owner, repo, args[3], "MInha issue", body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro updating issue: %s", err)
			os.Exit(1)
		}
		fmt.Println(issue)
	case "delete":
		fmt.Println("delete issue")
		err := DeleteIssue(owner, repo, args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro deleting issue: %s", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Not a valid CRUD command")
		return
	}
}

func GetTextEditor(body string) (string, error) {
	tmpFile, _ := os.CreateTemp("", "issue_body_*.md")
	defer os.Remove(tmpFile.Name())

	if body != "" {
		os.WriteFile(tmpFile.Name(), []byte(body), 0644)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, tmpFile.Name())

	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Run()

	content, _ := os.ReadFile(tmpFile.Name())
	return string(content), nil
}
func CreateIssue(owner, repo, body string) (*github.Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	newIssue := github.Issue{Title: "New Issue", Body: body}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(newIssue); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, buf)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Erro ao criar issue: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func ReadIssue(owner, repo, number string) (*github.Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro ao ler issue: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func UpdateIssue(owner, repo, number string, title, body string) (*github.Issue, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
	req, err := http.NewRequest("PATCH", url, strings.NewReader(fmt.Sprintf(`{"title": "%s", "body": "%s"}`, title, body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erro ao atualizar issue: %s", resp.Status)
	}

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func DeleteIssue(owner, repo, number string) error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
	data := `{"state":"closed"}`
	req, _ := http.NewRequest("PATCH", url, strings.NewReader(data))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Erro ao deletar issue: %s", resp.Status)
	}

	return nil
}
