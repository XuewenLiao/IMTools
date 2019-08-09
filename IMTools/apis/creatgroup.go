package apis

import (
	"IMTools/sdkconst"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

//var AllAccountsId []string

type Multiaccount struct {
	Accounts []string
}

type ReMultiaccount struct {
	ActionStatus string
	ErrorCode    int64
	ErrorInfo    string
	FailAccounts []string
}

type AllAccountName struct {
	ActionStatus string
	ErrorCode    int
	ErrorInfo    string
	FailAccounts []string
}

type Creatgroup struct {
	Owner_Account string
	Type          string
	Name          string
}

type ReCreatgroup struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	GroupId      string
}

type GetGroup struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	TotalCount   int64
	GroupIdList  []Groupid
	Next         int64
}

type Groupid struct {
	GroupId string
}

type AddGroupMember struct {
	GroupId    string
	MemberList []MemberAccount
}

type ReAddGroupMember struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	MemberList   []MemberAccount
}

type MemberAccount struct {
	Member_Account string
}

type BatchAddFriend struct {
	From_Account  string
	AddFriendItem []AddFriendItem
}

type AddFriendItem struct {
	To_Account string
	AddSource  string
}

type ReBatchAddFriend struct {
	ResultItem      []ReAddFriendItem
	Fail_Account    []string
	Invalid_Account string
	ErrorCode       int64
	ErrorInfo       string
	ErrorDisplay    string
}

type ReAddFriendItem struct {
	To_Account string
	ResultCode int64
	ResultInfo string
}

type DeleteFriendAll struct {
	From_Account string
}

type SendC2CMes struct {
	SyncOtherMachine int
	To_Account       []string
	MsgRandom        int32
	MsgBody          []interface{}
}

type MsgBodyText struct {
	MsgType    string
	MsgContent MsgContentText
}

type MsgContentText struct {
	Text string
}

type MsgBodyFace struct {
	MsgType    string
	MsgContent MsgContentFace
}

type MsgContentFace struct {
	Index int
	Data  string
}

func (msb *SendC2CMes) Add(elem interface{}) []interface{} {
	//msgBody := msb.MsgBody
	msb.MsgBody = append(msb.MsgBody, elem)
	fmt.Printf("msgBody--%v\n", msb.MsgBody)

	return msb.MsgBody
}

type ReSendC2CMes struct {
	ErrorInfo    string
	ActionStatus string
	ErrorCode    int64
}

//type ReSendC2CMes struct {
//	ActionStatus string
//	ErrorInfo string
//	ErrorList []ErrorList
//}
//
//type ErrorList struct {
//	To_Account string
//	ErrorCode int64
//}

type SendGroupMes struct {
	GroupId string
	Random  int32
	MsgBody []interface{}
}

func (msb *SendGroupMes) Add(elem interface{}) []interface{} {
	//msgBody := msb.MsgBody
	msb.MsgBody = append(msb.MsgBody, elem)
	fmt.Printf("msgBody--%v\n", msb.MsgBody)

	return msb.MsgBody
}

type ReSendGroupMes struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	MsgTime      int64
	MsgSeq       int64
}

type PullFriendList struct {
	From_Account string
	TimeStamp	 int64
	StartIndex   int64
}

type ReFriendList struct {
	StartIndex   int64
	FriendNum    int64
	ActionStatus string
	ErrorCode    int64
	ErrorInfo    string
	ErrorDisplay string
}

type SendSysMsg struct {
	GroupId string
	Content string
}

type ReSysMsg struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
}

type DelGroup struct {
	GroupId string
}

type ReDelGroup struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
}

type GetGroupIdList struct {
	GroupIdList []string
}

type ReGetGroupIdList struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	GroupInfo	 []GroupNameById
}

type GroupNameById struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int64
	GroupId 	 string
	Name		 string
}

var allAccountsName []string
//var (
//	logFile, err = os.OpenFile("./IMTools/go.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
//	printLog     = log.New(logFile, "[print]", log.Ldate|log.Ltime|log.Llongfile)
//)

