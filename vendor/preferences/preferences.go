package preferences

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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

//GET *********************************************************************************
func GetAllPreferences(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetAllCategory(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	fmt.Println(results)
	buff, _ := json.Marshal(&results)
	fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetPrefencesbyCategory(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//results:=res{}
	//err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	result := getData{}
	category := c.FormValue("category")
	fmt.Println(category)
	//name :=c.FormValue("name")
	//fmt.Println(name)
	err = db.Find(bson.M{"category": category}).All(&result)
	if err != nil {
		//log.Fatal(err)
	}
	fmt.Println(result)
	buff, _ := json.Marshal(&result)
	fmt.Println(string(buff))

	json.Unmarshal(buff, &result)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)

}

//POST *********************************************************************************
func AddPreferences(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Add Preferences")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}
func AddCategory(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Add Category")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := getData{}
	err1 := db.Find(bson.M{"category": res.Category}).One(&result)
	if err1 == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "category already added")
	} else {
		db.Insert(res)
	}
	//fmt.Println(res)

	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}
func PutPreferences(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	//fmt.Println(name)
	//u:=new (postData)
	u := new(Res)
	if err = c.Bind(&u); err != nil {
	}
	res := Res{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	fmt.Println(res)
	//result := getData{}
	//var get_data getData
	//(db.Find(res.Data[0]).One(&get_data)
	//err = db.Find(bson.M{"billid": res.Data[0].Id,"totalitems":res.Data[0].TotalItems,"totalprice":res.Data[0].TotalPrice,"date":res.Data[0].Date}).One(&result)
	fmt.Println("new is", res.Data[0].Category, " ", res.Data[0].SubCategories, " ")
	//fmt.Println("result is",result.Id," ",result.TotalItems," ",result.TotalPrice)
	fmt.Println("new is", res.Data[1].Category, " ", res.Data[1].SubCategories, " ")
	db.Update(res.Data[0], res.Data[1])
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

//

func RemoveSubcategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}

	fmt.Println(res)
	result := getData{}

	err = db.Find(bson.M{"category": res.Category}).One(&result)

	result.removeFriend(res)

	result1 := getData{}

	err = db.Find(bson.M{"category": res.Category}).One(&result1)

	db.Update(result1, result)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

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

func Addsubcategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PREFERENCESCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}

	fmt.Println(res)
	result := getData{}

	err = db.Find(bson.M{"category": res.Category}).One(&result)
	newdata := getData{}
	newdata = result

	a := res.SubCategories[0].SubCategory

	item1 := getProduct{SubCategory: a}

	newdata.AddItem(item1)
	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.SubCategories = append(box.SubCategories, item)
	return box.SubCategories
}
