package main

import (
	"fmt"
	"golang-blockchain/models"
)

func main() {
	blockchain := models.NewBlockchain(2)
	blockchain.AddBlock("Alice", "Bob", 5)
	blockchain.AddBlock("John", "Bob", 2)
	fmt.Println(blockchain.IsValid())
}
