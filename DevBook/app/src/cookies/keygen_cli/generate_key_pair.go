package main

import (
	"devbook_app/src/cookies"
	"encoding/hex"
	"fmt"
)

func main() {
	hashKey, blockKey := cookies.GenerateKeyPair()
	hashKeyString := hex.EncodeToString(hashKey)
	blockKeyString := hex.EncodeToString(blockKey)
	fmt.Printf("HashKey: %v\nBlockKey: %v\n", hashKeyString, blockKeyString)
}
