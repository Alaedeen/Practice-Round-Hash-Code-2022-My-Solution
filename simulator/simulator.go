package simulator

import "github.com/Alaedeen/OnePizza/model"


func CheckScore(clients []model.Client, ingredients []string) int {
	score := 0
	for _,client := range clients {
		if subslice(client.Likes,ingredients) {
			if haveNoElements(client.Dislikes,ingredients) {
				score++
			}
		}
	}

	return score
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func subslice (s1 []string, s2 []string) bool {
    if len(s1) > len(s2) { return false }
    for _, e := range s1 {
        if ! contains(s2,e) {
            return false
        }
    }
    return true
}

func haveNoElements (s1 []string, s2 []string) bool {
    if len(s1) > len(s2) { return false }
    for _, e := range s1 {
        if contains(s2,e) {
            return false
        }
    }
    return true
}