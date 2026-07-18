# Servio 🚀
**A powerful, elegant, and interactive TUI to manage systemd services.**

Servio is a modern terminal-based tool designed for sysadmins and developers to manage systemd services with ease. Forget about typing long `systemctl` commands; use Servio to filter, inspect, and control your services in a beautiful UI.

## ✨ Key Features
- **Interactive UI:** Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), providing a smooth terminal experience.
- **Fuzzy Search:** Filter through hundreds of services in real-time.
- **Log Viewer:** View real-time service logs directly within the app (scrollable!).
- **Quick Actions:** Restart, Start, Stop, Enable, and Disable services with a single keypress.
- **Cross-Platform:** Binaries available for amd64 and arm64.

## 📥 Installation

### Debian/Ubuntu (.deb)
```bash
wget https://github.com/aminmadaniofficial/servio/releases/download/v1.0.0/servio_1.0.0_linux_amd64.deb
sudo dpkg -i servio_1.0.0_linux_amd64.deb
```

### From Source
```bash
go install github.com/aminmadaniofficial/servio/cmd/servio@latest
```

## ⌨️ Shortcuts
| Key | Action |
| :--- | :--- |
| `/` | Search/Filter services |
| `r` | Restart service |
| `s` | Stop service |
| `S` | Start service |
| `l` | View Logs |
| `q` | Quit |

## 🛠 Built With
- [Go](https://go.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)

---
*Built with ❤️ by Amin Madani*