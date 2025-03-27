package common

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// StepWaitForRemoteFile waits for the presence of a file on the remote VM.
// It periodically checks via SSH whether the file exists.
// If the SSH tunnel is down (e.g. after a reboot), it attempts to reestablish it.
type StepWaitForRemoteFile struct {
	Path    string        // The path to the signal file (e.g., installation_done)
	Timeout time.Duration // Maximum time to wait for the file
}

// reestablishSSHTunnel is a helper function to reestablish the SSH tunnel.
// In a complete implementation, this would re-run the logic from the SSH port-forward step.
func reestablishSSHTunnel(state multistep.StateBag) error {
	ui := state.Get("ui").(packer.Ui)
	ui.Message("Attempting to reestablish SSH tunnel...")

	// Example: You might call the StepForwardPortOverSSH logic here or reinitialize the tunnel.
	// For now, we simulate the process by waiting for a few seconds.
	time.Sleep(5 * time.Second)

	// After the wait, the tunnel should be reestablished (update state accordingly).
	// In a real implementation, ensure that the state key "local_ssh_port" is updated.

	return nil
}

// Run executes the step by periodically checking if the remote file exists.
// It also verifies that the SSH tunnel is up before each check.
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
			// Check if the SSH tunnel is active by attempting a TCP connection.
			localAddress, err := SSHLocalAddress(state)
			if err != nil {
				ui.Message("Local SSH address not available, waiting for tunnel setup...")
				continue
			}
			conn, err := net.DialTimeout("tcp", localAddress, 5*time.Second)
			if err != nil {
				ui.Message("SSH tunnel appears to be down, attempting to reestablish...")
				// Try to reestablish the SSH tunnel.
				if err := reestablishSSHTunnel(state); err != nil {
					ui.Message("Failed to reestablish SSH tunnel: " + err.Error())
					continue
				}
				// Wait a bit for the new tunnel to stabilize.
				time.Sleep(3 * time.Second)
				// Attempt to get the local address and connect again.
				localAddress, err = SSHLocalAddress(state)
				if err != nil {
					continue
				}
				conn, err = net.DialTimeout("tcp", localAddress, 5*time.Second)
				if err != nil {
					ui.Message("Tunnel still not available after reestablishment, waiting...")
					continue
				}
			}
			if conn != nil {
				conn.Close()
			}

			// Try executing the SSH command to check for the existence of the file.
			var out string
			var execErr error
			retryCount := 3
			for i := 0; i < retryCount; i++ {
				// Execute "test -f <file>" on the guest via SSH.
				out, execErr = ExecuteGuestSSHCmd(state, fmt.Sprintf("test -f '%s'", s.Path))
				if execErr == nil {
					break
				}
				ui.Message(fmt.Sprintf("SSH attempt %d failed: %s", i+1, execErr.Error()))
				time.Sleep(2 * time.Second)
			}
			if execErr == nil {
				ui.Message("Signal file found. Continuing...")
				return multistep.ActionContinue
			} else {
				ui.Message("Waiting... file not found yet.")
			}
		}
	}
}

// Cleanup is a no-op for this step.
func (s *StepWaitForRemoteFile) Cleanup(state multistep.StateBag) {}
