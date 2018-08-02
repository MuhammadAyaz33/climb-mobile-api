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
	MentorStatus   int
}
type Userres struct {
	Data []UsergetData
}

type UserinfogetData struct {
	ID             bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email          string
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
	MentorStatus   int
}
type Userinfores struct {
	Data []UserinfogetData
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
	MentorStatus   int           `json:"mentorstatus"`
}
type UserRes struct {
	Data []UserpostData `json:"Data"`
}

type UpdateData struct {
	old    UserpostData
	change UserpostData
}

// 	BECOME A MENTOR
type BMentorgetData struct {
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserEmail            string
	UserID               string
	UserName             string
	Industory            string
	SkilLevel            string
	Experience           string
	WorkedFor            string
	CompanyName          string
	NumberOfContribution int
	MotivationTxt        string
	Donation             int
	UserAge              int
	AdminStatus          int
	ParentStatus         int
}
type BMentorres struct {
	Data []BMentorgetData
}

type BMentorpostData struct {
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserEmail            string        `json:"useremail"`
	UserID               string        `json:"userid"`
	UserName             string        `json:"username"`
	Industory            string        `json:"industry"`
	SkilLevel            string        `json:"skillevel"`
	Experience           string        `json:"experience"`
	WorkedFor            string        `json:"workedfor"`
	CompanyName          string        `json:"companyname"`
	NumberOfContribution int           `json:"numberofcontribution"`
	MotivationTxt        string        `json:"motivationTxt"`
	Donation             bool          `json:"donation"`
	UserAge              int           `json:"userage"`
	AdminStatus          int           `json:"adminstatus"`
	ParentStatus         int           `json:"parentstatus"`
}
type BMentorRes struct {
	Data []UserpostData `json:"Data"`
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
	ID                   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	UserEmail            string
	UserID               string
	UserFullName         string
	UserProfilePicture   string
	Title                string
	MainCategory         string
	SubCategories        string
	ContributionText     string
	Videos               string
	AudioPath            string
	Images               []Getimageurl
	Website              []Getwebsiteurl
	Coverpage            string
	Tags                 []Gettag
	ViewCount            int
	ContributionStatus   string
	AdminStatus          int
	Date                 string
	ContributionType     string
	Location             string
	ContributionPostDate time.Time
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
	ID                   bson.ObjectId   `json:"_id" bson:"_id,omitempty"`
	UserEmail            string          `json:"useremail"`
	UserID               string          `json:"userid"`
	UserFullName         string          `json:"username"`
	UserProfilePicture   string          `json:"userprofilepicture"`
	Title                string          `json:"title"`
	MainCategory         string          `json:"maincategory"`
	SubCategories        string          `json:"subcategories"`
	ContributionText     string          `json:"contributiontext"`
	Videos               string          `json:"videos"`
	AudioPath            string          `json:"audiopath"`
	Images               []Postimageurl  `json:"images"`
	Website              []Getwebsiteurl `json:"website"`
	Coverpage            string          `json:"coverpage"`
	Tags                 []Posttag       `json:"tags"`
	ViewCount            int             `json:"view"`
	ContributionStatus   string          `json:"contributionstatus"`
	AdminStatus          int             `json:"adminstatus"`
	Date                 string          `json:"date"`
	ContributionType     string          `json:"contributiontype"`
	Location             string          `json:"location"`
	ContributionPostDate time.Time       `json:"contributionpostdate"`
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
	DBName                         = "cliimb"
	CName                          = "product"
	DBURL                          = "172.25.33.205:27017"
	FILEBUCKETURL                  = "https://s3.us-east-2.amazonaws.com/climbmentors/"
	USERCOLLECTION                 = "user"
	CONTRIBUTIONCOLLECTION         = "contribution"
	PREFERENCESCOLLECTION          = "preferences"
	MENTORCOLLECTION               = "mentor"
	VERIFICATIONCOLLECTION         = "verfication"
	USERPREFERENCECOLLECTION       = "userprefrences"
	MESSAGESCOLLECTION             = "msguserdetail"
	CHATCOLLECTION                 = "chat"
	FAVORITESCOLLECTION            = "favorites"
	PASSWORDVERIFICATIONCOLLECTION = "passwordverficiation"
	PARENTCOLLECTION               = "parent"
	VIEWCOLLECTION                 = "contributionview"
	MENTORHISTORYCOLLECTION        = "mentorhistory"
	MENTORREQUESTCOLLECTION        = "mentorrequest"
)
