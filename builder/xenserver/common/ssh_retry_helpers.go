package common

import (
	"fmt"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

func RetryableExecuteGuestSSHCmd(state multistep.StateBag, cmd string, timeout time.Duration) (stdout string, err error) {
	ui := state.Get("ui").(packer.Ui)
	start := time.Now()

	for {
		if time.Since(start) > timeout {
			return "", fmt.Errorf("SSH retry timeout exceeded while waiting to execute: %s", cmd)
		}

		stdout, err = ExecuteGuestSSHCmd(state, cmd)
		if err == nil {
			return stdout, nil
		}

		ui.Message(fmt.Sprintf("SSH not ready or command failed (%s), retrying in 5s...", err.Error()))
		time.Sleep(5 * time.Second)
	}
}