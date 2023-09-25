package api

import (
	"testing"

	"github.com/stretchr/testify/require"
	db "github.com/viralparmarme/simple-bank/db/sqlc"
	"github.com/viralparmarme/simple-bank/util"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}
