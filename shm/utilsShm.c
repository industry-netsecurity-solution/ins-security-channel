
#include <limits.h>
#include <string.h>
#include "utilsShm.h"

size_t getMaxAllocSize(key_t key) {
	if(key == (key_t)SK_ACCELEROMETER) {
		return 256;
	} else if(key == (key_t) SK_POSITIONPHASE_00) {
		return 256;
	} else if(key == (key_t) SK_POSITIONPHASE_01) {
		return 256;
	} else if(key == (key_t) SK_POSITIONPHASE_02) {
		return 256;
	} else if(key == (key_t) SK_POSITIONPHASE_03) {
		return 256;
	} else if(key == (key_t) SK_ACCELEROMETER_THRESHOLD) {
		return 256;
	} else if(key == (key_t) SK_EVENT_APPROACH_00) {
		return 256;
	} else if(key == (key_t) SK_EVENT_APPROACH_01) {
		return 256;
	} else if(key == (key_t) SK_EVENT_APPROACH_02) {
		return 256;
	} else if(key == (key_t) SK_EVENT_APPROACH_03) {
		return 256;
	} else if(key == (key_t) SK_EVENT_APPROACH_FF) {
		return 256;
	} else if(key == (key_t) SK_EVENT_CRASH_00) {
		return 256;
	} else if(key == (key_t) SK_EVENT_CRASH_01) {
		return 256;
	} else if(key == (key_t) SK_EVENT_CRASH_02) {
		return 256;
	} else if(key == (key_t) SK_EVENT_CRASH_03) {
		return 256;
	} else if(key == (key_t) SK_EVENT_CRASH_RISK) {
		return 256;
	} else if(key == (key_t) SK_EVENT_SPLIT_TIME_00) {
		return 256;
	} else if(key == (key_t) SK_EVENT_SPLIT_TIME_01) {
		return 256;
	} else if(key == (key_t) SK_EVENT_SPLIT_TIME_02) {
		return 256;
	} else if(key == (key_t) SK_EVENT_SPLIT_TIME_03) {
		return 256;
	} else if(key == (key_t) SK_TEGRASTATS) {
		return 256;
	} else if(key == (key_t) SK_DIAGNOSISSTATS) {
		return 256;
	} else if(key == (key_t) SK_CONTROL_SCREEN) {
		return 256;
	} else if(key == (key_t) SK_RADAR) {
		return 256;
	}

	return 0;
}

/**
 * 접근 감지 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setWarningApproach(Warning *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Warning);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(value->camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_00);
	} else if(value->camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_01);
	} else if(value->camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_02);
	} else if(value->camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_03);
	} else if(value->camera == -1) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_FF);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(value->camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == -1) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_FF, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(Warning));

	shmdt(shmPtr);
/*
	printf("====> APPROACH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/

	return 0;
}

int32_t getWarningApproach(Warning *value, int32_t camera) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Warning);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_00);
	} else if(camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_01);
	} else if(camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_02);
	} else if(camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_03);
	} else if(camera == -1) {
		size = getMaxAllocSize((key_t)SK_EVENT_APPROACH_FF);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == -1) {
		if((shmId = shmget((key_t)SK_EVENT_APPROACH_FF, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(Warning));
/*
	printf("====< APPROACH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/
	shmdt(shmPtr);

	return 0;
}

/**
 * 충돌 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setWarningCrash(Warning *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Warning);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(value->camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_00);
	} else if(value->camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_01);
	} else if(value->camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_02);
	} else if(value->camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(value->camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(Warning));

	shmdt(shmPtr);
/*
	printf("====> CRASH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/

	return 0;
}

int32_t getWarningCrash(Warning *value, int32_t camera) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Warning);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_00);
	} else if(camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_01);
	} else if(camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_02);
	} else if(camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_CRASH_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_CRASH_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(Warning));

	shmdt(shmPtr);
/*
	printf("====< CRASH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/
	return 0;
}

/**
 * 충돌 위험 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setWarningCrashRisk(CrashRisk *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(CrashRisk);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_EVENT_CRASH_RISK);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_EVENT_CRASH_RISK, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(CrashRisk));

	shmdt(shmPtr);
/*
	printf("====> CRASH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/

	return 0;
}

int32_t getWarningCrashRisk(CrashRisk *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(CrashRisk);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_EVENT_CRASH_RISK);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_EVENT_CRASH_RISK, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(CrashRisk));

	shmdt(shmPtr);
/*
	printf("====< CRASH: %ld.%ld %d, %d, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->camera, value->event, value->frameIndex);
*/
	return 0;
}

