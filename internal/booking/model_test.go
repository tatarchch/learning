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

	t.Run("source doesn`t change original", func(t *testing.T) {
		cloned.SetMetadata("source", "random")

		got := original.Metadata("source")
		want := "web"

		if got != want {
			t.Errorf("got %q, have %q", got, want)
		}

		got = cloned.Metadata("source")
		want = "random"

		if got != want {
			t.Errorf("got %q, have %q", got, want)
		}
	})

	t.Run("source doesn`t add at original", func(t *testing.T) {
		cloned.SetMetadata("device", "mobile")

		got := cloned.Metadata("device")
		want := "mobile"

		if got != want {
			t.Errorf("got %q, have %q", got, want)
		}

		got = original.Metadata("device")
		want = ""

		if got != want {
			t.Errorf("got %q, have %q", got, want)
		}
	})

}
