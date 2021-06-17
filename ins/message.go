package ins

import (
	"bytes"
	"encoding/binary"
)


var BBx0000 = []byte{0x00, 0x00}
var BBx0001 = []byte{0x00, 0x01}
var BBx0002 = []byte{0x00, 0x02}
var BBx0003 = []byte{0x00, 0x03}
var BBx0004 = []byte{0x00, 0x04}
var BBx0005 = []byte{0x00, 0x05}
var BBx0006 = []byte{0x00, 0x06}
var BBx0007 = []byte{0x00, 0x07}
var BBx0008 = []byte{0x00, 0x08}
var BBx0009 = []byte{0x00, 0x09}
var BBx000A = []byte{0x00, 0x0A}
var BBx000B = []byte{0x00, 0x0B}
var BBx000C = []byte{0x00, 0x0C}
var BBx000D = []byte{0x00, 0x0D}
var BBx000E = []byte{0x00, 0x0E}
var BBx000F = []byte{0x00, 0x0F}


var CODE_WRAPPED = []byte{0xEF, 0xF0}
var CODE_ELSSEN = []byte{0xEA, 0xCE}
var CODE_YMTECH = []byte{0xEF, 0xFE}
var CODE_TELEFIELD = []byte{0xFF, 0x00}
var CODE_ABRAIN = []byte{0xAA, 0x00}


// 전방/후방 영상 파일
var BB_FRONT_VIDEO = []byte{0x00, 0x01}
var BB_REAR_VIDEO = []byte{0x00, 0x02}
// 전방/후방 충돌 파일
var BB_FRONT_COLLISION = []byte{0x00, 0x03}
var BB_REAR_COLLISION = []byte{0x00, 0x04}
// 전방/후방 접근감지 파일
var BB_FRONT_APPROACH = []byte{0x00, 0x05}
var BB_REAR_APPROACH = []byte{0x00, 0x06}
// RAW 가속도 데이터 파일
var BB_RAW_ACCELEROMETER = []byte{0x00, 0x07}

var BB_UWB_LOCATION = []byte{0x00, 0x08}
// 블랙박스 충돌 이벤트
var BB_EVENT_COLLISION = []byte{0x00, 0x09}

// 제조현장
var GW_CENTER_FACTORY = []byte{0x00, 0x81}
var GW_RELAY_FACTORY = []byte{0x00, 0x82}
var GW_SMART_FACTORY = []byte{0x00, 0x83}

// 위험구역
var GW_CENTER_DANGERZONE = []byte{0x00, 0x84}
var GW_SMART_DANGERZONE = []byte{0x00, 0x85}

// 건설현장
var GW_PORTABLE_CONSTRUCTION = []byte{0x00, 0x86}
var GW_RELAY_CONSTRUCTION = []byte{0x00, 0x87}

// 텔레필드 레이다(RADAR)
var BB_RADAR_APPROACH = []byte{0x20, 0x01}
// 에이브레인 작업자 식별
var BB_WORKER_IDENTITY = []byte{0x30, 0x01}

const ELSSEN_WEARABLE_DEVICE int = 1
const ELSSEN_SAFETY_HOOK int = 2
const ELSSEN_TOXIC_GAS int = 3

// 건설 현장 생체
var GW_ELSSEN_WEARABLE_DEVICE = []byte{0x10, byte(ELSSEN_WEARABLE_DEVICE)}
// 건설 현장 안전고리
var GW_ELSSEN_SAFETY_HOOK = []byte{0x10, byte(ELSSEN_SAFETY_HOOK)}

// 건설 현장 유해가스
var GW_ELSSEN_TOXIC_GAS = []byte{0x10, byte(ELSSEN_TOXIC_GAS)}


// 중계 -> 집중
var BB_UWB_DISTANCE = []byte{0x00, 0x58}



var NAME_CODE_WRAPPED = "Wrapped"
var NAME_CODE_ELSSEN = "ELSSEN"
var NAME_CODE_YMTECH = "YMTECH"
var NAME_CODE_TELEFIELD = "TELEFIELD"
var NAME_CODE_ABRAIN = "ABRAIN"

