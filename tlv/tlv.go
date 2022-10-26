package tlv

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type TL16V struct {
	Type   []byte
	Length uint16
	Value  interface{}
}

type TL32V struct {
	Type   []byte
	Length uint32
	Value  interface{}
}

type TL64V struct {
	Type   []byte
	Length uint64
	Value  interface{}
}

func (obj TL16V) Encode(order binary.ByteOrder) ([]byte, error) {
	buffer := bytes.Buffer{}

	// Type
	buffer.Write(obj.Type)

	// Length
	bl := make([]byte, 2)
	order.PutUint16(bl, obj.Length)
	buffer.Write(bl)

	vb := bytes.Buffer{}

	switch t := obj.Value.(type) {
	case []byte:
		vb.Write(obj.Value.([]byte))
	case []TL16V:
		for _, e := range obj.Value.([]TL16V) {
			d, err := e.Encode(order)
			if err != nil {
				return nil, err
			}
			vb.Write(d)
		}
	case TL16V:
		d, err := obj.Value.(TL16V).Encode(order)
		if err != nil {
			return nil, err
		}
		vb.Write(d)
	default:
		return nil, errors.New(fmt.Sprintf("Not support type: %T", t))
	}

	buffer.Write(vb.Bytes())

	return buffer.Bytes(), nil
}

func (obj TL32V) Encode(order binary.ByteOrder) ([]byte, error) {
	buffer := bytes.Buffer{}

	// Type
	buffer.Write(obj.Type)

	// Length
	bl := make([]byte, 4)
	order.PutUint32(bl, obj.Length)
	buffer.Write(bl)

	vb := bytes.Buffer{}

	switch t := obj.Value.(type) {
	case []byte:
		buffer.Write(obj.Value.([]byte))
	case []TL32V:
		for _, e := range obj.Value.([]TL32V) {
			d, err := e.Encode(order)
			if err != nil {
				return nil, err
			}
			buffer.Write(d)
		}
	case TL32V:
		d, err := obj.Value.(TL32V).Encode(order)
		if err != nil {
			return nil, err
		}
		buffer.Write(d)
	default:
		return nil, errors.New(fmt.Sprintf("Not support type: %T", t))
	}

	buffer.Write(vb.Bytes())

	return buffer.Bytes(), nil
}

func (obj TL64V) Encode(order binary.ByteOrder) ([]byte, error) {
	buffer := bytes.Buffer{}

	// Type
	buffer.Write(obj.Type)

	// Length
	bl := make([]byte, 8)
	order.PutUint64(bl, obj.Length)
	buffer.Write(bl)

	vb := bytes.Buffer{}

	switch t := obj.Value.(type) {
	case []byte:
		buffer.Write(obj.Value.([]byte))
	case []TL64V:
		for _, e := range obj.Value.([]TL64V) {
			d, err := e.Encode(order)
			if err != nil {
				return nil, err
			}
			buffer.Write(d)
		}
	case TL64V:
		d, err := obj.Value.(TL64V).Encode(order)
		if err != nil {
			return nil, err
		}
		buffer.Write(d)
	default:
		return nil, errors.New(fmt.Sprintf("Not support type: %T", t))
	}

	buffer.Write(vb.Bytes())

	return buffer.Bytes(), nil
}

func (obj TL16V) Decode(order binary.ByteOrder, data []byte) (int64, error) {
	if data == nil || len(data) < 4 {
		return -1, errors.New("not enough data length")
	}

	length := int64(0)
	t := data[0:2]
	obj.Type = t
	length += int64(len(t))

	l := data[2:4]
	obj.Length = order.Uint16(l)
	length += int64(len(l))

	v := data[4 : 4+obj.Length]
	obj.Value = v
	length += int64(len(v))

	return length, nil
}

func (obj TL32V) Decode(order binary.ByteOrder, data []byte) (int64, error) {
	if data == nil || len(data) < 6 {
		return -1, errors.New("not enough data length")
	}

	length := int64(0)
	t := data[0:2]
	obj.Type = t
	length += int64(len(t))

	l := data[2:6]
	obj.Length = order.Uint32(l)
	length += int64(len(l))

	v := data[6 : 6+obj.Length]
	obj.Value = v
	length += int64(len(v))

	return length, nil
}

func (obj TL64V) Decode(order binary.ByteOrder, data []byte) (int64, error) {
	if data == nil || len(data) < 10 {
		return -1, errors.New("not enough data length")
	}

	length := int64(0)
	t := data[0:2]
	obj.Type = t
	length += int64(len(t))

	l := data[2:10]
	obj.Length = order.Uint64(l)
	length += int64(len(l))

	v := data[10 : 10+obj.Length]
	obj.Value = v
	length += int64(len(v))

	return length, nil
}

func DecTL16V(order binary.ByteOrder, data []byte) (*TL16V, error) {
	if data == nil || len(data) < 4 {
		return nil, errors.New("not enough data length")
	}

	ret := new(TL16V)
	_, err := ret.Decode(order, data)
	return ret, err
}

func DecTL32V(order binary.ByteOrder, data []byte) (*TL32V, error) {
	if data == nil || len(data) < 6 {
		return nil, errors.New("not enough data length")
	}

	ret := new(TL32V)
	_, err := ret.Decode(order, data)
	return ret, err
}

func DecTL64V(order binary.ByteOrder, data []byte) (*TL64V, error) {
	if data == nil || len(data) < 10 {
		return nil, errors.New("not enough data length")
	}

	ret := new(TL64V)
	_, err := ret.Decode(order, data)
	return ret, err
}
