package crud

import (
	"errors"

	"github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"
	"github.com/lib/pq"
)

func (d *DatabaseModel) AllGroups() ([]model.Groups, error) {
	var allGroups []model.Groups

	sqlStatement := `SELECT * FROM public."Groups"`

	rows, err := d.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var group model.Groups

		err = rows.Scan(
			&group.UUID,
			&group.GroupName,
			pq.Array(&group.Implants),
		)
		if err != nil {
			return nil, err
		}

		allGroups = append(allGroups, group)
	}

	return allGroups, nil
}

func (d *DatabaseModel) GetGroupById(id string) (model.Groups, error) {
	var group model.Groups

	sqlStatement := `SELECT * FROM public."Groups" WHERE "UUID"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&group.UUID,
		&group.GroupName,
		pq.Array(&group.Implants),
	)

	return group, err
}

func (d *DatabaseModel) GetGroupByName(name string) (model.Groups, error) {
	var group model.Groups

	sqlStatement := `SELECT * FROM public."Groups" WHERE "GroupName"=$1`

	row := d.db.QueryRow(sqlStatement, name)
	err := row.Scan(
		&group.UUID,
		&group.GroupName,
		pq.Array(&group.Implants),
	)

	return group, err
}

func (d *DatabaseModel) InsertGroup(group model.Groups) (bool, error) {
	sqlStatement := `INSERT INTO public."Groups"(
		"UUID", "GroupName")
		VALUES ($1, $2);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		group.UUID,
		group.GroupName)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) AddUUIDToGroup(id string, implantId string) (bool, error) {
	group, err := d.GetGroupById(id)
	if err != nil {
		return false, err
	}

	_, err = d.GetImplantById(implantId)
	if err != nil {
		return false, err
	}

	group.Implants = append(group.Implants, implantId)

	sqlStatement := `UPDATE public."Groups"
		SET "Implants"=$2
		WHERE "UUID"=$1;`

	check := true

	_, err = d.db.Exec(sqlStatement, id, pq.Array(group.Implants))

	if err != nil {
		check = false
	}
	return check, err

}

func (d *DatabaseModel) RemoveUUIDFromGroup(id string, implantId string) (bool, error) {
	group, err := d.GetGroupById(id)
	if err != nil {
		return false, err
	}

	idx := -1
	for i := 0; i < len(group.Implants); i++ {
		if group.Implants[i] == implantId {
			idx = i
		}
	}
	if idx == -1 {
		return false, errors.New("ImplantId not in database")
	}

	if len(group.Implants) > 1 {
		group.Implants[len(group.Implants)-1], group.Implants[idx] =
			group.Implants[idx], group.Implants[len(group.Implants)-1]

		group.Implants = group.Implants[:len(group.Implants)-1]
	} else {
		group.Implants = []string{}
	}

	sqlStatement := `UPDATE public."Groups"
		SET "Implants"=$2
		WHERE "UUID"=$1;`

	check := true

	_, err = d.db.Exec(sqlStatement, id, pq.Array(group.Implants))

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) RemoveGroupById(id string) error {
	sqlStatement := `DELETE FROM public."Groups" WHERE "UUID"=$1`

	_, err := d.db.Exec(sqlStatement, id)

	return err
}
