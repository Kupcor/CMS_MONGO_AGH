package api

import (
    "fmt"
    "time"

    "net/http"
    "context"
    "encoding/json"

	"historycznymonolog/constants"
    "historycznymonolog/datamodels"

    "github.com/gorilla/mux"
    
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
)


/*
GET METHODS
*/
//  ___ GetAllComments ___
func GetAllComments(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	comments, err := getAllComments(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(comments); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetAllComments method finished...")
}

func getAllComments(r *http.Request) ([]datamodels.Comment, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.COMMENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'COMMENTS' collection initialize...")
    fmt.Println("GetAllComments method...")

    //  Get all object from collection, filter is empty: bson.D{}
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error finding comments: %v", err)
    }
    defer cur.Close(nil)

    //  Define comments []Comment array
    var comments []datamodels.Comment
    for cur.Next(nil) {
        var comment datamodels.Comment
        if err := cur.Decode(&comment); err != nil {
            return nil, fmt.Errorf("error decoding comment: %v", err)
        }
        comments = append(comments, comment)
    }
    return comments, nil
}

//	___	Get Comments For Appropriate Post __
func GetCommentsForPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    comments, err := getCommentsForPost(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Comments not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(comments); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Println("getCommentsForPost method finished...")
}

func getCommentsForPost(r *http.Request) ([]datamodels.Comment, error) {
	collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.COMMENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'COMMENTS' collection initialize...")
    fmt.Println("getCommentsForPost method...")

    params := mux.Vars(r)
    postId, err := primitive.ObjectIDFromHex(params["postId"])
    if err != nil {
        return nil, err
    }

    filter := bson.M{"postId": postId}
    cursor, err := collection.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    // convert the cursor to an array of comments
    var comments []datamodels.Comment
    if err = cursor.All(context.Background(), &comments); err != nil {
        return nil, err
    }
    fmt.Println(comments)

    return comments, nil
}

//	___	Get Comments For Users __
func GetCommentsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    comments, err := getCommentsForUser(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Comment not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(comments); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Println("getCommentsForUser method finished...")

}

func getCommentsForUser(r *http.Request) ([]datamodels.Comment, error) {
	collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.COMMENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'COMMENTS' collection initialize...")
    fmt.Println("getCommentsForUser method...")

    params := mux.Vars(r)
    authorId, err := primitive.ObjectIDFromHex(params["authorId"])
    if err != nil {
        return nil, err
    }

    filter := bson.M{"authorId": authorId}
    cursor, err := collection.Find(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    // convert the cursor to an array of comments
    var comments []datamodels.Comment
    if err = cursor.All(context.Background(), &comments); err != nil {
        return nil, err
    }

    return comments, nil
}

/*
POST METHODS
*/
//Add Comment to collection 
func AddComment(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    var comment datamodels.Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    addedComment, err := addComment(comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(addedComment)
    fmt.Println("AddComment method finished...")
}

func addComment(comment datamodels.Comment) (*mongo.InsertOneResult, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.COMMENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'COMMENTS' collection initialize...")
    fmt.Println("AddComment method...")

    comment.CreatedAt = time.Now()
    addedComment, err := collection.InsertOne(context.Background(), comment)
    if err != nil {
        return nil, err
    }

    return addedComment, nil
}