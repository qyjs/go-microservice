package tenantInfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"ebaocloud.com/go-js-service/common/conf"
	"github.com/gin-gonic/gin"
)

// AutoGenerated ...
type AutoGenerated struct {
	ParentName string   `json:"parentName"`
	TagName    []string `json:"tagName"`
}

// BodyObj ...
type BodyObj struct {
	ReturnCode string          `json:"returnCode"`
	Body       []AutoGenerated `json:"body"`
}

// ResponseData ...
type ResponseData struct {
	status string
	body   []AutoGenerated
}

// GetTenant ...
func GetTenant() (*BodyObj, error) {
	client := &http.Client{}

	url := "http://" + conf.AdminServiceName + "/rest/v1/vendor/tag/definition/tree"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &BodyObj{}, err
	}
	req.Header.Add("content-type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return &BodyObj{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &BodyObj{}, err
	}

	bodyobj := &BodyObj{}

	json.Unmarshal([]byte(body), bodyobj)

	return bodyobj, nil
}

// GetAllTenant ...
func GetAllTenant(c *gin.Context) {

	token, err := GetTenant()
	if err != nil {
		c.JSON(http.StatusGatewayTimeout, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "body": token.Body})
	}
	// tenant, err := GetTenant(token)
	// if err != nil {
	// 	c.JSON(500, gin.H{"message": err.Error()})
	// 	return
	// }

	// result := &authResp{token, tenant}
	// c.JSON(200, gin.H{
	// 	"Content": result,
	// })
}
