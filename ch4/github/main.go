package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number int `json:"number,omitempty"`
	User
	HTMLURL   string    `json:"html_url,omitempty"`
	Title     string    `json:"title"`
	State     string    `json:"state,omitempty"` // "open" ou "closed"
	Body      string    `json:"body"`            // Aqui entra o texto do editor
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type User struct {
	Login   string
	HTMLURL string
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", resp.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

// exercicio 4.10
func DateFormat(date time.Time) string {
	now := time.Now()
	diff := now.Sub(date)
	seconds := int(diff.Seconds())

	if seconds < 0 {
		return "Future"
	}
	switch {
	case seconds < 60:
		return "Now"
	case seconds < (60 * 60):
		minutes := seconds / 60
		return fmt.Sprintf("%d min ago", minutes)
	case seconds < (60 * 60 * 60):
		hours := seconds / (60 * 60)

		return fmt.Sprintf("%d hours", hours)
	case seconds < (60 * 60 * 60 * 30):
		days := seconds / (60 * 60 * 60)
		return fmt.Sprintf("%d days", days)
	case seconds < 31536000:
		months := seconds / 2592000
		if months <= 1 {
			return "More month"
		}
		return fmt.Sprintf("%d months", months)

	default:
		years := seconds / 31536000
		if years <= 1 {
			return "More year"
		}
		return fmt.Sprintf("%d years", years)
	}
}
