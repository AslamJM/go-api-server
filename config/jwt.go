package config

import (
	"os"
	"time"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

const AccessTokenExpiry = time.Minute * 10
const RefreshTokenExpiry = time.Hour * 24 * 30
