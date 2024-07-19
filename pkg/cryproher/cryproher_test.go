package cryproher

// import (
// 	"testing"
// )

// func TestDecrypt(t *testing.T) {
// 	t.Parallel()
// 	tests := []struct {
// 		input, want string
// 	}{
// 		{"2A4C", "AACCCC"},
// 		{"A2(A2BC3(F4B)D)E", "AABBCFBBBBFBBBBFBBBBDABBCFBBBBFBBBBFBBBBDE"},
// 		{"AB4(CD)2E", "ABCDCDCDCDEE"},
// 	}
// 	for _, test := range tests {

// 		t.Run(test.input, func(t *testing.T) {
// 			decrypted := Decrypt(test.input)

// 			if decrypted != test.want {
// 				t.Errorf("got %s, want %s", decrypted, test.want)
// 			}

// 			t.Parallel()
// 			t.Log(test.input)
// 		})
// 	}
// }
