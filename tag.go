package main

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"index;<-:create"`
}

// BookTags adds book tags to books, by passing in the book id and creating the tags
func BookTags(conn *gorm.DB, tags []string, bookID uint) {
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
