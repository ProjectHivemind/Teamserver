package config

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

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
