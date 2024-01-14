package db

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"party_room_server/build_conf"
	"time"
)

var DB *gorm.DB

func Connect() error {
	var err error
	dsn := build_conf.DataBaseDSN
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			Colorful:                  false,
		},
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	err = _initTables()
	if err != nil {
		return err
	}
	return err
}

func _initTables() error {
	return DB.AutoMigrate(&Room{}, &RoomUser{})
}

type BaseTable struct {
	ID        uuid.UUID      `gorm:"type:uuid" json:"ID"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

func (base *BaseTable) BeforeCreate(tx *gorm.DB) error {
	u := uuid.New()
	tx.Statement.SetColumn("ID", u)
	return nil
}