// 전방/후방 영상 파일
var NAME_BB_FRONT_VIDEO = "front.video"
var NAME_BB_REAR_VIDEO = "rear.video"
// 전방/후방 충돌 파일
var NAME_BB_FRONT_COLLISION = "front.collision"
var NAME_BB_REAR_COLLISION = "rear.collision"
// 전방/후방 접근감지 파일
var NAME_BB_FRONT_APPROACH = "front.approach"
var NAME_BB_REAR_APPROACH = "rear.approach"
// RAW 가속도 데이터 파일
var NAME_BB_RAW_ACCELEROMETER = "RAW 6축 가속도 파일"

var NAME_BB_UWB_LOCATION = "event.location"
// 블랙박스 충돌 이벤트
var NAME_BB_EVENT_COLLISION = "event.collision"

// 제조현장
var NAME_GW_CENTER_FACTORY = "center.gw.status"
var NAME_GW_RELAY_FACTORY = "relay.gw.status"
var NAME_GW_SMART_FACTORY = "스마트 게이트웨이"

// 위험구역
var NAME_GW_CENTER_DANGERZONE = "center.gw.status"
var NAME_GW_SMART_DANGERZONE = "smart.gw.status"

// 건설현장
var NAME_GW_PORTABLE_CONSTRUCTION = "portable.gw.status"
var NAME_GW_RELAY_CONSTRUCTION = "relay.gw.status"

// 텔레필드 레이다(RADAR)
var NAME_BB_RADAR_APPROACH = "radar.approach"
// 에이브레인 작업자 식별
var NAME_BB_WORKER_IDENTITY = "worker.identify"

// 건설 현장 생체
var NAME_GW_ELSSEN_WEARABLE_DEVICE = "wearable"
// 건설 현장 안전고리
var NAME_GW_ELSSEN_SAFETY_HOOK = "safetyhook"
// 건설 현장 유해가스
var NAME_GW_ELSSEN_TOXIC_GAS = "toxicgas"


var NAME_UNKNOWN = "unknown"


var TYPE_CODE_WRAPPED = "Wrapped"
var TYPE_CODE_ELSSEN = "ELSSEN"
var TYPE_CODE_YMTECH = "YMTECH"
var TYPE_CODE_TELEFIELD = "TELEFIELD"
var TYPE_CODE_ABRAIN = "ABRAIN"

// 전방/후방 영상 파일
var TYPE_BB_FRONT_VIDEO = "전방 영상 파일"
var TYPE_BB_REAR_VIDEO = "후방 영상 파일"
// 전방/후방 충돌 파일
var TYPE_BB_FRONT_COLLISION = "전방 충돌 이력 파일"
var TYPE_BB_REAR_COLLISION = "후방 충돌 이력 파일"
// 전방/후방 접근감지 파일
var TYPE_BB_FRONT_APPROACH = "전방 접근감지 이력 파일"
var TYPE_BB_REAR_APPROACH = "후방 접근감지 이력 파일"
// RAW 가속도 데이터 파일
var TYPE_BB_RAW_ACCELEROMETER = "RAW 6축 가속도 파일"

var TYPE_BB_UWB_LOCATION = "UWB 위치/속도"
// 블랙박스 충돌 이벤트
var TYPE_BB_EVENT_COLLISION = "가속도 충돌 이벤트"

// 제조현장
var TYPE_GW_CENTER_FACTORY = "집중 게이트웨이"
var TYPE_GW_RELAY_FACTORY = "중계 게이트웨이"
var TYPE_GW_SMART_FACTORY = "스마트 게이트웨이"

// 위험구역
var TYPE_GW_CENTER_DANGERZONE = "집중 게이트웨이"
var TYPE_GW_SMART_DANGERZONE = "스마트 게이트웨이"

// 건설현장
var TYPE_GW_PORTABLE_CONSTRUCTION = "이동형 게이트웨이"
var TYPE_GW_RELAY_CONSTRUCTION = "중계 게이트웨이"

// 텔레필드 레이다(RADAR)
var TYPE_BB_RADAR_APPROACH = "접근 감지 레이다"
// 에이브레인 작업자 식별
var TYPE_BB_WORKER_IDENTITY = "작업자 식별"

// 건설 현장 생체
var TYPE_GW_ELSSEN_WEARABLE_DEVICE = "안전고리"
// 건설 현장 안전고리
var TYPE_GW_ELSSEN_SAFETY_HOOK = "생체정보"
// 건설 현장 유해가스
var TYPE_GW_ELSSEN_TOXIC_GAS = "유해가스"

