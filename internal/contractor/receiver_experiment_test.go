package contractor

import "testing"

func (c Contractor) renameByValue(name string) {
	c.name = name
}

func (c *Contractor) renameByPointer(name string) {
	c.name = name
}

func TestValueReceiverDoesNotChangeOriginal(t *testing.T) {
	c, err := NewContractor("Old Name", "123")
	if err != nil {
		t.Fatal(err)
	}

	c.renameByValue("New Name")

	if c.Name() != "Old Name" {
		t.Fatalf("expected %q, got %q", "Old Name", c.Name())
	}
}

func TestPointerReceiverChangesOriginal(t *testing.T) {
	c, err := NewContractor("Old Name", "123")
	if err != nil {
		t.Fatal(err)
	}

	c.renameByPointer("New Name")

	if c.Name() != "New Name" {
		t.Fatalf("expected %q, got %q", "New Name", c.Name())
	}
}
