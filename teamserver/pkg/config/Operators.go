package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

func (d *DatabaseModel) AllOperators() ([]model.Operators, error) {
	var allOperators []model.Operators

	sqlStatement := `SELECT * FROM public."Operators";`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var operator model.Operators

		err = rows.Scan(
			&operator.Username,
			&operator.Password,
			&operator.Permission,
		)
		if err != nil {
			return nil, err
		}

		allOperators = append(allOperators, operator)
	}

	return allOperators, nil
}
