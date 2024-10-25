package sqlormhelper

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormPostgresqlDB struct {
	db *gorm.DB
}

func NewGormPostgresqlDB(options *GormConnectionOptions) SqlGormDatabase {
	db, err := initGormPostgresqlDB(options)
	if err != nil {
		zap.S().Panic("Failed to init postgresql", err)
	}
	return &gormPostgresqlDB{
		db: db,
	}
}

func (h gormPostgresqlDB) Open() *gorm.DB {
	return h.db
}

func (h *gormPostgresqlDB) Close() error {
	return nil
}

func (h gormPostgresqlDB) Begin() *gorm.DB {
	return h.db.Begin()
}

func (h gormPostgresqlDB) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (h gormPostgresqlDB) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (h gormPostgresqlDB) GetConn() (*gorm.DB, error) {
	return h.db, nil
}

func initGormPostgresqlDB(options *GormConnectionOptions) (*gorm.DB, error) {
	tz0 := "Asia/Ho_Chi_Minh"
	if options.Timezone != "" {
		tz0 = options.Timezone
	}
	_, err := time.LoadLocation(tz0)
	if err != nil {
		return nil, err
	}

	dsnString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable search_path=%s TimeZone=%s",
		options.Host,
		options.Port,
		options.Username,
		options.Password,
		options.Database,
		options.Schema,
		tz0,
	)

	if options.UseTls {
		// Postgres sslmode: disable, allow, prefer, require, verify-ca and verify-full
		dsnString = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s search_path=%s sslmode=%s sslrootcert=%s sslkey=%s sslcert=%s TimeZone=%s",
			options.Host,
			options.Port,
			options.Username,
			options.Password,
			options.Database,
			options.Schema,
			options.TlsMode,
			options.TlsRootCACertFile,
			options.TlsKeyFile,
			options.TlsCertFile,
			tz0,
		)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsnString,
		PreferSimpleProtocol: true, // disables implicit prepared statement
		Conn:                 options.Conn,
	}), options.GormConfig)
	if err != nil {
		return nil, err
	}

	// Connection Pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// sqlDB.SetMaxIdleConns(10) // currently 2, maybe change in future
	// sqlDB.SetMaxOpenConns(100) // 0 unlimited
	// sqlDB.SetConnMaxIdleTime(time.Hour * 1)
	// sqlDB.SetConnMaxLifetime(time.Hour * 1)
	if options.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(options.MaxOpenConns)
	}

	zap.S().Info("Gorm Postgresql: Successfully connected!")
	return db, nil
}
