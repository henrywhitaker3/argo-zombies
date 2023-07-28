package dashboard

var (
	title  string   = "Argo Zombies Dashboard"
	labels []string = []string{"gitops/argo-zombies"}
)

type CreateOrUpdate func() error

type Dashboard struct{}

type GithubDashboard struct {
	Token string `yaml:"token"`
	Repo  string `yaml:"repo"`
}
