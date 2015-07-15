package pdp

import (
	"encoding/xml"
)

type PolicyCombinerParameters struct {
	XMLName            xml.Name            `xml:"PolicyCombinerParameters"`
	PolicyRefId        string              `xml:"PolicyRefId,attr"`
	CombinerParameters []CombinerParameter `xml:"CombinerParameter"`
}
