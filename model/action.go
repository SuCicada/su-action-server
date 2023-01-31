package model

type Action struct {
	Github *GithubContext
	Env    map[string]string
	Job    *JobContext
}
