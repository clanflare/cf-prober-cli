package types

type Manifest struct {
	Token      string            `json:"token"`
	Client_Uri string            `json:"client_uri"`
	Repos      map[string]string `json:"repos"`
	Projects   []Project         `json:"projects"`
}

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
