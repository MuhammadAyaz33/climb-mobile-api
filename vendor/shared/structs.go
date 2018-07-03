package shared

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// USERDATA  **********************************************************

type UsergetData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email          string
	Password       string
	CompanyName    string
	FullName       string
	Address        string
	City           string
	ZipCode        int
	Bio            string
	Age            string
	ParentPhone    int
	ParentEmail    string
	AboutMe        string
	Status         int
	ParentStatus   int
	ProfilePicture string
	UserType       string
}
type Userres struct {
	Data []UsergetData
}

type UserpostData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	CompanyName    string        `json:"companyName"`
	FullName       string        `json:"fullName"`
	Address        string        `json:"address"`
	City           string        `json:"city"`
	ZipCode        int           `json:"zipCode"`
	Bio            string        `json:"bio"`
	Age            string        `json:"age"`
	ParentPhone    int           `json:"parentPhone"`
	ParentEmail    string        `json:"parentEmail"`
	AboutMe        string        `json:"aboutMe"`
	Status         int           `json:"status"`
	ParentStatus   int           `json:"parentstatus"`
	ProfilePicture string        `json:"profilepicture"`
	UserType       string        `json:"usertype"`
}
type UserRes struct {
	Data []UserpostData `json:"Data"`
}
type UpdateData struct {
	old    UserpostData
	change UserpostData
}

// 	EMAIL VERIFICATION DATA ******************
type VerificationgetData struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	EmailID string
	Token   string
	Date    time.Time
}
type Verificationres struct {
	Data []VerificationgetData
}

type VerificationpostData struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	EmailID string        `json:"email"`
	Token   string        `json:"token"`
	Date    time.Time     `json:"date"`
}
type VerificationRes struct {
	Data []VerificationpostData `json:"Data"`
}

// PASSWORD UPDATE

type PasswordVerificationgetData struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	EmailID string
	Token   string
	Date    time.Time
}
type PasswordVerificationres struct {
	Data []VerificationgetData
}

type PasswordVerificationpostData struct {
	ID      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	EmailID string        `json:"email"`
	Token   string        `json:"token"`
	Date    time.Time     `json:"date"`
}
type PasswordVerificationRes struct {
	Data []VerificationpostData `json:"Data"`
}

// CONTRIBUTION VIEW **************

type ViewgetData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ContributionID string
	UserID         string
}
type Viewres struct {
	Data []ViewgetData
}

type ViewpostData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ContributionID string        `json:"contributionid"`
	UserID         string        `json:"userid"`
}
type ViewRes struct {
	Data []VerificationpostData `json:"Data"`
}

//CONTRIBUTION DATA ***********************************************

type Getimageurl struct {
	Imagestatus string
}
type Getwebsiteurl struct {
	Websiteurl string
}
type Gettag struct {
	Tag string
}
type Getsubcategory struct {
	Subcategory string
}
type ContributionData struct {
	ID                 bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserEmail          string
	UserID             string
	UserFullName       string
	UserProfilePicture string
	Title              string
	MainCategory       string
	SubCategories      string
	ContributionText   string
	Videos             string
	AudioPath          string
	Images             []Getimageurl
	Website            []Getwebsiteurl
	Coverpage          string
	Tags               []Gettag
	ViewCount          int
	ContributionStatus int
	AdminStatus        int
}
type Contributionres struct {
	Data []ContributionData
}

type Postimageurl struct {
	Imagestatus string `json:"imagestatus"`
}
type Postwebsiteurl struct {
	Websiteurl string `json:"websiteurl"`
}
type Posttag struct {
	Tag string `json:"tag"`
}
type Postsubcategory struct {
	Subcategory string `json:"subcategory"`
}
type ContributionPostData struct {
	ID                 bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	UserEmail          string          `json:"useremail"`
	UserID             string          `json:"userid"`
	UserFullName       string          `json:"username"`
	UserProfilePicture string          `json:"userprofilepicture"`
	Title              string          `json:"title"`
	MainCategory       string          `json:"maincategory"`
	SubCategories      string          `json:"subcategories"`
	ContributionText   string          `json:"contributiontext"`
	Videos             string          `json:"videos"`
	AudioPath          string          `json:"audiopath"`
	Images             []Postimageurl  `json:"images"`
	Website            []Getwebsiteurl `json:"website"`
	Coverpage          string          `json:"coverpage"`
	Tags               []Posttag       `json:"tags"`
	ViewCount          int             `json:"view"`
	ContributionStatus int             `json:"contributionstatus"`
	AdminStatus        int             `json:"adminstatus"`
}
type ContributionRes struct {
	Data []ContributionPostData `json:"Data"`
}

// PREFERENCES DATA  **************************************************
type CheckStatus struct {
	status string `json:"checkstatus"`
}

type ErrorCheckStatus struct {
	Status string `json:"errorstatus"`
}

//connection to mongo ***************************************************************
const (
	DBName                   = "cliimb"
	CName                    = "product"
	DBURL                    = "172.25.33.205:27017"
	FILEBUCKETURL            = "https://s3.us-east-2.amazonaws.com/climbmentors/"
	USERCOLLECTION           = "user"
	CONTRIBUTIONCOLLECTION   = "contribution"
	PREFERENCESCOLLECTION    = "preferences"
	MENTORCOLLECTION         = "mentor"
	VERIFICATIONCOLLECTION   = "verfication"
	USERPREFERENCECOLLECTION = "userprefrences"
	MESSAGESCOLLECTION       = "msguserdetail"
	CHATCOLLECTION           = "chat"

	FAVORITESCOLLECTION            = "favorites"
	PASSWORDVERIFICATIONCOLLECTION = "passwordverficiation"
	PARENTCOLLECTION               = "parent"
	VIEWCOLLECTION                 = "contributionview"
)
