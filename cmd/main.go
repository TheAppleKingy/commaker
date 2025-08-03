package main

import (
	"commiter/internal/application"
	"commiter/internal/commands"
	"flag"
	"log/slog"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		slog.Error("Expected 'commit' got ''")
		os.Exit(1)
	}
	if args[0] != "commit" {
		slog.Error("Incorrect command. Enter `commaker commit`")
		os.Exit(1)
	}
	push := flag.Bool(PUSH, false, "Add this flag to push changes to repote repository")
	lr := flag.String(LOCAL_REPOSITORY_DIRECTORY, commands.GetWorkingDir(), "May be provided only when push flag provided. Mean local repository directory what have to push. If not app will be try push from current directory")
	b := flag.String(BRANCH_NAME, "", "Branch that have to be pushed")
	flag.CommandLine.Parse(args[1:])
	commands.CheckIsRepo(*lr)
	changes := commands.GetGitDiff()
	message := application.GetCommitMessage(changes)
	commands.Commit(message)
	if *push {
		commands.PushToRemote(*b)
	}
}
