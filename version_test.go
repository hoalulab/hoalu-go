package hoalu

import "testing"

func TestVersion(t *testing.T) {
	if Version != "dev" {
		t.Fatalf("Version = %q, want %q", Version, "dev")
	}
}
