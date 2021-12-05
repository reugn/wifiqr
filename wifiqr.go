package wifiqr

import (
	"strconv"
	"strings"

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

// escapeString escapes the special characters with a backslash.
func escapeString(s string) string {
	// https://github.com/zxing/zxing/wiki/Barcode-Contents#wi-fi-network-config-android-ios-11

	for _, c := range []byte{'\\', ';', ',', '"', ':'} {
		s = strings.Replace(s, string(c), `\`+string(c), -1)
	}

	return s
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
		escapeString(config.SSID) +
		";T:" +
		config.Encryption.Code() +
		";P:" +
		escapeString(config.Key) +
		";H:" +
		strconv.FormatBool(config.Hidden) +
		";"
}
