package types

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const UserSignKey = ("GbpKLAzbKBfxJsbauHMdn7GNwR6XGfIL")
const TokenTTL = 1 * time.Hour

type UserDomain struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type UserClaim struct {
	Id    int64
	Name  string
	Email string
	jwt.RegisteredClaims
}
