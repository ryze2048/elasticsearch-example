package process

import "time"

var indexMapping = `{
		"mappings": {
			"properties": {
				"title": {
					"type": "text"
				},
				"content": {
					"type": "text"
				},
				"timestamp": {
					"type": "date"
				},
				"intro": {
					"type": "text"
				}
			}
		}
	}`

var IndexName = "article"

type Article struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Intro     string    `json:"intro"`
}
