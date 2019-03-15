package websever

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()

	//创建群组url——http://127.0.0.1:8000/creatgroup？accountsnum=xxx&accoutnumofgroup=xxx
	router.POST("/creatgroup", CreatGroup)

	return router
}
