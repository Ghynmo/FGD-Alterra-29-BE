package roles

import "fgd-alterra-29/drivers/databases/users"

type Roles struct {
	ID    int           `gorm:"primaryKey"`
	Roles string        `gorm:"not null"`
	Users []users.Users `gorm:"foreignKey:role_id"`
}
