package main // import "go.xitonix.io/sshx"

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

const identityEnvVariable = "SSH_IDENTITY_HOME"

func main() {
	log.SetFlags(0)
	withIdentity, command, args := getCommand()
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if withIdentity {
		err := setIdentitiesHome()
		if err != nil {
			log.Fatal(err)
		}
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to run the '%s' command: %s", command, err))
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func getCommand() (bool, string, []string) {
	var args []string
	command := "ssh"
	hasIdentity := false
	for _, arg := range os.Args[1:] {
		lArg := strings.ToLower(arg)
		if strings.HasPrefix(lArg, "-scp") {
			command = "scp"
			continue
		}
		if lArg == "-i" {
			hasIdentity = true
		}
		args = append(args, arg)
	}
	return hasIdentity, command, args
}

func setIdentitiesHome() error {
	path := os.Getenv(identityEnvVariable)
	trimmed := strings.TrimSpace(path)
	if len(trimmed) == 0 {
		return fmt.Errorf("failed to find the home for identity files. Make sure you set the %s environment variable", identityEnvVariable)
	}
	if strings.HasPrefix(trimmed, "~/") {
		usr, err := user.Current()
		if err != nil {
			return fmt.Errorf("failed to find the current user's profile: %s", err)
		}
		path = filepath.Join(usr.HomeDir, path[2:])
	}

	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("failed to load %s: %s", path, err)
	}
	return nil
}
