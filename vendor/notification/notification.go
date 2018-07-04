package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"shared"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type FollowerGet struct {
	UserFollowerID             string
	UserFollowerName           string
	UserFollowerProfilePicture string
}
type LikesGet struct {
	ContributionID         string
	ContributionTitle      string
	LikeUserID             string
	LikeUserName           string
	LikeUserProfilePicture string
}
type CommentsGet struct {
	ContributionID            string
	ContributionTitle         string
	CommentUserID             string
	CommentUserName           string
	CommentUserProfilePicture string
}
type MentoryHistorygetData struct {
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID               string
	Followers            []FollowerGet
	ContributionLikes    []LikesGet
	ContributionComments []CommentsGet
}
type res struct {
	Data []MentoryHistorygetData
}

//data from post********************************************************************

type FollowerPost struct {
	UserFollowerID             string `json:"userfollowerid"`
	UserFollowerName           string `json:"userfollowername"`
	UserFollowerProfilePicture string `json:"userfollowerprofilepicture"`
}
type LikesPost struct {
	ContributionID         string `json:"contributionid"`
	ContributionTitle      string `json:"contributiontitle"`
	LikeUserID             string `json:"likeuserid"`
	LikeUserName           string `json:"likeusername"`
	LikeUserProfilePicture string `json:"likeuserprofilepicture"`
}
type CommentsPost struct {
	ContributionID            string `json:"contributionid"`
	ContributionTitle         string `json:"contributiontitle"`
	CommentUserID             string `json:"commentuserid"`
	CommentUserName           string `json:"commentusername"`
	CommentUserProfilePicture string `json:"commentprofilepicture"`
}
type MentoryHistorypostData struct {
	ID                   bson.ObjectId  `json:"_id" bson:"_id,omitempty"`
	UserID               string         `json:"userid"`
	Followers            []FollowerPost `json:"followers"`
	ContributionLikes    []LikesPost    `json:"contributionlikes"`
	ContributionComments []CommentsPost `json:"contributioncomments"`
}
type Res struct {
	Data []MentoryHistorypostData `json:"Data"`
}
type GetUserData struct {
	UserID string `json:"userid"`
}

func AddMentorFollwerHistory(userid string, followerid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}
	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": userid}).One(&result)
	followeruserinfo := shared.UsergetData{}
	bsonObjectID := bson.ObjectIdHex(followerid)
	followeruserinfo = UserInfo(bsonObjectID)

	if err != nil {
		newfollower := MentoryHistorypostData{}
		newfollower.UserID = userid
		item := FollowerPost{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture}
		newfollower.AddItemPostFollow(item)
		db.Insert(newfollower)
		//fmt.Println(newfollower)
	} else {
		fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		item := FollowerGet{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture}
		newdata.AddItemGetFollow(item)
		db.Update(result, newdata)
	}

}
func (box *MentoryHistorypostData) AddItemPostFollow(item FollowerPost) []FollowerPost {
	box.Followers = append(box.Followers, item)
	return box.Followers
}
func (box *MentoryHistorygetData) AddItemGetFollow(item FollowerGet) []FollowerGet {
	box.Followers = append(box.Followers, item)
	return box.Followers
}

func AddMentorLikeHistory(contributionid string, likeuserid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}
	//get contribution detail
	ContributionIDconv := bson.ObjectIdHex(contributionid)
	contributiondetail := shared.ContributionData{}
	contributiondetail = ContributionInfo(ContributionIDconv)

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": contributiondetail.UserID}).One(&result)

	UserIDconv := bson.ObjectIdHex(likeuserid)
	userinfo := shared.UsergetData{}
	userinfo = UserInfo(UserIDconv)

	if err != nil {
		newlikes := MentoryHistorypostData{}
		newlikes.UserID = contributiondetail.UserID
		item := LikesPost{ContributionID: contributiondetail.ID.String(), ContributionTitle: contributiondetail.Title, LikeUserID: userinfo.ID.String(), LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture}
		newlikes.AddItemPostLike(item)
		db.Insert(newlikes)
		//fmt.Println(newfollower)
	} else {
		fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		item := LikesGet{ContributionID: contributiondetail.ID.String(), ContributionTitle: contributiondetail.Title, LikeUserID: likeuserid, LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture}
		newdata.AddItemGetLike(item)
		db.Update(result, newdata)
	}

}
func (box *MentoryHistorypostData) AddItemPostLike(item LikesPost) []LikesPost {
	box.ContributionLikes = append(box.ContributionLikes, item)
	return box.ContributionLikes
}
func (box *MentoryHistorygetData) AddItemGetLike(item LikesGet) []LikesGet {
	box.ContributionLikes = append(box.ContributionLikes, item)
	return box.ContributionLikes
}

func UserInfo(userid bson.ObjectId) shared.UsergetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.UsergetData{}

	if err != nil {
	}

	err = db.Find(bson.M{"_id": userid}).One(&results)

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
func ContributionInfo(contributionid bson.ObjectId) shared.ContributionData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.ContributionData{}

	if err != nil {
	}

	err = db.Find(bson.M{"_id": contributionid}).One(&results)

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

func GetUserMentorHistory(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	results := MentoryHistorygetData{}
	//newdata := getData{}

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

	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	if err != nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "no history found")
	}
	defer session.Close()

	return c.JSON(http.StatusOK, &results)

}
