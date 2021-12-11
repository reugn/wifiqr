package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/reugn/wifiqr"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
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

// validateAndGetFilename adds the .png extension to the
// filename if it doesn't already have one.
func validateAndGetFilename(filename string) string {
	const pngExt = ".png"

	if filepath.Ext(filename) != pngExt {
		filename = filename + pngExt
	}

	return filename
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
	encProtocol, err := wifiqr.NewEncryptionProtocol(protocol)
	if err != nil {
		fmt.Println("Invalid encryption protocol.")
	} else {
		return encProtocol, nil
	}

	prompt := promptui.Select{
		Label: "Select the encryption protocol type",
		Items: []string{
			wifiqr.WPA2.String(),
			wifiqr.WPA.String(),
			wifiqr.WEP.String(),
			wifiqr.NONE.String(),
		},
	}

	_, enc, err := prompt.Run()

	return wifiqr.NewEncryptionProtocol(enc)
}

// process generates the QR code given the parameters and can be
// considered to be a layer below that of the CLI.
func process(ssid, protocolIn, output string, pixels int, key string, keySet bool) int {
	var err error = nil

	ssid, err = validateSSID(ssid)
	if err != nil {
		return 1
	}

	protocol, err := validateEncryption(protocolIn)
	if err != nil {
		return 1
	}

	if protocol != wifiqr.NONE && !keySet {
		key, err = validateKey(key)
		if err != nil {
			return 1
		}
	}

	q, err := generateCode(ssid, key, protocol, false)
	if err != nil {
		return 1
	}

	if outputCode(q, output, pixels) != nil {
		return 1
	}

	return 0
}

// run processes the CLI arguments using Cobra then passes control
// to the process function, using the CLI options as function
// parameters.
func run() int {
	const optionKey = "key"
	var (
		ssid, key, protocolIn, output string
		pixels, exitVal               int
		hidden, keySet                bool
	)

	rootCmd := &cobra.Command{
		Use:   "wifiqr",
		Short: "wifiqr is a WiFi QR code generator",
		Long: `wifiqr is a WiFi QR code generator

It is used to create a QR code containing the login details such as
the name, password, and encryption type. This Android and iOS
compatible QR code can be scanned using Google Lens or other QR code
reader to connect to the network.

If the options necessary for creating the QR code are not given on
the command line, the user will be prompted for the information.`,
		Version: version,
	}

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		keySet = rootCmd.Flags().Changed(optionKey)

		exitVal = process(ssid, protocolIn, output, pixels, key, keySet)
	}

	rootCmd.Flags().StringVarP(&ssid, "ssid", "i", "", "Wireless network name")
	rootCmd.Flags().StringVarP(&key, optionKey, "k", "", "Wireless password (pre-shared key / PSK)")
	rootCmd.Flags().StringVarP(&protocolIn, "protocol", "p", wifiqr.WPA2.String(), "Wireless network encryption protocol ("+
		strings.Join([]string{
			wifiqr.WPA2.String(),
			wifiqr.WPA.String(),
			wifiqr.WEP.String(),
			wifiqr.NONE.String(),
		}, ", ")+
		").")
	rootCmd.Flags().BoolVarP(&hidden, "hidden", "", false, "Hidden SSID")
	rootCmd.Flags().StringVarP(&output, "output", "o", "", "PNG file for output (default stdout)")
	rootCmd.Flags().IntVarP(&pixels, "size", "s", 256, "Image width and height in pixels")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitVal = 1
	}

	return exitVal
}

func main() {
	os.Exit(run())
}
