package following

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notification"
	"os"
	"shared"
	"strconv"

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
	Age      string        `json:"userage"`
	FollowID string        `json:"followid"`
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
	//fmt.Println("/n", email)
	//password:=c.FormValue("password")

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	//fmt.Println(email)
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

func GetfollowerByEmail(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.UsergetData{}
	follower := res{}

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
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

	email := res.Email

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"email": email}).One(&results)

	if err != nil {
		//log.Fatal(err)
	}
	hexid := fmt.Sprintf("%x", string(results.ID))
	fmt.Println(hexid)
	follower = GetfollowerById(hexid)
	//fmt.Println(results)
	buff, _ := json.Marshal(&follower)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &follower)
	defer session.Close()
	return c.JSON(http.StatusOK, &follower)

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
	db := session.DB(shared.DBName).C(shared.MENTORCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	//fmt.Println(name)
	//u:=new (postData)
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
	result := getData{}
	//fmt.Println(res)
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	userfollower := postData{}
	fid := xid.New()
	if result.UserID == "" {
		//fmt.Println("ni match hova add kr do")
		userfollower.UserID = res.UserID
		// userfollower.Follower[0].Userfollowerid = res.Userfollowerid
		// userfollower.Follower[0].ParentStatus = 0
		// userfollower.Follower[0].MessageStatus = 0
		var agestatus int
		var msgstatus int
		age, err := strconv.Atoi(res.Age)
		if err != nil {
			fmt.Println(err)
			//fmt.Println(age)
		}

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
		db.Insert(userfollower)
	} else {
		//fmt.Println("match ho geya hai update kro")
		for x := range result.Follower {
			if result.Follower[x].Userfollowerid == res.Follower[0].Userfollowerid {
				//db.Insert(res)
				defer session.Close()
				return c.JSON(http.StatusOK, "user already follow")
			}
		}

		newdata := getData{}
		newdata = result

		a := res.Follower[0].Userfollowerid
		var agestatus int
		var msgstatus int
		age, err := strconv.Atoi(res.Age)
		if err != nil {
			fmt.Println(err)
			//fmt.Println(age)
		}
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
		db.Update(result, newdata)
	}

	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

//
func (box *postData) AddItem11(item postProduct) []postProduct {
	box.Follower = append(box.Follower, item)
	return box.Follower
}

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

func UpdateParentStatus(c echo.Context) error {

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
			notification.AddMentorFollwerHistory(results.Follower[i].Userfollowerid, res.UserID)
			newdata.Follower[i].ParentStatus = 1

			notification.AddMentorAproveHistory(res.UserID, results.Follower[i].Userfollowerid)
		}
	}
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	db.Update(results, newdata)

	defer session.Close()

	return c.JSON(http.StatusOK, 1)

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
