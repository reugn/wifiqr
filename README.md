# Wi-Fi QR Code generator
<img src="docs/images/qr.png" align='right'/>

[![Test Status](https://github.com/reugn/wifiqr/workflows/Test/badge.svg)](https://github.com/reugn/wifiqr/actions?query=workflow%3ATest)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/reugn/wifiqr?tab=doc)](https://pkg.go.dev/github.com/reugn/wifiqr?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/reugn/wifiqr)](https://goreportcard.com/report/github.com/reugn/wifiqr)

Create a QR code with your Wi-Fi login details.

Use Google Lens or other application to scan it and connect automatically.

## Installation
Pick a binary from the [releases](https://github.com/reugn/wifiqr/releases).

### Build from source 
Download and install Go https://golang.org/doc/install.

Get the package:
```sh
go get github.com/reugn/wifiqr
```

Read this [guide](https://golang.org/doc/tutorial/compile-install) on how to compile and install the application.

## Usage
```text
Usage of ./wifiqr:
  -enc string
        The wireless network encryption protocol (WEP, WPA, WPA2). (default "WPA2")
  -file string
        A png file to write the QR Code (prints to stdout if not set).
  -hidden
        Hidden SSID.
  -key string
        A pre-shared key (PSK). You'll be prompted to enter the key if not set.
  -size int
        Size is both the image width and height in pixels. (default 256)
  -ssid string
        The name of the wireless network. You'll be prompted to enter the SSID if not set.
  -version
        Show version.
```

## Usage example
```sh
./wifiqr -ssid some_ssid -key 1234 -file qr.png -size 128
```

## License
MIT