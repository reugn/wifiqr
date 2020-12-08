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
	schema := buildSchema(config)

	q, err := qrcode.New(schema, defaultRecoveryLevel)
	if err != nil {
		return nil, err
	}

	return q, nil
}

// WIFI:S:My_SSID;T:WPA;P:key goes here;H:false;
// ^    ^         ^     ^               ^
// |    |         |     |               +-- hidden SSID (true/false)
// |    |         |     +-- WPA key
// |    |         +-- encryption type
// |    +-- ESSID
// +-- code type
func buildSchema(config *Config) string {
	var sb strings.Builder
	sb.WriteString("WIFI:S:")
	sb.WriteString(config.SSID)
	sb.WriteString(";T:")
	sb.WriteString(config.Encryption)
	sb.WriteString(";P:")
	sb.WriteString(config.Key)
	sb.WriteString(";H:")
	sb.WriteString(strconv.FormatBool(config.Hidden))
	sb.WriteString(";")
	return sb.String()
}
