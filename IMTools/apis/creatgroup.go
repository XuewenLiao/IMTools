package apis

import (
	"IMTools/sdkconst"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var AllAccountsId []string

type Multiaccount struct {
	Accounts []string
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

type ReplyOfCreatgroup struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int
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

type DeleteFriendAll struct {
	From_Account string
}

type SendC2CMes struct {
	SyncOtherMachine int
	To_Account       []string
	MsgRandom        int
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
参数：userSig——用户签名,groupId——群id集合，accoutNumOfgroup——群组中需要添加的用户数量（后期需从前端获取）,allAccountsName——要添加的所有账户名
返回值：URL和请求包
*/
func AddFriend(userSig string, userId string, friendNum int) {
	httpUrl := "https://console.tim.qq.com/v4/sns/friend_add?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	var batchAddFriend = BatchAddFriend{}
	var addFriendItem = AddFriendItem{}
	var friendArray []AddFriendItem

	//通过批量添加账号获取所有用户名集合，这里默认添加账户数为系统上限（100）
	var userIdArray = Multiaccount_PostData(userSig, 100)

	for i := 0; i < friendNum; i++ {
		//if i == friendNum { //如果添加的好友为其本身，则添加其序号后一位的user
		//	addFriendItem.To_Account = userIdArray[i+1]
		//}else {
		addFriendItem.To_Account = userIdArray[i]
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

	//访问IM后台
	replydata, err := HTTP_Post(httpUrl, string(re))
	fmt.Printf("BatchAddFriend--%v\nerr--%v\n", replydata, err)

}

/**
功能：增加群组成员
参数：userSig——用户签名,groupId——群id集合，groupId—目标群组Id（后期需从前端获取），accoutNumOfgroup——群组中需要添加的用户数量（后期需从前端获取）,allAccountsName——要添加的所有账户名
返回值：URL和请求包
*/
func AddGroupAccount(userSig string, groupId string, accoutNumOfgroup int, allAccountsName []string) {
	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/add_group_member?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//假设群成员上限为2，每次最多添加账号数为2
	var memberLimit = 2
	var groupMember = AddGroupMember{}
	var memberAccount = MemberAccount{}
	var memberArry []MemberAccount

	for i := 0; i < accoutNumOfgroup; i = i + memberLimit {
		for j := i; j < i+memberLimit; j++ {
			memberAccount.Member_Account = allAccountsName[j]
			memberArry = append(memberArry, memberAccount)

		}

		//初始化数据结构体

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
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("AddGroupMember--%v\nerr--%v\n", replydata, err)

		//清空账号结构体
		memberArry = memberArry[:0]

	}

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
参数：userSig——用户签名,groupNum——需要添加的群组数量（后期需从前端获取）,allAccountsName——要添加的所有账户名
返回值：URL和请求包
*/
func BatchCreatgroup(userSig string, groupNum int, allAccountName []string) { //accoutNumOfgroup int

	httpUrl := "https://console.tim.qq.com/v4/group_open_http_svc/create_group?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	for i := 0; i < groupNum; i++ {

		name := "GroupOf" + allAccountName[i]
		//创建组
		creatGroup := Creatgroup{
			Owner_Account: allAccountName[i],
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
	}

	//解析应答包，获取groupId
	//groupInfo := ReplyOfCreatgroup{}
	//error := json.Unmarshal([]byte(replydata), &groupInfo)
	//if err != nil {
	//	fmt.Printf("Release groupInfo fail:%v\n", error)
	//}
	//fmt.Printf("Get the groupInfo : %v\n", groupInfo)
	//
	//groupId := groupInfo.GroupId

	//获取所有groupId
	//groupIdArry := GetAllGroup(userSig)
	//fmt.Printf("GroupIdArry--%v\n", groupIdArry)

	//添加群组成员
	//AddGroupAccount(userSig, groupId, accoutNumOfgroup, allAccountName)
}

/**
功能：批量添加账户添加
参数：userSig——用户签名,accountsnum——需要添加的账号数量（后期需从前端获取）
返回值：所有账户的用户名
*/
func Multiaccount_PostData(userSig string, accountsnum int) []string {
	httpUrl := "https://console.tim.qq.com/v4/im_open_login_svc/multiaccount_import?usersig=" + userSig + "&identifier=" + sdkconst.Identifier + "&sdkappid=" + strconv.Itoa(sdkconst.Appid) + "&random=99999999&contenttype=json"

	//假设单次请求导入账号数上限：10——需要添加的总账号数：accountsnum = 100
	var accountLimit = 10
	var multiaccount = Multiaccount{}
	var accounts []string
	var allAccounts []string

	var num = 0
	for i := 0; i < accountsnum; i = i + accountLimit {

		for j := 1; j <= accountLimit; j++ { //随机产生10个用户
			num++
			userName := "user" + strconv.Itoa(num)
			accounts = append(accounts, userName)
		}

		//记录所有账号 供后续获取全部成功的账号
		allAccounts = append(allAccounts, accounts...)

		//产生完10个用户后发起请求
		multiaccount = Multiaccount{Accounts: accounts}
		fmt.Printf("rebody--%v\n", multiaccount)

		//清空账号结构体
		accounts = accounts[:0]

		//封装json请求包
		re, err := json.Marshal(multiaccount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Multiaccount request json--%s\n", re)

		//访问IM后台
		replydata, err := HTTP_Post(httpUrl, string(re))
		fmt.Printf("AccountReplyData--%v\nerr--%v\n", replydata, err)

	}

	//这里应该加一个去除allAccounts中失败用户的功能（后续加）
	AllAccountsId = allAccounts
	fmt.Printf("allAccountsName--%v\n", allAccounts)

	return allAccounts

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
