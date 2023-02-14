package utils

import (
	"fmt"
	"testing"
)

func Test_string(t *testing.T) {
	key := "THIS IS THE KEY!"
	data := "Hello world"

	str, err := EncryptStr(data, key)
	if err != nil {
		t.Errorf("Encrypt failed %v", err)
	}
	fmt.Println(str)
	if "/k83U5B/SRFGtF3HEIbJPA==" != str {
		t.Errorf("Encrypt failed,  %v", str)
	}

	str, err = DecryptStr(str, key)
	if err != nil {
		t.Errorf("Decrypt failed %v", err)
	}
	if str != data {
		t.Errorf("Mismatched. origin %s vs decrypted %s", data, str)
	}
}

func Test_bytes(t *testing.T) {
	key := "THIS IS THE KEY!"
	data := "Hello world"

	encrypt, err := Encrypt([]byte(data), []byte(key))
	if err != nil {
		t.Errorf("Encrypt failed %v", err)
	}

	decrypt, err := Decrypt(encrypt, []byte(key))
	if err != nil {
		t.Errorf("Decrypt failed %v", err)
	}

	decryptStr := string(decrypt)

	if data != decryptStr {
		t.Errorf("Mismatched. origin %s vs decrypted %s", data, decryptStr)
	}
}
