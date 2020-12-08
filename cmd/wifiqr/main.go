package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/reugn/wifiqr"
)

const version = "0.1.0"

var (
	versionParam = flag.Bool("version", false, "Show version.")

	ssidParam   = flag.String("ssid", "", "The name of the wireless network. You'll be prompted to enter the SSID if not set.")
	keyParam    = flag.String("key", "", "A pre-shared key (PSK). You'll be prompted to enter the key if not set.")
	encParam    = flag.String("enc", "WPA2", "The wireless network encryption protocol (WEP, WPA, WPA2).")
	hiddenParam = flag.String("hidden", "false", "Hidden SSID true/false.")

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

	config := wifiqr.NewConfig(*ssidParam, *keyParam, *encParam, validateAndGetHidden())
	q, err := wifiqr.InitCode(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	if *fileNameParam == "" {
		fmt.Println(q.ToSmallString(false))
	} else {
		fileName := validateAndGetFileName()
		err := q.WriteFile(*sizeParam, fileName)
		if err != nil {
			fmt.Println(err)
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

func validateAndGetHidden() bool {
	hidden, err := strconv.ParseBool(*hiddenParam)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hidden
}

func validateArguments() {
	if *ssidParam == "" {
		fmt.Println("Enter the name of the wireless network (SSID):")
		fmt.Scan(ssidParam)
	}
	if *keyParam == "" {
		fmt.Println("Enter the network key (password):")
		fmt.Scan(keyParam)
	}
	if *ssidParam == "" || *keyParam == "" {
		flag.Usage()
		os.Exit(1)
	}
}
