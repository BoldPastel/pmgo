package pmgo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var dir string

func XmlUnpack(filename string) Permissions {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("okay opened that file")
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var permissions Permissions
	xml.Unmarshal(byteValue, &permissions)
	return permissions
}

func mkdirc() {
	if dir != "" {
		return
	}

	timef := time.Now()
	dir = "perms" + timef.Format("20060102150405") //should see whether perms can be changed based on the function that is calling this.

	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(dir)
}

func WriteXML(xmlStucture interface{}) {

	mkdirc()

	file, _ := xml.MarshalIndent(xmlStucture, "", "\t")
	file = []byte(xml.Header + string(file))

	// If the file doesn't exist, create it, or append to the file; cant be troubled creating dynamic naming
	f, err := os.OpenFile(xmlStucture.(Permissions).Label+".permissionset-meta.xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(file); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
