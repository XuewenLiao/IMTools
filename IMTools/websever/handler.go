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

//在群组发送系统通知
func SendGroupSysMsg(c *gin.Context) {
	groupName := c.Request.FormValue("groupname")
	content := c.Request.FormValue("content")
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	errorCode := apis.SendSystemMsg(userSig, groupName, content)
	if errorCode == 0 {
		c.String(http.StatusOK, "发送系统通知成功！")
	} else {
		c.String(http.StatusOK, "发送系统通知失败，errorCode："+strconv.FormatInt(errorCode, 10))
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"groupnum": "已成功向" + strconv.Itoa(groupNum) + "个群发消息",
	//})
}

//获取指定用户的好友列表
func GetFriendList(c *gin.Context) {
	userfrdid := c.Request.FormValue("userfrdid")
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	friendNum, errorCode := apis.GetFriendList(userSig, userfrdid)
	if errorCode == 0 {
		c.String(http.StatusOK, userfrdid+"的好友数："+strconv.FormatInt(friendNum, 10))
	} else {
		c.String(http.StatusOK, "拉取好友失败，errorCode："+strconv.FormatInt(errorCode, 10))
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"groupnum": "已成功向" + strconv.Itoa(groupNum) + "个群发消息",
	//})
}

//批量发群消息
func BatchSendGroupMsg(c *gin.Context) {
	groupNum, _ := strconv.Atoi(c.Request.FormValue("groupnum"))
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	errorCode := apis.SendGroupMsg(userSig, groupNum)
	if errorCode == 0 {
		c.String(http.StatusOK, "成功向"+strconv.Itoa(groupNum)+"个群发消息")
	} else {
		c.String(http.StatusOK, "发送失败，errorCode："+strconv.FormatInt(errorCode, 10))
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"groupnum": "已成功向" + strconv.Itoa(groupNum) + "个群发消息",
	//})
}

//批量发单聊消息
func BatchSendC2CMsg(c *gin.Context) {
	userNum, _ := strconv.Atoi(c.Request.FormValue("usernum"))
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	errorCode := apis.SendC2CMsg(userSig, userNum)
	if errorCode == 0 {
		c.String(http.StatusOK, "成功向"+strconv.Itoa(userNum)+"个用户发送消息")
	} else {
		c.String(http.StatusOK, "发送失败，errorCode："+strconv.FormatInt(errorCode, 10))
	}

	//errorList := apis.SendC2CMsg(userSig, userNum)
	//if len(errorList) == 0 {
	//	c.String(http.StatusOK,"成功向" + strconv.Itoa(userNum) + "个用户发送消息")
	//
	//}else {
	//	var errorString string
	//	for i := 0; i < len(errorList); i++ {
	//		errorString = errorString  + errorList[i].To_Account + "发送失败-错误码：" + strconv.FormatInt(errorList[i].ErrorCode,10) + "；"
	//	}
	//	c.String(http.StatusOK,errorString)
	//
	//}

	//c.JSON(http.StatusOK, gin.H{
	//	"usernum": "已成功向" + strconv.Itoa(userNum) + "个用户发消息",
	//})
}

//批量加群，是指提供一个群id，给这个群加很多的用户，实现群成员人数达到指定数量的目标
func BatchAddGroup(c *gin.Context) {
	groupId := c.Request.FormValue("groupid")
	accoutNumOfgroup, _ := strconv.Atoi(c.Request.FormValue("accoutnumofgroup"))

	fmt.Printf("groupId--%v\naccoutNumOfgroup--%v\n", groupId, accoutNumOfgroup)

	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)

	//获取所有群组的Id
	//allAccountsName := apis.AllAccountsId
	allAccountsName, errorCode1 := apis.Multiaccount_PostData(userSig, 2) //假设每次批量加群 注册的用户数为能容纳的用户上限

	if allAccountsName == nil {

		c.String(http.StatusOK, "还没创建账户，请先进行建群操作！")
		//c.JSON(http.StatusOK, gin.H{
		//	"error": "还没创建账户，请先进行建群操作！",
		//})
	} else {

		if errorCode1 == 0 {
			errorCode2 := apis.AddGroupAccount(userSig, groupId, accoutNumOfgroup)
			if errorCode2 == 0 {
				c.JSON(http.StatusOK, "操作成功-群ID："+groupId+"；人数："+strconv.Itoa(accoutNumOfgroup))

			} else {
				c.String(http.StatusOK, "加群失败，errorCode："+strconv.FormatInt(errorCode2, 100))

			}

		}
		//apis.AddGroupAccount(userSig, groupId, accoutNumOfgroup, allAccountsName)

	}
}

//批量建群
func BatchCreatGroup(c *gin.Context) {

	accountsNum, _ := strconv.Atoi(c.Request.FormValue("accountsnum"))
	groupNum, _ := strconv.Atoi(c.Request.FormValue("groupnum"))
	accountName := c.Request.FormValue("accountname")

	fmt.Printf("accountsNum--%v\ngroupNum--%v\n", accountsNum, groupNum)

	//生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)
	fmt.Printf("userSig: %v", userSig)

	//批量添加账户
	_, errorCode := apis.Multiaccount_PostData(userSig, accountsNum)

	if errorCode == 0 {
		//生成群组
		errorCode2 := apis.BatchCreatgroup(userSig, groupNum, accountName)
		if errorCode2 == 0 {
			c.String(http.StatusOK, "成功添加"+strconv.Itoa(groupNum)+"个群组")

		} else {
			c.String(http.StatusOK, "添加群组失败，errorCode："+strconv.FormatInt(errorCode2, 10))

		}

	} else {
		c.String(http.StatusOK, "创建账户失败，errorCode："+strconv.FormatInt(errorCode, 10))

	}

}

////批量建群
//func BatchCreatGroup(c *gin.Context) {
//
//	accountsNum, _ := strconv.Atoi(c.Request.FormValue("accountsnum"))
//	groupNum, _ := strconv.Atoi(c.Request.FormValue("groupnum"))
//
//	fmt.Printf("accountsNum--%v\ngroupNum--%v\n", accountsNum, groupNum)
//
//	//生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
//	userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)
//	//fmt.Printf("userSig: %v",userSig)
//
//	//批量添加账户
//	allAccountsName, errorCode := apis.Multiaccount_PostData(userSig, accountsNum)
//
//	if errorCode == 0 {
//		//生成群组
//		errorCode2 := apis.BatchCreatgroup(userSig, groupNum, allAccountsName)
//		if errorCode2 == 0 {
//			c.String(http.StatusOK, "成功添加"+strconv.Itoa(groupNum)+"个群组")
//
//		} else {
//			c.String(http.StatusOK, "添加群组失败，errorCode："+strconv.FormatInt(errorCode2, 10))
//
//		}
//
//	} else {
//		c.String(http.StatusOK, "创建账户失败，errorCode："+strconv.FormatInt(errorCode, 10))
//
//	}
//
//}

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
	errorCode := apis.AddFriend(userSig, userId, friendNum)

	if errorCode == 0 {
		c.String(http.StatusOK, "操作成功-目标用户："+userId+"；添加好友数："+strconv.Itoa(friendNum))

	} else {
		c.String(http.StatusOK, "添加好友失败，errorCode："+strconv.FormatInt(errorCode, 10))
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"userid":    userId,
	//	"friendnum": friendNum,
	//})

}
