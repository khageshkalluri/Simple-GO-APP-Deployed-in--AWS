package app

import(
	"dynamic/api"
	"dynamic/database"
)

type Appclient struct{
	Api api.ApiClient
}

func NewApp() Appclient{
	db:=database.NewDynamodb()
	apii:=api.NewApiClient(db)
    return Appclient{
		Api: apii,
	}
}