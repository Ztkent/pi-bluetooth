package pitooth

import (
	"testing"
	"time"
)

func Test_NewBluetoothManager(t *testing.T) {
	btm, err := NewBluetoothManager("SunlightMeter")
	if err != nil || btm == nil {
		t.Fatalf("Failed to create Bluetooth Manager: %v", err)
	}
	defer btm.Stop()
}

func Test_AcceptConnections(t *testing.T) {
	btm, err := NewBluetoothManager("SunlightMeter")
	if err != nil {
		t.Fatalf("Failed to create Bluetooth Manager: %v", err)
	}
	defer btm.Stop()

	connectedDevices, err := btm.AcceptConnections(time.Second * 30)
	if err != nil {
		t.Fatalf("Failed to accept connections: %v", err)
	}
	if len(connectedDevices) != 1 {
		t.Fatalf("Expected 1 connected device, got %d", len(connectedDevices))
	}
}

func Test_StartStopOBEXServer(t *testing.T) {
	btm, err := NewBluetoothManager("SunlightMeter")
	if err != nil {
		t.Fatalf("Failed to create Bluetooth Manager: %v", err)
	}
	defer btm.Stop()

	if err := btm.ControlOBEXServer(true, "/home/sunlight/sunlight-meter"); err != nil {
		t.Fatalf("Failed to start OBEX server: %v", err)
	}
	if err := btm.ControlOBEXServer(false, "/home/sunlight/sunlight-meter"); err != nil {
		t.Fatalf("Failed to stop OBEX server: %v", err)
	}
}
