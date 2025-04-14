package utils

import (
	"errors"
	"sync"
	"time"
)

type Snowflake struct {
	mu            sync.Mutex // 互斥锁，保证并发安全
	startTime     int64      // 起始时间戳（毫秒）
	machineID     int64      // 机器ID
	sequence      int64      // 序列号
	lastTimestamp int64      // 上次生成ID的时间戳
}

const (
	sequenceBits  = 12                           // 序列号位数
	machineIDBits = 10                           // 机器ID位数
	maxMachineID  = -1 ^ (-1 << machineIDBits)   // 最大机器ID（1023）
	maxSequence   = -1 ^ (-1 << sequenceBits)    // 最大序列号（4095）
	timeShift     = machineIDBits + sequenceBits // 时间戳左移位数
	machineShift  = sequenceBits                 // 机器ID左移位数
)

func NewSnowflake(machineID int64, startTime int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("machineID超出范围")
	}

	return &Snowflake{
		startTime:     startTime, // 转为毫秒
		machineID:     machineID,
		lastTimestamp: -1,
	}, nil
}

func (s *Snowflake) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := time.Now().UnixMilli()
	timestamp := currentTime - s.startTime

	if timestamp < s.lastTimestamp {
		return 0, errors.New("时钟回拨，拒绝生成ID")
	}

	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 { // 当前毫秒序列号用完，等待下一毫秒
			for currentTime <= s.lastTimestamp {
				currentTime = time.Now().UnixMilli()
			}
			timestamp = currentTime - s.startTime
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = timestamp

	id := (timestamp << timeShift) | (s.machineID << machineShift) | s.sequence
	return id, nil
}

func ParseID(id int64) (timestamp int64, machineID int64, sequence int64) {
	timestamp = id >> timeShift
	machineID = (id >> machineShift) & maxMachineID
	sequence = id & maxSequence
	return
}
