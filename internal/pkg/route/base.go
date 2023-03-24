package route

import (
	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine) {
	SetUpLineMessage(r)
}
