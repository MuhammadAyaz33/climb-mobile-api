package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"
	"time"
	"user"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//Mentor history Get data
type FollowerGet struct {
	UserFollowerID             string
	UserFollowerName           string
	UserFollowerProfilePicture string
	NotificationTime           time.Time
}
type FollowerGetWithType struct {
	FollowerGet
	Type string
}
type LikesGet struct {
	ContributionID         string
	ContributionTitle      string
	LikeUserID             string
	LikeUserName           string
	LikeUserProfilePicture string
	NotificationTime       time.Time
}
type LikesGetWithType struct {
	LikesGet
	Type string
}
type CommentsGet struct {
	ContributionID            string
	ContributionTitle         string
	CommentUserID             string
	CommentUserName           string
	CommentUserProfilePicture string
	NotificationTime          time.Time
}
type CommentsGetWithType struct {
	CommentsGet
	Type string
}
type MentorCreateContributionGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
	ContributionID       string
	ContributionTitle    string
	NotificationTime     time.Time
}
type MentorCreateContributionGetWithType struct {
	MentorCreateContributionGet
	Type string
}
type ChildCreateContributionGet struct {
	ChildID             string
	ChildUserName       string
	ChildProfilePicture string
	ContributionID      string
	ContributionTitle   string
	NotificationTime    time.Time
}
type ChildCreateContributionGetWithType struct {
	ChildCreateContributionGet
	Type string
}
type AproveMentorGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
	NotificationTime     time.Time
}
type AproveMentorGetWithType struct {
	AproveMentorGet
	Type string
}
type AproveMentorMsgGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
	NotificationTime     time.Time
}
type AproveMentorMsgGetWithType struct {
	AproveMentorMsgGet
	Type string
}
type AdminAproveContributionGet struct {
	ContributionID    string
	ContributionTitle string
	ContributionType  string
	NotificationTime  time.Time
}
type AdminAproveContributionGetWithType struct {
	AdminAproveContributionGet
	Type string
}
type AdminRejectContributionGet struct {
	ContributionID    string
	ContributionTitle string
	ContributionType  string
	NotificationTime  time.Time
}
type AdminRejectContributionGetWithType struct {
	AdminRejectContributionGet
	Type string
}
type ChildSubmitMentorFormGet struct {
	ChildID             string
	ChildUserName       string
	ChildProfilePicture string
	NotificationTime    time.Time
}
type ChildSubmitMentorFormGetWithType struct {
	ChildSubmitMentorFormGet
	Type string
}
type RequestAproveGet struct {
	RequestAprove    bool
	NotificationTime time.Time
}
type RequestAproveGetWithType struct {
	RequestAproveGet
	Type string
}
type RequestRejectGet struct {
	RequestReject    bool
	NotificationTime time.Time
}
type RequestRejectGetWithType struct {
	RequestRejectGet
	Type string
}
type Detail struct {
	UserID          string
	NewNotification bool
	Result          []interface{}
}
type MentoryHistorygetData struct {
	ID                        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID                    string
	AdminMentorRequest        RequestAproveGet
	AdminMentorRequestReject  RequestRejectGet
	ParentMentorRequest       RequestAproveGet
	ParentMentorRequestReject RequestRejectGet
	NewNotification           bool
	Followers                 []FollowerGet
	ContributionLikes         []LikesGet
	ContributionComments      []CommentsGet
	MentorCreateContribution  []MentorCreateContributionGet
	ChildCreateContribution   []ChildCreateContributionGet
	ChildSubmitMentorForm     []ChildSubmitMentorFormGet
	MentorAprovel             []AproveMentorGet
	MentorMsgAprovel          []AproveMentorMsgGet
	AdminAproveContribution   []AdminAproveContributionGet
	AdminRejectContribution   []AdminRejectContributionGet
}
type res struct {
	Data []MentoryHistorygetData
}

//Mentor history post data

