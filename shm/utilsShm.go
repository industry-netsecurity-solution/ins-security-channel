package shm

import "C"

/*
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>
#include <memory.h>
#include <sys/types.h>
#include <sys/time.h>

#pragma pack(push, 1)
typedef struct warning_ {
	struct  timeval startTime;
	struct  timeval evtTime;
	int32_t camera;
	int32_t event;
	int32_t frameIndex;
	uint8_t dataLen;
	uint8_t data[128];
} Warning;

typedef struct accelerometerthreshold_ {
	int32_t total;
	int32_t x;
	int32_t y;
	int32_t z;
} AccelerometerThreshold;

typedef struct accelerometer_ {
	struct timeval tv;
	int16_t x;
	int16_t y;
	int16_t z;
	int16_t vx;
	int16_t vy;
	int16_t vz;
	int16_t vt;
} Accelerometer;

typedef struct radarunit_ {
	uint8_t id;
	int8_t x;
	uint8_t y;
	uint8_t dist;
	int8_t d_speed;
	int8_t o_speed;
	uint8_t size;
	uint8_t flag;
} RadarUnit;

typedef struct radar {
	struct timeval tv;
	uint8_t id[2];
	RadarUnit u1;
	RadarUnit u2;
	RadarUnit u3;
	RadarUnit u4;
} Radar;

typedef struct positionphase {
	int32_t update;
	int32_t camera;
	double p1;
	double p2;
} PositionPhase;

typedef struct crashrisk_ {
	struct timeval alertTime;
	struct timeval recvTime;
	uint8_t  sendGwId[64];
	uint32_t uwbTagId;
	uint32_t speed;
	uint32_t count;
} CrashRisk;

typedef struct splittime_ {
	struct timeval splitTime;
	uint32_t fragment;
	int32_t  camera;
} SplitTime;

typedef struct tegraStats_ {
	struct timeval tv;
	uint32_t ram;
	uint32_t swap;
	float    cpu;
	float    iwlwifi;
	float    pmic;
	float    gpu;
	float    ao;
	float    thermal;
} TegraStats;

typedef struct diagnosisStats_ {
	struct timeval tvBattery;
	struct timeval tvUsbStorage;
	struct timeval tvCamera;
	struct timeval tvAccelerometer;
	struct timeval tvFan;
	int32_t batteryLevel;
	int32_t usbStorage;
	int32_t camera00;
	int32_t camera01;
	int32_t camera02;
	int32_t camera03;
	int32_t accelerometer;
	int32_t fan;
} DiagnosisStats;

typedef struct contrlscreen_ {
	struct timeval manualTime;
	struct timeval gpioTime;
	int32_t screen_manual;
	int32_t screen_gpio;
} ControlScreen;

typedef struct timeval Timeval;

#pragma pack(pop)

size_t getMaxAllocSize(key_t key);

int32_t setWarningApproach(Warning *value);
int32_t getWarningApproach(Warning *value, int32_t camera);

int32_t setWarningCrash(Warning *value);
int32_t getWarningCrash(Warning *value, int32_t camera);

int32_t setWarningCrashRisk(CrashRisk *value);
int32_t getWarningCrashRisk(CrashRisk *value);

int32_t setSplitTime(SplitTime *value);
int32_t getSplitTime(SplitTime *value, int32_t camera);

int32_t setAccelerometerThreshold(AccelerometerThreshold *value);
int32_t getAccelerometerThreshold(AccelerometerThreshold *value);

int32_t setAccelerometer(Accelerometer *value);
int32_t getAccelerometer(Accelerometer *value);

int32_t setRadar(Radar *value);
int32_t getRadar(Radar *value);

int32_t setTegraStats(TegraStats *value);
int32_t getTegraStats(TegraStats *value);

int32_t setPositionPhase(PositionPhase *value);
int32_t getPositionPhase(PositionPhase *value, int32_t camera);

int32_t setTegraStats(TegraStats *value);
int32_t getTegraStats(TegraStats *value);

int32_t setDiagnosisStats(DiagnosisStats *value);
int32_t getDiagnosisStats(DiagnosisStats *value);

int32_t setDiagnosisBattery(DiagnosisStats *value);
int32_t setDiagnosisUsbStorage(DiagnosisStats *value);
int32_t setDiagnosisCamera(DiagnosisStats *value);
int32_t setDiagnosisAccelerometer(DiagnosisStats *value);
int32_t setDiagnosisFan(DiagnosisStats *value);

int32_t setControlScreen(ControlScreen *value);
int32_t getControlScreen(ControlScreen *value);

int32_t setManualScreen(ControlScreen *value);
int32_t setGpioScreen(ControlScreen *value);

*/
import "C"
import "unsafe"