/**
 * 충돌 시점의 데이터를 공유메모리에 저장한다.
 */
int32_t setSplitTime(SplitTime *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(SplitTime);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(value->camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_00);
	} else if(value->camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_01);
	} else if(value->camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_02);
	} else if(value->camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(value->camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(SplitTime));

	shmdt(shmPtr);
/*
	printf("====< Split Time: %ld.%ld %u, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->fragment, value->camera);
*/

	return 0;
}

int32_t getSplitTime(SplitTime *value, int32_t camera) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(SplitTime);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(camera == 0) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_00);
	} else if(camera == 1) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_01);
	} else if(camera == 2) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_02);
	} else if(camera == 3) {
		size = getMaxAllocSize((key_t)SK_EVENT_SPLIT_TIME_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(camera == 0) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 1) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 2) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 3) {
		if((shmId = shmget((key_t)SK_EVENT_SPLIT_TIME_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(SplitTime));

	shmdt(shmPtr);
/*
	printf("====< Split Time: %ld.%ld %u, %d\n", value->evtTime.tv_sec, value->evtTime.tv_usec,
			value->fragment, value->camera);
*/
	return 0;
}

/**
 * accelerometer의 threshold값을 공유메모리에 저장한다.
 */
int32_t setAccelerometerThreshold(AccelerometerThreshold *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(AccelerometerThreshold);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_ACCELEROMETER_THRESHOLD);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_ACCELEROMETER_THRESHOLD, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(AccelerometerThreshold));

	shmdt(shmPtr);

	return 0;
}


/**
 * 공유메모리에서 accelerometer의 값을 읽어온다.
 */
int32_t getAccelerometerThreshold(AccelerometerThreshold *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(AccelerometerThreshold);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_ACCELEROMETER_THRESHOLD);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_ACCELEROMETER_THRESHOLD, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(AccelerometerThreshold));

	shmdt(shmPtr);

	return 0;

}


/**
 * 가속도 값을 공유메모리에 저장한다.
 */
int32_t setAccelerometer(Accelerometer *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Accelerometer);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_ACCELEROMETER);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_ACCELEROMETER, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(Accelerometer));

	shmdt(shmPtr);

	return 0;
}

/**
 * 공유메모리에서 가속도 값을 읽어온다.
 */
int32_t getAccelerometer(Accelerometer *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Accelerometer);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_ACCELEROMETER);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_ACCELEROMETER, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(Accelerometer));

	shmdt(shmPtr);

	return 0;
}

/**
 * 레이다 접근 감지 센서 값을 공유메모리에 저장한다.
 */
int32_t setRadar(Radar *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Radar);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_RADAR);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_RADAR, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(Radar));

	shmdt(shmPtr);

	return 0;
}

/**
 * 레이다 접근 감지 센서 값을 읽어온다.
 */
int32_t getRadar(Radar *value) {
	int shmId;
	int8_t *shmPtr;

	size_t shm_size = sizeof(Radar);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_RADAR);

	if(shm_size < size) {
		shm_size = size;
	}

	if((shmId = shmget((key_t)SK_RADAR, shm_size, IPC_CREAT|0666)) == -1) {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(value, shmPtr, sizeof(Radar));

	shmdt(shmPtr);

	return 0;
}

