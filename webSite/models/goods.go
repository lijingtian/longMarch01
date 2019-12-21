package models

import(
	"longMarch01/manageSystem/common"
	"fmt"
)

var tableName = "goods"

type Goods struct{
	Id int32	`form:"id" json:"id"`
	Name string		`form:"name" json:"name"`
	Price int64		`form:"price" json:"price"`
	Stock int64		`form:"stock" json:"stock"`
	Status int32	`form:"status" json:"status"`
}

func NewGoods() *Goods {
	return new(Goods)
}

func (g *Goods) List() []Goods {
	var (
		datas []Goods
		where string
	)	
	if g.Status != 0{
		where = fmt.Sprintf(" AND status=%s", g.Status)
	}
	sql := fmt.Sprintf(`
		SELECT id, name, price, stock, status FROM %s WHERE 1=1 %s`, tableName, where)
	fmt.Println(sql)
	rows, err := common.GetDbConn().Query(sql)
	if err != nil{
		fmt.Println(err)
		return datas
	}
	defer rows.Close()
	for rows.Next(){
		var tmp Goods
		rows.Scan(&tmp.Id, &tmp.Name, &tmp.Price, &tmp.Stock, &tmp.Status)
		datas = append(datas, tmp)
	}
	return datas
}