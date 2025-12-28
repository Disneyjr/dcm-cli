package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Disneyjr/dcm/utils"
	"github.com/Disneyjr/dcm/utils/messages"
)

func main() {
	fmt.Printf("\n%s DCM - Desinstalador\n\n", utils.Colorize("cyan", "🗑️"))
	defer messages.ExitMessage()

	if !utils.IsAdmin() {
		fmt.Println("🗑️  DCM - Desinstalador")
		fmt.Println("❌ ERRO: Necessário privilégios de Administrador")
		fmt.Println("💡 Clique direito no uninstall.exe > 'Executar como administrador'")
		return
	}

	var targetPath string
	var binaryName string

	switch runtime.GOOS {
	case "linux", "darwin":
		targetPath = "/usr/local/bin/dcm"
		binaryName = "dcm"
	case "windows":
		targetPath = filepath.Join(os.Getenv("WINDIR"), "System32", "dcm.exe")
		binaryName = "dcm.exe"
	default:
		fmt.Printf("%s SO não suportado: %s\n", utils.Colorize("red", "❌"), runtime.GOOS)
		return
	}

	// Verificar se o DCM está instalado
	_, err := os.Stat(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s DCM não está instalado\n", utils.Colorize("yellow", "⚠️"))
			fmt.Printf("Caminho verificado: %s\n", targetPath)
			return
		}
		fmt.Printf("%s Erro ao verificar instalação: %v\n", utils.Colorize("red", "❌"), err)
		return
	}

	fmt.Printf("%s DCM encontrado em: %s\n", utils.Colorize("green", "✅"), targetPath)
	fmt.Printf("%s Removendo %s...\n", utils.Colorize("blue", "🔧"), binaryName)

	// Remover o binário
	if err := os.Remove(targetPath); err != nil {
		fmt.Printf("%s Erro ao remover: %v\n", utils.Colorize("red", "❌"), err)
		fmt.Printf("\n%s Tente executar manualmente:\n", utils.Colorize("yellow", "💡"))
		if runtime.GOOS == "windows" {
			fmt.Printf("  del \"%s\"\n\n", targetPath)
		} else {
			fmt.Printf("  sudo rm %s\n\n", targetPath)
		}
		return
	}

	fmt.Printf("\n%s DCM desinstalado com sucesso!\n", utils.Colorize("green", "✅"))
	fmt.Printf("%s Obrigado por usar o DCM!\n\n", utils.Colorize("cyan", "👋"))
}
