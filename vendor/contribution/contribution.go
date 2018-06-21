package contribution

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"shared"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func ContributionGetAll(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)
	results := shared.Contributionres{}
	err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println(c)
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

	//fmt.Println(len(imagestatus))

	staticpath := shared.FILEBUCKETURL
	for i := range res.Images {
		res.Images[i].Imagestatus = staticpath + res.Images[i].Imagestatus
	}
	if audiopath != "" {
		res.AudioPath = staticpath + audiopath
	} else {
		res.AudioPath = ""
	}

	res.Coverpage = staticpath + converimage
	if profilepicture != "" {
		if strings.Contains(profilepicture, staticpath) {
			res.UserProfilePicture = profilepicture
		} else {
			res.UserProfilePicture = staticpath + profilepicture
		}

		//res.UserProfilePicture = staticpath + profilepicture
	} else {
		res.UserProfilePicture = ""
	}

	//res.Images[0].Imagestatus
	//fmt.Println(res)
	res.ViewCount = 0
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
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.ContributionRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.UserEmail
	fmt.Println("/n", email)
	//password:=c.FormValue("password")
	password := res.Title
	fmt.Println(password)

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)
	fmt.Println(email)
	err = db.Find(bson.M{"useremail": email}).All(&results.Data)

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
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

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
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

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
	res.ContributionStatus = 1
	newdata := shared.ContributionData{}
	newdata = result
	newdata.ContributionStatus = 1
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
	result := shared.ContributionData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"_id": res.ID}).One(&result)
	res.ContributionStatus = 1
	newdata := shared.ContributionData{}
	newdata = result
	newdata.AdminStatus = 1
	db.Update(result, newdata)
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
