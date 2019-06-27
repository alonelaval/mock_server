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
const PARAM  ="Param"
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
	router.POST("/API_GP_HC",  func(c *gin.Context) {
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
	router.POST("/API_DJ_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewPointInspection().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_GLYZL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewAdmin().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_GCSZL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewEngineer().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_XXYZL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewMessenger().GetData()
		c.JSON(http.StatusOK,d)
	})

	router.POST("/API_HSZX_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewAccountingCenter().GetData()
		c.JSON(http.StatusOK,d)
	})

	router.POST("/API_HSDQ_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewAccountingRegion().GetData()
		c.JSON(http.StatusOK,d)
	})


	router.POST("/API_HSPQ_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewAccountingArea().GetData()
		c.JSON(http.StatusOK,d)
	})


	router.POST("/API_FWDZL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewServiceShop().GetData()
		c.JSON(http.StatusOK,d)
	})

	router.POST("/API_CPPL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewProductCategory().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_CPLX_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewProductType().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_CPXL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewProductSeries().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_CPZL_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewProduct().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_QYDA_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewArea().GetData()
		c.JSON(http.StatusOK,d)
	})

	router.POST("/API_QYWDDY_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewAreaShop().GetData()
		c.JSON(http.StatusOK,d)
	})
	router.POST("/API_YBJGYJBZ_CX",  func(c *gin.Context) {
		sign := c.PostForm(SIGN)
		params:= c.PostForm(PARAM)
		fmt.Println(sign)
		fmt.Println(params)

		d := handler.NewProductPrice().GetData()
		c.JSON(http.StatusOK,d)
	})





	//改派
	router.GET("/test",  func(c *gin.Context) {
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
