package utils

import (
	"fmt"
	"os/exec"
)

func GenerateCertificate(domain string) error {
	cmd := exec.Command("sudo", "mkcert", domain)
	cmd.Dir = GetAppDataPath("./data/certs")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error generating certificate for %s: %v\nOutput: %s", domain, err, output)
		return err
	}
	return nil
}
