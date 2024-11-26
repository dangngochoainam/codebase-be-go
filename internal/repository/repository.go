package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

func isTransactionActive(db *gorm.DB) bool {
	return db.Statement != nil && db.Statement.ConnPool != db.Config.ConnPool
}

func runInTx(db *gorm.DB, fn func(db *gorm.DB) error) error {
	if isTransactionActive(db) {
		log.Println("In Transaction")
		return fn(db)
	}
	log.Println("Start Transaction")
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	err := fn(tx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return errors.Join(err, rollbackErr.Error)
		}
	}

	return tx.Commit().Error
}
