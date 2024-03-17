package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"st-gateway/configs"
	"st-gateway/pkg/auth/routes"
)

type endpoint struct {
	method, path string
	handler      func(ctx *gin.Context)
}

func RegisterAuthRoutes(r *gin.Engine, c *configs.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	endpoints := []endpoint{
		{"POST", "/auth/register", svc.Register},
		{"POST", "/auth/login", svc.Login},
		{"GET", "/auth/me", svc.FindMe},
		{"GET", "/users", svc.FindAllUsers},
		{"GET", "/users/:id", svc.FindOneUser},
		{"PATCH", "/users/:id", svc.UpdateUser},
		{"DELETE", "/user/:id", svc.DeleteUser},
	}

	group := r.Group(fmt.Sprintf("v%v/", c.ApiVersion))
	for _, e := range endpoints {
		switch e.method {
		case "POST":
			group.POST(e.path, e.handler)
		case "GET":
			group.GET(e.path, e.handler)
		case "PATCH":
			group.PATCH(e.path, e.handler)
		case "DELETE":
			group.DELETE(e.path, e.handler)
		}
	}

	return svc
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) { routes.DeleteUser(ctx, svc.Client) }
func (svc *ServiceClient) Login(ctx *gin.Context)      { routes.Login(ctx, svc.Client) }
func (svc *ServiceClient) Register(ctx *gin.Context)   { routes.Register(ctx, svc.Client) }
func (svc *ServiceClient) FindAllUsers(ctx *gin.Context) {
	routes.FindAllUsers(ctx, svc.Client)
}
func (svc *ServiceClient) FindOneUser(ctx *gin.Context) { routes.FindOneUser(ctx, svc.Client) }
func (svc *ServiceClient) FindMe(ctx *gin.Context)      { routes.Me(ctx, svc.Client) }
func (svc *ServiceClient) UpdateUser(ctx *gin.Context)  { routes.UpdateUser(ctx, svc.Client) }
