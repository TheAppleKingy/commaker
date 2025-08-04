package main

import (
	"commiter/internal/application"
	"commiter/internal/commands"
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" || args[0] != "commit" {
		fmt.Println(`usage: commaker commit [-p]

Commaker is a simple Git commit helper that generates concise commit messages based on git diff.
Supports optional push to the remote repository.

Commands:
	commit        Create a commit with an AI-generated message based on staged changes.

Options:
	-p            Push the commit to the remote repository after committing.

Examples:
	commaker commit
		Create a commit using AI-generated message from current staged diff.

	commaker commit -p
		Commit and push to the current branch on origin.`)
		os.Exit(0)
	}

	p := flag.Bool("p", false, "Add this flag to push changes to remote repository")
	flag.CommandLine.Parse(args[1:])
	commands.CheckIsRepo()
	changes := commands.GetGitDiff()
	message := application.GetCommitMessage(changes)
	commands.Commit(message)
	if *p {
		commands.PushToRemote()
	}
}
