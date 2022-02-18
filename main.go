package main

import (
	"fmt"
	"log"
	"os"

	fm "github.com/Alaedeen/OnePizza/filesmanagement"
	"github.com/Alaedeen/OnePizza/simulator"
	"github.com/Alaedeen/OnePizza/solution"
)

func main() {
	names, err := fm.ReadFilesNames()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
	for _, name := range names {
		c, l, d, err := fm.ReadInputFile(name)
		if err != nil {
			log.Fatal(err.Error())
			os.Exit(1)
		}

		pizza := solution.GeneratePizza(c,l, d)
		score := simulator.CheckScore(c, pizza)
		// fm.WriteOutputFile(name,pizza)

		fmt.Println(name,score)
	}

}
