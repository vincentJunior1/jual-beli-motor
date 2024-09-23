package repository

import (
	"time"

	"github.com/gin-gonic/gin"
)

// type = 0 (user)
// type = 1 (admin)
type User struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	Email     string     `gorm:"column:email" json:"email"`
	Password  string     `gorm:"column:password" json:"password"`
	Type      int        `gorm:"column:type" json:"type"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	Updated   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) GetTypeUser() string {
	userType := map[int]string{
		0: "user",
		1: "admin",
	}

	return userType[u.Type]
}

// TableName
func (User) TableName() string {
	return "users"
}

func GetUserByEmail(ctx *gin.Context, email string) (User, error) {
	var data User
	query := DB.Table("users")
	query = query.Where("email = ?", email)
	query.First(&data)

	return data, query.Error
}

func CreateUser(ctx *gin.Context, data User) error {
	query := DB.Model(&data)
	query.Save(&data)

	return query.Error
}
