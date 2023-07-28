package dashboard

var (
	title  string   = "Argo Zombies Dashboard"
	labels []string = []string{"gitops/argo-zombies"}
)

type GithubDashboard struct {
	Enabled bool   `yaml:"enabled"`
	Token   string `yaml:"token"`
	Repo    string `yaml:"repo"`
}
