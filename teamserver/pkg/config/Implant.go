package config

import (
	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllImplants() ([]model.Implant, error) {
	var allImplants []model.Implant

	sqlStatement := `SELECT * FROM public."Implant"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var implant model.Implant

		err = rows.Scan(
			&implant.UUID,
			&implant.UUIDImplantType,
			&implant.PrimaryIP,
			&implant.Hostname,
			&implant.MAC,
			&implant.ImplantOS,
			pq.Array(&implant.OtherIPs),
			pq.Array(&implant.SupportedModules),
		)
		if err != nil {
			return nil, err
		}

		allImplants = append(allImplants, implant)
	}

	return allImplants, nil
}

func (d *DatabaseModel) GetImplantById(id string) (model.Implant, error) {
	var implant model.Implant

	sqlStatement := `SELECT * FROM public."Implant" WHERE UUID=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&implant.UUID,
		&implant.UUIDImplantType,
		&implant.PrimaryIP,
		&implant.Hostname,
		&implant.MAC,
		&implant.ImplantOS,
		pq.Array(&implant.OtherIPs),
		pq.Array(&implant.SupportedModules),
	)

	return implant, err
}

func (d *DatabaseModel) GetImplantByType(id string) (model.Implant, error) {
	var implant model.Implant

	sqlStatement := `SELECT * FROM public."Implant" WHERE UUIDImplantType=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&implant.UUID,
		&implant.UUIDImplantType,
		&implant.PrimaryIP,
		&implant.Hostname,
		&implant.MAC,
		&implant.ImplantOS,
		pq.Array(&implant.OtherIPs),
		pq.Array(&implant.SupportedModules),
	)

	return implant, err
}
