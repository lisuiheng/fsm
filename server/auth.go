package server

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lisuiheng/esm"
	"github.com/lisuiheng/fsm"
	"net/http"
	"strings"
	"time"
)

var enforcer *casbin.Enforcer

func (s *Service) Login(c *gin.Context) {
	if !s.UseAuth {
		c.Header(fsm.AUTH_KEY, fsm.NO_AUTH)
		c.AbortWithStatus(http.StatusOK)
		return
	}

	if enforcer == nil {
		enforcer = casbin.NewEnforcer("config/auth/url_model.conf", "config/auth/url_policy.csv")
	}

	var user fsm.User
	if err := c.BindJSON(&user); err != nil {
		s.invalidJSON(c)
	}
	ctx := c.Request.Context()
	srcs, err := s.client.UsersStore.All(ctx)
	if err != nil {
		s.unknownError(c, err)
		return
	}
	for _, src := range srcs {
		if src.Username == user.Username && src.Password == user.Password {

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, fsm.AuthClaims{
				ID:       src.ID,
				Username: src.Username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
				},
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte(s.Signature))
			if err == nil {
				c.Header(esm.AUTH_KEY, tokenString)
				c.Header(esm.ROLE_KEY, strings.Join(enforcer.GetRolesForUser(src.Username), " "))
				c.AbortWithStatus(http.StatusOK)
				return
			}
		}
	}
	s.forbiddenError(c)
}

func (s *Service) auth(c *gin.Context) {
	if !s.UseAuth {
		return
	}

	if enforcer == nil {
		enforcer = casbin.NewEnforcer("config/auth/url_model.conf", "config/auth/url_policy.csv")
	}

	auth := c.Request.Header.Get(esm.AUTH_KEY)

	claims := esm.AuthClaims{}
	token, err := jwt.ParseWithClaims(auth, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return []byte(s.Signature), nil
	})
	if err != nil {
		s.unauthorizedError(c, err)
		return
	}

	if !token.Valid {
		s.unauthorizedError(c, fmt.Errorf("token valid"))
		return
	}

	if !enforcer.Enforce(claims.Username, c.Request.URL.Path, c.Request.Method) {
		s.unauthorizedError(c, fmt.Errorf("role valid"))
	}
}