var TYPE_UNKNOWN = "알 수 없음"


const METHOD_ERROR int = -1
const METHOD_UNKNOWN int = 0
const METHOD_SOCKET int = 1
const METHOD_MQTT int = 2

/**
 * 0xEF, 0xFE
 */
func GetTransmissionMethod4YMTECH(order binary.ByteOrder, data []byte) int {
	if data == nil || len(data) < 2 {
		return METHOD_ERROR
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return METHOD_MQTT
	}

	return METHOD_UNKNOWN
}

/**
 * 0xEF, 0xF0 메시지
 */
func GetTransmissionMethod4Wrap(order binary.ByteOrder, data []byte) int {
	if data == nil || len(data) < 2 {
		return METHOD_ERROR
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, BB_RADAR_APPROACH) {
		// 텔레필드
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, BB_WORKER_IDENTITY) {
		// 에이브레인
		return METHOD_MQTT
	} else if bytes.HasPrefix(data, GW_ELSSEN_WEARABLE_DEVICE) {
		// 건설현장 생체정보
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_ELSSEN_SAFETY_HOOK) {
		// 건설현장 안전고리
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, GW_ELSSEN_TOXIC_GAS) {
		// 건설현장 유해가스
		return METHOD_SOCKET
	}

	return METHOD_UNKNOWN
}

func GetTransmissionMethod(order binary.ByteOrder, data []byte) int {
	if data == nil || len(data) < 2 {
		return METHOD_ERROR
	}

	if bytes.HasPrefix(data, CODE_WRAPPED) {
		return GetTransmissionMethod4Wrap(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_ELSSEN) {
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, CODE_YMTECH) {
		return GetTransmissionMethod4YMTECH(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_TELEFIELD) {
		return METHOD_SOCKET
	} else if bytes.HasPrefix(data, CODE_ABRAIN) {
		return METHOD_MQTT
	}

	return METHOD_UNKNOWN
}

/**
 * 0xEA, 0xCE
 */
func GetMessageType4ELSSEN(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 7 {
		return TYPE_UNKNOWN
	}

	if bytes.HasPrefix(data, CODE_ELSSEN) == false {
		return TYPE_UNKNOWN
	}

	info := data[7]
	subid := info >> 4 & 0x0F

	switch subid {
	case 1:
		return TYPE_GW_ELSSEN_WEARABLE_DEVICE
	case 2:
		return TYPE_GW_ELSSEN_SAFETY_HOOK
	case 3:
		return TYPE_GW_ELSSEN_TOXIC_GAS
	}

	return TYPE_UNKNOWN
}

/**
 * 0xEF, 0xFE
 */
func GetMessageType4YMTECH(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return TYPE_UNKNOWN
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return TYPE_BB_FRONT_VIDEO
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return TYPE_BB_REAR_VIDEO
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return TYPE_BB_FRONT_COLLISION
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return TYPE_BB_REAR_COLLISION
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return TYPE_BB_FRONT_APPROACH
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return TYPE_BB_REAR_APPROACH
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return TYPE_BB_RAW_ACCELEROMETER
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return TYPE_BB_UWB_LOCATION
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return TYPE_BB_EVENT_COLLISION
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return TYPE_GW_CENTER_FACTORY
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return TYPE_GW_RELAY_FACTORY
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return TYPE_GW_SMART_FACTORY
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return TYPE_GW_CENTER_DANGERZONE
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return TYPE_GW_SMART_DANGERZONE
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return TYPE_GW_PORTABLE_CONSTRUCTION
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return TYPE_GW_RELAY_CONSTRUCTION
	}

	return TYPE_UNKNOWN
}

/**
 * 0xEF, 0xF0 메시지
 */
func GetMessageType4Wrap(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return TYPE_UNKNOWN
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return TYPE_BB_FRONT_VIDEO
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return TYPE_BB_REAR_VIDEO
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return TYPE_BB_FRONT_COLLISION
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return TYPE_BB_REAR_COLLISION
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return TYPE_BB_FRONT_APPROACH
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return TYPE_BB_REAR_APPROACH
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return TYPE_BB_RAW_ACCELEROMETER
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return TYPE_BB_UWB_LOCATION
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return TYPE_BB_EVENT_COLLISION
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return TYPE_GW_CENTER_FACTORY
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return TYPE_GW_RELAY_FACTORY
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return TYPE_GW_SMART_FACTORY
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return TYPE_GW_CENTER_DANGERZONE
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return TYPE_GW_SMART_DANGERZONE
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return TYPE_GW_PORTABLE_CONSTRUCTION
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return TYPE_GW_RELAY_CONSTRUCTION
	} else if bytes.HasPrefix(data, BB_RADAR_APPROACH) {
		// 텔레필드
		return TYPE_BB_RADAR_APPROACH
	} else if bytes.HasPrefix(data, BB_WORKER_IDENTITY) {
		// 에이브레인
		return TYPE_BB_WORKER_IDENTITY
	} else if bytes.HasPrefix(data, GW_ELSSEN_WEARABLE_DEVICE) {
		// 건설현장 생체정보
		return TYPE_GW_ELSSEN_WEARABLE_DEVICE
	} else if bytes.HasPrefix(data, GW_ELSSEN_SAFETY_HOOK) {
		// 건설현장 안전고리
		return TYPE_GW_ELSSEN_SAFETY_HOOK
	} else if bytes.HasPrefix(data, GW_ELSSEN_TOXIC_GAS) {
		// 건설현장 유해가스
		return TYPE_GW_ELSSEN_TOXIC_GAS
	}

	return TYPE_UNKNOWN
}

