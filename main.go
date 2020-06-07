package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type stringcount struct {
	String             string `json:"String"`
	Numberofcharacters int    `json:"Numberofcharacters"`
}

func (s stringcount) dostuff() {
	count := strconv.Itoa(s.Numberofcharacters)
	fmt.Println("The string is: " + s.String)
	fmt.Println("The number of characters in this string, including whitespace, is: " + count)
}

func main() {
	S := stringcount{}
	Scanner := bufio.NewScanner(os.Stdin)
	quitapp := false
start:
	for quitapp == false {
		fmt.Println("Type q to quit, or anything else to get API response")
		Scanner.Scan()
		query := Scanner.Text()
		if query == "q" {
			quitapp = true
			goto start
		}
		response, err := http.Get("https://stringcount.herokuapp.com/count/" + query)
		if err != nil {
			fmt.Println(err)
		}
		jsonFile, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal([]byte(jsonFile), &S)
		if err != nil {
			fmt.Println(err)
		}
		S.dostuff()
	}
}
