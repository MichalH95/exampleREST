package helper

import (
	"database/sql"
	"encoding/json"
	"github.com/MichalH95/exampleREST/model"
)

func NewClientAsCompany() model.Client {
	return model.Client{
		CompanyId:  sql.NullInt64{Valid: true},
		ClientType: model.ClientTypeCompany,
	}
}

func NewClientAsPerson() model.Client {
	return model.Client{
		PersonId:   sql.NullInt64{Valid: true},
		ClientType: model.ClientTypePerson,
	}
}

func MessageJson(msg string) []byte {
	ret, _ := json.Marshal(map[string]string{
		"Message": msg,
	})
	return ret
}

func ErrorMessageJson(msg string) []byte {
	ret, _ := json.Marshal(map[string]string{
		"Error": msg,
	})
	return ret
}
