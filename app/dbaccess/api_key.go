package dbaccess

import (
	"os"
)

func GetUser(api_key string) ([]User, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	query_string, err := os.ReadFile(wd + "/dbaccess/query/getApikey.go.sql")
	if err != nil {
		return nil, err
	}

	db, err := connect()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	
	rows, err := db.NamedQuery(string(query_string), map[string]interface{}{"api_key": api_key})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}