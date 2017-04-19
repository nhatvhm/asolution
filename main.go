package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql" // used for all session stores
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/nhatvhm/asolution/models"
	_ "github.com/nhatvhm/asolution/routers"
	_ "github.com/lib/pq"
	"log"
	"time"
	"fmt"
    "os"
)

var (
   dbConnectString string = beego.AppConfig.String("dbConnectString")
   dbDriverName string = beego.AppConfig.String("dbDriverName")
   dbMaxIdle, dbMaxIdleErr = beego.AppConfig.Int("dbMaxIdle")
   dbMaxConn, dbMaxConnErr = beego.AppConfig.Int("dbMaxConn")
 )

func init() {
	// Development Settings, adjust for production
	// mysql / sqlite3 / postgres driver registered by default already
	orm.RegisterDriver(dbDriverName, orm.DRPostgres)

	if dbMaxIdleErr != nil {
      panic("Main Init: - Unable to start the server can't parse dbMaxIdleErr from configuration file must be int.")
    }
   
    if dbMaxConnErr != nil {
      panic("Main Init: - Unable to start the server can't parse dbMaxConnErr from configuration file must be int.")
    }

	//                    db alias  drivername
	orm.RegisterDataBase("default", dbDriverName, dbConnectString, dbMaxIdle, dbMaxConn)
	orm.DefaultTimeLoc = time.UTC

}

func main() {
	//beego.TemplateLeft = "<<<" // set to make internal template compatible with most front ends i.e. Angular, Polymer, etc
	//beego.TemplateRight = ">>>"

	name := "default"
	// Whether to drop table and re-create.
	force := true
	// Print log.
	verbose := true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		log.Println(err)
	}
	// Reconfig port
	port := os.Getenv("PORT")
	fmt.Println("listening PORT: " + port)
    if err != nil {
        beego.Run(":" + port)
    }
}
