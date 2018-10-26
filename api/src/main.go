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

    // On the default page we will simply serve our static index page.
    r.Handle("/auth/v1/createuser", CreateUserHandler).Methods("POST")
    r.Handle("/auth/v1/getusers", GetUsersHandler).Methods("POST")
    r.Handle("/auth/v1/deleteuser", DeleteUserHandler).Methods("POST")
    //r.Handle("/auth/v1/login", login).Methods("POST")
    //r.Handle("/auth/v1/updateuser", updateUser).Methods("POST")
    r.Handle("/social/v1/follows/follow", FollowsFollowHandler).Methods("POST")
    r.Handle("/social/v1/follows/getfollowrequests", FollowsFollowRequestsHandler).Methods("POST")
    r.Handle("/social/v1/follows/allowfollow", FollowsAllowFollowHandler).Methods("POST")
    r.Handle("/social/v1/follows/removeafollower", FollowsRemoveFollowerHandler).Methods("POST")
    r.Handle("/social/v1/follows/stopfollowing", FollowsRemoveFollowingHandler).Methods("POST")
    r.Handle("/social/v1/likes/like", LikesLikeHandler).Methods("POST")
    r.Handle("/social/v1/likes/getlikes", LikesGetLikesHandler).Methods("POST")
    r.Handle("/social/v1/likes/unlike", LikesUnlikeHandler).Methods("POST")
    r.Handle("/social/v1/routes/create", RoutesCreateHandler).Methods("POST")
    r.Handle("/social/v1/routes/remove", RoutesRemoveHandler).Methods("POST")
    r.Handle("/social/v1/routes/get", RoutesGetHandler).Methods("POST")



    // Our application will run on port 3000. Here we declare the port and pass in our router.
    log.Printf("running at http://localhost:3000\n")
    http.ListenAndServe(":3000", r)
}

