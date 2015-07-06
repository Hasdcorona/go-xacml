package pdp

import (
	"encoding/xml"
	"pdp/policy"
	"pdp/policy_set"
	"pdp/rule"
	"strings"
	"bytes"
	"io"
)

type Decider interface {
	Decide(req *Request) *Decision, error
}

type Policer interface {
	Police() Match{}, error
}

type Match interface {
	Match(req *Request) *Response, error
}

type PolicyDecisionPoint struct {
	
}