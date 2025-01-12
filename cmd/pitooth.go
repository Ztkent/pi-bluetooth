package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ztkent/pitooth"
)

/*
	Supports command-line functionality for PiTooth.
*/

func main() {
	deviceAlias := flag.String("alias", "PiToothDevice", "Bluetooth device alias")
	logLevel := flag.String("log", "info", "Log level (debug, info, error)")
	enableObexFlag := flag.Bool("enableObex", false, "Enable OBEX server")
	disableObexFlag := flag.Bool("disableObex", false, "Disable OBEX server")
	obexPathFlag := flag.String("obexPath", "", "Path for OBEX server files")
	acceptConnectionsFlag := flag.Bool("acceptConnections", false, "Accept incoming connections")
	connectionWindowFlag := flag.Int("connectionWindow", 30, "Connection window in seconds")
	flag.Parse()

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

	manager, err := pitooth.NewBluetoothManager(*deviceAlias, pitooth.WithLogger(logger))
	if err != nil {
		fmt.Println("Error initializing Bluetooth manager:", err)
		os.Exit(1)
	}

	if *enableObexFlag {
		enableObex(manager, *obexPathFlag)
	} else if *disableObexFlag {
		disableObex(manager)
	} else if *acceptConnectionsFlag {
		acceptConnections(manager, *connectionWindowFlag)
	} else {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Println("PiTooth is a command-line tool for managing Bluetooth connections and OBEX server.")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("\tEnable OBEX server with a specific path for server files:")
		fmt.Println("\t\t" + os.Args[0] + " -enableObex -obexPath=/path/to/obex/files")
		fmt.Println("\tDisable OBEX server:")
		fmt.Println("\t\t" + os.Args[0] + " -disableObex")
		fmt.Println("\tAccept incoming connections with a custom connection window:")
		fmt.Println("\t\t" + os.Args[0] + " -acceptConnections -connectionWindow=60")
		os.Exit(1)
	}
}

func enableObex(btm pitooth.BluetoothManager, obexPath string) {
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

func disableObex(btm pitooth.BluetoothManager) {
	err := btm.ControlOBEXServer(false, "")
	if err != nil {
		fmt.Println("Error controlling OBEX server:", err)
		os.Exit(1)
	}
	fmt.Println("OBEX server controlled successfully.")
}

func acceptConnections(btm pitooth.BluetoothManager, window int) {
	windowDuration := int64(window)
	if windowDuration <= 0 {
		fmt.Println("Setting connection window to 30 seconds.")
		windowDuration = 30
	}
	err := btm.AcceptConnections(time.Duration(windowDuration) * time.Second)
	if err != nil {
		fmt.Println("Error accepting connections:", err)
		os.Exit(1)
	}
}
