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

		cleared := original.ClearedMetadata()

		got := original.Metadata("source")
		want := "web"
		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}

		got = cleared.Metadata("source")
		want = ""
		if got != want {
			t.Errorf("Metadata(%q) = %q, want %q", "source", got, want)
		}
	})
}

func TestPassenger(t *testing.T) {
	booking, err := New("ABC123", 2000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	t.Run("len cap def", func(t *testing.T) {
		want := 0
		got := len(booking.passenger)
		if got != want {
			t.Errorf("Passenger length = %d, want %d", got, want)
		}

		want = 4
		got = cap(booking.passenger)
		if got != want {
			t.Errorf("Passenger cap = %d, want %d", got, want)
		}

		booking.AddPassenger("Alice")
		booking.AddPassenger("Bob")

		want = 2
		got = len(booking.passenger)
		if got != want {
			t.Errorf("Passenger length = %d, want %d", got, want)
		}

		want = 4
		got = cap(booking.passenger)
		if got != want {
			t.Errorf("Passenger cap = %d, want %d", got, want)
		}
	})

	t.Run("len cap names", func(t *testing.T) {
		want := "Alice"
		got := booking.passenger[0]

		if got != want {
			t.Errorf("Passenger name[0] = %s, want %s", got, want)
		}

		want = "Bob"
		got = booking.passenger[1]

		if got != want {
			t.Errorf("Passenger name[1] = %s, want %s", got, want)
		}
	})

	t.Run("shared backing array", func(t *testing.T) {
		passenger := booking.passenger
		passenger[0] = "Charlie"

		want := "Charlie"
		got := passenger[0]

		if got != want {
			t.Errorf("Passenger cap = %s, want %s", got, want)
		}
	})

	t.Run("reslice", func(t *testing.T) {
		passenger := booking.passenger
		passenger = passenger[:1]

		want := 1
		got := len(passenger)
		if got != want {
			t.Errorf("Passenger cap = %d, want %d", got, want)
		}

		want = 2
		got = len(booking.passenger)
		if got != want {
			t.Errorf("Booking.passenger cap = %d, want %d", got, want)
		}
	})

}
