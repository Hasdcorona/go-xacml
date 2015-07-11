/**
	XACML Policy
	See: Syntax 5. Element <Policy> (page )
	http://docs.oasis-open.org/xacml/3.0/xacml-3.0-core-spec-os-en.pdf
*/
package pdp

import (
	"encoding/xml"
)

type Policy struct {
	XMLName xml.Name `xml:"Policy"`
	PolicyId string `xml:"PolicyId,attr"`
	RuleCombiningAlgId string `xml:"RuleCombiningAlgId,attr"`
	Description string
	Target Target
	Rules []Rule
	Obligations []Obligation
	Advice Advice
}