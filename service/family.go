package service

import (
	"example/desafio/database"
	"example/desafio/models"
	"errors"
	"github.com/gin-gonic/gin"
)

// @Summary Creates the relationship between two people to create a family and family tree
// @Description In this endpoint you can create a family, that is, put the ID of the reference person, another person and the relationship between the two in the database. All by ID
// @Tags Create a new family
// @Success 200 {object} models.Family
// @Failure 404 {object} object
// @Router /v1/person/family [post]
func CreateFamily(family *gin.Context){ 

	db := database.GetDataBase()

	var familyDb models.Family
	var genealogyDb models.Genealogy

	err := family.ShouldBindJSON(&familyDb)

	if err != nil{
		family.JSON(200, gin.H{
			"error": "Algo deu errado",
		})
		return 
	}

	err = db.Save(&familyDb).Error

	if err != nil{
		family.JSON(400, gin.H{
			"error": "Person not created:" + err.Error(),
		})
		return
	}

	if err != nil{
		family.JSON(200, gin.H{
			"error": "Algo deu errado",
		})
		return 
	}

	if db.Find(&genealogyDb, "person_id = ?", familyDb.PersonID) == nil {
		saveGenealogy := func(id int){

			genealogy := models.Genealogy{
				PersonID: id,
			}

			err := db.Select("PersonID").Save(&genealogy)

			if err != nil{
				errors.Unwrap(err.Error)
			}
		}

		saveGenealogy(familyDb.PersonID)
	}
	
	family.JSON(200, &familyDb)
}