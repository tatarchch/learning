package booking

import (
	"maps"
	"slices"
	"testing"
	"time"
	"unicode/utf8"
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

		got := booking.Metadata("source")
		want := "web"
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "source", got, want)
		}

		got = booking.Metadata("new")
		want = ""
		if got != want {
			t.Errorf("AllMetadata(%q) = %q, want %q", "new", got, want)
		}
	})
}

func TestMapAssignment(t *testing.T) {
	original := map[string]string{
		"source": "web",
	}

	t.Run("assignment shares state", func(t *testing.T) {
		alias := original
		alias["source"] = "api"

		got := original["source"]
		want := "api"
		if got != want {
			t.Errorf("got = %q, want %q", got, want)
		}
	})

	t.Run("clone has independent state", func(t *testing.T) {
		original["source"] = "web"

		clone := maps.Clone(original)
		clone["source"] = "api"

		got := original["source"]
		want := "web"
		if got != want {
			t.Errorf("got = %q, want %q", got, want)
		}

		got = clone["source"]
		want = "api"
		if got != want {
			t.Errorf("got = %q, want %q", got, want)
		}
	})
}

func TestValidationReference(t *testing.T) {
	tests := []struct {
		name string
		ref  string
		want bool
	}{
		{
			name: "valid letters and digits",
			ref:  "ABC123",
			want: true,
		},
		{
			name: "valid digits",
			ref:  "123123",
			want: true,
		},
		{
			name: "too short",
			ref:  "A12",
			want: false,
		},
		{
			name: "too long",
			ref:  "ABC123123",
			want: false,
		},
		{
			name: "lowercase",
			ref:  "abc123",
			want: false,
		},
		{
			name: "punctuation",
			ref:  "ABC!23",
			want: false,
		},
		{
			name: "space",
			ref:  "ABC 23",
			want: false,
		},
		{
			name: "rus",
			ref:  "АБС123",
			want: false,
		},
		{
			name: "unicode",
			ref:  "Привет",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validReference(tt.ref)

			if got != tt.want {
				t.Errorf("validReference(%q) = %v, want %v", tt.ref, got, tt.want)
			}
		})
	}
}

func TestStringBytesAndRunes(t *testing.T) {
	s := "Привет"

	got := len(s)
	want := 12
	if got != want {
		t.Errorf("len(%q) = %d, want %d", s, got, want)
	}

	got = utf8.RuneCountInString(s)
	want = 6
	if got != want {
		t.Errorf("utf8.RuneCountInString(%q) = %d, want %d", s, got, want)
	}

	offsets := make([]int, 0, 6)
	for i := range s {
		offsets = append(offsets, i)
	}

	wantOffsets := []int{0, 2, 4, 6, 8, 10}
	if !slices.Equal(offsets, wantOffsets) {
		t.Errorf("range indexes = %v, want %v", offsets, wantOffsets)
	}
}
