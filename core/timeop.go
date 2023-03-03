package core

import (
	"errors"
	"strconv"
	"time"
)

//const EPOCH = int64(1610351662000) //起始时间，预计可用34+34年

type ITimeOp interface {
	TimeNow() uint64
	ParseDuration(duration uint64) (int64, string)
	ParseTimestampToDuration(ts int64) (uint64, error)
	ValidateTimestamp(lastTs, nowTs uint64) bool
	WaitNextTime(lastTs uint64) uint64
}

var _ ITimeOp = &TimeOp{}

type TimeOp struct {
	epoch         int64
	genTimeFunc   func() uint64
	transTimeFunc func(duration uint64) (int64, string)
}

func (t *TimeOp) getTimeNow4Second() uint64 {
	return uint64((time.Now().UnixNano()/1000000 - t.epoch) / 1000)
}
func (t *TimeOp) getTimeNow4Milli() uint64 {
	return uint64(time.Now().UnixNano()/1000000 - t.epoch)
}
func (t *TimeOp) TimeNow() uint64 {
	return t.genTimeFunc()
}

func (t *TimeOp) transTimeWithMaxPeak(timeDuration uint64) (int64, string) {
	// 从时间间隔中计算出时间戳
	//return timeDuration
	timestamp := int64(timeDuration*1000) + t.epoch
	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000000
	return timestamp, time.Unix(sec, nsec).Format("2006-01-02 15:04:05.000")
}
func (t *TimeOp) transTimeWithMinGranularity(timeDuration uint64) (int64, string) {
	// 从时间间隔中计算出时间戳
	timestamp := int64(timeDuration) + t.epoch
	sec := timestamp / 1000
	nsec := timestamp % 1000 * 1000000
	return timestamp, time.Unix(sec, nsec).Format("2006-01-02 15:04:05.000")
}
func (t *TimeOp) ParseDuration(duration uint64) (int64, string) {
	return t.transTimeFunc(duration)
}

func (t *TimeOp) ValidateTimestamp(lastTimestamp uint64, timestamp uint64) bool {
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
func (t *TimeOp) WaitNextTime(lastTimestamp uint64) uint64 {
	// if (log.isInfoEnabled())
	// 	log.info(String
	// 			.format("Ids are used out during %d. Waiting till next second/milisencond.",
	// 					lastTimestamp))

	timestamp := t.genTimeFunc()
	for timestamp <= lastTimestamp {
		timestamp = t.genTimeFunc()
	}

	// if (log.isInfoEnabled())
	// 	log.info(String.format("Next second/milisencond %d is up.",
	// 			timestamp))

	return timestamp
}
func (t *TimeOp) ParseTimestampToDuration(timestamp int64) (uint64, error) {
	timeStr := strconv.FormatInt(timestamp, 10)
	if len(timeStr) == 10 {
		return uint64((timestamp*1000 - t.epoch) / 1000), nil
	} else if len(timeStr) == 13 {
		return uint64((timestamp - t.epoch) / 1000), nil
	} else {
		return 0, errors.New("the timestamp is invalided")
	}
}

func NewTimeOpWithMaxPeak(epoch int64) *TimeOp {
	t := &TimeOp{epoch: epoch}
	t.genTimeFunc = t.getTimeNow4Second
	t.transTimeFunc = t.transTimeWithMaxPeak
	return t
}
func NewTimeOpWithMinGranularity(epoch int64) *TimeOp {
	t := &TimeOp{epoch: epoch}
	t.genTimeFunc = t.getTimeNow4Milli
	t.transTimeFunc = t.transTimeWithMinGranularity
	return t
}
