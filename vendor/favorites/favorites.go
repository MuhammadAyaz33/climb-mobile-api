package favorites

import (
	"encoding/json"
	"fmt"
	"log"
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
	CommentID     string
	CommentUserID string
	Comment       string
}
type likesgetproduct struct {
	LikeUserID string
}
type getData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ContributionID string
	Likes          []likesgetproduct
	Comments       []getProduct
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	CommentID     string `json:"commentid"`
	CommentUserID string `json:"commentuserid"`
	Comment       string `json:"comment"`
}
type likespostProduct struct {
	LikeUserID string `json:"likeuserid"`
}
type postData struct {
	ID             bson.ObjectId      `json:"_id" bson:"_id,omitempty"`
	ContributionID string             `json:"contributionid"`
	Likes          []likespostProduct `json:"likes"`
	Comments       []postProduct      `json:"comments"`
}
type Res struct {
	Data []postData `json:"Data"`
}

var response shared.Response

//GET *********************************************************************************
func GetAllFvrtData(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
	results := res{}
	err = db.Find(bson.M{}).All(&results.Data)
	if err != nil {
		log.Fatal(err)
	}
	if results.Data == nil {
		var d = postData{
			ContributionID: "",
			Likes:          []likespostProduct{},
			Comments:       []postProduct{},
		}
		defer session.Close()
		response = shared.ReturnMessage(false, "Database Not Connected", 401, d)
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

func GetLikesAndComments(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
	results := res{}
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	err = db.Find(bson.M{"contributionid": res.ContributionID}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data[0])
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

//POST *********************************************************************************
func AddComments(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
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
	result := getData{}
	fmt.Println(res)
	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result)
	mid := xid.New()
	if result.ContributionID == "" {
		//fmt.Println("new data added")
		res.Comments[0].CommentID = mid.String()
		db.Insert(res)
		notification.AddMentorcommentHistory(res.ContributionID, res.Comments[0].CommentUserID)
		defer session.Close()
		return c.JSON(http.StatusOK, "comment added")
	} else {
		//fmt.Println("data update")

		newdata := getData{}
		newdata = result

		a := res.Comments[0].Comment

		item1 := getProduct{CommentID: mid.String(), CommentUserID: res.Comments[0].CommentUserID, Comment: a}

		newdata.AddItem(item1)
		db.Update(result, newdata)
		notification.AddMentorcommentHistory(res.ContributionID, res.Comments[0].CommentUserID)
		defer session.Close()
		return c.JSON(http.StatusOK, "comment added")
	}

	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

//

func UnLike(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u

	result := getData{}
	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result)
	result.removelike(res)

	result1 := getData{}
	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result1)
	if err != nil {
		response = shared.ReturnMessage(false, "Record not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	db.Update(result1, result)
	UpdateUnLikeCount(res.ContributionID)
	likecount := GetLikeCount(res.ContributionID)
	response = shared.ReturnMessage(true, "", 200, likecount)

	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func (self *getData) removelike(item postData) {
	for i := range self.Likes {
		if self.Likes[i].LikeUserID == item.Likes[0].LikeUserID {
			self.Likes = append(self.Likes[:i], self.Likes[i+1:]...)
			//fmt.Println(i)
			fmt.Println("match ho geya")
			break
		}
	}
}
func DeleteComment(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
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

	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result)

	result.removeFriend(res)

	result1 := getData{}

	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result1)
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
	for i := range self.Comments {
		if self.Comments[i].CommentID == item.Comments[0].CommentID {
			self.Comments = append(self.Comments[:i], self.Comments[i+1:]...)
			fmt.Println(i)
			fmt.Println("match ho geya")
			break
		}
	}
}

func AddLikes(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
	u := new(postData)
	if err = c.Bind(&u); err != nil {
	}
	res := postData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)
	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}

	result := getData{}
	err = db.Find(bson.M{"contributionid": res.ContributionID}).One(&result)
	if result.ContributionID != "" {
		newdata := getData{}
		newdata = result
		a := res.Likes[0].LikeUserID

		item2 := likesgetproduct{LikeUserID: a}

		newdata.AddItemlikes(item2)
		db.Update(result, newdata)

		notification.AddMentorLikeHistory(res.ContributionID, res.Likes[0].LikeUserID)
		UpdateLikeCount(res.ContributionID)
		likecount := GetLikeCount(res.ContributionID)
		response = shared.ReturnMessage(true, "", 200, likecount)

		// defer session.Close()
		// return c.JSON(http.StatusOK, likecount)
	} else {
		//fmt.Println("new data add")
		db.Insert(res)
		notification.AddMentorLikeHistory(res.ContributionID, res.Likes[0].LikeUserID)
		UpdateLikeCount(res.ContributionID)
		likecount := GetLikeCount(res.ContributionID)
		response = shared.ReturnMessage(true, "", 200, likecount)

		// defer session.Close()
		// return c.JSON(http.StatusOK, likecount)
	}

	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func (box *getData) AddItem(item getProduct) []getProduct {
	box.Comments = append(box.Comments, item)
	return box.Comments
}
func (box *getData) AddItemlikes(item likesgetproduct) []likesgetproduct {
	box.Likes = append(box.Likes, item)
	return box.Likes
}

func UpdateLikeCount(contributionid string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	result := shared.ContributionData{}
	contributionObjectID := bson.ObjectIdHex(contributionid)
	err = db.Find(bson.M{"_id": contributionObjectID}).One(&result)
	newdata := shared.ContributionData{}
	newdata = result
	likes := newdata.Likes
	likes++
	newdata.Likes = likes
	if err != nil {
		fmt.Println("some thing goes wrong")
	}

	db.Update(result, newdata)
}
func UpdateUnLikeCount(contributionid string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	result := shared.ContributionData{}
	contributionObjectID := bson.ObjectIdHex(contributionid)
	err = db.Find(bson.M{"_id": contributionObjectID}).One(&result)
	newdata := shared.ContributionData{}
	newdata = result
	likes := newdata.Likes
	likes--
	newdata.Likes = likes
	if likes < 0 {
		newdata.Likes = 0
	}
	if err != nil {
		fmt.Println("some thing goes wrong")
	}

	db.Update(result, newdata)
}
func GetLikeCount(contributionid string) int {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	result := shared.ContributionData{}
	contributionObjectID := bson.ObjectIdHex(contributionid)
	err = db.Find(bson.M{"_id": contributionObjectID}).One(&result)

	if err != nil {
		fmt.Println("some thing goes wrong")
	}

	return result.Likes
}
