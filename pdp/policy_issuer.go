package pdp

import (
	"encoding/xml"
)

type PolicyIssuer struct {
	XMLName    xml.Name    `xml:"PolicyIssuer"`
	Content    string      `xml:"Content,innerxml"`
	Attributes []Attribute `xml:"Attribute"`
}
