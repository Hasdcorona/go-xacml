package pdp

import (
	"encoding/xml"
)

type PolicySet struct {
	XMLName              xml.Name `xml:"PolicySet"`
	Xmlns                string   `xml:"xmlns,attr"`
	Xsi                  string   `xml:"xmlns:xsi,attr"`
	SchemaLocation       string   `xml:"schemaLocation,attr"`
	PolicySetId          string   `xml:"PolicySetId,attr"`
	PolicyCombiningAlgId string   `xml:"PolicyCombiningAlgId,attr"`
	Version              string   `xml:"Version,attr"`
	MaxDelegationDepth   uint16   `xml:"MaxDelegationDepth"`
	Description          string
	PolicyIssuer         PolicyIssuer      `xml:"PolicyIssuer"`
	PolicySetDefaults    PolicySetDefaults `xml:"PolicySetDefaults"`
	Target               Target            `xml:"Target"`
	Policies             []Policy          `xml:"Policy"`
	PolicySets           []PolicySet       `xml:"PolicySet"`
	PolicySetIdReference PolicySetIdReference
	PolicyIdReference    PolicyIdReference
	Obligations          []Obligation `xml:"Obligations>Obligation"`
	//Advice []Advice `xml:"Advice"`
	CombinerParameters          []CombinerParameter `xml:"CombinerParameters"`
	PolicyCombinerParameters    []CombinerParameter `xml:"PolicyCombinerParameters"`
	PolicySetCombinerParameters []CombinerParameter `xml:"PolicySetCombinerParameters"`
}
