package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kalifun/ginblog/pkg/setting"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.GET("/test", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"test",
		})
	})

	s := &http.Server{
		Addr:              fmt.Sprintf(":%d",setting.HTTPPort),
		Handler:           router,
		ReadTimeout:       setting.ReadTimeout,
		WriteTimeout:      setting.WriteTimeout,
		MaxHeaderBytes:    1 << 20,
	}

	s.ListenAndServe()
}
