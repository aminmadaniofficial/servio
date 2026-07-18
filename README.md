<p align="center">
  <h1 align="center">🚀 Servio</h1>
  <p align="center">
    <strong>A powerful, elegant, and interactive TUI to manage systemd services.</strong>
  </p>
  <p align="center">
    <img src="https://img.shields.io/github/v/release/aminmadaniofficial/servio?style=flat-square&color=7D56F4" alt="Release">
    <img src="https://img.shields.io/github/license/aminmadaniofficial/servio?style=flat-square&color=00BFFF" alt="License">
    <img src="https://img.shields.io/github/go-mod/go-version/aminmadaniofficial/servio?style=flat-square" alt="Go Version">
  </p>
</p>

---

Servio is a modern terminal-based tool designed for sysadmins and developers to manage systemd services with ease. Forget about typing long and tedious `systemctl` commands; use Servio to filter, inspect, and control your services in a beautiful, highly responsive UI.

<p align="center">
  <!-- 🟢 این همون جاییه که عکس یا گیف شما قرار میگیره -->
  <img src="assets/demo.gif" alt="Servio Demo" width="800">
</p>

## ✨ Key Features

- **Interactive UI:** Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), providing a buttery-smooth terminal experience.
- **Fuzzy Search:** Press `/` to instantly filter through hundreds of services in real-time.
- **Log Viewer:** View real-time service logs directly within the app (scrollable up to 100 lines!).
- **Quick Actions:** Restart, Start, Stop, Enable, and Disable services with a single keypress.
- **Cross-Platform:** Binaries available for `amd64` and `arm64` Linux architectures.

## 📥 Installation

### Debian / Ubuntu (.deb)
```bash
wget https://github.com/aminmadaniofficial/servio/releases/download/v1.0.0/servio_1.0.0_linux_amd64.deb
sudo dpkg -i servio_1.0.0_linux_amd64.deb
```

### RHEL / CentOS / Fedora (.rpm)
```bash
wget https://github.com/aminmadaniofficial/servio/releases/download/v1.0.0/servio_1.0.0_linux_amd64.rpm
sudo rpm -i servio_1.0.0_linux_amd64.rpm
```

### Build from Source
```bash
go install github.com/aminmadaniofficial/servio/cmd/servio@latest
```

## 🚀 Usage

Simply type `servio` in your terminal. Note: Action commands (like restart/stop) will prompt for `sudo` privileges.

```bash
servio
```

## ⌨️ Shortcuts

| Key | Action | Description |
| :--- | :--- | :--- |
| `/` | **Search** | Open the search bar to filter services |
| `r` | **Restart** | Restart the currently selected service |
| `s` / `S` | **Stop / Start**| Stop or Start the selected service |
| `e` / `d` | **Enable / Disable**| Enable or Disable service on boot |
| `l` | **Logs** | Open the viewport to see service logs |
| `j` / `k` | **Navigate** | Move cursor Down / Up |
| `q` | **Quit** | Exit the application |

## 🛠 Built With

- [Go](https://go.dev/) - The programming language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The powerful TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for the terminal

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
<p align="center">Built with ❤️ by <a href="https://github.com/aminmadaniofficial">Amin Madani</a></p>
