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

#pragma pack(pop)

size_t getMaxAllocSize(key_t key);

/**
 * 접근 감지 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setWarningApproach(Warning *value);
int32_t getWarningApproach(Warning *value, int32_t camera);

/**
 * 충돌 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setWarningCrash(Warning *value);
int32_t getWarningCrash(Warning *value, int32_t camera);

/**
 * accelerometer의 threshold값을 공유메모리에 저장한다.
 */
int32_t setAccelerometerThreshold(AccelerometerThreshold *value);

/**
 * 공유메모리에서 accelerometer의 값을 읽어온다.
 */
int32_t getAccelerometerThreshold(AccelerometerThreshold *value);

/**
 * 가속도 값을 공유메모리에 저장한다.
 */
int32_t setAccelerometer(Accelerometer *value);

/**
 * 공유메모리에서 가속도 값을 읽어온다.
 */
int32_t getAccelerometer(Accelerometer *value);

/**
 * 접근 감지 비율을 공유메모리에 저장한다.
 */
int32_t setPositionPhase(PositionPhase *value);

/**
 * 공유메모리에 저장된 1,2 단계의 접근 감지 비율을 가져온다.
 * 접근 감지를 판단한다..
 */
int32_t getPositionPhase(PositionPhase *value, int32_t camera);

#endif
