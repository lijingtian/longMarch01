package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/webSite/models"
	"strconv"
)

type Goods struct{}

func NewGoods() *Goods{
	return new(Goods)
}

func (g *Goods) List(c *gin.Context) {
	datas := models.NewGoods().List()
	c.HTML(200, "Goods/List.html", gin.H{"datas": datas})
}

func (g *Goods) Buy(c *gin.Context) {
	num := c.DefaultQuery("num", "0")
	numInt, _ := strconv.ParseInt(num, 10, 64)
	m := models.NewGoods()
	tmpId := c.DefaultQuery("goodId", "0")
	t, _ := strconv.ParseInt(tmpId, 10, 64)
	m.Id = int32(t)
	msg := m.Buy(numInt)
	c.JSON(200, msg)
}