type CTimeval struct {
	Sec  int64
	Usec int64
}

type CWarning struct {
	StartTime  CTimeval
	EvtTime    CTimeval
	Camera     int32
	Event      int32
	FrameIndex int32
	DataLen    uint8
	Data       [128]uint8
}

type CAccelerometerThreshold struct {
	Total int32
	X     int32
	Y     int32
	Z     int32
}

type CAccelerometer struct {
	Tv CTimeval
	X  int16
	Y  int16
	Z  int16
	Vx int16
	Vy int16
	Vz int16
	Vt int16
}

type CRadarUnit struct {
	Id     uint8
	X      int8 // -127 ~ 127 : -12.7m ~ 12.7m
	Y      uint8
	Dist   uint8 // 0 ~ 255 : 0.0m ~ 25.5m
	DSpeed int8  // -127 ~ 127 : -12.7m/s ~ 12.7m/s
	OSpeed int8  // -127 ~ 127 : -12.7m/s ~ 12.7m/s
	Size   uint8
	Flag   uint8
}

type CRadar struct {
	Tv CTimeval
	Id [2]uint8
	U1 CRadarUnit
	U2 CRadarUnit
	U3 CRadarUnit
	U4 CRadarUnit
}

type CCrashRisk struct {
	AlertTime CTimeval
	RecvTime  CTimeval
	SendGwId  [64]uint8
	UwbTagId  uint32
	Speed     uint32
	Count     uint32
}

type CSplitTime struct {
	SplitTime CTimeval
	Fragment  uint32
	Camera    int32
}

type CTegraStats struct {
	Tv      CTimeval
	Ram     uint32
	Swap    uint32
	Cpu     float32
	Iwlwifi float32
	Pmic    float32
	Gpu     float32
	Ao      float32
	Thermal float32
}

type CDiagnosisStats struct {
	TvBattery       CTimeval
	BatteryLevel    int32
	TvUsbStorage    CTimeval
	UsbStorage      int32
	TvCamera        CTimeval
	Camera00        int32
	Camera01        int32
	Camera02        int32
	Camera03        int32
	TvAccelerometer CTimeval
	Accelerometer   int32
	TvFan           CTimeval
	Fan             int32
}

func getMaxAllocSize(key C.key_t) C.size_t {
	return C.getMaxAllocSize(key)
}

/**
 * 접근 감지 시점의 데이터를 공유메모리에 저장한다.
 */
func SetWarningApproach(value *CWarning) C.int32_t {

	warning := C.Warning{}
	warning.startTime = C.Timeval{tv_sec: C.long(value.StartTime.Sec),
		tv_usec: C.long(value.StartTime.Usec)}
	warning.evtTime = C.Timeval{tv_sec: C.long(value.EvtTime.Sec),
		tv_usec: C.long(value.EvtTime.Usec)}
	warning.camera = -1
	warning.event = C.int(value.Event)
	warning.frameIndex = C.int(value.FrameIndex)
	warning.dataLen = C.uint8_t(value.DataLen)

	//var bs []byte = make([]byte, len(value.Data))
	//for i, v := range value.Data {
	//	bs[i] = v
	//}
	//
	//dataPtr := unsafe.Pointer(&warning.data[0])
	//cs := C.CString(string(bs))
	//
	//C.memcpy(dataPtr, unsafe.Pointer(cs), C.ulong(128))
	//
	//C.free(unsafe.Pointer(cs))

	dataPtr := unsafe.Pointer(&warning.data[0])

	C.memcpy(dataPtr, unsafe.Pointer(&value.Data[0]), C.ulong(128))

	return C.setWarningApproach(&warning)
}

/**
 * 접근 감지 데이터를 공유메모리에서 읽는다.
 */
