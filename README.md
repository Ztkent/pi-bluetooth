# PiTooth
Go Bluetooth manager for Raspberry Pi devices.   
Quickly enable simple Bluetooth connectivity and file transfer capabilities.

You can import it into your projects, or use it as a standalone tool.

## Features
- Manage the bluetooth service on the Raspberry Pi.
- Accept incoming Bluetooth connections.
- Discover nearby and connected Bluetooth devices.
- Control the OBEX server to support file transfers.

## Requirements
- Any Raspberry Pi device with Bluetooth
- Go 1.21 or later

```bash
## Setup Golang
wget https://go.dev/dl/go1.21.11.linux-armv6l.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.11.linux-armv6l.tar.gz && rm go1.21.11.linux-armv6l.tar.gz

## Add obexd for file transfers
sudo apt install bluez-obexd
```

## Usage

### CLI
```bash
## Build the tool
cd pitooth/cmd
go build -v -o pitooth

## Accept incoming connections with a specified window:
./pitooth -alias=PiToothDevice -acceptConnections -connectionWindow=60 -log=debug

## Enable OBEX server with a path to store received files:
./pitooth -enableObex -obexPath=/path/to/obex/files

## Disable OBEX server:
./pitooth -disableObex
```

### Library
```go
import (
    "log"
    "time"
    "github.com/ztkent/pitooth"
)

// Validate bluetooth functionality, then create a new Bluetooth Manager
btm, err := NewBluetoothManager("YourDeviceName")
if err != nil {
    log.Fatalf("Failed to create Bluetooth Manager: %v", err)
} 

// Become discoverable, and accept incoming connections for 30 seconds
connectedDevices, err := btm.AcceptConnections(time.Second * 30)
if err != nil {
    log.Fatalf("Failed to accept connections: %v", err)
}

// Enable the obexd server, and set the file transfer directory
if err := btm.ControlOBEXServer(true, "/home/sunlight/sunlight-meter"); err != nil {
    log.Fatalf("Failed to start OBEX server: %v", err)
}

// At this point, any connected devices can send files to the Raspberry Pi.
```