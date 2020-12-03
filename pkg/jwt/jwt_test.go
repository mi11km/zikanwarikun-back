package jwt

import (
	"testing"
)

func TestGenerateAndParseTokenSuccess(t *testing.T) {
	const id = "7ce27222-ee75-41f0-96eb-ddcc7df9f64e"
	token, err := GenerateToken(id)
	if err != nil {
		t.Errorf("got: %s, want: nil", err)
	}
	parseId, err := ParseToken(token)
	if err != nil {
		t.Errorf("got: %s, want: nil", err)
	}
	if id != parseId {
		t.Errorf("failed to generate token and parse the token from id")
	}
}