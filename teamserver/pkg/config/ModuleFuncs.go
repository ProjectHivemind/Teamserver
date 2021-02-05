package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

func (d *DatabaseModel) GetModuleFuncById(id string) (model.ModulesFuncs, error) {
	var moduleFunc model.ModulesFuncs

	sqlStatement := `SELECT * FROM public."ModuleFuncs" WHERE ModuleFuncName=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&moduleFunc.ModuleFuncName,
		&moduleFunc.NumOfParameters,
		&moduleFunc.ParameterTypes,
		&moduleFunc.ParameterNames,
	)

	return moduleFunc, err
}
