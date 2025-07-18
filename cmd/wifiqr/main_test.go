package main

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"testing"

	"github.com/reugn/wifiqr"
)

func byteString(b [32]byte) string {
	s := "[32]byte{ "
	l := len(b) - 1
	for i, x := range b {
		s += strconv.Itoa(int(x))
		if i != l {
			s += ", "
		}
	}

	return s + " }"
}

//nolint:gosec
func Test_generateCode(t *testing.T) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, 7100)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	longString := string(b)

	type args struct {
		ssid     string
		key      string
		encoding wifiqr.EncryptionProtocol
		hidden   bool
	}
	tests := []struct {
		name    string
		args    args
		want    [32]byte
		wantErr bool
	}{
		{
			name: "encoder success",
			args: args{
				ssid:     "testssid",
				key:      "testkeytestkey",
				encoding: wifiqr.WPA2,
				hidden:   false,
			},
			want: [32]byte{117, 113, 240, 31, 70, 131, 178, 237, 61, 56, 190, 135, 145, 86, 173, 81,
				244, 78, 103, 173, 103, 188, 82, 70, 79, 180, 149, 217, 5, 113, 227, 25},
			wantErr: false,
		},
		{
			name: "encoder failure",
			args: args{
				ssid:     longString,
				key:      "test",
				encoding: wifiqr.WPA2,
				hidden:   false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCode(tt.args.ssid, tt.args.key, tt.args.encoding, tt.args.hidden)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if data, err := got.PNG(512); err != nil {
					t.Errorf("generateCode() error generating png: %v", err)
				} else {
					hash := sha256.Sum256(data)
					if tt.want != hash {
						t.Errorf("generateCode() png data does not match wanted hash, got: %v, want %v",
							byteString(hash), byteString(tt.want))
					}
				}
			}
		})
	}
}

func Test_validateAndGetFileName(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no png suffix",
			args: args{
				filename: "imagefilename",
			},
			want: "imagefilename.png",
		},
		{
			name: "with png suffix",
			args: args{
				filename: "imagefilename.png",
			},
			want: "imagefilename.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateAndGetFilename(tt.args.filename); got != tt.want {
				t.Errorf("validateAndGetFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
