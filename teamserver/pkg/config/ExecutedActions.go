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
