package config

import (
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

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

func (d *DatabaseModel) InsertExecutedAction(executedAction model.ExecutedActions) (bool, error) {
	sqlStatement := `INSERT INTO public."ExecutedActions"(
		id, "UUIDofAction", "TimeRan", "Successful", "ActionResponse")
		VALUES ($1, $2, $3, $4, $5);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		executedAction.Id,
		executedAction.UUIDofAction,
		executedAction.TimeRan,
		executedAction.Successful,
		executedAction.ActionResponse)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteExecutedAction(id int) (bool, error) {
	sqlStatement := `DELETE FROM public."ExecutedActions"
		WHERE id=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionResponse(id int, response string) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "ActionResponse"=$2
		WHERE id=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, response)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionTimeRan(id int, t time.Time) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "TimeRan"=$2
		WHERE id=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, t)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionSuccessful(id int, successful bool) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "Successful"=$2
		WHERE id=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, successful)

	if err != nil {
		check = false
	}
	return check, err
}
