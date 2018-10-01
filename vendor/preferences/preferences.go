package preferences

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shared"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type getProduct struct {
	SubCategory string
}
type getData struct {
	ID            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Category      string
	SubCategories []getProduct
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	SubCategory string `json:"subcategory"`
}
type postData struct {
	ID            bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Category      string        `json:"category"`
	SubCategories []postProduct `json:"SubCategories"`
}
type Res struct {
	Data []postData `json:"Data"`
}

type getCategory struct {
	Category string
}
type Category struct {
	Data []getCategory `json:"Data"`
}

var response shared.Response

//GET *********************************************************************************
func GetAllPreferences(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	results := res{}
	err = db.Find(bson.M{}).All(&results.Data)
	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record Not found", 404, "")
		return c.JSON(http.StatusNotFound, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func GetAllCategory(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	results := Category{}
	err = db.Find(bson.M{}).All(&results.Data)
	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record Not found", 404, "")
		return c.JSON(http.StatusNotFound, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func GetPrefencesbyCategory(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//results:=res{}
	//err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	result := getData{}
	category := c.FormValue("category")
	//name :=c.FormValue("name")
	//fmt.Println(name)
	err = db.Find(bson.M{"category": category}).All(&result)
	if err != nil {
		//log.Fatal(err)
	}
	buff, _ := json.Marshal(&result)
	json.Unmarshal(buff, &result)
	response = shared.ReturnMessage(true, "Record found", 200, result)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

//POST *********************************************************************************
func AddPreferences(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	//fmt.Println(name)
	//u:=new (postData)

	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	if len(res.SubCategories) < 1 {
		response = shared.ReturnMessage(false, "Select Atleast One SubCategories", 401, "")
		return c.JSON(http.StatusNotFound, response)
	}
	db.Insert(res)
	response = shared.ReturnMessage(true, "Record Inserted", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func AddCategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	if len(res.Category) < 1 {
		response = shared.ReturnMessage(false, "Category not be empty", 401, "")
		return c.JSON(http.StatusNotFound, response)
	}
	result := getData{}
	err1 := db.Find(bson.M{"category": res.Category}).One(&result)
	if err1 == nil {
		response = shared.ReturnMessage(false, "Category already Exist", 409, "")
	} else {
		db.Insert(res)
		response = shared.ReturnMessage(true, "Category saded", 200, "")
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func PutPreferences(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	u := new(Res)
	if err = c.Bind(&u); err != nil {
	}
	res := Res{}
	res = *u
	//result := getData{}
	//var get_data getData
	//(db.Find(res.Data[0]).One(&get_data)
	//err = db.Find(bson.M{"billid": res.Data[0].Id,"totalitems":res.Data[0].TotalItems,"totalprice":res.Data[0].TotalPrice,"date":res.Data[0].Date}).One(&result)
	fmt.Println("new is", res.Data[0].Category, " ", res.Data[0].SubCategories, " ")
	//fmt.Println("result is",result.Id," ",result.TotalItems," ",result.TotalPrice)
	fmt.Println("new is", res.Data[1].Category, " ", res.Data[1].SubCategories, " ")
	err = db.Update(res.Data[0], res.Data[1])
	if err != nil {
		response = shared.ReturnMessage(false, "Preferences not updated", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Preferences updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

//

func RemoveSubcategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	result := getData{}

	err = db.Find(bson.M{"category": res.Category}).One(&result)
	result.removeFriend(res)
	result1 := getData{}
	err = db.Find(bson.M{"category": res.Category}).One(&result1)
	err = db.Update(result1, result)
	if err != nil {
		response = shared.ReturnMessage(false, "SubCategory Not Removed", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "SubCategory Removed", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func (self *getData) removeFriend(item postData) {
	for i := range self.SubCategories {
		if self.SubCategories[i].SubCategory == item.SubCategories[0].SubCategory {
			self.SubCategories = append(self.SubCategories[:i], self.SubCategories[i+1:]...)
			fmt.Println(i)
			fmt.Println("match ho geya")
			break
		}
	}
}

func RemoveCategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	err = db.Remove(bson.M{"category": res.Category})
	if err != nil {
		response = shared.ReturnMessage(false, "Category Not Removed", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Category Removed", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func Addsubcategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	result := getData{}

	err = db.Find(bson.M{"category": res.Category}).One(&result)
	newdata := getData{}
	newdata = result
	a := res.SubCategories[0].SubCategory
	item1 := getProduct{SubCategory: a}
	newdata.AddItem(item1)
	err = db.Update(result, newdata)
	if err != nil {
		response = shared.ReturnMessage(false, "SubCategory not added", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "SubCategory added", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.SubCategories = append(box.SubCategories, item)
	return box.SubCategories
}
