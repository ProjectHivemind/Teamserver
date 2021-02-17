package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

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

func (d *DatabaseModel) GetOperatorByUsername(username string) (model.Operators, error) {
	var operator model.Operators

	sqlStatement := `SELECT * FROM public."Operators" WHERE "Username"=$1`

	row := d.db.QueryRow(sqlStatement, username)
	err := row.Scan(
		&operator.Username,
		&operator.Password,
		&operator.Permission,
	)

	return operator, err
}

func (d *DatabaseModel) InsertOperator(operator model.Operators) (bool, error) {
	sqlStatement := `INSERT INTO public."Operators"(
		"Username", "Password", "Permission")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		operator.Username,
		operator.Password,
		operator.Permission)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteOperator(username string) (bool, error) {
	sqlStatement := `DELETE FROM public."Operators"
		WHERE "Username"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, username)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) ChangeOperatorPassword(username string, password string) (bool, error) {
	sqlStatement := `UPDATE public."Operators"
		SET "Password"=$2
		WHERE "Username"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, username, password)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) ChangeOperatorPermission(username string, permission int) (bool, error) {
	sqlStatement := `UPDATE public."Operators"
		SET "Permission"=$2
		WHERE "Username"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, username, permission)

	if err != nil {
		check = false
	}
	return check, err
}
