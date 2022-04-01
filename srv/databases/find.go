package databases

import (
	"fmt"
)

type (
	Contents struct {
		Contents []Content
	}

	Content struct {
		Content string
	}
)

func FindAllContent() (*Contents, error) {
	q := `
	SELECT content.content FROM content;
	`
	rows, err := GetMysql().Query(q)
	if err != nil {
		return nil, fmt.Errorf("FindAllContent.Query: %v", err)
	}

	result := Contents{}
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			return nil, fmt.Errorf("FindAllContent.Scan: %v", err)
		}
		if len(result.Contents) != 0 {
			result.Contents = append(result.Contents, Content{content})
		} else {
			result.Contents = []Content{{content}}
		}
	}

	return &result, nil
}
