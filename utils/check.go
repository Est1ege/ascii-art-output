package utils

import (
	"fmt"
	"hash/fnv"
	"os"
)

func Checkhash(style string) {
	// 1766917683
	// мы получаем хэш файла
	hasher := fnv.New32a()
	hasher.Write([]byte(style))
	hashValue := hasher.Sum32()
	if hashValue != 1766917683 && hashValue != 4281396044 && hashValue != 3075161722 {
		fmt.Println("ERROR : wrong style file")
		os.Exit(0)
	}
}

func CheckInput(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] > 127 {
			if s[i] == 10 {
				continue
			}
			fmt.Println("ERROR: wrong format")
			return false
		}
	}
	return true
}
