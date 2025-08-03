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
		fmt.Println(`usage: commaker commit [-p] [-b <branch>]

Commaker is a simple Git commit helper that generates concise commit messages based on git diff.
Supports optional push to the remote repository.

Commands:
	commit        Create a commit with an AI-generated message based on staged changes.

Options:
	-p            Push the commit to the remote repository after committing.
	-b <branch>   Specify the branch. If not provided, the current branch is used.

Examples:
	commaker commit
		Create a commit using AI-generated message from current staged diff.

	commaker commit -p
		Commit and push to the current branch on origin.

	commaker commit -p -b feature-branch
		Commit and push to origin/feature-branch.`)
		os.Exit(0)
	}

	push := flag.Bool(PUSH, false, "Add this flag to push changes to repote repository")
	b := flag.String(BRANCH_NAME, "", "Branch that have to be pushed")
	flag.CommandLine.Parse(args[1:])
	commands.CheckIsRepo()
	if *b != "" {
		commands.SwitchToBranch(*b)
	}
	changes := commands.GetGitDiff()
	message := application.GetCommitMessage(changes)
	commands.Commit(message)
	if *push {
		commands.PushToRemote()
	}
}
