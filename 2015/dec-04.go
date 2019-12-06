package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var key []byte
	_, err := fmt.Scanf("%s\n", &key)
	if err != nil {
		panic(err)
	}

	tA := time.Now()
	fmt.Printf("A: %d (in %v)\n", solve(key, 5), time.Since(tA))
	tB := time.Now()
	fmt.Printf("B: %d (in %v)\n", solve(key, 6), time.Since(tB))
}

func solve(key []byte, n int) int64 {
	prefix := strings.Repeat("0", n)
	for n := int64(0); ; n++ {
		data := strconv.AppendInt(key, n, 10)
		hash := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(hash, prefix) {
			return n
		}
	}
}
