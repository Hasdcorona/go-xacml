package pdp

import (
	"encoding/xml"
)

type PolicySetCombinerParameters struct {
	XMLName xml.Name `xml:"PolicySetCombinerParameters"`
	PolicySetIdRef string `xml:"PolicySetIdRef,attr"`
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}