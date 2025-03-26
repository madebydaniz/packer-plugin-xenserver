package common

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"golang.org/x/crypto/ssh"
)

type StepWaitForRemoteFile struct {
	Path    string
	Timeout time.Duration
}

func (s *StepWaitForRemoteFile) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	if s.Path == "" {
		s.Path = "/home/xcoorp/installation_done" // fallback default path for testing
	}

	if s.Timeout == 0 {
		s.Timeout = 15 * time.Minute // fallback timeout for testing
	}

	ui.Say(fmt.Sprintf("=== Waiting for signal file: '%s' ===", s.Path))
	timeout := time.After(s.Timeout)
	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			ui.Say("Context cancelled while waiting for signal file.")
			return multistep.ActionHalt

		case <-timeout:
			err := fmt.Errorf("Timeout reached while waiting for file '%s'", s.Path)
			state.Put("error", err)
			return multistep.ActionHalt

		case <-tick.C:
			cmd := fmt.Sprintf("test -f '%s'", s.Path)
			_, err := ExecuteGuestSSHCmd(state, cmd)

			if err == nil {
				ui.Message("Signal file found. Continuing...")
				return multistep.ActionContinue
			}

			if isSshDisconnected(err) {
				ui.Message("SSH disconnected. Trying to reconnect...")
				if reconnectErr := reconnectSSHInternal(ctx, state); reconnectErr != nil {
					ui.Error("Reconnect failed: " + reconnectErr.Error())
					state.Put("error", reconnectErr)
					return multistep.ActionHalt
				}
				continue
			}

			ui.Message("Waiting... file not found yet.")
		}
	}
}

func (s *StepWaitForRemoteFile) Cleanup(state multistep.StateBag) {}

func isSshDisconnected(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	return strings.Contains(s, "connection reset") ||
		strings.Contains(s, "broken pipe") ||
		strings.Contains(s, "EOF") ||
		strings.Contains(s, "connection refused") ||
		strings.Contains(s, "no route to host")
}

func reconnectSSHInternal(ctx context.Context, state multistep.StateBag) error {
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(map[string]interface{})
	sshUser := config["ssh_username"].(string)
	sshPass := config["ssh_password"].(string)
	host := state.Get("instance_ssh_address").(string)
	port := 22

	deadline := time.Now().Add(10 * time.Minute)

	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout reconnecting SSH")
		}

		addr := fmt.Sprintf("%s:%d", host, port)
		clientConfig := &ssh.ClientConfig{
			User:            sshUser,
			Auth:            []ssh.AuthMethod{ssh.Password(sshPass)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         10 * time.Second,
		}

		conn, err := ssh.Dial("tcp", addr, clientConfig)
		if err == nil {
			defer conn.Close()
			ui.Message("SSH reconnect successful")
			return nil
		}

		ui.Message("Still waiting for SSH...")
		time.Sleep(5 * time.Second)
	}
}
