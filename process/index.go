package process

import (
	"context"
)

type Index struct{}

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
				}
			}
		}
	}`

// Create 创建索引
func (i *Index) Create(ctx context.Context) {
}
