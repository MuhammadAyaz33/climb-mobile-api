package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"shared"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	gomail "gopkg.in/gomail.v2"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//	PARENT PROFILE DATA ****************

type GetKid struct {
	KidID string
}
type ParentgetData struct {
	ID          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ParentEmail string
	Token       string
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
	Token       string        `json:"token"`
	Kids        []PostKid     `json:"kids"`
}
type ParentRes struct {
	Data []ParentpostData `json:"Data"`
}

//GET *********************************************************************************
func GetAll(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.Userres{}
	err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)
	fmt.Println(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
	buff, _ := json.Marshal(&results)
	fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}

//POST *********************************************************************************
func Adduser(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	// name := c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	//fmt.Println(name)
	//u:=new (postData)
	u := new(shared.UserUpdateData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserUpdateData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	results := shared.Userres{}
	res.Email = strings.ToLower(res.Email)
	err = db.Find(bson.M{"email": res.Email}).All(&results.Data)

	if results.Data == nil {

		var maintoken string
		maintoken = sendemail(res.Email, "adduser", "user")
		//fmt.Println("this is token /n")
		//fmt.Println(maintoken)
		VerificationTokenSave(res.Email, maintoken, "adduser")
		res.Status = 0
		//parent := res.ParentStatus
		res.ParentStatus = 0
		res.Age = 0
		res.ParentPhone = 0
		hash := hashAndSalt([]byte(res.Password))
		res.Password = hash
		res.UserType = "user"
		//res.MentorStatus = 0
		db.Insert(res)

	} else {
		//fmt.Println("user available try to login")

		var resultEmailStatus = shared.ErrorCheckStatus{
			Status: "1",
		}
		defer session.Close()
		return c.JSON(http.StatusOK, resultEmailStatus)

	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, &r)

}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func AddAdmin(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	//fmt.Println(name)
	//u:=new (postData)
	u := new(shared.UserUpdateData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserUpdateData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	results := shared.Userres{}
	err = db.Find(bson.M{"email": res.Email}).All(&results.Data)
	response := Response{}
	if results.Data == nil {

		var mySigningKey = []byte(res.Email)
		token := jwt.New(jwt.SigningMethodHS256)
		maintoken, _ := token.SignedString(mySigningKey)

		// var maintoken string
		//maintoken = sendemail(res.Email, "adduser", "user")
		//fmt.Println("this is token /n")
		//fmt.Println(maintoken)
		VerificationTokenSave(res.Email, maintoken, "adduser")
		res.Status = 1
		//parent := res.ParentStatus
		res.ParentStatus = 1
		res.Age = 21
		res.ParentPhone = 0
		hash := hashAndSalt([]byte(res.Password))
		res.Password = hash
		res.UserType = "admin"
		res.FullName = "Cliiimb"
		res.Bio = "News Room"
		res.ProfilePicture = "https://s3.us-east-2.amazonaws.com/climbmentors/2bad8322b2af2cb0513a503b0346d881d7737337.jpg"
		//res.MentorStatus = 0
		db.Insert(res)
		response.Message = "Admin Added"
		response.Status = true

	} else {
		//fmt.Println("user available try to login")
		response.Message = "User Already Exits"
		response.Status = false
		defer session.Close()
		return c.JSON(http.StatusOK, response)

	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, response)

}

//encrypt password
func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// decrypt password
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func PasswordVerification(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	fmt.Println(res)
	results := shared.Userres{}
	err = db.Find(bson.M{"email": res.Email}).All(&results.Data)

	if results.Data != nil {

		fmt.Println("/n")
		fmt.Println("user not available send email")
		var maintoken string
		maintoken = sendemail(res.Email, "password", "user")
		//fmt.Println("this is token /n")
		//fmt.Println(maintoken)
		VerificationTokenSave(res.Email, maintoken, "password")
		//res.Status = "unverified"
		//db.Insert(res)

	} else {
		fmt.Println("user available try to login")
		defer session.Close()
		return c.JSON(http.StatusOK, "user not available")
	}
	//db.Insert(res)
	defer session.Close()
	return c.JSON(http.StatusOK, "email send")

}

func VerificationTokenSave(email string, token string, check string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	var db *mgo.Collection
	if check == "adduser" {
		db = session.DB(shared.DBName).C(shared.VERIFICATIONCOLLECTION)
		res := shared.VerificationpostData{}
		res.Date = time.Now()
		res.EmailID = email
		res.Token = token
		//fmt.Println(vres)
		db.Insert(res)
		defer session.Close()
	} else if check == "password" {
		fmt.Println("password update data save")
		db = session.DB(shared.DBName).C(shared.PASSWORDVERIFICATIONCOLLECTION)
		res := shared.PasswordVerificationpostData{}
		res.Date = time.Now()
		res.EmailID = email
		res.Token = token
		//fmt.Println(vres)
		db.Insert(res)
		defer session.Close()
	}

	if err != nil {
	}

}

func RegistrationVerfication(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	var db *mgo.Collection
	db = session.DB(shared.DBName).C(shared.VERIFICATIONCOLLECTION)
	fmt.Println("/n")
	fmt.Println(reflect.TypeOf(db))
	results := shared.Verificationres{}

	u := new(shared.VerificationpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.VerificationpostData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//token := c.FormValue("token")
	token := res.Token

	err = db.Find(bson.M{"token": token}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}

	t := time.Now()
	fmt.Println("this is time /n")
	fmt.Println(t)

	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "invalid input")

	} else {
		UpdateStatus(1, results.Data[0].EmailID, "registration")
		defer session.Close()

		//open.Start("https://google.com")
		return c.JSON(http.StatusOK, 1)
	}
	//	buff, _ := json.Marshal(&results)

	//	json.Unmarshal(buff, &results)
	//	return c.JSON(http.StatusOK, &results)

}

func PasswordResetVerfication(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	var db *mgo.Collection
	db = session.DB(shared.DBName).C(shared.PASSWORDVERIFICATIONCOLLECTION)

	results := shared.PasswordVerificationres{}

	u := new(shared.PasswordVerificationpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.PasswordVerificationpostData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.PasswordVerificationRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//token := c.FormValue("token")

	err = db.Find(bson.M{"token": res.Token}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}

	//t := time.Now()
	fmt.Println("this is time /n")
	fmt.Println(results)

	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "invalid input")

	} else {
		defer session.Close()
		//UpdateStatus("verified", results.Data[0].EmailID)
		return c.JSON(http.StatusOK, "you can change your password now enter new password")
	}

}
func ParentVerfication(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	var db *mgo.Collection
	db = session.DB(shared.DBName).C(shared.PARENTCOLLECTION)

	results := Parentres{}

	u := new(shared.VerificationpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.VerificationpostData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	// os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//token := c.FormValue("token")
	//token := res.Token

	// token := c.FormValue("token")
	// useremail := c.FormValue("useremail")
	token := res.Token
	useremail := res.EmailID
	fmt.Println(useremail)

	err = db.Find(bson.M{"token": token, "kids": bson.M{"kidid": useremail}}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}

	t := time.Now()
	fmt.Println("this is time /n")
	fmt.Println(t)

	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "Invalid Link")

	} else {
		defer session.Close()
		UpdateParentStatus(1, useremail, "parent")
		return c.JSON(http.StatusOK, "you are successfully verify your childer now you can see your childern activity")
	}

}

