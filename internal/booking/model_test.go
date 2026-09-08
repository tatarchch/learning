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
			t.Errorf("Passengers()[0]  = %s, want %s", got, want)
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

func TestFirstPassengers(t *testing.T) {
	booking, err := New("ABC123", 2000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	booking.AddPassenger("Alice")
	booking.AddPassenger("Bob")
	booking.AddPassenger("Charlie")
	booking.AddPassenger("David")

	t.Run("len cap def", func(t *testing.T) {

		part := booking.FirstPasangers(2)
		want := 2
		got := len(part)
		if got != want {
			t.Errorf("FirstPassengers() len = %d, want %d", got, want)
		}

		want = 4
		got = cap(booking.passengers)
		if got != want {
			t.Errorf("FirstPassengers() cap = %d, want %d", got, want)
		}
	})

	t.Run("append dangerous with mount back array", func(t *testing.T) {
		part := booking.FirstPasangers(2)
		part = append(part, "Eve")

		got := booking.Passengers()[2]
		want := "Eve"
		if got != want {
			t.Errorf("FirstPassengers()[2] = %s, want %s", got, want)
		}
	})

	t.Run("reallocation with unmount back array", func(t *testing.T) {
		all := booking.Passengers()
		all = append(all, "Eve")
		all[0] = "Change"

		got := all[0]
		want := "Change"
		if got != want {
			t.Errorf("all[0] = %s, want %s", got, want)
		}

		got = booking.Passengers()[0]
		want = "Alice"
		if got != want {
			t.Errorf("booking.Passengers()[0] = %s, want %s", got, want)
		}
	})

	t.Run("reallocation with mount back array with max cap in subslice", func(t *testing.T) {
		newBooking, err := New("ABC157", 200, time.Now())
		if err != nil {
			t.Fatalf("New() error = %v", err)
		}

		newBooking.AddPassenger("Alice")
		newBooking.AddPassenger("Bob")
		newBooking.AddPassenger("Charlie")
		newBooking.AddPassenger("David")

		passengers := newBooking.Passengers()

		limited := passengers[:2:2]
		limited = append(limited, "Eve")
		limited[0] = "Changed"

		got := newBooking.Passengers()[0]
		want := "Alice"
		if got != want {
			t.Errorf("newBooking.Passengers()[0] = %s, want %s", got, want)
		}

		got = booking.Passengers()[2]
		want = "Charlie"
		if got != want {
			t.Errorf("newBooking.Passengers()[2] = %s, want %s", got, want)
		}
	})
}
