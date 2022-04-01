package databases

import (
	"fmt"
)

type (
	InsertContentRequest struct {
		Content string
	}
)

func InsertContent(r *InsertContentRequest) error {
	_, err := GetMysql().Exec("INSERT INTO content (content) VALUES (?)", r.Content)
	if err != nil {
		return fmt.Errorf("InsertContent: %v", err)
	}

	return nil
}
