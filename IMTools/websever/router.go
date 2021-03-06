package websever

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//跨域
	router.Use(Cors())
	//批量建群url——http://127.0.0.1:8000/batchcreatgroup？accountsnum=xxx&accountname=xxx&groupnum=xxx
	router.POST("/batchcreatgroup", BatchCreatGroup)

	//批量加群url——http://127.0.0.1:8000/batchaddgroup？groupid=xxx&accoutnumofgroup=xxx
	router.POST("/batchaddgroup", BatchAddGroup)

	//批量加好友url——http://127.0.0.1:8000/batchaddfriend？userid=xxx&friendnum=xxx&friendNumFrom=xxx&friendnumto=xxx
	router.POST("/batchaddfriend", BatchAddFriend)

	//批量发单聊信息url——http://127.0.0.1:8000/batchsendc2cmsg？usernum=xxx
	router.POST("/batchsendc2cmsg", BatchSendC2CMsg)

	//批量发群聊信息url——http://127.0.0.1:8000/batchsendgroupmsg？groupnum=xxx
	router.POST("/batchsendgroupmsg", BatchSendGroupMsg)

	//拉取好友列表url——http://127.0.0.1:8000/getfriendlist？userfrdid=xxx
	router.POST("/getfriendlist", GetFriendList)

	//在指定群组发送系统通知url——http://127.0.0.1:8000/sendgroupsysmsg？groupname=xxx&content=xxx
	router.POST("/sendgroupsysmsg", SendGroupSysMsg)

	//批量添加账户url——http://127.0.0.1:8000/batchaddaccounts？allaccountsnum=xxx
	router.POST("/batchaddaccounts", BatchAddAccounts)

	//批量添加账户url——http://127.0.0.1:8000/deletegroupbyname？sdkappid=xxx&identifier=xxx
	router.POST("/deletegroupbyname", DeleteGroupByName)


	return router
}

//跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "text/plain")                                                                                                                                                                    // 设置返回格式是json text/plain
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
