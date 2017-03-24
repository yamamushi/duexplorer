package models

type Users [] struct {
	/*ID struct {
		Oid string `bson:"$oid" json:"$oid"`
	} `json:"_id"` */
	User         string `bson:"user" json:"user"`
	CreatedDate  string `bson:"createdDate" json:"createdDate"`
	PledgeStatus string `bson:"pledgeStatus" json:"pledgeStatus"`
	Organizations []struct {
		Orgname  string `bson:"orgname" json:"orgname"`
		JoinDate string `bson:"joinDate" json:"joinDate"`
		Status   string `bson:"status" json:"status"`
	} `bson:"organizations" json:"organizations"`
}