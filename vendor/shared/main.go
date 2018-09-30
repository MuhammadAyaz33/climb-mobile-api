package shared

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

func ConnectMongo(url string) (session *mgo.Session, err error) {
	MongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"18.216.55.104:27017"},
		Timeout:  60 * time.Second,
		Database: "cliimb",
		Username: "mahar",
		Password: "hello123",
	}
	if session == nil {
		session, err = mgo.DialWithInfo(MongoDBDialInfo)
		// Optional. Switch the session to a monotonic behavior.
		//session.SetMode(mgo.Monotonic, true)
		if err != nil {
			return session, err
		} else {
			return session, nil
		}
	} else {
		return session, err
	}
}

// ReturnMessage returns a response
func ReturnMessage(status bool, message string, code int64, data interface{}) Response {
	return Response{Status: status, Message: message, Code: code, Data: data}
}
