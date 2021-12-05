package wifiqr

import (
	"testing"
)

func TestNewEncryptionProtocol(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    EncryptionProtocol
		wantErr bool
	}{
		{
			name: "WPA",
			args: args{
				t: "WPA",
			},
			want:    WPA,
			wantErr: false,
		},
		{
			name: "WPA2",
			args: args{
				t: "WPA2",
			},
			want:    WPA2,
			wantErr: false,
		},
		{
			name: "WEP",
			args: args{
				t: "WEP",
			},
			want:    WEP,
			wantErr: false,
		},
		{
			name: "NONE",
			args: args{
				t: "NONE",
			},
			want:    NONE,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				t: "invalid",
			},
			want:    WPA2,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEncryptionProtocol(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEncryptionProtocol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewEncryptionProtocol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptionProtocol_String(t *testing.T) {
	tests := []struct {
		name string
		ep   EncryptionProtocol
		want string
	}{
		{
			name: "WPA",
			ep:   WPA,
			want: "WPA",
		},
		{
			name: "WPA2",
			ep:   WPA2,
			want: "WPA2",
		},
		{
			name: "WEP",
			ep:   WEP,
			want: "WEP",
		},
		{
			name: "NONE",
			ep:   NONE,
			want: "NONE",
		},
		{
			name: "NONE",
			ep:   99,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ep.String(); got != tt.want {
				t.Errorf("EncryptionProtocol.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptionProtocol_Code(t *testing.T) {
	tests := []struct {
		name string
		ep   EncryptionProtocol
		want string
	}{
		{
			name: "WPA",
			ep:   WPA,
			want: "WPA",
		},
		{
			name: "WPA2",
			ep:   WPA2,
			want: "WPA2",
		},
		{
			name: "WEP",
			ep:   WEP,
			want: "WEP",
		},
		{
			name: "NONE",
			ep:   NONE,
			want: "nopass",
		},
		{
			name: "NONE",
			ep:   99,
			want: "",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ep.Code(); got != tt.want {
				t.Errorf("EncryptionProtocol.Code() = %v, want %v", got, tt.want)
			}
		})
	}
}
