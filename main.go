package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
)

func main() {
	//read file
	content, _ := ioutil.ReadFile("reader.csv")

	reader := csv.NewReader(bytes.NewBuffer(content))

	for {
		row, err := reader.Read()
		if err != nil {
			// any error breaks the loop. EOF counted as error
			fmt.Println(err)
			break
		}
		fmt.Println(row)
	}

	//Write file to byte and string
	empData := [][]string{
		{"Id", "Name", "Amount"},
		{"39a4b556-e2fa-4885-845c-8b937f753885", "Ezza", "105000"},
		{"f8c0f279-eedd-44e6-83c5-1f33d5661f46", "Arif", "900000"},
		{"314f34ec-f282-4f80-8d91-defeff6149f4", "Galih", "690000"},
	}

	var bufferFile bytes.Buffer

	writePerLine := csv.NewWriter(&bufferFile)

	for _, empRow := range empData {
		err := writePerLine.Write(empRow)
		if err != nil {
			fmt.Println(err)
		}
	}

	//must flush first before send csv
	writePerLine.Flush()
	fmt.Println("BUFFER", bufferFile)
	fmt.Println("BYTES", bufferFile.Bytes())
	fmt.Println("STRING", bufferFile.String())

	//write all with [][]string
	writeAll := csv.NewWriter(&bufferFile)
	err := writeAll.WriteAll(empData)
	if err != nil {
		fmt.Println(err)
	}
	writeAll.Flush()
	fmt.Println("BUFFER", bufferFile)
	fmt.Println("BYTES", bufferFile.Bytes())
	fmt.Println("STRING", bufferFile.String())

}
