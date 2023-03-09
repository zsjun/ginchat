package models

import (
	"ginchat/common"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `gorm:"unique" json:"name"`
	PassWord      string    `json:"pass_word"`
	Phone         string    `valid:"matches(^1[3-9]{1}\\d{9})" json:"phone"`
	Email         string    `valid:"email" json:"email"`
	Identity      string    `json:"identity"`
	ClientIP      string    `json:"client_ip"`
	ClientPort    string    `json:"client_port"`
	LoginTime     time.Time `gorm:"column:login_time;default:2023-03-02 15:08:25.427" json:"login_time"`
	HeartbeatTime time.Time `gorm:"column:heart_beat_time; default:2023-03-02 15:08:25.427" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"column:login_out_time; default:2023-03-02 15:08:25.427" json:"login_out_time"`
	IsLogout      bool      `json:"is_logout"`
	DeviceInfo    string    `json:"device_info"`
}
type User struct {
	Name     string `gorm:"unique" json:"name"`
	PassWord string `json:"pass_word"`
}

func (table *UserBasic) CreateTableName() string {
	return "user_basic"
}

func (u UserBasic) List() ([]*UserBasic, error) {
	userList := make([]*UserBasic, 10)
	err := common.DB.Model(UserBasic{}).Where("id <> 0").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func FindUserByName(name string) (*UserBasic, error) {
	user := UserBasic{}

	err := common.DB.Model(UserBasic{}).Where("name = ?", name).First(&user).Error

	if err != nil {
		return nil, err
	}
	// fmt.Println(12, common.DB.Model(UserBasic{}).Where("name = ?", name).First(&user))
	return &user, err
}
func FindUserByPhone(phone string) error {
	user := UserBasic{}
	return common.DB.Model(UserBasic{}).Where("phone = ?", phone).First(&user).Error
}

func FindUserByEmail(email string) error {
	user := UserBasic{}
	return common.DB.Model(UserBasic{}).Where("email = ?", email).First(&user).Error
}
func (u UserBasic) Create(user UserBasic) error {
	return common.DB.Model(u).Create(&user).Error
}
func (u *UserBasic) Delete() error {
	return common.DB.Model(u).Delete(u).Error
}
func (u UserBasic) Update(user UserBasic) error {
	return common.DB.Model(u).Updates(&user).Error
}
