package filesmanagement

import (
	"fmt"
	"log"
	"os"
)

func WriteOutputFile(filename string, pizza []string) {
	var out string = fmt.Sprint(len(pizza))
	for _, ingredient := range pizza {
		out += " " + ingredient
	}
	err := os.WriteFile("./output_files/" + filename + ".out", []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
