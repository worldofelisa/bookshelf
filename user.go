package main

import (
	"gorm.io/gorm"
)

//addABook adds a book if it is not in the Book db table, if it does exist, it returns
func addABook(conn *gorm.DB, bookData BookInfo, returnedAuthors []APIAuthor, barcode string, tags []string) {
	book := Book{}
	bookTableData := conn.Where(&Book{Title: bookData.Title}).Find(&book)
	if bookTableData != nil {
		//TODO allow book to be selected by front end?
		return
	} else {
		storeBookInDB(conn, bookData, returnedAuthors, barcode, "")
	}

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
		bookTagResult := conn.Create(&BookTag{BookID: book.ID, TagID: bookTag.ID})
		rowsAddedResponse(bookTagResult.RowsAffected)
		printErrorHandler(bookTagResult.Error)
	}
}
