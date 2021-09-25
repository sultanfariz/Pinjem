package authors

import (
	"Pinjem/businesses/authors"
	"time"

	"gorm.io/gorm"
)

type Authors struct {
	ID           uint   `gorm:"primary_key"`
	AuthorId     string `gorm:"type:varchar(100);not null"`
	Name         string `gorm:"type:varchar(13);not null"`
	PersonalName string `gorm:"type:varchar(100);not null"`
	Birthdate    string `gorm:"type:varchar(50);not null"`
	Bio          string `gorm:"type:text;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *Authors) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

func (u *Authors) ToDomain() authors.Domain {
	return authors.Domain{
		Id:           u.ID,
		AuthorId:     u.AuthorId,
		Name:         u.Name,
		PersonalName: u.PersonalName,
		Birthdate:    u.Birthdate,
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func FromDomain(domain authors.Domain) Authors {
	return Authors{
		ID:           domain.Id,
		AuthorId:     domain.AuthorId,
		Name:         domain.Name,
		PersonalName: domain.PersonalName,
		Birthdate:    domain.Birthdate,
		Bio:          domain.Bio,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func ToListDomain(data []Authors) []authors.Domain {
	var listDomain []authors.Domain
	for _, d := range data {
		listDomain = append(listDomain, d.ToDomain())
	}
	return listDomain
}
