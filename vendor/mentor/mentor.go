package mentor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notification"
	"shared"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type userRequest struct {
	UserEmail string `json:"parentemail"`
}
type getMentorRequest struct {
	UserID string `json:"userid"`
}
type mentorRequestResponse struct {
	Status int `json:"status"`
}

type statusChangeRequest struct {
	ID bson.ObjectId `json:"_id"`
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

var response shared.Response

func BecomeMentorRequest(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	u := new(shared.BMentorpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.BMentorpostData{}
	res = *u

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
		notification.AddChildMentorRequestFormHistory(res.UserID)
		response = shared.ReturnMessage(true, "Request Submited", 200, "")
	} else {
		response = shared.ReturnMessage(false, "User Already Submit Request", 409, "")
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func GetAllMentorAdminRequest(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
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

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

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
	fmt.Println("parent get kid mentor request")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	kiddata := ParentgetData{}
	kiddata = GetParentKids(res.UserEmail)
	results := shared.BMentorres{}
	kidrequest := shared.BMentorgetData{}
	for x := range kiddata.Kids {
		fmt.Println(kiddata.Kids[x].KidID)
		kidemail := kiddata.Kids[x].KidID
		err = db.Find(bson.M{"useremail": kidemail, "parentstatus": 0}).One(&kidrequest)
		if err == nil {
			results.Data = append(results.Data, kidrequest)
		}

	}

	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetParentKids(parentemail string) ParentgetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PARENTCOLLECTION)
	results := ParentgetData{}

	err = db.Find(bson.M{"parentemail": parentemail}).One(&results)

	if err != nil {
		//log.Fatal(err)
	}
	fmt.Println(results)
	defer session.Close()
	return results
}

func UpdateParentStatus(c echo.Context) error {
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
	fmt.Println("update parent status of mentor request")
	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := shared.BMentorpostData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		defer session.Close()
		return c.JSON(http.StatusOK, 0)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.ParentStatus = 1
	db.Update(result, newdata)
	notification.AddParentMentorRequestApprove(result.UserID)

	defer session.Close()
	return c.JSON(http.StatusOK, 1)

}
func UpdateRejectParentMentorStatus(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	u := new(shared.BMentorpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.BMentorpostData{}
	res = *u
	result := shared.BMentorpostData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusNotFound, response)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.ParentStatus = 2
	db.Update(result, newdata)
	notification.AddParentMentorRequestReject(result.UserID)

	response = shared.ReturnMessage(true, "Record Updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func UpdateAdminStatus(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(shared.BMentorpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.BMentorpostData{}
	//fmt.Println("this is C:", postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("update admin status of mentor request")
	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := shared.BMentorpostData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		defer session.Close()
		return c.JSON(http.StatusOK, 0)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.AdminStatus = 1
	db.Update(result, newdata)
	notification.AddAdminMentorRequestApprove(result.UserID)

	defer session.Close()
	return c.JSON(http.StatusOK, 1)
}
func UpdateAdminRejectStatus(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	u := new(shared.BMentorpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.BMentorpostData{}
	res = *u
	result := shared.BMentorpostData{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusNotFound, response)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.AdminStatus = 2
	db.Update(result, newdata)
	notification.AddAdminMentorRequestReject(result.UserID)
	response = shared.ReturnMessage(true, "Record Updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func GetMentorRequest(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(getMentorRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := getMentorRequest{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("get mentor request by user id")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := shared.BMentorgetData{}
	response := mentorRequestResponse{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		response.Status = 0
		defer session.Close()
		return c.JSON(http.StatusOK, &response)
		//results.Data = append(results.Data, kidrequest)
	}
	response.Status = 1
	buff, _ := json.Marshal(&response)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &response)
	defer session.Close()
	return c.JSON(http.StatusOK, &response)

}
func GetMentorRequestDetail(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(getMentorRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := getMentorRequest{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("get mentor request by user id")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	result := shared.BMentorgetData{}
	// response := mentorRequestResponse{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		// response.Status = 0
		defer session.Close()
		return c.JSON(http.StatusOK, "no request found")
		//results.Data = append(results.Data, kidrequest)
	}
	// response.Status = 1
	buff, _ := json.Marshal(&result)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &result)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)

}
