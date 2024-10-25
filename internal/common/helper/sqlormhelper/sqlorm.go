package sqlormhelper

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type (
	SqlGormDatabase interface {
		Open() *gorm.DB
		Close() error
		Begin() *gorm.DB
		Commit(tx *gorm.DB) *gorm.DB
		Rollback(tx *gorm.DB) *gorm.DB
		GetConn() (*gorm.DB, error)
	}

	GormConnectionOptions struct {
		Host               string
		Port               int
		Username           string
		Password           string
		Database           string
		Schema             string
		GormConfig         *gorm.Config
		Conn               gorm.ConnPool
		UseTls             bool
		TlsMode            string
		TlsRootCACertFile  string
		TlsKeyFile         string
		TlsCertFile        string
		InsecureSkipVerify bool
		MaxOpenConns       int
		Timezone           string
	}

	BaseEntity struct {
		CreatedTime      time.Time      `gorm:"column:created_date;type:timestamptz(3);not null;autoCreateTime;comment:created date"`
		CreatedUser      sql.NullString `gorm:"column:created_by;type:varchar;comment:created by"`
		LastModifiedTime time.Time      `gorm:"column:updated_date;type:timestamptz(3);not null;autoUpdateTime;comment:updated date"`
		LastModifiedUser sql.NullString `gorm:"column:updated_by;type:varchar;comment:updated by"`
	}
)
