package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

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
