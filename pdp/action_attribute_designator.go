package pdp

import (
	"encoding/xml"
)

type ActionAttributeDesignator struct {
	XMLName xml.Name `xml:"ActionAttributeDesignator"`
	AttributeId string `xml:"AttributeId,attr"`
	DataType string `xml:"DataType,attr"`
}