package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

func (d *DatabaseModel) AllExecutedActions() ([]model.ExecutedAction, error) {
	var allExecutedActions []model.ExecutedAction

	sqlStatement := `SELECT * FROM public."ExecutedActions"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var executedAction model.ExecutedAction

		err = rows.Scan(
			&executedAction.Id,
			&executedAction.UUIDofImplant,
			&executedAction.UUIDofAction,
			&executedAction.TimeSent,
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

func (d *DatabaseModel) GetExecutedActionById(id string) (model.ExecutedAction, error) {
	var executedAction model.ExecutedAction

	sqlStatement := `SELECT * FROM public."ExecutedActions" WHERE "id"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&executedAction.Id,
		&executedAction.UUIDofImplant,
		&executedAction.UUIDofAction,
		&executedAction.TimeSent,
		&executedAction.TimeRan,
		&executedAction.Successful,
		&executedAction.ActionResponse,
	)

	return executedAction, err
}

func (d *DatabaseModel) InsertExecutedAction(executedAction model.ExecutedAction) (bool, error) {
	sqlStatement := `INSERT INTO public."ExecutedActions"(
		id, "UUIDofImplant", "UUIDofAction", "TimeSent", "TimeRan", "Successful", "ActionResponse")
		VALUES ($1, $2, $3, $4, $5, $6, $7);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		executedAction.Id,
		executedAction.UUIDofImplant,
		executedAction.UUIDofAction,
		executedAction.TimeSent,
		executedAction.TimeRan,
		executedAction.Successful,
		executedAction.ActionResponse)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteExecutedAction(id string) (bool, error) {
	sqlStatement := `DELETE FROM public."ExecutedActions"
		WHERE "id"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionResponse(id string, response string) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "ActionResponse"=$2
		WHERE "id"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, response)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionTimeRan(id string, t string) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "TimeRan"=$2
		WHERE "id"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, t)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateExecutedActionSuccessful(id string, successful bool) (bool, error) {
	sqlStatement := `UPDATE public."ExecutedActions"
		SET "Successful"=$2
		WHERE "id"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, successful)

	if err != nil {
		check = false
	}
	return check, err
}
