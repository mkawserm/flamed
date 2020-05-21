package crypto

import (
	"errors"
	"strings"

	"github.com/mkawserm/flamed/pkg/iface"
)

var ErrUnimplementedPasswordHashAlgorithm = errors.New("password hash algorithm is not" +
	" implemented yet")

type PasswordHashAlgorithmFactory struct {
	mAlgorithmMap map[string]iface.IPasswordHashAlgorithm
}

func (p *PasswordHashAlgorithmFactory) AppendPasswordHashAlgorithm(pha iface.IPasswordHashAlgorithm) {
	p.mAlgorithmMap[pha.Algorithm()] = pha
}

func (p *PasswordHashAlgorithmFactory) CheckPassword(password, encoded string) (bool, error) {
	algorithm := identifyPasswordHashAlgorithm(encoded)
	a, found := p.mAlgorithmMap[algorithm]
	if found {
		return a.Verify(password, encoded)
	} else {
		return false, ErrUnimplementedPasswordHashAlgorithm
	}
}

func (p *PasswordHashAlgorithmFactory) MakePassword(password, salt, algorithm string) (string, error) {
	a, found := p.mAlgorithmMap[algorithm]
	if found {
		return a.Encode(password, salt)
	} else {
		return "", ErrUnimplementedPasswordHashAlgorithm
	}
}

func NewPasswordHashAlgorithmFactory() *PasswordHashAlgorithmFactory {
	return &PasswordHashAlgorithmFactory{mAlgorithmMap: make(map[string]iface.IPasswordHashAlgorithm)}
}

func DefaultPasswordHashAlgorithmFactory() *PasswordHashAlgorithmFactory {
	factory := NewPasswordHashAlgorithmFactory()
	factory.AppendPasswordHashAlgorithm(NewArgon2PasswordHashAlgorithm())
	return factory
}

func identifyPasswordHashAlgorithm(encoded string) string {
	size := len(encoded)

	if size == 32 && !strings.Contains(encoded, "$") {
		return "unsalted_md5"
	}

	if size == 37 && strings.HasPrefix(encoded, "md5$$") {
		return "unsalted_md5"
	}

	if size == 46 && strings.HasPrefix(encoded, "sha1$$") {
		return "unsalted_sha1"
	}

	return strings.SplitN(encoded, "$", 2)[0]
}
