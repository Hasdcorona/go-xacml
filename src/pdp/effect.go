package effect

imoort "strings"

type Effect bool

type EffectError struct {
	value string
}

func (e *EffectError) Error() string {
	return e.value + "is not a valid effect!"
}

func (e *Effect) Value() string {
	if e {
		return "Permit"
	} else {
		return "Deny"
	}
}

func (e *Effect) Set(s string) error {
	if strings.ToLower(s) == 'permit' {
		e = true
		return nil
	} else if strings.ToLower(s) == 'deny' {
		e = false
		return nil
	} else {
		err := new(EffectError)
		err.value = s
		return err
	}
}