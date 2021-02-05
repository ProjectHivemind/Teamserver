package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
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

func (d *DatabaseModel) GetStagedActionById(id string) (model.StagedActions, error) {
	var stagedAction model.StagedActions

	sqlStatement := `SELECT * FROM public."StagedActions" WHERE id=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&stagedAction.Id,
		&stagedAction.UUIDofAction,
		&stagedAction.UUIDofImplant,
		&stagedAction.TimeStaged,
	)

	return stagedAction, err
}
