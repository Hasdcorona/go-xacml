package policy

import (
	"encoding/xml"
	"pdp/target"
	"pdp/rule"
	"pdp/obligation"
	"pdp/advice"
)

type Policy struct {
	XMLName xml.Name `xml:"Policy"`
	PolicyId string `xml:"PolicyId,attr"`
	RuleCombiningAlgId string `xml:"RuleCombiningAlgId,attr"`
	Description string
	Target Target
	Rules []Rule
	Obligations []Obligation
}