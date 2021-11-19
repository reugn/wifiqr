package wifiqr

import (
	"errors"
	"strings"
)

type EncryptionProtocol int

const (
	WPA2 EncryptionProtocol = iota
	WPA
	WEP

	wpa2Str = "WPA2"
	wpaStr  = "WPA"
	wepStr  = "WEP"
)

func (ep EncryptionProtocol) String() string {
	switch ep {
	case WPA2:
		return wpa2Str
	case WPA:
		return wpaStr
	case WEP:
		return wepStr
	}
	return ""
}

func NewEncryptionProtocol(t string) (EncryptionProtocol, error) {
	switch strings.ToUpper(t) {
	case wpa2Str:
		return WPA2, nil
	case wpaStr:
		return WPA, nil
	case wepStr:
		return WEP, nil
	}
	return WPA2, errors.New("no such protocol")
}

// Config is the Wi-Fi network configuration parameters.
type Config struct {
	// The Service Set Identifier (SSID) is the name of the wireless network.
	// It can be contained in the beacons sent out by APs, or it can be ‘hidden’ so that clients
	// who wish to associate must first know the name of the network. Early security guidance was
	// to hide the SSID of your network, but modern networking tools can detect the SSID by simply
	// watching for legitimate client association, as SSIDs are transmitted in cleartext.
	SSID string
	// A pre-shared key (PSK).
	Key string
	// The wireless network encryption protocol (WEP, WPA, WPA2).
	Encryption EncryptionProtocol
	// Defines if the SSID is ‘hidden’.
	Hidden bool
}

// NewConfig returns a new Config.
func NewConfig(ssid string, key string, enc EncryptionProtocol, hidden bool) *Config {
	return &Config{
		SSID:       ssid,
		Key:        key,
		Encryption: enc,
		Hidden:     hidden,
	}
}
