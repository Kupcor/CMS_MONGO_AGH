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
//  ___ GetAllContents ___
func GetAllContents(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	contents, err := getAllContents(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(contents); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetAllContents method finished...")
}

func getAllContents(r *http.Request) ([]datamodels.Content, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CONTENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'CONTENTS' collection initialize...")
    fmt.Println("GetAllContents method...")

    //  Get all object from collection, filter is empty: bson.D{}
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error finding contents: %v", err)
    }
    defer cur.Close(nil)

    //  Define contents []Content array
    var contents []datamodels.Content
    for cur.Next(nil) {
        var content datamodels.Content
        if err := cur.Decode(&content); err != nil {
            return nil, fmt.Errorf("error decoding content: %v", err)
        }
        contents = append(contents, content)
    }
    return contents, nil
}

//	___	Get Contents For Users __
func GetContentsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    contents, err := getContentsForUser(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Content not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(contents); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Println("getContentsForUser method finished...")

}

func getContentsForUser(r *http.Request) ([]datamodels.Content, error) {
	collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CONTENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'CONTENTS' collection initialize...")
    fmt.Println("getContentsForUser method...")

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

    // convert the cursor to an array of contents
    var contents []datamodels.Content
    if err = cursor.All(context.Background(), &contents); err != nil {
        return nil, err
    }

    return contents, nil
}

// Second get method -> get content by its id
func GetContentByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	content, err := getContentByID(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Content not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(content); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetContentByID method finished...")
}

func getContentByID(r *http.Request) (datamodels.Content, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CONTENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'CONTENTS' collection initialize...")
    fmt.Println("GetContentByID method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        return datamodels.Content{}, err
    }

    filter := bson.M{"_id": id}

    var content datamodels.Content
    if err := collection.FindOne(context.Background(), filter).Decode(&content); err != nil {
        return datamodels.Content{}, err
    }

    return content, nil
}

/*
POST METHODS
*/
//Add Content to collection 
func AddContent(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    var content datamodels.Content
    err := json.NewDecoder(r.Body).Decode(&content)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    addedContent, err := addContent(content)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(addedContent)
    fmt.Println("AddContent method finished...")
}

func addContent(content datamodels.Content) (*mongo.InsertOneResult, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CONTENTS_COLLECTION_NAME)
    fmt.Println("Connection to 'CONTENTS' collection initialize...")
    fmt.Println("AddContent method...")

    content.CreatedAt = time.Now()
    addedContent, err := collection.InsertOne(context.Background(), content)
    if err != nil {
        return nil, err
    }

    return addedContent, nil
}