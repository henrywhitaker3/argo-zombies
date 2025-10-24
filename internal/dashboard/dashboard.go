// Package dashboard
package dashboard

var (
	title  string   = "Argo Zombies Dashboard"
	labels []string = []string{"gitops/argo-zombies"}
)

type GithubDashboard struct {
	Enabled        bool   `yaml:"enabled"`
	Token          string `yaml:"token"`
	ClientID       string `yaml:"client_id"`
	InstallationID int    `yaml:"installation_id"`
	PrivateKey     string `yaml:"private_key"`
	Repo           string `yaml:"repo"`
}

type GitlabDashboard struct {
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Repo    string `yaml:"repo"`
}
