/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"framagit.org/InfoLibre/rapido/models/settings"

	//	"github.com/golang-jwt/jwt"
	//	"github.com/golang-jwt/jwt/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// Encodes/decodes token
type Payload struct {
	Iat          int    `json:"iat"`
	Exp          int    `json:"exp"`
	Iss          string `json:"iss"`
	Sub          string `json:"sub"`
	ID           int    `json:"id"`
	Pseudonym    string `json:"pseudonym"`
	EmailAddress string `json:"email_address"`
	LastLoggedIn int    `json:"last_logged_in"`
	AccessLevel  int    `json:"access_level"`
}

const (
	// Administrator access level
	AdminAL = 10
	// User access level
	UserAL = 1
)

// Generates JWT token
func GenerateToken(payload map[string]interface{}, expireHour time.Duration, issuer string) (string, error) {
	subject := fmt.Sprintf("access granted for %d hour(s)", expireHour)
	claims := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * expireHour).Unix(),
		"iss": issuer,
		"sub": subject,
	}

	for key, value := range payload {
		claims[key] = value
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(settings.SecretKey()))
}

// Extracts JWT token payload
func GetClaims(tokenString string) (jwt.MapClaims, error) {
	str := strings.Split(tokenString, " ")
	hmacSecret := []byte(settings.SecretKey())
	token, err := jwt.Parse(str[1], func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}
	return nil, err
}

// Extracts token to payload
func Extract(token string) (result Payload, err error) {
	claims, err1 := GetClaims(token)
	if err1 != nil {
		return
	}

	byteJSON, err2 := json.Marshal(claims)
	if err2 != nil {
		return
	}

	err2 = json.Unmarshal(byteJSON, &result)
	return
}

// Checks JWT token
func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{"message": "you_must_connect", "error": err.Error()})
		}
	}
}

// Checks access level
func RBAC(accessLevel int) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, _ := Extract(c.GetHeader("Authorization"))
		if payload.AccessLevel != accessLevel {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		}
	}
}
