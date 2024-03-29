package contribution

import (
	"encoding/json"
	"favorites"
	"fmt"
	"net/http"
	"notification"
	"os"
	"shared"
	"strings"
	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type UserRequest struct {
	userRequest string `json:"userRequest"`
}

var response shared.Response

type UserDetail struct {
	UserID         string
	UserBio        string
	UserType       string
	CommentsCount  int
	LikesCount     int
	ProContributor bool
}

func ContributionGetAll(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.GetContributionres{}
	err = db.Find(bson.M{"contributionstatus": "Publish"}).Sort("-contributionpostdate").All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Error while finding record", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}

	userinfo := shared.UserinfoUpdategetData{}
	userDetail := []UserDetail{}

	for x := range results.Data {
		if len(userDetail) <= 0 {
			contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
			var contributionDetail favorites.GetFavrtData
			contributionDetail = ContributionFavrt(contributionid)
			UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)

			fmt.Println("database request")
			// fmt.Println("user id : ", UserIDconv)
			userinfo = notification.UserInfo(UserIDconv)
			// fmt.Println("user type : ", userinfo.UserType)
			results.Data[x].UserBio = userinfo.Bio
			results.Data[x].UserType = userinfo.UserType
			results.Data[x].CommentsCount = len(contributionDetail.Comments)
			if len(contributionDetail.Likes) > 0 {
				for aa := range contributionDetail.Likes {
					var userid shared.LikesUserID
					userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
					results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
				}
			}

			// fmt.Println("contribution id : ", contributionid)
			// fmt.Println("comment count : ", len(contributionDetail.Comments))
			// fmt.Println("like count : ", len(contributionDetail.Likes))

			var data UserDetail
			data.UserID = results.Data[x].UserID
			data.UserBio = userinfo.Bio
			data.UserType = userinfo.UserType
			data.LikesCount = len(contributionDetail.Likes)
			data.CommentsCount = len(contributionDetail.Comments)
			userDetail = append(userDetail, data)
		} else {
			for a := range userDetail {
				if userDetail[a].UserID == results.Data[x].UserID {
					results.Data[x].UserBio = userDetail[a].UserBio
					results.Data[x].UserType = userDetail[a].UserType

					contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
					var contributionDetail favorites.GetFavrtData
					contributionDetail = ContributionFavrt(contributionid)

					results.Data[x].CommentsCount = len(contributionDetail.Comments)
					// results.Data[x].LikesUserID = contributionDetail.Likes
					// fmt.Println("contribution id : ", contributionid)
					// fmt.Println("comment count : ", len(contributionDetail.Comments))
					// fmt.Println("like count : ", len(contributionDetail.Likes))
					if len(contributionDetail.Likes) > 0 {
						for aa := range contributionDetail.Likes {
							var userid shared.LikesUserID
							userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
							results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
						}
					}
					break
				}
			}
			if results.Data[x].UserType == "" {
				contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
				var contributionDetail favorites.GetFavrtData
				contributionDetail = ContributionFavrt(contributionid)
				UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
				fmt.Println("data request ")
				// fmt.Println("user id : ", UserIDconv)
				userinfo = notification.UserInfo(UserIDconv)
				// fmt.Println("user type : ", userinfo.UserType)
				results.Data[x].UserBio = userinfo.Bio
				results.Data[x].UserType = userinfo.UserType
				results.Data[x].CommentsCount = len(contributionDetail.Comments)
				// results.Data[x].LikesUserID = contributionDetail.Likes

				// fmt.Println("contribution id : ", contributionid)
				// fmt.Println("comment count : ", len(contributionDetail.Comments))
				// fmt.Println("like count : ", len(contributionDetail.Likes))
				if len(contributionDetail.Likes) > 0 {
					for aa := range contributionDetail.Likes {
						var userid shared.LikesUserID
						userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
						results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
					}
				}
				var data UserDetail
				data.UserID = results.Data[x].UserID
				data.UserBio = userinfo.Bio
				data.UserType = userinfo.UserType
				data.LikesCount = len(contributionDetail.Likes)
				data.CommentsCount = len(contributionDetail.Comments)
				userDetail = append(userDetail, data)
			}
		}
		// UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
		// fmt.Println("user id : ", UserIDconv)
		// userinfo = notification.UserInfo(UserIDconv)
		// fmt.Println("user type : ", userinfo.UserType)
		// results.Data[x].UserBio = userinfo.Bio
		// results.Data[x].UserType = userinfo.UserType

	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func MainContributionGetAll(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.GetContributionres{}
	err = db.Find(bson.M{"contributionstatus": "Publish", "adminstatus": 1}).Sort("-contributionpostdate").All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println("Get All Contribution")
	if err != nil {

	}
	if results.Data == nil {
		return c.JSON(http.StatusOK, 0)
	}

	userinfo := shared.UserinfoUpdategetData{}
	userDetail := []UserDetail{}

	for x := range results.Data {
		if len(userDetail) <= 0 {
			contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
			var contributionDetail favorites.GetFavrtData
			contributionDetail = ContributionFavrt(contributionid)
			UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
			pro := CheckProContributor(results.Data[x].UserEmail)

			fmt.Println("database request")
			// fmt.Println("user id : ", UserIDconv)
			userinfo = notification.UserInfo(UserIDconv)
			// fmt.Println("user type : ", userinfo.UserType)
			results.Data[x].UserBio = userinfo.Bio
			results.Data[x].UserType = userinfo.UserType
			results.Data[x].ProContributor = pro
			results.Data[x].CommentsCount = len(contributionDetail.Comments)
			if len(contributionDetail.Likes) > 0 {
				for aa := range contributionDetail.Likes {
					var userid shared.LikesUserID
					userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
					results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
				}
			}

			// fmt.Println("contribution id : ", contributionid)
			// fmt.Println("comment count : ", len(contributionDetail.Comments))
			// fmt.Println("like count : ", len(contributionDetail.Likes))

			var data UserDetail
			data.UserID = results.Data[x].UserID
			data.UserBio = userinfo.Bio
			data.UserType = userinfo.UserType
			data.LikesCount = len(contributionDetail.Likes)
			data.CommentsCount = len(contributionDetail.Comments)
			data.ProContributor = pro
			userDetail = append(userDetail, data)
		} else {
			for a := range userDetail {
				if userDetail[a].UserID == results.Data[x].UserID {
					results.Data[x].UserBio = userDetail[a].UserBio
					results.Data[x].UserType = userDetail[a].UserType
					results.Data[x].ProContributor = userDetail[a].ProContributor

					contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
					var contributionDetail favorites.GetFavrtData
					contributionDetail = ContributionFavrt(contributionid)

					results.Data[x].CommentsCount = len(contributionDetail.Comments)
					// results.Data[x].LikesUserID = contributionDetail.Likes
					// fmt.Println("contribution id : ", contributionid)
					// fmt.Println("comment count : ", len(contributionDetail.Comments))
					// fmt.Println("like count : ", len(contributionDetail.Likes))
					if len(contributionDetail.Likes) > 0 {
						for aa := range contributionDetail.Likes {
							var userid shared.LikesUserID
							userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
							results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
						}
					}
					break
				}
			}
			if results.Data[x].UserType == "" {
				contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
				var contributionDetail favorites.GetFavrtData
				contributionDetail = ContributionFavrt(contributionid)
				UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)

				pro := CheckProContributor(results.Data[x].UserEmail)

				fmt.Println("data request ")
				// fmt.Println("user id : ", UserIDconv)
				userinfo = notification.UserInfo(UserIDconv)
				// fmt.Println("user type : ", userinfo.UserType)
				results.Data[x].UserBio = userinfo.Bio
				results.Data[x].UserType = userinfo.UserType
				results.Data[x].CommentsCount = len(contributionDetail.Comments)
				results.Data[x].ProContributor = pro
				// results.Data[x].LikesUserID = contributionDetail.Likes

				// fmt.Println("contribution id : ", contributionid)
				// fmt.Println("comment count : ", len(contributionDetail.Comments))
				// fmt.Println("like count : ", len(contributionDetail.Likes))
				if len(contributionDetail.Likes) > 0 {
					for aa := range contributionDetail.Likes {
						var userid shared.LikesUserID
						userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
						results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
					}
				}
				var data UserDetail
				data.UserID = results.Data[x].UserID
				data.UserBio = userinfo.Bio
				data.UserType = userinfo.UserType
				data.LikesCount = len(contributionDetail.Likes)
				data.CommentsCount = len(contributionDetail.Comments)
				data.ProContributor = pro
				userDetail = append(userDetail, data)
			}
		}
		// UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
		// fmt.Println("user id : ", UserIDconv)
		// userinfo = notification.UserInfo(UserIDconv)
		// fmt.Println("user type : ", userinfo.UserType)
		// results.Data[x].UserBio = userinfo.Bio
		// results.Data[x].UserType = userinfo.UserType

	}

	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func ContributionFavrt(contributionid string) favorites.GetFavrtData {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.FAVORITESCOLLECTION)
	results := favorites.GetFavrtData{}

	if err != nil {
	}

	err = db.Find(bson.M{"contributionid": contributionid}).One(&results)

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

func GetAllRejectedContribution(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}
	err = db.Find(bson.M{"contributionstatus": "Reject"}).Sort("-contributionpostdate").All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println("Get All Rejected Contribution")
	if err != nil {

	}
	if results.Data == nil {
		return c.JSON(http.StatusOK, 0)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetAllEvent(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.GetContributionres{}
	err = db.Find(bson.M{"contributiontype": "event", "contributionstatus": "Publish"}).Sort("-contributionpostdate").All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println("Get All Event")
	if err != nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusOK, response)
	}

	userinfo := shared.UserinfoUpdategetData{}
	userDetail := []UserDetail{}

	for x := range results.Data {
		if len(userDetail) <= 0 {
			contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
			var contributionDetail favorites.GetFavrtData
			contributionDetail = ContributionFavrt(contributionid)
			UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
			pro := CheckProContributor(results.Data[x].UserEmail)
			fmt.Println("database request")
			// fmt.Println("user id : ", UserIDconv)
			userinfo = notification.UserInfo(UserIDconv)
			// fmt.Println("user type : ", userinfo.UserType)
			results.Data[x].UserBio = userinfo.Bio
			results.Data[x].UserType = userinfo.UserType
			results.Data[x].ProContributor = pro
			results.Data[x].CommentsCount = len(contributionDetail.Comments)
			if len(contributionDetail.Likes) > 0 {
				for aa := range contributionDetail.Likes {
					var userid shared.LikesUserID
					userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
					results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
				}
			}

			var data UserDetail
			data.UserID = results.Data[x].UserID
			data.UserBio = userinfo.Bio
			data.UserType = userinfo.UserType
			data.LikesCount = len(contributionDetail.Likes)
			data.CommentsCount = len(contributionDetail.Comments)
			data.ProContributor = pro
			userDetail = append(userDetail, data)
		} else {
			for a := range userDetail {
				if userDetail[a].UserID == results.Data[x].UserID {

					contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
					var contributionDetail favorites.GetFavrtData
					contributionDetail = ContributionFavrt(contributionid)

					results.Data[x].UserBio = userDetail[a].UserBio
					results.Data[x].UserType = userDetail[a].UserType
					results.Data[x].ProContributor = userDetail[a].ProContributor
					results.Data[x].CommentsCount = len(contributionDetail.Comments)
					if len(contributionDetail.Likes) > 0 {
						for aa := range contributionDetail.Likes {
							var userid shared.LikesUserID
							userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
							results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
						}
					}
					break
				}
			}
			if results.Data[x].UserType == "" {
				contributionid := fmt.Sprintf("%x", string(results.Data[x].ID))
				var contributionDetail favorites.GetFavrtData
				contributionDetail = ContributionFavrt(contributionid)
				UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
				pro := CheckProContributor(results.Data[x].UserEmail)
				fmt.Println("data request ")
				// fmt.Println("user id : ", UserIDconv)
				userinfo = notification.UserInfo(UserIDconv)
				// fmt.Println("user type : ", userinfo.UserType)
				results.Data[x].UserBio = userinfo.Bio
				results.Data[x].UserType = userinfo.UserType
				results.Data[x].ProContributor = pro
				results.Data[x].CommentsCount = len(contributionDetail.Comments)

				if len(contributionDetail.Likes) > 0 {
					for aa := range contributionDetail.Likes {
						var userid shared.LikesUserID
						userid.LikeUserID = contributionDetail.Likes[aa].LikeUserID
						results.Data[x].LikesUserID = append(results.Data[x].LikesUserID, userid)
					}
				}

				var data UserDetail
				data.UserID = results.Data[x].UserID
				data.UserBio = userinfo.Bio
				data.UserType = userinfo.UserType
				data.LikesCount = len(contributionDetail.Likes)
				data.CommentsCount = len(contributionDetail.Comments)
				data.ProContributor = pro
				userDetail = append(userDetail, data)
			}
		}
		// UserIDconv := bson.ObjectIdHex(results.Data[x].UserID)
		// fmt.Println("user id : ", UserIDconv)
		// userinfo = notification.UserInfo(UserIDconv)
		// fmt.Println("user type : ", userinfo.UserType)
		// results.Data[x].UserBio = userinfo.Bio
		// results.Data[x].UserType = userinfo.UserType

	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	response = shared.ReturnMessage(true, "Record Found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

//POST *********************************************************************************
func Addcontribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u
	audiopath := res.AudioPath
	converimage := res.Coverpage
	profilepicture := res.UserProfilePicture
	staticpath := shared.FILEBUCKETURL
	// for i := range res.Images {
	// 	res.Images[i].Imagestatus = staticpath + res.Images[i].Imagestatus
	// }
	if audiopath != "" {
		res.AudioPath = staticpath + audiopath
	} else {
		res.AudioPath = ""
	}

	if res.Coverpage != "" {
		if strings.Contains(converimage, staticpath) {
			res.UserProfilePicture = converimage
		} else {
			res.Coverpage = staticpath + converimage
		}
	}

	if profilepicture != "" {
		if strings.Contains(profilepicture, staticpath) {
			res.UserProfilePicture = profilepicture
		} else {
			res.UserProfilePicture = staticpath + profilepicture
		}
	} else {
		res.UserProfilePicture = ""
	}
	if res.Tags != nil {
		for x := range res.Tags {
			res.Tags[x].Tag = strings.ToLower(res.Tags[x].Tag)
		}
	}
	res.ViewCount = 0
	currentdate := time.Now().UTC()
	res.Likes = 0
	//date := currentdate.Format("2006-01-02 3:4:5 PM")
	res.ContributionPostDate = currentdate
	if res.UserFullName == "Cliiimb" {
		res.AdminStatus = 1
	}
	if res.ID == "" {
		db.Insert(res)
		response = shared.ReturnMessage(true, "Contribution Added", 200, "")
	} else if res.ID != "" {
		result := shared.ContributionPostData{}
		err = db.Find(bson.M{"_id": res.ID}).One(&result)
		newdata := shared.ContributionPostData{}
		newdata = result
		//		staticpath := shared.FILEBUCKETURL

		// newdata.AdminStatus = 0
		if res.UserFullName == "Cliiimb Article" {
			res.AdminStatus = 1
		}
		if len(res.Website) > 0 {
			newdata.Website = res.Website
		}
		if res.Coverpage != "" {
			newdata.Coverpage = res.Coverpage
		}
		if len(res.Tags) > 0 {
			newdata.Tags = res.Tags
		}
		if res.Title != "" {
			newdata.Title = res.Title
		}
		if res.MainCategory != "" {
			newdata.MainCategory = res.MainCategory
		}
		if res.SubCategories != "" {
			newdata.SubCategories = res.SubCategories
		}
		if res.Videos != "" {
			newdata.Videos = res.Videos
		}
		if len(res.Images) > 0 {
			newdata.Images = res.Images
		}
		if res.ContributionStatus != "" {
			newdata.ContributionStatus = res.ContributionStatus
		}

		if res.ContributionText != "" {
			newdata.ContributionText = res.ContributionText
		}
		db.Update(result, newdata)
		response = shared.ReturnMessage(true, "Contribution updated", 200, "")
	} else {
		response = shared.ReturnMessage(false, "Process Faild", 401, "")
	}
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func SearchContribution(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u

	email := res.UserEmail
	err = db.Find(bson.M{"useremail": email, "contributiontype": "contribution"}).Sort("-contributionpostdate").All(&results.Data)
	if err != nil {
		//log.Fatal(err)
	}
	buff, _ := json.Marshal(&results)
	json.Unmarshal(buff, &results)
	if results.Data == nil {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		defer session.Close()
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Record found", 200, results.Data)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func SearchEventByEmail(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Search event by email")
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.UserEmail

	fmt.Println(email)
	err = db.Find(bson.M{"useremail": email, "contributiontype": "event"}).Sort("-contributionpostdate").All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	var a [0]string
	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, &a)
	}
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func SearchContributionByCategory(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("search contribution by category")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	category := res.MainCategory

	err = db.Find(bson.M{"maincategory": category}).Sort("-contributionpostdate").All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}

func SearchContributionBySubCategory(c echo.Context) error {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("search contribution by sub category")
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	subcategory := res.SubCategories

	err = db.Find(bson.M{"subcategories": subcategory}).Sort("-contributionpostdate").All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}

func Editcontribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	result := shared.ContributionPostData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	newdata := shared.ContributionPostData{}
	newdata = result
	staticpath := shared.FILEBUCKETURL

	// newdata.AdminStatus = 0
	if len(res.Website) > 0 {
		newdata.Website = res.Website
	}

	if res.Coverpage != "" {
		a := strings.Contains(res.Coverpage, staticpath)
		if a == false {
			newdata.Coverpage = staticpath + res.Coverpage
		} else {
			newdata.Coverpage = result.Coverpage
		}

	}
	if len(res.Tags) > 0 {
		newdata.Tags = res.Tags
	}
	if res.Title != "" {
		newdata.Title = res.Title
	}
	if res.MainCategory != "" {
		newdata.MainCategory = res.MainCategory
	}
	if res.SubCategories != "" {
		newdata.SubCategories = res.SubCategories
	}
	if res.Videos != "" {
		newdata.Videos = res.Videos
	}
	if len(res.Images) > 0 {
		newdata.Images = res.Images
	}
	if res.ContributionStatus != "" {
		newdata.ContributionStatus = res.ContributionStatus
	}
	if res.ContributionText != "" {
		newdata.ContributionText = res.ContributionText
	}
	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func SearchContributionById(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u

	result := shared.ContributionData{}
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Record not found", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	response = shared.ReturnMessage(true, "Record found", 200, result)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func UpdateContributionStatus(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	result := shared.ContributionData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	res.ContributionStatus = ""
	newdata := shared.ContributionData{}
	newdata = result
	newdata.ContributionStatus = ""
	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func UpdateAdminStatus(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	result := shared.ContributionPostData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	// res.ContributionStatus = 1
	newdata := shared.ContributionPostData{}
	newdata = result
	newdata.AdminStatus = 1
	db.Update(result, newdata)
	notification.AddMentorCreatContributionHistory(result.UserID)
	notification.AddChildCreatContributionHistory(result.UserID)
	contributionid := fmt.Sprintf("%x", string(result.ID))
	notification.AddAdminAproveContributionHistory(result.UserID, contributionid, result.Title, result.ContributionType)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func AddView(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u

	result := shared.ContributionData{}
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	if result.UserID != "" {
		newdata := shared.ContributionData{}
		newdata = result
		session1, err1 := shared.ConnectMongo(shared.DBURL)
		db1 := session1.DB(shared.DBName).C(shared.VIEWCOLLECTION)
		if err1 != nil {
			fmt.Println("error:", err1)
		}

		result1 := shared.ViewgetData{}
		err = db1.Find(bson.M{"contributionid": res.ID.Hex(), "userid": res.UserID}).One(&result1)
		if result1.ContributionID == "" {

			view := newdata.ViewCount
			view++
			newdata.ViewCount = view
			db.Update(result, newdata)
			res1 := shared.ViewpostData{}

			res1.ContributionID = res.ID.Hex()
			res1.UserID = res.UserID
			db1.Insert(res1)
			response = shared.ReturnMessage(true, "", 200, "")
			defer session.Close()
			return c.JSON(http.StatusOK, response)

		}
	}
	response = shared.ReturnMessage(false, "View Not Added", 400, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func RemoveOneContribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u

	result := shared.ContributionData{}
	db.Find(bson.M{"_id": res.ID}).One(&result)

	err = db.Remove(result)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusNotFound, response)
	}
	response = shared.ReturnMessage(true, "Contribution removed", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}
func RemainingContributionCheck(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u
	err = db.Find(bson.M{"userid": res.UserID, "adminstatus": 1}).All(&results.Data)
	if err != nil {
		response = shared.ReturnMessage(false, "Record Not Found", 404, "")
		return c.JSON(http.StatusNotFound, response)
	}
	userContributionCount := GetMentorRequest(res.UserID)
	// if results.Data == nil {
	// 	response = shared.ReturnMessage(false, "Record Not Found", 404, userContributionCount)
	// 	defer session.Close()
	// 	return c.JSON(http.StatusOK, response)
	// }
	currentdate := time.Now().UTC()
	currentyear, currentmonth, _ := currentdate.Date()
	contributionCount := 0
	for x := range results.Data {
		contributiondate := results.Data[x].ContributionPostDate
		//t, _ := time.Parse("2006-01-02", contributiondate)
		contributionyear, contributionmonth, _ := contributiondate.Date()
		if currentmonth == contributionmonth && currentyear == contributionyear {
			contributionCount++
		}

	}
	remainingContribuiton := userContributionCount - contributionCount
	response = shared.ReturnMessage(true, "Contribution Count", 200, remainingContribuiton)
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}
func GetMentorRequest(userid string) int {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	result := shared.BMentorgetData{}
	//response := mentorRequestResponse{}

	err = db.Find(bson.M{"userid": userid}).One(&result)
	if err != nil {
		defer session.Close()
		return 0
		//results.Data = append(results.Data, kidrequest)
	}
	defer session.Close()
	return result.NumberOfContribution

}

func SearchEvent(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u
	result := shared.Contributionres{}
	//locationresult := shared.Contributionres{}
	//fmt.Println("%T \n", result)
	//alldata := shared.Contributionres{}
	query := bson.M{}
	if res.Date != "" {
		query["date"] = res.Date
	}
	if res.MainCategory != "" {
		query["maincategory"] = res.MainCategory
	}
	if res.Location != "" {
		query["location"] = bson.RegEx{Pattern: res.Location, Options: "i"}
	}
	query["contributiontype"] = "event"
	query["adminstatus"] = 1

	err = db.Find(query).Sort("-contributionpostdate").All(&result.Data)
	//err = db.Find(bson.M{"maincategory": res.MainCategory, "date": res.Date, "location": bson.RegEx{Pattern: res.Location, Options: "i"}, "contributiontype": "event"}).All(&result.Data)
	//db.Update(result, res)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)
}
func SearchSubContribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	data := shared.Contributionres{}
	resultsubcategory := shared.Contributionres{}
	resulttag := shared.Contributionres{}
	resultmaincategory := shared.Contributionres{}
	//fmt.Println("%T \n", result)
	//err = db.Find(bson.M{"$or": []bson.M{bson.M{"subcategories": bson.RegEx{"^.*" + res.SubCategories + "", "sim"}}, bson.M{"tags": bson.M{"tag": bson.RegEx{"^.*" + res.SubCategories + "", "sm"}}}, bson.M{"maincategory": res.SubCategories}}}).All(&resultsubcategory.Data)
	err = db.Find(bson.M{"subcategories": bson.RegEx{"^.*" + res.SubCategories + "", "im"}, "contributiontype": "contribution"}).Sort("-contributionpostdate").All(&resultsubcategory.Data)
	if resultsubcategory.Data != nil {
		for x := range resultsubcategory.Data {
			data.Data = append(data.Data, resultsubcategory.Data[x])
		}
	}

	err = db.Find(bson.M{"tags": bson.M{"tag": strings.ToLower(res.SubCategories)}, "contributiontype": "contribution"}).Sort("-contributionpostdate").All(&resulttag.Data)
	if resulttag.Data != nil {
		for x := range resulttag.Data {
			data.Data = append(data.Data, resulttag.Data[x])
		}
	}
	err = db.Find(bson.M{"maincategory": bson.RegEx{"^.*" + res.SubCategories + "", "i"}, "contributiontype": "contribution"}).Sort("-contributionpostdate").All(&resultmaincategory.Data)
	if resultmaincategory.Data != nil {
		for x := range resultmaincategory.Data {
			data.Data = append(data.Data, resultmaincategory.Data[x])
		}
	}

	defer session.Close()
	return c.JSON(http.StatusOK, &data)
}

func RejectContribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	if err != nil || session == nil {
		response = shared.ReturnMessage(false, "Server error", 501, "")
		return c.JSON(http.StatusOK, response)
	}
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	u := new(shared.ContributionPostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.ContributionPostData{}
	res = *u

	result := shared.ContributionData{}
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	if err != nil {
		response = shared.ReturnMessage(false, "Error finding data", 404, "")
		return c.JSON(http.StatusOK, response)
	}
	newdata := shared.ContributionData{}
	newdata = result
	newdata.ContributionStatus = "Reject"
	newdata.AdminStatus = 0
	db.Update(result, newdata)
	if err != nil {
		response = shared.ReturnMessage(false, "Error rejecting contrinution", 409, "")
		return c.JSON(http.StatusOK, response)
	}
	contributionid := fmt.Sprintf("%x", string(result.ID))
	notification.AddAdminRejectContributionHistory(result.UserID, contributionid, result.Title, result.ContributionType)
	response = shared.ReturnMessage(true, "Successfuly Rejected", 200, "")
	defer session.Close()
	return c.JSON(http.StatusOK, response)
}

func CheckProContributor(useremail string) bool {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	result := shared.BMentorgetData{}
	// response := mentorRequestResponse{}

	err = db.Find(bson.M{"useremail": useremail}).One(&result)
	if err != nil {
		fmt.Println("user mentor request form not submitted")
		defer session.Close()
		return false
	}
	var contributions shared.Contributionres
	contributions = UserContribution(useremail)
	if len(contributions.Data) >= 5 {
		defer session.Close()
		return true
	}

	defer session.Close()
	return false
}
func UserContribution(useremail string) shared.Contributionres {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}

	err = db.Find(bson.M{"useremail": useremail, "contributiontype": "contribution", "adminstatus": 1}).Sort("-contributionpostdate").All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}

	defer session.Close()
	return results
}
