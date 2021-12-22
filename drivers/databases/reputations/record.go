package reputations

import "fgd-alterra-29/drivers/databases/users"

type Reputations struct {
	ID         int           `gorm:"primaryKey"`
	Reputation string        `gorm:"not null"`
	Users      []users.Users `gorm:"foreignKey:reputation_id"`
}
