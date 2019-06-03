package main

import (
	"fmt"
	"handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SIGN ="SIGN"
const APPKEY="APPKEY"
const PGDH,YYSJ= "PGDH","YYSJ"
const GCSBH = "GCSBH"
func main() {


	router := gin.Default()
	//订单
	router.POST("/API_PG_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		appkey := c.PostForm(APPKEY)
		fmt.Println(sign)
		fmt.Println(appkey)
		d := handler.NewOrder().GetData()
		c.JSON(http.StatusOK,d)
	})
	//预约
	router.POST("/API_YY_HC",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		appkey:= c.PostForm(APPKEY)
		pgdh := c.PostForm(PGDH)
		yysj:= c.PostForm(YYSJ)
		fmt.Println(sign)
		fmt.Println(appkey)
		fmt.Println(pgdh)
		fmt.Println(yysj)

		d := handler.NewAppointment().GetData()
		c.JSON(http.StatusOK,d)
	})
	//改派
	router.POST("/API_YY_HC",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		appkey:= c.PostForm(APPKEY)
		pgdh := c.PostForm(PGDH)
		yysj:= c.PostForm(YYSJ)
		gcsbh:= c.PostForm(GCSBH)
		fmt.Println(sign)
		fmt.Println(appkey)
		fmt.Println(pgdh)
		fmt.Println(yysj)
		fmt.Println(gcsbh)

		d := handler.NewAppointment().GetData()
		c.JSON(http.StatusOK,d)
	})



	router.Run(":8888") //for a hard coded port
}
