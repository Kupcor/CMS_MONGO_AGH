package api
//Temp disabled

/*
package api

import (
    "fmt"
    "net/http"
    "context"
    "encoding/json"
	"time"

	"historycznymonolog/constants"
    "historycznymonolog/datamodels"

    "golang.org/x/crypto/bcrypt"

    "go.mongodb.org/mongo-driver/bson"

	"github.com/dgrijalva/jwt-go"
)

//Login user into application
func Login(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    fmt.Println("LOGIN METHOD")

    var user datamodels.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.USERS_COLLECTION_NAME)
	filter := bson.M{"username": user.Username}
	var result datamodels.User
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := generateJWTToken(result.ID.Hex(), result.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    response := map[string]string{"token": token}
    json.NewEncoder(w).Encode(response)
    
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Authorization", "Bearer "+token)
    w.WriteHeader(http.StatusOK)
}


func generateJWTToken(userID string, email string) (string, error) {
    expirationTime := time.Now().Add(12 * time.Hour)

    claims := jwt.MapClaims{
        "userID": userID,
        "email":  email,
        "exp":    expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(constants.SECRET_KEY))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
*/