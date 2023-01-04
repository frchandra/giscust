package app

import "bitbucket.org/frchandra/giscust/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.Message{}},
	}
}
