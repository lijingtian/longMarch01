package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/webSite/models"
)

type Goods struct{}

func NewGoods() *Goods{
	return new(Goods)
}

func (g *Goods) List(c *gin.Context) {
	datas := models.NewGoods().List()
	c.HTML(200, "Goods/List.html", gin.H{"datas": datas})
}