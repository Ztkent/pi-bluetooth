package pitooth

import (
	"fmt"
	"os"
	"os/exec"
)

/*
	OBEX is a protocol for transferring files between devices over Bluetooth.
	To use OBEX, we need to start the obexd service on the Raspberry Pi.
	Then we can use [UUID: OBEX Object Push] or [UUID: OBEX File Transfer] to communicate.
*/

// Raspberry Pi as OBEX Server: Configure obexd to accept incoming connections and receive files.
// Sending the files must be handled by the client. Any files received will be saved in the specified directory.
func (btm *bluetoothManager) ControlOBEXServer(start bool, outputPath string) error {
	// Check the current status of obexd
	pgrepCmd := exec.Command("pgrep", "obexd")
	pidBytes, err := pgrepCmd.Output()
	isActive := err == nil

	// Decide whether to start or stop based on the desired state and current status
	if start && isActive {
		btm.l.Debugln("obexd is already running.")
		return nil
	} else if !start && !isActive {
		btm.l.Debugln("obexd is already stopped.")
		return nil
	}

	// Check if the output folder exists, if it doesnt create it
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	var cmd *exec.Cmd
	if start {
		// Command to start the obexd service
		btm.l.Debugln("Starting obexd service... ")
		cmd = exec.Command("obexd", "-a", "-r", outputPath)
	} else {
		// Command to stop the obexd service
		btm.l.Debugln("Stopping obexd service...")
		cmd = exec.Command("killall", "obexd")
	}

	// Execute the command
	if err := cmd.Start(); err != nil {
		if start {
			return fmt.Errorf("failed to start obexd: %v", err)
		} else {
			return fmt.Errorf("failed to stop obexd: %v", err)
		}
	} else {
		if start {
			btm.l.Infof("obexd [PID: %d] started successfully\n", cmd.Process.Pid)
		} else {
			btm.l.Infof("obexd [PID: %s] stopped successfully\n", string(pidBytes))
		}
	}
	return nil
}
