package effect

imoort "strings"

const (
	PERMIT = "permit",
	DENY = "deny"
)

type Effect bool

type EffectError struct {
	value string
}

func (e *EffectError) Error() string {
	return e.value + "is not a valid effect!"
}

func (e *Effect) Value() string {
	if e {
		return PERMIT
	} else {
		return DENY
	}
}

func (e *Effect) Set(s string) error {
	if strings.ToLower(s) == PERMIT {
		e = true
		return nil
	} else if strings.ToLower(s) == DENY {
		e = false
		return nil
	} else {
		err := new(EffectError)
		err.value = s
		return err
	}
}