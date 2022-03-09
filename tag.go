package main

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"index;<-:create"`
}

type BookTag struct {
	gorm.Model
	TagID  uint
	BookID uint
}

// Create adds book tags to Books, by passing in the book id and creating the tags
func (t *Tag) Create(conn *gorm.DB, tags []string, bookID uint) {
	//creates an empty array
	bookTags := []Tag{}
	//for every tag within the array tags, add to the empty tag array in bookTags
	for _, tag := range tags {
		bookTags = append(bookTags, Tag{Name: tag})
	}

	for _, bookTag := range bookTags {
		tagsResult := conn.Create(&bookTag)
		rowsAddedResponse(tagsResult.RowsAffected)
		printErrorHandler(tagsResult.Error)
	}

	for _, bookTag := range bookTags {
		bookTagResult := conn.Create(&BookTag{BookID: bookID, TagID: bookTag.ID})
		rowsAddedResponse(bookTagResult.RowsAffected)
		printErrorHandler(bookTagResult.Error)
	}
}

// Retrieve see the tags on a book
func (t *Tag) Retrieve(conn gorm.DB) *gorm.DB {
	return conn.Where(&t).Find(&t)
}

// Update the tags on a book
func (t *Tag) Update(conn gorm.DB, newTags []string) {
	//search if tag exists, if it does, return loop, else add
	tag := t.Name
	for _, newtag := range newTags {
		if newtag == tag {
			return
		} else {
			newTags = append(newTags, Tag{Name: newtag})
		}
		}
	}
}
