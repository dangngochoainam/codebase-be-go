package logwriterhelper

import (
	"encoding/json"
	"example/internal/common/helper/sqlormhelper"
)

type (
	sqlPostgresWriter struct {
		gormDb sqlormhelper.SqlGormDatabase
	}

	sqlPostgresAuditLogWriter struct {
		gormDb sqlormhelper.SqlGormDatabase
	}
)

type (
	Log struct {
		sqlormhelper.BaseEntity
		Id      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Level   string `gorm:"column:level" json:"level"`
		TraceId string `gorm:"column:trace_id;index:idx__logs__trace_id" json:"traceId"`
		Message string `gorm:"column:message" json:"message"`
	}

	AuditLog struct {
		sqlormhelper.BaseEntity
		Id      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
		Level   string `gorm:"column:level" json:"level"`
		TraceId string `gorm:"column:trace_id;index:idx__logs__trace_id" json:"traceId"`
		Message string `gorm:"column:message" json:"message"`
	}
)

func (Log) TableName() string {
	return "log"
}

func (AuditLog) TableName() string {
	return "audit_log"
}

func NewSqlPostgresWriter(gormDb sqlormhelper.SqlGormDatabase) Writer {
	return &sqlPostgresWriter{
		gormDb: gormDb,
	}
}

func (h *sqlPostgresWriter) Write(p []byte) (n int, err error) {
	db := h.gormDb.Open()
	var jsonData Log
	jsonErr := json.Unmarshal(p, &jsonData)
	if jsonErr != nil {
		return 0, jsonErr
	}
	log := &Log{
		TraceId: jsonData.TraceId,
		Level:   jsonData.Level,
		Message: string(p),
	}
	result := db.Create(log)
	return int(result.RowsAffected), result.Error
}

func NewSqlPostgresAuditLogWriter(gormDb sqlormhelper.SqlGormDatabase) Writer {
	return &sqlPostgresAuditLogWriter{
		gormDb: gormDb,
	}
}

func (h *sqlPostgresAuditLogWriter) Write(p []byte) (n int, err error) {
	db := h.gormDb.Open()
	auditLog := &AuditLog{
		Message: string(p),
	}
	result := db.Create(auditLog)
	return int(result.RowsAffected), result.Error
}
