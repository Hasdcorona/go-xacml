package main

import (
	"./pdp"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	xmlFile, err := os.Open("./ConformanceTests/IIIG006Policy.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	var p pdp.Policy
	s, err := xmlFile.Stat()
	if err != nil {
		fmt.Println("Error stating file:", err)
		return
	}
	size := s.Size()
	b := make([]byte, size)
	xmlFile.Read(b)
	err = xml.Unmarshal([]byte(b), &p)
	fmt.Printf("%+v\n", p)
}
