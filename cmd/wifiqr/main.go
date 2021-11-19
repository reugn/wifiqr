package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/reugn/wifiqr"
	"github.com/skip2/go-qrcode"
)

const (
	ssidDesc = "network name (SSID)"
	pskDesc  = "network key (password)"
)

var version string = "develop"

// generateCode creates a WiFi QR code.
func generateCode(ssid, key string, encoding wifiqr.EncryptionProtocol, hidden bool) (*qrcode.QRCode, error) {
	q, err := wifiqr.InitCode(
		wifiqr.NewConfig(ssid, key, encoding, hidden),
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error generating WiFi QR code:", err)
	}

	return q, err
}

// saveCode saves the QR code to a file.
func saveCode(q *qrcode.QRCode, filename string, size int) error {
	if err := q.WriteFile(size, validateAndGetFilename(filename)); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving WiFi QR code:", err)
		return err
	}

	fmt.Println("WiFi QR code saved to " + filename + ".")

	return nil
}

// outputCode outputs the QR code to stdout or to a file.
func outputCode(q *qrcode.QRCode, filename string, size int) error {
	if filename == "" {
		fmt.Println(q.ToSmallString(false))
		return nil
	}

	return saveCode(q, filename, size)
}

// validateAndGetFilename adds the .png extension to the
// filename if it doesn't already have one.
func validateAndGetFilename(filename string) string {
	const pngExt = ".png"

	if filepath.Ext(filename) != pngExt {
		filename = filename + pngExt
	}

	return filename
}

// getInput gets user input using promptui.
func getInput(prompt string, validate func(string) error) (string, error) {
	return (&promptui.Prompt{
		Label:    prompt,
		Validate: validate,
	}).Run()
}

// inputValidator is a generic function for getting user input if the value is empty.
func inputValidator(value, prompt string, validate func(string) error) (string, error) {
	var err error = nil

	if value == "" {
		value, err = getInput(prompt, validate)
	}

	return value, err
}

// validateSSID	gets the SSID from the user if it is empty.
func validateSSID(ssid string) (string, error) {
	return inputValidator(ssid, "Enter the "+ssidDesc, func(input string) error {
		if len(input) == 0 {
			return errors.New("empty")
		}

		// https://serverfault.com/questions/45439/what-is-the-maximum-length-of-a-wifi-access-points-ssid
		if len(input) > 32 {
			return errors.New("maximum length exceeded")
		}

		return nil
	})
}

// validateKey gets the key from the user if it is empty.
func validateKey(key string) (string, error) {
	return inputValidator(key, "Enter the "+pskDesc, func(input string) error {
		if len(input) == 0 {
			return errors.New("empty")
		}

		// no check for maximum length because it can get messy
		return nil
	})
}

// validateEncryption gets the encryption protocol from the user
// if the provided protocol is empty or not valid.
func validateEncryption(protocol string) (wifiqr.EncryptionProtocol, error) {
	if protocol != "" {
		encProtocol, err := wifiqr.NewEncryptionProtocol(protocol)
		if err != nil {
			fmt.Println("Invalid encryption protocol.")
		} else {
			return encProtocol, nil
		}
	}

	prompt := promptui.Select{
		Label: "Select the encryption type",
		Items: []string{
			wifiqr.WPA2.String(),
			wifiqr.WPA.String(),
			wifiqr.WEP.String(),
		},
	}

	_, enc, err := prompt.Run()

	encProtocol, err := wifiqr.NewEncryptionProtocol(enc)
	if err != nil {
		return encProtocol, err
	}

	return encProtocol, err
}

// run is tne primary function for the program.
func run() int {
	var (
		versionParam = flag.Bool("version", false, "Show version.")

		ssidParam = flag.String("ssid", "", "The name of the wireless network. You'll be prompted to enter the SSID if not set.")
		keyParam  = flag.String("key", "", "A pre-shared key (PSK). You'll be prompted to enter the key if not set.")
		encParam  = flag.String("enc", wifiqr.WPA2.String(),
			"The wireless network encryption protocol ("+
				strings.Join([]string{wifiqr.WPA2.String(), wifiqr.WPA.String(), wifiqr.WEP.String()}, ", ")+
				").")
		hiddenParam = flag.Bool("hidden", false, "Hidden SSID.")

		fileNameParam = flag.String("file", "", "A png file to write the QR Code (prints to stdout if not set).")
		sizeParam     = flag.Int("size", 256, "Size is both the image width and height in pixels.")

		err error
	)

	flag.Parse()

	if *versionParam {
		fmt.Println("Version:", version)
		return 0
	}

	*ssidParam, err = validateSSID(*ssidParam)
	if err != nil {
		return 1
	}

	protocol, err := validateEncryption(*encParam)
	if err != nil {
		return 1
	}

	*keyParam, err = validateKey(*keyParam)
	if err != nil {
		return 1
	}

	q, err := generateCode(*ssidParam, *keyParam, protocol, *hiddenParam)
	if err != nil {
		return 1
	}

	if outputCode(q, *fileNameParam, *sizeParam) != nil {
		return 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
