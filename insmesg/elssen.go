package insmesg

import (
	"bytes"
	"encoding/binary"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"time"
)

/**
 * 엘센 메시지 전달하기
 */
func MakeWrappedPacketFor0xEACE(data []byte, addtion ins.Map) *bytes.Buffer {

	// level 3 payload
	l3payload := bytes.Buffer{}

	// 원본
	l3payload.Write(data)

	// GW 식별
	gwid := addtion.Get(ins.MapKey([]byte {0x80, 0x01}))
	if gwid == nil {
		l3payload.Write(ins.EncTagLnV(binary.LittleEndian, []byte{0x80, 0x01}, 32, []byte{}))
	} else {
		l3payload.Write(ins.EncTagLnV(binary.LittleEndian, []byte{0x80, 0x01}, 32, gwid.([]byte)))
	}
	// 시간 추가
	unix32 := addtion.Get(ins.MapKey([]byte {0x80, 0x02}))
	if unix32 == nil {
		unix32 = uint32(time.Now().Unix())
	}
	l3payload.Write(ins.EncTagLnUInt32(binary.LittleEndian, []byte{0x80, 0x02}, 32, unix32.(uint32)))

	// level 2 payload
	l2payload := bytes.Buffer{}

	subid := ins.GetSubID4ELSSEN(binary.LittleEndian, data)
	if subid == ins.ELSSEN_WEARABLE_DEVICE {
		//0x10, 01, payloadLength, payload
		l2payload.Write(ins.EncTagLnV(binary.LittleEndian, ins.GW_ELSSEN_WEARABLE_DEVICE, 32, l3payload.Bytes()))
	} else 	if subid == ins.ELSSEN_SAFETY_HOOK {
		//0x10, 02, payloadLength, payload
		l2payload.Write(ins.EncTagLnV(binary.LittleEndian, ins.GW_ELSSEN_SAFETY_HOOK, 32, l3payload.Bytes()))
	} else 	if subid == ins.ELSSEN_TOXIC_GAS {
		//0x10, 03, payloadLength, payload
		l2payload.Write(ins.EncTagLnV(binary.LittleEndian, ins.GW_ELSSEN_TOXIC_GAS, 32, l3payload.Bytes()))
	}

	buffer := &bytes.Buffer{}

	//0xEF, F0, payloadLength, payload
	buffer.Write(ins.EncTagLnV(binary.LittleEndian, ins.CODE_WRAPPED, 32, l2payload.Bytes()))

	return buffer
}
