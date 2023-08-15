package greetings

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		nameArg string
		want    string
	}{
		{
			name:    "Empty name",
			nameArg: "",
			want:    "Hello ! Welcome!",
		},
		{
			name:    "Simple name",
			nameArg: "Eduardo",
			want:    "Hello Eduardo! Welcome!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.nameArg); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
