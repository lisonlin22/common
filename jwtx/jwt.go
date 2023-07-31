/*
@File    :   jwt.go
@Time    :   2023/07/31 18:17:59
@Author  :   Lison LIN
@Version :   1.0
@Contact :   lisonlin22@gmail.com
@Desc    :
@WiKi    :
*/

package jwtx

import "github.com/golang-jwt/jwt"

func GetToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
