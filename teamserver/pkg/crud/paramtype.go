package crud

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) GetParamType(name string) (model.ParamType, error) {
	var paramType model.ParamType

	sqlStatement := `SELECT * FROM public."ParamTypes" WHERE "TypeName"=$1`

	row := d.db.QueryRow(sqlStatement, name)
	err := row.Scan(
		&paramType.TypeName,
		&paramType.IsCombo,
		pq.Array(&paramType.ComboOptions),
	)

	return paramType, err
}

func (d *DatabaseModel) InsertParamType(paramType model.ParamType) (bool, error) {
	sqlStatement := `INSERT INTO public."ParamTypes"(
		"TypeName", "IsComboOption", "ComboOptions")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		paramType.TypeName,
		paramType.IsCombo,
		pq.Array(paramType.ComboOptions))

	if err != nil {
		check = false
	}
	return check, err
}