func GetWarningApproach(value *CWarning, camera int32) int32 {

	warning := C.Warning{}

	if C.getWarningApproach(&warning, C.int(camera)) != 0 {
		return -1
	}

	value.StartTime = CTimeval{
		Sec:  int64(warning.startTime.tv_sec),
		Usec: int64(warning.startTime.tv_usec)}

	value.EvtTime = CTimeval{
		Sec:  int64(warning.evtTime.tv_sec),
		Usec: int64(warning.evtTime.tv_usec)}

	value.Camera = int32(warning.camera)
	value.Event = int32(warning.event)
	value.FrameIndex = int32(warning.frameIndex)

	value.DataLen = uint8(warning.dataLen)

	dataPtr := unsafe.Pointer(&warning.data[0])
	bs := C.GoBytes(dataPtr, 128)
	for i, v := range bs {
		value.Data[i] = v
	}

	return 0
}

/**
 * 접근 감지 시점의 데이터를 공유메모리에 저장한다.
 */
func SetSplitTime(value *CSplitTime) C.int32_t {

	splittime := C.SplitTime{}
	splittime.splitTime = C.Timeval{tv_sec: C.long(value.SplitTime.Sec),
		tv_usec: C.long(value.SplitTime.Usec)}
	splittime.fragment = C.uint32_t(value.Fragment)
	splittime.camera = C.int(value.Camera)

	return C.setSplitTime(&splittime)
}

/**
 * 접근 감지 데이터를 공유메모리에서 읽는다.
 */
func GetSplitTime(value *CSplitTime, camera int32) int32 {

	splittime := C.SplitTime{}

	if C.getSplitTime(&splittime, C.int(camera)) != 0 {
		return -1
	}

	value.SplitTime = CTimeval{
		Sec:  int64(splittime.splitTime.tv_sec),
		Usec: int64(splittime.splitTime.tv_usec)}
	value.Fragment = uint32(splittime.fragment)
	value.Camera = int32(splittime.camera)

	return 0
}

/**
 * 접근 감지 시점의 데이터를 공유메모리에 저장한다.
 */
func SetWarningCrashRisk(value *CCrashRisk) C.int32_t {

	warning := C.CrashRisk{}
	warning.alertTime = C.Timeval{tv_sec: C.long(value.AlertTime.Sec),
		tv_usec: C.long(value.AlertTime.Usec)}
	warning.recvTime = C.Timeval{tv_sec: C.long(value.RecvTime.Sec),
		tv_usec: C.long(value.RecvTime.Usec)}

	var bs []byte = make([]byte, len(value.SendGwId))
	for i, v := range value.SendGwId {
		bs[i] = v
	}
	bs[63] = 0 // null char in C

	gwIdPtr := unsafe.Pointer(&warning.sendGwId[0])
	cs := C.CString(string(bs))

	C.memcpy(gwIdPtr, unsafe.Pointer(cs), C.ulong(64))

	C.free(unsafe.Pointer(cs))

	//warning.sendGwId =
	warning.uwbTagId = C.uint32_t(value.UwbTagId)
	warning.speed = C.uint32_t(value.Speed)
	warning.count = C.uint32_t(value.Count)

	return C.setWarningCrashRisk(&warning)
}

/**
 * 접근 감지 데이터를 공유메모리에서 읽는다.
 */
func GetWarningCrashRisk(value *CCrashRisk) int32 {

	warning := C.CrashRisk{}

	if C.getWarningCrashRisk(&warning) != 0 {
		return -1
	}

	value.AlertTime = CTimeval{
		Sec:  int64(warning.alertTime.tv_sec),
		Usec: int64(warning.alertTime.tv_usec)}

	value.RecvTime = CTimeval{
		Sec:  int64(warning.recvTime.tv_sec),
		Usec: int64(warning.recvTime.tv_usec)}

	gwIdPtr := unsafe.Pointer(&warning.sendGwId[0])
	bs := C.GoBytes(gwIdPtr, 64)
	for i, v := range bs {
		value.SendGwId[i] = v
	}
	value.UwbTagId = uint32(warning.uwbTagId)
	value.Speed = uint32(warning.speed)
	value.Count = uint32(warning.count)

	return 0
}

func GetAccelerometerThreshold(value *CAccelerometerThreshold) int32 {

	threshold := C.AccelerometerThreshold{}
	if C.getAccelerometerThreshold(&threshold) != 0 {
		return -1
	}

	value.Total = int32(threshold.total)
	value.X = int32(threshold.x)
	value.Y = int32(threshold.y)
	value.Z = int32(threshold.z)

	return 0
}

