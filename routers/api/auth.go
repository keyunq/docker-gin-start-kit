package api

import (
	"net/http"
	"userInfoService/models"
	"userInfoService/pkg/e"
	"userInfoService/pkg/util"

	"userInfoService/pkg/logging"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)" form:"username"`
	Password string `valid:"Required; MaxSize(50)" form:"password"`
}

// swagger:operation GET /auth auth get
// ---
// summary: 获取jwt-token
// description: 前端系统通过分发的账号密码，从本接口获取有效期为7天的token
// get:
//
//	responses:
//	  200:
//	    description: 获取token成功
//	    schema:
//	    type : AuthResponse
//
// parameters:
//   - name: username
//     in: query
//     description: 账号
//     type: string
//     required: true
//   - name: password
//     in: query
//     description: 密码
//     type: string
//     required: true
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	logging.SugarLogger.Debugf("username: %s", username)
	logging.SugarLogger.Debugf("password: %s", password)
	if username == "" || password == "" {
		logging.SugarLogger.Errorf("username or password 为空")
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		return
	}

	a := auth{Username: username, Password: password}

	err := c.ShouldBind(&a)
	if err != nil {
		logging.SugarLogger.Errorf("ShouldBind error : %s", err)
		return
	}

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	type resp struct {
	}

	isExist := models.CheckAuth(username, password)
	if isExist {
		logging.SugarLogger.Debugf("find auth : %s", username)

		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
			logging.SugarLogger.Errorf("GenerateToken error : %s", err)
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		logging.SugarLogger.Debugf("cannot find auth : %s", username)
		code = e.ERROR_USER
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
