package journal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/journal/routes"
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

func (svc *ServiceClient) ListJournals(ctx *gin.Context) {
	routes.ListJournals(ctx, svc.JournalClient)
}
func (svc *ServiceClient) GetJournal(ctx *gin.Context) {
	routes.FineOneJournal(ctx, svc.JournalClient)
}
func (svc *ServiceClient) CreateJournal(ctx *gin.Context) {
	routes.CreateJournal(ctx, svc.JournalClient)
}
func (svc *ServiceClient) UpdateJournal(ctx *gin.Context) {
	routes.UpdateJournal(ctx, svc.JournalClient)
}
func (svc *ServiceClient) RemoveJournal(ctx *gin.Context) {
	routes.RemoveJournal(ctx, svc.JournalClient)
}
