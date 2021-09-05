package ins

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type TL16V struct {
	Type []byte
	Length uint16
	Value []byte
}

type TL32V struct {
	Type []byte
	Length uint32
	Value []byte
}

type TL64V struct {
	Type []byte
	Length uint64
	Value []byte
}

func (v TL32V) Size() int {
	length := len(v.Type)
	length += 4
	length += int(v.Length)

	return length
}

func (v TL32V) Bytes(order binary.ByteOrder) []byte {
	var buffer = new(bytes.Buffer)
	binary.Write(buffer, order, v.Type)
	binary.Write(buffer, order, v.Length)
	binary.Write(buffer, order, v.Value)

	return buffer.Bytes()
}

func DecTL16V(order binary.ByteOrder, data []byte) (*TL16V, error) {
	if data == nil || len(data) < 4 {
		return nil, errors.New("not enough data length")
	}

	tlv := TL16V{}
	tlv.Type = data[0:2]
	tlv.Length = order.Uint16(data[2:4])
	tlv.Value = data[4:4+tlv.Length]

	return &tlv, nil
}

func DecTL32V(order binary.ByteOrder, data []byte) (*TL32V, error) {
	if data == nil || len(data) < 6 {
		return nil, errors.New("not enough data length")
	}

	tlv := TL32V{}
	tlv.Type = data[0:2]
	tlv.Length = order.Uint32(data[2:6])
	tlv.Value = data[6:6+tlv.Length]

	return &tlv, nil
}

func DecTL64V(order binary.ByteOrder, data []byte) (*TL64V, error) {
	if data == nil || len(data) < 10 {
		return nil, errors.New("not enough data length")
	}

	tlv := TL64V{}
	tlv.Type = data[0:2]
	tlv.Length = order.Uint64(data[2:10])
	tlv.Value = data[10:10+tlv.Length]

	return &tlv, nil
}

func TraceTLVMessage(order binary.ByteOrder, data []byte, handler func(tl32v *TL32V) int) (int, error) {
	offset := 0
	for offset < len(data) {
		tl32v, err := DecTL32V(order, data[offset:])
		if err != nil {
			return 0, err
		}
		offset += tl32v.Size()

		result := handler(tl32v)
		if result < 0 {
			return result, nil
		} else if result == 0 {
			continue
		} else {
			r, e := TraceTLVMessage(order, tl32v.Value, handler)
			if e != nil {
				return r, e
			}
			if r < 0 {
				return r, nil
			}
		}
	}

	return 0, nil
}

func EncodeMap(order binary.ByteOrder, params map[int][]byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	for key := range params {

		// encode key
		order.PutUint16(buf2, uint16(key))
		buffer.Write(buf2)

		// encode length
		paramLen := len(params[key])
		order.PutUint32(buf4, uint32(paramLen))
		buffer.Write(buf4)

		// encode data
		buffer.Write(params[key])
	}

	return buffer.Bytes()
}

func EncTagLn(order binary.ByteOrder, tag []byte, size uint, length uint32) []byte {
	var buffer = new(bytes.Buffer)
	binary.Write(buffer, order, tag)

	if size == 16 {
		binary.Write(buffer, order, uint16(length))
	} else if size == 32 {
		binary.Write(buffer, order, uint32(length))
	} else if size == 64 {
		binary.Write(buffer, order, uint64(length))
	} else {
		return nil
	}

	return buffer.Bytes()
}

func EncTagLnV(order binary.ByteOrder, tag []byte, size uint, data []byte) []byte {
	var buffer = new(bytes.Buffer)
	binary.Write(buffer, order, tag)

	if data == nil {
		if size == 16 {
			binary.Write(buffer, order, uint16(0))
		} else if size == 32 {
			binary.Write(buffer, order, uint32(0))
		} else if size == 64 {
			binary.Write(buffer, order, uint64(0))
		} else {
			return nil
		}

		return buffer.Bytes()
	}

	if size == 16 {
		binary.Write(buffer, order, uint16(len(data)))
	} else if size == 32 {
		binary.Write(buffer, order, uint32(len(data)))
	} else if size == 64 {
		binary.Write(buffer, order, uint64(len(data)))
	} else {
		return nil
	}

	buffer.Write(data)

	return buffer.Bytes()
}

