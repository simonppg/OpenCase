package hardware

import (
	"net"
	"os"
	"path/filepath"
	"strings"
)

// CollectNetwork discovers network interfaces.
func CollectNetwork() ([]NetworkInfo, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []NetworkInfo

	for _, iface := range interfaces {

		state := "DOWN"

		statePath := filepath.Join(
			"/sys/class/net",
			iface.Name,
			"operstate",
		)

		data, err := os.ReadFile(statePath)
		if err == nil {
			state = strings.TrimSpace(string(data))
		}

		result = append(result, NetworkInfo{
			Name:  iface.Name,
			State: strings.ToUpper(state),
		})
	}

	return result, nil
}
