package pdp

import (
	"encoding/xml"
)

type PolicySetDefaults struct {
	XMLName xml.Name `xml:"PolicySetDefaults"`
	// XPathVersion is REQUIRED if PolicySet or Policy contains
	// <AttributeSelector> slements or XPath-based functions
	XPathVersion string `xml:"XPathVersion"`
}
