package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/manageSystem/models"
)

type Goods struct{}

func NewGoods() *Goods{
	return new(Goods)
}

func (g *Goods) GoodsList(c *gin.Context) {
	c.HTML(200, "Goods/List.html", gin.H{})
}