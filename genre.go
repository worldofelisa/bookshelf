package main

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `gorm:"unique;index;<-:create"`
}

//seedGenres will fill the table with the genres and create the genre IDs - loop to skip over once created
func seedGenres(conn *gorm.DB) {
	genres := []string{
		"Children's",
		"YA",
		"Classics",
		"Contemporary Fiction",
		"Drama",
		"Poetry",
		"Crime, Mystery, Thriller",
		"Horror",
		"Science Fiction",
		"Fantasy",
		"Arts and Design",
		"Comics, Graphic Novels and Manga",
		"Biography and Entertainment",
		"Food and Drink",
		"Health and Lifestyle",
		"Historical Fiction",
		"History and Politics",
		"Reference and Languages",
		"Travel",
		"Science, Study and Work",
		"Religion and Philosophy",
		"Home and Nature",
	}

	for _, g := range genres {
		genre := Genre{Name: g}
		genreResult := conn.Where(&genre).Find(&genre)
		if genreResult.RowsAffected == 0 {
			conn.Create(&genre)
		}
	}
}

func (g *Genre) Create(conn *gorm.DB) *gorm.DB {
	return conn.Create(&g)
}

//Retrieve to be used to call the genres for the front end
func (g *Genre) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&g).Find(&g)
}

func (g *Genre) Update(conn *gorm.DB) *gorm.DB {
	return conn.Save(&g)
}

func (g *Genre) Delete(conn *gorm.DB) *gorm.DB {
	return conn.Delete(&g)
}
