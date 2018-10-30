package main

import (
	"contribution"
	"favorites"
	"fmt"
	"following"
	"mentor"
	message "messages"
	"notification"
	"preferences"
	"user"
	"userpreference"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//e.GET("/", user.GetAll)
	// USER FUNCTIONS
	e.GET("/getalluser", user.GetAll)
	e.POST("/registration", user.Adduser)
	e.POST("/addadmin", user.AddAdmin)
	e.POST("/login", user.Login)
	e.PUT("/editprofile", user.EditProfile)
	e.PUT("/editaboutme", user.Updateaboutme)
	e.PUT("/updateprofile", user.UpdateProfile)
	e.POST("/viewprofile", user.ViewProfile)
	e.POST("/viewprofilebyid", user.ViewProfileById)

	// CONTRIBUTION FUNCITON
	//e.GET("/", getAll)
	e.GET("/showallcontribution", contribution.ContributionGetAll)
	e.GET("/showallevent", contribution.GetAllEvent)
	e.POST("/addcontribution", contribution.Addcontribution)
	e.POST("/searchcontribution", contribution.SearchContribution)                           //done
	e.POST("/searcheventbyemail", contribution.SearchEventByEmail)                           //done
	e.POST("/searchcontributionbyid", contribution.SearchContributionById)                   // done
	e.POST("/searchcontributionbycategory", contribution.SearchContributionByCategory)       // done
	e.POST("/searchcontributionbysubcategory", contribution.SearchContributionBySubCategory) // done
	e.PUT("/editcontribution", contribution.Editcontribution)                                // done
	e.POST("/deletecontribution", contribution.RemoveOneContribution)                        // done
	e.POST("/addview", contribution.AddView)                                                 //done
	e.POST("/updatecontributionstatus", contribution.UpdateContributionStatus)               //done
	e.POST("/updateadminstatus", contribution.UpdateAdminStatus)                             //done
	e.POST("/searchevent", contribution.SearchEvent)                                         // done
	e.POST("/searchingcontribution", contribution.SearchSubContribution)                     //done
	e.POST("/rejectcontribution", contribution.RejectContribution)                           //done
	e.GET("/getallrejectedcontribution", contribution.GetAllRejectedContribution)            //done

	// PREFERENCES FUNCTION

	//e.GET("/", preferences.GetAllPreferences)
	e.GET("/getallpreferences", preferences.GetAllPreferences)
	e.GET("/getallcategory", preferences.GetAllCategory)
	e.GET("/getallpreferencesbycategory", preferences.GetPrefencesbyCategory)
	e.POST("/addpreference", preferences.AddPreferences)
	e.POST("/addcategory", preferences.AddCategory)
	e.PUT("/removesubcategory", preferences.RemoveSubcategory)
	e.PUT("/addsubcategory", preferences.Addsubcategory)
	e.POST("/removecategory", preferences.RemoveCategory)

	e.PUT("/updateall", preferences.PutPreferences)

	// MENTOR / FOLLOWER FUNCTION

	//e.GET("/", following.GetAllData)
	e.GET("/getallfollwerdata", following.GetAllData)
	e.POST("/getfollower", following.Getfollower)
	e.POST("/getfollowerbyemail", following.GetfollowerByEmail)
	e.POST("/addmentor", following.AddMentor)
	e.POST("/updateparentstatus", following.UpdateParentStatus)
	e.POST("/updatemessagestatus", following.UpdateMessageStatus)
	e.PUT("/unfollow", following.Unfollow)
	e.PUT("/addfollower", following.Addfollower)

	// EMAIL VERIFICATION FUNCTION

	e.POST("/registrationverification", user.RegistrationVerfication)

	// PASSWORD UPDATE

	e.POST("/passwordupdate", user.PasswordVerification)
	e.POST("/passwordupdateverification", user.PasswordResetVerfication)
	e.POST("/passwordchange", user.PasswordChange)

	// PARENT VERIFICATION FUNCTION *****

	e.POST("/parentverification", user.ParentVerfication)
	e.POST("/getparentkids", user.GetParentKids)

	// USER PREFERENCES / SUGGESTION

	e.GET("/getalluserpreferences", userpreference.GetAllUserPreferences)
	e.POST("/adduserpreferences", userpreference.AddUserPreferences)
	e.POST("/getuserpreferences", userpreference.GetUserPrefences)
	e.PUT("/removeuserpreferences", userpreference.RemoveUserPreferences)
	e.POST("/usersuggestionpreferences", userpreference.UserPreferenceSuggestionContribution)
	//e.PUT("/removeuserpreferences", userpreference.RemoveUserPreferences

	// USER MESSAGES

	e.GET("/getallmessages", message.GetAllMessages)
	e.POST("/addusermessages", message.AddUserMessages)
	e.POST("/getusermessagesdetail", message.GetUserMessagesDetail)
	e.POST("/getusermessages", message.GetUserChat)
	e.PUT("/removeusermessages", message.RemoveUserMessages)
	e.POST("/getuserchatstatus", message.GetUserChatStatus)
	//e.PUT("/markasread", message.MarkAsRead)

	// CONTRIBUTION LIKES / COMMENTS

	e.GET("/getallfvrts", favorites.GetAllFvrtData)
	e.POST("/addlikes", favorites.AddLikes)
	e.POST("/addcomments", favorites.AddComments)
	e.POST("/getlikesandcomments", favorites.GetLikesAndComments)
	e.POST("/unlike", favorites.UnLike)
	e.POST("/deletecomments", favorites.DeleteComment)

	// HISTORY / NOTIFICATION
	e.POST("/getmentorhistory", notification.GetUserMentorHistory)
	e.POST("/changenotificationstatus", notification.ChangeNotificationStatus)

	// Become a Mentor

	e.POST("/becomementor", mentor.BecomeMentorRequest)
	e.GET("/admingetmentorrequests", mentor.GetAllMentorAdminRequest)
	e.POST("/getparentmentorrequest", mentor.GetMentorParentsRequest)
	e.POST("/updatementorparentrequest", mentor.UpdateParentStatus)
	e.POST("/updatementoradminrequest", mentor.UpdateAdminStatus)
	e.POST("/getmentorrequeststatus", mentor.GetMentorRequest)
	e.POST("/getmentorrequestdetail", mentor.GetMentorRequestDetail)
	e.POST("/remainingcontributioncheck", contribution.RemainingContributionCheck)
	e.POST("/adminmentorrequestreject", mentor.UpdateAdminRejectStatus)
	e.POST("/parentmentorrequestreject", mentor.UpdateRejectParentMentorStatus)

	// *****************************************************************
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("start...")
}
