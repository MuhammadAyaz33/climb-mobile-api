package message

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"shared"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/xid"
	"gopkg.in/mgo.v2/bson"
)

//data to get from db ***********************************************************
type getProduct struct {
	MessageID           string
	Message             string
	MessageStatus       string
	MessageParentStatus int
	MessageTime         string
	SenderUserId        bson.ObjectId
}
type Message struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ChatID   string
	Messages []getProduct
}
type getData struct {
	ID                  bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User1ID             bson.ObjectId
	User1UserName       string
	User1ProfilePicture string
	User1ChatStatus     int
	User2ID             bson.ObjectId
	User2UserName       string
	User2ProfilePicture string
	User2ChatStatus     int
	ChatID              string
}
type res struct {
	Data []getData
}

//data from post********************************************************************
type postProduct struct {
	MessageID           string        `json:"messageid"`
	Message             string        `json:"message"`
	MessageStatus       string        `json:"messagestatus"`
	MessageParentStatus int           `json:"messageparentstatus"`
	MessageTime         string        `json:"messagetime"`
	SenderUserId        bson.ObjectId `json:"senderuserid"`
}
type Chat struct {
	ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ChatID   string        `json:"chatid"`
	Messages []postProduct `json:"messages"`
}
type postData struct {
	ID                  bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	User1ID             bson.ObjectId `json:"user1id"`
	User1UserName       string        `json:"user1username"`
	User1ProfilePicture string        `json:"user1profilepicture"`
	User1ChatStatus     int           `json:"user1chatstatus"`
	User2ID             bson.ObjectId `json:"user2id"`
	User2UserName       string        `json:"user2username"`
	User2ProfilePicture string        `json:"user2profilepicture"`
	User2ChatStatus     int           `json:"user2chatstatus"`
	ChatID              string        `json:"userchatid"`
}
type Res struct {
	Data []postData `json:"Data"`
}

type GetDataFromUser struct {
	SenderUserID   bson.ObjectId `json:"senderuserid"`
	ReceiverUserID bson.ObjectId `json:"receiveruserid"`
	Message        string        `json:"message"`
}

type GetChatDetail struct {
	ChatID   string        `json:"chatid"`
	SenderID bson.ObjectId `json:"senderid"`
}

//GET *********************************************************************************
func GetAllMessages(c echo.Context) error {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
	results := res{}
	err = db.Find(bson.M{}).All(&results.Data)

	//  |  for one result
	//  V
	//result := getData{}
	//err = db.Find(bson.M{"name": "two"}).One(&result)

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
func GetUserChat(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CHATCOLLECTION)
	results := Message{}

	u := new(GetChatDetail)
	if err = c.Bind(&u); err != nil {
	}
	res := GetChatDetail{}
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

	err = db.Find(bson.M{"chatid": res.ChatID}).One(&results)

	if err != nil {
		defer session.Close()
		return c.JSON(http.StatusOK, "data not found")
	}
	MarkAsRead(res.ChatID, res.SenderID)
	//fmt.Println(results)
	buff, _ := json.Marshal(&results)
	//fmt.Println(string(buff))

	json.Unmarshal(buff, &results)
	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}

