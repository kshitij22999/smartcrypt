package smartcrypt

import (
	"errors"
	"fmt"
	"smartcrypt/compress"
	//_ "github.com/kshitij22999/smartcrypt/algo"
	//"github.com/kshitij22999/smartcrypt/compress"
	// other imports
)

func EncryptFile(ctx EncryptionContext, plain []byte) ([]byte, error) {
	strategy := SelectStrategy(ctx)
	fmt.Println("strat", strategy)
	if strategy == nil {
		return nil, errors.New("no valid encryptor found")
	}

	// Optional compression step
	if ctx.Compress {
		var err error
		plain, err = compress.Compress(plain)
		if err != nil {
			return nil, err
		}
	}

	encrypted, err := strategy.Encrypt(plain)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

func DecryptFile(data []byte, decompress bool) ([]byte, error) {
	strategy := LookupEncryptor("AES-GCM")
	if strategy == nil {
		return nil, errors.New("decrypt: no strategy found")
	}

	decrypted, err := strategy.Decrypt(data)
	if err != nil {
		return nil, err
	}

	if decompress {
		decrypted, err = compress.Decompress(decrypted)
		if err != nil {
			return nil, err
		}
	}

	return decrypted, nil
}
