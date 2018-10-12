package server

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUrlAuth(t *testing.T) {
	e := casbin.NewEnforcer("../config/auth/url_model.conf", "../config/auth/url_policy.csv")
	user := "admin"
	method := "GET"
	path := "/api/users"
	pass := e.Enforce(user, path, method)
	assert.True(t, pass)
}

func TestRole(t *testing.T) {
	e := casbin.NewEnforcer("../config/auth/url_model.conf", "../config/auth/test_url_policy.csv")
	//roles := e.GetAllRoles()
	roles := e.GetRolesForUser("admin")
	fmt.Println(strings.Join(roles, " "))
}
