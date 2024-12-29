package pitooth

import (
	"github.com/godbus/dbus/v5"
	"github.com/muka/go-bluetooth/bluez/profile/agent"
	"github.com/sirupsen/logrus"
)

/*
	An Agent is how bluetooth controls the pairing process.
	It is responsible for displaying the passkey, pincode, etc. to the user and handling the user's response.
	Implementing this custom agent allows us to tap into the pairing process.
	In this case, the goal is to allow trusted pairing without additional user interaction.
*/

type PiToothAgent struct {
	*agent.SimpleAgent
	l *logrus.Logger
}

func (a *PiToothAgent) RequestPinCode(device dbus.ObjectPath) (string, *dbus.Error) {
	a.l.Println("RequestPinCode called, returning empty string")
	return "", nil
}

func (a *PiToothAgent) RequestPasskey(device dbus.ObjectPath) (uint32, *dbus.Error) {
	a.l.Println("RequestPasskey called, returning zero")
	return 0, nil
}

func (a *PiToothAgent) DisplayPasskey(device dbus.ObjectPath, passkey uint32, entered uint16) *dbus.Error {
	a.l.Println("DisplayPasskey called, ignoring")
	return nil
}

func (a *PiToothAgent) RequestConfirmation(device dbus.ObjectPath, passkey uint32) *dbus.Error {
	a.l.Println("RequestConfirmation called, auto-confirming")
	return nil
}

func (a *PiToothAgent) AuthorizeService(device dbus.ObjectPath, uuid string) *dbus.Error {
	a.l.Println("AuthorizeService called, auto-authorizing")
	return nil
}
