#ifndef __UTILS_SHM_H__
#define __UTILS_SHM_H__

#include <sys/time.h>
#include <stdint.h>
#include <sys/ipc.h>
#include <sys/shm.h>


#define SK_ACCELEROMETER           20190902 // 0x013416B6
#define SK_POSITIONPHASE_00        20201010 // 0x01343E32
#define SK_POSITIONPHASE_01        20201011 // 0x01343E33
#define SK_POSITIONPHASE_02        20200902 // 0x01348BE6
#define SK_POSITIONPHASE_03        20200903 // 0x01348BE7
#define SK_ACCELEROMETER_THRESHOLD 20200301 // 0x01343B6D
#define SK_EVENT_APPROACH_00       20200800 // 0x01343D60
#define SK_EVENT_APPROACH_01       20200801 // 0x01343D61
#define SK_EVENT_APPROACH_02       20220900 // 0x01348BE4
#define SK_EVENT_APPROACH_03       20200901 // 0x01348BE5
#define SK_EVENT_APPROACH_FF       20200899 //
#define SK_EVENT_CRASH_00          20200802 // 0x01343D62
#define SK_EVENT_CRASH_01          20200803 // 0x01343D63
#define SK_EVENT_CRASH_02          20200904 // 0x01348BE8
#define SK_EVENT_CRASH_03          20200905 // 0x01348BE9
#define SK_EVENT_CRASH_RISK        20211118 // 0x013465AE
#define SK_EVENT_SPLIT_TIME_00     20211200 // 0x01346600
#define SK_EVENT_SPLIT_TIME_01     20211201 // 0x01346601
#define SK_EVENT_SPLIT_TIME_02     20200906 // 0x01348BEA
#define SK_EVENT_SPLIT_TIME_03     20200907 // 0x01348BEB
#define SK_TEGRASTATS              20220301 // 0x0134898D
#define SK_DIAGNOSISSTATS          20220330 // 0x013489AA
#define SK_CONTROL_SCREEN          20220603 // 0x01348ABB
#define SK_RADAR                   20221025 // 0x01348C61

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
	uint8_t x;
	uint8_t y;
	uint8_t dist;
	uint8_t d_speed;
	uint8_t o_speed;
	uint8_t size;
	uint8_t flag;
} RadarUnit;

typedef struct radar {
	struct timeval tv;
	uint8_t   id;
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
 * 레이다 접근 감지 센서 값을 공유메모리에 읽고/쓴다.
 */
int32_t setRadar(Radar *value);
int32_t getRadar(Radar *value);

/**
 * 접근 감지 비율을 공유메모리에 읽고/쓴다.
 */
int32_t setPositionPhase(PositionPhase *value);
int32_t getPositionPhase(PositionPhase *value, int32_t camera);

/**
 * Jetson Nano 상태 정보를 공유메모리에 읽고/쓴다.
 */
int32_t setTegraStats(TegraStats *value);
int32_t getTegraStats(TegraStats *value);

/**
 * 상태 정보를 공유메모리에 읽고/쓴다.
 */
int32_t setDiagnosisStats(DiagnosisStats *value);
int32_t getDiagnosisStats(DiagnosisStats *value);

int32_t setDiagnosisBattery(DiagnosisStats *value);
int32_t setDiagnosisUsbStorage(DiagnosisStats *value);
int32_t setDiagnosisCamera(DiagnosisStats *value);
int32_t setDiagnosisAccelerometer(DiagnosisStats *value);
int32_t setDiagnosisFan(DiagnosisStats *value);

/**
 * 화면 출력 상태를 조정한다.
 */
int32_t setControlScreen(ControlScreen *value);
int32_t getControlScreen(ControlScreen *value);

/*
 * 사용자가 터치 스크린으로 조작한 화면 방향
 */
int32_t setManualScreen(ControlScreen *value);
/*
 * GPIO에 의해 조작한 화면 방향
 */
int32_t setGpioScreen(ControlScreen *value);

#endif
