package encription

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {

	testCase := []struct {
		name     string
		password string
		checkout func(t *testing.T, password string, hasedpassword string, err error)
	}{
		{
			name:     "ok",
			password: "thisismysecretepassword",
			checkout: func(t *testing.T, password string, hasedpassword string, err error) {
				require.NoError(t, err)
				err = CheckPassword(password, hasedpassword)
				require.NoError(t, err)
			},
		},
		{
			name:     "empty password",
			password: "",
			checkout: func(t *testing.T, password string, hasedpassword string, err error) {
				require.Error(t, err)
				err = CheckPassword(password, hasedpassword)
				require.Error(t, err)

			},
		},
	}
	for _, tc := range testCase {
		hashedpassword, err := GenerateHashedPassword(tc.password)
		t.Run(tc.name, func(t *testing.T) {
			tc.checkout(t, tc.password, hashedpassword, err)
		})
	}

}
