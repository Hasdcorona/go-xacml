package pdp

import (
	"encoding/xml"
)

type Target struct {
	XMLName xml.Name `xml:"Target"`
	Subjects Subjects `xml:"Subjects"`
	Resources Resources `xml:"Resources"`
	Actions Actions `xml:"Actions"`
}