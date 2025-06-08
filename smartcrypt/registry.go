package smartcrypt

// Holds registered algorithms
var encryptorRegistry = make(map[string]Encryptor)

// RegisterEncryptor registers an encryptor for strategy selection and lookup.
func RegisterEncryptor(enc Encryptor) {
	encryptorRegistry[enc.Name()] = enc
}

// LookupEncryptor returns the registered encryptor by name.
func LookupEncryptor(name string) Encryptor {
	return encryptorRegistry[name]
}

// ListAvailable returns a list of available algorithms.
func ListAvailable() []string {
	keys := make([]string, 0, len(encryptorRegistry))
	for k := range encryptorRegistry {
		keys = append(keys, k)
	}
	return keys
}
