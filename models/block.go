package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data         map[string]interface{}
	Hash         string
	PreviousHash string
	TimeStamp    time.Time
	Pow          int
}

// Creamos una funcion para hashear el identificador de un bloque con el algoritmo SHA256
func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.TimeStamp.String() + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.CalculateHash()
	}
}
