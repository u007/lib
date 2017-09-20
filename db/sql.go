package db

import (
	"fmt"
	// "github.com/jmoiron/sqlx"
	"github.com/u007/lib/tools"
	// "reflect"
	"strings"
)

const DB_INSERT = 1
const DB_UPDATE = 2

// 1 = insert, 2 = update
func UpdateOrInsert(model map[string]interface{}) (int, error) {
	id := 0
	if val, ok := model["id"]; ok {
		switch val.(type) {
		case int:
			if val.(int) > 0 {
				return DB_UPDATE, nil
			} else {
				return DB_INSERT, nil
			}

		case float64:
			if val.(float64) > 0.0 {
				return DB_UPDATE, nil
			} else {
				return DB_INSERT, nil
			}
		}

		id, _ = tools.ParseIntFromString(val.(string), 0)
		if id > 0 {
			return DB_UPDATE, nil
		} else {
			return DB_INSERT, nil
		}
	}

	if val, ok := model["Id"]; ok {
		switch val.(type) {
		case int:
			if val.(int) > 0 {
				return DB_UPDATE, nil
			} else {
				return DB_INSERT, nil
			}

		case float64:
			if val.(float64) > 0.0 {
				return DB_UPDATE, nil
			} else {
				return DB_INSERT, nil
			}
		}
		id, _ = tools.ParseIntFromString(val.(string), 0)
		if id > 0 {
			return DB_UPDATE, nil
		} else {
			return DB_INSERT, nil
		}
	}

	return -1, fmt.Errorf("id missing")
}

func PrimaryField(model map[string]interface{}) (string, int, error) {
	id := 0
	if val, ok := model["id"]; ok {
		switch val.(type) {
		case int:
			return "id", val.(int), nil

		case float64:
			return "id", int(val.(float64)), nil
		}
		id, _ = tools.ParseIntFromString(val.(string), 0)
		return "id", id, nil
	}

	if val, ok := model["Id"]; ok {
		switch val.(type) {
		case int:
			return "Id", val.(int), nil

		case float64:
			return "Id", int(val.(float64)), nil
		}
		id, _ = tools.ParseIntFromString(val.(string), 0)
		return "Id", id, nil
	}

	return "", -1, fmt.Errorf("id missing")
}

// return sql, params array, error
func SqlxUpdate(model map[string]interface{}, table string, db_type string) (string, []interface{}, error) {
	var fields []string
	// var values_name []string
	var values []interface{}
	//TODO dB type
	field_delimiter := "\""
	primary_field, id, _ := PrimaryField(model)
	// value_delimiter := "'"
	for k, v := range model {
		field := field_delimiter + k + field_delimiter

		if v == nil {
			fields = append(fields, fmt.Sprintf("%s=null", field))
		} else if k != primary_field { //dont update primary field
			switch v.(type) {
			case int:
				values = append(values, v.(int))
			case int64:
				values = append(values, v.(int64))
			case float64:
				values = append(values, v.(float64))
			case string:
				values = append(values, v.(string))
			case bool:
				values = append(values, v.(bool))
			case []interface{}:
				return "", make([]interface{}, 0), fmt.Errorf("unsupported array type of field %s", k)
				// for i, n := range t {
				// 	fmt.Printf("Item: %v= %v\n", i, n)
				// }
			default:
				return "", make([]interface{}, 0), fmt.Errorf("unsupported type of field %s", k)
				// var r = reflect.TypeOf(t)
				// values = append(values, v.(t))
				// fmt.Printf("Other:%v\n", r)
			}
			fields = append(fields, fmt.Sprintf("%s=?", field))
		} //end if

	} //each

	values = append(values, fmt.Sprintf("%d", id))

	sql := fmt.Sprintf("UPDATE %s set %s where %s%s%s=?", table, strings.Join(fields, ","), field_delimiter, primary_field, field_delimiter)
	return sql, values, nil
}

// return sql, params array, error
func SqlxInsert(model map[string]interface{}, table string, db_type string) (string, []interface{}, error) {
	var fields []string
	var values_name []string
	var values []interface{}
	//TODO dB type
	field_delimiter := "\""
	// value_delimiter := "'"
	for k, v := range model {
		field := field_delimiter + k + field_delimiter

		if v == nil {
			// skip this field
			// fields = append(fields, fmt.Sprintf("%s=null", field))
		} else {
			fields = append(fields, field)
			switch v.(type) {
			case int:
				values = append(values, v.(int))
			case int64:
				values = append(values, v.(int64))
			case float64:
				values = append(values, v.(float64))
			case string:
				values = append(values, v.(string))
			case bool:
				values = append(values, v.(bool))
			case []interface{}:
				return "", make([]interface{}, 0), fmt.Errorf("unsupported array type of field %s", k)
				// for i, n := range t {
				// 	fmt.Printf("Item: %v= %v\n", i, n)
				// }
			default:
				return "", make([]interface{}, 0), fmt.Errorf("unsupported type field %s", k)
				// var r = reflect.TypeOf(t)
				// values = append(values, v.(t))
				// fmt.Printf("Other:%v\n", r)
			}
			values_name = append(values_name, "?")
		} // end if

	} //foreach

	sql := fmt.Sprintf("INSERT INTO %s (%s) values (%s)", table, strings.Join(fields, ","), strings.Join(values_name, ","))
	return sql, values, nil
}
