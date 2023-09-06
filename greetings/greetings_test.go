package greetings

import (
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		nameArg string
		want    string
		wantErr bool
	}{
		{
			name:    "Empty name",
			nameArg: "",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Simple name",
			nameArg: "Eduardo",
			want:    "Hello Eduardo! Welcome!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hello(tt.nameArg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
