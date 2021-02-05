package config

import (
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
