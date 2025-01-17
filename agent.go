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
*/

type PiToothAgent struct {
	*agent.SimpleAgent
	l *logrus.Logger
}

func (a *PiToothAgent) AuthorizeService(device dbus.ObjectPath, uuid string) *dbus.Error {
	a.l.Debugf("AuthorizeService called for device %s with UUID %s", device, uuid)
	return a.SimpleAgent.AuthorizeService(device, uuid)
}

func (a *PiToothAgent) DisplayPasskey(device dbus.ObjectPath, passkey uint32, entered uint16) *dbus.Error {
	a.l.Debugf("DisplayPasskey called for device %s with passkey %d and entered %d", device, passkey, entered)
	return a.SimpleAgent.DisplayPasskey(device, passkey, entered)
}

func (a *PiToothAgent) DisplayPinCode(device dbus.ObjectPath, pincode string) *dbus.Error {
	a.l.Debugf("DisplayPinCode called for device %s with pincode %s", device, pincode)
	return a.SimpleAgent.DisplayPinCode(device, pincode)
}

func (a *PiToothAgent) RequestAuthorization(device dbus.ObjectPath) *dbus.Error {
	a.l.Debugf("RequestAuthorization called for device %s", device)
	return a.SimpleAgent.RequestAuthorization(device)
}

func (a *PiToothAgent) RequestConfirmation(path dbus.ObjectPath, passkey uint32) *dbus.Error {
	a.l.Debugf("RequestConfirmation called for path %s with passkey %d", path, passkey)
	return a.SimpleAgent.RequestConfirmation(path, passkey)
}

func (a *PiToothAgent) RequestPasskey(path dbus.ObjectPath) (uint32, *dbus.Error) {
	a.l.Debugf("RequestPasskey called for path %s", path)
	return a.SimpleAgent.RequestPasskey(path)
}

func (a *PiToothAgent) RequestPinCode(path dbus.ObjectPath) (string, *dbus.Error) {
	a.l.Debugf("RequestPinCode called for path %s", path)
	return a.SimpleAgent.RequestPinCode(path)
}

func (a *PiToothAgent) SetPassCode(pinCode string) {
	a.l.Debugf("SetPassCode called with pinCode %s", pinCode)
	a.SimpleAgent.SetPassCode(pinCode)
}

func (a *PiToothAgent) SetPassKey(passkey uint32) {
	a.l.Debugf("SetPassKey called with passkey %d", passkey)
	a.SimpleAgent.SetPassKey(passkey)
}

func (a *PiToothAgent) Interface() string {
	return a.SimpleAgent.Interface()
}

func (a *PiToothAgent) PassCode() string {
	return a.SimpleAgent.PassCode()
}

func (a *PiToothAgent) PassKey() uint32 {
	return a.SimpleAgent.PassKey()
}

func (a *PiToothAgent) Path() dbus.ObjectPath {
	return a.SimpleAgent.Path()
}

func (a *PiToothAgent) Release() *dbus.Error {
	return a.SimpleAgent.Release()
}

func (a *PiToothAgent) Cancel() *dbus.Error {
	return a.SimpleAgent.Cancel()
}
