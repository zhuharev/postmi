package postmi

type Config struct {
	App struct {
		Port          int
		TemplatesPath string
		Prefix        string
	}
}
