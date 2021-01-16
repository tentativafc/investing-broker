package repo

import (
	"errors"
	"fmt"
	"time"

	"github.com/tentativafc/investing-broker/app/backend/sts-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClientCredentials struct {
	ClientName   string `gorm:"primarykey"`
	ClientId     string `gorm:"not null"`
	ClientSecret string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (ClientCredentials) TableName() string {
	return "client_credentials"
}

type ClientCredentialsRepository struct {
	db *gorm.DB
}

func (ccr ClientCredentialsRepository) CreateClientCredentials(cr *ClientCredentials) (*ClientCredentials, error) {
	err := ccr.db.Create(cr).Error
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (ccr ClientCredentialsRepository) FindByClientId(clientId string) (*ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.db.Where("client_id = ?", clientId).First(&cr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cr, err
}

func (ccr ClientCredentialsRepository) FindByClientName(clientName string) (*ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.db.Where("client_name = ?", clientName).First(&cr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cr, err
}

func NewClientCredentialsRepository() ClientCredentialsRepository {
	fmt.Printf("Starting connection DB: %v ...", config.GetDbConfig())
	var err error
	db, err := gorm.Open(postgres.Open(config.GetDbConfig()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&ClientCredentials{})
	ccr := ClientCredentialsRepository{db: db}
	return ccr
}
