package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hoehwa/gopkg/bubbletea/listsimple"
)

func readResBody(res *http.Response) []byte {
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	return body
}

func SearchTopRepos(q Query) {
	data := JSONData{}
	var formattedData []string

	body := readResBody(q.fetchReposDesc())
	json.Unmarshal(body, &data)

	for _, i := range data.Items {
		formattedData = append(formattedData,
			fmt.Sprintf("%s, %d stars", i.FullName, i.StargazersCount))
	}

	listsimple.BrowseInit(formattedData)
}
