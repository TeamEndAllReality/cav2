package cav2

import (
	"io/ioutil"

	"github.com/aviddiviner/go-murmur"
)

//GetByteArrayHash Computes a murmur2 byte array hash
func GetByteArrayHash(bytes []byte) int {
	return int(murmur.MurmurHash2(computeNormalizedArray(bytes), 1))
}

//GetFileHash Computes a murmur2 hash of a file
func GetFileHash(file string) (int, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	result := GetByteArrayHash(bytes)
	return result, nil
}

func computeNormalizedArray(bytes []byte) []byte {
	var newArray []byte
	for _, byte := range bytes {
		if !isWhitespaceCharacter(byte) {
			newArray = append(newArray, byte)
		}
	}
	return newArray
}

func isWhitespaceCharacter(b byte) bool {
	return b == 9 || b == 10 || b == 13 || b == 32
}
