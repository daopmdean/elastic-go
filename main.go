package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	fmt.Println("----CLOUD_ID---", os.Getenv("CLOUD_ID"))
	fmt.Println("----API_KEY---", os.Getenv("API_KEY"))

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		CloudID: os.Getenv("CLOUD_ID"),
		APIKey:  os.Getenv("API_KEY"),
	})
	if err != nil {
		fmt.Println("err at NewClient", err)
		return
	}

	info, err := client.Info()
	if err != nil {
		fmt.Println("err get client info", err)
		return
	}

	fmt.Println(info)

	// creating an index
	client.Indices.Create("my_index")

	// indexing documents
	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch-daopmdean",
	}
	data, _ := json.Marshal(document)
	client.Index("my_index", bytes.NewReader(data))

	// getting document
	res, err := client.Get("my_index", "id")
	if err != nil {
		fmt.Println("err client.Get", err)
	}
	fmt.Println("---client.Get---")
	fmt.Println(res)
	fmt.Println("---client.Get---")

	// searching documents
	query := `{
		"query": {
			"match_all": {}
		}
	}`
	res, err = client.Search(
		client.Search.WithIndex("my_index"),
		client.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		fmt.Println("err client.Search", err)
	}
	fmt.Println("---client.Search---")
	fmt.Println(res)
	fmt.Println("---client.Search---")

	// updating documents
	client.Update("my_index", "id", strings.NewReader(`{doc: {language: "Gooo"}}`))

	// deleting documents
	client.Delete("my_index", "id")

	// deleting index
	client.Indices.Delete([]string{"my_index"})

}
