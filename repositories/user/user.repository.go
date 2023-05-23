package user

import (
	"context"

	"github.com/rohitq50/go-lang-project/config"
	model "github.com/rohitq50/go-lang-project/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {

	// Create, will perform db opration to save user
	// Returns modified user and error if occurs
	Create(context.Context, *model.User) error

	// FildAll, returns all users in the system
	// It will return error also if occurs
	FindAll(context.Context) ([]*model.User, error)

	// FindOneById, find the user by the provided id
	// return matched user and error if any
	// FindOneById(context.Context, string) (*model.User, error)

	// Update, will update user data by id
	// return error if any
	Update(context.Context, interface{}, interface{}) error

	// Delete, will remove user entry from DB
	// Return error if any
	Delete(context.Context, *model.User) error

	// FindOne, will find one entry of user matched by the query
	// Query object is an interface type that can accept any object
	// return matched user and error if any
	FindOne(context.Context, interface{}) (*model.User, error)

	IsUserAlreadyExists(context.Context, string) bool
}

type UserRepositoryImp struct {
	db     *mongo.Client
	config *config.Configuration
}

func New(db *mongo.Client, c *config.Configuration) UserRepository {
	return &UserRepositoryImp{db: db, config: c}
}

func (service *UserRepositoryImp) Create(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserRepositoryImp) FindAll(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (service *UserRepositoryImp) Update(ctx context.Context, query, change interface{}) error {
	return nil
}

// func (service *UserRepositoryImp) FindOneById(ctx context.Context, id string) (*model.User, error) {
// 	var user model.User
// 	e := service.collection().Find({"_id": id}).Select({"password": 0, "salt": 0}).One(&user)
// 	return &user, e
// }

func (service *UserRepositoryImp) Delete(ctx context.Context, user *model.User) error {
	return nil
}

func (service *UserRepositoryImp) FindOne(ctx context.Context, query interface{}) (*model.User, error) {
	return nil, nil
}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *UserRepositoryImp) IsUserAlreadyExists(ctx context.Context, email string) bool {
	// query := {"email": email}
	// _, e := service.FindOne(ctx, query)
	// if e != nil {
	// 	return false
	// }
	return true
}

func (service *UserRepositoryImp) collection() *mongo.Collection {
	return service.db.Database(service.config.DBName).Collection("users")
}
