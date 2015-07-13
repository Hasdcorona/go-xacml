package pdp

import (
	"encoding/xml"
)

type SubjectAttributeDesignator struct {
	XMLName xml.Name `xml:"SubjectAttributeDesignator"`
	AttributeId string `xml:"AttributeId,attr"`
	DataType string `xml:"DataType,attr"`
}