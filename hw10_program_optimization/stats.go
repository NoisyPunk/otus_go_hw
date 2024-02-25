package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users []User

func getUsers(r io.Reader) (result users, err error) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		var user User
		if err = easyjson.Unmarshal(s.Bytes(), &user); err != nil {
			return
		}
		result = append(result, user)
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)

	for _, user := range u {
		if user.Email == "" {
			continue
		}
		mailDomain := strings.SplitAfter(user.Email, "@")
		if strings.HasSuffix(mailDomain[1], domain) {
			result[strings.ToLower(mailDomain[1])]++
		}
	}
	return result, nil
}
