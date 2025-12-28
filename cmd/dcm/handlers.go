package main

import (
	"fmt"
	"strings"

	"github.com/Disneyjr/dcm/internal/commands"
	"github.com/Disneyjr/dcm/internal/workspace"
	"github.com/Disneyjr/dcm/utils/messages"
)

func handleUpCommand(ws *workspace.Workspace, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("especifique um projeto ou grupo")
	}

	projectOrGroup := args[1]
	extraArgs := []string{}
	for i := 2; i < len(args); i++ {
		if args[i] == "--build" {
			extraArgs = append(extraArgs, "--build")
		}
		if args[i] == "--dry-run" {
			commands.DryRun = true
		}
	}

	return commands.UpGroup(ws, projectOrGroup, extraArgs...)
}

func handleDownCommand(ws *workspace.Workspace, args []string) error {
	removeVolumes := false
	groupName := ""

	// Parse arguments
	for i := 1; i < len(args); i++ {
		if args[i] == "-v" {
			removeVolumes = true
		} else if !strings.HasPrefix(args[i], "-") {
			groupName = args[i]
		}
	}

	// If group is specified, use DownGroup
	if groupName != "" {
		return commands.DownGroup(ws, groupName, removeVolumes)
	}

	// Otherwise, use DownAll
	return commands.DownAll(ws, removeVolumes)
}

func handleRestartCommand(ws *workspace.Workspace) error {
	return commands.RestartAll(ws)
}

func handleLogsCommand(ws *workspace.Workspace) error {
	return commands.LogsAll(ws)
}

func handleStatusCommand(ws *workspace.Workspace) error {
	return commands.StatusAll(ws)
}

func handleListCommand(ws *workspace.Workspace) error {
	commands.ListAll(ws)
	return nil
}

func handleInspectCommand(ws *workspace.Workspace, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("especifique um grupo para inspecionar")
	}
	commands.InspectGroup(ws, args[1])
	return nil
}

func handleValidateCommand(ws *workspace.Workspace) error {
	commands.ValidateWorkspace(ws)
	return nil
}

func handleInitCommand() error {
	return commands.InitWorkspace()
}

func handleVersionCommand() {
	messages.VersionMessage()
}
