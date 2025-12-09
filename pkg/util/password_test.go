package util

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
    pwd := "Secret123!"
    hash, err := HashPassword(pwd)
    if err != nil {
        t.Fatalf("HashPassword error: %v", err)
    }
    if hash == "" || hash == pwd {
        t.Fatalf("unexpected hash: %q", hash)
    }
    if !CheckPasswordHash(pwd, hash) {
        t.Fatalf("expected password to match hash")
    }
    if CheckPasswordHash("wrong", hash) {
        t.Fatalf("expected wrong password to fail")
    }
}
