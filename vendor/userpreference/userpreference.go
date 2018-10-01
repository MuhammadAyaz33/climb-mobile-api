package userpreference

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"shared"

	"github.com/ftloc/exception"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type getProduct struct {
	SubCategory string
}
type getData struct {
	ID              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID          string
	UserPreferences []getProduct
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	SubCategory string `json:"subcategory"`
}
type postData struct {
	ID              bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID          string        `json:"userid"`
	UserPreferences []postProduct `json:"userpreferences"`
}
type Res struct {
	Data []postData `json:"Data"`
}

//CONTRIBUTION DATA ***********************************************

type Getimageurl struct {
	Imagestatus string
}
type Getwebsiteurl struct {
	Websiteurl string
}
type Gettag struct {
	Tag string
}
type Getsubcategory struct {
	Subcategory string
}
type ContributionData struct {
	ID                 bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserEmail          string
	UserID             string
	UserFullName       string
	UserProfilePicture string
	Title              string
	MainCategory       string
	SubCategories      string
	ContributionText   string
	Videos             string
	AudioPath          string
	Images             []Getimageurl
	Website            []Getwebsiteurl
	Coverpage          string
	Tags               []Gettag
	ViewCount          int
	ContributionStatus int
	AdminStatus        int
}
type Contributionres struct {
	Data []ContributionData
}

type Postimageurl struct {
	Imagestatus string `json:"imagestatus"`
}
type Postwebsiteurl struct {
	Websiteurl string `json:"websiteurl"`
}
type Posttag struct {
	Tag string `json:"tag"`
}
type Postsubcategory struct {
	Subcategory string `json:"subcategory"`
}
type ContributionPostData struct {
	ID                 bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	UserEmail          string          `json:"useremail"`
	UserID             string          `json:"userid"`
	UserFullName       string          `json:"username"`
	UserProfilePicture string          `json:"userprofilepicture"`
	Title              string          `json:"title"`
	MainCategory       string          `json:"maincategory"`
	SubCategories      string          `json:"subcategories"`
	ContributionText   string          `json:"contributiontext"`
	Videos             string          `json:"videos"`
	AudioPath          string          `json:"audiopath"`
	Images             []Postimageurl  `json:"images"`
	Website            []Getwebsiteurl `json:"website"`
	Coverpage          string          `json:"coverpage"`
	Tags               []Posttag       `json:"tags"`
	ViewCount          int             `json:"view"`
	ContributionStatus int             `json:"contributionstatus"`
	AdminStatus        int             `json:"adminstatus"`
}
type ContributionRes struct {
	Data []ContributionPostData `json:"Data"`
}

var response shared.Response

//GET *********************************************************************************
func GetAllUserPreferences(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.USERPREFERENCECOLLECTION)
	results := res{}
	err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	if len(results.Data) < 1 {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "User Preferences Updated", 200, results.Data[0])
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func GetUserPrefences(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.USERPREFERENCECOLLECTION)
	results := res{}
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	//email :=c.FormValue("email")
	email := res.UserID
	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	err = db.Find(bson.M{"userid": email}).All(&results.Data)
	if err != nil {
		//log.Fatal(err)
	}
	if len(results.Data) < 1 {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "User Preferences", 200, results.Data[0])
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

//POST *********************************************************************************
func AddUserPreferences(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.USERPREFERENCECOLLECTION)

	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	if len(res.UserPreferences) < 1 {
		response = shared.ReturnMessage(false, "Preferences not be empty", 400, "")
		return c.JSON(http.StatusOK, response)
	}
	result := getData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if result.UserID == "" {
		db.Insert(res)
		response = shared.ReturnMessage(true, "User Preferences Added", 200, "")
	} else {
		newdata := getData{}
		newdata = result
		a := res.UserPreferences[0].SubCategory
		item1 := getProduct{SubCategory: a}
		newdata.AddItem(item1)
		db.Update(result, res)
		response = shared.ReturnMessage(true, "User Preferences Updated", 200, "")

	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

func RemoveUserPreferences(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.USERPREFERENCECOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	result := getData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	result.removeFriend(res)
	result1 := getData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result1)
	if err != nil {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	db.Update(result1, result)
	response = shared.ReturnMessage(true, "User preferences removed", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func (self *getData) removeFriend(item postData) {
	for i := range self.UserPreferences {
		if self.UserPreferences[i].SubCategory == item.UserPreferences[0].SubCategory {
			self.UserPreferences = append(self.UserPreferences[:i], self.UserPreferences[i+1:]...)
			//fmt.Println(i)
			//fmt.Println("match ho geya")
			break
		}
	}
}

func Addsubcategory(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERPREFERENCECOLLECTION)
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

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	newdata := getData{}
	newdata = result

	a := res.UserPreferences[0].SubCategory

	item1 := getProduct{SubCategory: a}

	newdata.AddItem(item1)
	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.UserPreferences = append(box.UserPreferences, item)
	return box.UserPreferences
}

func UserPreferenceSuggestionContribution(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	var subcat = []postProduct{}
	subcat = res.UserPreferences
	newdata := Contributionres{}
	if subcat == nil {
		response = shared.ReturnMessage(false, "SubCategory not be empty", 409, "")
	} else {
		for i := range subcat {
			exception.Try(func() {
				results := Contributionres{}
				err = db.Find(bson.M{"subcategories": subcat[i].SubCategory, "viewcount": bson.M{"$gt": 1}}).All(&results.Data)
				if results.Data != nil {
					for x := range results.Data {
						newdata.AddItem11(results.Data[x])
					}
					response = shared.ReturnMessage(true, "User preferences added", 200, newdata.Data)
				} else {
					response = shared.ReturnMessage(false, "Record not found", 404, "")
				}
			}).CatchAll(func(e interface{}) {
				fmt.Println("no data")
			}).Finally(func() {})
		}
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func (box *Contributionres) AddItem11(item ContributionData) []ContributionData {
	box.Data = append(box.Data, item)
	return box.Data
}
