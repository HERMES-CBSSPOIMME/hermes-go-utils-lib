package identifiergenerator

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateLoggingIdentifier : Generate a logging identifier to store in logging system
func GenerateLoggingIdentifier(prefix string) (string, error) {

	var err error

	randBytes, err := GenerateCryptoSecureRandomBytes(8)

	if err != nil {
		return "", err
	}

	timestampAsBytesArray, err := time.Now().MarshalBinary()

	if err != nil {
		return "", err
	}

	seed := append(randBytes[:], timestampAsBytesArray[:]...)

	randHash, err := GetMD5HashString(seed)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", prefix, randHash), nil
}

// GenerateCryptoSecureRandomBytes : Generate n-sized random bytes array
func GenerateCryptoSecureRandomBytes(n int) ([]byte, error) {
	r := make([]byte, n)
	_, err := rand.Read(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetMD5HashString : Get MD5 hash string of bytes array
func GetMD5HashString(s []byte) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write(s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
