package wifiqr

import (
	"hash/fnv"
	"reflect"
	"testing"
)

func TestCode(t *testing.T) {
	config := NewConfig("ssid1", "1234", WPA2, false)
	qrCode, err := InitCode(config)
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(t, hashCode(qrCode.ToString(false)), 550955445)
}

func hashCode(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("%v != %v", a, b)
	}
}

func Test_escapeString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all escapes",
			args: args{
				in: `abc\;,":xyz`,
			},
			want: `abc\\\;\,\"\:xyz`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeString(tt.args.in); got != tt.want {
				t.Errorf("escapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildSchema(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no encryption protocol",
			args: args{
				config: &Config{
					SSID:       "ssid1",
					Encryption: NONE,
				},
			},
			want: `WIFI:S:ssid1;T:nopass;P:;H:false;`,
		},
		{
			name: "WEP encryption protocol",
			args: args{
				config: &Config{
					SSID:       "ssid1",
					Key:        "key1",
					Encryption: WEP,
				},
			},
			want: `WIFI:S:ssid1;T:WEP;P:key1;H:false;`,
		},
		{
			name: "WPA encryption protocol",
			args: args{
				config: &Config{
					SSID:       "ssid1",
					Key:        "key1",
					Encryption: WPA,
				},
			},
			want: `WIFI:S:ssid1;T:WPA;P:key1;H:false;`,
		},
		{
			name: "WPA2 encryption protocol",
			args: args{
				config: &Config{
					SSID:       "ssid1",
					Key:        "key1",
					Encryption: WPA2,
				},
			},
			want: `WIFI:S:ssid1;T:WPA2;P:key1;H:false;`,
		},
		{
			name: "WPA2 encryption protocol, hidden",
			args: args{
				config: &Config{
					SSID:       "ssid1",
					Key:        "key1",
					Encryption: WPA2,
					Hidden:     true,
				},
			},
			want: `WIFI:S:ssid1;T:WPA2;P:key1;H:true;`,
		},
		{
			name: "escaped characters, WPA2 encryption protocol",
			args: args{
				config: &Config{
					SSID:       `abc\;,":xyz`,
					Key:        `xyz\;,":abc`,
					Encryption: WPA2,
					Hidden:     false,
				},
			},
			want: `WIFI:S:abc\\\;\,\"\:xyz;T:WPA2;P:xyz\\\;\,\"\:abc;H:false;`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildSchema(tt.args.config); got != tt.want {
				t.Errorf("buildSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}
