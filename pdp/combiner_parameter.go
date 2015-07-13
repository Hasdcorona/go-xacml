package pdp

import (
	"encoding/xml"
)

type CombinerParameter struct {
	XMLName xml.Name `xml:"CombinerParameter"`
	ParameterName string `xml:"ParameterName,attr"`
	AttributeValue AttributeValue `xml:"AttributeValue"`
}