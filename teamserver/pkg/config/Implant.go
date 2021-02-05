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
