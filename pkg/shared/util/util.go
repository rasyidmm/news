package util

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
)

func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func HashSha512(Texts string) string {
	sha_512 := sha512.New()
	sha_512.Write([]byte(Texts))
	return fmt.Sprintf("%x", sha_512.Sum(nil))
}
