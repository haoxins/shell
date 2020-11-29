package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/haoxins/tools/v2"
)

const defaultUser = "haoxins"

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Invalid params")
	}

	param := os.Args[1]

	if !strings.Contains(param, "/") {
		log.Fatalln("Invalid repo")
	}

	params := strings.Split(param, "/")
	if len(params) != 2 {
		log.Fatalln("Invalid params")
	}

	owner := params[0]
	repo := params[1]

	user := ""

	extra := ""
	if owner != "airwallex" {
		user = defaultUser
		extra = defaultUser
	}

	dest := fmt.Sprintf("git@github.com-%s:%s/%s.git", extra, owner, repo)
	fmt.Printf("Clone to: %s\n", dest)

	githubDir := os.Getenv("GITHUBDIR")

	if githubDir == "" {
		log.Fatalln("Pls set env: GITHUBDIR")
	}

	fmt.Printf("Go to: %s\n", githubDir)
	err := os.Chdir(githubDir)
	tools.FatalError(err)

	githubDir = path.Join(githubDir, owner)
	exists := tools.Exists(githubDir)
	if !exists {
		err := os.Mkdir(owner, os.ModePerm)
		tools.FatalError(err)
	}

	err = os.Chdir(githubDir)
	tools.FatalError(err)

	exists = tools.Exists(path.Join(githubDir, repo))
	if exists {
		log.Fatalf("The %s/%s exists", owner, repo)
	}

	// Clone

	result, err := tools.ExecCommand("git", "clone", dest)
	tools.FatalError(err)
	fmt.Println(result)

	if user == defaultUser {
		err := os.Chdir(path.Join(githubDir, repo))
		tools.FatalError(err)

		_, _ = tools.ExecCommand("git", "config", "user.email", "\"haoxinst@gmail.com\"")
		_, _ = tools.ExecCommand("git", "config", "user.name", "\"haoxin\"")

		result, err := tools.ExecCommand("git", "config", "--list", "--local")
		tools.FatalError(err)
		fmt.Println(result)
	}
}
