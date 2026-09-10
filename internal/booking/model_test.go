package booking

import (
	"maps"
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

func TestCloneHasIndependentState(t *testing.T) {
	original, err := New("ABC123", 1000, time.Now())

	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	original.SetMetadata("source", "web")
	original.AddPassenger("Alice")
	original.AddPassenger("Bob")

	cloned := original.Clone()

	cloned.SetMetadata("source", "random")
	cloned.SetMetadata("device", "mobile")
	cloned.AddPassenger("Charlie")

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

	t.Run("adding passenger does not affect original", func(t *testing.T) {
		got := len(original.Passengers())
		want := 2

		if got != want {
			t.Errorf("len(original.Passengers()) = %d, want %d", got, want)
		}

		got = len(cloned.Passengers())
		want = 3

		if got != want {
			t.Errorf("len(cloned.Passengers()) = %d, want %d", got, want)
		}

		gotPassenger := cloned.Passengers()[2]
		wantPassenger := "Charlie"

		if gotPassenger != wantPassenger {
			t.Errorf("Passengers()[2] = %q, want %q", gotPassenger, wantPassenger)
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

	t.Run("returns independent copy", func(t *testing.T) {
		passengers := booking.Passengers()
		passengers[0] = "Charlie"

		got := booking.Passengers()[0]
		want := "Alice"

		if got != want {
			t.Errorf("Passengers()[0]  = %s, want %s", got, want)
		}
	})
}

func TestAllMetadata(t *testing.T) {
	booking, err := New("ABC123", 2000, time.Now())
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	booking.SetMetadata("source", "web")
	booking.SetMetadata("device", "mobile")

	metadata := booking.AllMetadata()

	t.Run("test data from AllMetadata()", func(t *testing.T) {
		got := metadata["source"]
		want := "web"
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "source", got, want)
		}

		got = metadata["device"]
		want = "mobile"
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "device", got, want)
		}

		gotLen := len(metadata)
		wantLen := 2
		if gotLen != wantLen {
			t.Errorf("gotLen = %d, want %d", gotLen, wantLen)
		}
	})

	t.Run("test defensive copy", func(t *testing.T) {
		metadata["source"] = "changed"
		metadata["new"] = "value"

		got := booking.metadata["source"]
		want := "web"
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "source", got, want)
		}

		got = booking.metadata["new"]
		want = ""
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "new", got, want)
		}
	})
}

func TestMapAssigment(t *testing.T) {
	original := map[string]string{
		"source": "web",
	}

	t.Run("test aliasing have a shared state", func(t *testing.T) {
		alias := original
		alias["source"] = "api"

		got := original["source"]
		want := "api"
		if got != want {
			t.Errorf("got = %q, want %q", got, want)
		}
	})

	t.Run("test clone haven`t a shared state", func(t *testing.T) {
		original["source"] = "web"

		clone := maps.Clone(original)
		clone["source"] = "api"

		got := original["source"]
		want := "web"
		if got != want {
			t.Errorf("got = %q, want %q", got, want)
		}
	})
}
