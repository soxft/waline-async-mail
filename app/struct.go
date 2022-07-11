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
			InsertedAt string `json:"inserted_at"`
			Status     string `json:"status"`
		} `json:"comment"`
		Reply struct {
			Comment    string `json:"comment"`
			InsertedAt string `json:"inserted_at"`
			Status     string `json:"status"`
			Mail       string `json:"mail"`
		} `json:"reply"`
	} `json:"data"`
}
