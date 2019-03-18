package websever

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	//批量建群url——http://127.0.0.1:8000/batchcreatgroup？accountsnum=xxx&groupnum=xxx
	router.POST("/batchcreatgroup", BatchCreatGroup)

	//批量加群url——http://127.0.0.1:8000/batchaddgroup？groupid=xxx&accoutnumofgroup=xxx
	router.POST("/batchaddgroup", BatchAddGroup)

	//批量加好友url——http://127.0.0.1:8000/batchaddfriend？userid=xxx&friendnum=xxx
	router.POST("/batchaddfriend", BatchAddFriend)

	return router
}
