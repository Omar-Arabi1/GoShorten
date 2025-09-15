package utils

import (
	"encoding/base64"
	"math/rand"
)

func ShortenUrl(longUrl string) string {
	longShortUrlKey := base64.StdEncoding.EncodeToString([]byte(longUrl))
	shortUrlKey := ""
	for range 5 {
		randIndex := rand.Intn(5)
		shortUrlKey += string(longShortUrlKey[randIndex])
	}
	return shortUrlKey
}
