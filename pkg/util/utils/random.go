package utils

import "math/rand"

// GetRandom 获取随机数,范围是[min,max)
func GetRandom(min, max int) int {
	return min + rand.Intn(max-min)
}
