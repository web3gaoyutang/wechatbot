package handlers

import "testing"

func TestA(t *testing.T) {
	a := UserMessageHandler{
		info: map[string]interface{}{
			"user_id": "123",
			"year":    1999,
			"month":   5,
			"day":     29,
			"hour":    20,
			"min":     30,
			"gender":  "男",
		},
	}
	create_gpt(a.info)
}
