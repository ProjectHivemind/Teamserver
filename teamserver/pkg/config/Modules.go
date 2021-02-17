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

func (d *DatabaseModel) GetModuleByName(name string) (model.Modules, error) {
	var module model.Modules

	sqlStatement := `SELECT * FROM public."Modules" WHERE "ModuleName"=$1`

	row := d.db.QueryRow(sqlStatement, name)
	err := row.Scan(
		&module.ModuleName,
		pq.Array(&module.ModuleFuncNames),
	)

	return module, err
}

func (d *DatabaseModel) InsertModule(module model.Modules) (bool, error) {
	sqlStatement := `INSERT INTO public."Modules"(
		"ModuleName", "ModuleFuncNames")
		VALUES ($1, $2);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		module.ModuleName,
		pq.Array(module.ModuleFuncNames))

	if err != nil {
		check = false
	}
	return check, err
}
