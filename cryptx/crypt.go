/*
@File    :   crypt.go
@Time    :   2023/07/31 18:18:07
@Author  :   Lison LIN
@Version :   1.0
@Contact :   lisonlin22@gmail.com
@Desc    :
@WiKi    :
*/

package cryptx

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
