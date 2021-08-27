package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/reugn/wifiqr"
)

const version = "0.2.1"

var (
	versionParam = flag.Bool("version", false, "Show version.")

	ssidParam   = flag.String("ssid", "", "The name of the wireless network. You'll be prompted to enter the SSID if not set.")
	keyParam    = flag.String("key", "", "A pre-shared key (PSK). You'll be prompted to enter the key if not set.")
	encParam    = flag.String("enc", "WPA2", "The wireless network encryption protocol (WEP, WPA, WPA2).")
	hiddenParam = flag.Bool("hidden", false, "Hidden SSID.")

	fileNameParam = flag.String("file", "", "A png file to write the QR Code (prints to stdout if not set).")
	sizeParam     = flag.Int("size", 256, "Size is both the image width and height in pixels.")
)

func main() {
	flag.Parse()

	if *versionParam {
		fmt.Println("Version: " + version)
		return
	}

	validateArguments()

	config := wifiqr.NewConfig(*ssidParam, *keyParam, *encParam, *hiddenParam)
	q, err := wifiqr.InitCode(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	if *fileNameParam == "" {
		fmt.Println(q.ToSmallString(false))
	} else {
		fileName := validateAndGetFileName()
		err := q.WriteFile(*sizeParam, fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		} else {
			fmt.Println("QR Code was successfully saved to " + fileName + ".")
		}
	}
}

func validateAndGetFileName() string {
	if filepath.Ext(*fileNameParam) != ".png" {
		return *fileNameParam + ".png"
	}
	return *fileNameParam
}

func validateArguments() {
	if *ssidParam == "" {
		fmt.Println("Enter the name of the wireless network (SSID):")
		*ssidParam = readLine()
	}
	if *keyParam == "" {
		fmt.Println("Enter the network key (password):")
		*keyParam = readLine()
	}
	if *ssidParam == "" || *keyParam == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	// convert CRLF to LF
	line = strings.Replace(line, "\n", "", -1)
	return line
}
