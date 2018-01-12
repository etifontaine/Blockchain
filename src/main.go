package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Add Inception to list")
	bc.AddBlock("Add Star Wars 8 to list")

	for _, block := range bc.blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
