package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// base64
	var encoded = base64.StdEncoding.EncodeToString([]byte("Angga Ari Wijaya"))
	fmt.Println(encoded)
	
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	// csv
	csvString := "angga,ari,wijaya\n" +
	"keenan,evander,alastair\n" +
	"diana,eka,wulandari\n";
	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read();
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"angga", "ari", "wijaya"})
	_ = writer.Write([]string{"keenan", "evander", "alastair"})
	writer.Flush()

}