package api

import (
    "fmt"
    "net/http"
    "context"
    "encoding/json"

	"historycznymonolog/constants"
    "historycznymonolog/datamodels"

    "github.com/gorilla/mux"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

//	First get method -> get all available categories from MongoDB Cluster
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	categories, err := getAllCategories(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(categories); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	fmt.Println("getAllCategories method finished...")
}

func getAllCategories(r *http.Request) ([]datamodels.Category, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CATEGORIES_COLLECTION_NAME)
    fmt.Println("Connection to 'CATEGORIES' collection initialize...")
    fmt.Println("getAllCategories method...")

    //  Get all object from collection, filter is empty: bson.D{}
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, fmt.Errorf("error finding categories: %v", err)
    }
    defer cur.Close(nil)

    var categories []datamodels.Category
    for cur.Next(nil) {
        var category datamodels.Category
        if err := cur.Decode(&category); err != nil {
            return nil, fmt.Errorf("error decoding category: %v", err)
        }
        categories = append(categories, category)
    }
    return categories, nil
}

// Second get method -> get category by its id
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	category, err := getCategoryByID(r)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            http.Error(w, "Category not found", http.StatusNotFound)
        } else {
            http.Error(w, "Internal server error", http.StatusInternalServerError)
        }
        return
    }
    if err := json.NewEncoder(w).Encode(category); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("GetCategoryByID method finished...")
}

func getCategoryByID(r *http.Request) (datamodels.Category, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CATEGORIES_COLLECTION_NAME)
    fmt.Println("Connection to 'CATEGORIES' collection initialize...")
    fmt.Println("GetCategoryByID method...")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        return datamodels.Category{}, err
    }

    filter := bson.M{"_id": id}

    var category datamodels.Category
    if err := collection.FindOne(context.Background(), filter).Decode(&category); err != nil {
        return datamodels.Category{}, err
    }

    return category, nil
}

//	One put method -> add new category to categories Collection in MongoDB Cluster
func AddCategory(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    var category datamodels.Category
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


    addedCategory, err := addCategory(category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(addedCategory)
    fmt.Println("AddCategory method finished...")
}

func addCategory(category datamodels.Category) (*mongo.InsertOneResult, error) {
    collection := client.Database(constants.DATA_BASE_NAME).Collection(constants.CATEGORIES_COLLECTION_NAME)
    fmt.Println("Connection to 'CATEGORIES' collection initialize...")
    fmt.Println("AddCategory method...")

    addedCategory, err := collection.InsertOne(context.Background(), category)
    if err != nil {
        return nil, err
    }

    return addedCategory, nil
}