package repo

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tentativafc/investing-broker/app/backend/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Email     string `gorm:"not null,unique,uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}

type RecoverLogin struct {
	ID                string `gorm:"primarykey"`
	UserID            string
	User              User   `gorm:"foreignKey:UserID"`
	TemporaryPassword string `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (RecoverLogin) TableName() string {
	return "recover_password"
}

type ClientCredentials struct {
	ClientName   string `gorm:"primarykey"`
	ClientId     string `gorm:"not null"`
	ClientSecret string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository struct {
	db    *gorm.DB
	dbSts *gorm.DB
}

func (ur UserRepository) CreateUser(u User) (User, error) {
	err := ur.db.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (ur UserRepository) UpdateUser(u User) (User, error) {
	err := ur.db.Updates(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (ur UserRepository) FindByEmail(email string) (User, error) {
	var user User
	ur.db.Where("email = ?", email).First(&user)
	err := ur.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (ur UserRepository) FindById(userId string) (*User, error) {
	var user User
	err := ur.db.First(&user, &userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (ur UserRepository) CreateRecoverPassword(u User, id uuid.UUID, tempPassword string) (RecoverLogin, error) {
	recoverLogin := RecoverLogin{ID: id.String(), UserID: u.ID, User: u, TemporaryPassword: tempPassword, CreatedAt: time.Now()}
	err := ur.db.Create(&recoverLogin).Error
	return recoverLogin, err
}

func (ccr UserRepository) FindClientCredentialsByClientName(clientName string) (*ClientCredentials, error) {
	var cr ClientCredentials
	err := ccr.dbSts.Where("client_name = ?", clientName).First(&cr).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cr, err
}

func NewUserRepository() UserRepository {
	var err error
	db, err := gorm.Open(postgres.Open(config.GetDbConfig()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&RecoverLogin{})

	dbSts, err := gorm.Open(postgres.Open(config.GetDbConfigSts()), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database of Sts")
	}

	ccr := UserRepository{db: db, dbSts: dbSts}
	return ccr
}
