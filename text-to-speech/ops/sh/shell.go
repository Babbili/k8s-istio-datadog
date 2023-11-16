package shell

import (
	"log"
	"os/exec"
)

func Native(target string) string {
	cmd, err := exec.Command("./ops/sh/script.sh", target).Output()

	if err != nil {
		log.Fatal(err)
	}

	output := string(cmd)
	return output
}

func GetText(target string) string {
	cmd, err := exec.Command("./ops/sh/curl.sh", target).Output()

	if err != nil {
		log.Fatal(err)
	}

	output := string(cmd)
	return output
}
