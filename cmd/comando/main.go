package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stanislav-zeman/comando/internal/comando"
	"github.com/stanislav-zeman/comando/internal/config"
	"github.com/stanislav-zeman/comando/internal/navigation"
)

const defaultConfigPath = "config/comando.yaml"

var (
	ErrNoCommands      = errors.New("config contains no commands")
	ErrUnexpectedModel = errors.New("unexpected model")
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running: %v\n", err)
		fmt.Fprintln(os.Stderr, "Usage: gocut [config-file]")
		os.Exit(1)
	}
}

func run() error {
	configPath := defaultConfigPath
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if !filepath.IsAbs(configPath) {
		configPath = filepath.Join(".", configPath)
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed loading config from %s: %v", configPath, err)
	}

	rootNodes := navigation.ParseTree(config.Commands)
	if len(rootNodes) == 0 {
		return ErrNoCommands
	}

	model := comando.NewModel(rootNodes)
	program := tea.NewProgram(model, tea.WithAltScreen())
	finalModel, err := program.Run()
	if err != nil {
		return err
	}

	m, ok := finalModel.(comando.Model)
	if !ok {
		return ErrUnexpectedModel
	}

	selectedCommand := m.GetSelectedCommand()
	if selectedCommand == "" {
		return nil
	}

	args := strings.Fields(selectedCommand)
	if len(args) == 0 {
		return nil
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed running command: %w", err)
	}

	return nil
}
