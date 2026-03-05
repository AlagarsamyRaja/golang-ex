package models

type Datas struct {
	Host     string
	User     string
	Dbname   string
	Password string
	Port     string
	Mode     string
}

type Users struct {
	RollNumber      uint            `gorm:"primaryKey" json:"rollnumber"`
	Name            string          `json:"name"`
	DOB             string          `json:"dob"`
	Address         string          `json:"address"`
	UserRoleMapping UserRoleMapping `gorm:"foreignKey:RollNumber;references:RollNumber;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"userrolemapping"`
}

type UserRoles struct {
	RoleId          uint            `gorm:"primaryKey" json:"roleid"`
	Role            string          `json:"role"`
	Permission      string          `json:"permission"`
	UserRoleMapping UserRoleMapping `gorm:"foreignKey:RoleId;references:RoleId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"userrolemapping"`
}

type UserRoleMapping struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	RollNumber uint `json:"rollnumber"`
	RoleId     uint `json:"roleid"`
}

type Error struct {
	Message    string `json:"message"`
	Errors     error  `json:"errors"`
	Statuscode int    `json:"statuscode"`
}
