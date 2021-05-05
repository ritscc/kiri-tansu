package model

import (
	"time"
)

type Ldap_user struct {
	User_id int
	Ldap_id int
	Created_at *time.Time
	Created_by int
	Updated_at *time.Time
	Updated_by int
}