func UpdateStatus(status int, email string, check string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	if err != nil {
	}

	result := shared.UserinfoUpdategetData{}
	err = db.Find(bson.M{"email": email}).One(&result)
	newresult := shared.UserinfoUpdategetData{}
	newresult = result
	newresult.Status = status
	db.Update(result, newresult)
	defer session.Close()
	//err = db.Find(bson.M{"email": email}).One(&result)

}
func UpdateParentStatus(status int, email string, check string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	if err != nil {
	}
	fmt.Println(email)
	result := shared.UserUpdateData{}
	err = db.Find(bson.M{"email": email}).One(&result)
	fmt.Println("helloloooooooo")
	fmt.Println(result)
	newresult := shared.UserUpdateData{}
	newresult = result
	newresult.ParentStatus = 1

	db.Update(result, newresult)
	defer session.Close()

	//err = db.Find(bson.M{"email": email}).One(&result)

}
func Login(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.Userres{}

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("login ")
	//	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	email := strings.ToLower(res.Email)

	password := res.Password

	err = db.Find(bson.M{"email": email}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	mentorstatus := GetMentorRequest(email)
	fmt.Println("status changed : ", mentorstatus)
	if results.Data == nil {
		var Status = shared.ErrorCheckStatus{
			Status: "0",
		}
		defer session.Close()
		return c.JSON(http.StatusOK, Status)
	} else {
		hash := results.Data[0].Password
		check := comparePasswords(hash, []byte(password))
		if check == true {
			results.Data[0].MentorStatus = mentorstatus
			buff, _ := json.Marshal(&results)
			//fmt.Println(string(buff))
			defer session.Close()
			json.Unmarshal(buff, &results)
			return c.JSON(http.StatusOK, &results)
		} else {
			var Status = shared.ErrorCheckStatus{
				Status: "0",
			}
			defer session.Close()
			return c.JSON(http.StatusOK, Status)
		}
	}
	return c.JSON(http.StatusOK, err)
}
func GetMentorRequest(useremail string) int {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MENTORREQUESTCOLLECTION)

	result := shared.BMentorgetData{}
	//response := mentorRequestResponse{}

	err = db.Find(bson.M{"useremail": useremail}).One(&result)
	if err != nil {
		defer session.Close()
		return 0
		//results.Data = append(results.Data, kidrequest)
	}

	if result.AdminStatus == 1 && result.ParentStatus == 1 {

		defer session.Close()
		return 1
	} else {
		defer session.Close()
		return 0
	}

}
func EditProfile(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
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
	//fmt.Println(res)
	//fmt.Println(res.Data)
	fmt.Println(res)
	result := shared.UsergetData{}
	fmt.Println("%T \n", result)
	err = db.Find(bson.M{"email": res.Email}).One(&result)
	newdata := shared.UsergetData{}
	newdata = result
	if res.FullName != "" {
		newdata.FullName = res.FullName
	}
	if res.CompanyName != "" {
		newdata.CompanyName = res.CompanyName
	}
	if res.Password != "" {
		newdata.Password = res.Password
	}

	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &newdata)
}
func UpdateProfile(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)

	u := new(shared.UserUpdateData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserUpdateData{}
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r shared.UserRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res)
	result := shared.UserUpdateData{}

	err = db.Find(bson.M{"email": res.Email}).One(&result)

	newdata := shared.UserUpdateData{}
	newdata = result
	if res.Password != "" {
		hash := hashAndSalt([]byte(res.Password))
		//res.Password = hash
		newdata.Password = hash
	}
	if res.FullName != "" {
		newdata.FullName = res.FullName
	}

	//fmt.Println(newdata.FullName)
	if res.CompanyName != "" {
		newdata.CompanyName = res.CompanyName
	}
	if res.Address != "" {
		newdata.Address = res.Address
	}
	if res.City != "" {
		newdata.City = res.City
	}
	if res.ZipCode != 0 {
		newdata.ZipCode = res.ZipCode
	}
	if res.Bio != "" {
		newdata.Bio = res.Bio
	}

	if res.Age < 18 {
		newdata.ParentStatus = 0
	} else {
		newdata.ParentStatus = 1
	}
	newdata.Age = res.Age

	if res.ParentPhone != 0 {
		newdata.ParentPhone = res.ParentPhone

	}

	if res.ProfilePicture != "" {
		staticpath := shared.FILEBUCKETURL
		var profilepic string
		if strings.Contains(res.ProfilePicture, staticpath) {
			newdata.ProfilePicture = res.ProfilePicture
			profilepic = res.ProfilePicture
		} else {
			newdata.ProfilePicture = staticpath + res.ProfilePicture
			profilepic = staticpath + res.ProfilePicture
		}
		UpdateContributionProfilePic(profilepic, res.Email)

	}
	if res.ParentEmail != "" {
		if result.ParentEmail == "" {
			newdata.ParentEmail = res.ParentEmail
			var maintoken string
			maintoken = sendemail(res.ParentEmail, "parent", res.Email)

			ParentVerificationTokenSave(res.ParentEmail, maintoken, res.Email)
			//fmt.Println(maintoken)
		}
		if res.ParentEmail == result.ParentEmail {
			newdata.ParentEmail = res.ParentEmail
		} else {
			newdata.ParentEmail = res.ParentEmail
			var maintoken string
			maintoken = sendemail(res.ParentEmail, "parent", res.Email)

			ParentVerificationTokenSave(res.ParentEmail, maintoken, res.Email)
		}

	}
	if res.ParentPhone <= 0 {
		newdata.ParentPhone = res.ParentPhone
	}
	//newdata.ParentPhone = res.ParentPhone
	// fmt.Println("result ---------", newdata)
	// fmt.Println("result ---------", result)
	//err = db.Find(bson.M{"email": res.Email}).One(&result)
	db.Update(result, newdata)
	//fmt.Println("result ---------", err1)
	result1 := shared.UsergetData{}
	err = db.Find(bson.M{"email": res.Email}).One(&result1)
	mentorstatus := GetMentorRequest(res.Email)
	result1.MentorStatus = mentorstatus
	defer session.Close()
	return c.JSON(http.StatusOK, result1)
}
func UpdateContributionProfilePic(picture string, email string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CONTRIBUTIONCOLLECTION)

	result := shared.Contributionres{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"useremail": email}).All(&result.Data)
	if result.Data == nil {

	} else {
		//fmt.Println("contribution profile pic update /n")
		dataset := shared.ContributionData{}
		newdata := shared.ContributionData{}

		for x := range result.Data {
			//	fmt.Println(x)
			err = db.Find(bson.M{"_id": result.Data[x].ID}).One(&dataset)
			newdata = dataset
			newdata.UserProfilePicture = picture
			db.Update(dataset, newdata)
		}

	}
	// newdata := shared.UsergetData{}
	// newdata = result
	// newdata.AboutMe = res.AboutMe
	// db.Update(result, newdata)
	defer session.Close()
	if err != nil {
	}
	//return c.JSON(http.StatusOK, &newdata)

}
func ParentVerificationTokenSave(parentemail string, token string, email string) {
	session, err := shared.ConnectMongo(shared.DBURL)
	var db *mgo.Collection

	db = session.DB(shared.DBName).C(shared.PARENTCOLLECTION)

	res := ParentpostData{}
	res.ParentEmail = parentemail
	res.Token = token
	item := PostKid{KidID: email}
	res.AddItemPost(item)
	//fmt.Println(vres)
	result := ParentgetData{}
	err = db.Find(bson.M{"parentemail": res.ParentEmail}).One(&result)
	if result.ParentEmail == "" {
		//fmt.Println("ni match hova add kr do")
		db.Insert(res)
		defer session.Close()
	} else {
		//fmt.Println("match ho geya hai update kro")

		newdata := ParentgetData{}
		newdata = result
		newdata.Token = token
		a := res.Kids[0].KidID

		item1 := GetKid{KidID: a}

		newdata.AddItem(item1)
		db.Update(result, newdata)
		defer session.Close()
	}
	//db.Insert(res)

	if err != nil {
	}
}
func (box *ParentgetData) AddItem(item GetKid) []GetKid {
	box.Kids = append(box.Kids, item)
	return box.Kids
}
func (box *ParentpostData) AddItemPost(item PostKid) []PostKid {
	box.Kids = append(box.Kids, item)
	return box.Kids
}
func sendemail(email string, check string, useremail string) (s string) {

	var mySigningKey = []byte(email)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, error := token.SignedString(mySigningKey)
	if error != nil {
		fmt.Println(error)
	}
	//fmt.Println(tokenString)
	var sendmessage string
	var emailsubject string
	if check == "adduser" {
		emailsubject = "Cliiimb Registration"
		// sendmessage = fmt.Sprintf("Hello <b>testing</b> </br> click here to verify the email <a href='http://18.216.55.104:4200/email-verified?token=%s&useremail=%s'>Click here</a>", tokenString, email)
		cleanstring := `'<DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml">
		   <head>
			  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			  <title>Cliiimb</title>
			  
			  <style type="text/css">
				 /* Client-specific Styles */
				 #outlook a {padding:0;} /* Force Outlook to provide a "view in browser" menu link. */
				 body{width:100%% ; -webkit-text-size-adjust:100%%; -ms-text-size-adjust:100%%; margin:0; padding:0;background: #F5F5F4;}
				 /* Prevent Webkit and Windows Mobile platforms from changing default font sizes, while not breaking desktop design. */
				 .ExternalClass {width:100%%;} /* Force Hotmail to display emails at full width */
				 .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div {line-height: 100%%;} /* Force Hotmail to display normal line spacing.*/
				 #backgroundTable {margin:0; padding:0; width:100%% ; line-height: 100%% ;background: #F5F5F4;}
				 img {outline:none; text-decoration:none;border:none; -ms-interpolation-mode: bicubic;}
				 a img {border:none;}
				 .image_fix {display:block;}
				 p {margin: 0px 0px ;}
				 table td {border-collapse: collapse;}
				 table { border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt; }
				 a {color: #0a8cce;text-decoration: none;text-decoration:none;}
				 /*STYLES*/
				 table[class=full] { width: 100%%; clear: both; }
				 /*IPAD STYLES*/
				 @media only screen and (max-width: 640px) {
				 a[href^="tel"], a[href^="sms"] {
				 text-decoration: none;
				 color: #0a8cce; /* or whatever your want */
				 pointer-events: none;
				 cursor: default;
				 }
				 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
				 text-decoration: default;
				 color: #0a8cce ;
				 pointer-events: auto;
				 cursor: default;
				 }
				 table[class=devicewidth] {width: 440px;text-align:center;}
				 table[class=devicewidthinner] {width: 420px;text-align:center;}
				 img[class=banner] {width: 440px;height:220px;}
				 img[class=colimg2] {width: 440px;height:220px;}
				 
				 
				 }
				 /*IPHONE STYLES*/
				 @media only screen and (max-width: 480px) {
				 a[href^="tel"], a[href^="sms"] {
				 text-decoration: none;
				 color: #0a8cce; /* or whatever your want */
				 pointer-events: none;
				 cursor: default;
				 }
				 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
				 text-decoration: default;
				 color: #0a8cce ; 
				 pointer-events: auto;
				 cursor: default;
				 }
				 table[class=devicewidth] {width: 280px;text-align:center;}
				 table[class=devicewidthinner] {width: 260px;text-align:center;}
				 img[class=banner] {width: 280px;height:140px;}
				 img[class=colimg2] {width: 280px;height:140px;}
				 td[class=mobile-hide]{display:none;}
				 td[class="padding-bottom25"]{padding-bottom:25px;}
				
				 }
			  </style>
		   </head>
		   <body bgcolor="#F5F5F4" style="background: #F5F5F4 ;"><br /><br />
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td align="center" height="20" style="font-size:1px; line-height:1px;">&nbsp;</td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
			 
	
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="header">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
								   
									  <tr>
										 <td>
											
											<table width="200" align="" border="0" cellpadding="0" cellspacing="0" class="devicewidth">
											   <tbody>
												  <tr>
													 <td width="200" height="45" align="">
														<div class="imgpop">
														   <a target="_blank" href="#">
														   <img src="https://s3.us-east-2.amazonaws.com/climbmentors/logo.svg" alt="" border="0" width="" height="54" style="display:block; border:none; outline:none; text-decoration:none;">
														   </a>
														</div>
													 </td>
												  </tr>
											   </tbody>
											</table>
										
										 </td>
									  </tr>
									  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
									  
									  <tr>
										 <td>
											<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
											   <tbody>
											
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 30px; color: #0F1A59; text-align:left; line-height: 40px;" st-title="fulltext-heading">
														Thanks for joining Cliiimb
													 </td>
												  </tr>
												
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #0F1A59; text-align:left; line-height: 25px;
			font-weight: bold;" st-content="fulltext-content">
														Please confirm that your email address is correct to continue. Click the link below to get started. 
													 </td>
												  </tr>
											
											   </tbody>
											</table>
										 </td>
									  </tr>
									 
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable">
		   <tbody>
			  <tr>
				 <td>
				   
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="2columns">
		   <tbody>
			  <tr>
				 <td>
				   
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
									 
									  <tr>
										 <td>
											<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
											   <tbody>
												 
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #fff; text-align:left; line-height: 10px;" st-title="fulltext-title">
														<a href="https://cliiimb.com/email-verified?token=%s&useremail=%s" style="padding: 16px 30px;background:#ec1c24;border-radius: 5px;font-weight:600;border:1px solid #ec1c24;color: #fff ;">Confirm Email Address</a>
													 </td>
												  </tr>
												
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #666666; text-align:center; line-height: 30px;" st-content="fulltext-content">
														
													 </td>
												  </tr>
												 
											   </tbody>
											</table>
										 </td>
									  </tr>
									  
									  <tr>
										 <td ></td>
									  </tr>
									  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
					   <tbody>
						  
						  <tr>
							 <td width="550" align="center" height="1" bgcolor="#d1d1d1" style="font-size:1px; line-height:1px;">&nbsp;</td>
						  </tr>
					 
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="postfooter" >
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="left" class="devicewidth" style="background: #fff;">
								   <tbody>
									  <tr>
										 <td align="left" valign="middle" style="font-family: Helvetica, arial, sans-serif; font-weight:300;font-size: 16px;color: #252C65" st-content="postfooter">
											Visit <a href="#" style="text-decoration: none; color: #ec1c24;font-weight:300;"> Cliiimb Site</a> 
										 </td>
									  </tr>
								  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		   
		   </body>
		   </html>'`
		sendmessage = fmt.Sprintf(trimQuotes(cleanstring), tokenString, email)

	} else if check == "password" {
		emailsubject = "Cliiimb Password Reset"
		//sendmessage = fmt.Sprintf("Hello <b>testing</b> </br> click here to change password <a href='http://18.216.55.104:4200/reset-password?token=%s&useremail=%s'>Click here</a>", tokenString, email)

		newstring := `'<DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml">
		   <head>
			  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
			  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			  <title>Cliiimb</title>
			  
			  <style type="text/css">
				 /* Client-specific Styles */
				 #outlook a {padding:0;} /* Force Outlook to provide a "view in browser" menu link. */
				 body{width:100%% ; -webkit-text-size-adjust:100%%; -ms-text-size-adjust:100%%; margin:0; padding:0;background: #F5F5F4;}
				 /* Prevent Webkit and Windows Mobile platforms from changing default font sizes, while not breaking desktop design. */
				 .ExternalClass {width:100%%;} /* Force Hotmail to display emails at full width */
				 .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div {line-height: 100%%;} /* Force Hotmail to display normal line spacing.*/
				 #backgroundTable {margin:0; padding:0; width:100%% ; line-height: 100%% ;background: #F5F5F4;}
				 img {outline:none; text-decoration:none;border:none; -ms-interpolation-mode: bicubic;}
				 a img {border:none;}
				 .image_fix {display:block;}
				 p {margin: 0px 0px ;}
				 table td {border-collapse: collapse;}
				 table { border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt; }
				 a {color: #0a8cce;text-decoration: none;text-decoration:none;}
				 /*STYLES*/
				 table[class=full] { width: 100%%; clear: both; }
				 /*IPAD STYLES*/
				 @media only screen and (max-width: 640px) {
				 a[href^="tel"], a[href^="sms"] {
				 text-decoration: none;
				 color: #0a8cce; /* or whatever your want */
				 pointer-events: none;
				 cursor: default;
				 }
				 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
				 text-decoration: default;
				 color: #0a8cce ;
				 pointer-events: auto;
				 cursor: default;
				 }
				 table[class=devicewidth] {width: 440px;text-align:center;}
				 table[class=devicewidthinner] {width: 420px;text-align:center;}
				 img[class=banner] {width: 440px;height:220px;}
				 img[class=colimg2] {width: 440px;height:220px;}
				 
				 
				 }
				 /*IPHONE STYLES*/
				 @media only screen and (max-width: 480px) {
				 a[href^="tel"], a[href^="sms"] {
				 text-decoration: none;
				 color: #0a8cce; /* or whatever your want */
				 pointer-events: none;
				 cursor: default;
				 }
				 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
				 text-decoration: default;
				 color: #0a8cce ; 
				 pointer-events: auto;
				 cursor: default;
				 }
				 table[class=devicewidth] {width: 280px;text-align:center;}
				 table[class=devicewidthinner] {width: 260px;text-align:center;}
				 img[class=banner] {width: 280px;height:140px;}
				 img[class=colimg2] {width: 280px;height:140px;}
				 td[class=mobile-hide]{display:none;}
				 td[class="padding-bottom25"]{padding-bottom:25px;}
				
				 }
			  </style>
		   </head>
		   <body bgcolor="#F5F5F4" style="background: #F5F5F4 ;"><br /><br />
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td align="center" height="20" style="font-size:1px; line-height:1px;">&nbsp;</td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
			 
	
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="header">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
								   
									  <tr>
										 <td>
											
											<table width="200" align="" border="0" cellpadding="0" cellspacing="0" class="devicewidth">
											   <tbody>
												  <tr>
													 <td width="200" height="45" align="">
														<div class="imgpop">
														   <a target="_blank" href="#">
														   <img src="https://s3.us-east-2.amazonaws.com/climbmentors/logo.svg" alt="" border="0" width="" height="54" style="display:block; border:none; outline:none; text-decoration:none;">
														   </a>
														</div>
													 </td>
												  </tr>
											   </tbody>
											</table>
										
										 </td>
									  </tr>
									  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
									  
									  <tr>
										 <td>
											<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
											   <tbody>
											
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 30px; color: #0F1A59; text-align:left; line-height: 40px;" st-title="fulltext-heading">
														Thanks for joining Cliiimb
													 </td>
												  </tr>
												
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #0F1A59; text-align:left; line-height: 25px;
			font-weight: bold;" st-content="fulltext-content">
			You recently requested to reset your password for your Cliiimb Account. Use the button to reset it.
													 </td>
												  </tr>
											
											   </tbody>
											</table>
										 </td>
									  </tr>
									 
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable">
		   <tbody>
			  <tr>
				 <td>
				   
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="2columns">
		   <tbody>
			  <tr>
				 <td>
				   
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
								   <tbody>
									 
									  <tr>
										 <td>
											<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
											   <tbody>
												 
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #fff; text-align:left; line-height: 10px;" st-title="fulltext-title">
														<a href="https://cliiimb.com/reset-password?token=%s&useremail=%s" style="padding: 16px 30px;background:#ec1c24;border-radius: 5px;font-weight:600;border:1px solid #ec1c24;color: #fff ;">Reset Your Password</a>
													 </td>
												  </tr>
												
												  <tr>
													 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #666666; text-align:center; line-height: 30px;" st-content="fulltext-content">
														
													 </td>
												  </tr>
												 
											   </tbody>
											</table>
										 </td>
									  </tr>
									  
									  <tr>
										 <td ></td>
									  </tr>
									  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
			  <tr>
				 <td>
					<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
					   <tbody>
						  
						  <tr>
							 <td width="550" align="center" height="1" bgcolor="#d1d1d1" style="font-size:1px; line-height:1px;">&nbsp;</td>
						  </tr>
					 
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="postfooter" >
		   <tbody>
			  <tr>
				 <td>
					<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
					   <tbody>
						  <tr>
							 <td width="100%%">
								<table width="600" cellpadding="20" cellspacing="0" border="0" align="left" class="devicewidth" style="background: #fff;">
								   <tbody>
									  <tr>
										 <td align="left" valign="middle" style="font-family: Helvetica, arial, sans-serif; font-weight:300;font-size: 16px;color: #252C65" st-content="postfooter">
											Visit <a href="#" style="text-decoration: none; color: #ec1c24;font-weight:300;"> Cliiimb Site</a> 
										 </td>
									  </tr>
								  
								   </tbody>
								</table>
							 </td>
						  </tr>
					   </tbody>
					</table>
				 </td>
			  </tr>
		   </tbody>
		</table>
	
		   
		   </body>
		   </html>'`
		sendmessage = fmt.Sprintf(trimQuotes(newstring), tokenString, email)
		//fmt.Println(sendmessage)

	} else if check == "parent" {
		emailsubject = "Cliiimb Registration"
		//sendmessage = fmt.Sprintf("Hello <b>testing</b> </br> your childern is try to register to cliimb please click here to to verifity your childern <a href='http://18.216.55.104:4200/email-verified?token=%s&useremail=%s'>Click here</a>", tokenString, useremail)
		newstring := `'<DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml">
		   <head>
		  <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		  <title>Cliiimb</title>
		  
		  <style type="text/css">
		 /* Client-specific Styles */
		 #outlook a {padding:0;} /* Force Outlook to provide a "view in browser" menu link. */
		 body{width:100%% ; -webkit-text-size-adjust:100%%; -ms-text-size-adjust:100%%; margin:0; padding:0;background: #F5F5F4;}
		 /* Prevent Webkit and Windows Mobile platforms from changing default font sizes, while not breaking desktop design. */
		 .ExternalClass {width:100%%;} /* Force Hotmail to display emails at full width */
		 .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div {line-height: 100%%;} /* Force Hotmail to display normal line spacing.*/
		 #backgroundTable {margin:0; padding:0; width:100%% ; line-height: 100%% ;background: #F5F5F4;}
		 img {outline:none; text-decoration:none;border:none; -ms-interpolation-mode: bicubic;}
		 a img {border:none;}
		 .image_fix {display:block;}
		 p {margin: 0px 0px ;}
		 table td {border-collapse: collapse;}
		 table { border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt; }
		 a {color: #0a8cce;text-decoration: none;text-decoration:none;}
		 /*STYLES*/
		 table[class=full] { width: 100%%; clear: both; }
		 /*IPAD STYLES*/
		 @media only screen and (max-width: 640px) {
		 a[href^="tel"], a[href^="sms"] {
		 text-decoration: none;
		 color: #0a8cce; /* or whatever your want */
		 pointer-events: none;
		 cursor: default;
		 }
		 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
		 text-decoration: default;
		 color: #0a8cce ;
		 pointer-events: auto;
		 cursor: default;
		 }
		 table[class=devicewidth] {width: 440px;text-align:center;}
		 table[class=devicewidthinner] {width: 420px;text-align:center;}
		 img[class=banner] {width: 440px;height:220px;}
		 img[class=colimg2] {width: 440px;height:220px;}
		 
		 
		 }
		 /*IPHONE STYLES*/
		 @media only screen and (max-width: 480px) {
		 a[href^="tel"], a[href^="sms"] {
		 text-decoration: none;
		 color: #0a8cce; /* or whatever your want */
		 pointer-events: none;
		 cursor: default;
		 }
		 .mobile_link a[href^="tel"], .mobile_link a[href^="sms"] {
		 text-decoration: default;
		 color: #0a8cce ; 
		 pointer-events: auto;
		 cursor: default;
		 }
		 table[class=devicewidth] {width: 280px;text-align:center;}
		 table[class=devicewidthinner] {width: 260px;text-align:center;}
		 img[class=banner] {width: 280px;height:140px;}
		 img[class=colimg2] {width: 280px;height:140px;}
		 td[class=mobile-hide]{display:none;}
		 td[class="padding-bottom25"]{padding-bottom:25px;}
		
		 }
		  </style>
		   </head>
		   <body bgcolor="#F5F5F4" style="background: #F5F5F4 ;"><br /><br />
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
		  <tr>
		 <td>
			<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
			   <tbody>
				  <tr>
					 <td align="center" height="20" style="font-size:1px; line-height:1px;">&nbsp;</td>
				  </tr>
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		 
		
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="header">
		   <tbody>
		  <tr>
		 <td>
			<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
			   <tbody>
				  <tr>
					 <td width="100%%">
						<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
						   <tbody>
						   
							  <tr>
								 <td>
									
									<table width="200" align="" border="0" cellpadding="0" cellspacing="0" class="devicewidth">
									   <tbody>
										  <tr>
											 <td width="200" height="45" align="">
												<div class="imgpop">
												   <a target="_blank" href="#">
												   <img src="https://s3.us-east-2.amazonaws.com/climbmentors/logo.svg" alt="" border="0" width="" height="54" style="display:block; border:none; outline:none; text-decoration:none;">
												   </a>
												</div>
											 </td>
										  </tr>
									   </tbody>
									</table>
								
								 </td>
							  </tr>
							  
						   </tbody>
						</table>
					 </td>
				  </tr>
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
		  <tr>
		 <td>
			<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
			   <tbody>
				  <tr>
					 <td width="100%%">
						<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
						   <tbody>
							  
							  <tr>
								 <td>
									<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
									   <tbody>
									
										  <tr>
											 <td style="font-family: Helvetica, arial, sans-serif; font-size: 30px; color: #0F1A59; text-align:left; line-height: 40px;" st-title="fulltext-heading">
												Thanks for joining Cliiimb
											 </td>
										  </tr>
										
										  <tr>
											 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #0F1A59; text-align:left; line-height: 25px;
		font-weight: bold;" st-content="fulltext-content">
		Your Childern is recently try to register on Cliiimb. Please Click the below button and allow to join Cliiimb.
											 </td>
										  </tr>
									
									   </tbody>
									</table>
								 </td>
							  </tr>
							 
						   </tbody>
						</table>
					 </td>
				  </tr>
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable">
		   <tbody>
		  <tr>
		 <td>
		   
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="2columns">
		   <tbody>
		  <tr>
		 <td>
		   
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="full-text">
		   <tbody>
		  <tr>
		 <td>
			<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth">
			   <tbody>
				  <tr>
					 <td width="100%%">
						<table width="600" cellpadding="20" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
						   <tbody>
							 
							  <tr>
								 <td>
									<table width="560" align="left" cellpadding="0" cellspacing="0" border="0" class="devicewidthinner" style="background: #fff;">
									   <tbody>
										 
										  <tr>
											 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #fff; text-align:left; line-height: 10px;" st-title="fulltext-title">
												<a href="https://cliiimb.com/parent-verification?token=%s&useremail=%s" style="padding: 16px 30px;background:#ec1c24;border-radius: 5px;font-weight:600;border:1px solid #ec1c24;color: #fff ;">Allow</a>
											 </td>
										  </tr>
										
										  <tr>
											 <td style="font-family: Helvetica, arial, sans-serif; font-size: 16px; color: #666666; text-align:center; line-height: 30px;" st-content="fulltext-content">
												
											 </td>
										  </tr>
										 
									   </tbody>
									</table>
								 </td>
							  </tr>
							  
							  <tr>
								 <td ></td>
							  </tr>
							  
						   </tbody>
						</table>
					 </td>
				  </tr>
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="seperator">
		   <tbody>
		  <tr>
		 <td>
			<table width="600" align="center" cellspacing="0" cellpadding="0" border="0" class="devicewidth" style="background: #fff;">
			   <tbody>
				  
				  <tr>
					 <td width="550" align="center" height="1" bgcolor="#d1d1d1" style="font-size:1px; line-height:1px;">&nbsp;</td>
				  </tr>
			 
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		<table width="100%%" bgcolor="#ffffff" cellpadding="0" cellspacing="0" border="0" id="backgroundTable" st-sortable="postfooter" >
		   <tbody>
		  <tr>
		 <td>
			<table width="600" cellpadding="0" cellspacing="0" border="0" align="center" class="devicewidth" style="background: #fff;">
			   <tbody>
				  <tr>
					 <td width="100%%">
						<table width="600" cellpadding="20" cellspacing="0" border="0" align="left" class="devicewidth" style="background: #fff;">
						   <tbody>
							  <tr>
								 <td align="left" valign="middle" style="font-family: Helvetica, arial, sans-serif; font-weight:300;font-size: 16px;color: #252C65" st-content="postfooter">
									Visit <a href="#" style="text-decoration: none; color: #ec1c24;font-weight:300;"> Cliiimb Site</a> 
								 </td>
							  </tr>
						  
						   </tbody>
						</table>
					 </td>
				  </tr>
			   </tbody>
			</table>
		 </td>
		  </tr>
		   </tbody>
		</table>
		
		   
		   </body>
		   </html>'`
		sendmessage = fmt.Sprintf(trimQuotes(newstring), tokenString, useremail)
	}
	//fmt.Println(sendmessage)
	m := gomail.NewMessage()
	m.SetHeader("From", "testproject628@gmail.com")
	m.SetHeader("To", email)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", emailsubject)
	m.SetBody("text/html", sendmessage)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "testproject628@gmail.com", "hello1234@")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return tokenString
}
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
func Updateaboutme(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(shared.UserUpdateData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserUpdateData{}
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
	//fmt.Println(res)
	//fmt.Println(res.Data)
	//fmt.Println(res)
	result := shared.UserUpdateData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"email": res.Email}).One(&result)
	newdata := shared.UserUpdateData{}
	newdata = result
	fmt.Println(res.AboutMe)
	newdata.AboutMe = res.AboutMe
	db.Update(result, newdata)

	//fmt.Println(string(buff))
	result1 := shared.UsergetData{}
	err = db.Find(bson.M{"email": res.Email}).One(&result1)
	mentorstatus := GetMentorRequest(res.Email)
	result1.MentorStatus = mentorstatus

	defer session.Close()

	return c.JSON(http.StatusOK, &result1)
}
func ViewProfile(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)

	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.Userinfores{}

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
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
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.Email
	//fmt.Println(email)
	//password:=c.FormValue("password")
	password := res.Password
	fmt.Println(password)

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"email": email}).All(&results.Data)

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
func ViewProfileById(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	results := shared.Userinfores{}

	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
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
	//fmt.Println(res.Email)

	//email := c.FormValue("id")
	email := res.ID
	//fmt.Println(email)
	//password:=c.FormValue("password")
	password := res.Password
	fmt.Println(password)

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"_id": email}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	//fmt.Println(results)
	mentorstatus := GetMentorRequest(results.Data[0].Email)

	results.Data[0].MentorStatus = mentorstatus
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)
}
func GetParentKids(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.PARENTCOLLECTION)
	results := Parentres{}

	u := new(ParentpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := ParentpostData{}
	//fmt.Println("this is C:",postData{})
	res = *u
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println("this is res=", res)
	os.Stdout.Write(b)

	var jsonBlob = []byte(b)
	var r ParentRes
	error := json.Unmarshal(jsonBlob, &r)
	if error != nil {
		fmt.Println("error:", error)
	}
	//fmt.Println(res.Email)

	//email :=c.FormValue("email")
	email := res.ParentEmail

	//err = db.Find(bson.M{"$or":[]bson.M{bson.M{"cms":cms},bson.M{"name":name}}}).All(&results.Data)

	err = db.Find(bson.M{"parentemail": email}).All(&results.Data)

	if err != nil {
		//log.Fatal(err)
	}
	if results.Data == nil {
		defer session.Close()
		return c.JSON(http.StatusOK, 0)
	}
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)
}
func PasswordChange(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.USERCOLLECTION)
	//name:=c.FormValue("Cms")
	//fmt.Println(name)
	//name =c.FormValue("name")
	u := new(shared.UserpostData)
	if err = c.Bind(&u); err != nil {
	}
	res := shared.UserpostData{}
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
	//fmt.Println(res)
	//fmt.Println(res.Data)
	//fmt.Println(res)
	result := shared.UserUpdateData{}
	//fmt.Println("%T \n", result)
	err = db.Find(bson.M{"email": res.Email}).One(&result)
	newdata := shared.UserUpdateData{}
	newdata = result

	if res.Password != "" {
		hash := hashAndSalt([]byte(res.Password))
		//res.Password = hash
		newdata.Password = hash
	}

	db.Update(result, newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &newdata)
}

// docker run --name mongodb \
//   -e MONGODB_USERNAME=mahar -e MONGODB_PASSWORD=hello123 \
//   -e MONGODB_DATABASE=cliimb bitnami/mongodb:latest
