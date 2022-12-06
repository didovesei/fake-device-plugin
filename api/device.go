package api

type Device struct {
	id           string `yaml:"id"`
	resourceType string `yaml:"resourceType"`
}

type Empty struct{}
