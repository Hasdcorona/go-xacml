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
	XMLName                xml.Name       `xml:"Policy"`
	PolicyId               string         `xml:"PolicyId,attr"`
	RuleCombiningAlgId     string         `xml:"RuleCombiningAlgId,attr"`
	MaxDelegationDepth     uint16         `xml:"MaxDelegationDepth,attr"`
	PolicyDefaults         PolicyDefaults `xml:"PolicyDefaults"`
	Description            string
	Target                 Target                 `xml:"Target"`
	CombinerParameters     []CombinerParameter    `xml:"CombinerParameters"`
	RuleCombinerParameters RuleCombinerParameters `xml:"RuleCombinerParameters"`
	VariableDefinition     VariableDefinition     `xml:"VariableDefinition"`
	Rules                  []Rule                 `xml:"Rule"`
	Obligations            []Obligation           `xml:"Obligations>Obligation"`
	Advice                 Advice                 `xml:"Advice"`
}