type FollowerPost struct {
	UserFollowerID             string    `json:"userfollowerid"`
	UserFollowerName           string    `json:"userfollowername"`
	UserFollowerProfilePicture string    `json:"userfollowerprofilepicture"`
	NotificationTime           time.Time `json:"notificationtime"`
}
type LikesPost struct {
	ContributionID         string    `json:"contributionid"`
	ContributionTitle      string    `json:"contributiontitle"`
	LikeUserID             string    `json:"likeuserid"`
	LikeUserName           string    `json:"likeusername"`
	LikeUserProfilePicture string    `json:"likeuserprofilepicture"`
	NotificationTime       time.Time `json:"notificationtime"`
}
type CommentsPost struct {
	ContributionID            string    `json:"contributionid"`
	ContributionTitle         string    `json:"contributiontitle"`
	CommentUserID             string    `json:"commentuserid"`
	CommentUserName           string    `json:"commentusername"`
	CommentUserProfilePicture string    `json:"commentprofilepicture"`
	NotificationTime          time.Time `json:"notificationtime"`
}
type MentorCreateContributionPost struct {
	MentorID             string    `json:"mentorid"`
	MentorUserName       string    `json:"mentorusername"`
	MentorProfilePicture string    `json:"mentorprofilepicture"`
	ContributionID       string    `json:"contributionid"`
	ContributionTitle    string    `json:"contributiontitle"`
	NotificationTime     time.Time `json:"notificationtime"`
}
type ChildCreateContributionPost struct {
	ChildID             string    `json:"childid"`
	ChildUserName       string    `json:"childusername"`
	ChildProfilePicture string    `json:"childprofilepicture"`
	ContributionID      string    `json:"contributionid"`
	ContributionTitle   string    `json:"contributiontitle"`
	NotificationTime    time.Time `json:"notificationtime"`
}
type AproveMentorPost struct {
	MentorID             string    `json:"mentorid"`
	MentorUserName       string    `json:"mentorusername"`
	MentorProfilePicture string    `json:"mentorprofilepicture"`
	NotificationTime     time.Time `json:"notificationtime"`
}
type AproveMentorMsgPost struct {
	MentorID             string    `json:"mentorid"`
	MentorUserName       string    `json:"mentorusername"`
	MentorProfilePicture string    `json:"mentorprofilepicture"`
	NotificationTime     time.Time `json:"notificationtime"`
}
type AdminAproveContributionPost struct {
	ContributionID    string    `json:"contributionid"`
	ContributionTitle string    `json:"contributiontitle"`
	ContributionType  string    `json:"contributiontype"`
	NotificationTime  time.Time `json:"notificationtime"`
}
type AdminRejectContributionPost struct {
	ContributionID    string    `json:"contributionid"`
	ContributionTitle string    `json:"contributiontitle"`
	ContributionType  string    `json:"contributiontype"`
	NotificationTime  time.Time `json:"notificationtime"`
}
type ChildSubmitMentorFormPost struct {
	ChildID             string    `json:"childid"`
	ChildUserName       string    `json:"childusername"`
	ChildProfilePicture string    `json:"childprofilepicture"`
	NotificationTime    time.Time `json:"notificationtime"`
}
type RequestAprovePost struct {
	RequestAprove    bool      `json:"requestaprove"`
	NotificationTime time.Time `json:"notificationtime"`
}
type RequestRejectPost struct {
	RequestReject    bool      `json:"requestreject"`
	NotificationTime time.Time `json:"notificationtime"`
}
type MentoryHistorypostData struct {
	ID                        bson.ObjectId                  `json:"_id" bson:"_id,omitempty"`
	UserID                    string                         `json:"userid"`
	AdminMentorRequest        RequestAprovePost              `json:"adminmentorrequestaprove"`
	AdminMentorRequestReject  RequestRejectPost              `json:"adminmentorrequestreject"`
	ParentMentorRequest       RequestAprovePost              `json:"parentmentorrequestaprove"`
	ParentMentorRequestReject RequestRejectPost              `json:"parentmentorrequestreject"`
	NewNotification           bool                           `json:"newnotification"`
	Followers                 []FollowerPost                 `json:"followers"`
	ContributionLikes         []LikesPost                    `json:"contributionlikes"`
	ContributionComments      []CommentsPost                 `json:"contributioncomments"`
	MentorCreateContribution  []MentorCreateContributionPost `json:"mentorcreatecontribution"`
	ChildCreateContribution   []ChildCreateContributionPost  `json:"childcreatecontribution"`
	ChildSubmitMentorForm     []ChildSubmitMentorFormPost    `json:"childsubmitmentorform"`
	MentorAprovel             []AproveMentorPost             `json:"mentoraprovel"`
	MentorMsgAprovel          []AproveMentorMsgPost          `json:"mentormsgaprovel"`
	AdminAproveContribution   []AdminAproveContributionPost  `json:"adminaprovecontribution"`
	AdminRejectContribution   []AdminRejectContributionPost  `json:"adminrejectcontribution"`
}
type Res struct {
	Data []MentoryHistorypostData `json:"Data"`
}
type GetUserData struct {
	UserID string `json:"userid"`
}

