package booking

import (
	"testing"
	"time"
)

func TestBookingCopySharesMetadata(t *testing.T) {
	original, err := New("ABC123", 1000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	original.SetMetadata("source", "web")

	copied := original

	copied.SetMetadata("source", "random")

	got := original.Metadata("source")
	want := "random"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBookingCloneHasIndependentMetadata(t *testing.T) {
	original, err := New("ABC123", 1000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	original.SetMetadata("source", "web")

	cloned := original.Clone()
	cloned.SetMetadata("source", "random")

	got := original.Metadata("source")
	have := cloned.Metadata("source")

	if got == have {
		t.Errorf("got %q, have %q", got, have)
	}
}
