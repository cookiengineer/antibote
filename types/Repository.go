package types

type Repository struct {
	Identifier  uint64 `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	IsFork      bool   `json:"fork"`
	Created     string `json:"created_at"`
	Updated     string `json:"updated_at"`
	Pushed      string `json:"pushed_at"`
	Owner       struct {
		Identifier uint64 `json:"id"`
		Name       string `json:"login"`
		Type       string `json:"type"`
	} `json:"owner"`
	Commits     map[string]*Commit `json:"commits"`
}

func NewRepository(name string) Repository {

	var repo Repository

	repo.Name = name
	repo.Commits = make(map[string]*Commit)

	return repo

}

func (repo *Repository) AddCommit(commit Commit) {
	repo.Commits[commit.Hash] = &commit
}
