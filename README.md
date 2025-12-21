# CLIFolio

A terminal-based portfolio accessible via SSH. Built with Go and Bubbletea.

## Overview

clifolio is an interactive terminal user interface that serves as a personal portfolio. Users can SSH into your server and navigate through your projects, skills, experience, and contact information using keyboard controls.

## Features

- Interactive terminal UI with smooth navigation
- GitHub integration for live project data
- Real-time statistics dashboard
- Multiple theme support (Hacker, Dracula, Solarized)
- Matrix rain easter egg
- SSH server for remote access
- Markdown rendering for project READMEs
- Responsive layout with clean design

## Technology Stack

**Language:** Go 1.24.4

**Core Libraries:**
- Bubbletea - Terminal UI framework
- Lipgloss - Styling and layout
- Bubbles - Reusable TUI components
- Glamour - Markdown rendering

**Integrations:**
- GitHub API (go-github)
- OAuth2 authentication
- Wish - SSH server

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/clifolio.git
cd clifolio
```

Install dependencies:

```bash
go mod download
```

Create a `.env` file for GitHub integration:

```bash
GITHUB_TOKEN=your_github_personal_access_token
```

Build the application:

```bash
go build -o clifolio .
```

## Usage

### Local Mode

Run the TUI locally in your terminal:

```bash
./clifolio
```

With a specific theme:

```bash
./clifolio --theme hacker
```

### SSH Mode

Start the SSH server:

```bash
./clifolio --ssh-mode
```

Users can then connect via:

```bash
ssh username@your-server-address -p 23234
```


## Navigation Controls

- `↑/↓` or `j/k` - Navigate lists
- `←/→` or `h/l` - Switch tabs
- `Enter` - Select item
- `/` - Open menu (from any screen)
- `m` - Activate Matrix easter egg
- `ESC` - Go back
- `q` or `Ctrl+C` - Quit

## Configuration

Edit the following files to customize your portfolio:

- `internal/ui/skills.go` - Update your skills and technologies
- `internal/ui/experience.go` - Add your work experience
- `internal/ui/contact.go` - Update contact information
- `internal/ui/projects.go` - Change GitHub username
- `internal/ui/stats.go` - Change GitHub username for stats
- `assets/intro.txt` - Customize intro text
- `assets/ascii.txt` - Add custom ASCII art

## Development

Run locally with auto-reload during development:

```bash
go run main.go
```

Run tests:

```bash
go test ./...
```

Format code:

```bash
go fmt ./...
```

## Deployment

### VPS Deployment

1. Build for your target platform:

```bash
GOOS=linux GOARCH=amd64 go build -o clifolio
```

2. Transfer to your VPS:

```bash
scp clifolio user@your-vps:/path/to/app/
```

3. Set up as a systemd service:

```ini
[Unit]
Description=CLIFolio SSH Portfolio
After=network.target

[Service]
Type=simple
User=clifolio
WorkingDirectory=/path/to/app
ExecStart=/path/to/app/clifolio --ssh-mode
Restart=always

[Install]
WantedBy=multi-user.target
```

4. Enable and start:

```bash
sudo systemctl enable clifolio
sudo systemctl start clifolio
```

### Docker Deployment

Build the Docker image:

```bash
docker build -t clifolio .
```

Run the container:

```bash
docker run -d -p 23234:23234 --name clifolio clifolio
```

## Security Considerations

When deploying the SSH server:

- Use SSH key authentication
- Disable password authentication
- Run as non-root user
- Configure firewall rules
- Use `ForceCommand` in SSH config to prevent shell access
- Implement rate limiting

## License

MIT

## Author

Janpol Hidalgo
