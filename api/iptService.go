package api

import (
	"fmt"
	"main/db"
	"main/model"

	"github.com/gin-gonic/gin"
)


func setupIptAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/ipt")
	{
      
		authenAPI.GET("/admitnow", getadmitNow)
		//authenAPI.GET("/login", test)
		//authenAPI.POST("/register", register)
	}
}




/*func getadmitNow(c *gin.Context) {
	var Ipt []model.Ipt
	tx:=db.GetDB5().Where("dchtype IS NULL").Find(&Ipt)
	if tx.Error !=nil{
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200,Ipt)
}*/


func getadmitNow(c *gin.Context) {
	var Ipt []model.Ipt
	se :="ipt.hn,ipt.an,concat(pt.pname,pt.fname,' ',pt.lname) AS fullname,ward.name AS wardname"
	join :="left join patient pt on pt.hn=ipt.hn left join ward on ward.ward=ipt.ward"
	where :="dchtype IS NULL"
	tx:=db.GetDB5().Model(&Ipt).Where(where).Select(se).Joins(join).Scan(&Ipt)
	if tx.Error !=nil{
		fmt.Println(tx.Error)
		return
	}
	c.JSON(200, gin.H{"result": Ipt})
}
