package uniqwords

import (
	"testing"
)

func TestUnique(t *testing.T) {
	got := Unique("a b a c b")
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("Unique=%v want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Unique[%d]=%q want %q", i, got[i], want[i])
		}
	}
}

func TestCount(t *testing.T) {
	got := Count("a b a c")
	if len(got) != 3 {
		t.Fatalf("len=%d want 3", len(got))
	}
	if got[0].Word != "a" || got[0].Count != 2 {
		t.Errorf("first=%v want a:2", got[0])
	}
	if got[1].Word != "b" || got[2].Word != "c" {
		t.Errorf("order wrong: %v", got)
	}
}
