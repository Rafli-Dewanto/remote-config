package model

type FirebaseRemoteConfig struct {
	Conditions map[string]Condition `json:"conditions"`
	Parameters map[string]Parameter `json:"parameters"`
	Etag       string               `json:"etag"`
	Version    Version              `json:"version"`
}

type Condition struct {
	Name       string   `json:"name"`
	Expression string   `json:"expression"`
	TagColor   string   `json:"tagColor"`
}

type Parameter struct {
	DefaultValue DefaultValue `json:"defaultValue"`
	Description string        `json:"description"`
	ValueType   string        `json:"valueType"`
}

type DefaultValue struct {
	Value string `json:"value"`
}

type Version struct {
	VersionNumber string     `json:"versionNumber"`
	UpdateOrigin  string     `json:"updateOrigin"`
	UpdateType    string     `json:"updateType"`
	UpdateUser    UpdateUser `json:"updateUser"`
	UpdateTime    string     `json:"updateTime"`
}

type UpdateUser struct {
	Email string `json:"email"`
}
