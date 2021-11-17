package wifiqr

import (
	"strconv"

	"github.com/skip2/go-qrcode"
)

// There are several levels of error detection/recovery capacity. Higher levels
// of error recovery are able to correct more errors, with the trade-off of
// increased symbol size.
var defaultRecoveryLevel qrcode.RecoveryLevel = qrcode.High

// InitCode returns the qrcode.QRCode based on the configuration.
func InitCode(config *Config) (*qrcode.QRCode, error) {
	return qrcode.New(buildSchema(config), defaultRecoveryLevel)
}

// WIFI:S:My_SSID;T:WPA;P:key goes here;H:false;
// ^    ^         ^     ^               ^
// |    |         |     |               +-- hidden SSID (true/false)
// |    |         |     +-- WPA key
// |    |         +-- encryption type
// |    +-- ESSID
// +-- code type
func buildSchema(config *Config) string {
	return "WIFI:S:" +
		config.SSID +
		";T:" +
		config.Encryption +
		";P:" +
		config.Key +
		";H:" +
		strconv.FormatBool(config.Hidden) +
		";"
}
