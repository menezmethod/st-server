package journal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/routes"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
)

func RegisterJournalRoutes(r *gin.Engine, config *configs.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		JournalClient: InitJournalServiceClient(config),
	}

	journalGroup := r.Group(fmt.Sprintf("v%v/", config.ApiVersion))

	protectedGroup := journalGroup.Group("/")
	protectedGroup.Use(a.AuthRequired)

	protectedEndpoints := []struct {
		method, path string
		handler      gin.HandlerFunc
	}{
		{"GET", "/journals", svc.ListJournals},
		{"POST", "/journal/", svc.CreateJournal},
		{"PATCH", "/journal/:id", svc.UpdateJournal},
		{"GET", "/journal/:id", svc.GetJournal},
		{"DELETE", "/journal/:id", svc.RemoveJournal},
	}

	for _, e := range protectedEndpoints {
		switch e.method {
		case "GET":
			protectedGroup.GET(e.path, e.handler)
		case "POST":
			protectedGroup.POST(e.path, e.handler)
		case "PATCH":
			protectedGroup.PATCH(e.path, e.handler)
		case "DELETE":
			protectedGroup.DELETE(e.path, e.handler)
		}
	}
	return svc
}

func (svc *ServiceClient) journalHandler(ctx *gin.Context, routeFunc func(*gin.Context, pb.JournalServiceClient, *authPb.User)) {
	user, err := util.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	routeFunc(ctx, svc.JournalClient, user)
}

func (svc *ServiceClient) ListJournals(ctx *gin.Context) {
	svc.journalHandler(ctx, routes.ListJournals)
}

func (svc *ServiceClient) GetJournal(ctx *gin.Context) {
	svc.journalHandler(ctx, routes.FineOneJournal)
}

func (svc *ServiceClient) CreateJournal(ctx *gin.Context) {
	svc.journalHandler(ctx, routes.CreateJournal)
}

func (svc *ServiceClient) UpdateJournal(ctx *gin.Context) {
	svc.journalHandler(ctx, routes.UpdateJournal)
}

func (svc *ServiceClient) RemoveJournal(ctx *gin.Context) {
	svc.journalHandler(ctx, routes.RemoveJournal)
}
