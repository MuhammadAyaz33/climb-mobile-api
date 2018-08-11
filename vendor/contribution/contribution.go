package contribution

import (
	"encoding/json"
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

func ContributionGetAll(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}
	err = db.Find(bson.M{"contributiontype": "contribution", "contributionstatus": "Publish"}).All(&results.Data)

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
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetAllEvent(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}
	err = db.Find(bson.M{"contributiontype": "event", "contributionstatus": "Publish"}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println("Get All Event")
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

//POST *********************************************************************************
func Addcontribution(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	audiopath := res.AudioPath
	converimage := res.Coverpage
	profilepicture := res.UserProfilePicture
	staticpath := shared.FILEBUCKETURL
	for i := range res.Images {
		res.Images[i].Imagestatus = staticpath + res.Images[i].Imagestatus
	}
	if audiopath != "" {
		res.AudioPath = staticpath + audiopath
	} else {
		res.AudioPath = ""
	}
	if res.Coverpage != "" {
		res.Coverpage = staticpath + converimage
	} else {
		res.Coverpage = ""
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
	//date := currentdate.Format("2006-01-02 3:4:5 PM")
	res.ContributionPostDate = currentdate
	db.Insert(res)
	//fmt.Println(db)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

func SearchContribution(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	fmt.Println("Search contribution by email")
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
	err = db.Find(bson.M{"useremail": email, "contributiontype": "contribution"}).All(&results.Data)

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
func SearchEventByEmail(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
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
	err = db.Find(bson.M{"useremail": email, "contributiontype": "event"}).All(&results.Data)

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

	err = db.Find(bson.M{"maincategory": category}).All(&results.Data)

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

	err = db.Find(bson.M{"subcategories": subcategory}).All(&results.Data)

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
	db.Update(result, res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func SearchContributionById(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
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
	//db.Update(result, res)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)
}
func UpdateContributionStatus(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
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
	notification.AddAdminAproveContributionHistory(result.UserID, contributionid, result.Title)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}
func AddView(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
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
	//defer session.Close()
	if result.UserID != "" {
		//fmt.Println("contribution exist")
		newdata := shared.ContributionData{}
		newdata = result
		session1, err1 := shared.ConnectMongo(shared.DBURL)
		db1 := session1.DB(shared.DBName).C(shared.VIEWCOLLECTION)
		if err1 != nil {
			fmt.Println("error:", error)
		}

		result1 := shared.ViewgetData{}

		err = db1.Find(bson.M{"contributionid": res.ID.Hex(), "userid": res.UserID}).One(&result1)
		fmt.Println(result1)
		if result1.ContributionID == "" {

			view := newdata.ViewCount
			view++
			newdata.ViewCount = view
			db.Update(result, newdata)
			res1 := shared.ViewpostData{}

			res1.ContributionID = res.ID.Hex()
			res1.UserID = res.UserID
			db1.Insert(res1)
			defer session.Close()

			return c.JSON(http.StatusOK, "view added")
		}

	}

	return c.JSON(http.StatusOK, "no view add")
}
func RemoveOneContribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
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
	//fmt.Println(res.Data)
	fmt.Println(res)
	result := shared.ContributionData{}
	fmt.Println("%T \n", result)
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	db.Remove(result)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}
func RemainingContributionCheck(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
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

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)
	err = db.Find(bson.M{"userid": res.UserID, "adminstatus": 1}).All(&results.Data)

	// if err != nil {
	// 	//log.Fatal(err)
	// }
	userContributionCount := GetMentorRequest(res.UserID)

	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, &userContributionCount)
	}
	currentdate := time.Now().UTC()

	currentyear, currentmonth, _ := currentdate.Date()
	//fmt.Println(len(results.Data))
	contributionCount := 0
	for x := range results.Data {
		contributiondate := results.Data[x].ContributionPostDate
		//t, _ := time.Parse("2006-01-02", contributiondate)
		contributionyear, contributionmonth, _ := contributiondate.Date()
		if currentmonth == contributionmonth && currentyear == contributionyear {
			contributionCount++
		}

	}
	fmt.Println("total contribution count : ", contributionCount)

	fmt.Println("user contribution count: ", userContributionCount)

	remainingContribuiton := userContributionCount - contributionCount
	fmt.Println("remaing contribution : ", remainingContribuiton)

	defer session.Close()
	return c.JSON(http.StatusOK, &remainingContribuiton)
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
	result := shared.Contributionres{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"maincategory": res.MainCategory, "date": res.Date, "location": bson.RegEx{Pattern: res.Location, Options: "i"}, "contributiontype": "event"}).All(&result.Data)
	//db.Update(result, res)
	defer session.Close()
	return c.JSON(http.StatusOK, &result)
}
func SearchSubContribution(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
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
	err = db.Find(bson.M{"subcategories": bson.RegEx{"^.*" + res.SubCategories + "", "im"}, "contributiontype": "contribution"}).All(&resultsubcategory.Data)
	if resultsubcategory.Data != nil {
		for x := range resultsubcategory.Data {
			data.Data = append(data.Data, resultsubcategory.Data[x])
		}
	}

	err = db.Find(bson.M{"tags": bson.M{"tag": strings.ToLower(res.SubCategories)}, "contributiontype": "contribution"}).All(&resulttag.Data)
	if resulttag.Data != nil {
		for x := range resulttag.Data {
			data.Data = append(data.Data, resulttag.Data[x])
		}
	}
	err = db.Find(bson.M{"maincategory": res.SubCategories, "contributiontype": "contribution"}).All(&resultmaincategory.Data)
	if resultmaincategory.Data != nil {
		for x := range resultmaincategory.Data {
			data.Data = append(data.Data, resultmaincategory.Data[x])
		}
	}

	defer session.Close()
	return c.JSON(http.StatusOK, &data)
}
