package main

import (
    "log"
    "net/http"
    "strconv"
    "fmt"
    "github.com/gorilla/mux"
)

type server struct{}

// this is comments in golang
func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "[GET] HTTP method called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "[POST] HTTP method called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func patch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "[PATCH] HTTP method called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

func options(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "[OPTIONS] HTTP method called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "Unknown HTTP method"}`))
}

/// ----------------------------------------------------------------------------
/// Function to deal with query paramters
/// for the [GET /user/{userID}/comment/{commentID}] REST API endpoint
/// (GET HTTP method)
///
func getParams(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    userID := -1
    var err error
    if val, ok := pathParams["userID"]; ok {
        userID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "userID needs a number as value"}`))
            return
        }
    }

    commentID := -1
    if val, ok := pathParams["commentID"]; ok {
        commentID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "commentID needs a number as value"}`))
            return
        }
    }

    query := r.URL.Query()
    location := query.Get("location")

    w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}


/// ----------------------------------------------------------------------------
/// Function to deal with query paramters
/// for the [GET /user/{userID}/comment/{commentID}] REST API endpoint
/// (GET HTTP method)
/// returns {commentID} of the created comment
///
func postParams(w http.ResponseWriter, r *http.Request) {
    pathParams := mux.Vars(r)
    w.Header().Set("Content-Type", "application/json")

    userID := -1
    var err error
    if val, ok := pathParams["userID"]; ok {
        userID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "userID needs a number as value"}`))
            return
        }
    }

    /// Here I create the comment where it should be created (a comment on a github issue, using the Gihub v4 API ?)
    /// And in return, I get a commentID from the Github API
    /// mocked value
    commentID := 154577545445

    query := r.URL.Query()
    location := query.Get("location")

    w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s", "message": "Successfully added comment as ${commentID}" }`, userID, commentID, location)))
}


func main() {
  r := mux.NewRouter()
  api := r.PathPrefix("/api/v1").Subrouter()

  /// those below handle the root endpoint :
  ///         [GET /]
  ///         [POST /]
  ///         [PUT /]
  ///         [PATCH /]
  ///         [DELETE /]
  ///         [OPTIONS /]

  api.HandleFunc("/", get).Methods(http.MethodGet)
  api.HandleFunc("/", post).Methods(http.MethodPost)
  api.HandleFunc("/", put).Methods(http.MethodPut)
  api.HandleFunc("/", patch).Methods(http.MethodPatch)
  api.HandleFunc("/", delete).Methods(http.MethodDelete)
  api.HandleFunc("/", options).Methods(http.MethodOptions)
  api.HandleFunc("/", notFound)

  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the GET HTTP method :
  ///         [GET /user/{userID}/comment/{commentID}]
  api.HandleFunc("/user/{userID}/comment/{commentID}", getParams).Methods(http.MethodGet)

  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the POST HTTP method :
  ///         [POST /user/{userID}/comment] returns {commentID} of the created comment
  api.HandleFunc("/user/{userID}/comment", postParams).Methods(http.MethodPost)
  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the PUT HTTP method :
  ///         [PUT /user/{userID}/comment/{commentID}]

  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the PATCH HTTP method :
  ///         [PATCH /user/{userID}/comment/{commentID}]

  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the DELETE HTTP method :
  ///         [DELETE /user/{userID}/comment/{commentID}]

  /// the guy below handle the [/user/{userID}/comment/{commentID}] endpoint, for the OPTIONS HTTP method :
  ///         [OPTIONS /user/{userID}/comment/{commentID}]


  log.Fatal(http.ListenAndServe(":10101", r))
}
