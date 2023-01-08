package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestValidate(t *testing.T) {
	v := yav.NewPlayground()

	test := func(account Account) func(t *testing.T) {
		return func(t *testing.T) {
			assert.Equal(t, v.Validate(account), account.ChainValidate(), "incompatible with playground")
		}
	}

	tests := []struct {
		name    string
		account Account
	}{{
		name:    "valid",
		account: ValidAccount(),
	}, {
		name: "empty id",
		account: func() Account {
			a := ValidAccount()
			a.ID = ""
			return a
		}(),
	}, {
		name: "invalid id",
		account: func() Account {
			a := ValidAccount()
			a.ID += "a"
			return a
		}(),
	}, {
		name: "empty login",
		account: func() Account {
			a := ValidAccount()
			a.Login = ""
			return a
		}(),
	}, {
		name: "invalid login",
		account: func() Account {
			a := ValidAccount()
			a.Login = "#####"
			return a
		}(),
	}, {
		name: "upper login",
		account: func() Account {
			a := ValidAccount()
			a.Login = "LOGIN"
			return a
		}(),
	}, {
		name: "empty password",
		account: func() Account {
			a := ValidAccount()
			a.Password = ""
			return a
		}(),
	}, {
		name: "short password",
		account: func() Account {
			a := ValidAccount()
			a.Password = "pass"
			return a
		}(),
	}, {
		name: "long password",
		account: func() Account {
			a := ValidAccount()
			a.Password += strings.Repeat("0", 100)
			return a
		}(),
	}, {
		name: "password no lower",
		account: func() Account {
			a := ValidAccount()
			a.Password = strings.ToUpper(a.Password)
			return a
		}(),
	}, {
		name: "password no upper",
		account: func() Account {
			a := ValidAccount()
			a.Password = strings.ToLower(a.Password)
			return a
		}(),
	}, {
		name: "password no digit",
		account: func() Account {
			a := ValidAccount()
			a.Password = "DasPassword#"
			return a
		}(),
	}, {
		name: "password no special character",
		account: func() Account {
			a := ValidAccount()
			a.Password = "DasPassword123"
			return a
		}(),
	}, {
		name: "password with space",
		account: func() Account {
			a := ValidAccount()
			a.Password = "DasPassword# #123"
			return a
		}(),
	}, {
		name: "password not text",
		account: func() Account {
			a := ValidAccount()
			a.Password += string([]rune{rune(0)})
			return a
		}(),
	}, {
		name: "invalid email",
		account: func() Account {
			a := ValidAccount()
			a.Email = "yav+123"
			return a
		}(),
	}, {
		name: "invalid phone",
		account: func() Account {
			a := ValidAccount()
			a.Phone = a.Email
			return a
		}(),
	}, {
		name: "invalid secret",
		account: func() Account {
			a := ValidAccount()
			a.Secret = "insecure"
			return a
		}(),
	}, {
		name: "valid promo code",
		account: func() Account {
			a := ValidAccount()
			a.PromoCode = "BlackFriday2022"
			return a
		}(),
	}, {
		name: "invalid promo code",
		account: func() Account {
			a := ValidAccount()
			a.PromoCode = "Administrator"
			return a
		}(),
	}, {
		name: "invalid first name start",
		account: func() Account {
			a := ValidAccount()
			a.FirstName = "yaV"
			return a
		}(),
	}, {
		name: "invalid last name end",
		account: func() Account {
			a := ValidAccount()
			a.LastName = "YaV"
			return a
		}(),
	}, {
		name: "valid names",
		account: func() Account {
			a := ValidAccount()
			a.FirstName = "Yav"
			a.LastName = "Yav"
			return a
		}(),
	}, {
		name: "display name not title 1",
		account: func() Account {
			a := ValidAccount()
			a.DisplayName = "YAV123\n"
			return a
		}(),
	}, {
		name: "display name not title 2",
		account: func() Account {
			a := ValidAccount()
			a.DisplayName = "YAV\t123"
			return a
		}(),
	}, {
		name: "display name not alpha",
		account: func() Account {
			a := ValidAccount()
			a.DisplayName = "YAV123"
			return a
		}(),
	}, {
		name: "lower display name",
		account: func() Account {
			a := ValidAccount()
			a.DisplayName = "yav"
			return a
		}(),
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.account))
	}
}
