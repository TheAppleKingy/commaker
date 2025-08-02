package main

import (
	"commiter/internal/commands"
	"flag"
	"fmt"
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
	// push := flag.Bool(PUSH, false, "Add this flag to push changes to repote repository")
	lr := flag.String(LOCAL_REPOSITORY_DIRECTORY, commands.GetWorkingDir(), "May be provided only when push flag provided. Mean local repository directory what have to push. If not app will be try push from current directory")
	// lb := flag.String(LOCAL_BRANCH_NAME, "main", "Local branch that have to be pushed")
	// rb := flag.String(REMOTE_BRANCH_NAME, "main", "Remote branch for pushing")
	flag.CommandLine.Parse(args[1:])
	commands.CheckIsRepo(*lr)
	changes := commands.GetGitDiff()
	fmt.Print(changes)
}
