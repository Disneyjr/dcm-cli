package main

import (
	"flag"
	"fmt"

	"github.com/Disneyjr/dcm/internal/workspace"
	"github.com/Disneyjr/dcm/utils"
	"github.com/Disneyjr/dcm/utils/messages"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		messages.PrintHelp()
		return
	}

	if err := runDcm(args); err != nil {
		fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
		messages.ExitMessage()
	}
}

func runDcm(args []string) error {
	var ws *workspace.Workspace
	// Comandos que não precisam de workspace
	if args[0] != "version" && args[0] != "init" {
		ws = workspace.NewWorkspace()
		if err := workspace.LoadWorkspace(ws); err != nil {
			return err
		}
	}

	switch args[0] {
	case "version":
		handleVersionCommand()
		return nil

	case "init":
		return handleInitCommand()

	case "validate":
		return handleValidateCommand(ws)

	case "up":
		return handleUpCommand(ws, args)

	case "down":
		return handleDownCommand(ws, args)

	case "restart":
		return handleRestartCommand(ws)

	case "logs":
		return handleLogsCommand(ws)

	case "status":
		return handleStatusCommand(ws)

	case "list":
		return handleListCommand(ws)

	case "inspect":
		return handleInspectCommand(ws, args)

	default:
		return fmt.Errorf("comando desconhecido: %s", args[0])
	}
}
