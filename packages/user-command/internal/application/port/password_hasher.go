package port

type PasswordHasher interface {
	Hash(plainText string) (string, error)
	Verify(plainText string, hashedText string) (bool, error)
}
