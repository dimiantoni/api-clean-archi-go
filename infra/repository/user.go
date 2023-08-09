package repository

import (
	"context"
	"fmt"

	"github.com/dimiantoni/api-clean-archi-go/domain/entity"
	"github.com/dimiantoni/api-clean-archi-go/infra/database"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Create(e *entity.User) (entity.ID, error)
}

type UserRepositoryMongoDb struct {
	Db *database.MongodbDatabase
}

func NewUserRepository(db *database.MongodbDatabase) UserRepositoryMongoDb {
	repository := UserRepositoryMongoDb{db}
	repository.createIndexes()
	return repository
}

func (repository UserRepositoryMongoDb) createIndexes() error {
	_, err := repository.getCollection().Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{Keys: bson.D{{"email", 1}}},
		{Keys: bson.D{{"name", 1}}},
	}, &options.CreateIndexesOptions{})

	if err != nil {
		return err
	}
	return nil
}

func (repository UserRepositoryMongoDb) getCollection() *mongo.Collection {
	return repository.Db.Client.Database("local_dev_database").Collection("users")
}

// Create an user
func (repository UserRepositoryMongoDb) Create(e *entity.User) (entity.ID, error) {
	fmt.Println("create user here")
	// Create a user
	user := bson.M{
		"_id":      primitive.NewObjectID(),
		"name":     e.Name,
		"email":    e.Email,
		"password": e.Password,
		"address":  e.Address,
		"age":      e.Age,
	}
	// Insert the user into the collection
	_, err := repository.getCollection().InsertOne(context.Background(), user)

	log.Infoln("User CREATED", e.Name)

	if err != nil {
		log.Errorln(err)
		return e.ID, err
	}

	return e.ID, nil
}

func (repository UserRepositoryMongoDb) Get(id entity.ID) (*entity.User, error) {
	// return getUser(id)
	return &entity.User{}, nil
}

func (repository UserRepositoryMongoDb) getUser(id entity.ID) (*entity.User, error) {
	fmt.Println("get user here")
	return nil, nil
}

// Update an user
func (repository UserRepositoryMongoDb) Update(e *entity.User) error {
	fmt.Println("update user here")
	return nil
}

// Search users
func (repository UserRepositoryMongoDb) Search(param string) ([]*entity.User, error) {

	query := bson.M{
		"email": param,
	}

	result, err := repository.getCollection().Find(context.Background(), query)
	if err != nil {
		log.Errorln(err)
		res := make([]*entity.User, 0)
		return res, err
	}

	// Loop over the results
	users := []*entity.User{}
	for result.Next(context.TODO()) {
		elem := bson.M{}
		pos := &entity.User{}
		if err := result.Decode(&elem); err != nil {
			return nil, err
		}
		err = repository.transformBsonToModel(elem, &pos)
		if err != nil {
			return nil, err
		}
		users = append(users, pos)
	}
	return users, nil
}

func (repository UserRepositoryMongoDb) transformBsonToModel(b interface{}, m interface{}) error {
	doc, err := bson.Marshal(b)
	if err != nil {
		return err
	}
	return bson.Unmarshal(doc, m)
}

// List users
func (repository UserRepositoryMongoDb) List() ([]*entity.User, error) {
	result, err := repository.getCollection().Find(context.Background(), bson.M{})
	users := []*entity.User{}
	for result.Next(context.TODO()) {
		elem := bson.M{}
		pos := &entity.User{}
		if err := result.Decode(&elem); err != nil {
			return nil, err
		}
		err = repository.transformBsonToModel(elem, &pos)
		if err != nil {
			return nil, err
		}
		users = append(users, pos)
	}
	return users, nil
}

// Delete an user
func (repository UserRepositoryMongoDb) Delete(id entity.ID) error {
	_, err := repository.getCollection().DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
