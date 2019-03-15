package websever

import (
	"IMTools/apis"
	"IMTools/apis/TLSSigAPI"
	"IMTools/sdkconst"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreatGroup(c *gin.Context) {

	accountsNum, _ := strconv.Atoi(c.Request.FormValue("accountsnum"))
	accountNumOfGroup, _ := strconv.Atoi(c.Request.FormValue("accoutnumofgroup"))

	fmt.Printf("accountsNum--%v\naccountNumOfGroup--%v\n", accountsNum, accountNumOfGroup)

	//生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)
	//fmt.Printf("userSig: %v",userSig)

	//批量添加账户
	allAccountsName := apis.Multiaccount_PostData(userSig, accountsNum)

	//生成群组
	apis.Creatgroup_PostData(userSig, accountNumOfGroup, allAccountsName)

	c.JSON(http.StatusOK, gin.H{
		"accountsnum":      accountsNum,
		"accoutnumofgroup": accountNumOfGroup,
	})

}
