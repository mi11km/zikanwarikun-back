package users

import "testing"

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	const password = ")Xa+i$4*q(&3"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("got: %s, want: nil", err)
	}

	correct := CheckPasswordHash(password, hash)
	if !correct {
		t.Errorf("got: %t, want: true", correct)
	}
}
