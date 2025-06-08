package simple

type Configuration struct {
	Name string
}

type Application struct {
	Config *Configuration
}

func NewApplication() *Application {
	return &Application{
		Config: &Configuration{
			Name: "Simple Application",
		},
	}
}