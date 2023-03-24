package repository

import (
	"os"
  model "crudgengui/model"
)


type YAMLModel struct {
  yamlfile string
}

func NewYAMLModel(fname string) *YAMLModel {
  ym:=new(YAMLModel)
  ym.yamlfile=fname
  return ym
}


func (ym YAMLModel) WriteModel(m *model.Model) error {
	file, err := os.Create(ym.yamlfile)
	if err != nil {
		return err
	}
	err = m.WriteYAML(file)
	if err != nil {
		return err
	}
	return nil
}

func (ym *YAMLModel)ReadModel(m *model.Model) (err error) {
	file, err := os.Open(ym.yamlfile)
	if err != nil {
		return err
	}
	err = m.ReadYAML(file)
	if err != nil {
		return err
	}
	return nil
}
