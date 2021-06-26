package helper

import (
	"database/sql"
	"github.com/MichalH95/exampleREST/model"
)

func NewClientWithCompanyId() model.Client {
	return model.Client{CompanyId: sql.NullInt64{
		Valid: true,
	}}
}

func NewClientWithPersonId() model.Client {
	return model.Client{PersonId: sql.NullInt64{
		Valid: true,
	}}
}
