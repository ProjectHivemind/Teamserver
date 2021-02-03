package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllImplantTypes() ([]model.ImplantType, error) {
	var allImplantTypes []model.ImplantType

	sqlStatement := `SELECT * FROM public."ImplantType"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var implantType model.ImplantType

		err = rows.Scan(
			&implantType.UUID,
			&implantType.ImplantName,
			&implantType.ImplantVersion,
		)
		if err != nil {
			return nil, err
		}

		allImplantTypes = append(allImplantTypes, implantType)
	}

	return allImplantTypes, nil
}

func (d *DatabaseModel) AllImplants() ([]model.Implant, error) {
	var allImplants []model.Implant

	sqlStatement := `SELECT * FROM public."Implant"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var implant model.Implant

		err = rows.Scan(
			&implant.UUID,
			&implant.UUIDImplantType,
			&implant.PrimaryIP,
			&implant.Hostname,
			&implant.MAC,
			&implant.ImplantOS,
			pq.Array(&implant.OtherIPs),
			pq.Array(&implant.SupportedModules),
		)
		if err != nil {
			return nil, err
		}

		allImplants = append(allImplants, implant)
	}

	return allImplants, nil
}

func (d *DatabaseModel) AllModules() ([]model.Modules, error) {
	var allModules []model.Modules

	sqlStatement := `SELECT * FROM public."Modules"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var model model.Modules

		err = rows.Scan(
			&model.ModuleName,
			pq.Array(&model.ModuleFuncNames),
		)
		if err != nil {
			return nil, err
		}

		allModules = append(allModules, model)
	}

	return allModules, nil
}

func (d *DatabaseModel) AllGroups() ([]model.Groups, error) {
	var allGroups []model.Groups

	sqlStatement := `SELECT * FROM public."Groups"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var group model.Groups

		err = rows.Scan(
			&group.UUID,
			&group.GroupName,
			pq.Array(&group.Implants),
		)
		if err != nil {
			return nil, err
		}

		allGroups = append(allGroups, group)
	}

	return allGroups, nil
}

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