func GetAccelerometer(value *CAccelerometer) int32 {

	accelerometer := C.Accelerometer{}
	if C.getAccelerometer(&accelerometer) != 0 {
		return -1
	}

	value.Tv = CTimeval{Sec: int64(accelerometer.tv.tv_sec), Usec: int64(accelerometer.tv.tv_usec)}
	value.X = int16(accelerometer.x)
	value.Y = int16(accelerometer.y)
	value.Z = int16(accelerometer.z)
	value.Vx = int16(accelerometer.vx)
	value.Vy = int16(accelerometer.vy)
	value.Vz = int16(accelerometer.vz)
	value.Vt = int16(accelerometer.vt)

	return 0
}

func GetRadar(value *CRadar) int32 {

	radar := C.Radar{}
	if C.getRadar(&radar) != 0 {
		return -1
	}

	value.Tv = CTimeval{
		Sec:  int64(radar.tv.tv_sec),
		Usec: int64(radar.tv.tv_usec),
	}

	idPtr := unsafe.Pointer(&radar.id[0])
	bs := C.GoBytes(idPtr, 2)
	for i, v := range bs {
		value.Id[i] = v
	}

	value.U1 = CRadarUnit{
		Id:     uint8(radar.u1.id),
		X:      int8(radar.u1.x),
		Y:      uint8(radar.u1.y),
		Dist:   uint8(radar.u1.dist),
		DSpeed: int8(radar.u1.d_speed),
		OSpeed: int8(radar.u1.o_speed),
		Size:   uint8(radar.u1.size),
		Flag:   uint8(radar.u1.flag),
	}
	value.U2 = CRadarUnit{
		Id:     uint8(radar.u2.id),
		X:      int8(radar.u2.x),
		Y:      uint8(radar.u2.y),
		Dist:   uint8(radar.u2.dist),
		DSpeed: int8(radar.u2.d_speed),
		OSpeed: int8(radar.u2.o_speed),
		Size:   uint8(radar.u2.size),
		Flag:   uint8(radar.u2.flag),
	}
	value.U3 = CRadarUnit{
		Id:     uint8(radar.u3.id),
		X:      int8(radar.u3.x),
		Y:      uint8(radar.u3.y),
		Dist:   uint8(radar.u3.dist),
		DSpeed: int8(radar.u3.d_speed),
		OSpeed: int8(radar.u3.o_speed),
		Size:   uint8(radar.u3.size),
		Flag:   uint8(radar.u3.flag),
	}
	value.U4 = CRadarUnit{
		Id:     uint8(radar.u4.id),
		X:      int8(radar.u4.x),
		Y:      uint8(radar.u4.y),
		Dist:   uint8(radar.u4.dist),
		DSpeed: int8(radar.u4.d_speed),
		OSpeed: int8(radar.u4.o_speed),
		Size:   uint8(radar.u4.size),
		Flag:   uint8(radar.u4.flag),
	}

	return 0
}

func SetRadar(value *CRadar) int32 {
	radar := C.Radar{}

	radar.tv = C.Timeval{
		tv_sec:  C.long(value.Tv.Sec),
		tv_usec: C.long(value.Tv.Usec),
	}

	var bs []byte = make([]byte, len(value.Id))
	for i, v := range value.Id {
		bs[i] = v
	}

	idPtr := unsafe.Pointer(&radar.id[0])
	cs := C.CBytes(bs)

	C.memcpy(idPtr, cs, C.ulong(len(bs)))

	C.free(cs)

	radar.u1 = C.RadarUnit{
		id:      C.uint8_t(value.U1.Id),
		x:       C.int8_t(value.U1.X),
		y:       C.uint8_t(value.U1.Y),
		dist:    C.uint8_t(value.U1.Dist),
		d_speed: C.int8_t(value.U1.DSpeed),
		o_speed: C.int8_t(value.U1.OSpeed),
		size:    C.uint8_t(value.U1.Size),
		flag:    C.uint8_t(value.U1.Flag),
	}

	radar.u2 = C.RadarUnit{
		id:      C.uint8_t(value.U2.Id),
		x:       C.int8_t(value.U2.X),
		y:       C.uint8_t(value.U2.Y),
		dist:    C.uint8_t(value.U2.Dist),
		d_speed: C.int8_t(value.U2.DSpeed),
		o_speed: C.int8_t(value.U2.OSpeed),
		size:    C.uint8_t(value.U2.Size),
		flag:    C.uint8_t(value.U2.Flag),
	}

	radar.u3 = C.RadarUnit{
		id:      C.uint8_t(value.U3.Id),
		x:       C.int8_t(value.U3.X),
		y:       C.uint8_t(value.U3.Y),
		dist:    C.uint8_t(value.U3.Dist),
		d_speed: C.int8_t(value.U3.DSpeed),
		o_speed: C.int8_t(value.U3.OSpeed),
		size:    C.uint8_t(value.U3.Size),
		flag:    C.uint8_t(value.U3.Flag),
	}

	radar.u4 = C.RadarUnit{
		id:      C.uint8_t(value.U4.Id),
		x:       C.int8_t(value.U4.X),
		y:       C.uint8_t(value.U4.Y),
		dist:    C.uint8_t(value.U4.Dist),
		d_speed: C.int8_t(value.U4.DSpeed),
		o_speed: C.int8_t(value.U4.OSpeed),
		size:    C.uint8_t(value.U4.Size),
		flag:    C.uint8_t(value.U4.Flag),
	}

	ret := C.setRadar(&radar)

	return int32(ret)
}

