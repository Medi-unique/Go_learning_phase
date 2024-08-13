package domain

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Username string `json:"username" bson:"username"`
    Password string `json:"password" bson:"password"`
}

type Task struct {
    ID          string `json:"id" bson:"_id,omitempty"`
    Title       string `json:"title" bson:"title"`
    Description string `json:"description" bson:"description"`
    UserID      string `json:"user_id" bson:"user_id"`
}