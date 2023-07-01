package service

import (
	"math/rand"
	"time"
)

type JitterService struct {
}

type IJitterService interface {
	GetRandomNumber(low, high uint64) uint64
}

func NewJitterService() IJitterService {
	return &JitterService{}
}

func (jitter *JitterService) GetRandomNumber(low, high uint64) uint64 {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	res := random.Intn(int(high)) + int(low)
	return uint64(res)
}
