package iface

type PasswordHashAlgorithm interface {
	Encode(password string, salt string) (string, error)
	Verify(password string, encoded string) (bool, error)
}

type PasswordHashAlgorithmFactory interface {
	AppendPasswordHashAlgorithm(ph PasswordHashAlgorithm)

	CheckPassword(password, encoded string) (bool, error)
	MakePassword(password, salt, algorithm string) (string, error)
}
