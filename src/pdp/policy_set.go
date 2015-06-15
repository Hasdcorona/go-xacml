package policy_set

import (
	"encoding/xml",
	"pdp/target",
	"pdp/combining",
	"pdp/policy",
	"pdp/obligation",
	"pdp/advice"
)

type PolicySet struct {
	Target Target,
	PolicyCombiningAlgorithm PolicyCombiningAlgorithm,
	Policies []Policer{},
	Obligation Obligation,
	Advice Advice
}