package common

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepWaitForRemoteFile struct {
	Path    string
	Timeout time.Duration
}

func (s *StepWaitForRemoteFile) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	if s.Path == "" {
		ui.Say("No installation_done_file specified, skipping wait step.")
		return multistep.ActionContinue
	}

	ui.Say(fmt.Sprintf("=== Waiting for signal file: '%s' ===", s.Path))

	timeout := time.After(s.Timeout)
	tick := time.Tick(5 * time.Second)

	for {
		select {
		case <-ctx.Done():
			ui.Say("Context cancelled while waiting for signal file.")
			return multistep.ActionHalt

		case <-timeout:
			err := fmt.Errorf("Timeout reached while waiting for file '%s'", s.Path)
			state.Put("error", err)
			return multistep.ActionHalt

		case <-tick:
			cmd := fmt.Sprintf("test -f '%s'", s.Path)
			_, err := ExecuteGuestSSHCmd(state, cmd)
			if err == nil {
				ui.Message("Signal file found. Continuing...")
				return multistep.ActionContinue
			} else {
				ui.Message("Waiting... file not found yet.")
			}
		}
	}
}

func (s *StepWaitForRemoteFile) Cleanup(state multistep.StateBag) {}
