package mentor

import (
	"encoding/json"
	"fmt"
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
		response = shared.ReturnMessage(false, "Database Not Connected", 501, "")
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
		response = shared.ReturnMessage(true, "Mentor added", 200, "")
	} else {
		response = shared.ReturnMessage(false, "Mentor already exist", 409, "")
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func GetAllMentorAdminRequest(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)
	results := shared.BMentorres{}
	err = db.Find(bson.M{"adminstatus": 0}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Error finding  admin status", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Admin status not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Admin status", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func GetMentorParentsRequest(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(userRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := userRequest{}
	res = *u

	kiddata := ParentgetData{}
	kiddata = GetParentKids(res.UserEmail)
	results := shared.BMentorres{}
	kidrequest := shared.BMentorgetData{}
	for x := range kiddata.Kids {
		kidemail := kiddata.Kids[x].KidID
		err = db.Find(bson.M{"useremail": kidemail, "parentstatus": 0}).One(&kidrequest)
		if err == nil {
			results.Data = append(results.Data, kidrequest)
		}
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Parent status not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Parent status", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
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
		response = shared.ReturnMessage(false, "Error finding record", 404, "")
		defer session.Close()
		return c.JSON(http.StatusOK, response)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.ParentStatus = 1
	db.Update(result, newdata)
	notification.AddParentMentorRequestApprove(result.UserID)
	response = shared.ReturnMessage(true, "Status updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
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

	response = shared.ReturnMessage(true, "Status updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func UpdateAdminStatus(c echo.Context) error {
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
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusNotFound, response)
		//results.Data = append(results.Data, kidrequest)
	}
	// res.ContributionStatus = 1
	newdata := shared.BMentorpostData{}
	newdata = result
	newdata.AdminStatus = 1
	db.Update(result, newdata)
	notification.AddAdminMentorRequestApprove(result.UserID)
	response = shared.ReturnMessage(true, "Status updated", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
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
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(getMentorRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := getMentorRequest{}
	res = *u

	result := shared.BMentorgetData{}
	resp := mentorRequestResponse{}

	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		resp.Status = 0
		response = shared.ReturnMessage(false, "Error finding data", 404, resp)
	} else {
		resp.Status = 1
		response = shared.ReturnMessage(true, "Mentor dtatus", 200, resp)
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func GetMentorRequestDetail(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	u := new(getMentorRequest)
	if err := c.Bind(&u); err != nil {
	}
	res := getMentorRequest{}
	res = *u

	result := shared.BMentorgetData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Error finding data", 404, "")
	} else {
		response = shared.ReturnMessage(true, "Mentor dtatus", 200, result)
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
