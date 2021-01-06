package repo

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ClientCredentials struct {
	ID           string `gorm:"primarykey"`
	ClientName   string `gorm:"not null"`
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

func (ccr ClientCredentialsRepository) UpdateClientCredentials(cr ClientCredentials) (ClientCredentials, error) {
	err := ccr.db.Updates(&cr).Error
	if err != nil {
		return ClientCredentials{}, err
	}
	return cr, nil
}

func (cr ClientCredentialsRepository) FindByClientId(clientId string) (ClientCredentials, error) {
	var cr ClientCredentials
	err := cr.db.Where("client_id = ?", ClientId).First(&cr).Error
	return cr, err
}

func NewClientCredentialsRepository() ClientCredentialsRepository {
	db := DbInit()
	ur := ClientCredentialsRepository{db: db}
	return ur
}

func DbInit() *gorm.DB {
	dbConfigs := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
	var err error
	db, err := gorm.Open(postgres.Open(dbConfigs), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&ClientCredentials{})
	return db
}
