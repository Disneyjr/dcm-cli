package messages

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Disneyjr/dcm/utils"
)

var Version = "dev"

func ExitMessage() {
	fmt.Println("\nPressione ENTER para sair...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
func InstallSuccessful() {
	fmt.Printf("%s DCM instalado com sucesso!\n", utils.Colorize("green", "✅"))
	fmt.Printf("%s Você pode usar 'dcm' em qualquer terminal/pasta.\n\n", utils.Colorize("green", "✨"))
	fmt.Printf("Exemplo:\n")
	fmt.Printf("  dcm list\n")
	fmt.Printf("  dcm up dev\n")
	fmt.Printf("  dcm version\n\n")
}
func PrintHelp() {
	fmt.Printf("%s DCM - Docker Compose Manager\n\n", utils.Colorize("cyan", "📌"))
	fmt.Printf("Versão: %s\n\n", Version)
	fmt.Println("Uso:")
	fmt.Println("  dcm up <grupo> [--build] [--dry-run] - Inicia grupo")
	fmt.Println("  dcm down                      - Para todos os serviços")
	fmt.Println("  dcm restart                   - Reinicia todos")
	fmt.Println("  dcm logs                      - Mostra logs")
	fmt.Println("  dcm status                    - Status dos serviços")
	fmt.Println("  dcm list                      - Lista projetos e grupos")
	fmt.Println("  dcm inspect <grupo>           - Detalha composição de um grupo")
	fmt.Println("  dcm validate                  - Valida o arquivo workspace.json")
	fmt.Println("  dcm init                      - Cria configuração inicial")
	fmt.Println("  dcm version                   - Mostra versão")
}

func VersionMessage() {
	fmt.Printf("dcm v%s\n", Version)
}
