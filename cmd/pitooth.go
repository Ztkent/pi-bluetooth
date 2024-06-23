package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Ztkent/pitooth"
	"github.com/sirupsen/logrus"
)

func main() {
	// Device Settings
	deviceAlias := flag.String("alias", "PiToothDevice", "Bluetooth device alias")
	logLevel := flag.String("log", "info", "Log level (debug, info, error)")
	// Obex Options
	enableObex := flag.Bool("enableObex", false, "Enable OBEX server")
	disableObex := flag.Bool("disableObex", false, "Disable OBEX server")
	obexPath := flag.String("obexPath", "", "Path for OBEX server files")
	// Connection Options
	acceptConnections := flag.Bool("acceptConnections", false, "Accept incoming connections")
	connectionWindow := flag.Int("connectionWindow", 30, "Connection window in seconds")

	flag.Parse()

	// Set up logger
	logger := logrus.New()
	switch *logLevel {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	// Initialize Bluetooth manager with options
	manager, err := pitooth.NewBluetoothManager(*deviceAlias, pitooth.WithLogger(logger))
	if err != nil {
		fmt.Println("Error initializing Bluetooth manager:", err)
		os.Exit(1)
	}

	// Handle options:
	if *enableObex {
		enableObexFunc(manager, *obexPath)
	} else if *disableObex {
		disableObexFunc(manager)
	} else if *acceptConnections {
		acceptConnectionsFunc(manager, *connectionWindow)
	} else {
		fmt.Println("No action specified. Please use -enableObex or -disableObex.")
		os.Exit(1)
	}
}

func enableObexFunc(btm pitooth.BluetoothManager, obexPath string) {
	if obexPath == "" {
		fmt.Println("Error: OBEX path is required when enabling OBEX server.")
		os.Exit(1)
	}
	err := btm.ControlOBEXServer(true, obexPath)
	if err != nil {
		fmt.Println("Error controlling OBEX server:", err)
		os.Exit(1)
	}
	fmt.Println("OBEX server controlled successfully.")
}

func disableObexFunc(btm pitooth.BluetoothManager) {
	err := btm.ControlOBEXServer(false, "")
	if err != nil {
		fmt.Println("Error controlling OBEX server:", err)
		os.Exit(1)
	}
	fmt.Println("OBEX server controlled successfully.")
}

func acceptConnectionsFunc(btm pitooth.BluetoothManager, window int) {
	windowDuration := int64(window)
	if windowDuration <= 0 {
		fmt.Println("Setting connection window to 30 seconds.")
		windowDuration = 30
	}
	devices, err := btm.AcceptConnections(time.Duration(windowDuration))
	if err != nil {
		fmt.Println("Error accepting connections:", err)
		os.Exit(1)
	}
	fmt.Printf("%d active connections.\n", len(devices))
}
