package main

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `gorm:"index;<-:create"`
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
		genreSeed := genre.Retrieve(conn)
		if genreSeed == nil {
			conn.Create(&genre)
		}
	}
}

//Retrieve to be used to call the genres for the front end
func (g *Genre) Retrieve(conn *gorm.DB) *gorm.DB {
	return conn.Where(&g).Find(&g)
}
