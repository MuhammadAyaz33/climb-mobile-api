package mentor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"shared"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func Adduser(c echo.Context) (err error) {

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
