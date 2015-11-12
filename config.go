package postmi

type Config struct {
	App struct {
		Port          int
		TemplatesPath string
		AssetPath     string
		Prefix        string
	}
	Admin struct {
		Login    string
		Password string
	}
	DataBase struct {
		Driver  string
		Setting string
	}
}
