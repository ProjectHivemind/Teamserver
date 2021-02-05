package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

func (d *DatabaseModel) GetCallBackById(id string) (model.CallBack, error) {
	var callBack model.CallBack

	sqlStatement := `SELECT * FROM public."CallBack" WHERE UUIDImplant=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&callBack.UUIDImplant,
		&callBack.FirstCall,
		&callBack.LastCall,
	)

	return callBack, err
}
