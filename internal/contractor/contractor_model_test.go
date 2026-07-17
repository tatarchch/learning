package contractor

import "testing"

func TestContractor_ChangeDescription_WithValue(t *testing.T) {
	c, err := NewContractor("Old Name", "123")
	if err != nil {
		t.Fatal(err)
	}

	c.ChangeDescription("New Description")

	if c.Description() != "New Description" {
		t.Fatalf("expected %q, got %q", "New Description", c.Description())
	}
}

func TestContractor_ChangeDescription_WithEmpty(t *testing.T) {
	c, err := NewContractor("Old Name", "123")
	if err != nil {
		t.Fatal(err)
	}

	c.ChangeDescription("")

	if c.Description() != "" {
		t.Fatalf("expected %q, got %q", "", c.Description())
	}
}
