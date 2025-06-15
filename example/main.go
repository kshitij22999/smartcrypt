package main

import (
	"fmt"
	_ "smartcrypt/algo"
	"smartcrypt/smartcrypt"
	// _ "github.com/kshitij22999/smartcrypt/algo"     // for init() to register AES-GCM
	// "github.com/kshitij22999/smartcrypt/smartcrypt" // adjust if your module name differs
)

func main() {
	ctx := smartcrypt.EncryptionContext{
		FileType:    "txt",
		Sensitivity: "high",
		Compress:    true,
		Key:         []byte("8/FURn3gTxqL4IalWf+6m2bN85P6G1wFfa5mqPuptXJ9VRoE15lUEXAjGvdVh7Sf"), //to set the key in algos flow
	}

	data := []byte("akjndkasjnfalfnasdnasljkfnas")

	encrypted, _ := smartcrypt.EncryptFile(ctx, data)
	fmt.Println("Encrypted:", encrypted)

	decrypted, _ := smartcrypt.DecryptFile(encrypted, ctx.Compress)
	fmt.Println("Decrypted:", string(decrypted))
}
