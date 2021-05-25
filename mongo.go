package main

import (
	"context"
	"fmt"
	logger "logger"
)

// UpdateTokenInDB - updates token in cache and DB
func (au *AllUsers) UpdateTokenInDB(user User, tokenDB string, endless bool) bool {
	if err := UpdateID(user.ID, obj{"$set": obj{"token": token{tokenDB, endless}}}); err != nil {
		logger.Error.Println("users.go -> UpdateTokenInDB -> UpdateID: err =", err)
		return false
	}
	user.Token.Token = tokenDB
	user.Token.Endless = endless
	au.UpdateUserInMap(user)
	return true
}

// UpdateBuildingsInDB - updates building in cache and DB
func (au *AllUsers) UpdateBuildingsInDB(user User, dbb []DBBuilding) bool {
	if err := UpdateID(user.ID, obj{"$set": obj{"fields": dbb}}); err != nil {
		logger.Error.Println("users.go -> UpdateBuildingsInDB -> UpdateID: err =", err)
		return false
	}
	au.UpdateUserInMap(user)
	return true
}

// AddNewUser - Adds new user to DB
func (au *AllUsers) AddNewUser(user User) error {
	// Inserting user to DB
	insertResult, err := Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("mongo.go -> AddNewUser -> InsertOne: user = %+v; err = %s", user, err)
	}
	logger.Info.Println("mongo.go -> AddNewUser -> InsertOne: ", insertResult)
	user.Token.Endless = false
	au.UpdateUserInMap(user)
	return nil
}

