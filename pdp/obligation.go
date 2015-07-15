package pdp

import (
	"encoding/xml"
)

type Obligation struct {
	XMLName              xml.Name `xml:"Obligation"`
	ObligationId         string   `xml:"ObligationId,attr"`
	AttributeAssignments []AttributeAssignment
}
