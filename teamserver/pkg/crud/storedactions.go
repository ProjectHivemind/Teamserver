package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

func (d *DatabaseModel) AllStoredActions() ([]model.StoredAction, error) {
	var allStoredActions []model.StoredAction

	sqlStatement := `SELECT * FROM public."StoredActions"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var storedAction model.StoredAction

		err = rows.Scan(
			&storedAction.UUID,
			&storedAction.Name,
			&storedAction.ModuleToRun,
			&storedAction.ModuleFunc,
			&storedAction.Arguments,
		)
		if err != nil {
			return nil, err
		}

		allStoredActions = append(allStoredActions, storedAction)
	}

	return allStoredActions, nil
}

func (d *DatabaseModel) GetStoredActionById(id string) (model.StoredAction, error) {
	var storedAction model.StoredAction

	sqlStatement := `SELECT * FROM public."StoredActions" WHERE "UUID"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&storedAction.UUID,
		&storedAction.Name,
		&storedAction.ModuleToRun,
		&storedAction.ModuleFunc,
		&storedAction.Arguments,
	)

	return storedAction, err
}

func (d *DatabaseModel) InsertStoredAction(storedAction model.StoredAction) (bool, error) {
	sqlStatement := `INSERT INTO public."StoredActions"(
		"UUID", "Name", "ModuleToRun", "ModuleFunc", "Arguments")
		VALUES ($1, $2, $3, $4, $5);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		storedAction.UUID,
		storedAction.Name,
		storedAction.ModuleToRun,
		storedAction.ModuleFunc,
		storedAction.Arguments)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteStoredAction(id string) (bool, error) {
	sqlStatement := `DELETE FROM public."StoredActions"
		WHERE "UUID"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}
