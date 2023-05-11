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

    "golang.org/x/crypto/bcrypt"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)


/*
GET METHODS
*/
//  ___ GetAllUsers ___
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	users, err := getAllUsers(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetAllUsers method finished...")
}

func getAllUsers(r *http.Request) ([]datamodels.User, error){
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    fmt.Println("Connection to 'USERS' collection initialize...")
    fmt.Println("GetAllUsers method...")

    //  Get all object from collection, filter is empty: bson.D{}
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error finding users: %v", err)
    }
    defer cur.Close(nil)

    //  Define users []User array
    var users []datamodels.User
    for cur.Next(nil) {
        var user datamodels.User
        if err := cur.Decode(&user); err != nil {
            return nil, fmt.Errorf("error decoding user: %v", err)
        }
        users = append(users, user)
    }
    return users, nil
}

//  ___ GetUserByUsername ___
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	user, err := getUserByUsername(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetUserByUsername method finished...")
}

func getUserByUsername(r *http.Request) (datamodels.User, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    fmt.Println("Connection to 'USER' collection initialize...")
    fmt.Println("GetUserByUsername method...")

    params := mux.Vars(r)
    username := params["username"]

    filter := bson.M{"username": username}

    var user datamodels.User
    if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
        return datamodels.User{}, err
    }

    return user, nil
}

//  ___ GetUserByID ___
func GetUserByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	user, err := getUserById(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetUserByID method finished...")
}

func getUserById(r *http.Request) (datamodels.User, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    fmt.Println("Connection to 'USER' collection initialize...")
    fmt.Println("GetUserByID method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        return datamodels.User{}, err
    }

    filter := bson.M{"_id": id}

    var user datamodels.User
    if err := collection.FindOne(context.Background(), filter).Decode(&user); err != nil {
        return datamodels.User{}, err
    }

    return user, nil
}

/*
POST METHODS
*/
//Add User to collection
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    var user datamodels.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    var result datamodels.User
    //  Check for dublicates
    err = collection.FindOne(context.Background(), bson.M{"$or": []bson.M{{"username": user.Username}, {"email": user.Email}}}).Decode(&result)
    if err == nil {
        http.Error(w, "Username or email already exists", http.StatusBadRequest)
        return
    }

    //  Hashing password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)
    user.CreatedAt = time.Now()

    addedUser, err := addUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(addedUser)
    fmt.Println("AddUser method finished...")
}

func addUser(user datamodels.User) (*mongo.InsertOneResult, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    fmt.Println("Connection to 'USER' collection initialize...")
    fmt.Println("AddUser method...")

    addedUser, err := collection.InsertOne(context.Background(), user)
    if err != nil {
        return nil, err
    }

    return addedUser, nil
}

/*
    Delete methods
*/
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("Start DeleteUserByID method...")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User deleted successfully")
}


//  Temporary disableds
/*
PUT METHODS

//  ChangeUserPassword
func ChangeUserPassword(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    fmt.Println("Start ChangePassword method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user datamodels.User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    var result datamodels.User
    //  Check if user exist
    err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
    if err != nil {
        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    update := bson.M{"$set": bson.M{"password": string(hashedPassword)}}
    _, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Println("ChangePassword method...")

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Password updated successfully")
}

// ChangeUserEmail
func ChangeUserEmail(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    fmt.Println("Start ChangeUserEmail method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user datamodels.User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println("Received user data: ", user)

    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
    var result datamodels.User
    err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
    if err != nil {
        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    fmt.Println("Found user with ID: ", id)

    update := bson.M{"$set": bson.M{"email": user.Email}}
    _, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Println("Updated user with ID: ", id)

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Email updated successfully")
}
*/