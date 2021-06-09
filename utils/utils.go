package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
	"milkteapool.com/pool-server/crypto/bech32"
)

func DecodePuzzleHash(address string) []byte {
	_, data, err := bech32.Decode(address)
	if err != nil {
		log.Error("Error:", err)
	}
	return data
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
