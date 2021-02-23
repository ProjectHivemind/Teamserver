package crud

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

	sqlStatement := `SELECT * FROM public."Implant" WHERE "UUID"=$1`

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

func (d *DatabaseModel) GetImplantByType(id string) ([]model.Implant, error) {
	var allImplants []model.Implant

	sqlStatement := `SELECT * FROM public."Implant" WHERE "UUIDImplantType"=$1`

	rows, err := d.db.Query(sqlStatement, id)
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

func (d *DatabaseModel) GetImplantByIp(id string) ([]model.Implant, error) {
	var allImplants []model.Implant

	sqlStatement := `SELECT * FROM public."Implant" WHERE "PrimaryIP"=$1`

	rows, err := d.db.Query(sqlStatement, id)
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

func (d *DatabaseModel) InsertImplant(implant model.Implant) (bool, error) {
	sqlStatement := `INSERT INTO public."Implant"(
		"UUID",
		"UUIDImplantType",
		"PrimaryIP",
		"Hostname",
		"MAC",
		"ImplantOS",
		"OtherIPs",
		"SupportedModules")
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		implant.UUID,
		implant.UUIDImplantType,
		implant.PrimaryIP,
		implant.Hostname,
		implant.MAC,
		implant.ImplantOS,
		pq.Array(implant.OtherIPs),
		pq.Array(implant.SupportedModules))

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) DeleteImplant(id string) (bool, error) {
	sqlStatement := `DELETE FROM public."Implant"
		WHERE "UUID"=$1`

	check := true

	_, err := d.db.Exec(sqlStatement, id)

	if err != nil {
		check = false
	}
	return check, err
}
