package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bytedance/sonic"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	m := map[string]interface{}{
		"name": "z3",
		"age":  20,
		"desc": "我来自CD",
	}

	movie := Movie{Title: "给爸爸的信", Year: 1998, Color: true, Actors: []string{"李连杰", "梅艳芳"}}

	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "给爸爸的信", Year: 1998, Color: true,
			Actors: []string{"李连杰", "梅艳芳"}},
	}

	// sonic序列化
	byt, err := sonic.Marshal(&m)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("json: %+v\n", string(byt))

	// sonic反序列化
	um := make(map[string]interface{})
	err = sonic.Unmarshal(byt, &um)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("unjson: %+v\n", um)

	// 使用encoding/json库
	// Marshal返回一个编码后的字节slice
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 格式化输出，第二个参数是前缀，第三个参数是缩进
	data2, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	data, err = json.Marshal(&movie)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	byt, err = sonic.Marshal(&movies)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("json: %+v\n", string(byt))
}
