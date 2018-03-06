package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func main() {

	f, err := os.Open(path.Join(path.Dir(os.Args[0]), "../data/iris.cvs"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	//rawCSVData, err := reader.ReadAll()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(rawCSVData)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			log.Println("reached EOF")
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(record)
	}

}
