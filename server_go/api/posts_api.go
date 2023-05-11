package api


import (
    "fmt"
    "net/http"
    "context"
    "encoding/json"
    "time"

	"historycznymonolog/constants"
    "historycznymonolog/datamodels"

    "github.com/gorilla/mux"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)




//  ___ GetAllPosts ___
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	posts, err := getAllPosts(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetAllPosts method finished...")
}

func getAllPosts(r *http.Request) ([]datamodels.Post, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
    fmt.Println("Connection to 'POSTS' collection initialize...")
    fmt.Println("GetAllPosts method...")

    //  Get all object from collection, filter is empty: bson.D{}
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error finding posts: %v", err)
    }
    defer cur.Close(nil)

    //  Define posts []Post array
    var posts []datamodels.Post
    for cur.Next(nil) {
        var post datamodels.Post
        if err := cur.Decode(&post); err != nil {
            return nil, fmt.Errorf("error decoding post: %v", err)
        }
        posts = append(posts, post)
    }
    return posts, nil
}

//  ___ GetPostByTitle ___
func GetPostByTitle(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	posts, err := getPostByTitle(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetPostByTitle method finished...")
}

func getPostByTitle(r *http.Request) (datamodels.Post, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
    fmt.Println("Connection to 'POSTS' collection initialize...")
    fmt.Println("GetPostByTitle method...")

    params := mux.Vars(r)
    title := params["title"]

    filter := bson.M{"title": title}

    var post datamodels.Post
    if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
        return datamodels.Post{}, err
    }

    return post, nil
}

//  ___ GetPostById ___
func GetPostById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	post, err := getPostById(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(post); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetPostById method finished...")
}

func getPostById(r *http.Request) (datamodels.Post, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
    fmt.Println("Connection to 'POSTS' collection initialize...")
    fmt.Println("GetPostById method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        return datamodels.Post{}, err
    }

    filter := bson.M{"_id": id}

    var post datamodels.Post
    if err := collection.FindOne(context.Background(), filter).Decode(&post); err != nil {
        return datamodels.Post{}, err
    }

    return post, nil
}

//  __ GetPostByAuthor __
func GetPostByAuthor(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    posts, err := getPostByAuthor(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    fmt.Println("GetPostByAuthor method finished...")  
}

func getPostByAuthor(r *http.Request) ([]datamodels.Post, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
    fmt.Println("Connection to 'POSTS' collection initialize...")
    fmt.Println("GetPostByAuthor method...")

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

    // convert the cursor to an array of posts
    var posts []datamodels.Post
    if err = cursor.All(context.Background(), &posts); err != nil {
        return nil, err
    }

    return posts, nil
}

//  Get post by category


//Add Single Post to collection 
func AddSinglePost(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    var post datamodels.Post
    err := json.NewDecoder(r.Body).Decode(&post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    post.CreatedAt = time.Now()
    addedPost, err := addSinglePost(post)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(addedPost)
    fmt.Println("AddSinglePost method finished...")
}

func addSinglePost(post datamodels.Post) (*mongo.InsertOneResult, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
    fmt.Println("Connection to 'POSTS' collection initialize...")
    fmt.Println("AddSinglePost method...")

    addedPost, err := collection.InsertOne(context.Background(), post)
    if err != nil {
        return nil, err
    }

    return addedPost, nil
}


func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Start DeletePost method...")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.POSTS_COLLECTION_NAME)
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Delete Post method...")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Post deleted successfully")
}