func GetUserChatStatus(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
	results := getData{}

	u := new(GetChatDetail)
	if err = c.Bind(&u); err != nil {
	}
	res := GetChatDetail{}
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

	err = db.Find(bson.M{"user1id": res.SenderID}).One(&results)

	if err != nil {
		err1 := db.Find(bson.M{"user2id": res.SenderID}).One(&results)
		if err1 != nil {
			defer session.Close()
			return c.JSON(http.StatusOK, 0)
		} else {
			if results.User2ChatStatus == 1 {
				defer session.Close()
				return c.JSON(http.StatusOK, 1)
			} else {
				defer session.Close()
				return c.JSON(http.StatusOK, 0)
			}
		}

	} else {
		if results.User1ChatStatus == 1 {
			defer session.Close()
			return c.JSON(http.StatusOK, 1)
		} else {
			defer session.Close()
			return c.JSON(http.StatusOK, 0)
		}
	}

	defer session.Close()
	return c.JSON(http.StatusOK, &results)

}
func GetUserMessagesDetail(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
	results := res{}
	results2 := res{}
	newdata := res{}

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

	err = db.Find(bson.M{"user1id": res.User1ID}).All(&results.Data)

	if err != nil {
		fmt.Println("user1 no data")
		err = db.Find(bson.M{"user2id": res.User1ID}).All(&results.Data)
		if err != nil {
			defer session.Close()
			return c.JSON(http.StatusOK, "data not found")
		} else {
			newdata = results
			//fmt.Println(newdata)
		}

	} else {
		newdata = results
		err = db.Find(bson.M{"user2id": res.User1ID}).All(&results2.Data)
		if err != nil {
			fmt.Println("user2 no data")
			defer session.Close()
			return c.JSON(http.StatusOK, newdata)
		} else {
			fmt.Println("user2 data found")
			for i := range results2.Data {
				newdata.AddItem22(results2.Data[i])
			}
		}
	}

	// //fmt.Println(results)
	// buff, _ := json.Marshal(&newdata)
	// //fmt.Println(string(buff))

	// json.Unmarshal(buff, &newdata)
	defer session.Close()
	return c.JSON(http.StatusOK, &newdata)

}
func (box *res) AddItem22(item getData) []getData {
	box.Data = append(box.Data, item)
	return box.Data
}

//POST *********************************************************************************
func AddUserMessages(c echo.Context) (err error) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)

	u := new(GetDataFromUser)
	if err = c.Bind(&u); err != nil {
	}
	res := GetDataFromUser{}
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
	//fmt.Println(res)
	senderuserinfo := shared.UsergetData{}
	receiveruserinfo := shared.UsergetData{}

	senderuserinfo = UserInfo(res.SenderUserID)
	receiveruserinfo = UserInfo(res.ReceiverUserID)

	//fmt.Println("userinfo")
	//fmt.Println(receiveruserinfo)
	age, err := strconv.Atoi(senderuserinfo.Age)
	if err != nil {
		fmt.Println(err)
		fmt.Println(age)

	}

	result := getData{}
	err = db.Find(bson.M{"user1id": res.SenderUserID, "user2id": res.ReceiverUserID}).One(&result)
	mid := xid.New()
	//send message data save
	if err != nil {

		err1 := db.Find(bson.M{"user1id": res.ReceiverUserID, "user2id": res.SenderUserID}).One(&result)
		if err1 != nil {
			userdetail := postData{}
			userdetail.User1ID = res.SenderUserID
			userdetail.User1UserName = senderuserinfo.FullName
			userdetail.User1ProfilePicture = senderuserinfo.ProfilePicture
			userdetail.User2ID = res.ReceiverUserID
			userdetail.User2UserName = receiveruserinfo.FullName
			userdetail.User2ProfilePicture = receiveruserinfo.ProfilePicture
			userdetail.ChatID = mid.String()
			//read = 0
			//unread = 1
			userdetail.User1ChatStatus = 0
			userdetail.User2ChatStatus = 1
			db.Insert(userdetail)
			AddChatMessages(mid.String(), res.Message, res.SenderUserID, age)
			// if a == true {
			// 	fmt.Println("chat added")
			// }
			//fmt.Println(age)
			defer session.Close()
			return c.JSON(http.StatusOK, "new msg create")
		} else {
			newdata := getData{}
			newdata = result
			//read = 0
			//unread = 1
			newdata.User1ChatStatus = 1
			newdata.User2ChatStatus = 0

			db.Update(result, newdata)
			AddChatMessages(result.ChatID, res.Message, res.SenderUserID, age)
			// if a == true {
			// 	fmt.Println("chat added")
			// }
			fmt.Println("user found")
			defer session.Close()
			return c.JSON(http.StatusOK, "chat update")
		}

	} else {
		newdata := getData{}
		newdata = result
		//read = 0
		//unread = 1
		newdata.User1ChatStatus = 0
		newdata.User2ChatStatus = 1

		db.Update(result, newdata)
		AddChatMessages(result.ChatID, res.Message, res.SenderUserID, age)
		// if a == true {
		// 	fmt.Println("chat added")
		// }
		fmt.Println("user found")
		defer session.Close()
		return c.JSON(http.StatusOK, "chatupdate")
	}

	defer session.Close()
	return c.JSON(http.StatusOK, &r)
}

