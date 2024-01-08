package github

import (
	"fmt"
	"log"
	"net/http"
)

type Query struct {
	keyword string

	// optional fields
	language string
	topics   string
}

func NewQuery(keyword string) Query {
	sq := Query{}
	return sq
}

func (q Query) WithLanguage(language string) Query {
	q.language = fmt.Sprintf("+language=%s", language)
	return q
}

func (q Query) WithTopic(topic string) Query {
	q.topics = fmt.Sprintf("+topics=%s", topic)
	return q
}

func (q Query) fetchReposDesc() *http.Response {
	SearchAPIendpoint := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=%s%s%s+stars:>=0&sort=stars&order=desc",
		q.keyword, q.language, q.topics)

	res, err := http.Get(SearchAPIendpoint)
	if err != nil {
		log.Fatal(err)
	}

	return res
}
