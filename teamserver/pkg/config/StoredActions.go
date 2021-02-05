package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

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
		)
		if err != nil {
			return nil, err
		}

		allStoredActions = append(allStoredActions, storedAction)
	}

	return allStoredActions, nil
}
