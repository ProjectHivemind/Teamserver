package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
)

func (d *DatabaseModel) AllImplantTypes() ([]model.ImplantType, error) {
	var allImplantTypes []model.ImplantType

	sqlStatement := `SELECT * FROM public."ImplantType"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var implantType model.ImplantType

		err = rows.Scan(
			&implantType.UUID,
			&implantType.ImplantName,
			&implantType.ImplantVersion,
		)
		if err != nil {
			return nil, err
		}

		allImplantTypes = append(allImplantTypes, implantType)
	}

	return allImplantTypes, nil
}

func (d *DatabaseModel) GetImplantTypeById(id string) (model.ImplantType, error) {
	var implantType model.ImplantType

	sqlStatement := `SELECT * FROM public."ImplantType" WHERE "UUID"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&implantType.UUID,
		&implantType.ImplantName,
		&implantType.ImplantVersion,
	)

	return implantType, err
}

func (d *DatabaseModel) InsertImplantType(implantType model.ImplantType) (bool, error) {
	sqlStatement := `INSERT INTO public."ImplantType"(
		"UUID", "ImplantName", "ImplantVersion")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		implantType.UUID,
		implantType.ImplantName,
		implantType.ImplantVersion)

	if err != nil {
		check = false
	}
	return check, err
}
