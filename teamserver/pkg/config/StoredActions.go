package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllStoredActions() ([]model.StoredActions, error) {
	var allStoredActions []model.StoredActions

	sqlStatement := `SELECT * FROM public."StoredActions"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var storedAction model.StoredActions

		err = rows.Scan(
			&storedAction.UUID,
			&storedAction.ModuleToRun,
			&storedAction.ModuleFunc,
			pq.Array(&storedAction.Arguments),
		)
		if err != nil {
			return nil, err
		}

		allStoredActions = append(allStoredActions, storedAction)
	}

	return allStoredActions, nil
}

func (d *DatabaseModel) GetStoredActionById(id string) (model.StoredActions, error) {
	var storedAction model.StoredActions

	sqlStatement := `SELECT * FROM public."StagedActions" WHERE id=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&storedAction.UUID,
		&storedAction.ModuleToRun,
		&storedAction.ModuleFunc,
		pq.Array(&storedAction.Arguments),
	)

	return storedAction, err
}

func (d *DatabaseModel) InsertStoredAction(storedAction model.StoredActions) (bool, error) {
	sqlStatement := `INSERT INTO public."StoredActions"(
		"UUID", "ModuleToRun", "ModuleFunc", "Arguments")
		VALUES ($1, $2, $3, $4);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		storedAction.UUID,
		storedAction.ModuleToRun,
		storedAction.ModuleFunc,
		pq.Array(storedAction.Arguments))

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteStoredAction(id string) (bool, error) {
	sqlStatement := `DELETE FROM public."StoredActions"
		WHERE UUID=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}
