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

func TestCheckPasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"normal",
			args{"password", "$2a$14$5oDMaSP7zxfjVvVForxIbeA6S6dMvvKc5iAy5L8pmvBCgfmynglZu"},
			true,
		},
		{
			"fail",
			args{"password", "$2a$14$5oDMaSP7zxfjVvVForxIbeA6S6dMvvKc5iAy5L8pmvBCgfaebglZu"},
			false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := CheckPasswordHash(tt.args.password, tt.args.hash); got != tt.want {
				t.Errorf("CheckPasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnauthenticatedUserAccessError_Error(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"normal", "unauthenticated user access denied"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &UnauthenticatedUserAccessError{}
			if got := m.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
