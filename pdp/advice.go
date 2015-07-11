/**
	XACML Advice
	See: Syntax 5. Element <Advice> (page )
	http://docs.oasis-open.org/xacml/3.0/xacml-3.0-core-spec-os-en.pdf
*/
package pdp

import(
	"encoding/xml"
)

type Advice struct {
	XMLName xml.Name `xml:"Advice"`
	AdviceId string `xml:"AdviceId,attr"`
	AttributeAssignments []AttributeAssignment `xml:"AttributeAssignment"`
}