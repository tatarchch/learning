package booking

import (
	"testing"
	"time"
)

func TestCopySharesMetadata(t *testing.T) {
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
		t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
	}
}

func TestCloneHasIndependentMetadata(t *testing.T) {
	original, err := New("ABC123", 1000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	original.SetMetadata("source", "web")

	cloned := original.Clone()

	cloned.SetMetadata("source", "random")
	cloned.SetMetadata("device", "mobile")

	t.Run("changing existing metadata does not affect original", func(t *testing.T) {

		got := original.Metadata("source")
		want := "web"

		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}

		got = cloned.Metadata("source")
		want = "random"

		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}
	})

	t.Run("adding metadata does not affect original", func(t *testing.T) {

		got := cloned.Metadata("device")
		want := "mobile"

		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "device", got, want)
		}

		got = original.Metadata("device")
		want = ""

		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "device", got, want)
		}
	})
}

func TestClearedMetadata(t *testing.T) {
	t.Run("cleared metadata by pointer receiver", func(t *testing.T) {
		original, err := New("ABC123", 1000, time.Now())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}

		original.SetMetadata("source", "web")
		original.ClearMetadata()

		got := original.Metadata("source")
		want := ""

		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}
	})

	t.Run("new Booking with cleared metadata", func(t *testing.T) {
		original, err := New("ABC123", 1000, time.Now())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}

		original.SetMetadata("source", "web")

		valueBooking := original.ClearedMetadata()

		got := original.Metadata("source")
		want := "web"
		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}

		got = valueBooking.Metadata("source")
		want = ""
		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}
	})
}
