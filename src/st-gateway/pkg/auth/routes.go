package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth/routes"
)

type endpoint struct {
	method, path string
	handler      func(ctx *gin.Context)
}

func RegisterAuthRoutes(r *gin.Engine, c *configs.Config) *ServiceClient {
	svc := &ServiceClient{
		AuthServiceClient: InitServiceClient(c),
	}

	a := InitAuthMiddleware(svc)

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

	authRoutes := group.Group("/")
	authRoutes.Use(a.AuthRequired)
	for _, e := range endpoints {
		switch e.method {
		case "POST":
			if e.path == "/auth/register" || e.path == "/auth/login" {
				group.POST(e.path, e.handler)
			} else {
				authRoutes.POST(e.path, e.handler)
			}
		case "GET":
			authRoutes.GET(e.path, e.handler)
		case "PATCH":
			authRoutes.PATCH(e.path, e.handler)
		case "DELETE":
			authRoutes.DELETE(e.path, e.handler)
		}
	}

	return svc
}

func (svc *ServiceClient) DeleteUser(ctx *gin.Context) { routes.DeleteUser(ctx, svc.AuthServiceClient) }
func (svc *ServiceClient) Login(ctx *gin.Context)      { routes.Login(ctx, svc.AuthServiceClient) }
func (svc *ServiceClient) Register(ctx *gin.Context)   { routes.Register(ctx, svc.AuthServiceClient) }
func (svc *ServiceClient) FindAllUsers(ctx *gin.Context) {
	routes.FindAllUsers(ctx, svc.AuthServiceClient)
}
func (svc *ServiceClient) FindOneUser(ctx *gin.Context) {
	routes.FindOneUser(ctx, svc.AuthServiceClient)
}
func (svc *ServiceClient) FindMe(ctx *gin.Context)     { routes.Me(ctx, svc.AuthServiceClient) }
func (svc *ServiceClient) UpdateUser(ctx *gin.Context) { routes.UpdateUser(ctx, svc.AuthServiceClient) }
