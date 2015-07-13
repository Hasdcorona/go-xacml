package main

import (
	"./pdp"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	xmlFile, err := os.Open("./ConformanceTests/IIIA020Policy.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	var p pdp.PolicySet
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return
	}
	size := s.Size()
	b := make([]byte,size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	// decode a whole chunk of following XML into the
	// variable p which is a Page (se above)
	//decoder.DecodeElement(&p, &se)

	// Do some stuff with the page.
	// p.Title = CanonicalizeTitle(p.Title)
	// m := filter.MatchString(p.Title)
	// if !m && p.Redir.Title == "" {
	// 	WritePage(p.Title, p.Text)
	// 	total++
	// }
	fmt.Printf("%+v\n",p)

	//fmt.Printf("PolicySet: %+v \n", p)
}
