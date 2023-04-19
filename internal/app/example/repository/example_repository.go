package repository

import "cocodingding/keyword-recommender/internal/app/example/model"

func GetExampleList() []model.Example {
	examples := []model.Example{{
		SeqNo: 0,
		Name:  "a",
	}, {
		SeqNo: 1,
		Name:  "b",
	}}
	return examples
}
