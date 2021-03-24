package crud

import "github.com/ProjectHivemind/Teamserver/teamserver/pkg/model"

func (d *DatabaseModel) GetSessionById(id string) (model.Sessions, error) {
	var session model.Sessions

	sqlStatement := `SELECT * FROM public."Sessions" WHERE "SessionToken"=$1`

	row := d.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&session.SessionToken,
		&session.Username,
		&session.ExpTime,
	)

	return session, err
}

func (d *DatabaseModel) InsertSession(sessions model.Sessions) (bool, error) {
	sqlStatement := `INSERT INTO public."Sessions"(
		"SessionToken", "Username", "ExpTime")
		VALUES ($1, $2, $3);`

	check := true

	_, err := d.db.Exec(sqlStatement,
		sessions.SessionToken,
		sessions.Username,
		sessions.ExpTime)

	if err != nil {
		check = false
	}
	return check, err
}

func (d *DatabaseModel) RemoveSessionById(id string) error {
	sqlStatement := `DELETE FROM public."Sessions" WHERE "SessionToken"=$1`

	_, err := d.db.Exec(sqlStatement, id)

	return err
}
