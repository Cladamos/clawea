# clawea

A terminal based weather forecast application written in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss)

![clawea](screenshots/fullUi.png)

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

- `q` or `ctrl+c`: Quit the application
- `l` or `→`: Next Page
- `h` or `←`: Previous Page

## Configuration

Clawea follows standard OS conventions for storing configuration files. On your first run, it will create automatically

### File Path
* **macOS:** `/Users/username/Library/Application\ Support/clawea/config.conf`
* **Linux:** `~/.config/clawea/config.conf`
* **Windows:** `%AppData%\clawea\config.conf`

### Configuration Settings
You can manually set your coordinates or leave them as `0.0` to enable automatic location detection via your IP address. Also you can change metrics to your preferred one.

```ini
[location]
latitude  = 41.67
longitude = 26.56
country   = Türkiye
region    = Edirne

[metrics]
is_imperial = false
```

## APIs Used

- I used [Open-Meteo](https://open-meteo.com/) for all weather data in the application.
- I used [ip-api](https://ip-api.com/) to find the user's location.
