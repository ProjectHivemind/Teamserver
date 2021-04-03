package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
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
			&executedAction.Stager,
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

func (d *DatabaseModel) AllExecutedActionsFrontend() ([]model.ExecutedActionsFrontend, error) {
	var allExecutedActions []model.ExecutedActionsFrontend

	sqlStatement := `SELECT * FROM public."ExecutedActions" JOIN public."Implant" ON public."Implant"."UUID" = public."ExecutedActions"."UUIDofImplant" JOIN public."StoredActions" ON public."StoredActions"."UUID" = public."ExecutedActions"."UUIDofAction"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var executedAction model.ExecutedActionsFrontend

		err = rows.Scan(
			&executedAction.ExecutedAction.Id,
			&executedAction.ExecutedAction.Stager,
			&executedAction.ExecutedAction.UUIDofImplant,
			&executedAction.ExecutedAction.UUIDofAction,
			&executedAction.ExecutedAction.TimeSent,
			&executedAction.ExecutedAction.Successful,
			&executedAction.ExecutedAction.ActionResponse,
			&executedAction.Implant.UUID,
			&executedAction.Implant.UUIDImplantType,
			&executedAction.Implant.PrimaryIP,
			&executedAction.Implant.Hostname,
			&executedAction.Implant.MAC,
			&executedAction.Implant.ImplantOS,
			pq.Array(&executedAction.Implant.OtherIPs),
			pq.Array(&executedAction.Implant.SupportedModules),
			&executedAction.StoredAction.UUID,
			&executedAction.StoredAction.ModuleToRun,
			&executedAction.StoredAction.ModuleFunc,
			&executedAction.StoredAction.Arguments,
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
		&executedAction.Stager,
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
		id, "Stager", "UUIDofImplant", "UUIDofAction", "TimeSent", "TimeRan", "Successful", "ActionResponse")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		executedAction.Id,
		executedAction.Stager,
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
