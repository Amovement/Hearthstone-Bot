package utils

import (
	"math/rand"
	"time"
)

// Rand 生成 x 到 y 范围内的随机整数
func Rand(x, y int) int {
	if x > y {
		x, y = y, x // 确保 x 小于等于 y
	}
	rand.Seed(time.Now().UnixNano())
	return x + rand.Intn(y-x+1)
}

// RandNormal 生成符合正态分布的随机数
func RandNormal(mean, stdDev float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.NormFloat64()*stdDev + mean
}
