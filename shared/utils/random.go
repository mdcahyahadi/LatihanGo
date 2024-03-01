package utils

import (
	"crypto/rand"
	"fmt"
)

func RandomExternalID() string {
	randomCrypto, _ := rand.Prime(rand.Reader, 32)

	return fmt.Sprintf("%v", randomCrypto)
}
