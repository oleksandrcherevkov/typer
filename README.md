# Typer

Small TUI application for typing training over local text files.

## About

To open a file pass its path as the first argument to the application start command.

From root of the project:

```sh
go run ./cil <path>
```

![image](https://github.com/user-attachments/assets/8d5946d9-519a-4152-aed1-f66cfd0ad1a4)

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
