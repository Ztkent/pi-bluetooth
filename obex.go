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
func (btm *bluetoothManager) ControlOBEXServer(start bool, outputPath string) error {
	pgrepCmd := exec.Command("pgrep", "obexd")
	pidBytes, err := pgrepCmd.Output()
	isActive := err == nil
	if start && isActive {
		btm.l.Debugln("obexd is already running.")
		return nil
	} else if !start && !isActive {
		btm.l.Debugln("obexd is already stopped.")
		return nil
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		if err := os.MkdirAll(outputPath, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	var cmd *exec.Cmd
	if start {
		btm.l.Debugln("Starting obexd service... ")
		cmd = exec.Command("obexd", "-a", "-r", outputPath)
	} else {
		btm.l.Debugln("Stopping obexd service...")
		cmd = exec.Command("killall", "obexd")
	}

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
