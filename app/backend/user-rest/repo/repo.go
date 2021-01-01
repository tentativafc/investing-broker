package repo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDB struct {
	ID        string `gorm:"primarykey"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Email     string `gorm:"not null,unique,uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (UserDB) TableName() string {
	return "user"
}

type RecoverLoginDB struct {
	ID                string `gorm:"primarykey"`
	UserID            string
	User              UserDB `gorm:"foreignKey:UserID"`
	TemporaryPassword string `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (RecoverLoginDB) TableName() string {
	return "recover_password"
}

type UserRepository struct {
	db *gorm.DB
}

func (ur UserRepository) CreateUser(u UserDB) (UserDB, error) {
	err := ur.db.Create(&u).Error
	if err != nil {
		return UserDB{}, err
	}
	return u, nil
}

func (ur UserRepository) UpdateUser(u UserDB) (UserDB, error) {
	err := ur.db.Updates(&u).Error
	if err != nil {
		return UserDB{}, err
	}
	return u, nil
}

func (ur UserRepository) FindByEmail(email string) (UserDB, error) {
	var userDb UserDB
	ur.db.Where("email = ?", email).First(&userDb)
	err := ur.db.Where("email = ?", email).First(&userDb).Error
	return userDb, err
}

func (ur UserRepository) FindById(userId string) (UserDB, error) {
	var userDb UserDB
	err := ur.db.First(&userDb, &userId).Error
	return userDb, err
}

func (ur UserRepository) CreateRecoverPassword(u UserDB, id uuid.UUID, tempPassword string) (RecoverLoginDB, error) {
	recoverLoginDB := RecoverLoginDB{ID: id.String(), UserID: u.ID, User: u, TemporaryPassword: tempPassword, CreatedAt: time.Now()}
	err := ur.db.Create(&recoverLoginDB).Error
	return recoverLoginDB, err
}

func NewUserRepository() UserRepository {
	db := DbInit()
	ur := UserRepository{db: db}
	return ur
}

func DbInit() *gorm.DB {
	dbConfigs := "host=localhost user=postgres password=123456 dbname=postgres port=5432"
	var err error
	db, err := gorm.Open(postgres.Open(dbConfigs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserDB{})
	db.AutoMigrate(&RecoverLoginDB{})
	return db
}
