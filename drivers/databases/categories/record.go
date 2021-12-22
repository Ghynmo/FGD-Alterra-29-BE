package categories

import "fgd-alterra-29/drivers/databases/threads"

type Categories struct {
	ID       int               `gorm:"primaryKey"`
	Category string            `gorm:"not null"`
	Threads  []threads.Threads `gorm:"foreignKey:Category_id"`
	Title    string            `gorm:"-:migration;->"`
}
