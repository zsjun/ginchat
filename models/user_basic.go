package models

import (
	"fmt"
	"ginchat/common"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `json:"name"`
	PassWord      string `json:"pass_word"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Identity      string `json:"identity"`
	ClientIP      string `json:"client_ip"`
	ClientPort    string `json:"client_port"`
	LoginTime     time.Time `gorm:"column:login_time;default:2023-03-02 15:08:25.427" json:"login_time"`
	HeartbeatTime time.Time `gorm:"column:heart_beat_time; default:2023-03-02 15:08:25.427" json:"heart_beat_time"`
	LoginOutTime  time.Time `gorm:"column:login_out_time; default:2023-03-02 15:08:25.427" json:"login_out_time"`
	IsLogout      bool   `json:"is_logout"`
	DeviceInfo    string `json:"device_info"`

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
func (u UserBasic) Create(user UserBasic) error {
	fmt.Println(common.DB)
	return common.DB.Model(u).Create(&user).Error
}
func (u *UserBasic) Delete() error {
	return common.DB.Model(u).Delete(u).Error
}
func (u UserBasic) Update(user UserBasic) error {
	return common.DB.Model(u).Updates(&user).Error
}




