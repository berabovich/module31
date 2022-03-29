package entity

type User struct {
	Id      string   `json:"id" bson:"_id,omitempty"`
	Name    string   `json:"name" bson:"Name"`
	Age     int      `json:"age" bson:"Age"`
	Friends []string `json:"friends" bson:"Friends"`
}

type Id struct {
	TargetId string `json:"target_id"`
	SourceId string `json:"source_id"`
}

type UserUpgrade struct {
	NewName string `json:"new_name"`
	NewAge  int    `json:"new_age"`
}
