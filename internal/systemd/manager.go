package systemd

import (
	"os/exec"
	"strings"
)

type Service struct {
	Name   string
	Status string
}

func GetServices() ([]Service, error) {
	out, _ := exec.Command("systemctl", "list-units", "--type=service", "--no-legend", "--no-pager").Output()
	var services []Service
	for _, line := range strings.Split(string(out), "\n") {
		f := strings.Fields(line)
		if len(f) >= 4 {
			services = append(services, Service{Name: f[0], Status: f[3]})
		}
	}
	return services, nil
}