func AddChatMessages(chatid string, msg string, senderid bson.ObjectId, age int) {

	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.CHATCOLLECTION)
	fmt.Println(chatid)
	result := Message{}
	err = db.Find(bson.M{"chatid": chatid}).One(&result)
	fmt.Println(result)
	//send message data save
	newdata := Message{}
	if err != nil {
		fmt.Println("chat not found")
		messagedetail := Chat{}
		messagedetail.ChatID = chatid
		msgstatus := "unread"
		msgtime := time.Now().Format("2006-01-02 3:4:5 PM")
		var parentstatus int
		if age < 18 {
			parentstatus = 1
		} else {
			parentstatus = 0
		}
		//msgid := bson.NewObjectId()
		mid := xid.New()
		item1 := postProduct{MessageID: mid.String(), Message: msg, MessageStatus: msgstatus, MessageTime: msgtime, MessageParentStatus: parentstatus, SenderUserId: senderid}
		messagedetail.AddItem(item1)
		db.Insert(messagedetail)

	} else {

		newdata = result
		msgstatus := "unread"
		msgtime := time.Now().Format("2006-01-02 3:4:5 PM")
		var parentstatus int
		if age < 18 {
			parentstatus = 1
		} else {
			parentstatus = 0
		}
		//msgid1 := bson.NewObjectId()
		mid := xid.New()
		item11 := getProduct{MessageID: mid.String(), Message: msg, MessageStatus: msgstatus, MessageTime: msgtime, MessageParentStatus: parentstatus, SenderUserId: senderid}
		newdata.AddMessage(item11)
		//fmt.Println(newdata)

		db.Update(result, newdata)
	}

	defer session.Close()

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

func RemoveUserMessages(c echo.Context) (err error) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)
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

	err = db.Find(bson.M{"userid": res.ID}).One(&result)

	//result.removeFriend(res)

	result1 := getData{}

	err = db.Find(bson.M{"userid": res.ID}).One(&result1)
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

// func (self *getData) removeFriend(item postData) {
// 	for i := range self.SendMessages {
// 		if self.SendMessages[i].MessageID == item.SendMessages[0].MessageID {
// 			self.SendMessages = append(self.SendMessages[:i], self.SendMessages[i+1:]...)
// 			//fmt.Println(i)
// 			//fmt.Println("match ho geya")
// 			break
// 		}
// 	}
// }

func MarkAsRead(chatid string, senderid bson.ObjectId) {
	session, err := shared.ConnectMongo(shared.DBURL)
	db := session.DB(shared.DBName).C(shared.MESSAGESCOLLECTION)

	result := getData{}
	newdata := getData{}

	err = db.Find(bson.M{"chatid": chatid}).One(&result)
	newdata = result
	//read = 0
	//unread = 1
	if result.User1ID == senderid {
		newdata.User1ChatStatus = 0
	}
	if result.User2ID == senderid {
		newdata.User2ChatStatus = 0
	}
	//newdata.ChatStatus = 0
	db.Update(result, newdata)
	//err = db.Update(result, newdata)
	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

}

func (box *Chat) AddItem(item postProduct) []postProduct {
	box.Messages = append(box.Messages, item)
	return box.Messages
}
func (box *Message) AddMessage(items getProduct) []getProduct {
	box.Messages = append(box.Messages, items)
	return box.Messages
}

// func (newdata *getData) markRead(res postData) {
// 	for i := range newdata.SendMessages {
// 		if newdata.SendMessages[i].MessageID == res.SendMessages[0].MessageID {
// 			fmt.Println("matched")
// 			newdata.SendMessages[i].MessageStatus = "read"
// 			fmt.Println(newdata.SendMessages[i].MessageStatus)
// 			break
// 		}
// 	}
// }
