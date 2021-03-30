package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllStagedActions() ([]model.StagedActions, error) {
	var allStagedActions []model.StagedActions

	sqlStatement := `SELECT * FROM public."StagedActions"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var stagedAction model.StagedActions

		err = rows.Scan(
			&stagedAction.Id,
			&stagedAction.UUIDofAction,
			&stagedAction.UUIDofImplant,
			&stagedAction.TimeStaged,
		)
		if err != nil {
			return nil, err
		}

		allStagedActions = append(allStagedActions, stagedAction)
	}

	return allStagedActions, nil
}

func (d *DatabaseModel) AllStagedActionsFrontend() ([]model.StagedActionsFrontend, error) {
	var allStagedActions []model.StagedActionsFrontend

	sqlStatement := `SELECT * FROM public."StagedActions" JOIN public."Implant" ON public."Implant"."UUID" = public."StagedActions"."UUIDofImplant" JOIN public."StoredActions" ON public."StoredActions"."UUID" = public."StagedActions"."UUIDofAction"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var stagedAction model.StagedActionsFrontend

		err = rows.Scan(
			&stagedAction.StagedAction.Id,
			&stagedAction.StagedAction.UUIDofAction,
			&stagedAction.StagedAction.UUIDofImplant,
			&stagedAction.StagedAction.TimeStaged,
			&stagedAction.Implant.UUID,
			&stagedAction.Implant.UUIDImplantType,
			&stagedAction.Implant.PrimaryIP,
			&stagedAction.Implant.Hostname,
			&stagedAction.Implant.MAC,
			&stagedAction.Implant.ImplantOS,
			pq.Array(&stagedAction.Implant.OtherIPs),
			pq.Array(&stagedAction.Implant.SupportedModules),
			&stagedAction.StoredAction.UUID,
			&stagedAction.StoredAction.ModuleToRun,
			&stagedAction.StoredAction.ModuleFunc,
			pq.Array(&stagedAction.StoredAction.Arguments),
		)
		if err != nil {
			return nil, err
		}

		allStagedActions = append(allStagedActions, stagedAction)
	}

	return allStagedActions, nil
}

func (d *DatabaseModel) GetStagedActionById(id string) (model.StagedActions, error) {
	var stagedAction model.StagedActions

	sqlStatement := `SELECT * FROM public."StagedActions" WHERE "id"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&stagedAction.Id,
		&stagedAction.UUIDofAction,
		&stagedAction.UUIDofImplant,
		&stagedAction.TimeStaged,
	)

	return stagedAction, err
}

func (d *DatabaseModel) GetStagedActionByImplant(id string) ([]model.StagedActions, error) {
	var allStagedActions []model.StagedActions

	sqlStatement := `SELECT * FROM public."StagedActions" WHERE "UUIDofImplant"=$1`

	rows, err := d.db.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var stagedAction model.StagedActions

		err = rows.Scan(
			&stagedAction.Id,
			&stagedAction.UUIDofAction,
			&stagedAction.UUIDofImplant,
			&stagedAction.TimeStaged,
		)
		if err != nil {
			return nil, err
		}

		allStagedActions = append(allStagedActions, stagedAction)
	}

	return allStagedActions, nil
}

func (d *DatabaseModel) InsertStagedAction(stagedAction model.StagedActions) (bool, error) {
	sqlStatement := `INSERT INTO public."StagedActions"(
		"id", "UUIDofAction", "UUIDofImplant", "TimeStaged")
		VALUES ($1, $2, $3, $4);`

	check := true

	_, err := d.GetImplantById(stagedAction.UUIDofImplant)
	if err != nil {
		return false, err
	}

	_, err = d.db.Exec(sqlStatement,
		stagedAction.Id,
		stagedAction.UUIDofAction,
		stagedAction.UUIDofImplant,
		stagedAction.TimeStaged)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteStagedAction(id string) (bool, error) {
	sqlStatement := `DELETE FROM public."StagedActions"
		WHERE "id"=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}