// GetAllUsers - gets all users from DB and return slice of them
func GetAllUsers() []User {

	var results []User

	cur, err := Collection.Find(context.TODO(), obj{})
	if err != nil {
		logger.Error.Println(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// Create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			logger.Error.Println(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		logger.Error.Println(err)
	}

	// Closing the cursor
	cur.Close(context.TODO())

	return results
}

// GetAllSurveys - gets all users from DB and return slice of them
func GetAllSurveys() []Survey {

	var results []Survey

	cur, err := CollectionInfo.Find(context.TODO(), obj{})
	if err != nil {
		logger.Error.Println(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// Create a value into which the single document can be decoded
		var elem Survey
		err := cur.Decode(&elem)
		if err != nil {
			logger.Error.Println(err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		logger.Error.Println(err)
	}

	// Closing the cursor
	cur.Close(context.TODO())

	return results
}

// UpdateID - updates user by id in DB
func UpdateID(id uint64, query interface{}) error {
	filter := obj{"_id": id}

	updateResult, err := Collection.UpdateOne(context.TODO(), filter, query)

	if err != nil {
		return fmt.Errorf("mongo.go -> UpdateID -> UpdateOne: id = %d; query = %+v; err = %s", id, query, err)
	}
	logger.Info.Printf("mongo.go -> UpdateID -> UpdateOne: %+v; id = %d; query = %+v\n", updateResult, id, query)
	return nil
}

// AddNewSurvey - Adds new survey to DB
func (as *AllSurveys) AddNewSurvey(sur Survey) error {
	// Inserting survey to DB
	insertResult, err := CollectionInfo.InsertOne(context.TODO(), sur)
	if err != nil {
		return fmt.Errorf("mongo.go -> AddNewSurvey -> InsertOne: survey = %+v; err = %s", sur, err)
	}
	logger.Info.Println("mongo.go -> AddNewSurvey -> InsertOne: ", insertResult)
	as.UpdateSurveyInMap(sur)
	return nil
}

// UpdateAskedInDB - updates token in cache and DB
func (au *AllUsers) UpdateAskedInDB(user User, asked bool) bool {
	if err := UpdateID(user.ID, obj{"$set": obj{"asked": asked}}); err != nil {
		logger.Error.Println("users.go -> UpdateAskedInDB -> UpdateID: err =", err)
		return false
	}
	user.Asked = asked
	au.UpdateUserInMap(user)
	return true
}

// /////////////////////////////////////////////
// // User manipulations
// /////////////////////////////////////////////

// // Deletes user's token
// func delToken(username string) error {

// 	collection := *Client.Database("courses").Collection("users")
// 	filter := bson.M{"username": username}

// 	update := bson.M{
// 		"$set": bson.M{"token": token{}},
// 	}

// 	deleteResult, err := collection.UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("Failed to delete")
// 	}
// 	log.Println(deleteResult)
// 	return nil
// }

// // Adds token to user
// func addToken(username, tokenStr string, endless bool) error {

// 	collection := *Client.Database("courses").Collection("users")
// 	filter := bson.M{"username": username}

// 	update := bson.M{
// 		"$set": bson.M{"token": token{Name: tokenStr, Endless: endless}},
// 	}

// 	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("Failed to delete")
// 	}
// 	log.Println(updateResult)
// 	return nil
// }

// // Checks if entered username and password are the same in DB
// func checkFromDB(username string, password string) (bool, user) {

// 	// Getting the collection (table)
// 	collection := Client.Database("courses").Collection("users")

// 	// Filter by name
// 	filter := bson.M{"username": username}

// 	var result user

// 	// Writing result of filtration to result var
// 	err := collection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		log.Println(err)
// 		return false, user{}
// 	}

// 	// Checking if entered password is the same with hashed
// 	if CheckPasswordHash(password, result.Password) {
// 		return true, result
// 	}

// 	log.Println("Wrong password")
// 	return false, user{}

// }

// // Adds new user to DB
// func addUserToDB(user user) error {

// 	collection := Client.Database("courses").Collection("users")

// 	// Inserting user to DB
// 	insertResult, err := collection.InsertOne(context.TODO(), user)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	log.Println(insertResult)
// 	return nil
// }

// // Checks if user with this name - exists in DB
// func checkUserExist(username string) bool {

// 	collection := Client.Database("courses").Collection("users")

// 	filter := bson.M{"username": username}

// 	var result user
// 	//var ret bool

// 	err := collection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		log.Println(err)
// 		return true
// 	}

// 	return false
// }

// // Checks if user with this email - exists in DB
// func checkEmailExist(email string) bool {

// 	collection := Client.Database("courses").Collection("users")

// 	filter := bson.M{"email": email}

// 	var result user

// 	err := collection.FindOne(context.TODO(), filter).Decode(&result)
// 	if err != nil {
// 		log.Println(err)
// 		return true
// 	}

// 	return false
// }

// /////////////////////////////////////////////
// // Article manipulations
// /////////////////////////////////////////////

// // Gets articles from DB and return slice of them
// func getArticleFromDB() []article {

// 	collection := Client.Database("courses").Collection("articles")

// 	// Here's an array in which we store the decoded documents
// 	var results []article

// 	cur, err := collection.Find(context.TODO(), bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Finding multiple documents returns a cursor
// 	// Iterating through the cursor allows us to decode documents one at a time
// 	for cur.Next(context.TODO()) {

// 		// Create a value into which the single document can be decoded
// 		var elem article
// 		err := cur.Decode(&elem)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		results = append(results, elem)
// 	}

// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Closing the cursor
// 	cur.Close(context.TODO())

// 	return results
// }

// // Adds article to DB
// func insertArticleToDB(a article) {

// 	collection := Client.Database("courses").Collection("articles")

// 	insertResult, err := collection.InsertOne(context.TODO(), a)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	log.Println(insertResult)
// }

// // Deletes article from DB
// func deleteArticleFromDB(title string) error {

// 	collection := Client.Database("courses").Collection("articles")
// 	filter := bson.M{"title": title}

// 	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("Failed to delete")
// 	}
// 	log.Println(deleteResult)
// 	return nil
// }

// /////////////////////////////////////////////
// // Comment manipulations
// /////////////////////////////////////////////

// // // Gets comments from DB and return slice of them
// // func getCommentFromDB(title string) []comment {

// // 	collection := Client.Database("courses").Collection("articles")

// // 	filter := bson.M{"title": title}
// // 	var result article

// // 	err := collection.FindOne(context.TODO(), filter).Decode(&result)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	var comments []comment
// // 	for _, comm := range result.Comment {
// // 		comments = append(comments, comm)
// // 	}
// // 	return comments

// // }

// // Adds comment to DB
// func commentToDB(comtitle, commentStr, time, name string) error {

// 	collection := Client.Database("courses").Collection("articles")
// 	filter := bson.M{"title": comtitle}

// 	update := bson.M{
// 		"$push": bson.M{"comment": comment{ComTime: time, ComContent: commentStr, ComName: name}},
// 	}

// 	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("Failed to update")
// 	}
// 	log.Println(updateResult)
// 	return nil
// }

// // Deletes all coments in article
// func delComment(comtitle string) error {

// 	collection := *Client.Database("courses").Collection("articles")
// 	filter := bson.M{"title": comtitle}

// 	update := bson.M{
// 		"$set": bson.M{"comment": []comment{}},
// 	}

// 	deleteResult, err := collection.UpdateOne(context.TODO(), filter, update)

// 	if err != nil {
// 		log.Println(err)
// 		return errors.New("Failed to delete")
// 	}
// 	log.Println(deleteResult)
// 	return nil
// }
