package main

import (
	"fmt"
	_ "image/jpeg"
)

func main() {
	barcode := scanImage("./IMG_4779.jpg")
	bookInfo := getBookInfo(barcode)
	//fmt.Println(string(bookInfo))
	bookData := parseBookJson(bookInfo)
	authInfo := bookData.Authors
	returnedAuthors := []APIAuthor{}
	for _, author := range authInfo {
		//adding author information to a slice of new authors
		returnedAuthors = append(returnedAuthors, parseAuthInfo(getAuthorInfo(author)))
	}

	coverPicURL(barcode)
	fmt.Println(bookData.Title, bookData.ISBN10, bookData.NumberOfPages, bookData.Covers, bookData.Series)

	conn := connectToDB()

	//migrates the DB tables!
	migrateDB(conn)
	//seedGenres(conn)
	//
	//addABook(conn, bookData, returnedAuthors, barcode, 9)

	user := User{}
	//user.Name = "Elisa"
	//user.Password = "123456"
	user.Email = "elisa@elisa.com"
	//user.Create(conn)
	user.Retrieve(conn)
	//fmt.Println(user.Name, user.Password, user.Email)
	//user.Name = "Princess"
	//user.Update(conn)
	////newUserInfo := user.Retrieve(conn)
	////rowsAddedResponse(newUserInfo.RowsAffected)
	////printErrorHandler(newUserInfo.Error)
	////fmt.Println(user.Name, user.Password, user.Email)
	//user.Delete(conn)

	//pages := PageTracker{}
	book := Book{}
	book.Title = "Blood and Chocolate"
	book.Retrieve(conn)
	//pages.UserID = user.ID
	//pages.BookID = book.ID
	//pages.CurrentPage = 75
	//pages.Create(conn)
	//pages.Retrieve(conn)
	//fmt.Println(pages.CurrentPage)
	//pages.Update(conn)
	//tags := []string{
	//	"werewolf",
	//	"romance",
	//	"fantasy",
	//	"enemies to lovers",
	//}
	//for _, tag := range tags {
	//	t := Tag{}
	//	t.Name = tag
	//	t.Create(conn)
	//	bt := BookTag{}
	//	bt.BookID = book.ID
	//	bt.TagID = t.ID
	//	bt.Create(conn)
	//}
	newTags := []string{
		"paranormal",
		"YA",
		"werewolf",
		"werewolves",
	}
	for _, newTag := range newTags {
		t := Tag{}
		t.Name = newTag
		result := t.Retrieve(conn)
		fmt.Println(result.RowsAffected, newTag)
		if result.RowsAffected != 0 {
			continue
		} else {
			t.Create(conn)
			bt := BookTag{}
			bt.BookID = book.ID
			bt.TagID = t.ID
			bt.Create(conn)
		}
	}

	//shelf := Shelf{}
	//shelf.Name = "To Read"
	//shelf.UserID = user.ID
	////shelf.Create(conn)
	//shelf.Retrieve(conn)
	//shelf.Delete(conn)
	t := []Tag{}
	conn.Model(&book).Association("Tag").Find(&t)
	fmt.Print(t)
}
