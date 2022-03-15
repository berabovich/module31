package entity

type User struct {
	Id      int      `json:"id" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"Name"`
	Age     int      `json:"age" bson:"Age"`
	Friends []string `json:"friends" bson:"Friends"`
}

type Id struct {
	TargetId int `json:"target_id"`
	SourceId int `json:"source_id"`
}

type UserUpgrade struct {
	NewName string `json:"new_name"`
	NewAge  int    `json:"new_age"`
}
