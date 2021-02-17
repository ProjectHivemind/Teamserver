package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) GetModuleFuncById(id string) (model.ModulesFuncs, error) {
	var moduleFunc model.ModulesFuncs

	sqlStatement := `SELECT * FROM public."ModuleFuncs" WHERE "UUID"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&moduleFunc.UUID,
		&moduleFunc.ModuleFuncName,
		&moduleFunc.NumOfParameters,
		pq.Array(&moduleFunc.ParameterTypes),
		pq.Array(&moduleFunc.ParameterNames),
	)

	return moduleFunc, err
}

func (d *DatabaseModel) InsertModuleFunc(moduleFunc model.ModulesFuncs) (bool, error) {
	sqlStatement := `INSERT INTO public."ModuleFuncs"(
		"UUID", "ModuleFuncName", "NumOfParameters", "ParameterTypes", "ParameterNames")
		VALUES ($1, $2, $3, $4, $5);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		moduleFunc.UUID,
		moduleFunc.ModuleFuncName,
		moduleFunc.NumOfParameters,
		pq.Array(moduleFunc.ParameterTypes),
		pq.Array(moduleFunc.ParameterNames))

	if err != nil {
		check = false
	}
	return check, err
}
