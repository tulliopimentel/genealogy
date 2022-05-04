package models


type Family struct {
	
	ID             		uint         	`json:"id"`
	PersonID       		int      	 	`json:"person"`
	Person				*Person		 	`gorm:"foreignKey:PersonID;references:ID;" json:"person"`
	PersonFamilyID 		int          	`json:"personFamily"`
	PersonFamily		*Person			`gorm:"foreignKey:PersonFamilyID;references:ID; json:"personFamily, omitempty"`
	RelationshipID 		int          	`json:"relationship"`
	Relationship		*Relationship 	`gorm:"foreignKey:RelationshipID;references:ID"`
	Family				[]*Family		`gorm:"foreignKey:PersonID; references:PersonFamilyID;" json:"family, omitempty"`
}