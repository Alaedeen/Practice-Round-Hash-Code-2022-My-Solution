package solution

import (
	"github.com/Alaedeen/OnePizza/model"
	"github.com/Alaedeen/OnePizza/simulator"
)

type Doc []model.Client

var BestScore int = 0
var BestPizza []string

func (a Doc) Len() int           { return len(a) }
func (a Doc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Doc) Less(i, j int) bool { return len(a[i].Dislikes) < len(a[j].Dislikes) }

func GeneratePizza(clients []model.Client, likes model.IngredientsMap, dislikes model.IngredientsMap) []string {
	pizza := []string{}

	//use all possible ingredients

	// for like := range likes {
	// 	if !contains(pizza, like) {
	// 		pizza = append(pizza, like)
	// 	}
	// }
	// for dislike := range dislikes {
	// 	if !contains(pizza, dislike) {
	// 		pizza = append(pizza, dislike)
	// 	}
	// }

	/*
		a_an_example.in.txt 1
		b_basic.in.txt 5
		c_coarse.in.txt 2
		d_difficult.in.txt 1441
		e_elaborate.in.txt 840
	*/

	// add all ingredients with 0 dislikes

	// for l := range likes {
	// 	if dislikes[l]==0 {
	// 		pizza = append(pizza, l)
	// 	}
	// }

	/*
		a_an_example.in.txt 2
		b_basic.in.txt 5
		c_coarse.in.txt 1
		d_difficult.in.txt 1420
		e_elaborate.in.txt 412
	*/

	// add all ingredients with likes > dislikes

	// for name, likeOcc := range likes {
	// 	if likeOcc > dislikes[name] {
	// 		pizza = append(pizza, name)
	// 	}
	// }

	/*
		a_an_example.in.txt 2
		b_basic.in.txt 5
		c_coarse.in.txt 4
		d_difficult.in.txt 1697
		e_elaborate.in.txt 799
	*/

	// sort clients by number of dislikes and choose top 50% with less dislikes

	// sort.Sort(Doc(clients))

	// for i, client := range clients {
	// 	if i > len(clients)/2 {
	// 		break
	// 	}
	// 	for _, like := range client.Likes {
	// 		if !contains(pizza, like) {
	// 			pizza = append(pizza, like)
	// 		}
	// 	}

	// }

	/*
		a_an_example.in.txt 2
		b_basic.in.txt 5
		c_coarse.in.txt 4
		d_difficult.in.txt 1441
		e_elaborate.in.txt 1306
	*/

	// test all possible cases

	allIngredients := []string{}

	for like := range likes {
		if !contains(allIngredients, like) {
			allIngredients = append(allIngredients, like)
		}
	}
	for dislike := range dislikes {
		if !contains(allIngredients, dislike) {
			allIngredients = append(allIngredients, dislike)
		}
	}

	for r := 1; r <= len(allIngredients); r++ {
		data := make([]string, r)

		testAllPossibleCases(clients, allIngredients, data, 0, len(allIngredients)-1, 0, r)

	}
	pizza = BestPizza

	return pizza
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func testAllPossibleCases(c []model.Client, ingredients []string, data []string, start int, end int, index int, r int) {
	if index == r {
		var pizza []string
		for i := 0; i < r; i++ {
			pizza = append(pizza, data[i])
		}
		checkIfBestScore(pizza, simulator.CheckScore(c, pizza))
		return
	}

	for i := start; i <= end && end-i+1 >= r-index; i++ {
		data[index] = ingredients[i]
		testAllPossibleCases(c, ingredients, data, i+1, end, index+1, r)
	}

}

func checkIfBestScore(pizza []string, score int) {
	if score > BestScore {
		BestScore = score
		BestPizza = pizza
	}
}
