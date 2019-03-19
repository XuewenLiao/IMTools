package main

import "IMTools/websever"

func main() {

	router := websever.InitRouter()
	router.Run(":8000")

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
