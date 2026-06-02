package mcmauthorizer

import "testing"

func TestBuildClusterAPIEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected string
	}{
		{
			name:     "plain ipv6",
			address:  "2001:cafe:43::1",
			expected: "https://[2001:cafe:43::1]",
		},
		{
			name:     "unbracketed ipv6 with port",
			address:  "2001:cafe:43::1:443",
			expected: "https://[2001:cafe:43::1]:443",
		},
		{
			name:     "bracketed ipv6 with port",
			address:  "[2001:cafe:43::1]:443",
			expected: "https://[2001:cafe:43::1]:443",
		},
		{
			name:     "ipv4 with port",
			address:  "10.0.0.5:6443",
			expected: "https://10.0.0.5:6443",
		},
		{
			name:     "hostname with port",
			address:  "api.example.com:6443",
			expected: "https://api.example.com:6443",
		},
		{
			name:     "hostname only",
			address:  "api.example.com",
			expected: "https://api.example.com",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := buildClusterAPIEndpoint(tc.address)
			if actual != tc.expected {
				t.Fatalf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}