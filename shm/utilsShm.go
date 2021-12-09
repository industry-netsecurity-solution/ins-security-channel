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
