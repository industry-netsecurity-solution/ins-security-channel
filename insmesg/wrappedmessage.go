package insmesg

import (
	"bytes"
	"encoding/binary"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"time"
)

func MakeWrappedPacketFor0xEFF0(data []byte, additional ins.Map) *bytes.Buffer {

	// level 3 payload
	l3payload := bytes.Buffer{}

	// 원본
	l3payload.Write(data)

	// GW 식별
	gwid := additional.Get(ins.MapKey([]byte {0x80, 0x01}))
	if gwid == nil {
		l3payload.Write(ins.EncTagLnV(binary.LittleEndian, []byte{0x80, 0x01}, 32, []byte{}))
	} else {
		l3payload.Write(ins.EncTagLnV(binary.LittleEndian, []byte{0x80, 0x01}, 32, gwid.([]byte)))
	}
	// 시간 추가
	unix32 := additional.Get(ins.MapKey([]byte {0x80, 0x02}))
	if unix32 == nil {
		unix32 = uint32(time.Now().Unix())
	}
	l3payload.Write(ins.EncTagLnUInt32(binary.LittleEndian, []byte{0x80, 0x02}, 32, unix32.(uint32)))

	// level 2 payload
	l2payload := bytes.Buffer{}

	//0x00, 00, payloadLength, payload
	l2payload.Write(ins.EncTagLnV(binary.LittleEndian, []byte{0xF0, 0x00}, 32, l3payload.Bytes()))

	buffer := &bytes.Buffer{}

	//0xEF, F0, payloadLength, payload
	buffer.Write(ins.EncTagLnV(binary.LittleEndian, ins.CODE_WRAPPED, 32, l2payload.Bytes()))

	return buffer

}

func MakeWrappedPacket(data []byte, addtion ins.Map) *bytes.Buffer {
	var newdata *bytes.Buffer
	if bytes.HasPrefix(data, ins.CODE_WRAPPED) {
		// 중계 메시지
		newdata = MakeWrappedPacketFor0xEFF0(data, addtion)
	} else if bytes.HasPrefix(data, ins.CODE_YMTECH) {
		// 유미테크
		newdata = MakeWrappedPacketFor0xEFFE(data, addtion)
	} else if bytes.HasPrefix(data, ins.CODE_ELSSEN) {
		// 엘센
		newdata = MakeWrappedPacketFor0xEACE(data, addtion)
	} else if bytes.HasPrefix(data, ins.CODE_TELEFIELD) {
		// 텔레필드
		newdata = MakeWrappedPacketFor0x8F8F(data, addtion)
	} else if bytes.HasPrefix(data, ins.CODE_ABRAIN) {
		// 에이브레인
		newdata = MakeWrappedPacketFor0xABAB(data, addtion)
	} else {
		return nil
	}

	return newdata
}