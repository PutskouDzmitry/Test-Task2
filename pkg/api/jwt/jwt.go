package jwt

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
	"os"
	"time"
)

type JWT struct {
	AccessExpiration int
	SecretKey        string
}

func NewJWT(
	accessExpiration int,
	secretKey string,
) *JWT {
	return &JWT{
		AccessExpiration: accessExpiration,
		SecretKey:        secretKey,
	}
}

// Claims represents the JWT claims for the access and refresh tokens.
type claims struct {
	// UserID is the user ID associated with the token.
	UserID string `json:"user_id"`
	Role   string `json:"role"`

	// StandardClaims contains the standard JWT claims.
	jwt.StandardClaims
}

// CreateAccessToken creates a JWT access token for the given user ID.
// The token expires in 15 minutes.
// Returns the signed token string and an error if there's any.
func (j *JWT) CreateAccessToken(userID string, role string) (string, error) {

	t := time.Duration(j.AccessExpiration) * time.Minute
	if os.Getenv("GIN_MODE") == "debug" {
		t = time.Duration(j.AccessExpiration) * time.Hour
	}
	// Create the access token claims
	accessClaims := &claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Add(t).Unix())),
			IssuedAt:  jwt.NewTime(float64(time.Now().Unix())),
			Issuer:    "Test-App2",
			Subject:   "access",
		},
	}

	// Create the access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Sign the token with a secret key
	accessString, err := accessToken.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return accessString, nil
}

// GetToken GetRefreshToken retrieves a JWT token from the Authorization header for token refreshing.
//
// c: A *gin.Context representing the HTTP request/response context.
//
// Returns:
// - A *jwt.Token representing the refreshed token.
// - An error if the Authorization header is not found or if the token is invalid.
func (j *JWT) GetToken(authHeader string) (*jwt.Token, error) {
	// Verify the token
	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error while parsing token")
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JWT) Parse(strToken string) (jwt.MapClaims, error) {
	tokenString := strToken[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error while parsing token")
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("couldn't parse jwt claims")
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
