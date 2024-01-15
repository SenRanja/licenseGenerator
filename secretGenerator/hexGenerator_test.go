package secretGenerator

import "testing"

func TestCipherKeyGenerator(t *testing.T) {
	tests := []struct {
		name       string
		ByteLength int
	}{
		{
			name:       "Generate IV(nonce), 12bytes",
			ByteLength: 12,
		},
		{
			name:       "Generate Cipher Key, 16bytes",
			ByteLength: 16,
		},
		{
			name:       "Generate Cipher Key, 32bytes",
			ByteLength: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CipherKeyGenerator(tt.ByteLength)
		})
	}
}
