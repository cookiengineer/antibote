package constants

var GitHubKeys []string = []string{
	"4AEE18F83AFDEB23", // GitHub Web UI, 2023
	"B5690EEEBB952194", // GitHub Web UI, 2024
}

func IsGitHubKey(keyid string) bool {

	found := false

	for k := 0; k < len(GitHubKeys); k++ {

		if keyid == GitHubKeys[k] {
			found = true
			break
		}

	}

	return found

}
