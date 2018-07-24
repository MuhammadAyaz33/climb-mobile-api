package mentor

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

type userRequest struct {
	UserEmail string `json:"parentemail"`
}

//	PARENT PROFILE DATA ****************

type GetKid struct {
	KidID string
}
type ParentgetData struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ParentEmail string
	Kids        []GetKid
}
type Parentres struct {
	Data []ParentgetData
}
type PostKid struct {
	KidID string `json:"kidid"`
}
type ParentpostData struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ParentEmail string        `json:"parentemail"`
	Kids        []PostKid     `json:"kids"`
}
type ParentRes struct {
	Data []ParentpostData `json:"Data"`
}

func BecomeMentorRequest(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(shared.BMentorpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.BMentorpostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//userid: = fmt.Sprintf("%x", string(res.UserID))
	//fmt.Println(res)
	results := shared.BMentorres{}
	err = db.Find(bson.M{"userid": res.UserID}).All(&results.Data)

	if results.Data == nil {
		if res.UserAge < 18 {
			res.ParentStatus = 0
		} else {
			res.ParentStatus = 1
		}
		res.AdminStatus = 0

		db.Insert(res)

	} else {

		defer session.Close()
		return c.JSON(http.StatusOK, "user already submit request")

	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}
func GetAllMentorAdminRequest(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	results := shared.BMentorres{}
	err = db.Find(bson.M{"adminstatus": 0}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetMentorParentsRequest(c echo.Context) error {

	//session, err := shared.ConnectMongo(shared.DBURL)
	//db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	//results := shared.Userinfores{}

	u := new(userRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := userRequest{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	kiddata := ParentgetData{}
	kiddata = GetParentKids(res.UserEmail)

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	// err = db.Find(bson.M{"_id": email}).All(&results.Data)

	// if err != nil {
	// 	//log.Fatal(err)
	// }
	//fmt.Println(results)
	buff, _ := json.Marshal(&kiddata)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &kiddata)
	//defer session.Close()
	return c.JSON(http.StatusOK, &kiddata)

}
func GetParentKids(parentemail string) ParentgetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PARENTCOLLECTION)
	results := ParentgetData{}

	//email :=c.FormValue("email")

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"parentemail": parentemail}).One(&results)

	if err != nil {
		//log.Fatal(err)
	}
	fmt.Println(results)
	defer session.Close()
	return results
}
