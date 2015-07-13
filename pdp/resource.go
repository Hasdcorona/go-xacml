package pdp

import (
	"encoding/xml"
)

type Resources struct {
	XMLName xml.Name `xml:"Resources"`
	AnyResource AnyResource
	Resources []Resource `xml:"Resource"`
}

type Resource struct {
	XMLName xml.Name `xml:"Resource"`
	ResourceMatch ResourceMatch `xml:"ResourceMatch"`
}