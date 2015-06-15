package effect

import "strings"

type Effect struct {
	value bool
}

type EffectError struct {
	value string
}

func (e *EffectError) Error() string {
	return e.value + " is not a valid effect!"
}

func (e *Effect) Value() string {
	if e.value == true {
		return "Permit"
	} else {
		return "Deny"
	}
}

func (e *Effect) Set(s string) error {
	if strings.ToLower(s) == "permit" {
		e.value = true
		return nil
	} else if strings.ToLower(s) == "deny" {
		e.value = false
		return nil
	} else {
		err := new(EffectError)
		err.value = s
		return err
	}
}