/**
 * 접근 감지 비율을 공유메모리에 저장한다.
 */
int32_t setPositionPhase(PositionPhase *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(PositionPhase);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(value->camera == 0) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_00);
	} else if(value->camera == 1) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_01);
	} else if(value->camera == 2) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_02);
	} else if(value->camera == 3) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(value->camera == 0) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 1) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 2) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(value->camera == 3) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(PositionPhase));

	shmdt(shmPtr);

	return 0;
}

/**
 * 공유메모리에 저장된 1,2 단계의 접근 감지 비율을 가져온다.
 * 접근 감지를 판단한다..
 */
int32_t getPositionPhase(PositionPhase *value, int32_t camera) {
	int shmId;
	int8_t *shmPtr;
	PositionPhase *pp;
	size_t shm_size = sizeof(PositionPhase);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	if(camera == 0) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_00);
	} else if(camera == 1) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_01);
	} else if(camera == 2) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_02);
	} else if(camera == 3) {
		size = getMaxAllocSize((key_t)SK_POSITIONPHASE_03);
	} else {
		return -1;
	}

	if(shm_size < size) {
		shm_size = size;
	}

	if(camera == 0) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_00, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 1) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_01, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 2) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_02, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else if(camera == 3) {
		if((shmId = shmget((key_t)SK_POSITIONPHASE_03, shm_size, IPC_CREAT|0666)) == -1) {
			return -1;
		}
	} else {
		return -1;
	}

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	// 공유메모리 주소
	pp = (PositionPhase *)shmPtr;
	if(pp->update != 1) {
		memcpy(value, pp, sizeof(PositionPhase));

		shmdt(shmPtr);
		return 0;
	}

	// 공유 메모리의 d1, d2
	double tmp1 = pp->p1;
	double tmp2 = pp->p2;

	// d1의 범위 확인
	if(tmp1 < 0.0 || 1.0 < tmp1) {
		shmdt(shmPtr);
		return -1;
	}

	// d2의 범위 확인
	if(tmp2 < 0.0 || 1.0 < tmp2) {
		shmdt(shmPtr);
		return -1;
	}

	// d1, d2의 값을 비교하여, 작은 값을 d1, 큰 값을 d2로
	if(tmp2 < tmp1) {
		double tmp = tmp1;
		tmp1 = tmp2;
		tmp2 = tmp;
	}

	value->update = pp->update;
	value->camera = pp->camera;
	value->p1 = tmp1;
	value->p2 = tmp2;

	pp->update = 0;
	//pp->camera = pp->camera;
	pp->p1 = 0.0;
	pp->p2 = 0.0;

	shmdt(shmPtr);

	return 1;
}

/**
 * etson Nano 상태 정보를 공유메모리에 저장한다.
 */
int32_t setTegraStats(TegraStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(TegraStats);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_TEGRASTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_TEGRASTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(TegraStats));

	shmdt(shmPtr);

	return 0;
}


/**
 * 공유메모리에 저장된 etson Nano 상태 정보를 가져온다.
 */
int32_t getTegraStats(TegraStats *value) {
	int shmId;
	int8_t *shmPtr;
	TegraStats *pp;
	size_t shm_size = sizeof(TegraStats);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_TEGRASTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_TEGRASTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	// 공유메모리 주소
	pp = (TegraStats *)shmPtr;
	memcpy(value, pp, sizeof(TegraStats));

    shmdt(shmPtr);
    return 0;
}

/**
 * 상태 정보를 공유메모리에 저장한다.
 */
int32_t setDiagnosisStats(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(DiagnosisStats));

	shmdt(shmPtr);

	return 0;
}


/**
 * 공유메모리에 저장된 상태 정보를 가져온다.
 */
int32_t getDiagnosisStats(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	DiagnosisStats *pp;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	// 공유메모리 주소
	pp = (DiagnosisStats *)shmPtr;
	memcpy(value, pp, sizeof(DiagnosisStats));

    shmdt(shmPtr);
    return 0;
}

