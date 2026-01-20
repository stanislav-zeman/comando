# Comando

A tree-based TUI (Terminal User Interface) application for organizing and executing shell commands.
Navigate through hierarchical command structures and quickly execute commands directly from your terminal.

## Features

- **Tree-based navigation**: Organize commands into folders and subfolders
- **Intuitive interface**: Navigate with arrow keys, execute with enter
- **YAML configuration**: Simple and flexible command organization
- **Direct execution**: Selected commands run in your shell with full environment
- **Breadcrumb navigation**: Always know where you are in the tree
- **Visual indicators**: Clear distinction between folders (📁) and commands (⚡)

## Installation

### From Source

```bash
git install github.com/stanislav-zeman/comando/cmd/comando
```

### Requirements

- Go 1.24.0 or later

## Usage

```bash
# Run with default config
comando

# Run with custom config file
comando /path/to/your/config.yaml
```

### Key Bindings

| Key                             | Action                                  |
| ------------------------------- | --------------------------------------- |
| `↑` / `k`                       | Move cursor up                          |
| `↓` / `j`                       | Move cursor down                        |
| `↵` (Enter)                     | Navigate into folder or execute command |
| `←` / `Backspace` / `h` / `Esc` | Navigate back to parent folder          |
| `q` / `Ctrl+C`                  | Quit without executing                  |

## Configuration

Create a `config/comando.yaml` file (or any location you prefer) with your commands organized in a tree structure:

```yaml
commands:
  - name: Development
    children:
      - name: Start local server
        command: npm run dev
      - name: Run tests
        command: npm test
      - name: Build project
        command: go build

  - name: Servers
    children:
      - name: Production
        children:
          - name: SSH to API Server
            command: ssh api.example.com
          - name: SSH to DB Server
            command: ssh db.example.com
      - name: Staging
        command: ssh staging.example.com

  - name: Docker
    children:
      - name: List containers
        command: docker ps -a
      - name: Clean up
        command: docker system prune -f

  - name: Quick Commands
    children:
      - name: Edit config
        command: vim config/gocut.yaml
      - name: List directory
        command: ls -la
```

### Configuration Structure

Each node in the tree can be either a **folder** or a **command**:

- **Folder**: Has a `name` and `children` array
- **Command**: Has a `name` and `command` string (the shell command to execute)

You can nest folders as deeply as needed to organize your commands logically.

## Use Cases

- **Server Management**: Organize SSH connections by environment, service, or region
- **Development Workflows**: Quick access to build, test, and deployment commands
- **Docker Operations**: Group container management commands
- **Database Tasks**: Organize connection strings and maintenance scripts
- **Custom Scripts**: Create a personal command palette for frequent operations