func EncTagLnString(order binary.ByteOrder, tag []byte, size uint, data string) []byte {
	return EncTagLnV(order, tag, size, []byte(data))
}

func EncTagLnUInt16(order binary.ByteOrder, tag []byte, size uint, data uint16) []byte {
	buf := make([]byte, 2)
	order.PutUint16(buf, data)
	return EncTagLnV(order, tag, size, buf)
}

func EncTagLnUInt32(order binary.ByteOrder, tag []byte, size uint, data uint32) []byte {
	buf := make([]byte, 4)
	order.PutUint32(buf, data)
	return EncTagLnV(order, tag, size, buf)
}

func EncTagLnUInt64(order binary.ByteOrder, tag []byte, size uint, data uint64) []byte {
	buf := make([]byte, 8)
	order.PutUint64(buf, data)
	return EncTagLnV(order, tag, size, buf)
}

/*
 * LittleEndian
 */
func EncLETL(tag uint16, length uint32) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	binary.LittleEndian.PutUint32(buf4, length)
	buffer.Write(buf4)

	return buffer.Bytes()
}

func EncLETLnV(tag uint16, size uint, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)
	buf8 := make([]byte, 8)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		if size == 16 {
			binary.LittleEndian.PutUint16(buf2, uint16(0))
			buffer.Write(buf2)
		} else if size == 32 {
			binary.LittleEndian.PutUint32(buf4, uint32(0))
			buffer.Write(buf4)
		} else if size == 64 {
			binary.LittleEndian.PutUint64(buf8, uint64(0))
			buffer.Write(buf8)
		} else {
			return nil
		}

		return buffer.Bytes()
	}

	if size == 16 {
		binary.LittleEndian.PutUint16(buf2, uint16(len(data)))
		buffer.Write(buf2)
	} else if size == 32 {
		binary.LittleEndian.PutUint32(buf4, uint32(len(data)))
		buffer.Write(buf4)
	} else if size == 64 {
		binary.LittleEndian.PutUint64(buf8, uint64(len(data)))
		buffer.Write(buf8)
	} else {
		return nil
	}

	buffer.Write(data)

	return buffer.Bytes()
}

func EncLETLV(tag uint16, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.LittleEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		binary.LittleEndian.PutUint32(buf4, uint32(0))
		buffer.Write(buf4)
		return buffer.Bytes()
	}

	binary.LittleEndian.PutUint32(buf4, uint32(len(data)))
	buffer.Write(buf4)

	buffer.Write(data)

	return buffer.Bytes()
}

func EncLEString(tag uint16, data string) []byte {
	return EncLETLV(tag, []byte(data))
}

func EncLEUint16(tag uint16, data uint16) []byte {
	buf2 := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf2, data)

	return EncLETLV(tag, buf2)
}

func EncLEUint32(tag uint16, data uint32) []byte {
	buf4 := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf4, data)

	return EncLETLV(tag, buf4)
}

/*
 * BigEndian
 */
func EncBETL(tag uint16, length uint32) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.BigEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	binary.BigEndian.PutUint32(buf4, length)
	buffer.Write(buf4)

	return buffer.Bytes()
}

func EncBETLV(tag uint16, data []byte) []byte {
	buf2 := make([]byte, 2)
	buf4 := make([]byte, 4)

	var buffer bytes.Buffer

	binary.BigEndian.PutUint16(buf2, tag)
	buffer.Write(buf2)

	if data == nil {
		binary.BigEndian.PutUint32(buf4, uint32(0))
		buffer.Write(buf4)
		return buffer.Bytes()
	}

	binary.BigEndian.PutUint32(buf4, uint32(len(data)))
	buffer.Write(buf4)

	buffer.Write(data)

	return buffer.Bytes()
}

func EncBEString(tag uint16, data string) []byte {
	return EncBETLV(tag, []byte(data))
}

func EncBEUint16(tag uint16, data uint16) []byte {
	buf2 := make([]byte, 2)
	binary.BigEndian.PutUint16(buf2, data)

	return EncBETLV(tag, buf2)
}

func EncBEUint32(tag uint16, data uint32) []byte {
	buf4 := make([]byte, 4)
	binary.BigEndian.PutUint32(buf4, data)

	return EncBETLV(tag, buf4)
}
