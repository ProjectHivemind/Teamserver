package config

import (
	"time"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

func (d *DatabaseModel) GetCallBackById(id string) (model.CallBack, error) {
	var callBack model.CallBack

	sqlStatement := `SELECT * FROM public."CallBack" WHERE "UUIDImplant"=$1`

	var time1, time2 string
	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&callBack.UUIDImplant,
		&time1,
		&time2,
	)

	return callBack, err
}

func (d *DatabaseModel) InsertCallBack(callBack model.CallBack) (bool, error) {
	sqlStatement := `INSERT INTO public."CallBack"(
		"UUIDImplant", "FirstCall", "LastCall")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		callBack.UUIDImplant,
		callBack.FirstCall.String(),
		callBack.LastCall.String())

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) UpdateCallBackTime(id string, t time.Time) (bool, error) {
	sqlStatement := `UPDATE public."CallBack"
		SET "LastCall"=$2
		WHERE UUIDImplant=$1;`

	check := true

	_, err := d.db.Exec(sqlStatement, id, t)

	if err != nil {
		check = false
	}
	return check, err
}
