package list

import (
	"database/sql"
	"fmt"
)

//List struct
type List struct {
	UserID    int    `json:"userID"`
	ListTitle string `json:"listTitle"`
	Create    string `json:"create"`
	Update    string `json:"update"`
	Delete    string `json:"delete"`
	Status    string `json:"status"`
}

// ListDetails struct for storing details of list to be deleted
type ListDetails struct {
	ProductID   int    `json:"productID"`
	ProductName string `json:"productName"`
	ListTitle   string `json:"listTitle"`
	Create      string `json:"create"`
	Update      string `json:"update"`
	Delete      string `json:"delete"`
	ModifiedBy  string `json:"modifiedBy"`
}

// CreateList func
func CreateList(list List, db *sql.DB) error {

	stmt, err := db.Prepare("insert into list (userID,listTitlename,creation_time,modified_time,deletion_time,status) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(list.UserID, list.ListTitle, list.Create, list.Update, list.Delete, list.Status)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}

// AddItemsList func
func AddItemsList(listDetails ListDetails, db *sql.DB) error {

	stmt, err := db.Prepare("insert into listDetails (productID,productName,listTitle,creation_time,modified_time,deletion_time,modifiedBy) values(?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(listDetails.ProductID, listDetails.ProductName, listDetails.ListTitle, listDetails.Create, listDetails.Update, listDetails.Delete, listDetails.ModifiedBy)

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}

// DeleteItemList func
func DeleteItemList(productID int, productName string, listTitle string, db *sql.DB) error {

	stmt, err := db.Prepare("Delete from listDetails where productID=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}

// DeleteList func
func DeleteList(listID int, listTitle string, db *sql.DB) error {

	stmt, err := db.Prepare("delete from list where ListID=? ;")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		fmt.Print(err.Error())
	}

	return err
}
