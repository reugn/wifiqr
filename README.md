# Wi-Fi QR Code generator
<img src="docs/images/qr.png" align='right'/>
Create a QR code with your Wi-Fi login details.

Use Google Lens or other application to scan it and connect automatically.

## Installation
Download and install Go https://golang.org/doc/install.

Clone the repository:  
```sh
git clone https://github.com/reugn/wifiqr.git
```

Build:  
```sh
cd wifiqr
go build ./cmd/wifiqr
```

## Usage
```text
Usage of ./wifiqr:
  -enc string
        The wireless network encryption protocol (WEP, WPA, WPA2). (default "WPA2")
  -file string
        A png file to write the QR Code (prints to stdout if not set).
  -hidden string
        Hidden SSID true/false. (default "false")
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