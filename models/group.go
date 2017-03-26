package models

type Group struct {
	CreatedDate string `bson:"createdDate" json:"createdDate"`
	Members     []struct {
		JoinDate string `bson:"joinDate" json:"joinDate"`
		Name     string `bson:"name" json:"name"`
		Status   string `bson:"status" json:"status"`
	} `bson:"members" json:"members"`
	Name string `bson:"name" json:"name"`
}
