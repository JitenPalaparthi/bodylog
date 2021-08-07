package database

import (
	"context"
	"errors"
	"time"

	"github.com/jitenpalaparthi/bodylog/models"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserDB is to maintain database related methods
type UserDB struct {
	DB *Database
}

const (
	// ErrUserExists is an error constant
	ErrUserExists        = "user already exists or already registered with the given mobile number or email address"
	ErrUserDoesNotExists = "user does not exist"
	ErrWrongInputType    = "wrong input type"
)

//IsUserExists is to check vendor exists or not
func (u *UserDB) IsUserExists(email *string) bool {
	if email == nil {
		return false
	}
	if *email == "" {
		return false
	}
	filter := make(map[string]interface{})
	filter["email"] = email
	count, err := u.DB.GetCount("users", filter)
	if err != nil {
		if err.Error() == "not found" {
			return false
		}
	}
	if count > 0 {
		return true
	}
	return false
}

//IsUserExistsForLogin is to check user exists or not
func (u *UserDB) IsUserExistsForLogin(userLogin *models.User) bool {
	filter := make(map[string]interface{})
	filter["email"] = userLogin.Email
	filter["password"] = userLogin.Password
	filter["status"] = "active"
	count, err := u.DB.GetCount("users", filter)
	if err != nil {
		if err.Error() == "not found" {
			return false
		}
	}
	if count > 0 {
		return true
	}
	return false
}

//IsUserExistsForLogin is to check user exists or not
func (u *UserDB) IsUserActive(userLogin *models.User) error {
	filter := make(map[string]interface{})
	filter["email"] = userLogin.Email
	filter["password"] = userLogin.Password
	filter["status"] = userLogin.Status
	count, err := u.DB.GetCount("users", filter)
	if err != nil {
		if err.Error() == "not found" {
			return err
		}
	}
	if count > 0 {
		return nil
	}
	return nil
}

// //GetUser is to fetch user based on email and mobile as inputs
func (u *UserDB) GetUser(email, mobile *string) (*models.User, error) {
	user := &models.User{}
	myData, err := u.DB.FindRecord("users", primitive.M{"mobile": mobile, "email": email})
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(myData, &user); err != nil {
		return nil, err
	}
	//_id := myData.(map[string]interface{})["_id"].(primitive.ObjectID).Hex()
	//user.ID = _id
	return user, nil
}

// Register is to register a user to the system
func (u *UserDB) Register(user *models.User) error {
	if !u.IsUserExists(user.Email) {
		if user.Role == nil {
			role := "anonymous"
			user.Role = &role
		}

		if _, err := u.DB.InsertRecord("users", user); err != nil {
			return err
		}

	} else {
		return errors.New(ErrUserExists)
	}

	return nil
}

// //Signin is used to sings a user into the system
func (u *UserDB) Signin(userLogin *models.User) bool {
	return u.IsUserExistsForLogin(userLogin)
}

// // ResetPassword is to reset the password.
func (u *UserDB) ResetPassword(ResetPassword *models.UserPasswodReset) error {
	if u.IsUserExists(ResetPassword.Email) {
		//Todo change the code after verification code logic has been brought
		if *ResetPassword.VerifyCode == "00000" {
			filter := make(map[string]interface{})
			filter["email"] = ResetPassword.Email
			data := make(map[string]interface{})
			data["password"] = ResetPassword.Password

			if _, err := u.DB.UpdateRecord("users", filter, data); err != nil {
				return err
			}
		} else {
			return errors.New("invalid Verification Code")
		}
	} else {
		return errors.New(ErrUserExists)
	}
	return nil
}

// //GetOrganization is to fetch individual based on email  as inputs
func (u *UserDB) GetUserBy(email *string) (*models.User, error) {
	user := &models.User{}
	filter := make(map[string]interface{})
	filter["email"] = email
	myData, err := u.DB.FindRecord("users", filter)
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(myData, &user); err != nil {
		return nil, err
	}
	return user, nil
}
func (u *UserDB) GetUsers(skip int64, limit int64, selector interface{}) ([]models.User, error) {
	if _, ok := selector.(map[string]interface{}); !ok {
		return nil, errors.New("invalid input type")
	}
	var result []models.User
	filter := make(map[string]interface{})
	for k, _ := range selector.(map[string]interface{}) {
		if string(k[0]) == "_" {

			objID, err := primitive.ObjectIDFromHex(selector.(map[string]interface{})[k].(string))
			if err != nil {
				return nil, err
			}
			filter[k] = objID
		} else {
			filter[k] = selector.(map[string]interface{})[k]
		}
	}
	colleection := u.DB.Client.(*mongo.Client).Database(u.DB.Name).Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	findOptions := options.Find()
	findOptions.SetLimit(limit).SetSkip(skip)
	cur, err := colleection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	result = make([]models.User, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		iresult := models.User{}
		err := cur.Decode(&iresult)
		if err != nil {
			return nil, err
		}
		result = append(result, iresult)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserDB) UpdateById(id string, data interface{}) error {
	if _, ok := data.(map[string]interface{}); !ok {
		return errors.New(ErrWrongInputType)
	}

	if _, err := u.DB.UpdateRecordByID("users", id, data); err != nil {
		return err
	}
	return nil
}

// func (u *UserDB) GetSummary() (map[string]interface{} /* []struct{ string int },*/, error) {
// 	data := make(map[string]interface{})
// 	no_of_projects, err := u.DB.GetCount("templates", primitive.M{"status": "active"})
// 	if err != nil {
// 		no_of_projects = 0
// 	}
// 	data["no_of_projects"] = no_of_projects
// 	no_of_users, err := u.DB.GetCount("users", primitive.M{"status": "active"})
// 	if err != nil {
// 		no_of_users = 0
// 	}
// 	data["no_of_users"] = no_of_users

// 	//matchStage := primitive.D{{"$match", primitive.D{{"status", "active"}}}}
// 	loopUpStage1 := primitive.D{{"$lookup", primitive.M{"from": "templates", "localField": "_templateId", "foreignField": "_id", "as": "templates"}}}

// 	groupStage := primitive.D{{"$group", primitive.D{{"_id", primitive.D{{"project", "$templates.project"}}}, {"count", primitive.D{{"$sum", 1}}}}}}

// 	collection := u.DB.Client.(*mongo.Client).Database(u.DB.Name).Collection("projectData")
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
// 	cur, _ := collection.Aggregate(ctx, mongo.Pipeline{loopUpStage1, groupStage})
// 	result := []primitive.M{}
// 	//results := []models.ProjectDataGroup{}

// 	err = cur.All(ctx, &result)

// 	for _, v := range result {
// 		results = append(results, models.ProjectDataGroup{v["_id"].(primitive.M)["project"].(primitive.A)[0], v["count"]})
// 	}

// 	fmt.Println(results)

// 	data["projectData_groupby_project_count"] = results
// 	return data, nil
// }
