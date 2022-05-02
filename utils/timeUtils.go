package utils

import (
	"errors"
	"github.com/ml444/gid"
	"strconv"
	"time"
)

// type TimeUtils struct {
// 	EPOCH = 1610351662000
// }
const EPOCH = int64(1610351662000) //起始时间，预计可用34+34年

func ValidateTimestamp(lastTimestamp uint64, timestamp uint64) bool {
	if timestamp < lastTimestamp {
		// if (log.isErrorEnabled())
		// 	log.error(String
		// 			.format("Clock moved backwards.  Refusing to generate id for %d second/milisecond.",
		// 					lastTimestamp - timestamp));

		// throw new IllegalStateException(
		// 		String.format(
		// 				"Clock moved backwards.  Refusing to generate id for %d second/milisecond.",
		// 				lastTimestamp - timestamp));
		return false
	}
	return true
}

func TillNextTimeUnit(lastTimestamp uint64, idType uint64) uint64 {
	// if (log.isInfoEnabled())
	// 	log.info(String
	// 			.format("Ids are used out during %d. Waiting till next second/milisencond.",
	// 					lastTimestamp))

	timestamp := GenTime(idType)
	for timestamp <= lastTimestamp {
		timestamp = GenTime(idType)
	}

	// if (log.isInfoEnabled())
	// 	log.info(String.format("Next second/milisencond %d is up.",
	// 			timestamp))

	return timestamp
}

func GenTime(idType uint64) uint64 {
	// 通过EPOCH对时间压缩
	if idType == gid.MaxPeak {
		return uint64((time.Now().UnixNano()/1000000 - EPOCH) / 1000)
	} else if idType == gid.MinGranularity {
		return uint64(time.Now().UnixNano()/1000000 - EPOCH)
	} else {
		return uint64((time.Now().UnixNano()/1000000 - EPOCH) / 1000)
	}
	// return (1509238640744 - EPOCH) / 1000
}

func TransDuration(timestamp int64) (uint64, error) {
	timeStr := strconv.FormatInt(timestamp, 10)
	if len(timeStr) == 10 {
		return uint64((timestamp*1000 - EPOCH) / 1000), nil
	} else if len(timeStr) == 13 {
		return uint64((timestamp - EPOCH) / 1000), nil
	} else {
		return 0, errors.New("the timestamp is invalided")
	}
}

func TransTime(timeDuration, idType uint64) (int64, string) {
	// 从时间间隔中计算出时间戳
	//return timeDuration
	var timestamp int64
	if idType == gid.MaxPeak {
		timestamp = int64(timeDuration*1000) + EPOCH
		//fmt.Println(timestamp, "MAX_PEAK")
	} else {
		timestamp = int64(timeDuration) + EPOCH
		//fmt.Println(timestamp, "MIN_GRANULARITY")
	}
	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000000
	return timestamp, time.Unix(sec, nsec).Format("2006-01-02 15:04:05.000")
}
