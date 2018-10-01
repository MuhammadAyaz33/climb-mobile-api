package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"
	"user"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//Mentor history Get data
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
type MentorCreateContributionGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
	ContributionID       string
	ContributionTitle    string
}
type ChildCreateContributionGet struct {
	ChildID             string
	ChildUserName       string
	ChildProfilePicture string
	ContributionID      string
	ContributionTitle   string
}
type AproveMentorGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
}
type AproveMentorMsgGet struct {
	MentorID             string
	MentorUserName       string
	MentorProfilePicture string
}
type AdminAproveContributionGet struct {
	ContributionID    string
	ContributionTitle string
	ContributionType  string
}
type AdminRejectContributionGet struct {
	ContributionID    string
	ContributionTitle string
	ContributionType  string
}
type ChildSubmitMentorFormGet struct {
	ChildID             string
	ChildUserName       string
	ChildProfilePicture string
}
type MentoryHistorygetData struct {
	ID                        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserID                    string
	AdminMentorRequest        bool
	AdminMentorRequestReject  bool
	ParentMentorRequest       bool
	ParentMentorRequestReject bool
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
type MentorCreateContributionPost struct {
	MentorID             string `json:"mentorid"`
	MentorUserName       string `json:"mentorusername"`
	MentorProfilePicture string `json:"mentorprofilepicture"`
	ContributionID       string `json:"contributionid"`
	ContributionTitle    string `json:"contributiontitle"`
}
type ChildCreateContributionPost struct {
	ChildID             string `json:"childid"`
	ChildUserName       string `json:"childusername"`
	ChildProfilePicture string `json:"childprofilepicture"`
	ContributionID      string `json:"contributionid"`
	ContributionTitle   string `json:"contributiontitle"`
}
type AproveMentorPost struct {
	MentorID             string `json:"mentorid"`
	MentorUserName       string `json:"mentorusername"`
	MentorProfilePicture string `json:"mentorprofilepicture"`
}
type AproveMentorMsgPost struct {
	MentorID             string `json:"mentorid"`
	MentorUserName       string `json:"mentorusername"`
	MentorProfilePicture string `json:"mentorprofilepicture"`
}
type AdminAproveContributionPost struct {
	ContributionID    string `json:"contributionid"`
	ContributionTitle string `json:"contributiontitle"`
	ContributionType  string `json:"contributiontype"`
}
type AdminRejectContributionPost struct {
	ContributionID    string `json:"contributionid"`
	ContributionTitle string `json:"contributiontitle"`
	ContributionType  string `json:"contributiontype"`
}
type ChildSubmitMentorFormPost struct {
	ChildID             string `json:"childid"`
	ChildUserName       string `json:"childusername"`
	ChildProfilePicture string `json:"childprofilepicture"`
}
type MentoryHistorypostData struct {
	ID                        bson.ObjectId                  `json:"_id" bson:"_id,omitempty"`
	UserID                    string                         `json:"userid"`
	AdminMentorRequest        bool                           `json:"adminmentorrequestaprove"`
	AdminMentorRequestReject  bool                           `json:"adminmentorrequestreject"`
	ParentMentorRequest       bool                           `json:"parentmentorrequestaprove"`
	ParentMentorRequestReject bool                           `json:"parentmentorrequestreject"`
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

	if err != nil {
		newfollower := MentoryHistorypostData{}
		newfollower.UserID = userid
		newfollower.NewNotification = true
		item := FollowerPost{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture}
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
					item := FollowerGet{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture}
					newdata.AddItemGetFollow(item)
					newdata.NewNotification = true
					db.Update(result, newdata)
				}
			}
		} else {
			fmt.Println("no data available")
			newdata := MentoryHistorygetData{}
			newdata = result
			item := FollowerGet{UserFollowerID: followerid, UserFollowerName: followeruserinfo.FullName, UserFollowerProfilePicture: followeruserinfo.ProfilePicture}
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

	if err != nil {
		newlikes := MentoryHistorypostData{}
		newlikes.UserID = contributiondetail.UserID
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := LikesPost{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, LikeUserID: hexuserid, LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture}
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
		item := LikesGet{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, LikeUserID: hexuserid, LikeUserName: userinfo.FullName, LikeUserProfilePicture: userinfo.ProfilePicture}
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

	if err != nil {
		newlikes := MentoryHistorypostData{}
		newlikes.UserID = contributiondetail.UserID
		hexuserid := fmt.Sprintf("%x", string(userinfo.ID))
		hexcontributionid := fmt.Sprintf("%x", string(contributiondetail.ID))
		item := CommentsPost{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, CommentUserID: hexuserid, CommentUserName: userinfo.FullName, CommentUserProfilePicture: userinfo.ProfilePicture}
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
		item := CommentsGet{ContributionID: hexcontributionid, ContributionTitle: contributiondetail.Title, CommentUserID: hexuserid, CommentUserName: userinfo.FullName, CommentUserProfilePicture: userinfo.ProfilePicture}
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
	results := res{}

	u := new(GetUserData)
	if err = c.Bind(&u); err != nil {
	}
	res := GetUserData{}
	res = *u

	err = db.Find(bson.M{"userid": res.UserID}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Record found", 404, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

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
				item := MentorCreateContributionPost{MentorID: hexmentorid, MentorUserName: mentoruserinfo.FullName, MentorProfilePicture: mentoruserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle}
				createcontribution.AddItemPostCreateContribution(item)
				createcontribution.NewNotification = true
				db.Insert(createcontribution)
				fmt.Println("user not found add new data")
			} else {
				newdata := MentoryHistorygetData{}
				newdata = followerresult
				item := MentorCreateContributionGet{MentorID: hexmentorid, MentorUserName: mentoruserinfo.FullName, MentorProfilePicture: mentoruserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle}
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

			item := ChildCreateContributionPost{ChildID: childid, ChildUserName: kiduserinfo.FullName, ChildProfilePicture: kiduserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle}
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
			item := ChildCreateContributionGet{ChildID: childid, ChildUserName: kiduserinfo.FullName, ChildProfilePicture: kiduserinfo.ProfilePicture, ContributionID: contributionid, ContributionTitle: contributiontitle}
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

	if err != nil {
		mentoraprove := MentoryHistorypostData{}
		mentoraprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AproveMentorPost{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture}
		mentoraprove.AddItemPostAproveMentor(item)
		mentoraprove.NewNotification = true
		db.Insert(mentoraprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AproveMentorGet{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture}
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

	if err != nil {
		mentoraprove := MentoryHistorypostData{}
		mentoraprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AproveMentorMsgPost{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture}
		mentoraprove.AddItemPostAproveMentorMsg(item)
		mentoraprove.NewNotification = true
		db.Insert(mentoraprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AproveMentorMsgGet{MentorID: followerid, MentorUserName: followerinfo.FullName, MentorProfilePicture: followerinfo.ProfilePicture}
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AdminAproveContributionPost{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype}
		adminaprove.AddItemPostAdminAprove(item)
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AdminAproveContributionGet{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype}
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		item := AdminRejectContributionPost{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype}
		adminaprove.AddItemPostAdminReject(item)
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result

		item := AdminRejectContributionGet{ContributionID: contributionid, ContributionTitle: contributiontitle, ContributionType: contributiontype}
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		adminaprove.AdminMentorRequest = true
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		newdata.AdminMentorRequest = true
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		adminaprove.AdminMentorRequestReject = true
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		newdata.AdminMentorRequestReject = true
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		adminaprove.ParentMentorRequest = true
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		newdata.ParentMentorRequest = true
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

	if err != nil {
		adminaprove := MentoryHistorypostData{}
		adminaprove.UserID = Userid
		//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
		adminaprove.ParentMentorRequestReject = true
		adminaprove.NewNotification = true
		db.Insert(adminaprove)
		//fmt.Println(newfollower)
	} else {
		//fmt.Println("user exit update history")
		newdata := MentoryHistorygetData{}
		newdata = result
		newdata.ParentMentorRequestReject = true
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
	if userinfo.ParentEmail != "" {
		result := MentoryHistorygetData{}
		parentinfo := UserInfoByEmail(userinfo.ParentEmail)

		parentid := fmt.Sprintf("%x", string(parentinfo.ID))

		err = db.Find(bson.M{"userid": parentid}).One(&result)

		if err != nil {
			mentoraprove := MentoryHistorypostData{}
			mentoraprove.UserID = parentid
			//followerid := fmt.Sprintf("%x", string(followerinfo.ID))
			item := ChildSubmitMentorFormPost{ChildID: Userid, ChildUserName: userinfo.FullName, ChildProfilePicture: userinfo.ProfilePicture}
			mentoraprove.AddMentorFormRequestPost(item)
			mentoraprove.NewNotification = true
			db.Insert(mentoraprove)
			//fmt.Println(newfollower)
		} else {
			//fmt.Println("user exit update history")
			newdata := MentoryHistorygetData{}
			newdata = result

			item := ChildSubmitMentorFormGet{ChildID: Userid, ChildUserName: userinfo.FullName, ChildProfilePicture: userinfo.ProfilePicture}
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
