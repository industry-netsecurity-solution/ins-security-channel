#ifndef __UTILS_SHM_H__
#define __UTILS_SHM_H__

#include <sys/time.h>
#include <stdint.h>
#include <sys/ipc.h>
#include <sys/shm.h>


#define SK_ACCELEROMETER           20190902 // 0x013416B6
#define SK_POSITIONPHASE_00        20201010 // 0x01343E32
#define SK_POSITIONPHASE_01        20201011 // 0x01343E33
#define SK_ACCELEROMETER_THRESHOLD 20200301 // 0x01343B6D
#define SK_EVENT_APPROACH_00       20200800 // 0x01343D60
#define SK_EVENT_APPROACH_01       20200801 // 0x01343D61
#define SK_EVENT_APPROACH_FF       20200899 //
#define SK_EVENT_CRASH_00          20200802 // 0x01343D62
#define SK_EVENT_CRASH_01          20200803 // 0x01343D63
#define SK_EVENT_CRASH_RISK        20211118 // 0x013465AE
#define SK_EVENT_SPLIT_TIME_00     20211200 // 0x01346600
#define SK_EVENT_SPLIT_TIME_01     20211201 // 0x01346601

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

#pragma pack(pop)

size_t getMaxAllocSize(key_t key);

/**
 * 접근 감지 시점의 데이터를 공유메모리에 읽고/쓴다.
 */
int32_t setWarningApproach(Warning *value);
int32_t getWarningApproach(Warning *value, int32_t camera);

/**
 * 충돌 시점의 데이터를 공유메모리에 읽고/쓴다.
 */
int32_t setWarningCrash(Warning *value);
int32_t getWarningCrash(Warning *value, int32_t camera);

/**
 * 충돌 위험 시점의 데이터를 공유메모리에 읽고/쓴다.
 */
int32_t setWarningCrashRisk(CrashRisk *value);
int32_t getWarningCrashRisk(CrashRisk *value);

/**
 * 영상 파일 분리 시점의 시간과 fragment 를 공유메모리에 읽고/쓴다.
 */
int32_t setSplitTime(SplitTime *value);
int32_t getSplitTime(SplitTime *value, int32_t camera);

/**
 * accelerometer의 threshold값을 공유메모리에 읽고/쓴다.
 */
int32_t setAccelerometerThreshold(AccelerometerThreshold *value);
int32_t getAccelerometerThreshold(AccelerometerThreshold *value);

/**
 * 가속도 값을 공유메모리에 읽고/쓴다.
 */
int32_t setAccelerometer(Accelerometer *value);
int32_t getAccelerometer(Accelerometer *value);

/**
 * 접근 감지 비율을 공유메모리에 읽고/쓴다.
 */
int32_t setPositionPhase(PositionPhase *value);
int32_t getPositionPhase(PositionPhase *value, int32_t camera);

#endif
