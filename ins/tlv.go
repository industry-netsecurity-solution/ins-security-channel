package ins

import (
	"bytes"
	"encoding/binary"
)


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

func EncTagLnUInt32(order binary.ByteOrder, tag []byte, size uint, data uint32) []byte {
	buf4 := make([]byte, 4)
	order.PutUint32(buf4, data)
	return EncTagLnV(order, tag, size, buf4)
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
