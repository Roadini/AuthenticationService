package main

// Import our dependencies. We'll use the standard HTTP library as well as the gorilla router for this app
import (
	"net/http/httputil"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func PrintRequest(r *http.Request){

    requestDump, err := httputil.DumpRequest(r, true)
    if err != nil {
        log.Println(err)
    }
    log.Println(string(requestDump))

}

func main() {
    // Here we are instantiating the gorilla/mux router
    r := mux.NewRouter()
    PingDB();

    /* html */
    // http.Handle("/static/style.css", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/html"))))

    /* Gui */
    r.HandleFunc("/login", loginPageHandler)
    r.HandleFunc("/logout", logoutPageHandler)

    // On the default page we will simply serve our static index page.
    r.Handle("/auth/v1/createuser", CreateUserHandler).Methods("POST")
    r.Handle("/auth/v1/getusers", GetUsersHandler).Methods("POST")
    r.Handle("/auth/v1/getselfuser", GetSelfUser).Methods("POST")
    r.Handle("/auth/v1/deleteuser", DeleteUserHandler).Methods("POST")
    r.Handle("/auth/v1/login", LoginUserHandler).Methods("POST")
    r.Handle("/auth/v1/logout", LogoutUserHandler).Methods("POST")
    r.Handle("/auth/v1/updateuser", UpdateUserHandler).Methods("POST")
    r.Handle("/auth/v1/getallusers", GetAllUsersHandler).Methods("POST")

    r.Handle("/auth/v1/loginfb", LoginFb).Methods("POST")

    r.Handle("/social/v1/follows/follow", FollowsFollowHandler).Methods("POST")
    r.Handle("/social/v1/follows/getfollowrequests", FollowsFollowRequestsHandler).Methods("POST")
    r.Handle("/social/v1/follows/allowfollow", FollowsAllowFollowHandler).Methods("POST")
    r.Handle("/social/v1/follows/getfollowers", FollowsGetFollowers).Methods("POST")
    r.Handle("/social/v1/follows/getfollowing", FollowsGetFollowing).Methods("POST")
    r.Handle("/social/v1/follows/removeafollower", FollowsRemoveFollowerHandler).Methods("POST")
    r.Handle("/social/v1/follows/stopfollowing", FollowsRemoveFollowingHandler).Methods("POST")

    r.Handle("/social/v1/likes/like", LikesLikeHandler).Methods("POST")
    r.Handle("/social/v1/likes/getlikes", LikesGetLikesHandler).Methods("POST")
    r.Handle("/social/v1/likes/unlike", LikesUnlikeHandler).Methods("POST")
    
    r.Handle("/social/v1/publication/create", RoutesCreateHandler).Methods("POST")
    r.Handle("/social/v1/publication/remove", RoutesRemoveHandler).Methods("POST")
    r.Handle("/social/v1/publication/get", RoutesGetHandler).Methods("POST")


    // Our application will run on port 3000. Here we declare the port and pass in our router.
    log.Printf("running at http://localhost:3000\n")
    log.Fatal(http.ListenAndServe(":3000", r))
    
    log.Printf("running at http://localhost:3000\n")
}

