package rule

import (
	"encoding/xml"
	"pdp/target",
	"pdp/effect",
	"pdp/condition",
	"pdp/obligation",
	"pdp/advice"

)

type Rule struct {
	Target Target,
	Effect Effect,
	Condition Condition,
	Obligation Obligation,
	Advice Advice
}