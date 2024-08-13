package repositories

import (
    "context"
    "task-manager/Domain"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type UserRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
    return &UserRepository{
        collection: db.Collection("users"),
    }
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
    _, err := r.collection.InsertOne(ctx, user)
    return err
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
    var user domain.User
    err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
    return &user, err
}