package filesmanagement

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/Alaedeen/OnePizza/model"
)

func ReadFilesNames() ([]string, error) {
	files, err := ioutil.ReadDir("./input_files")
	if err != nil {
		return nil, err
	}
	filenames := []string{}
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames, nil
}

func ReadInputFile(filename string) (clients []model.Client,likes model.IngredientsMap, dislikes model.IngredientsMap, err error) {

	likes = make(model.IngredientsMap)
	dislikes = make(model.IngredientsMap)
	clients = []model.Client{}

	f, err := os.Open("./input_files/" + filename)
	if err != nil {
		return nil,nil, nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	line := 1
	var client model.Client
	for scanner.Scan() {
		
		if line%2 == 0 {
			client = model.Client{}
			like := strings.Split(scanner.Text(), " ")
			if like[0] != "0" {
				for l, _ := strconv.Atoi(like[0]); l >= 1; l-- {
					likes[like[l]]++
					client.Likes = append(client.Likes, like[l])
				}
			}
		} else if line%2 != 0 && line != 1 {
			dislike := strings.Split(scanner.Text(), " ")
			if dislike[0] != "0" {
				for l, _ := strconv.Atoi(dislike[0]); l >= 1; l-- {
					dislikes[dislike[l]]++
					client.Dislikes = append(client.Dislikes, dislike[l])
				}
			}
			clients = append(clients, client)
		}
		
		line++
	}

	return  clients, likes, dislikes, nil
}
