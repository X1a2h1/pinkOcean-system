package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Context *gin.Context
	Errors  error
}

func (e *Api) addError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		fmt.Println("12")
	}
}
