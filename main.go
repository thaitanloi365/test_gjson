package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {
	page1, err := ioutil.ReadFile("page1.json")
	if err != nil {
		panic(err)
	}

	page2, err := ioutil.ReadFile("page2.json")
	if err != nil {
		panic(err)
	}

	page3, err := ioutil.ReadFile("page3.json")
	if err != nil {
		panic(err)
	}

	value1 := gjson.Get(string(page1), "_embedded.categories.#.name")

	value2 := gjson.Get(string(page2), "_embedded.categories.#.name")

	value3 := gjson.Get(string(page3), "_embedded.categories.#.name")

	var value = []string{}
	for _, v := range value1.Array() {
		value = append(value, v.String())
	}
	for _, v := range value2.Array() {
		value = append(value, v.String())
	}
	for _, v := range value3.Array() {
		value = append(value, v.String())
	}

	var m = map[string]bool{}
	for _, v := range value {
		var trimValue = strings.TrimSpace(v)
		if trimValue == "" || trimValue == "null" {
			continue
		}

		m[v] = true
	}

	var finalResult = []string{}
	for key, v := range m {
		if v {
			finalResult = append(finalResult, key)
		}
	}

	fmt.Println(finalResult)
	data, err := json.MarshalIndent(&finalResult, "", "   ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
