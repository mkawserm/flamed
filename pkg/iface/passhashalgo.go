package iface

type IPasswordHashAlgorithm interface {
	Algorithm() string

	Encode(password string, salt string) (string, error)
	Verify(password string, encoded string) (bool, error)
}

type IPasswordHashAlgorithmFactory interface {
	AppendPasswordHashAlgorithm(pha IPasswordHashAlgorithm)

	CheckPassword(password, encoded string) (bool, error)
	MakePassword(password, salt, algorithm string) (string, error)
}
