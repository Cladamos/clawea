# clawea

A terminal based weather forecast application written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss)

## Installation

```bash
go install github.com/cladamos/clawea@latest
```

If the command isn't recognized after installation, ensure your Go bin folder is in your PATH

## Usage

Simply run the command to see the forecast for your current location:

```bash
clawea
```

### Controls

- `q` or `ctrl+c`: Quit the application.

## APIs Used

- [Open-Meteo](https://open-meteo.com/)
- [ip-api](https://ip-api.com/)
