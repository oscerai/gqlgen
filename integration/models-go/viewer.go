package models

import "github.com/oscerai/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
