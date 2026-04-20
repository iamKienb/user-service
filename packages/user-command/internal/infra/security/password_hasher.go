package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"
	"strings"

	configx "github.com/iamKienb/shopify-go-platform/config"
	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type argon2Hasher struct {
	cfg configx.Argon2Config
}

func NewArgon2Hasher(cfg configx.Argon2Config) port.PasswordHasher {
	return &argon2Hasher{cfg: cfg}
}

func (h *argon2Hasher) Hash(plainText string) (string, error) {
	p := h.cfg

	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(plainText), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil

}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (h *argon2Hasher) Verify(plainText, hashedText string) (bool, error) {
	p, salt, hash, err := decodeHash(hashedText)
	if err != nil {
		return false, fmt.Errorf("decode password: %w", err)
	}

	otherHash := argon2.IDKey([]byte(plainText), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, identity.ErrInvalidEmailPassword
}

func decodeHash(encodedHash string) (p *configx.Argon2Config, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &configx.Argon2Config{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = len(salt)

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
