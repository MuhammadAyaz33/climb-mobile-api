package following

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification"
	"os"
	"shared"

	"github.com/labstack/echo"
	"github.com/rs/xid"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type getProduct struct {
	Followid       string
	Userfollowerid string
	ParentStatus   int
	MessageStatus  int
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
	Followid       string `json:"followid"`
	Userfollowerid string `json:"followersid"`
	ParentStatus   int    `json:"parentstatus"`
	MessageStatus  int    `json:"messagestatus"`
}
type postData struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID   string        `json:"userid"`
	Follower []postProduct `json:"follower"`
}
type Res struct {
	Data []postData `json:"Data"`
}

type GetUserData struct {
	UserID   string        `json:"userid"`
	Follower []postProduct `json:"follower"`
	Age      int           `json:"userage"`
	FollowID string        `json:"followid"`
}

var response shared.Response

type GetFollowing struct {
	UserID string `json:"userid"`
}

//GET *********************************************************************************
func GetAllData(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := res{}
	err = db.Find(bson.M{}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, &results)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func Getfollower(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)

	results := res{}
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	email := res.UserID

	err = db.Find(bson.M{"userid": email}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	if results.Data == nil || len(results.Data[0].Follower) <= 0 {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data[0])
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

func GetUserfollower(c echo.Context) error {

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
	fmt.Println("get following")
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.UserID
	//fmt.Println("/n", email)
	//password:=c.FormValue("password")

	// err = db.Find(bson.M{"$or": []bson.M{bson.M{"cms": cms}, bson.M{"name": name}}}).All(&results.Data)
	// fmt.Println(email)
	err = db.Find(bson.M{"follower.userfollowerid": email}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	if results.Data == nil {
		defer session.Close()

		return c.JSON(http.StatusOK, 0)
	}
	//fmt.Println(results)
	var following []GetFollowing
	for x := range results.Data {
		var a GetFollowing
		a.UserID = results.Data[x].UserID
		following = append(following, a)
	}
	buff, _ := json.Marshal(&following)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &following)

	defer session.Close()

	return c.JSON(http.StatusOK, &following)

}

func GetfollowerByEmail(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.UsergetData{}
	follower := res{}

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
	res = *u

	email := res.Email

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"email": email}).One(&results)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	hexid := fmt.Sprintf("%x", string(results.ID))
	follower = GetfollowerById(hexid)
	buff, _ := json.Marshal(&follower)
	json.Unmarshal(buff, &follower)
	response = shared.ReturnMessage(true, "Record Found", 200, follower)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func GetfollowerById(userid string) res {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := res{}

	err = db.Find(bson.M{"userid": userid}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)

	defer session.Close()

	return results

}

//POST *********************************************************************************
func AddMentor(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
	res = *u
	result := getData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	userfollower := postData{}
	fid := xid.New()
	if result.UserID == "" {
		userfollower.UserID = res.UserID
		// userfollower.Follower[0].Userfollowerid = res.Userfollowerid
		// userfollower.Follower[0].ParentStatus = 0
		// userfollower.Follower[0].MessageStatus = 0
		var agestatus int
		var msgstatus int
		age := res.Age

		if age < 18 {
			agestatus = 0
			msgstatus = 0
		} else {
			agestatus = 1
			msgstatus = 1
			notification.AddMentorFollwerHistory(res.Follower[0].Userfollowerid, res.UserID)
		}
		item1 := postProduct{Followid: fid.String(), Userfollowerid: res.Follower[0].Userfollowerid, ParentStatus: agestatus, MessageStatus: msgstatus}
		userfollower.AddItem11(item1)
		err = db.Insert(userfollower)
		if err != nil {
			response = shared.ReturnMessage(false, "Record not added", 400, "")
			return c.JSON(http.StatusOK, response)
		}
		response = shared.ReturnMessage(true, "Record added", 200, "")
	} else {
		for x := range result.Follower {
			if result.Follower[x].Userfollowerid == res.Follower[0].Userfollowerid {
				defer session.Close()
				response = shared.ReturnMessage(false, "User already follow", 409, "")
				return c.JSON(http.StatusOK, response)
			}
		}
		newdata := getData{}
		newdata = result

		a := res.Follower[0].Userfollowerid
		var agestatus int
		var msgstatus int
		age := res.Age

		if age < 18 {
			agestatus = 0
			msgstatus = 0
		} else {
			agestatus = 1
			msgstatus = 1
			notification.AddMentorFollwerHistory(res.Follower[0].Userfollowerid, res.UserID)
		}
		item1 := getProduct{Followid: fid.String(), Userfollowerid: a, ParentStatus: agestatus, MessageStatus: msgstatus}
		newdata.AddItem(item1)
		err = db.Update(result, newdata)
		if err != nil {
			response = shared.ReturnMessage(false, "Record not updated", 400, "")
			return c.JSON(http.StatusOK, response)
		}
		response = shared.ReturnMessage(true, "Record updated", 200, "")
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

//
func (box *postData) AddItem11(item postProduct) []postProduct {
	box.Follower = append(box.Follower, item)
	return box.Follower
}

func Unfollow(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)

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
	err = db.Update(result1, result)
	if err != nil {
		response = shared.ReturnMessage(false, "Can't unfollow, error", 400, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Unfollow successful", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func (self *getData) removeFriend(item postData) {
	for i := range self.Follower {
		if self.Follower[i].Userfollowerid == item.Follower[0].Userfollowerid {
			self.Follower = append(self.Follower[:i], self.Follower[i+1:]...)
			break
		}
	}
}

func Addfollower(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	result := getData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Can't follow, error", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	newdata := getData{}
	newdata = result
	a := res.Follower[0].Userfollowerid
	item1 := getProduct{Userfollowerid: a}
	newdata.AddItem(item1)
	db.Update(result, newdata)
	response = shared.ReturnMessage(true, "Successfully followed", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.Follower = append(box.Follower, item)
	return box.Follower
}

func UpdateParentStatus(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := getData{}
	newdata := getData{}

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
	res = *u
	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	newdata = results
	if err != nil {
		response = shared.ReturnMessage(false, "Record not added", 400, "")
		return c.JSON(http.StatusOK, response)
	}

	for i := range results.Follower {
		if results.Follower[i].Followid == res.FollowID {
			notification.AddMentorFollwerHistory(results.Follower[i].Userfollowerid, res.UserID)
			newdata.Follower[i].ParentStatus = 1

			notification.AddMentorAproveHistory(res.UserID, results.Follower[i].Userfollowerid)
		}
	}
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	db.Update(results, newdata)
	response = shared.ReturnMessage(true, "Parent status updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

func UpdateMessageStatus(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	results := getData{}
	newdata := getData{}

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
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
	//email := res.UserID
	//fmt.Println("/n", email)
	//password:=c.FormValue("password")

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	//fmt.Println(email)
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	newdata = results
	if err != nil {
		//log.Fatal(err)
	}

	for i := range results.Follower {
		if results.Follower[i].Followid == res.FollowID {
			newdata.Follower[i].MessageStatus = 1
			notification.AddMentorMsgAproveHistory(res.UserID, results.Follower[i].Userfollowerid)
		}
	}
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	db.Update(results, newdata)

	defer session.Close()

	return c.JSON(http.StatusOK, 1)

}
