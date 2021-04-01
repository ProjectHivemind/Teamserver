package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllModules() ([]model.Module, error) {
	var allModules []model.Module

	sqlStatement := `SELECT * FROM public."Modules"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var module model.Module

		err = rows.Scan(
			&module.ModuleName,
			&module.ModuleDesc,
			pq.Array(&module.ModuleFuncIds),
		)
		if err != nil {
			return nil, err
		}

		allModules = append(allModules, module)
	}

	return allModules, nil
}

func (d *DatabaseModel) GetModuleByName(name string) (model.Module, error) {
	var module model.Module

	sqlStatement := `SELECT * FROM public."Modules" WHERE "ModuleName"=$1`

	row := d.db.QueryRow(sqlStatement, name)
	err := row.Scan(
		&module.ModuleName,
		&module.ModuleDesc,
		pq.Array(&module.ModuleFuncIds),
	)

	return module, err
}

func (d *DatabaseModel) InsertModule(module model.Module) (bool, error) {
	sqlStatement := `INSERT INTO public."Modules"(
		"ModuleName", "ModuleDesc", "ModuleFuncIds")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		module.ModuleName,
		module.ModuleDesc,
		pq.Array(module.ModuleFuncIds))

	if err != nil {
		check = false
	}
	return check, err
}
