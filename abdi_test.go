package abdi

import "testing"

func TestHash(t *testing.T) {

	_, err := Hash("less8", []byte("mykey"))
	if err == nil {
		t.Error("password shorter than MinPasswordLength allowed")
	}

	_, err = Hash("trustno1", []byte("mykey"))
	if err == nil {
		t.Error("password not caught by blacklist")
	}

	str, err := Hash("thisismypassword", []byte("mykey"))
	if err != nil {
		t.Error(err)
	}

	str2, err := Hash("thisismypassword", []byte("mykey"))
	if err != nil {
		t.Error(err)
	}

	if str == str2 {
		t.Error("hashing the same thing twice should have different results due to different salts")
	}
}

func TestCheck(t *testing.T) {
	str, err := Hash("thisismypassword", []byte("mykey"))
	if err != nil {
		t.Error(err)
	}

	err = Check("thisismypassword", *str, []byte("mykey"))
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkHash(b *testing.B) {
	// hash b.N times
	for n := 0; n < b.N; n++ {
		_, err := Hash("thisismypassword", []byte("mykey"))
		if err != nil {
			b.Error(err)
		}
	}
}