/**
功能：在指定群组发送系统通知
参数：userSig——用户签名,groupId——目标群，content要发送的通知内容
返回值：好友数，错误码
*/
func SendSystemMsg(userSig string, groupId string, content string) int64 {
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/send_group_system_notification?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//var errorCode int64

	var sendSysMsg = SendSysMsg{}
	sendSysMsg.GroupId = groupId
	sendSysMsg.Content = content

	//封装json应答包
	re, err := json.Marshal(sendSysMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sendSysMsg request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, string(re))
	fmt.Printf("sendSysMsg--%v\nerr--%v\n", replydata, err)
	fmt.Printf("SendSystemMsg_httpUrl--%v",httpUrl)

	reSendSysMsg := ReSysMsg{}
	json.Unmarshal([]byte(replydata), &reSendSysMsg)

	errorCode := reSendSysMsg.ErrorCode

	return errorCode

}

/**
功能：拉取好友列表
参数：userSig——用户签名,userId——目标用户
返回值：好友数，错误码
*/
func GetFriendList(userSig string, userId string) (int64, int64) {
	httpUrl := "https://test.tim.qq.com/v4/sns/friend_get_all?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//var errorCode int64

	var getFriendAll = PullFriendList{}
	getFriendAll.From_Account = userId
	getFriendAll.StartIndex = 0
	getFriendAll.TimeStamp = 0

	//封装json应答包
	re, err := json.Marshal(getFriendAll)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GetFriendAll request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, string(re))
	fmt.Printf("GetFriendAll--%v\nerr--%v\n", replydata, err)
	fmt.Printf("GetFriendAll_url--%v\n",httpUrl)

	reFriendLis := ReFriendList{}
	json.Unmarshal([]byte(replydata), &reFriendLis)

	errorCode := reFriendLis.ErrorCode
	friendNum := reFriendLis.FriendNum

	return friendNum, errorCode

}

