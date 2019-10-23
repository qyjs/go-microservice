package auth

import (
	"github.com/gin-gonic/gin"
)

// Username 用户名
// Password 密码
const (
	Username = "Naturami.admin"
	Password = "450HiF3snj"
)

// Token ...
type Token struct {
	AccessToken string `json:"AccessToken"`
}

// Profile ...
type Profile struct {
	ID         string `json:"id"`
	Attributes []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"attributes"`
}

type authResp struct {
	Token  string
	Tenant string
}

// GetTokenAndTenant ...
func GetTokenAndTenant(c *gin.Context) {
	token, err := GetToken()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	tenant, err := GetTenant(token)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	result := &authResp{token, tenant}
	c.JSON(200, gin.H{
		"Content": result,
	})
}
