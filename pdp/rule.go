/**
	XACML Rule
	See: Syntax 5. Element <Rule> (page )
	http://docs.oasis-open.org/xacml/3.0/xacml-3.0-core-spec-os-en.pdf
*/
package pdp

import (
	"encoding/xml"
)

type Rule struct {
	XMLName xml.Name `xml:"Rule"`
	RuleId string `xml:"RuleId,attr"`
	Effect Effect `xml:"Effect,attr"`
	Description string
	Target Target
}