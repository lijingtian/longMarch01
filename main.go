package main

import(
	"fmt"
	"os"
	"github.com/urfave/cli"
	
	"longMarch01/services"
	
)

func main(){
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name: "test",
			Usage: "this is a test",
			Action: func(c *cli.Context) error {
				fmt.Println("this is urfave cli test")
				return nil
			},
		},
		{
			Name: "webSite",
			Usage: "this is web site",
			Action: func(c *cli.Context) error {
				fmt.Println("web site start")
				services.StartWebSite()
				return nil
			},
		},
		{
			Name: "manageSystem",
			Usage: "this is manage system",
			Action: func(c *cli.Context) error{
				fmt.Println("manage system start")
				services.StartManageSystem()
				return nil
			},
		},
	}
	
	app.Run(os.Args)
}

/*
func StartWebSite(){
	r := gin.Default()
	r.GET("/index", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "this is web site index"})
	})
	r.Run(":8081")
}


func StartManagerSystem(){
	r := gin.Default()
	r.GET("/index", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"this is manager system index"})
	})
	r.Run(":8082")
}
*/