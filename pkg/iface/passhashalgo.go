package iface

type PasswordHashAlgorithm interface {
	Encode(password string, salt string) (string, error)
	Verify(password string, encoded string) (bool, error)
}
