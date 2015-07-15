package pdp

import (
	"encoding/xml"
)

type ResourceAttributeDesignator struct {
	XMLName     xml.Name `xml:"ResourceAttributeDesignator"`
	AttributeId string   `xml:"AttributeId,attr"`
	DataType    string   `xml:"DataType"`
}
