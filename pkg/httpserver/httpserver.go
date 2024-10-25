package httpserver

import (
	"example/config"
	"example/entity"
	"example/internal/common/helper/loghelper"
	"example/internal/common/helper/logwriterhelper"
	"example/internal/common/helper/sqlormhelper"
	"example/internal/diregistry"
	"example/internal/router"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func StartHTTPServer() {
	diregistry.BuildDIContainer()
	cfg := diregistry.GetDependency(diregistry.ConfigDIName).(*config.Config)

	err := loghelper.InitZap(cfg.App, cfg.Env)
	if err != nil {
		log.Panic("Can't init zap logger", err)
	}

	sqlGormLogDb := diregistry.GetDependency(diregistry.SqlGormLogHelperDIName).(sqlormhelper.SqlGormDatabase)
	if cfg.DatabaseLog.AutoMigration {
		dbLogs, err := sqlGormLogDb.GetConn()
		if err != nil {
			loghelper.Logger.Panic("Can't get gorm log connection", err)
		}
		err = autoMigrationLog(dbLogs)
		if err != nil {
			loghelper.Logger.Panic("Failed to auto migration log", err)
		}
	}
	sqlOrmWriter := logwriterhelper.NewSqlPostgresWriter(sqlGormLogDb)
	err = loghelper.InitZapWithSql(cfg.App, cfg.Env, sqlOrmWriter)
	if err != nil {
		loghelper.Logger.Panic("Can't init zap logger into db", err)
	}
	loghelper.Logger.Infof("Log helper created")
	loghelper.DBLogger.Infof("Log helper into db created")

	sqlGormPostgresDb := diregistry.GetDependency(diregistry.SqlGormPostgresHelperDIName).(sqlormhelper.SqlGormDatabase)
	if cfg.DatabasePostgres.AutoMigration {
		dbPostgres, err := sqlGormPostgresDb.GetConn()
		if err != nil {
			loghelper.Logger.Panic("Can't get gorm log connection", err)
		}
		err = autoMigrationPostgres(dbPostgres)
		if err != nil {
			loghelper.Logger.Panic("Failed to auto migration log", err)
		}
	}

	gin.SetMode(cfg.Mode)

	c := diregistry.GetDependency(diregistry.CronSchedulerDIName).(*cron.Cron)
	defer c.Stop()

	routersInit := router.InitRouter()

	port := fmt.Sprintf(":%d", cfg.HttpAddress)
	server := &http.Server{
		Addr:    port,
		Handler: routersInit,
	}

	loghelper.Logger.Infof("[INFO] Start http server listening %s", port)
	err = server.ListenAndServe()
	if err != nil {
		loghelper.Logger.Fatal("ListenAndServe: ", err)
	}
}

func autoMigrationLog(db *gorm.DB) error {
	err := db.AutoMigrate(
		&logwriterhelper.Log{},
		&logwriterhelper.AuditLog{},
	)
	if err != nil {
		return err
	}
	return nil
}

func autoMigrationPostgres(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.Product{},
		&entity.User{})
	if err != nil {
		return err
	}
	return nil
}
