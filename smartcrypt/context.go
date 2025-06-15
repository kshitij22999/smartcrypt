package smartcrypt

type EncryptionContext struct {
	FileType    string
	Sensitivity string
	Compress    bool
	Key         []byte
}
