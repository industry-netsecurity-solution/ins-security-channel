package ins

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

func NativeEndian() binary.ByteOrder {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xEFFE)

	var nativeEndian binary.ByteOrder
	switch buf {
	case [2]byte{0xFE, 0xEF}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xEF, 0xFE}:
		nativeEndian = binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}

	return nativeEndian
}

func AsBytes(value interface{}) ([]byte, error) {
	if value == nil {
		return nil, errors.New("Can not convert type 'nil' to '[]byte'.")
	}

	buf := [16]byte{}

	switch value.(type) {
	case bool:
		*(*bool)(unsafe.Pointer(&buf[0])) = value.(bool)
		return buf[:1], nil
	case int:
		sz := unsafe.Sizeof(value.(int))
		*(*int)(unsafe.Pointer(&buf[0])) = value.(int)
		return buf[:sz], nil
	case int8:
		sz := unsafe.Sizeof(value.(int8))
		*(*int8)(unsafe.Pointer(&buf[0])) = value.(int8)
		return buf[:sz], nil
	case int16:
		sz := unsafe.Sizeof(value.(int16))
		*(*int16)(unsafe.Pointer(&buf[0])) = value.(int16)
		return buf[:sz], nil
	case int32:
		sz := unsafe.Sizeof(value.(int32))
		*(*int32)(unsafe.Pointer(&buf[0])) = value.(int32)
		return buf[:sz], nil
	case int64:
		sz := unsafe.Sizeof(value.(int64))
		*(*int64)(unsafe.Pointer(&buf[0])) = value.(int64)
		return buf[:sz], nil
	case uint:
		sz := unsafe.Sizeof(value.(uint))
		*(*uint)(unsafe.Pointer(&buf[0])) = value.(uint)
		return buf[:sz], nil
	case uint8:
		sz := unsafe.Sizeof(value.(uint8))
		*(*uint8)(unsafe.Pointer(&buf[0])) = value.(uint8)
		return buf[:sz], nil
	case uint16:
		sz := unsafe.Sizeof(value.(uint16))
		*(*uint16)(unsafe.Pointer(&buf[0])) = value.(uint16)
		return buf[:sz], nil
	case uint32:
		sz := unsafe.Sizeof(value.(uint32))
		*(*uint32)(unsafe.Pointer(&buf[0])) = value.(uint32)
		return buf[:sz], nil
	case uint64:
		sz := unsafe.Sizeof(value.(uint64))
		*(*uint64)(unsafe.Pointer(&buf[0])) = value.(uint64)
		return buf[:sz], nil
	case float32:
		sz := unsafe.Sizeof(value.(float32))
		*(*float32)(unsafe.Pointer(&buf[0])) = value.(float32)
		return buf[:sz], nil
	case float64:
		sz := unsafe.Sizeof(value.(float64))
		*(*float64)(unsafe.Pointer(&buf[0])) = value.(float64)
		return buf[:sz], nil
	case string:
		return []byte(value.(string)), nil
	case []byte:
		return value.([]byte), nil
	//... etc
	default:
		return nil, errors.New(fmt.Sprintf("Can not convert type '%s' to 'string'.", GetType(value)))
	}

	return value.([]byte), nil
}

func AsString(value interface{}) (string, error) {
	if value == nil {
		return "", errors.New("Can not convert type 'nil' to 'string'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return "true", nil
		}
		return "false", nil
	case int:
		return strconv.FormatInt(int64(value.(int)), 10), nil
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10), nil
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10), nil
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10), nil
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10), nil
	case uint:
		return strconv.FormatUint(uint64(value.(uint)), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(value.(uint64)), 10), nil
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(float64(value.(float64)), 'f', -1, 64), nil
	case string:
		return value.(string), nil
	case []byte:
		return string(value.([]byte)), nil
	//... etc
	default:
		return "", errors.New(fmt.Sprintf("Can not convert type '%s' to 'string'.", GetType(value)))
	}

	return value.(string), nil
}

