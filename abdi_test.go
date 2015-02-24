package abdi

import "testing"

func TestHash(t *testing.T) {
	Key = []byte("mykey")

	_, err := Hash("less8")
	if err == nil {
		t.Error("password shorter than MinPasswordLength allowed")
	}

	_, err = Hash("trustno1")
	if err == nil {
		t.Error("password not caught by blacklist")
	}

	str, err := Hash("thisismypassword")
	if err != nil {
		t.Error(err)
	}

	str2, err := Hash("thisismypassword")
	if err != nil {
		t.Error(err)
	}

	if str == str2 {
		t.Error("hashing the same thing twice should have different results due to different salts")
	}
}

func TestCheck(t *testing.T) {
	Key = []byte("mykey")

	str, err := Hash("thisismypassword")
	if err != nil {
		t.Error(err)
	}

	err = Check("thisismypassword", str)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkHash(b *testing.B) {
	Key = []byte("mykey")

	// hash b.N times
	for n := 0; n < b.N; n++ {
		_, err := Hash("thisismypassword")
		if err != nil {
			b.Error(err)
		}
	}
}
