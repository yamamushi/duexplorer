package routers

type Users struct {
	Users []struct {
		CreatedDate   string `json:"createdDate"`
		Organizations struct {
			GalacticCourierServices struct {
				JoinDate string `json:"joinDate"`
				Status   string `json:"status"`
			} `json:"galactic-courier-services"`
			TheOrderOfMeru struct {
				JoinDate string `json:"joinDate"`
				Status   string `json:"status"`
			} `json:"the-order-of-meru"`
		} `json:"organizations"`
		PledgeStatus string `json:"pledgeStatus"`
		User         string `json:"user"`
	} `json:"users"`
}