func AsInt(value interface{}) (int, error) {
	if value == nil {
		return 0, errors.New("Can not convert type 'nil' to 'int'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case int:
		return value.(int), nil
	case int8:
		return int(value.(int8)), nil
	case int16:
		return int(value.(int16)), nil
	case int32:
		return int(value.(int32)), nil
	case int64:
		return int(value.(int64)), nil
	case uint:
		return int(value.(uint)), nil
	case uint8:
		return int(value.(uint8)), nil
	case uint16:
		return int(value.(uint16)), nil
	case uint32:
		return int(value.(uint32)), nil
	case uint64:
		return int(value.(uint64)), nil
	case float32:
		return int(value.(float32)), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		i, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	//... etc
	default:
		return 0, errors.New(fmt.Sprintf("Can not convert type '%s' to 'int'.", GetType(value)))
	}

	return value.(int), nil
}

func AsInt16(value interface{}) (int16, error) {
	if value == nil {
		return int16(0), errors.New("Can not convert type 'nil' to 'int16'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return int16(1), nil
		}
		return int16(0), nil
	case int:
		return int16(value.(int)), nil
	case int8:
		return int16(value.(int8)), nil
	case int16:
		return int16(value.(int16)), nil
	case int32:
		return int16(value.(int32)), nil
	case int64:
		return int16(value.(int64)), nil
	case uint:
		return int16(value.(uint)), nil
	case uint8:
		return int16(value.(uint8)), nil
	case uint16:
		return int16(value.(uint16)), nil
	case uint32:
		return int16(value.(uint32)), nil
	case uint64:
		return int16(value.(uint64)), nil
	case float32:
		return int16(value.(float32)), nil
	case float64:
		return int16(value.(float64)), nil
	case string:
		i, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return int16(0), err
		}
		return int16(i), nil
	//... etc
	default:
		return int16(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'int16'.", GetType(value)))
	}

	return value.(int16), nil
}

func AsInt32(value interface{}) (int32, error) {
	if value == nil {
		return int32(0), errors.New("Can not convert type 'nil' to 'int32'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return int32(1), nil
		}
		return int32(0), nil
	case int:
		return int32(value.(int)), nil
	case int8:
		return int32(value.(int8)), nil
	case int16:
		return int32(value.(int16)), nil
	case int32:
		return int32(value.(int32)), nil
	case int64:
		return int32(value.(int64)), nil
	case uint:
		return int32(value.(uint)), nil
	case uint8:
		return int32(value.(uint8)), nil
	case uint16:
		return int32(value.(uint16)), nil
	case uint32:
		return int32(value.(uint32)), nil
	case uint64:
		return int32(value.(uint64)), nil
	case float32:
		return int32(value.(float32)), nil
	case float64:
		return int32(value.(float64)), nil
	case string:
		i, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return int32(0), err
		}
		return int32(i), nil
	//... etc
	default:
		return int32(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'int32'.", GetType(value)))
	}

	return value.(int32), nil
}

func AsInt64(value interface{}) (int64, error) {
	if value == nil {
		return int64(0), errors.New("Can not convert type 'nil' to 'int64'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return int64(1), nil
		}
		return int64(0), nil
	case int:
		return int64(value.(int)), nil
	case int8:
		return int64(value.(int8)), nil
	case int16:
		return int64(value.(int16)), nil
	case int32:
		return int64(value.(int32)), nil
	case int64:
		return int64(value.(int64)), nil
	case uint:
		return int64(value.(uint)), nil
	case uint8:
		return int64(value.(uint8)), nil
	case uint16:
		return int64(value.(uint16)), nil
	case uint32:
		return int64(value.(uint32)), nil
	case uint64:
		return int64(value.(uint64)), nil
	case float32:
		return int64(value.(float32)), nil
	case float64:
		return int64(value.(float64)), nil
	case string:
		i, err := strconv.ParseInt(value.(string), 10, 64)
		if err != nil {
			return int64(0), err
		}
		return int64(i), nil
	//... etc
	default:
		return int64(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'int64'.", GetType(value)))
	}

	return value.(int64), nil
}

func AsUint(value interface{}) (uint, error) {
	if value == nil {
		return uint(0), errors.New("Can not convert type 'nil' to 'uint'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return uint(1), nil
		}
		return uint(0), nil
	case int:
		return uint(value.(int)), nil
	case int8:
		return uint(value.(int8)), nil
	case int16:
		return uint(value.(int16)), nil
	case int32:
		return uint(value.(int32)), nil
	case int64:
		return uint(value.(int64)), nil
	case uint:
		return uint(value.(uint)), nil
	case uint8:
		return uint(value.(uint8)), nil
	case uint16:
		return uint(value.(uint16)), nil
	case uint32:
		return uint(value.(uint32)), nil
	case uint64:
		return uint(value.(uint64)), nil
	case float32:
		return uint(value.(float32)), nil
	case float64:
		return uint(value.(float64)), nil
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return uint(0), err
		}
		return uint(i), nil
	//... etc
	default:
		return uint(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'uint'.", GetType(value)))
	}

	return value.(uint), nil
}

func AsUint16(value interface{}) (uint16, error) {
	if value == nil {
		return uint16(0), errors.New("Can not convert type 'nil' to 'uint16'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return uint16(1), nil
		}
		return uint16(0), nil
	case int:
		return uint16(value.(int)), nil
	case int8:
		return uint16(value.(int8)), nil
	case int16:
		return uint16(value.(int16)), nil
	case int32:
		return uint16(value.(int32)), nil
	case int64:
		return uint16(value.(int64)), nil
	case uint:
		return uint16(value.(uint)), nil
	case uint8:
		return uint16(value.(uint8)), nil
	case uint16:
		return uint16(value.(uint16)), nil
	case uint32:
		return uint16(value.(uint32)), nil
	case uint64:
		return uint16(value.(uint64)), nil
	case float32:
		return uint16(value.(float32)), nil
	case float64:
		return uint16(value.(float64)), nil
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return uint16(0), err
		}
		return uint16(i), nil
	//... uint32
	default:
		return uint16(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'uint16'.", GetType(value)))
	}

	return value.(uint16), nil
}

func AsUint32(value interface{}) (uint32, error) {
	if value == nil {
		return uint32(0), errors.New("Can not convert type 'nil' to 'uint32'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return uint32(1), nil
		}
		return uint32(0), nil
	case int:
		return uint32(value.(int)), nil
	case int8:
		return uint32(value.(int8)), nil
	case int16:
		return uint32(value.(int16)), nil
	case int32:
		return uint32(value.(int32)), nil
	case int64:
		return uint32(value.(int64)), nil
	case uint:
		return uint32(value.(uint)), nil
	case uint8:
		return uint32(value.(uint8)), nil
	case uint16:
		return uint32(value.(uint16)), nil
	case uint32:
		return uint32(value.(uint32)), nil
	case uint64:
		return uint32(value.(uint64)), nil
	case float32:
		return uint32(value.(float32)), nil
	case float64:
		return uint32(value.(float64)), nil
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return uint32(0), err
		}
		return uint32(i), nil
	//... uint32
	default:
		return uint32(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'uint32'.", GetType(value)))
	}

	return value.(uint32), nil
}

func AsUint64(value interface{}) (uint64, error) {
	if value == nil {
		return uint64(0), errors.New("Can not convert type 'nil' to 'uint64'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return uint64(1), nil
		}
		return uint64(0), nil
	case int:
		return uint64(value.(int)), nil
	case int8:
		return uint64(value.(int8)), nil
	case int16:
		return uint64(value.(int16)), nil
	case int32:
		return uint64(value.(int32)), nil
	case int64:
		return uint64(value.(int64)), nil
	case uint:
		return uint64(value.(uint)), nil
	case uint8:
		return uint64(value.(uint8)), nil
	case uint16:
		return uint64(value.(uint16)), nil
	case uint32:
		return uint64(value.(uint32)), nil
	case uint64:
		return uint64(value.(uint64)), nil
	case float32:
		return uint64(value.(float32)), nil
	case float64:
		return uint64(value.(float64)), nil
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return uint64(0), err
		}
		return uint64(i), nil
	//... etc
	default:
		return uint64(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'uint64'.", GetType(value)))
	}

	return value.(uint64), nil
}

func AsByte(value interface{}) (byte, error) {
	if value == nil {
		return byte(0), errors.New("Can not convert type 'nil' to 'byte'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return byte(1), nil
		}
		return byte(0), nil
	case int:
		return byte(value.(int)), nil
	case int8:
		return byte(value.(int8)), nil
	case int16:
		return byte(value.(int16)), nil
	case int32:
		return byte(value.(int32)), nil
	case int64:
		return byte(value.(int64)), nil
	case uint:
		return byte(value.(uint)), nil
	case uint8:
		return byte(value.(uint8)), nil
	case uint16:
		return byte(value.(uint16)), nil
	case uint32:
		return byte(value.(uint32)), nil
	case uint64:
		return byte(value.(uint64)), nil
	case float32:
		return byte(value.(float32)), nil
	case float64:
		return byte(value.(float64)), nil
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err != nil {
			return byte(0), err
		}
		return byte(i), nil
	//... etc
	default:
		return byte(0), errors.New(fmt.Sprintf("Can not convert type '%s' to 'byte'.", GetType(value)))
	}

	return value.(byte), nil
}

func AsFloat32(value interface{}) (float32, error) {
	if value == nil {
		return float32(0), errors.New("Can not convert type 'nil' to 'float32'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case int:
		return float32(value.(int)), nil
	case int8:
		return float32(value.(int8)), nil
	case int16:
		return float32(value.(int16)), nil
	case int32:
		return float32(value.(int32)), nil
	case int64:
		return float32(value.(int64)), nil
	case uint:
		return float32(value.(uint)), nil
	case uint8:
		return float32(value.(uint8)), nil
	case uint16:
		return float32(value.(uint16)), nil
	case uint32:
		return float32(value.(uint32)), nil
	case uint64:
		return float32(value.(uint64)), nil
	case float32:
		return float32(value.(float32)), nil
	case float64:
		return float32(value.(float64)), nil
	case string:
		i, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}
		return float32(i), nil
	//... etc
	default:
		return 0, errors.New(fmt.Sprintf("Can not convert type '%s' to 'float32'.", GetType(value)))
	}

	return value.(float32), nil
}

func AsFloat64(value interface{}) (float64, error) {
	if value == nil {
		return float64(0), errors.New("Can not convert type 'nil' to 'float64'.")
	}

	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}
		return 0, nil
	case int:
		return float64(value.(int)), nil
	case int8:
		return float64(value.(int8)), nil
	case int16:
		return float64(value.(int16)), nil
	case int32:
		return float64(value.(int32)), nil
	case int64:
		return float64(value.(int64)), nil
	case uint:
		return float64(value.(uint)), nil
	case uint8:
		return float64(value.(uint8)), nil
	case uint16:
		return float64(value.(uint16)), nil
	case uint32:
		return float64(value.(uint32)), nil
	case uint64:
		return float64(value.(uint64)), nil
	case float32:
		return float64(value.(float32)), nil
	case float64:
		return float64(value.(float64)), nil
	case string:
		i, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}
		return float64(i), nil
	//... etc
	default:
		return 0, errors.New(fmt.Sprintf("Can not convert type '%s' to 'float64'.", GetType(value)))
	}

	return value.(float64), nil
}
