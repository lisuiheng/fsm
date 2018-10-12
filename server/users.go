package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/chronograf"
	"github.com/lisuiheng/fsm"
	"net/http"
)

func (s *Service) User(c *gin.Context) {
	ctx := c.Request.Context()

	srcs, err := s.client.UsersStore.All(ctx)
	if err != nil {
		s.unknownError(c, err)
		return
	}
	c.Header("content-range", fmt.Sprintf("%d/%d", 0, len(srcs)))
	if len(srcs) == 0 {
		srcs = []fsm.User{}
	}
	c.JSON(http.StatusOK, srcs)
}

func (s *Service) NewUser(c *gin.Context) {
	var user fsm.User
	ctx := c.Request.Context()
	if c.BindJSON(&user) == nil {
		if err := s.client.UsersStore.Add(ctx, &user); err != nil {
			s.unknownError(c, err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}

func (s *Service) UserID(c *gin.Context) {
	id, hasID := c.Params.Get("id")
	if !hasID {
		s.errorw(c, http.StatusUnprocessableEntity, "id un translation")
	}

	ctx := c.Request.Context()
	src, err := s.client.UsersStore.Get(ctx, id)
	if err != nil {
		s.unknownError(c, err)
		return
	}
	c.JSON(http.StatusOK, src)
}

func (s *Service) UpdateUser(c *gin.Context) {
	id, hasID := c.Params.Get("id")
	if !hasID {
		s.errorw(c, http.StatusUnprocessableEntity, "id un translation")
	}

	ctx := c.Request.Context()
	src, err := s.client.UsersStore.Get(ctx, id)
	if err != nil {
		s.notFound(c, id)
		return
	}

	var req = fsm.User{}
	if c.BindJSON(&req) == nil {
		if req.Name != "" {
			src.Name = req.Name
		}

		if req.Username != "" {
			src.Username = req.Username
		}

		if req.Password != "" {
			src.Password = req.Password
		}

		if req.Role != 0 {
			src.Role = req.Role
		}

		if err := s.client.UsersStore.Update(ctx, src); err != nil {
			msg := fmt.Sprintf("Error updating source ID %d", id)
			s.errorw(c, http.StatusInternalServerError, msg)
			return
		}
		c.JSON(http.StatusOK, src)
		return
	}
	s.invalidJSON(c)
}

func (s *Service) RemoveUser(c *gin.Context) {
	id, hasID := c.Params.Get("id")
	if !hasID {
		s.errorw(c, http.StatusUnprocessableEntity, "id un translation")
	}

	src := fsm.User{ID: id}
	ctx := c.Request.Context()
	if err := s.client.UsersStore.Delete(ctx, &src); err != nil {
		if err == chronograf.ErrSourceNotFound {
			s.notFound(c, id)
		} else {
			s.unknownError(c, err)
		}
		return
	}
	c.AbortWithStatus(http.StatusOK)
}
