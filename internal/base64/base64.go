package base64

import (
	"math/rand"
	"time"
)

const symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

func Generate(length int) string {
	res := make([]byte, length)

	r := rand.New(rand.NewSource(time.Now().UnixMilli()))

	for i := range length {
		res[i] = symbols[r.Intn(len(symbols))]
	}

	return string(res)

}