func GetMessageType(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return TYPE_UNKNOWN
	}

	if bytes.HasPrefix(data, CODE_WRAPPED) {
		return GetMessageType4Wrap(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_ELSSEN) {
		return GetMessageType4ELSSEN(order, data)
	} else if bytes.HasPrefix(data, CODE_YMTECH) {
		return GetMessageType4YMTECH(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_TELEFIELD) {
		return TYPE_CODE_TELEFIELD
	} else if bytes.HasPrefix(data, CODE_ABRAIN) {
		return TYPE_CODE_ABRAIN
	}

	return TYPE_UNKNOWN
}


/**
 * 0xEF, 0xFE
 */
func GetMessageName4YMTECH(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return NAME_UNKNOWN
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return NAME_BB_FRONT_VIDEO
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return NAME_BB_REAR_VIDEO
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return NAME_BB_FRONT_COLLISION
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return NAME_BB_REAR_COLLISION
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return NAME_BB_FRONT_APPROACH
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return NAME_BB_REAR_APPROACH
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return NAME_BB_RAW_ACCELEROMETER
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return NAME_BB_UWB_LOCATION
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return NAME_BB_EVENT_COLLISION
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return NAME_GW_CENTER_FACTORY
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return NAME_GW_RELAY_FACTORY
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return NAME_GW_SMART_FACTORY
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return NAME_GW_CENTER_DANGERZONE
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return NAME_GW_SMART_DANGERZONE
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return NAME_GW_PORTABLE_CONSTRUCTION
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return NAME_GW_RELAY_CONSTRUCTION
	}

	return NAME_UNKNOWN
}

/**
 * 0xEA, 0xCE
 */
func GetMessageName4ELSSEN(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 7 {
		return NAME_UNKNOWN
	}

	if bytes.HasPrefix(data, CODE_ELSSEN) == false {
		return NAME_UNKNOWN
	}

	info := data[7]
	subid := info >> 4 & 0x0F

	switch subid {
	case 1:
		return NAME_GW_ELSSEN_WEARABLE_DEVICE
	case 2:
		return NAME_GW_ELSSEN_SAFETY_HOOK
	case 3:
		return NAME_GW_ELSSEN_TOXIC_GAS
	}

	return NAME_UNKNOWN
}

/**
 * 0xEF, 0xF0 메시지
 */