var response shared.Response

// user history
func AddMentorFollwerHistory(userid string, followerid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}
	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": userid}).One(&result)
	followeruserinfo := shared.UserinfoUpdategetData{}
	bsonObjectID := bson.ObjectIdHex(followerid)
	followeruserinfo = UserInfo(bsonObjectID)
	currentdate := time.Now().UTC()
	if err != nil {
		newfollower := MentoryHistorypostData{}
		newfollower.UserID = userid
		newfollower.NewNotification = true
		item := FollowerPost{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture, NotificationTime: currentdate}
		newfollower.AddItemPostFollow(item)
		db.Insert(newfollower)
		//fmt.Println(newfollower)
	} else {
		if len(result.Followers) > 0 {
			fmt.Println("data available")
			for i := range result.Followers {
				if result.Followers[i].UserFollowerID == followerid {
					fmt.Println("follwer already added")
				} else {
					fmt.Println("user exit update history")
					newdata := MentoryHistorygetData{}
					newdata = result
					item := FollowerGet{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture, NotificationTime: currentdate}
					newdata.AddItemGetFollow(item)
					newdata.NewNotification = true
					db.Update(result, newdata)
				}
			}
		} else {
			fmt.Println("no data available")
			newdata := MentoryHistorygetData{}
			newdata = result
			item := FollowerGet{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture, NotificationTime: currentdate}
			newdata.AddItemGetFollow(item)
			newdata.NewNotification = true
			db.Update(result, newdata)
		}

	}
	defer session.Close()
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
	userinfo := shared.UserinfoUpdategetData{}
	userinfo = UserInfo(UserIDconv)
	currentdate := time.Now().UTC()

	if err != nil {
		newlikes := MentoryHistorypostData{}
		newlikes.UserID = contributiondetail.UserID
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := LikesPost{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, LikeUserID: hexuserid, LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
		newlikes.AddItemPostLike(item)
		newlikes.NewNotification = true
		db.Insert(newlikes)
		//fmt.Println(newfollower)
	} else {
		fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := LikesGet{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, LikeUserID: hexuserid, LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
		newdata.AddItemGetLike(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
	}
	defer session.Close()
}
func (box *MentoryHistorypostData) AddItemPostLike(item LikesPost) []LikesPost {
	box.ContributionLikes = append(box.ContributionLikes, item)
	return box.ContributionLikes
}
func (box *MentoryHistorygetData) AddItemGetLike(item LikesGet) []LikesGet {
	box.ContributionLikes = append(box.ContributionLikes, item)
	return box.ContributionLikes
}
func AddMentorcommentHistory(contributionid string, commentuserid string) {

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

	UserIDconv := bson.ObjectIdHex(commentuserid)
	userinfo := shared.UserinfoUpdategetData{}
	userinfo = UserInfo(UserIDconv)
	currentdate := time.Now().UTC()
	if err != nil {
		newlikes := MentoryHistorypostData{}
		newlikes.UserID = contributiondetail.UserID
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := CommentsPost{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, CommentUserID: hexuserid, CommentUserName: userinfo.FullName, CommentUserProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
		newlikes.AddItemPostComment(item)
		newlikes.NewNotification = true
		db.Insert(newlikes)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := CommentsGet{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, CommentUserID: hexuserid, CommentUserName: userinfo.FullName, CommentUserProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
		newdata.AddItemGetComment(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}

func (box *MentoryHistorypostData) AddItemPostComment(item CommentsPost) []CommentsPost {
	box.ContributionComments = append(box.ContributionComments, item)
	return box.ContributionComments
}
func (box *MentoryHistorygetData) AddItemGetComment(item CommentsGet) []CommentsGet {
	box.ContributionComments = append(box.ContributionComments, item)
	return box.ContributionComments
}

func UserInfo(userid bson.ObjectId) shared.UserinfoUpdategetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.UserinfoUpdategetData{}

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
func UserInfoByEmail(useremail string) shared.UserinfoUpdategetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.UserinfoUpdategetData{}

	if err != nil {
	}

	err = db.Find(bson.M{"email": useremail}).One(&results)

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
func MentorContributionInfo(mentorid string) shared.Contributionres {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	if err != nil {
	}

	err = db.Find(bson.M{"userid": mentorid}).All(&results.Data)

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
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)

	//newdata := getData{}

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("get notification")
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r Res
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	results := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	if err != nil {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusOK, 0)
	}
	var all []interface{}

	if results.ParentMentorRequestReject.RequestReject == true {

		a := RequestRejectGetWithType{}
		a.RequestRejectGet = results.ParentMentorRequestReject
		a.Type = "ParentMentorRequestReject"
		s := []interface{}{a}
		all = append(s, all...)
	}

	if results.ParentMentorRequest.RequestAprove == true {

		a := RequestAproveGetWithType{}
		a.RequestAproveGet = results.ParentMentorRequest
		a.Type = "ParentMentorRequest"
		s := []interface{}{a}
		all = append(s, all...)
	}

	if results.AdminMentorRequestReject.RequestReject == true {

		a := RequestRejectGetWithType{}
		a.RequestRejectGet = results.AdminMentorRequestReject
		a.Type = "AdminMentorRequestReject"
		s := []interface{}{a}
		all = append(s, all...)
	}

	if results.AdminMentorRequest.RequestAprove == true {

		a := RequestAproveGetWithType{}
		a.RequestAproveGet = results.AdminMentorRequest
		a.Type = "AdminMentorRequest"
		s := []interface{}{a}
		all = append(s, all...)
	}
	if len(results.MentorMsgAprovel) > 0 {
		s := make([]interface{}, len(results.MentorMsgAprovel))
		for i, v := range results.MentorMsgAprovel {
			a := AproveMentorMsgGetWithType{}
			a.AproveMentorMsgGet = v
			a.Type = "MentorMsgAprovel"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.MentorAprovel) > 0 {
		s := make([]interface{}, len(results.MentorAprovel))
		for i, v := range results.MentorAprovel {
			a := AproveMentorGetWithType{}
			a.AproveMentorGet = v
			a.Type = "MentorAprovel"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.ChildSubmitMentorForm) > 0 {
		s := make([]interface{}, len(results.ChildSubmitMentorForm))
		for i, v := range results.ChildSubmitMentorForm {
			a := ChildSubmitMentorFormGetWithType{}
			a.ChildSubmitMentorFormGet = v
			a.Type = "ChildSubmitMentorForm"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.ChildCreateContribution) > 0 {
		s := make([]interface{}, len(results.ChildCreateContribution))
		for i, v := range results.ChildCreateContribution {
			a := ChildCreateContributionGetWithType{}
			a.ChildCreateContributionGet = v
			a.Type = "ChildCreateContribution"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.MentorCreateContribution) > 0 {
		s := make([]interface{}, len(results.MentorCreateContribution))
		for i, v := range results.MentorCreateContribution {
			a := MentorCreateContributionGetWithType{}
			a.MentorCreateContributionGet = v
			a.Type = "MentorCreateContribution"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.ContributionComments) > 0 {
		s := make([]interface{}, len(results.ContributionComments))
		for i, v := range results.ContributionComments {
			a := CommentsGetWithType{}
			a.CommentsGet = v
			a.Type = "ContributionComment"
			s[i] = a
		}
		all = append(s, all...)
	}

	if len(results.Followers) > 0 {
		s := make([]interface{}, len(results.Followers))
		for i, v := range results.Followers {
			a := FollowerGetWithType{}
			a.FollowerGet = v
			a.Type = "Followers"
			s[i] = a
		}
		all = append(s, all...)
	}
	if len(results.ContributionLikes) > 0 {
		s := make([]interface{}, len(results.ContributionLikes))
		for i, v := range results.ContributionLikes {
			a := LikesGetWithType{}
			a.LikesGet = v
			a.Type = "ContributionLike"
			s[i] = a
		}
		all = append(s, all...)
	}
	if len(results.AdminAproveContribution) > 0 {
		s := make([]interface{}, len(results.AdminAproveContribution))
		for i, v := range results.AdminAproveContribution {
			a := AdminAproveContributionGetWithType{}
			a.AdminAproveContributionGet = v
			a.Type = "AdminAproveContribution"
			s[i] = a
		}
		all = append(s, all...)
	}
	if len(results.AdminRejectContribution) > 0 {
		s := make([]interface{}, len(results.AdminRejectContribution))
		for i, v := range results.AdminRejectContribution {
			a := AdminRejectContributionGetWithType{}
			a.AdminRejectContributionGet = v
			a.Type = "AdminRejectContribution"
			s[i] = a
		}
		all = append(s, all...)
	}

	var detail Detail
	detail.NewNotification = results.NewNotification
	detail.UserID = results.UserID
	detail.Result = all
	// ss := []interface{}{detail}
	// all = append(ss, all...)

	//ParentInfo("mohd.kasimnazesser@gmail.com")
	// x := all[0].([]interface{})["NotificationTime"].(string)
	// spew.Dump(all)
	// fmt.Println(x)
	defer session.Close()

	return c.JSON(http.StatusOK, &detail)

}

func AddMentorCreatContributionHistory(mentorid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}
	result := MentoryHistorygetData{}
	followerresult := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": mentorid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {

	} else {
		mentoridconv := bson.ObjectIdHex(mentorid)
		mentoruserinfo := shared.UserinfoUpdategetData{}
		mentoruserinfo = UserInfo(mentoridconv)
		hexmentorid := fmt.Sprintf("%x", string(mentoruserinfo.ID))

		mentorcontributiondetail := shared.Contributionres{}
		mentorcontributiondetail = MentorContributionInfo(mentorid)
		l := len(mentorcontributiondetail.Data)

		contributionid := fmt.Sprintf("%x", string(mentorcontributiondetail.Data[l-1].ID))
		contributiontitle := mentorcontributiondetail.Data[l-1].Title
		fmt.Println(mentorcontributiondetail.Data[l-1])
		for i := range result.Followers {
			userid := result.Followers[i].UserFollowerID
			err = db.Find(bson.M{"userid": userid}).One(&followerresult)
			if err != nil {
				createcontribution := MentoryHistorypostData{}
				createcontribution.UserID = userid
				item := MentorCreateContributionPost{MentorID: hexmentorid, MentorUserName: mentoruserinfo.FullName, MentorProfilePicture: mentoruserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle, NotificationTime: currentdate}
				createcontribution.AddItemPostCreateContribution(item)
				createcontribution.NewNotification = true
				db.Insert(createcontribution)
				fmt.Println("user not found add new data")
			} else {
				newdata := MentoryHistorygetData{}
				newdata = followerresult
				item := MentorCreateContributionGet{MentorID: hexmentorid, MentorUserName: mentoruserinfo.FullName, MentorProfilePicture: mentoruserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle, NotificationTime: currentdate}
				newdata.AddItemGetCreateContribution(item)
				newdata.NewNotification = true
				db.Update(followerresult, newdata)
				fmt.Println("user found update data")
			}

		}
	}
	defer session.Close()
}

func (box *MentoryHistorypostData) AddItemPostCreateContribution(item MentorCreateContributionPost) []MentorCreateContributionPost {
	box.MentorCreateContribution = append(box.MentorCreateContribution, item)
	return box.MentorCreateContribution
}
func (box *MentoryHistorygetData) AddItemGetCreateContribution(item MentorCreateContributionGet) []MentorCreateContributionGet {
	box.MentorCreateContribution = append(box.MentorCreateContribution, item)
	return box.MentorCreateContribution
}
func ParentInfo(kidemail string) user.ParentgetData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PARENTCOLLECTION)
	results := user.ParentgetData{}

	if err != nil {
	}

	err = db.Find(bson.M{"kids": bson.M{"kidid": kidemail}}).One(&results)

	if err != nil {
		fmt.Println("no data")
	} else {
		fmt.Println(results)
	}

	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return results
}
func AddChildCreatContributionHistory(mentorid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	kididconv := bson.ObjectIdHex(mentorid)
	kiduserinfo := shared.UserinfoUpdategetData{}
	kiduserinfo = UserInfo(kididconv)
	kidemail := kiduserinfo.Email

	childinfo := user.ParentgetData{}
	childinfo = ParentInfo(kidemail)
	currentdate := time.Now().UTC()
	if childinfo.ParentEmail != "" {
		parentinfo := shared.UserinfoUpdategetData{}
		parentinfo = UserInfoByEmail(childinfo.ParentEmail)

		parentid := fmt.Sprintf("%x", string(parentinfo.ID))
		result := MentoryHistorygetData{}
		err = db.Find(bson.M{"userid": parentid}).One(&result)
		if err != nil {
			createcontribution := MentoryHistorypostData{}
			createcontribution.UserID = parentid
			childid := fmt.Sprintf("%x", string(kiduserinfo.ID))

			childcontributiondetail := shared.Contributionres{}
			childcontributiondetail = MentorContributionInfo(mentorid)
			l := len(childcontributiondetail.Data)
			contributionid := fmt.Sprintf("%x", string(childcontributiondetail.Data[l-1].ID))
			contributiontitle := childcontributiondetail.Data[l-1].Title

			item := ChildCreateContributionPost{ChildID: childid, ChildUserName: kiduserinfo.FullName, ChildProfilePicture: kiduserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle, NotificationTime: currentdate}
			createcontribution.AddItemPostCreateContributionKid(item)
			createcontribution.NewNotification = true
			db.Insert(createcontribution)
		} else {
			childid := fmt.Sprintf("%x", string(kiduserinfo.ID))
			childcontributiondetail := shared.Contributionres{}
			childcontributiondetail = MentorContributionInfo(mentorid)
			l := len(childcontributiondetail.Data)
			contributionid := fmt.Sprintf("%x", string(childcontributiondetail.Data[l-1].ID))
			contributiontitle := childcontributiondetail.Data[l-1].Title

			newdata := MentoryHistorygetData{}
			newdata = result
			item := ChildCreateContributionGet{ChildID: childid, ChildUserName: kiduserinfo.FullName, ChildProfilePicture: kiduserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle, NotificationTime: currentdate}
			newdata.AddItemGetCreateContributionKid(item)
			newdata.NewNotification = true
			db.Update(result, newdata)
		}
	}
	defer session.Close()
}

func (box *MentoryHistorypostData) AddItemPostCreateContributionKid(item ChildCreateContributionPost) []ChildCreateContributionPost {
	box.ChildCreateContribution = append(box.ChildCreateContribution, item)
	return box.ChildCreateContribution
}
func (box *MentoryHistorygetData) AddItemGetCreateContributionKid(item ChildCreateContributionGet) []ChildCreateContributionGet {
	box.ChildCreateContribution = append(box.ChildCreateContribution, item)
	return box.ChildCreateContribution
}
func AddMentorAproveHistory(Userid string, followerid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	// UserIDconv := bson.ObjectIdHex(Userid)
	// userinfo := shared.UsergetData{}
	// userinfo = UserInfo(UserIDconv)

	followerIDconv := bson.ObjectIdHex(followerid)
	followerinfo := shared.UserinfoUpdategetData{}
	followerinfo = UserInfo(followerIDconv)

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		mentoraprove := MentoryHistorypostData{}
		mentoraprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AproveMentorPost{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture, NotificationTime: currentdate}
		mentoraprove.AddItemPostAproveMentor(item)
		mentoraprove.NewNotification = true
		db.Insert(mentoraprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AproveMentorGet{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture, NotificationTime: currentdate}
		newdata.AddItemGetAproveMentor(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func (box *MentoryHistorypostData) AddItemPostAproveMentor(item AproveMentorPost) []AproveMentorPost {
	box.MentorAprovel = append(box.MentorAprovel, item)
	return box.MentorAprovel
}
func (box *MentoryHistorygetData) AddItemGetAproveMentor(item AproveMentorGet) []AproveMentorGet {
	box.MentorAprovel = append(box.MentorAprovel, item)
	return box.MentorAprovel
}

func AddMentorMsgAproveHistory(Userid string, followerid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	// UserIDconv := bson.ObjectIdHex(Userid)
	// userinfo := shared.UsergetData{}
	// userinfo = UserInfo(UserIDconv)

	followerIDconv := bson.ObjectIdHex(followerid)
	followerinfo := shared.UserinfoUpdategetData{}
	followerinfo = UserInfo(followerIDconv)

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()

	if err != nil {
		mentoraprove := MentoryHistorypostData{}
		mentoraprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AproveMentorMsgPost{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture, NotificationTime: currentdate}
		mentoraprove.AddItemPostAproveMentorMsg(item)
		mentoraprove.NewNotification = true
		db.Insert(mentoraprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AproveMentorMsgGet{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture, NotificationTime: currentdate}
		newdata.AddItemGetAproveMentorMsg(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func (box *MentoryHistorypostData) AddItemPostAproveMentorMsg(item AproveMentorMsgPost) []AproveMentorMsgPost {
	box.MentorMsgAprovel = append(box.MentorMsgAprovel, item)
	return box.MentorMsgAprovel
}
func (box *MentoryHistorygetData) AddItemGetAproveMentorMsg(item AproveMentorMsgGet) []AproveMentorMsgGet {
	box.MentorMsgAprovel = append(box.MentorMsgAprovel, item)
	return box.MentorMsgAprovel
}

func AddAdminAproveContributionHistory(Userid string, contributionid string, contributiontitle string, contributiontype string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AdminAproveContributionPost{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype, NotificationTime: currentdate}
		adminaprove.AddItemPostAdminAprove(item)
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AdminAproveContributionGet{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype, NotificationTime: currentdate}
		newdata.AddItemGetAdminAprove(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func AddAdminRejectContributionHistory(Userid string, contributionid string, contributiontitle string, contributiontype string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AdminRejectContributionPost{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype, NotificationTime: currentdate}
		adminaprove.AddItemPostAdminReject(item)
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AdminRejectContributionGet{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype, NotificationTime: currentdate}
		newdata.AddItemGetAdminReject(item)
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func (box *MentoryHistorypostData) AddItemPostAdminAprove(item AdminAproveContributionPost) []AdminAproveContributionPost {
	box.AdminAproveContribution = append(box.AdminAproveContribution, item)
	return box.AdminAproveContribution
}
func (box *MentoryHistorygetData) AddItemGetAdminAprove(item AdminAproveContributionGet) []AdminAproveContributionGet {
	box.AdminAproveContribution = append(box.AdminAproveContribution, item)
	return box.AdminAproveContribution
}

func (box *MentoryHistorypostData) AddItemPostAdminReject(item AdminRejectContributionPost) []AdminRejectContributionPost {
	box.AdminRejectContribution = append(box.AdminRejectContribution, item)
	return box.AdminRejectContribution
}
func (box *MentoryHistorygetData) AddItemGetAdminReject(item AdminRejectContributionGet) []AdminRejectContributionGet {
	box.AdminRejectContribution = append(box.AdminRejectContribution, item)
	return box.AdminRejectContribution
}

func AddAdminMentorRequestApprove(Userid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		aproverequest := RequestAprovePost{}
		aproverequest.RequestAprove = true
		aproverequest.NotificationTime = currentdate
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		adminaprove.AdminMentorRequest = aproverequest
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		aproverequest := RequestAproveGet{}
		aproverequest.RequestAprove = true
		aproverequest.NotificationTime = currentdate
		newdata.AdminMentorRequest = aproverequest
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func AddAdminMentorRequestReject(Userid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		rejectrequest := RequestRejectPost{}
		rejectrequest.RequestReject = true
		rejectrequest.NotificationTime = currentdate

		adminaprove.AdminMentorRequestReject = rejectrequest
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		rejectrequest := RequestRejectGet{}
		rejectrequest.RequestReject = true
		rejectrequest.NotificationTime = currentdate

		newdata.AdminMentorRequestReject = rejectrequest
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func AddParentMentorRequestApprove(Userid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))

		aproverequest := RequestAprovePost{}
		aproverequest.RequestAprove = true
		aproverequest.NotificationTime = currentdate

		adminaprove.ParentMentorRequest = aproverequest
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		aproverequest := RequestAproveGet{}
		aproverequest.RequestAprove = true
		aproverequest.NotificationTime = currentdate

		newdata.ParentMentorRequest = aproverequest
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func AddParentMentorRequestReject(Userid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	result := MentoryHistorygetData{}
	err = db.Find(bson.M{"userid": Userid}).One(&result)
	currentdate := time.Now().UTC()
	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		rejectrequest := RequestRejectPost{}
		rejectrequest.RequestReject = true
		rejectrequest.NotificationTime = currentdate

		adminaprove.ParentMentorRequestReject = rejectrequest
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		rejectrequest := RequestRejectGet{}
		rejectrequest.RequestReject = true
		rejectrequest.NotificationTime = currentdate
		newdata.ParentMentorRequestReject = rejectrequest
		newdata.NewNotification = true
		db.Update(result, newdata)
		//AddMentorCreatContributionHistory(contributiondetail.UserID)
	}
	defer session.Close()
}
func ChangeNotificationStatus(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Database Not Connected", 401, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	results := MentoryHistorygetData{}

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
	res = *u

	err = db.Find(bson.M{"userid": res.UserID}).One(&results)
	if err != nil {
		response = shared.ReturnMessage(false, "Status Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	newdata := MentoryHistorygetData{}
	newdata = results
	newdata.NewNotification = false
	db.Update(results, newdata)
	//ParentInfo("mohd.kasimnazesser@gmail.com")
	response = shared.ReturnMessage(true, "Status Found", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func AddChildMentorRequestFormHistory(Userid string) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORHISTORYCOLLECTION)
	if err != nil {
		fmt.Println(err)
	}

	UserIDconv := bson.ObjectIdHex(Userid)
	userinfo := shared.UserinfoUpdategetData{}
	userinfo = UserInfo(UserIDconv)
	currentdate := time.Now().UTC()
	if userinfo.ParentEmail != "" {
		result := MentoryHistorygetData{}
		parentinfo := UserInfoByEmail(userinfo.ParentEmail)

		parentid := fmt.Sprintf("%x", string(parentinfo.ID))

		err = db.Find(bson.M{"userid": parentid}).One(&result)

		if err != nil {
			mentoraprove := MentoryHistorypostData{}
			mentoraprove.UserID = parentid
			//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
			item := ChildSubmitMentorFormPost{ChildID: Userid, ChildUserName: userinfo.FullName, ChildProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
			mentoraprove.AddMentorFormRequestPost(item)
			mentoraprove.NewNotification = true
			db.Insert(mentoraprove)
			//fmt.Println(newfollower)
		} else {
			//fmt.Println("user exit update history")
			newdata := MentoryHistorygetData{}
			newdata = result

			item := ChildSubmitMentorFormGet{ChildID: Userid, ChildUserName: userinfo.FullName, ChildProfilePicture: userinfo.ProfilePicture, NotificationTime: currentdate}
			newdata.AddMentorFormRequestGet(item)
			newdata.NewNotification = true
			db.Update(result, newdata)
			//AddMentorCreatContributionHistory(contributiondetail.UserID)
		}
	}

	defer session.Close()
}
func (box *MentoryHistorypostData) AddMentorFormRequestPost(item ChildSubmitMentorFormPost) []ChildSubmitMentorFormPost {
	box.ChildSubmitMentorForm = append(box.ChildSubmitMentorForm, item)
	return box.ChildSubmitMentorForm
}
func (box *MentoryHistorygetData) AddMentorFormRequestGet(item ChildSubmitMentorFormGet) []ChildSubmitMentorFormGet {
	box.ChildSubmitMentorForm = append(box.ChildSubmitMentorForm, item)
	return box.ChildSubmitMentorForm
}
