package main

import (
	"flag"
	"fmt"

	"github.com/Disneyjr/dcm/internal/commands"
	"github.com/Disneyjr/dcm/internal/workspace"
	"github.com/Disneyjr/dcm/utils"
	"github.com/Disneyjr/dcm/utils/messages"
)

func main() {
	flag.Parse()
	args := flag.Args()
	defer messages.ExitMessage()
	if len(args) == 0 {
		messages.PrintHelp()
		return
	}
	ws := workspace.NewWorkspace()
	if err := workspace.LoadWorkspace(ws); err != nil {
		if args[0] != "version" {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}
	}

	switch args[0] {
	case "up":
		if len(args) < 2 {
			fmt.Println(utils.Colorize("red", "❌ Especifique um projeto ou grupo"))
			return
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

		if groupErr := commands.UpGroup(ws, projectOrGroup, extraArgs...); groupErr == nil {
			return
		}

	case "down":
		if err := commands.DownAll(ws); err != nil {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}

	case "restart":
		if err := commands.RestartAll(ws); err != nil {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}

	case "logs":
		if err := commands.LogsAll(ws); err != nil {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}

	case "status":
		if err := commands.StatusAll(ws); err != nil {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}

	case "list":
		commands.ListAll(ws)

	case "inspect":
		if len(args) < 2 {
			fmt.Println(utils.Colorize("red", "❌ Especifique um grupo para inspecionar"))
			return
		}
		commands.InspectGroup(ws, args[1])

	case "validate":
		commands.ValidateWorkspace(ws)

	case "init":
		if err := commands.InitWorkspace(); err != nil {
			fmt.Printf("%s %v\n", utils.Colorize("red", "❌"), err)
			return
		}

	case "version":
		messages.VersionMessage()

	default:
		fmt.Printf("%s Comando desconhecido: %s\n", utils.Colorize("yellow", "⚠️"), args[0])
		return
	}
}