/**
 * Jetson Nano 상태 정보를 공유메모리에 저장한다.
 */
func SetTegraStats(value *CTegraStats) C.int32_t {

	tegra := C.TegraStats{}
	tegra.tv = C.Timeval{tv_sec: C.long(value.Tv.Sec),
		tv_usec: C.long(value.Tv.Usec)}
	tegra.ram = C.uint32_t(value.Ram)
	tegra.swap = C.uint32_t(value.Swap)
	tegra.cpu = C.float(value.Cpu)
	tegra.iwlwifi = C.float(value.Iwlwifi)
	tegra.pmic = C.float(value.Pmic)
	tegra.gpu = C.float(value.Gpu)
	tegra.ao = C.float(value.Ao)
	tegra.thermal = C.float(value.Thermal)

	return C.setTegraStats(&tegra)
}

/**
 * 접근 감지 데이터를 공유메모리에서 읽는다.
 */
func GetTegraStats(value *CTegraStats) int32 {

	tegra := C.TegraStats{}

	if C.getTegraStats(&tegra) != 0 {
		return -1
	}

	value.Tv = CTimeval{
		Sec:  int64(tegra.tv.tv_sec),
		Usec: int64(tegra.tv.tv_usec)}

	value.Ram = uint32(tegra.ram)
	value.Swap = uint32(tegra.swap)
	value.Cpu = float32(tegra.cpu)
	value.Iwlwifi = float32(tegra.iwlwifi)
	value.Pmic = float32(tegra.pmic)
	value.Gpu = float32(tegra.gpu)
	value.Ao = float32(tegra.ao)
	value.Thermal = float32(tegra.thermal)

	return 0
}

/**
 * Jetson Nano 상태 정보를 공유메모리에 저장한다.
 */
func SetDiagnosisStats(value *CDiagnosisStats) C.int32_t {

	diag := C.DiagnosisStats{}
	diag.tvBattery = C.Timeval{tv_sec: C.long(value.TvBattery.Sec),
		tv_usec: C.long(value.TvBattery.Usec)}
	diag.batteryLevel = C.int32_t(value.BatteryLevel)
	diag.tvUsbStorage = C.Timeval{tv_sec: C.long(value.TvUsbStorage.Sec),
		tv_usec: C.long(value.TvUsbStorage.Usec)}
	diag.usbStorage = C.int32_t(value.UsbStorage)
	diag.tvCamera = C.Timeval{tv_sec: C.long(value.TvCamera.Sec),
		tv_usec: C.long(value.TvCamera.Usec)}
	diag.camera00 = C.int32_t(value.Camera00)
	diag.camera01 = C.int32_t(value.Camera01)
	diag.camera02 = C.int32_t(value.Camera02)
	diag.camera03 = C.int32_t(value.Camera03)
	diag.tvAccelerometer = C.Timeval{tv_sec: C.long(value.TvAccelerometer.Sec),
		tv_usec: C.long(value.TvAccelerometer.Usec)}
	diag.accelerometer = C.int32_t(value.Accelerometer)

	diag.tvFan = C.Timeval{tv_sec: C.long(value.TvFan.Sec),
		tv_usec: C.long(value.TvFan.Usec)}
	diag.fan = C.int32_t(value.Fan)

	return C.setDiagnosisStats(&diag)
}

