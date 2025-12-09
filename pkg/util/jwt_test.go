package util

import (
    "os"
    "testing"

    "github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
    os.Setenv("JWT_SECRET", "testsecret")
    tokenStr, err := GenerateToken(42, "admin")
    if err != nil {
        t.Fatalf("GenerateToken error: %v", err)
    }
    if tokenStr == "" {
        t.Fatalf("expected non-empty token")
    }

    // Parse and validate claims
    parsed, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            t.Fatalf("unexpected signing method")
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil || !parsed.Valid {
        t.Fatalf("token not valid: %v", err)
    }
}
