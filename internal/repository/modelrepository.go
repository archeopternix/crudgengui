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
	m       *model.Model
}

// NewModelRepository createas an new instance of ModelRepository with injected persistence functionality
func NewModelRepository(mrw ModelReaderWriter) *ModelRepository {
	mrep := &ModelRepository{
		modelRW: mrw,
		m:       model.NewModel(),
	}
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		log.Panic(err)
	}
	return mrep
}

// GetModel returns the model
func (mrep *ModelRepository) GetModel() *model.Model {
	return mrep.m
}

// SaveModel saves the model
func (mrep *ModelRepository) SaveModel(name string, settings model.Settings) error {
	mrep.m.Settings = settings
	mrep.m.Name = name
	return mrep.modelRW.WriteModel(mrep.m)
}

// SaveOrUpdateEntity saves or updates an entity in the model
func (mrep *ModelRepository) SaveOrUpdateEntity(e *model.Entity) error {
	name := strings.ToLower(e.Name)
	mrep.m.Entities[name] = *e
	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteEntity deletes one entity from the model
func (mrep *ModelRepository) DeleteEntity(name string) error {
	name = strings.ToLower(name)
	if _, ok := mrep.m.Entities[name]; !ok {
		return fmt.Errorf("entity '%s' not found", name)
	}
	if mrep.m.EntityInRealtions(name) {
		return fmt.Errorf("cannot delete as entity '%s' is linked in a relation", name)
	}
	delete(mrep.m.Entities, name)
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
	ename = strings.ToLower(ename)
	e, ok := mrep.GetEntity(ename)
	if !ok {
		return fmt.Errorf("entity '%s' not found", ename)
	}
	idx := e.GetFieldIndexByName(f.Name)
	if idx == -1 {
		e.Fields = append(e.Fields, *f)
	} else {
		e.Fields[idx] = *f
	}
	mrep.m.Entities[ename] = *e
	return mrep.modelRW.WriteModel(mrep.m)

}

// DeleteField deletes one field from the model
func (mrep *ModelRepository) DeleteField(ename string, fname string) error {
	e, ok := mrep.GetEntity(strings.ToLower(ename))
	if !ok {
		return fmt.Errorf("Entity '%s' not found", ename)
	}
	e.DeleteFieldByName(fname)
	mrep.m.Entities[strings.ToLower(ename)] = *e

	return mrep.modelRW.WriteModel(mrep.m)
}

// GetAllFields gets all fields from the model
func (mrep *ModelRepository) GetAllFields(ename string) ([]model.Field, error) {
	ename = strings.ToLower(ename)
	e, ok := mrep.GetEntity(ename)
	if !ok {
		return nil, fmt.Errorf("entity '%s' not found", ename)
	}
	return e.Fields, nil
}

// SaveOrUpdateRelation saves or updates a relation in the model
func (mrep *ModelRepository) SaveOrUpdateRelation(rname string, r *model.Relation) error {
	rname = strings.ToLower(rname)
	mrep.m.Relations[rname] = *r
	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteRelation deletes one relation from the model
func (mrep *ModelRepository) DeleteRelation(name string) error {
	name = strings.ToLower(name)
	delete(mrep.m.Relations, name)
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
	name = strings.ToLower(name)
	r, ok := mrep.m.Relations[name]
	return &r, ok
}

// GetField of an Entity from the model
func (mrep *ModelRepository) GetField(entityname string, name string) (*model.Field, bool) {
	entityname = strings.ToLower(entityname)
	ent, ok := mrep.m.Entities[entityname]
	if !ok {
		return nil, false
	}
	field := ent.GetFieldByName(name)
	return field, field != nil
}

// GetAllLookups gets all entities from the model
func (mrep ModelRepository) GetAllLookups() map[string]model.Lookup {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil
	}
	return mrep.m.Lookups
}

// GetLookup gets one single entity from the model
func (mrep *ModelRepository) GetLookup(name string) (*model.Lookup, bool) {
	name = strings.ToLower(name)
	r, ok := mrep.m.Lookups[name]
	return &r, ok
}

// SaveOrUpdateLookup saves or updates a relation in the model
func (mrep *ModelRepository) SaveOrUpdateLookup(name string, lookup *model.Lookup) error {
	name = strings.ToLower(name)
	mrep.m.Lookups[name] = *lookup
	return mrep.modelRW.WriteModel(mrep.m)
}

// DeleteLookup deletes one relation from the model
func (mrep *ModelRepository) DeleteLookup(name string) error {
	name = strings.ToLower(name)
	delete(mrep.m.Lookups, name)
	return mrep.modelRW.WriteModel(mrep.m)
}

// GetAllLookupNames gets all names of lookups from the model
func (mrep ModelRepository) GetAllLookupNames() (names []string) {
	if err := mrep.modelRW.ReadModel(mrep.m); err != nil {
		return nil
	}
	for key, _ := range mrep.m.Lookups {
		names = append(names, key)
	}
	return names
}

// GetAllLookupNames gets all names of lookups from the model
func (mrep ModelRepository) StartGeneration() error {
	copy, err := mrep.m.ParseDependencies()
	if err != nil {
		fmt.Println("Fehler beim Parsen des Models")
	}
	model.WriteToFile(copy, "../../data/generated.yaml")

	generator := NewGenerator()

	basepath := "../../"
	modules := []string{
		basepath + "modules/application/app.yaml",
		basepath + "modules/model/models.yaml",
		basepath + "modules/mockdatabase/mockdatabase.yaml",
		basepath + "modules/view/view.yaml",
	}
	if err := generator.AddModules(modules...); err != nil {
		fmt.Println("Fehler beim Hinzufügen des Models: ", err)
	}
	if err := generator.GenerateAll(copy, "/Users/Andreas Eisner/go/src"); err != nil {
		fmt.Println("Fehler beim Hinzufügen des Models: ", err)
	}

	fmt.Println("Model geladen")
	return nil
}
