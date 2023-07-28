package dashboard

type CreateOrUpdate func() error

type Dashboard struct{}

type GithubDashboard struct {
	Token string `yaml:"token"`
	Repo  string `yaml:"repo"`
}
