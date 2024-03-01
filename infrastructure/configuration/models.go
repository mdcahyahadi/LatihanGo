package configuration

// AppConfig struct model, the base configuration struct for this web service
type AppConfig struct {
	App       App
	Postgre   Postgre
	SlaveDB   SlaveDB
	Statement Statement
}

// App struct model, the configuration file related to basic webservice info such as name environment
type App struct {
	Name        string
	Environment string
	Debug       bool
	Host        string
	Port        string
	Protocol    string
}

// Postgre struct model, this is configuration for postgresql, storing relational data
type Postgre struct {
	Connection string
	Name       string
	User       string
	Ps         string
	Port       string
	Host       string
}

// Postgre struct model, this is configuration for postgresql, storing relational data
type SlaveDB struct {
	Connection string
	Name       string
	User       string
	Ps         string
	Port       string
	Host       string
}

type Statement struct {
	Limit int
}
