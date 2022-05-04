package models

type Genealogy struct{
	ID						uint			`json:"id"`
	PersonID				int				
	Person					*Person			`gorm:"foreignKey:PersonID;references:ID;"`		
	PersonFamily			[]Family		`gorm:"foreignKey:PersonID;"`
}