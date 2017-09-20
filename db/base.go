package db

import (
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/structs"
	"github.com/u007/lib/tools"
	"gopkg.in/gorp.v1"
)

var models = make(map[string]ModelInfo)

type ModelInfo struct {
	Dates  map[string]int
	Fields []string
}

type ModelDeletable interface {
	MarkedDeleted(time.Time)
}

//cache current fields
var current_fields []string

func HasDeletedAt(model interface{}) (bool, error) {
	i, err := ModelDeletedAt(model)
	if err != nil {
		return false, err
	}
	if i >= 0 {
		return true, nil
	}
	return false, nil
}

func ModelMarkDelete(model interface{}) error {
	now := time.Now().UTC()
	i, err := ModelDeletedAt(model)
	if err != nil {
		return err
	}
	if i >= 0 {
		ps := reflect.ValueOf(&model)
		field := ps.Elem().Field(i)
		field.Set(reflect.ValueOf(now))
	}
	return nil
}

func ModelMarkUpdate(model interface{}) error {
	now := time.Now().UTC()
	i, err := ModelUpdatedAt(model)
	if err != nil {
		return err
	}
	if i >= 0 {
		ps := reflect.ValueOf(&model)
		field := ps.Elem().Field(i)
		field.Set(reflect.ValueOf(now))
	}
	return nil
}

func ModelMarkCreate(model interface{}) error {
	now := time.Now().UTC()
	i, err := ModelCreatedAt(model)
	if err != nil {
		return err
	}
	if i >= 0 {
		ps := reflect.ValueOf(&model)
		field := ps.Elem().Field(i)
		field.Set(reflect.ValueOf(now))
	}

	i, err = ModelUpdatedAt(model)
	if err != nil {
		return err
	}
	if i >= 0 {
		ps := reflect.ValueOf(&model)
		field := ps.Elem().Field(i)
		field.Set(reflect.ValueOf(now))
	}
	return nil
}

func ModelCreatedAt(model interface{}) (int, error) {
	struct_name := structs.Name(&model)
	if m, ok := models[struct_name]; ok {
		return m.Dates["created"], nil // already declare
	}
	return -1, fmt.Errorf("Model not loaded")
}

func ModelUpdatedAt(model interface{}) (int, error) {
	struct_name := structs.Name(&model)
	if m, ok := models[struct_name]; ok {
		return m.Dates["updated"], nil // already declare
	}
	return -1, fmt.Errorf("Model not loaded")
}

func ModelDeletedAt(model interface{}) (int, error) {
	struct_name := structs.Name(&model)
	if m, ok := models[struct_name]; ok {
		return m.Dates["updated"], nil // already declare
	}
	return -1, fmt.Errorf("Model not loaded")
}

func InitializeTableAdv(gorpdb *gorp.DbMap, model interface{}, table string, primary_field string) error {
	struct_name := structs.Name(model)
	if _, ok := models[struct_name]; ok {
		return nil // already declare
	}

	tbl := gorpdb.AddTableWithName(model, table)
	if primary_field != "" {
		tbl.SetKeys(true, primary_field)
	}
	if len(current_fields) == 0 {
		current_fields = structs.Names(&model)
	}

	model_ob := &ModelInfo{Dates: make(map[string]int)}
	//fields of all deleted fields
	model_ob.Dates["created"] = int(tools.IndexOf(current_fields, "Created_at"))
	model_ob.Dates["updated"] = int(tools.IndexOf(current_fields, "Updated_at"))
	model_ob.Dates["deleted"] = tools.IndexOf(current_fields, "Deleted_at")
	model_ob.Fields = current_fields

	models[struct_name] = *model_ob
	return nil
}

func InitializeTable(gorpdb *gorp.DbMap, model interface{}, table string) error {
	current_fields = structs.Names(model)
	primary_field := ""

	if tools.InStringArray(current_fields, "Id") {
		primary_field = "Id"
	}
	return InitializeTableAdv(gorpdb, model, table, primary_field)
}
