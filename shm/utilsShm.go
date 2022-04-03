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

typedef struct segraStats_ {
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

*/
import "C"
import "unsafe"

type CTimeval struct {
	Sec int64
	Usec int64
}

type CWarning struct {
	StartTime CTimeval
	EvtTime CTimeval
	Camera int32
	Event int32
	FrameIndex int32
	DataLen uint8;
	Data [128]uint8;
}

type CAccelerometerThreshold struct {
	Total int32
	X int32
	Y int32
	Z int32
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

type CCrashRisk struct {
	AlertTime     CTimeval
	RecvTime      CTimeval
	SendGwId  [64]uint8
	UwbTagId      uint32
	Speed         uint32
	Count         uint32
}

type CSplitTime struct {
	SplitTime CTimeval
	Fragment uint32
	Camera int32
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
	warning.startTime = C.Timeval{ tv_sec: C.long(value.StartTime.Sec),
									tv_usec: C.long(value.StartTime.Usec)}
	warning.evtTime = C.Timeval{ tv_sec: C.long(value.EvtTime.Sec),
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
		Sec: int64(warning.startTime.tv_sec),
		Usec: int64(warning.startTime.tv_usec)}

	value.EvtTime = CTimeval{
		Sec: int64(warning.evtTime.tv_sec),
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
	splittime.splitTime = C.Timeval{ tv_sec: C.long(value.SplitTime.Sec),
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
		Sec: int64(splittime.splitTime.tv_sec),
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
	warning.alertTime = C.Timeval{ tv_sec: C.long(value.AlertTime.Sec),
		tv_usec: C.long(value.AlertTime.Usec)}
	warning.recvTime = C.Timeval{ tv_sec: C.long(value.RecvTime.Sec),
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
		Sec: int64(warning.alertTime.tv_sec),
		Usec: int64(warning.alertTime.tv_usec)}

	value.RecvTime = CTimeval{
		Sec: int64(warning.recvTime.tv_sec),
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

	value.Tv = CTimeval{ Sec: int64(accelerometer.tv.tv_sec), Usec: int64(accelerometer.tv.tv_usec)}
	value.X = int16(accelerometer.x)
	value.Y = int16(accelerometer.y)
	value.Z = int16(accelerometer.z)
	value.Vx = int16(accelerometer.vx)
	value.Vy = int16(accelerometer.vy)
	value.Vz = int16(accelerometer.vz)
	value.Vt = int16(accelerometer.vt)

	return 0
}


/**
 * Jetson Nano 상태 정보를 공유메모리에 저장한다.
 */
func SetTegraStats(value *CTegraStats) C.int32_t {

	tegra := C.TegraStats{}
	tegra.tv = C.Timeval{ tv_sec: C.long(value.Tv.Sec),
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
		Sec: int64(tegra.tv.tv_sec),
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
	diag.tvBattery = C.Timeval{ tv_sec: C.long(value.TvBattery.Sec),
		tv_usec: C.long(value.TvBattery.Usec)}
	diag.batteryLevel = C.int32_t(value.BatteryLevel)
	diag.tvUsbStorage = C.Timeval{ tv_sec: C.long(value.TvUsbStorage.Sec),
		tv_usec: C.long(value.TvUsbStorage.Usec)}
	diag.usbStorage = C.int32_t(value.UsbStorage)
	diag.tvCamera = C.Timeval{ tv_sec: C.long(value.TvCamera.Sec),
		tv_usec: C.long(value.TvCamera.Usec)}
	diag.camera00 = C.int32_t(value.Camera00)
	diag.camera01 = C.int32_t(value.Camera01)
	diag.camera02 = C.int32_t(value.Camera02)
	diag.camera03 = C.int32_t(value.Camera03)
	diag.tvAccelerometer = C.Timeval{ tv_sec: C.long(value.TvAccelerometer.Sec),
		tv_usec: C.long(value.TvAccelerometer.Usec)}
	diag.accelerometer = C.int32_t(value.Accelerometer)

	diag.tvFan = C.Timeval{ tv_sec: C.long(value.TvFan.Sec),
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
		Sec: int64(diag.tvBattery.tv_sec),
		Usec: int64(diag.tvBattery.tv_usec)}
	value.BatteryLevel = int32(diag.batteryLevel)

	value.TvUsbStorage = CTimeval{
		Sec: int64(diag.tvUsbStorage.tv_sec),
		Usec: int64(diag.tvUsbStorage.tv_usec)}
	value.UsbStorage = int32(diag.usbStorage)

	value.TvCamera = CTimeval{
		Sec: int64(diag.tvCamera.tv_sec),
		Usec: int64(diag.tvCamera.tv_usec)}
	value.Camera00 = int32(diag.camera00)
	value.Camera01 = int32(diag.camera01)
	value.Camera02 = int32(diag.camera02)
	value.Camera03 = int32(diag.camera03)

	value.TvAccelerometer = CTimeval{
		Sec: int64(diag.tvAccelerometer.tv_sec),
		Usec: int64(diag.tvAccelerometer.tv_usec)}
	value.Accelerometer = int32(diag.accelerometer)

	value.TvFan = CTimeval{
		Sec: int64(diag.tvFan.tv_sec),
		Usec: int64(diag.tvFan.tv_usec)}
	value.Fan = int32(diag.fan)

	return 0
}
