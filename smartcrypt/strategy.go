package smartcrypt

func SelectStrategy(ctx EncryptionContext) Encryptor {
	// Very simple strategy logic for now
	if ctx.Sensitivity == "high" {
		if aes := LookupEncryptor("AES-GCM"); aes != nil {
			return aes
		}
	}

	// Default fallback
	for _, name := range ListAvailable() {
		return LookupEncryptor(name)
	}

	return nil
}