/**
 * 접근 감지 데이터를 공유메모리에서 읽는다.
 */
func GetDiagnosisStats(value *CDiagnosisStats) int32 {

	diag := C.DiagnosisStats{}

	if C.getDiagnosisStats(&diag) != 0 {
		return -1
	}

	value.TvBattery = CTimeval{
		Sec:  int64(diag.tvBattery.tv_sec),
		Usec: int64(diag.tvBattery.tv_usec)}
	value.BatteryLevel = int32(diag.batteryLevel)

	value.TvUsbStorage = CTimeval{
		Sec:  int64(diag.tvUsbStorage.tv_sec),
		Usec: int64(diag.tvUsbStorage.tv_usec)}
	value.UsbStorage = int32(diag.usbStorage)

	value.TvCamera = CTimeval{
		Sec:  int64(diag.tvCamera.tv_sec),
		Usec: int64(diag.tvCamera.tv_usec)}
	value.Camera00 = int32(diag.camera00)
	value.Camera01 = int32(diag.camera01)
	value.Camera02 = int32(diag.camera02)
	value.Camera03 = int32(diag.camera03)

	value.TvAccelerometer = CTimeval{
		Sec:  int64(diag.tvAccelerometer.tv_sec),
		Usec: int64(diag.tvAccelerometer.tv_usec)}
	value.Accelerometer = int32(diag.accelerometer)

	value.TvFan = CTimeval{
		Sec:  int64(diag.tvFan.tv_sec),
		Usec: int64(diag.tvFan.tv_usec)}
	value.Fan = int32(diag.fan)

	return 0
}

func SetDiagnosisBattery(value *CDiagnosisStats) int32 {
	diag := CDiagnosisStats{}
	if GetDiagnosisStats(&diag) == 0 {
		diag.TvBattery = CTimeval{
			Sec:  int64(value.TvBattery.Sec),
			Usec: int64(value.TvBattery.Usec)}
		diag.BatteryLevel = int32(value.BatteryLevel)

		return int32(SetDiagnosisStats(&diag))
	}

	return -1
}

func SetDiagnosisUsbStorage(value *CDiagnosisStats) int32 {
	diag := CDiagnosisStats{}
	if GetDiagnosisStats(&diag) == 0 {
		diag.TvUsbStorage = CTimeval{
			Sec:  int64(value.TvUsbStorage.Sec),
			Usec: int64(value.TvUsbStorage.Usec)}
		diag.UsbStorage = int32(value.UsbStorage)

		return int32(SetDiagnosisStats(&diag))
	}

	return -1
}

func SetDiagnosisCamera(value *CDiagnosisStats) int32 {
	diag := CDiagnosisStats{}
	if GetDiagnosisStats(&diag) == 0 {
		diag.TvCamera = CTimeval{
			Sec:  int64(value.TvCamera.Sec),
			Usec: int64(value.TvCamera.Usec)}
		diag.Camera00 = int32(value.Camera00)
		diag.Camera01 = int32(value.Camera01)
		diag.Camera02 = int32(value.Camera02)
		diag.Camera03 = int32(value.Camera03)

		return int32(SetDiagnosisStats(&diag))
	}

	return -1
}

func SetDiagnosisAccelerometer(value *CDiagnosisStats) int32 {
	diag := CDiagnosisStats{}
	if GetDiagnosisStats(&diag) == 0 {
		diag.TvAccelerometer = CTimeval{
			Sec:  int64(value.TvAccelerometer.Sec),
			Usec: int64(value.TvAccelerometer.Usec)}
		diag.Accelerometer = int32(value.Accelerometer)

		return int32(SetDiagnosisStats(&diag))
	}

	return -1
}

func SetDiagnosisFan(value *CDiagnosisStats) int32 {
	diag := CDiagnosisStats{}
	if GetDiagnosisStats(&diag) == 0 {
		diag.TvFan = CTimeval{
			Sec:  int64(value.TvFan.Sec),
			Usec: int64(value.TvFan.Usec)}
		diag.Fan = int32(value.Fan)

		return int32(SetDiagnosisStats(&diag))
	}

	return -1
}
