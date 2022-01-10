package categories

import (
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/drivers/databases/threads"
)

type Categories struct {
	ID       int    `gorm:"primaryKey"`
	Category string `gorm:"not null"`
	IconUrl  string
	Threads  []threads.Threads `gorm:"foreignKey:Category_id"`
	Q_Title  int               `gorm:"-:migration;->"`
}

func (category *Categories) ToDomain() categories.Domain {
	return categories.Domain{
		ID:       category.ID,
		Category: category.Category,
		IconUrl:  category.IconUrl,
		Threads:  threads.ToListDomain(category.Threads),
		Q_Title:  category.Q_Title,
	}
}

func ToListDomain(u []Categories) []categories.Domain {
	var Domains []categories.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
