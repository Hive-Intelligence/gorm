package clause_test

import (
	"fmt"
	"testing"

	"github.com/Hive-Intelligence/gorm/clause"
)

func TestSelect(t *testing.T) {
	results := []struct {
		Clauses []clause.Interface
		Result  string
		Vars    []interface{}
	}{
		{
			[]clause.Interface{clause.Select{}, clause.From{}},
			"SELECT * FROM `users`", nil,
		},
		{
			[]clause.Interface{clause.Select{
				Columns: []clause.Column{clause.PrimaryColumn},
			}, clause.From{}},
			"SELECT `users`.`id` FROM `users`", nil,
		},
		{
			[]clause.Interface{clause.Select{
				Columns: []clause.Column{clause.PrimaryColumn},
			}, clause.Select{
				Columns: []clause.Column{{Name: "name"}},
			}, clause.From{}},
			"SELECT `name` FROM `users`", nil,
		},
	}

	for idx, result := range results {
		t.Run(fmt.Sprintf("case #%v", idx), func(t *testing.T) {
			checkBuildClauses(t, result.Clauses, result.Result, result.Vars)
		})
	}
}
