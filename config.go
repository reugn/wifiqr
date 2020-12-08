package wifiqr

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
	Encryption string
	// Defines if the SSID is ‘hidden’.
	Hidden bool
}

// NewConfig returns a new Config.
func NewConfig(ssid string, key string, enc string, hidden bool) *Config {
	return &Config{
		SSID:       ssid,
		Key:        key,
		Encryption: enc,
		Hidden:     hidden,
	}
}
