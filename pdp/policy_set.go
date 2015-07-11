package pdp

import (
	"encoding/xml"
)

type PolicySet struct {
	XMLName xml.Name `xml:"PolicySet"`
	Xmlns string `xml:"xmlns,attr"`
	Xsi string `xml:"xmlns:xsi,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
	PolicySetId string `xml:"PolicySetId,attr"`
	PolicyCombiningAlgId string `xml:"PolicyCombiningAlgId,attr"`
	Description string
	Target Target
	Policies []Policy `xml:"Policy"`
	Obligations []Obligation
	Advice Advice
}