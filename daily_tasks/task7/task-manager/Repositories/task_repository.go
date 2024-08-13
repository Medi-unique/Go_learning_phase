package repositories

import (
    "context"
    "task-manager/Domain"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)

type TaskRepository struct {
    collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) *TaskRepository {
    return &TaskRepository{
        collection: db.Collection("tasks"),
    }
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) error {
    _, err := r.collection.InsertOne(ctx, task)
    return err
}

func (r *TaskRepository) GetTasksByUserID(ctx context.Context, userID string) ([]*domain.Task, error) {
    var tasks []*domain.Task
    cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var task domain.Task
        if err := cursor.Decode(&task); err != nil {
            return nil, err
        }
        tasks = append(tasks, &task)
    }

    return tasks, nil
}