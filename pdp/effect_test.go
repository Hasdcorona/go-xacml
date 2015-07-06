package effect_test

import (
	. "pdp/effect"
	"testing"
)

func TestEffect(t *testing.T) {
	e := new(Effect)
	if e.Value() != "Deny" {
		t.Error("Empty effect should evaluate to 'Deny'")
	}

	err := e.Set("Permit")
	if err != nil {
		t.Error("Set threw an error when a valid parameter 'Permit' was passed in.")
	}

	if e.Value() != "Permit" {
		t.Errorf("Set() should return 'Permit' but got %s instead.", e.Value())
	}

	err = e.Set("Deny")
	if err != nil {
		t.Error("Set thgrew an error when a valid parameter 'Deny' was passed in.")
	}

	if e.Value() != "Deny" {
		t.Errorf("Set() should return 'Deny' by got %s instead.", e.Value())
	}
}