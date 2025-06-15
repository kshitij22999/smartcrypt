package smartcrypt

import "fmt"

func SelectStrategy(ctx EncryptionContext) Encryptor {
	// Very simple strategy logic for now
	if ctx.Sensitivity == "high" {
		fmt.Println("select", LookupEncryptor("AES-GCM"))
		if aes := LookupEncryptor("AES-GCM"); aes != nil {
			aes.SetKey(ctx.Key)
			return aes
		}
	}

	//Default fallback
	for _, name := range ListAvailable() {
		return LookupEncryptor(name)
	}

	return nil
}
