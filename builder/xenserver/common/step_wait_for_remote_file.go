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

	cmd := fmt.Sprintf("test -f '%s'", s.Path)
	_, err := RetryableExecuteGuestSSHCmd(state, cmd, s.Timeout)

	if err != nil {
		ui.Error(fmt.Sprintf("Error waiting for remote file: %s", err))
		state.Put("error", err)
		return multistep.ActionHalt
	}

	ui.Message("Signal file found. Continuing...")
	return multistep.ActionContinue
}

func (s *StepWaitForRemoteFile) Cleanup(state multistep.StateBag) {}
