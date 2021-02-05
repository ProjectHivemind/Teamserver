package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

func (d *DatabaseModel) AllExecutedActions() ([]model.ExecutedActions, error) {
	var allExecutedActions []model.ExecutedActions

	sqlStatement := `SELECT * FROM public."ExecutedActions"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var executedAction model.ExecutedActions

		err = rows.Scan(
			&executedAction.Id,
			&executedAction.UUIDofAction,
			&executedAction.TimeRan,
			&executedAction.Successful,
			&executedAction.ActionResponse,
		)
		if err != nil {
			return nil, err
		}

		allExecutedActions = append(allExecutedActions, executedAction)
	}

	return allExecutedActions, nil
}

func (d *DatabaseModel) GetExecutedActionById(id string) (model.ExecutedActions, error) {
	var executedAction model.ExecutedActions

	sqlStatement := `SELECT * FROM public."ExecutedActions" WHERE id=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&executedAction.Id,
		&executedAction.UUIDofAction,
		&executedAction.TimeRan,
		&executedAction.Successful,
		&executedAction.ActionResponse,
	)

	return executedAction, err
}
