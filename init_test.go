package iso6391

import "testing"

func TestCodes(t *testing.T) {
	if len(Codes) != 184 {
		t.Fatalf("Codes must have 184 entries, but was: %v", len(Codes))
	}
}

func TestNames(t *testing.T) {
	if len(Names) != 184 {
		t.Fatalf("Names must have 184 entries, but was: %v", len(Codes))
	}
}

func TestNativeNames(t *testing.T) {
	if len(NativeNames) != 184 {
		t.Fatalf("NativeNames must have 184 entries, but was: %v", len(Codes))
	}
}
