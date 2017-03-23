package models

type UserModel struct {
	Users []struct {
		CreatedDate   string `json:"createdDate"`
		Organizations []struct {
			JoinDate string `json:"joinDate"`
			Orgname  string `json:"orgname"`
			Status   string `json:"status"`
		} `json:"organizations"`
		PledgeStatus string `json:"pledgeStatus"`
		User         string `json:"user"`
	} `json:"users"`
}
