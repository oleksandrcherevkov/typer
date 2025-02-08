# Typer

Small TUI application for typing training over local text files.

## About

To open a file pass its path as the first argument to the application start command.

From root of the project:

```sh
go run ./cil <path>
```

![image](https://github.com/user-attachments/assets/b749fc9d-c9b4-4c3c-bff6-0bbd1cb6a05a)

### Keybindings

| Keys | Action |
| ---- | ------ |
| ctrl + c | exit |
| down arrow | next line |

## Architecture

Contains one TUI component now. Rendering UI elements in nested way:
Character > Line > Box > Program.

Libraries used:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - for building TUI by blocks.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - for styling application: drawing box and coloring letters.
