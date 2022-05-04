package models


type Family struct {
	
	ID             		uint         	`json:"id"`
	PersonID       		int      	 	
	Person				*Person		 	`gorm:"foreignKey:PersonID;references:ID;"`
	PersonFamilyID 		int          	`json:"-"`
	PersonFamily		*Person			`gorm:"foreignKey:PersonFamilyID;references:ID;"`
	RelationshipID 		int          	`json:"-"`
	Relationship		*Relationship 	`gorm:"foreignKey:RelationshipID;references:ID"`
	Family				[]*Family		`gorm:"foreignKey:PersonID;references:PersonFamilyID;"`
}