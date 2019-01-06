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

	command, args, err := getCommand()
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("sshx: failed to access the current working directory: %s", err)
	}
	cmd.Dir = wd
	err = cmd.Run()
	if err != nil {
		log.Fatalf("sshx: failed to run the '%s' command: %s", command, err)
	}
}

func getCommand() (string, []string, error) {
	var args []string
	command := "ssh"

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		lArg := strings.ToLower(arg)
		if strings.HasPrefix(lArg, "-scp") {
			command = "scp"
			continue
		}
		args = append(args, arg)
		if lArg != "-i" {
			continue
		}
		identityFile := os.Args[i+1]
		if len(identityFile) == 0 {
			continue
		}

		id, err := prefixIdentity(identityFile)
		if err != nil {
			return "", nil, err
		}
		args = append(args, id)
		i++
	}
	return command, args, nil
}

func prefixIdentity(idArg string) (string, error) {
	var err error
	idArg, err = replaceHomeDir(idArg)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(idArg); err == nil {
		return idArg, nil
	}

	identityDir := os.Getenv(identityEnvVariable)
	trimmed := strings.TrimSpace(identityDir)
	if len(trimmed) == 0 {
		return "", fmt.Errorf("sshx: failed to find the home for identity files. Make sure you set the %s environment variable", identityEnvVariable)
	}

	identityDir, err = replaceHomeDir(trimmed)
	if err != nil {
		return "", err
	}

	return filepath.Join(identityDir, idArg), nil
}

func replaceHomeDir(dir string) (string, error) {
	if !strings.HasPrefix(dir, "~/") {
		return dir, nil
	}
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("sshx: failed to find the current user's profile: %s", err)
	}
	return filepath.Join(usr.HomeDir, dir[2:]), nil
}
