package policy

import (
	"encoding/xml"
	"pdp/target",
	"pdp/combining",
	"pdp/rule",
	"pdp/obligation",
	"pdp/advice"
)

type Policy struct {
	Target Target,
	RuleCombiningAlgorithm RuleCOmbiningAlgorithm,
	Rules []Rule,
	Obligation Obligation,
	Advice Advice
}