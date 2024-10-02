package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var gist, codeword, cmd, logins string

type Comment struct {
	Body string `json:"body"`
	User struct {
		Login string `json:"login"`
	} `json:"user"`
}

func init() {
	flag.StringVar(&gist, "gist", "", "gist id to look on")
	flag.StringVar(&codeword, "codeword", "", "codeword to work with")
	flag.StringVar(&cmd, "cmd", "", "command to run")
	flag.StringVar(&logins, "logins", "", "possible name restrictions")

	flag.Parse()
}

func main() {
	if gist == "" || codeword == "" || cmd == "" {
		log.Fatalf("please provide all: gist, codeword and cmd")
	}

	comments, err := listGistComments(gist)
	if err != nil {
		log.Fatalf("error retrieving gist comments: %v", err)
	}
	command := strings.Split(cmd, " ")

	for _, comment := range comments {
		if strings.Contains(comment.Body, codeword) {
			if logins != "" && !slices.Contains(strings.Split(logins, ","), comment.User.Login) {
				break
			}

			if err := exec.Command(command[0], command[1:]...).Start(); err != nil {
				log.Fatalf("error executing command: %v", err)
			}

			log.Print("dead man trigger pulled!")
			os.Exit(0)
		}
	}

	log.Print("dead man still walking :3")
}

func listGistComments(id string) ([]Comment, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/gists/%s/comments", id), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	var comments []Comment
	if err := json.NewDecoder(res.Body).Decode(&comments); err != nil {
		return nil, fmt.Errorf("error decoding request: %w", err)
	}

	return comments, nil
}
