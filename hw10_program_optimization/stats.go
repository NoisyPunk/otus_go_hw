package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

var (
	ErrEmptyDomain = fmt.Errorf("empty domain")
	ErrEmptyData   = fmt.Errorf("empty data")
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
	u, err := getEmails(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type emails []string

func getEmails(r io.Reader) (result emails, err error) {
	if r == nil {
		return nil, ErrEmptyData
	}
	s := bufio.NewScanner(r)
	user := User{}

	for s.Scan() {
		if err = easyjson.Unmarshal(s.Bytes(), &user); err != nil {
			return
		}
		result = append(result, user.Email)
	}
	return
}

func countDomains(u emails, domain string) (DomainStat, error) {
	result := make(DomainStat)
	if domain == "" {
		return nil, ErrEmptyDomain
	}

	for _, email := range u {
		if email == "" {
			continue
		}
		mailDomain := strings.SplitAfter(email, "@")
		if strings.HasSuffix(mailDomain[1], domain) {
			result[strings.ToLower(mailDomain[1])]++
		}
	}
	return result, nil
}
