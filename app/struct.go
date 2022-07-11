package app

type CommentStruct struct {
	Type string `json:"type"`
	Data struct {
		Comment struct {
			Nick       string `json:"nick"`
			Mail       string `json:"mail"`
			Url        string `json:"url"`
			Comment    string `json:"comment"`
			RawComment string `json:"rawComment"`
			Ip         string `json:"ip"`
			InsertedAt string `json:"insertedAt"`
			Status     string `json:"status"`
		} `json:"comment"`
		Reply struct {
			Nick       string `json:"nick"`
			Mail       string `json:"mail"`
			Comment    string `json:"comment"`
			InsertedAt string `json:"insertedAt"`
			Status     string `json:"status"`
		} `json:"reply"`
	} `json:"data"`
}
