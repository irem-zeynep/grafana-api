package password

import (
	"context"
	"github.com/sethvargo/go-password/password"
)

type IPasswordGenerator interface {
	NewPassword(ctx context.Context) string
}

type passwordGenerator struct {
}

func NewPasswordGenerator() IPasswordGenerator {
	return &passwordGenerator{}
}

func (p passwordGenerator) NewPassword(ctx context.Context) string {
	pwd, err := password.Generate(8, 1, 1, false, false)
	if err != nil {
		return "pwd123"
	}

	return pwd
}
