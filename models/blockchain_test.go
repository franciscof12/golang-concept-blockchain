package models

import (
	"testing"
)

// TestNewBlockchain verifica la creación de una nueva blockchain con un bloque génesis.
func TestNewBlockchain(t *testing.T) {
	blockchain := NewBlockchain(1) // Asumiendo que la dificultad 1 es suficiente para pruebas rápidas.

	if len(blockchain.Chain) != 1 {
		t.Errorf("Expected blockchain length of 1, got %v", len(blockchain.Chain))
	}

	if blockchain.Chain[0].PreviousHash != "0" {
		t.Errorf("Expected genesis block previous hash to be '0', got '%v'", blockchain.Chain[0].PreviousHash)
	}
}

// TestAddBlock verifica la funcionalidad de añadir un nuevo bloque a la cadena.
func TestAddBlock(t *testing.T) {
	blockchain := NewBlockchain(1)
	blockchain.AddBlock("Alice", "Bob", 50)

	if len(blockchain.Chain) != 2 {
		t.Errorf("Expected blockchain length of 2, got %v", len(blockchain.Chain))
	}
}

// TestBlockchainValidity verifica que la blockchain sea válida después de añadir bloques.
func TestBlockchainValidity(t *testing.T) {
	blockchain := NewBlockchain(1)
	blockchain.AddBlock("Alice", "Bob", 50)
	blockchain.AddBlock("Bob", "Charlie", 25)

	if !blockchain.IsValid() {
		t.Error("Expected blockchain to be valid")
	}
}
