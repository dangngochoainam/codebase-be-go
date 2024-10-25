package copyhepler

import (
	"github.com/jinzhu/copier"
	"log"
)

type (
	ModelConverter interface {
		FromModel(to interface{}, from interface{})
		ToModel(to interface{}, from interface{})
	}
	modelConverter struct{}
)

func NewModelConverter() ModelConverter {
	return &modelConverter{}
}

func (m *modelConverter) FromModel(to interface{}, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		log.Fatalln("failed to copy from model: ", err)
	}
}

func (m *modelConverter) ToModel(to interface{}, from interface{}) {
	err := copier.Copy(to, from)
	if err != nil {
		log.Fatalln("failed to copy to model: ", err)
	}
}
