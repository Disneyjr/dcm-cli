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
		messages.ExitMessage()
		return
	}

	// Comandos que não precisam de workspace e não devem mostrar ExitMessage
	if args[0] == "version" {
		handleVersionCommand(nil, args)
		return
	}

	if args[0] == "init" {
		handleInitCommand(nil, args)
		return
	}

	// Para todos os outros comandos, mostrar ExitMessage ao final
	defer messages.ExitMessage()

	ws := workspace.NewWorkspace()
	if err := workspace.LoadWorkspace(ws); err != nil {
		fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
		return
	}

	switch args[0] {
	case "up":
		handleUpCommand(ws, args)

	case "down":
		handleDownCommand(ws, args)

	case "restart":
		handleRestartCommand(ws, args)

	case "logs":
		handleLogsCommand(ws, args)

	case "status":
		handleStatusCommand(ws, args)

	case "list":
		handleListCommand(ws, args)

	case "inspect":
		handleInspectCommand(ws, args)

	case "validate":
		handleValidateCommand(ws, args)

	default:
		fmt.Printf("%s Comando desconhecido: %s\n", utils.Colorize("yellow", "⚠️"), args[0])
		return
	}
}
