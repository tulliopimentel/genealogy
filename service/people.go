package service

import (
	"example/desafio/database"
	"example/desafio/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary ShowFamily() Returns the closest family members to the reference person
// @Description Returns the closest family members to the reference person, without delving into the family tree of the other members
// @Tags Show Family Tree
// @Success 200 {array} models.Genealogy
// @Failure 404 {object} object
// @Router /v1/person/{id} [get]
func ShowFamily(c *gin.Context){

	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(200, gin.H{
			"error": "ID has to be integer type",
		})

		return 
	}

	db := database.GetDataBase()

	var family []models.Family

	err = db.Preload("Person").
	Preload("Relationship").
	Preload("PersonFamily").
	Find(&family, "person_id = ?", newid).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error": "Person not found it:" + err.Error(),
		})
		return
	}

	c.JSON(200, family)
}

// @Summary CreatePerson() Create a new person
// @Description Insert a new person into the database, and you can create a new family tree with this person
// @Tags Create person
// @Success 200 {array} models.Person
// @Failure 404 {object} object
// @Router /v1/person [post]
func CreatePerson(c *gin.Context){

	db := database.GetDataBase()

	var person models.Person

	err := c.ShouldBindJSON(&person)

	if err != nil{
		c.JSON(400, gin.H{
			"error": "Something went wrong:" + err.Error(),
		})
		return
	}

	err = db.Create(&person).Error

	if err != nil{
		c.JSON(400, gin.H{
			"error": "Person not created:" + err.Error(),
		})
		return
	}

	c.JSON(200, person)
}

// @Summary Get all persons of the database
// @Description Get all persons of the database, even if is not related to each other
// @Tags Show all persons
// @Success 200 {array} models.Person
// @Failure 404 {object} object
// @Router /v1/person [get]
func ShowPeoples(c *gin.Context) {

	db := database.GetDataBase()

	var people []models.Person

	err := db.Find(&people).Error
	if err != nil{
		c.JSON(400, gin.H{
			"error": "Could not return any data:" + err.Error(),
		})
		return
	}

	c.JSON(200, people)
}