int32_t setDiagnosisBattery(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;
	DiagnosisStats *diag;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	diag = (DiagnosisStats *)shmPtr;
	diag->tvBattery = value->tvBattery;
	diag->batteryLevel = value->batteryLevel;

	shmdt(shmPtr);

	return 0;
}
int32_t setDiagnosisUsbStorage(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;
	DiagnosisStats *diag;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	diag = (DiagnosisStats *)shmPtr;
	diag->tvUsbStorage = value->tvUsbStorage;
	diag->usbStorage = value->usbStorage;

	shmdt(shmPtr);

	return 0;
}

int32_t setDiagnosisCamera(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;
	DiagnosisStats *diag;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	diag = (DiagnosisStats *)shmPtr;
	diag->tvCamera = value->tvCamera;
	diag->camera00 = value->camera00;
	diag->camera01 = value->camera01;
	diag->camera02 = value->camera02;
	diag->camera03 = value->camera03;

	shmdt(shmPtr);

	return 0;
}

int32_t setDiagnosisAccelerometer(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;
	DiagnosisStats *diag;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	diag = (DiagnosisStats *)shmPtr;
	diag->tvAccelerometer = value->tvAccelerometer;
	diag->accelerometer = value->accelerometer;

	shmdt(shmPtr);

	return 0;
}

int32_t setDiagnosisFan(DiagnosisStats *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(DiagnosisStats);
	size_t size = 0;
	DiagnosisStats *diag;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_DIAGNOSISSTATS);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_DIAGNOSISSTATS, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	diag = (DiagnosisStats *)shmPtr;
	diag->tvFan = value->tvFan;
	diag->fan = value->fan;

	shmdt(shmPtr);

	return 0;
}

/**
 * 화면 출력 상태를 조정한다.
 */
int32_t setControlScreen(ControlScreen *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(ControlScreen);
	size_t size = 0;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_CONTROL_SCREEN);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_CONTROL_SCREEN, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	memcpy(shmPtr, value, sizeof(ControlScreen));

	shmdt(shmPtr);

	return 0;
}

int32_t getControlScreen(ControlScreen *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(ControlScreen);
	size_t size = 0;
	ControlScreen *pp;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_CONTROL_SCREEN);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_CONTROL_SCREEN, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	// 공유메모리 주소
	pp = (ControlScreen *)shmPtr;
	memcpy(value, pp, sizeof(ControlScreen));

    shmdt(shmPtr);
    return 0;
}

/*
 * 사용자가 터치 스크린으로 조작한 화면 방향
 */
int32_t setManualScreen(ControlScreen *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(ControlScreen);
	size_t size = 0;
	ControlScreen *pp;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_CONTROL_SCREEN);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_CONTROL_SCREEN, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	pp = (ControlScreen *)shmPtr;
	pp->manualTime = value->manualTime;
	pp->screen_manual = value->screen_manual;

	shmdt(shmPtr);
    return 0;
}

/*
 * GPIO에 의해 조작한 화면 방향
 */
int32_t setGpioScreen(ControlScreen *value) {
	int shmId;
	int8_t *shmPtr;
	size_t shm_size = sizeof(ControlScreen);
	size_t size = 0;
	ControlScreen *pp;

	if(value == NULL) {
		return -1;
	}

	size = getMaxAllocSize((key_t)SK_CONTROL_SCREEN);

	if(shm_size < size) {
		shm_size = size;
	}

    if((shmId = shmget((key_t)SK_CONTROL_SCREEN, shm_size, IPC_CREAT|0666)) == -1) {
        return -1;
    }

	if((shmPtr = (int8_t *)shmat(shmId, (const void *)NULL, 0)) == (void *)-1) {
		return -1;
	}

	pp = (ControlScreen *)shmPtr;
	pp->gpioTime = value->gpioTime;
	pp->screen_gpio = value->screen_gpio;

	shmdt(shmPtr);
    return 0;
}

