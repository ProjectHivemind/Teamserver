package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) GetModuleFuncById(id string) (model.ModulesFunc, error) {
	var moduleFunc model.ModulesFunc

	sqlStatement := `SELECT * FROM public."ModuleFuncs" WHERE "UUID"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&moduleFunc.UUID,
		&moduleFunc.ModuleFuncName,
		&moduleFunc.ModuleFuncDesc,
		&moduleFunc.NumOfParams,
		pq.Array(&moduleFunc.ParamTypes),
		pq.Array(&moduleFunc.ParamNames),
	)

	return moduleFunc, err
}

func (d *DatabaseModel) InsertModuleFunc(moduleFunc model.ModulesFunc) (bool, error) {
	sqlStatement := `INSERT INTO public."ModuleFuncs"(
		"UUID", "ModuleFuncName", "ModuleFuncDesc", "NumOfParameters", "ParameterTypes", "ParameterNames")
		VALUES ($1, $2, $3, $4, $5, $6);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		moduleFunc.UUID,
		moduleFunc.ModuleFuncName,
		moduleFunc.ModuleFuncDesc,
		moduleFunc.NumOfParams,
		pq.Array(moduleFunc.ParamTypes),
		pq.Array(moduleFunc.ParamNames))

	if err != nil {
		check = false
	}
	return check, err
}