/**
功能：批量发群消息
参数：userSig——用户签名,userNum——要群发的用户数目
返回值：错误码
*/
func SendGroupMsg(userSig string, groupNum int) int64 {
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/send_group_msg?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	var errorCode int64

	sendGroupMes := SendGroupMes{}
	msgBodyText := MsgBodyText{}
	msgBodyFace := MsgBodyFace{}
	msgContentText := MsgContentText{}
	msgContentFace := MsgContentFace{}

	sendGroupMes.Random = rand.Int31()

	msgContentText.Text = "Hello Groups !"
	msgBodyText.MsgType = "TIMTextElem"
	msgBodyText.MsgContent = msgContentText
	sendGroupMes.Add(msgBodyText)

	msgContentFace.Index = 6
	msgContentFace.Data = "abc\u0000\u0001"
	msgBodyFace.MsgType = "TIMFaceElem"
	msgBodyFace.MsgContent = msgContentFace
	sendGroupMes.Add(msgBodyFace)

	fmt.Printf("sendGroupMes--%v\n", sendGroupMes)

	groupIdArray := GetAllGroup(userSig)

	for i := 0; i < groupNum; i++ {
		sendGroupMes.GroupId = groupIdArray[i]
		re, err := json.Marshal(sendGroupMes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("sendGroupMes request json--%s\n", re)

		//访问IM后台
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("sendGroupMes--%v\nerr--%v\n", replydata, err)

		reSendGroupMes := ReSendGroupMes{}
		json.Unmarshal([]byte(replydata), &reSendGroupMes)

		errorCode = reSendGroupMes.ErrorCode
	}
	return errorCode
}

/**
功能：批量发单聊消息
参数：userSig——用户签名,userNum——要群发的用户数目
返回：失败用户信息
*/
func SendC2CMsg(userSig string, userNum int) int64 {
	httpUrl := "https://console.tim.qq.com/v4/openim/batchsendmsg?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	var userNameArray []string
	var errorCode int64

	//假设一次对最多5个用户进行单发消息
	numLimit := 5
	if userNum >= numLimit {
		temp := userNum % numLimit
		if temp == 0 { //要发的用户数刚好是上限的整数倍

			for i := 0; i < userNum; i = i + numLimit {

				userNameArray = append(userNameArray, packC2CMsg(i, numLimit)...)
				errorCode = Post_SendC2CMsg(httpUrl, userNameArray)
				//清空临时数组
				userNameArray = userNameArray[:0]

			}
		} else {
			remaind := userNum - temp
			for i := 0; i < remaind; i = i + numLimit {
				userNameArray = append(userNameArray, packC2CMsg(i, numLimit)...)
				errorCode = Post_SendC2CMsg(httpUrl, userNameArray)
				//清空临时数组
				userNameArray = userNameArray[:0]
			}
			userNameArray = append(userNameArray, packC2CMsg(remaind, temp)...)
			errorCode = Post_SendC2CMsg(httpUrl, userNameArray)

		}
	} else {
		userNameArray = packC2CMsg(0, userNum)
		errorCode = Post_SendC2CMsg(httpUrl, userNameArray)

	}

	return errorCode

}

func packC2CMsg(index int, userNum int) []string {

	var userNameArray []string
	for j := index; j < index+userNum; j++ {
		userName := "user" + strconv.Itoa(j+1)
		userNameArray = append(userNameArray, userName)
	}

	return userNameArray

}

func Post_SendC2CMsg(url string, userNameArray []string) int64 {
	sendC2CMes := SendC2CMes{}
	msgBodyText := MsgBodyText{}
	msgBodyFace := MsgBodyFace{}
	msgContentText := MsgContentText{}
	msgContentFace := MsgContentFace{}

	sendC2CMes.To_Account = userNameArray
	sendC2CMes.SyncOtherMachine = 2
	sendC2CMes.MsgRandom = rand.Int31()

	msgContentText.Text = "Hello everbody !"
	msgBodyText.MsgType = "TIMTextElem"
	msgBodyText.MsgContent = msgContentText
	sendC2CMes.Add(msgBodyText)

	msgContentFace.Index = 6
	msgContentFace.Data = "content"
	msgBodyFace.MsgType = "TIMFaceElem"
	msgBodyFace.MsgContent = msgContentFace
	sendC2CMes.Add(msgBodyFace)

	fmt.Printf("sendC2CMes--%v\n", sendC2CMes)

	re, err := json.Marshal(sendC2CMes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("sendC2CMes request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(url, string(re))
	fmt.Printf("SendC2CMes--%v\nerr--%v\n", replydata, err)

	reSendC2CMes := ReSendC2CMes{}
	json.Unmarshal([]byte(replydata), &reSendC2CMes)

	errorCode := reSendC2CMes.ErrorCode
	//errorList := reSendC2CMes.ErrorList
	//fmt.Printf("reSendC2CMes--%v\n",reSendC2CMes)

	return errorCode
}

//删除指定用户的所有好友
func DeleteFriend(userSig string, userId string) {
	httpUrl := "https://console.tim.qq.com/v4/sns/friend_delete_all?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	var deleteFriendAll = DeleteFriendAll{}
	deleteFriendAll.From_Account = userId

	//封装json应答包
	re, err := json.Marshal(deleteFriendAll)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteFriendAll request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, string(re))
	fmt.Printf("DeleteFriendAll--%v\nerr--%v\n", replydata, err)
}

/**
功能：批量添加好友——指定一个用户Id，加指定个数的好友（没排除自己，自己也可以是自己的一个好友）
参数：userSig——用户签名,groupId——群id集合，accoutNumOfgroup——群组中需要添加的用户数量,allAccountsName——要添加的所有账户名
返回值：错误码
*/
func AddFriend(userSig string, userId string,friendNumFrom int,friendNumTo int) int64 {
	//httpUrl := "https://console.tim.qq.com/v4/sns/friend_add?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"
	httpUrl := "https://console.tim.qq.com/v4/sns/friend_add?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"


	var batchAddFriend = BatchAddFriend{}
	var addFriendItem = AddFriendItem{}
	var friendArray []AddFriendItem
	var errorCode int64
	var friendLimit = 1000

	//通过批量添加账号获取所有用户名集合，这里默认添加账户数为系统上限（100）
	//var userIdArray, _ = Multiaccount_PostData(userSig, friendNum)

	userIdArray := allAccountsName

	if friendNumTo-friendNumFrom+1 > friendLimit { //如果人数超过1000
		errorCode = 1000
	}else{
		for i := 0; i <= friendNumTo - friendNumFrom; i++ {
			//if i == friendNum { //如果添加的好友为其本身，则添加其序号后一位的user
			//	addFriendItem.To_Account = userIdArray[i+1]
			//}else {
			addFriendItem.To_Account = userIdArray[friendNumFrom - 1 + i]
			//}

			addFriendItem.AddSource = "AddSource_Type_Android" //默认好友来源都为Android
			friendArray = append(friendArray, addFriendItem)

		}

		batchAddFriend.From_Account = userId
		batchAddFriend.AddFriendItem = friendArray

		//封装json应答包
		re, err := json.Marshal(batchAddFriend)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("BatchAddFriend request json--%s\n", re)
		fmt.Printf("BatchAddFriend_url--%s\n", httpUrl)


		//访问IM后台
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("BatchAddFriend--%v\nerr--%v\n", replydata, err)

		reBatchAddFriend := ReBatchAddFriend{}
		json.Unmarshal([]byte(replydata), &reBatchAddFriend)

		errorCode = reBatchAddFriend.ErrorCode
	}

	//for i := 0; i < friendNum; i++ {
	//	//if i == friendNum { //如果添加的好友为其本身，则添加其序号后一位的user
	//	//	addFriendItem.To_Account = userIdArray[i+1]
	//	//}else {
	//	addFriendItem.To_Account = userIdArray[i]
	//	//}
	//
	//	addFriendItem.AddSource = "AddSource_Type_Android" //默认好友来源都为Android
	//	friendArray = append(friendArray, addFriendItem)
	//
	//}



	return errorCode

}

//func AddFriend(userSig string, userId string,friendNum int,friendNumFrom int,friendNumTo int) int64 {
//	httpUrl := "https://console.tim.qq.com/v4/sns/friend_add?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"
//
//	var batchAddFriend = BatchAddFriend{}
//	var addFriendItem = AddFriendItem{}
//	var friendArray []AddFriendItem
//	var errorCode int64
//	var friendLimit = 1000
//
//	//通过批量添加账号获取所有用户名集合，这里默认添加账户数为系统上限（100）
//	var userIdArray, _ = Multiaccount_PostData(userSig, friendNum)
//
//	if friendNumTo-friendNumFrom+1 > friendLimit { //如果人数超过1000
//		errorCode = 1000
//	}else{
//		for i := 0; i <= friendNumTo - friendNumFrom; i++ {
//			//if i == friendNum { //如果添加的好友为其本身，则添加其序号后一位的user
//			//	addFriendItem.To_Account = userIdArray[i+1]
//			//}else {
//			addFriendItem.To_Account = userIdArray[friendNumFrom - 1 + i]
//			//}
//
//			addFriendItem.AddSource = "AddSource_Type_Android" //默认好友来源都为Android
//			friendArray = append(friendArray, addFriendItem)
//
//		}
//
//		batchAddFriend.From_Account = userId
//		batchAddFriend.AddFriendItem = friendArray
//
//		//封装json应答包
//		re, err := json.Marshal(batchAddFriend)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("BatchAddFriend request json--%s\n", re)
//		fmt.Printf("BatchAddFriend_url--%s\n", httpUrl)
//
//
//		//访问IM后台
//		replydata, err := HTTP_Post(httpUrl, string(re))
//		fmt.Printf("BatchAddFriend--%v\nerr--%v\n", replydata, err)
//
//		reBatchAddFriend := ReBatchAddFriend{}
//		json.Unmarshal([]byte(replydata), &reBatchAddFriend)
//
//		errorCode = reBatchAddFriend.ErrorCode
//	}
//
//	//for i := 0; i < friendNum; i++ {
//	//	//if i == friendNum { //如果添加的好友为其本身，则添加其序号后一位的user
//	//	//	addFriendItem.To_Account = userIdArray[i+1]
//	//	//}else {
//	//	addFriendItem.To_Account = userIdArray[i]
//	//	//}
//	//
//	//	addFriendItem.AddSource = "AddSource_Type_Android" //默认好友来源都为Android
//	//	friendArray = append(friendArray, addFriendItem)
//	//
//	//}
//
//
//
//	return errorCode
//
//}

/**
功能：增加群组成员
参数：userSig——用户签名,groupId——群id集合，groupId—目标群组Id（后期需从前端获取），accoutNumOfgroup——群组中需要添加的用户数量（后期需从前端获取）,allAccountsName——要添加的所有账户名
返回值：错误码
*/
func AddGroupAccount(userSig string, groupId string, accoutNumOfgroup int) int64 { //, allAccountsName []string
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/add_group_member?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//假设群成员上限为2，每次最多添加账号数为2
	var memberLimit = 2
	var groupMember = AddGroupMember{}
	var memberAccount = MemberAccount{}
	var memberArry []MemberAccount
	var errorCode int64

	if accoutNumOfgroup >= memberLimit {
		temp := accoutNumOfgroup % memberLimit
		if temp == 0 { //要发的用户数刚好是上限的整数倍

			for i := 0; i < accoutNumOfgroup; i = i + memberLimit {

				memberArry = append(memberArry, dealGroupAccount(i, memberLimit, memberAccount)...)
				errorCode = Post_AddGroupAccount(httpUrl, groupId, memberArry, groupMember)
				//清空账号结构体
				memberArry = memberArry[:0]

			}
		} else {
			remaind := accoutNumOfgroup - temp
			for i := 0; i < remaind; i = i + memberLimit {
				memberArry = append(memberArry, dealGroupAccount(i, memberLimit, memberAccount)...)
				errorCode = Post_AddGroupAccount(httpUrl, groupId, memberArry, groupMember)
				//清空账号结构体
				memberArry = memberArry[:0]

			}
			memberArry = append(memberArry, dealGroupAccount(remaind, temp, memberAccount)...)
			errorCode = Post_AddGroupAccount(httpUrl, groupId, memberArry, groupMember)

		}
	} else {
		memberArry = dealGroupAccount(0, accoutNumOfgroup, memberAccount)
		errorCode = Post_AddGroupAccount(httpUrl, groupId, memberArry, groupMember)
	}

	return errorCode

	//初始化数据结构体

	//groupMember = AddGroupMember{
	//	GroupId:    groupId,
	//	MemberList: memberArry,
	//}
	//fmt.Printf("groupMember--%v\n", groupMember)
	//
	////封装json应答包
	//re, err := json.Marshal(groupMember)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("AddGroupMember request json--%s\n", re)
	//
	////访问IM后台
	//replydata, err := HTTP_Post(httpUrl, string(re))
	//fmt.Printf("AddGroupMember--%v\nerr--%v\n", replydata, err)
	//
	//reAddGroupMember := ReAddGroupMember{}
	//json.Unmarshal([]byte(replydata), &reAddGroupMember)
	//
	//errorCode = reAddGroupMember.ErrorCode
	//return errorCode

	//清空账号结构体
	//memberArry = memberArry[:0]

}

func Post_AddGroupAccount(url string, groupId string, memberArry []MemberAccount, groupMember AddGroupMember) int64 {

	groupMember = AddGroupMember{
		GroupId:    groupId,
		MemberList: memberArry,
	}
	fmt.Printf("groupMember--%v\n", groupMember)

	//封装json应答包
	re, err := json.Marshal(groupMember)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AddGroupMember request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(url, string(re))
	fmt.Printf("AddGroupMember--%v\nerr--%v\n", replydata, err)

	reAddGroupMember := ReAddGroupMember{}
	json.Unmarshal([]byte(replydata), &reAddGroupMember)

	errorCode := reAddGroupMember.ErrorCode

	return errorCode
}

//}

func dealGroupAccount(index int, userNum int, memberAccount MemberAccount) []MemberAccount {

	var userNameArray []MemberAccount
	for j := index; j < index+userNum; j++ {

		memberAccount.Member_Account = "user" + strconv.Itoa(j+1)
		userNameArray = append(userNameArray, memberAccount)
	}

	return userNameArray

}

/**
功能：获取所有群组
参数：userSig——用户签名
返回值：群ID集合
*/

func GetAllGroup(userSig string) []string {
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/get_appid_group_list?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, "{}")
	fmt.Printf("CreatgroupReplyData--%v\nerr--%v\n", replydata, err)

	//解析应答包
	groupInfo := GetGroup{}

	error := json.Unmarshal([]byte(replydata), &groupInfo)
	if err != nil {
		fmt.Printf("Release groupInfo fail:%v\n", error)
	}
	fmt.Printf("Get all groupInfo : %v\n", groupInfo)

	//取出所有groupId
	var groupIdArry []string
	groupId := groupInfo.GroupIdList
	for i := 0; i < len(groupId); i++ {
		groupIdArry = append(groupIdArry, groupId[i].GroupId)
	}

	return groupIdArry
}

/**
功能：批量创建群组
参数：userSig——用户签名,groupNum——需要添加的群组数量（后期需从前端获取）,accountName——指定群主的账户名
返回值：错误码
*/
func BatchCreatgroup(userSig string, groupNum int, accountName string) int64 {

	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/create_group?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	var errorCode int64

	for i := 0; i < groupNum; i++ {

		name := "GroupOf" + accountName
		//创建组
		creatGroup := Creatgroup{
			Owner_Account: accountName,
			Type:          "Public",
			Name:          name,
		}

		//封装json请求包
		re, err := json.Marshal(creatGroup)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Multiaccount request json--%s\n", re)

		//访问IM后台
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("CreatgroupReplyData--%v\nerr--%v\n", replydata, err)

		reCreatgroup := ReCreatgroup{}
		json.Unmarshal([]byte(replydata), &reCreatgroup)

		errorCode = reCreatgroup.ErrorCode
	}
	return errorCode
}

///**
//功能：批量创建群组
//参数：userSig——用户签名,groupNum——需要添加的群组数量（后期需从前端获取）,allAccountsName——要添加的所有账户名
//返回值：错误码
//*/
//func BatchCreatgroup(userSig string, groupNum int, allAccountName []string) int64 { //accoutNumOfgroup int
//
//	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/create_group?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"
//
//	var errorCode int64
//
//	for i := 0; i < groupNum; i++ {
//
//		name := "GroupOf" + allAccountName[i]
//		//创建组
//		creatGroup := Creatgroup{
//			Owner_Account: allAccountName[i],
//			Type:          "Public",
//			Name:          name,
//		}
//
//		//封装json请求包
//		re, err := json.Marshal(creatGroup)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("Multiaccount request json--%s\n", re)
//
//		//访问IM后台
//		replydata, err := HTTP_Post(httpUrl, string(re))
//		fmt.Printf("CreatgroupReplyData--%v\nerr--%v\n", replydata, err)
//
//		reCreatgroup := ReCreatgroup{}
//		json.Unmarshal([]byte(replydata), &reCreatgroup)
//
//		errorCode = reCreatgroup.ErrorCode
//	}
//	return errorCode
//}

/**
功能：批量添加账户添加
参数：userSig——用户签名,accountsnum——需要添加的账号数量（后期需从前端获取）
返回值：所有账户的用户名,错误码
*/
func Multiaccount_PostData(userSig string, accountsnum int) ([]string, int64) {
	httpUrl := "https://console.tim.qq.com/v4/im_open_login_svc/multiaccount_import?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//单次请求导入账号数上限：100——需要添加的总账号数：accountsnum = 100
	var accountLimit = 100
	var multiaccount = Multiaccount{}
	var accounts []string
	var allAccounts []string
	var errorCode int64

	if accountsnum >= accountLimit {
		temp := accountsnum % accountLimit
		if temp == 0 { //要发的用户数刚好是上限的整数倍

			for i := 0; i < accountsnum; i = i + accountLimit {

				accounts = append(accounts, packC2CMsg(i, accountLimit)...)
				errorCode = Post_Multiaccount(httpUrl, accounts, multiaccount)
				//记录所有账号 供后续获取全部成功的账号
				allAccounts = append(allAccounts, accounts...)
				//清空临时数组
				accounts = accounts[:0]

			}
		} else {
			remaind := accountsnum - temp
			for i := 0; i < remaind; i = i + accountLimit {
				accounts = append(accounts, packC2CMsg(i, accountLimit)...)
				errorCode = Post_Multiaccount(httpUrl, accounts, multiaccount)
				//记录所有账号 供后续获取全部成功的账号
				allAccounts = append(allAccounts, accounts...)
				//清空临时数组
				accounts = accounts[:0]
			}
			accounts = append(accounts, packC2CMsg(remaind, temp)...)
			errorCode = Post_Multiaccount(httpUrl, accounts, multiaccount)
			//记录所有账号 供后续获取全部成功的账号
			allAccounts = append(allAccounts, accounts...)

		}
	} else {
		accounts = packC2CMsg(0, accountsnum)
		errorCode = Post_Multiaccount(httpUrl, accounts, multiaccount)
		//记录所有账号 供后续获取全部成功的账号
		allAccounts = append(allAccounts, accounts...)

	}

	allAccountsName = allAccounts
	return allAccounts, errorCode

	////记录所有账号 供后续获取全部成功的账号
	//allAccounts = append(allAccounts, accounts...)

	//var num = 0
	//for i := 0; i < accountsnum; i = i + accountLimit {
	//
	//	for j := 1; j <= accountLimit; j++ { //随机产生10个用户
	//		num++
	//		userName := "user" + strconv.Itoa(num)
	//		accounts = append(accounts, userName)
	//	}
	//
	//	//记录所有账号 供后续获取全部成功的账号
	//	allAccounts = append(allAccounts, accounts...)
	//
	//	//产生完10个用户后发起请求
	//	multiaccount = Multiaccount{Accounts: accounts}
	//	fmt.Printf("rebody--%v\n", multiaccount)
	//
	//	//清空账号结构体
	//	accounts = accounts[:0]
	//
	//	//封装json请求包
	//	re, err := json.Marshal(multiaccount)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Printf("Multiaccount request json--%s\n", re)
	//
	//	//访问IM后台
	//	replydata, err := HTTP_Post(httpUrl, string(re))
	//	fmt.Printf("AccountReplyData--%v\nerr--%v\n", replydata, err)
	//
	//	reMultiaccount := ReMultiaccount{}
	//	json.Unmarshal([]byte(replydata), &reMultiaccount)
	//
	//	errorCode = reMultiaccount.ErrorCode

	//
	//}

	//这里应该加一个去除allAccounts中失败用户的功能（后续加）
	//AllAccountsId = allAccounts
	//fmt.Printf("allAccountsName--%v\n", allAccounts)

	//return allAccounts,errorCode

}

func Post_Multiaccount(url string, accounts []string, multiaccount Multiaccount) int64 {
	//产生完10个用户后发起请求
	multiaccount = Multiaccount{Accounts: accounts}
	fmt.Printf("rebody--%v\n", multiaccount)

	//封装json请求包
	re, err := json.Marshal(multiaccount)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Multiaccount request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(url, string(re))
	fmt.Printf("AccountReplyData--%v\nerr--%v\n", replydata, err)

	reMultiaccount := ReMultiaccount{}
	json.Unmarshal([]byte(replydata), &reMultiaccount)

	errorCode := reMultiaccount.ErrorCode

	return errorCode
}

/**
参数：功能URL，json请求包
返回值：json应到包，错误
*/
func HTTP_Post(url string, reqbody string) (string, error) {

	var result string

	//创建请求
	postReq, err := http.NewRequest("POST", url, strings.NewReader(string(reqbody)))
	if err != nil {
		fmt.Println("POST请求:创建请求失败", err)
	}

	//增加header
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")

	//执行请求
	client := http.Client{}
	resp, err := client.Do(postReq)

	if err != nil {
		fmt.Println("POST请求:创建执行请求失败", err)
	} else {
		//读取请求
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("POST请求:读取body失败", err)
		}

		result = string(body)
		fmt.Println("POST请求:创建成功", result)

	}

	defer resp.Body.Close()

	return result, err

}

