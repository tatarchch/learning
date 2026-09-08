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

func TestPassengers(t *testing.T) {
	booking, err := New("ABC123", 2000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	t.Run("len cap def", func(t *testing.T) {
		want := 0
		got := len(booking.passengers)
		if got != want {
			t.Errorf("Passenger length = %d, want %d", got, want)
		}

		got = cap(booking.passengers)
		want = 4
		if got != want {
			t.Errorf("Passenger cap = %d, want %d", got, want)
		}

		booking.AddPassenger("Alice")
		booking.AddPassenger("Bob")

		got = len(booking.passengers)
		want = 2
		if got != want {
			t.Errorf("Passenger length = %d, want %d", got, want)
		}

		got = cap(booking.passengers)
		want = 4
		if got != want {
			t.Errorf("Passenger cap = %d, want %d", got, want)
		}
	})

	t.Run("len cap names", func(t *testing.T) {
		got := booking.passengers[0]
		want := "Alice"

		if got != want {
			t.Errorf("Passenger name[0] = %s, want %s", got, want)
		}

		got = booking.passengers[1]
		want = "Bob"

		if got != want {
			t.Errorf("Passenger name[1] = %s, want %s", got, want)
		}
	})

	t.Run("shared backing array", func(t *testing.T) {
		passengers := booking.passengers
		passengers[0] = "Charlie"

		got := booking.Passengers()[0]
		want := "Charlie"

		if got != want {
			t.Errorf("Passenger cap = %s, want %s", got, want)
		}
	})

	t.Run("reslice", func(t *testing.T) {
		passengers := booking.passengers
		passengers = passengers[:1]

		want := 1
		got := len(passengers)
		if got != want {
			t.Errorf("Passenger len = %d, want %d", got, want)
		}

		want = 2
		got = len(booking.passengers)
		if got != want {
			t.Errorf("Booking.passenger len = %d, want %d", got, want)
		}
	})

}
