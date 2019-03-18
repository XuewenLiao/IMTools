package main

import "fmt"

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

func (msb *SendC2CMes) Add(elem interface{}) interface{} {
	msgBody := msb.MsgBody
	msgBody = append(msgBody, elem)
	fmt.Printf("msgBody--%v\n", msgBody)
	return msgBody
}

func main() {

	sendC2CMes := SendC2CMes{}
	msgBodyText := MsgBodyText{}
	msgBodyFace := MsgBodyFace{}
	msgContentText := MsgContentText{}
	msgContentFace := MsgContentFace{}

	sendC2CMes.SyncOtherMachine = 2
	sendC2CMes.To_Account = []string{"user1", "user2"}
	sendC2CMes.MsgRandom = 123

	msgContentText.Text = "red packet"
	msgBodyText.MsgType = "TIMTextElem"
	msgBodyText.MsgContent = msgContentText

	sendC2CMes.Add(msgBodyText)

	msgContentFace.Index = 6
	msgContentFace.Data = "content"
	msgBodyFace.MsgType = "TIMFaceElem"
	msgBodyFace.MsgContent = msgContentFace

	sendC2CMes.Add(msgBodyFace)

	fmt.Printf("sendC2CMes--%v\n", sendC2CMes)

	//router := websever.InitRouter()
	//router.Run(":8000")

	////生成userSig(注意expire：签名过期时间，设置小了请求几次就会抱7001签名过期错误)
	//userSig, _ := TLSSigAPI.GenerateUsersigWithExpire(sdkconst.PrivateKey, sdkconst.Appid, sdkconst.Identifier, 60*60*24*180)
	////fmt.Printf("userSig: %v",userSig)
	//
	////批量添加账户
	//allAccountsName := apis.Multiaccount_PostData(userSig, 100)
	//
	////生成群组
	//apis.Creatgroup_PostData(userSig, 20, allAccountsName)

}
