/**
XACML AttributeAssignment
See: Syntax 5. Element <AttributeAssignment> (page )
http://docs.oasis-open.org/xacml/3.0/xacml-3.0-core-spec-os-en.pdf
*/
package pdp

import (
	"encoding/xml"
)

type AttributeAssignment struct {
	XMLName     xml.Name `xml:"AttributeAssignment"`
	AttributeId string   `xml:"AttributeId,attr"`
	Value       string   `xml:",chardata"`
	Category    string
	Issuer      string
}
