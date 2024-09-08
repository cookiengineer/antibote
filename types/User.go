package types

import "antibote/constants"
import "antibote/gpg"

type User struct {
	Identifier   uint64                 `json:"id"`
	Name         string                 `json:"login"`
	Aliases      []string               `json:"aliases"`
	Emails       []string               `json:"emails"`
	Repositories map[string]*Repository `json:"repositories"`
	Type         string                 `json:"type"`
}

func NewUser(name string) User {

	var user User

	user.Name = name
	user.Aliases = make([]string, 0)
	user.Emails = make([]string, 0)
	user.Repositories = make(map[string]*Repository)

	return user

}

func (user *User) AddAlias(value string) {

	var found bool = false

	for a := 0; a < len(user.Aliases); a++ {

		if user.Aliases[a] == value {
			found = true
			break
		}

	}

	if found == false {
		user.Aliases = append(user.Aliases, value)
	}

}

func (user *User) AddEmail(value string) {

	var found bool = false

	for e := 0; e < len(user.Emails); e++ {

		if user.Emails[e] == value {
			found = true
			break
		}

	}

	if found == false {
		user.Emails = append(user.Emails, value)
	}

}

func (user *User) TrackRepository(value *Repository) {
	user.Repositories[value.Name] = value
}

func (user *User) GetRepository(value string) *Repository {
	return user.Repositories[value]
}

func (user *User) HasRepository(value string) bool {

	_, ok := user.Repositories[value]

	if ok {
		return true
	}

	return false

}

func (user *User) ToKeys() []Key {

	result := make([]Key, 0)

	for _, repository := range user.Repositories {

		for _, commit := range repository.Commits {

			author := commit.Commit.Author.Name
			email := commit.Commit.Author.Email
			verification := commit.Commit.Verification

			if verification.Verified == true && verification.Reason == "valid" {

				keyid := gpg.ToKeyID(verification.Signature)

				if keyid != "" && !constants.IsGitHubKey(keyid) {

					result = append(result, Key{
						ID:    keyid,
						Name:  author,
						Email: email,
					})

				}

			}

		}

	}


	return result

}

