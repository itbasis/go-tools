package docker

import "os"

func IsRunningInDockerContainer() bool {
	if len(os.Getenv("KUBERNETES_PORT")) > 0 {
		return true
	}

	if _, err := os.Stat("/.dockerenv"); err != nil {
		return false
	}

	return true
}
