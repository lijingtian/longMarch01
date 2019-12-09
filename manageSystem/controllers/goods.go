package controllers

import(
	"github.com/gin-gonic/gin"
	"longMarch01/manageSystem/models"
	"strconv"
	"fmt"
)

type Goods struct{}

func NewGoods() *Goods{
	return new(Goods)
}

func (g *Goods) List(c *gin.Context) {
	datas := models.NewGoods().List()
	c.HTML(200, "Goods/List.html", gin.H{"datas": datas})
}

func (g *Goods) Add(c *gin.Context){
	c.HTML(200, "Goods/Add.html", gin.H{})
}

func (g *Goods) AddJson(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	price := c.DefaultQuery("price", "")
	status := c.DefaultQuery("status", "")
	stock := c.DefaultQuery("stock", "")
	m := models.NewGoods()
	m.Name = name
	p, err := strconv.ParseFloat(price, 64)
	if err != nil{
		fmt.Println(err)
		return
	}
	m.Price = int64(p * 100)
	s, err := strconv.ParseInt(status, 10, 64)
	m.Status  = int32(s)
	if err != nil{
		fmt.Println(err)
		return
	}
	m.Stock, err = strconv.ParseInt(stock, 10, 64)
	if err != nil{
		fmt.Println(err)
		return
	}
	msg := m.Add()
	c.JSON(200, msg)
}