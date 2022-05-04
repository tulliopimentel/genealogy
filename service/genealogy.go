package service

import (
	"example/desafio/database"
	"example/desafio/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get the entire genealogy of the family of the referenced person
// @Description This endpoint will go deep, will get all the persons of the family tree. That means will get all persons related to each other
// @Tags Show the entire Genealogy
// @Success 200 {array} models.Genealogy
// @Failure 404 {object} object
// @Router /v1/person/genealogy/{id} [get]
func ShowGenealogy(c *gin.Context){

	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(200, gin.H{
			"error": "ID has to be integer type",
		})

		return 
	}

	db := database.GetDataBase()

	var genealogy models.Genealogy

	err = db.Model(&genealogy).
	Preload("Person").
	Preload("PersonFamily").
	Preload("PersonFamily.Family").
	Preload("PersonFamily.Person").
	Preload("PersonFamily.PersonFamily").
	Preload("PersonFamily.Relationship").
	Preload("PersonFamily.Family.Person").
	Preload("PersonFamily.Family.PersonFamily").
	Preload("PersonFamily.Family.Relationship").
	Find(&genealogy, "person_id = ?", newid).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error": "Person not found it:" + err.Error(),
		})
		return
	}

	c.JSON(200, genealogy)
}

// @Summary Create a new reference (person ID) for a new Genealogy, which is relate to the table 'family'
// @Description This table has just two columns in database, that means you will just put the ID of the referenced person in the table family.
// @Tags Create a new Genealogy
// @Success 200 {array} models.Genealogy
// @Failure 404 {object} object
// @Router /v1/person/genealogy/{id} [post]
func CreateGenealogy (newGenealogy *gin.Context){

	db := database.GetDataBase()

	var genealogy models.Genealogy

	err := newGenealogy.ShouldBindJSON(&genealogy)

	if err != nil{
		newGenealogy.JSON(400, gin.H{
			"error": "Something went wrong:" + err.Error(),
		})
		return
	}

	db.Select("PersonID", "PersonFamilyID", "RelationshipID").Save(&genealogy)

	if err != nil {
		newGenealogy.JSON(400, gin.H{
			"error":"genealogy not created",
		})
		return
	}

	newGenealogy.JSON(200, genealogy)
}
