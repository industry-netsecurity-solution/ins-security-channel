package ins

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type SMap map[string]interface{}

func NewSMap() SMap {
	d := make(SMap)
	return d
}

func ( d SMap ) Set(key string, value interface{}) {
	d[key] = value
}

func ( d SMap ) Get(key string) interface{} {
	if value, ok := d[key]; ok {
		return value
	}
	return nil
}

func ( d SMap ) Has(key string) bool {
	_, ok := d[key]
	return ok
}

func (d SMap) AsString(key string) (string, error) {
	if d[key] == nil {
		return "", errors.New("Not exist.")
	}

	switch d[key].(type) {
	case bool:
		if d[key].(bool) {
			return "true", nil
		}
		return "false", nil
	case int:
		return strconv.FormatInt(int64(d[key].(int)), 10), nil
	case int8:
		return strconv.FormatInt(int64(d[key].(int8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(d[key].(int16)), 10), nil
	case int32:
		return strconv.FormatInt(int64(d[key].(int32)), 10), nil
	case int64:
		return strconv.FormatInt(int64(d[key].(int64)), 10), nil
	case uint:
		return strconv.FormatUint(uint64(d[key].(uint)), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(d[key].(uint8)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(d[key].(uint16)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(d[key].(uint32)), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(d[key].(uint64)), 10), nil
	case float32:
		return strconv.FormatFloat(float64(d[key].(float32)), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(float64(d[key].(float32)), 'f', -1, 64), nil
	case string:
		return d[key].(string), nil
	//... etc
	default:
		return "", errors.New(fmt.Sprintf("Can not convert type '%s' to 'string'.", GetType(d[key])))
	}

	return d[key].(string), nil
}

func (d SMap) AsInt(key string) (int, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	switch d[key].(type) {
	case bool:
		if d[key].(bool) {
			return 1, nil
		}
		return 0, nil
	case int:
		return d[key].(int), nil
	case int8:
		return int(d[key].(int8)), nil
	case int16:
		return int(d[key].(int16)), nil
	case int32:
		return int(d[key].(int32)), nil
	case int64:
		return int(d[key].(int64)), nil
	case uint:
		return int(d[key].(uint)), nil
	case uint8:
		return int(d[key].(uint8)), nil
	case uint16:
		return int(d[key].(uint16)), nil
	case uint32:
		return int(d[key].(uint32)), nil
	case uint64:
		return int(d[key].(uint64)), nil
	case float32:
		return int(d[key].(float32)), nil
	case float64:
		return int(d[key].(float64)), nil
	case string:
		i, err := strconv.ParseInt(d[key].(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	//... etc
	default:
		return 0, errors.New(fmt.Sprintf("Can not convert type '%s' to 'int'.", GetType(d[key])))
	}

	return d[key].(int), nil
}

func (d SMap) AsInt32(key string) (int32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(int32), nil
}

func (d SMap) AsInt64(key string) (int64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(int64), nil
}

func (d SMap) AsUint(key string) (uint, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint), nil
}

func (d SMap) AsUint32(key string) (uint32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint32), nil
}

func (d SMap) AsUint64(key string) (uint64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint64), nil
}

func (d SMap) AsByte(key string) (byte, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(byte), nil
}

func (d SMap) AsFloat32(key string) (float32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(float32), nil
}

func (d SMap) AsFloat64(key string) (float64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(float64), nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		//return fmt.Errorf("No such field: %s in obj", name)
		return nil
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func SMap2Struct(dst interface{}, src map[string]interface{}) error {
	for k, v := range src {
		err := SetField(dst, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func Map2SMap(dst map[string]interface{}, src map[interface{}]interface{}) error {
	for k, v := range src {
		dst[k.(string)] = v
	}
	return nil
}
