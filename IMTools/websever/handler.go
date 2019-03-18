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

//批量加群，是只提供一个群id，给这个群加很多的用户，实现群成员人数达到指定数量的目标
func BatchAddGroup(c *gin.Context) {
	groupId := c.Request.FormValue("groupid")
	accoutNumOfgroup, _ := strconv.Atoi(c.Request.FormValue("accoutnumofgroup"))

	fmt.Printf("groupId--%v\naccoutNumOfgroup--%v\n", groupId, accoutNumOfgroup)

	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	//获取所有群组的Id
	allAccountsName := apis.AllAccountsId
	if allAccountsName == nil {
		c.JSON(http.StatusOK, gin.H{
			"error": "还没创建账户，请先进行建群操作！",
		})
	} else {

		apis.AddGroupAccount(userSig, groupId, accoutNumOfgroup, allAccountsName)

		c.JSON(http.StatusOK, gin.H{
			"groupid":          groupId,
			"accoutnumofgroup": accoutNumOfgroup,
		})
	}
}

//批量建群
func BatchCreatGroup(c *gin.Context) {

	accountsNum, _ := strconv.Atoi(c.Request.FormValue("accountsnum"))
	groupNum, _ := strconv.Atoi(c.Request.FormValue("groupnum"))

	fmt.Printf("accountsNum--%v\ngroupNum--%v\n", accountsNum, groupNum)

	//生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)
	//fmt.Printf("userSig: %v",userSig)

	//批量添加账户
	allAccountsName := apis.Multiaccount_PostData(userSig, accountsNum)

	//生成群组
	apis.BatchCreatgroup(userSig, groupNum, allAccountsName)

	c.JSON(http.StatusOK, gin.H{
		"accountsnum": accountsNum,
		"groupnum":    groupNum,
	})

}

//批量加好友，是只给一个用户账号，加很多的好友，实现好友数量达到指定数量的目标
func BatchAddFriend(c *gin.Context) {
	userId := c.Request.FormValue("userid")
	friendNum, _ := strconv.Atoi(c.Request.FormValue("friendnum"))

	fmt.Printf("userId--%v\nfriendNum--%v\n", userId, friendNum)

	//生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	//批量删除指定用户的好友，保证每次操作都能使该用户有新的一批好友
	apis.DeleteFriend(userSig, userId)

	//批量加好友
	apis.AddFriend(userSig, userId, friendNum)

	c.JSON(http.StatusOK, gin.H{
		"userid":    userId,
		"friendnum": friendNum,
	})

}
