package main

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateFloat() string {
	rand.Seed(time.Now().UnixNano())
	// 生成一个介于 0.0 到 1.0 之间的随机浮点数
	randomFloat := rand.Float64()

	return strconv.FormatFloat(randomFloat*10, 'f', 2, 64)
}

func GenerateInt() int {
	rand.Seed(time.Now().UnixNano())
	// 生成一个介于 0.0 到 1.0 之间的随机浮点数
	randomInt := rand.Intn(100)

	return randomInt
}

func GenerateUUID(maxLen int) string {
	raw := strings.ReplaceAll(uuid.NewString(), "-", "")
	if len(raw) <= maxLen {
		return raw
	} else {
		return raw[:maxLen]
	}
}
