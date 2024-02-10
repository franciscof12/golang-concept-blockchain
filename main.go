package main

import (
	"fmt"
	"golang-blockchain/models"
)

func main() {
	blockchain := models.NewBlockchain(2)
	blockchain.AddBlock("Edward", "Francisco", 5)
	blockchain.AddBlock("Francisco", "Alicia", 2)
	fmt.Println(blockchain.IsValid())
}
