package model

import (
	"fmt"
	"io"
	"strconv"
)

type Role string

const (
	RoleGeneral    Role = "GENERAL"
	RoleManager    Role = "MANAGER"
	RoleAdmin      Role = "ADMIN"
	RoleSuperAdmin Role = "SUPER_ADMIN"
)

var AllRole = []Role{
	RoleGeneral,
	RoleManager,
	RoleAdmin,
	RoleSuperAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleGeneral, RoleManager, RoleAdmin, RoleSuperAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
