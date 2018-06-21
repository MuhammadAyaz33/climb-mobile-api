package message

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"shared"

	"github.com/labstack/echo"
	"github.com/rs/xid"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type getProduct struct {
	MessageID     string
	Message       string
	MessageStatus string
}
type getData struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   string
	Messages []getProduct
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	MessageID     string `json:"messageid"`
	Message       string `json:"message"`
	MessageStatus string `json:"messagestatus"`
}
type postData struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   string        `json:"userid"`
	Messages []postProduct `json:"messages"`
}
type Res struct {
	Data []postData `json:"Data"`
}

//GET *********************************************************************************
func GetAllMessages(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
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
func GetUserMessages(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
	results := res{}

	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}

	email := res.UserID

	err = db.Find(bson.M{"userid": email}).All(&results.Data)

	if err != nil {
		return c.JSON(http.StatusOK, "data not found")
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}

//POST *********************************************************************************
func AddUserMessages(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)

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
	mid := xid.New()
	if result.UserID == "" {
		//fmt.Println("no data added")

		res.Messages[0].MessageID = mid.String()
		db.Insert(res)
		defer session.Close()
		return c.JSON(http.StatusOK, "data added successfully")
	} else {
		//fmt.Println("data update")
		newdata := getData{}
		newdata = result

		msg := res.Messages[0].Message
		msgstatus := res.Messages[0].MessageStatus

		item1 := getProduct{MessageID: mid.String(), Message: msg, MessageStatus: msgstatus}

		newdata.AddItem(item1)

		db.Update(result, newdata)
		defer session.Close()
		return c.JSON(http.StatusOK, "data updated successfully")
	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

func RemoveUserMessages(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
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
	if err != nil {
		//fmt.Println(err)
		defer session.Close()
		return c.JSON(http.StatusOK, "data not found")
	}
	db.Update(result1, result)
	//fmt.Println(check)
	defer session.Close()
	return c.JSON(http.StatusOK, "successfull deleted")

}
func (self *getData) removeFriend(item postData) {
	for i := range self.Messages {
		if self.Messages[i].MessageID == item.Messages[0].MessageID {
			self.Messages = append(self.Messages[:i], self.Messages[i+1:]...)
			//fmt.Println(i)
			//fmt.Println("match ho geya")
			break
		}
	}
}

func MarkAsRead(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)

	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}

	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}

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

	result.markRead(res)
	result1 := getData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result1)
	db.Update(result1, result)
	//err = db.Update(result, newdata)
	fmt.Println(err)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.Messages = append(box.Messages, item)
	return box.Messages
}
func (newdata *getData) markRead(res postData) {
	for i := range newdata.Messages {
		if newdata.Messages[i].MessageID == res.Messages[0].MessageID {
			fmt.Println("matched")
			newdata.Messages[i].MessageStatus = "read"
			fmt.Println(newdata.Messages[i].MessageStatus)
			break
		}
	}
}