/**
功能：解散所有群组
参数：userSig——用户签名,sdkappid,groupId
返回值：错误码
*/

func DeleteGroup(userSig string,sdkappid int,identifier string,groupId string)  int64{
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/destroy_group?usersig=" + userSig + "&identifier=" + identifier + "&sdkappid=" + strconv.Itoa(sdkappid) + "&random=99999999&contenttype=json"
	//httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/get_appid_group_list?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.FormatInt(sdkappid, 10) + "&random=99999999&contenttype=json"

	//groupIdArry := GetAllGroup(userSig)
	var errorCode int64

		var deleteGroup = DelGroup{}
		deleteGroup.GroupId = groupId

		//封装json应答包
		re, err := json.Marshal(deleteGroup)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("DeleteGroup request json--%s\n", re)

		//访问IM后台
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("DeleteGroup--%v\nerr--%v\n", replydata, err)

		//解析应答包
		reDelGroup := ReDelGroup{}

		json.Unmarshal([]byte(replydata), &reDelGroup)

		errorCode = reDelGroup.ErrorCode

		return errorCode


}

/**
功能：获取群组中固定格式的群组名
参数：userSig——用户签名,sdkappid
返回值：错误码
*/

func DeleteNameGroup(userSig string,sdkappid int,identifier string) int64 {
	//httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/get_group_info?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/get_group_info?usersig=" + userSig + "&identifier=" + identifier + "&sdkappid=" + strconv.Itoa(sdkappid) + "&random=99999999&contenttype=json"

	//获取所有群组idlist
	groupIdArry := GetAllGroup(userSig)
	getGroupIdList := GetGroupIdList{}
	getGroupIdList.GroupIdList = groupIdArry

	var errorCode int64

	//封装json应答包
	re, err := json.Marshal(getGroupIdList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteNameGroup request json--%s\n", re)

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, string(re))
	fmt.Printf("DeleteNameGroup--%v\nerr--%v\n", replydata, err)

	//解析应答包
	pattern := "autotest_"

	var groupNameByIdList []GroupNameById
	groupNameById := GroupNameById{}

	reGetGroupIdList := ReGetGroupIdList{}

	json.Unmarshal([]byte(replydata), &reGetGroupIdList)

	groupNameByIdList = reGetGroupIdList.GroupInfo

	for i := 0; i < len(groupNameByIdList);i++ {

		groupNameById.GroupId = groupNameByIdList[i].GroupId
		groupNameById.Name = groupNameByIdList[i].Name
		if strings.HasPrefix(groupNameById.Name, pattern) {
			DeleteGroup(userSig,sdkappid,identifier,groupNameById.GroupId)
		}else {
			continue
		}



	}
	errorCode = reGetGroupIdList.ErrorCode
	return errorCode

}
