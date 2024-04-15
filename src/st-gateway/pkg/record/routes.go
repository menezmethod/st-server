package record

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/routes"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
)

func RegisterRecordRoutes(r *gin.Engine, config *configs.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		RecordClient: InitRecordServiceClient(config),
	}

	journalGroup := r.Group(fmt.Sprintf("v%v/", config.ApiVersion))

	protectedGroup := journalGroup.Group("/")
	protectedGroup.Use(a.AuthRequired)

	protectedEndpoints := []struct {
		method, path string
		handler      gin.HandlerFunc
	}{
		{"GET", "/records/", svc.ListRecords},
		{"POST", "/record/", svc.CreateRecord},
		{"PATCH", "/record/:id", svc.UpdateRecord},
		{"GET", "/record/:id", svc.GetRecord},
		{"DELETE", "/record/:id", svc.RemoveRecord},
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
	return nil
}

func (svc *ServiceClient) ListRecords(ctx *gin.Context) {
	routes.ListRecords(ctx, svc.RecordClient)
}
func (svc *ServiceClient) GetRecord(ctx *gin.Context) {
	routes.FineOneRecord(ctx, svc.RecordClient)
}
func (svc *ServiceClient) CreateRecord(ctx *gin.Context) {
	routes.CreateRecord(ctx, svc.RecordClient)
}
func (svc *ServiceClient) UpdateRecord(ctx *gin.Context) {
	routes.UpdateRecord(ctx, svc.RecordClient)
}
func (svc *ServiceClient) RemoveRecord(ctx *gin.Context) {
	routes.RemoveRecord(ctx, svc.RecordClient)
}
