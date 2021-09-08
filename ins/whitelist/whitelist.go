package whitelist

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/industry-netsecurity-solution/ins-security-channel/fmterrors"
	"github.com/industry-netsecurity-solution/ins-security-channel/ins"
	"github.com/industry-netsecurity-solution/ins-security-channel/shared"
)

/*
 * 제조현장 스마트 게이트웨이 파일 전송을 확인한다.
 */
func IsAllowYM0x0001TO0006(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, nil
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowUWBLocation(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, nil
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowEventCollision(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, nil
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowCenterStatusInFactory(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, nil
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowRelayStatusInFactory(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, nil
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

/*
 * 제조현장 스마트 게이트웨이 상태 정보를 확인한다.
 */
func IsAllowSmartStatusInFactory(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowCenterStatusInDangerzone(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowSmartStatusInDangerzone(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowPortableStatusInConstruction(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowRelayStatusInConstruction(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	offset := uint32(0)

	for offset < tl32v.Length {
		data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += uint32(data.Size())

		if bytes.HasPrefix(data.Type, ins.BBx0000) {
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			} else {
				return true, nil
			}
		}
	}

	return true, nil
}

func IsAllowYM0xFFFF(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	return true, nil
}

func IsAllowYMTECH(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	data, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value)
	if err != nil {
		return false, err
	}
	if bytes.HasPrefix(data.Type, ins.BB_FRONT_VIDEO) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_REAR_VIDEO) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_FRONT_COLLISION) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_REAR_COLLISION) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_FRONT_APPROACH) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_REAR_APPROACH) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_RAW_ACCELEROMETER) {
		//return IsAllowYM0x0001TO0006(order, whiteGateway, whiteDevice, data)
		return true, nil
	} else if bytes.HasPrefix(data.Type, ins.BB_UWB_LOCATION) {
		return IsAllowUWBLocation(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.BB_EVENT_COLLISION) {
		return IsAllowEventCollision(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_CENTER_STATUS_FACTORY) {
		return IsAllowCenterStatusInFactory(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_RELAY_STATUS_FACTORY) {
		return IsAllowRelayStatusInFactory(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_SMART_STATUS_FACTORY) {
		return IsAllowSmartStatusInFactory(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_CENTER_STATUS_DANGERZONE) {
		return IsAllowCenterStatusInDangerzone(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_SMART_STATUS_DANGERZONE) {
		return IsAllowSmartStatusInDangerzone(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_PORTABLE_STATUS_CONSTRUCTION) {
		return IsAllowPortableStatusInConstruction(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, ins.GW_RELAY_STATUS_CONSTRUCTION) {
		return IsAllowRelayStatusInConstruction(order, whiteGateway, whiteDevice, data)
	} else if bytes.HasPrefix(data.Type, []byte{0xFF, 0xFF}) {
		return IsAllowYM0xFFFF(order, whiteGateway, whiteDevice, data)
	} else {
		return false, fmterrors.Error("unknown message: ", data.Type)
	}

	return true, nil
}

func IsAllowELSSEN(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	return true, nil
}

func IsAllowTELEFIELD(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	buf := bytes.NewBuffer(tl32v.Value)
	deviceId := make([]byte, 2)
	buf.Read(deviceId)

	if whiteDevice != nil {
		if whiteDevice.Has(hex.EncodeToString(deviceId)) == false {
			return false, nil
		}
	}

	return true, nil
}

func IsAllowABRAIN(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {
	buf := bytes.NewBuffer(tl32v.Value)

	status := make([]byte, 1)
	buf.Read(status)

	equip := make([]byte, 6)
	buf.Read(equip)

	if whiteDevice != nil {
		if whiteDevice.Has(hex.EncodeToString(equip)) == false {
			return false, nil
		}
	}

	return true, nil
}

func IsAllowWRAPPED(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	wrappedData, err := ins.DecTL32V(binary.LittleEndian, tl32v.Value)
	if err != nil {
		return false, err
	}

	var dataTlv *ins.TL32V = nil

	offset := 0
	for offset < int(wrappedData.Length) {
		data, err := ins.DecTL32V(binary.LittleEndian, wrappedData.Value[offset:])
		if err != nil {
			return false, err
		}
		offset += data.Size()

		if bytes.HasPrefix(data.Type, []byte{0x80, 0x01}) {
			//gwidTlv = data
			sourceId := string(data.Value)
			if whiteGateway.Has(sourceId) == false {
				return false, nil
			}
		} else if bytes.HasPrefix(data.Type, []byte{0x80, 0x02}) {
			// DO Nothing
		} else {
			dataTlv = data
		}
	}

	if dataTlv == nil {
		return false, fmterrors.Error("malformed message: ", tl32v.Type)
	}

	if bytes.HasPrefix(dataTlv.Type, ins.CODE_WRAPPED) {
		return IsAllowWRAPPED(order, whiteGateway, whiteDevice, dataTlv)
	} else if bytes.HasPrefix(dataTlv.Type, ins.CODE_YMTECH) {
		return IsAllowYMTECH(order, whiteGateway, whiteDevice, dataTlv)
	} else if bytes.HasPrefix(dataTlv.Type, ins.CODE_ELSSEN) {
		return IsAllowELSSEN(order, whiteGateway, whiteDevice, dataTlv)
	} else if bytes.HasPrefix(dataTlv.Type, ins.CODE_ABRAIN) {
		return IsAllowABRAIN(order, whiteGateway, whiteDevice, dataTlv)
	} else if bytes.HasPrefix(dataTlv.Type, ins.CODE_TELEFIELD) {
		return IsAllowTELEFIELD(order, whiteGateway, whiteDevice, dataTlv)
	} else if bytes.HasPrefix(dataTlv.Type, []byte{0x10, 0x01}) {
		// 엘센 메시지 포함 GW 메시지
		return false, fmterrors.Errorf("=---------------------------")
	} else if bytes.HasPrefix(dataTlv.Type, []byte{0x20, 0x01}) {
		// 텔레필드 메시지 포함 GW 메시지
		return false, fmterrors.Errorf("=---------------------------")
	} else if bytes.HasPrefix(dataTlv.Type, []byte{0x30, 0x01}) {
		// 에이브레인 메시지 포함 GW 메시지
		return false, fmterrors.Errorf("=---------------------------")
	}

	return true, nil
}

func IsAllowMessage(order binary.ByteOrder, whiteGateway, whiteDevice *shared.ConcurrentMap, tl32v *ins.TL32V) (bool, error) {

	if bytes.HasPrefix(tl32v.Type, ins.CODE_WRAPPED) {
		return IsAllowWRAPPED(order, whiteGateway, whiteDevice, tl32v)
	} else if bytes.HasPrefix(tl32v.Type, ins.CODE_YMTECH) {
		return IsAllowYMTECH(order, whiteGateway, whiteDevice, tl32v)
	} else if bytes.HasPrefix(tl32v.Type, ins.CODE_ELSSEN) {
		return IsAllowELSSEN(order, whiteGateway, whiteDevice, tl32v)
	} else if bytes.HasPrefix(tl32v.Type, ins.CODE_ABRAIN) {
		return IsAllowABRAIN(order, whiteGateway, whiteDevice, tl32v)
	} else if bytes.HasPrefix(tl32v.Type, ins.CODE_TELEFIELD) {
		return IsAllowTELEFIELD(order, whiteGateway, whiteDevice, tl32v)
	}

	return true, nil
}
