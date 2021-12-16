package main
import (
	"encoding/xml"
	"fmt"
)
/*
   subject : XML
*/

//Plant will be mapped to XML. Similarly to the JSON examples, field tags contain directives for the encoder and decoder. Here we use some special features of the XML package: the XMLName field name dictates the name of the XML element representing this struct; id,attr means that the Id field is an XML attribute rather than a nested element.

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}



func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}



func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	//coffee := &Plant{Id: 27, Name: "Coffee", Origin: []string{"Ethiopia", "Brazil"}}
	
	//Emit XML representing our plant; using MarshalIndent to produce a more human-readable output.
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	fmt.Println()



	//To add a generic XML header to the output, append it explicitly.
	fmt.Println(xml.Header + string(out))
	fmt.Println()



	//Use Unmarhshal to parse a stream of bytes with XML into a data structure. If the XML is malformed or cannot be mapped onto Plant, a descriptive error will be returned.
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)//Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Printf("%v\n", p)//Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Printf("%+v\n", p)//Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
	fmt.Printf("%#v\n", p)//main.Plant{XMLName:xml.Name{Space:"", Local:"plant"}, Id:27, Name:"Coffee", Origin:[]string{"Ethiopia", "Brazil"}}
	
	fmt.Println()


	
	//The parent>child>plant field tag tells the encoder to nest all plants under <parent><child>...
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
// -*- mode: compilation; default-directory: "~/go/src/practice/09json/" -*-
// Compilation started at Sat Oct  9 23:15:33
//  
// go run xml01.go
//  <plant id="27">
//    <name>Coffee</name>
//    <origin>Ethiopia</origin>
//    <origin>Brazil</origin>
//  </plant>
//  
// <?xml version="1.0" encoding="UTF-8"?>
//  <plant id="27">
//    <name>Coffee</name>
//    <origin>Ethiopia</origin>
//    <origin>Brazil</origin>
//  </plant>
//  
// Plant id=27, name=Coffee, origin=[Ethiopia Brazil]
//  
//  <nesting>
//    <parent>
//      <child>
//        <plant id="27">
//          <name>Coffee</name>
//          <origin>Ethiopia</origin>
//          <origin>Brazil</origin>
//        </plant>
//        <plant id="81">
//          <name>Tomato</name>
//          <origin>Mexico</origin>
//          <origin>California</origin>
//        </plant>
//      </child>
//    </parent>
//  </nesting>
//  
// Compilation finished at Sat Oct  9 23:15:34
