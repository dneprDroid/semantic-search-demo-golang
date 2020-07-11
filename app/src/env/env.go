package env 

import (
	"appserver/database"

	"common/rpc"
)

type Env struct {
	Db *database.Pool
	Rpc *rpc.Pool
}
