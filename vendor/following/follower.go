package following

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
	Userfollowerid string
}
type getData struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   string
	Follower []getProduct
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	Userfollowerid string `json:"followersid"`
}
type postData struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   string        `json:"userid"`
	Follower []postProduct `json:"follower"`
}
type Res struct {
	Data []postData `json:"Data"`
}

//GET *********************************************************************************
func GetAllData(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := res{}
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

func Getfollower(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := res{}

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
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.UserID
	fmt.Println("/n", email)
	//password:=c.FormValue("password")

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	fmt.Println(email)
	err = db.Find(bson.M{"userid": email}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)

	if results.Data == nil {
		defer session.Close()

		return c.JSON(http.StatusOK, 0)
	}
	defer session.Close()

	return c.JSON(http.StatusOK, &results)

}

//POST *********************************************************************************
func AddMentor(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
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
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := getData{}
	fmt.Println(res)
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)

	if result.UserID == "" {
		//fmt.Println("ni match hova add kr do")
		db.Insert(res)
	} else {
		//fmt.Println("match ho geya hai update kro")

		newdata := getData{}
		newdata = result

		a := res.Follower[0].Userfollowerid

		item1 := getProduct{Userfollowerid: a}

		newdata.AddItem(item1)
		db.Update(result, newdata)
	}

	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

//

func Unfollow(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
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

	result.removeFriend(res)

	result1 := getData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result1)

	db.Update(result1, result)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}
func (self *getData) removeFriend(item postData) {
	for i := range self.Follower {
		if self.Follower[i].Userfollowerid == item.Follower[0].Userfollowerid {
			self.Follower = append(self.Follower[:i], self.Follower[i+1:]...)
			fmt.Println(i)
			fmt.Println("match ho geya")
			break
		}
	}
}

func Addfollower(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
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

	a := res.Follower[0].Userfollowerid

	item1 := getProduct{Userfollowerid: a}

	newdata.AddItem(item1)
	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.Follower = append(box.Follower, item)
	return box.Follower
}
