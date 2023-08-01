package repository

import (
	"crudgengui/model"
	"fmt"
	"log"
	"strings"
)

// ModelReaderWriter must be implemented and injected in the creation of 'ModelRepository' is providing read and write functionalities
type ModelReaderWriter interface {
	WriteModel(m *model.Model) error
	ReadModel(m *model.Model) error
}

// ModelRepository provides all the data manipulation logic to the model
type ModelRepository struct {
	modelRW ModelReaderWriter
  m *model.Model
}

// NewModelRepository createas an new instance of ModelRepository with injected persistence functionality
func NewModelRepository(mrw ModelReaderWriter) *ModelRepository {
	mrep := new(ModelRepository)
	mrep.modelRW = mrw
  mrep.m = model.NewModel()
  mrep.modelRW.ReadModel(mrep.m)
	return mrep
}

// GetModel returns the model
func (mrep *ModelRepository)GetModel() *model.Model {
	return mrep.m
}

// SaveOrUpdateEntity saves or updates an entity in the model
func (mrep *ModelRepository) SaveOrUpdateEntity(e *model.Entity) error {
	mrep.m.Entities[strings.ToLower(e.Name)] = *e

	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteEntity deletes one entity from the model
func (mrep *ModelRepository) DeleteEntity(name string) error {
  _, ok := mrep.GetEntity(strings.ToLower(name))
	if !ok {
		return fmt.Errorf("Entity '%s' not found", name)
	}
  if mrep.m.EntityInRealtions(name) {
    return fmt.Errorf("Cannot delete as Entity '%s' is linked in a relation", name)
  }
  
	delete(mrep.m.Entities, strings.ToLower(name))

	return mrep.modelRW.WriteModel(mrep.m)
}

// GetAllEntities gets all entities from the model
func (mrep ModelRepository) GetAllEntities() map[string]model.Entity {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil
	}
	return mrep.m.Entities
}

// GetEntity gets one single entity from the model
func (mrep *ModelRepository) GetEntity(name string) (*model.Entity, bool) {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil, false
	}
	r, ok := mrep.m.Entities[strings.ToLower(name)]
	return &r, ok
}

// SaveOrUpdateField saves or updates a field in the model
func (mrep *ModelRepository) SaveOrUpdateField(ename string, f *model.Field) error {
	e, ok := mrep.GetEntity(strings.ToLower(ename))
	if !ok {
		return fmt.Errorf("Entity '%s' not found", ename)
	}
	e.Fields[strings.ToLower(f.Name)] = *f

	mrep.m.Entities[strings.ToLower(ename)] = *e

	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteField deletes one field from the model
func (mrep *ModelRepository) DeleteField(ename string, fname string) error {
	e, ok := mrep.GetEntity(strings.ToLower(ename))
	if !ok {
		return fmt.Errorf("Entity '%s' not found", ename)
	}

	delete(e.Fields, strings.ToLower(fname))
	mrep.m.Entities[strings.ToLower(ename)] = *e

	return mrep.modelRW.WriteModel(mrep.m)
}

// GetAllFields gets all fields from the model
func (mrep *ModelRepository) GetAllFields(ename string) (map[string]model.Field, error) {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil, err
	}
	e, ok := mrep.GetEntity(strings.ToLower(ename))
	if !ok {
		return nil, fmt.Errorf("Entity '%s' not found", ename)
	}
	return e.Fields, nil
}

// SaveOrUpdateRelation saves or updates a relation in the model
func (mrep *ModelRepository) SaveOrUpdateRelation(rname string, r *model.Relation) error {

	_, ok := mrep.GetRelation(strings.ToLower(rname))
	if !ok {
    log.Println("Create new relation with name: ",strings.ToLower(rname))		
	}

	mrep.m.Relations[strings.ToLower(rname)] = *r

	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteRelation deletes one relation from the model
func (mrep *ModelRepository) DeleteRelation(name string) error {
	delete(mrep.m.Relations, strings.ToLower(name))

	return mrep.modelRW.WriteModel(mrep.m)
}

// GetAllRelations gets all relations from the model
func (mrep *ModelRepository) GetAllRelations() map[string]model.Relation {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil
	}

	return mrep.m.Relations
}

// GetRelation gets one single relation from the model
func (mrep *ModelRepository) GetRelation(name string) (*model.Relation, bool) {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil, false
	}
	r, ok := mrep.m.Relations[strings.ToLower(name)]
	return &r, ok
}

// GetField of an Entity from the model
func (mrep *ModelRepository) GetField(entityname string, name string) (*model.Field, bool) {
  var ent model.Entity
  var field model.Field
  var ok bool
  
  if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil, false
	}

	if ent, ok = mrep.m.Entities[strings.ToLower(entityname)]; !ok {
    return nil,false
  }

  field, ok = ent.Fields[strings.ToLower(name)]
	return &field, ok
}
