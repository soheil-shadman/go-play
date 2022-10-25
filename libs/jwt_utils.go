package libs

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func JWTEncode(secret string, payload jwt.MapClaims, expiresAfter time.Duration) (string, error) {
	payload["exp"] = time.Now().UTC().Add(expiresAfter).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}
func JWTDecode(secret string, token string) (jwt.MapClaims, error) {
	tt, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tt.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}
	exp := claims["exp"].(float64)
	expT := time.Unix(int64(exp), 0)
	if time.Now().After(expT) {
		return nil, errors.New("expired token")
	}
	return claims, nil
}
