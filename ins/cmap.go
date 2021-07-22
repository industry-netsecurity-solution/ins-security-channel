package ins

import (
	"errors"
	"fmt"
	"strconv"
)

type Map map[interface{}]interface{}

func NewMap() Map {
	d := make(Map)
	return d
}

func ( d Map ) Set(key interface{}, value interface{}) {
	d[key] = value
}

func ( d Map ) Get(key interface{}) interface{} {
	if value, ok := d[key]; ok {
		return value
	}
	return nil
}

func ( d Map ) Has(key interface{}) bool {
	_, ok := d[key]
	return ok
}

func (d Map) AsString(key interface{}) (string, error) {
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

func (d Map) AsInt(key interface{}) (int, error) {
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

func (d Map) AsInt32(key interface{}) (int32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(int32), nil
}

func (d Map) AsInt64(key interface{}) (int64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(int64), nil
}

func (d Map) AsUint(key interface{}) (uint, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint), nil
}

func (d Map) AsUint32(key interface{}) (uint32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint32), nil
}

func (d Map) AsUint64(key interface{}) (uint64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(uint64), nil
}

func (d Map) AsByte(key interface{}) (byte, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(byte), nil
}

func (d Map) AsFloat32(key interface{}) (float32, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(float32), nil
}

func (d Map) AsFloat64(key interface{}) (float64, error) {
	if d[key] == nil {
		return 0, errors.New("Not exist.")
	}

	return d[key].(float64), nil
}

func Map2Map(dst map[interface{}]interface{}, src map[interface{}]interface{}) error {
	for k, v := range src {
		dst[k] = v
	}
	return nil
}