func GetMessageName4Wrap(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return NAME_UNKNOWN
	}

	if bytes.HasPrefix(data, BB_FRONT_VIDEO) {
		return NAME_BB_FRONT_VIDEO
	} else if bytes.HasPrefix(data, BB_REAR_VIDEO) {
		return NAME_BB_REAR_VIDEO
	} else if bytes.HasPrefix(data, BB_FRONT_COLLISION) {
		return NAME_BB_FRONT_COLLISION
	} else if bytes.HasPrefix(data, BB_REAR_COLLISION) {
		return NAME_BB_REAR_COLLISION
	} else if bytes.HasPrefix(data, BB_FRONT_APPROACH) {
		return NAME_BB_FRONT_APPROACH
	} else if bytes.HasPrefix(data, BB_REAR_APPROACH) {
		return NAME_BB_REAR_APPROACH
	} else if bytes.HasPrefix(data, BB_RAW_ACCELEROMETER) {
		return NAME_BB_RAW_ACCELEROMETER
	} else if bytes.HasPrefix(data, BB_UWB_LOCATION) {
		// 제조현장 지게차 UWB 위치 정보
		return NAME_BB_UWB_LOCATION
	} else if bytes.HasPrefix(data, BB_EVENT_COLLISION) {
		// 제조현장 지게차 충돌 이벤트
		return NAME_BB_EVENT_COLLISION
	} else if bytes.HasPrefix(data, GW_CENTER_FACTORY) {
		// 제조현장 지게차 집중 GW
		return NAME_GW_CENTER_FACTORY
	} else if bytes.HasPrefix(data, GW_RELAY_FACTORY) {
		// 제조현장 지게차 중계 GW
		return NAME_GW_RELAY_FACTORY
	} else if bytes.HasPrefix(data, GW_SMART_FACTORY) {
		// 제조현장 지게차 스마트 GW
		return NAME_GW_SMART_FACTORY
	} else if bytes.HasPrefix(data, GW_CENTER_DANGERZONE) {
		// 제조현장 위험구역 집중 GW
		return NAME_GW_CENTER_DANGERZONE
	} else if bytes.HasPrefix(data, GW_SMART_DANGERZONE) {
		// 제조현장 위험구역 스마트 GW
		return NAME_GW_SMART_DANGERZONE
	} else if bytes.HasPrefix(data, GW_PORTABLE_CONSTRUCTION) {
		// 건설현장 이동형 GW
		return NAME_GW_PORTABLE_CONSTRUCTION
	} else if bytes.HasPrefix(data, GW_RELAY_CONSTRUCTION) {
		// 건설현장 중계 GW
		return NAME_GW_RELAY_CONSTRUCTION
	} else if bytes.HasPrefix(data, BB_RADAR_APPROACH) {
		// 텔레필드
		return NAME_BB_RADAR_APPROACH
	} else if bytes.HasPrefix(data, BB_WORKER_IDENTITY) {
		// 에이브레인
		return NAME_BB_WORKER_IDENTITY
	} else if bytes.HasPrefix(data, GW_ELSSEN_WEARABLE_DEVICE) {
		// 건설현장 생체정보
		return NAME_GW_ELSSEN_WEARABLE_DEVICE
	} else if bytes.HasPrefix(data, GW_ELSSEN_SAFETY_HOOK) {
		// 건설현장 안전고리
		return NAME_GW_ELSSEN_SAFETY_HOOK
	} else if bytes.HasPrefix(data, GW_ELSSEN_TOXIC_GAS) {
		// 건설현장 유해가스
		return NAME_GW_ELSSEN_TOXIC_GAS
	}

	return NAME_UNKNOWN
}

func GetMessageName(order binary.ByteOrder, data []byte) string {
	if data == nil || len(data) < 2 {
		return NAME_UNKNOWN
	}

	if bytes.HasPrefix(data, CODE_WRAPPED) {
		return GetMessageName4Wrap(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_ELSSEN) {
		return GetMessageName4ELSSEN(order, data)
	} else if bytes.HasPrefix(data, CODE_YMTECH) {
		return GetMessageName4YMTECH(order, data[6:])
	} else if bytes.HasPrefix(data, CODE_TELEFIELD) {
		return NAME_CODE_TELEFIELD
	} else if bytes.HasPrefix(data, CODE_ABRAIN) {
		return NAME_CODE_ABRAIN
	}

	return NAME_UNKNOWN
}

/**
 * 0xEA, 0xCE
 */
func GetSubID4ELSSEN(order binary.ByteOrder, data []byte) int {
	if data == nil || len(data) < 7 {
		return -1
	}

	if bytes.HasPrefix(data, CODE_ELSSEN) == false {
		return -1
	}

	info := data[7]
	subid := info >> 4 & 0x0F

	return int(subid)
}

/**
 * 0xEF, 0xFE
 */
func GetSubID4YMTECH(order binary.ByteOrder, data []byte) []byte {
	if data == nil || len(data) < 8 {
		return nil
	}

	if bytes.HasPrefix(data, CODE_YMTECH) == false {
		return nil
	}

	return data[6:8]
}
