package repo

import (
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

func (ccr ClientCredentialsRepository) CreateClientCredentials(cr ClientCredentials) (ClientCredentials, error) {
	err := ccr.db.Create(&cr).Error
	if err != nil {
		return ClientCredentials{}, err
	}
	return cr, nil
}

func (ccr ClientCredentialsRepository) FindByClientId(clientId string) (ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.db.Where("client_id = ?", clientId).First(&cr).Error
	return cr, err
}

func NewClientCredentialsRepository() ClientCredentialsRepository {
	var err error
	db, err := gorm.Open(postgres.Open(config.GetDbConfig()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&ClientCredentials{})
	ccr := ClientCredentialsRepository{db: db}
	return ccr
}
