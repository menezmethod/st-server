package record

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/pb"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/record/routes"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/util"
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
	return svc
}

func (svc *ServiceClient) recordHandler(ctx *gin.Context, routeFunc func(*gin.Context, pb.RecordServiceClient, *authPb.User)) {
	user, err := util.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	routeFunc(ctx, svc.RecordClient, user)
}

func (svc *ServiceClient) ListRecords(ctx *gin.Context) {
	svc.recordHandler(ctx, routes.ListRecords)
}
func (svc *ServiceClient) GetRecord(ctx *gin.Context) {
	svc.recordHandler(ctx, routes.FineOneRecord)
}
func (svc *ServiceClient) CreateRecord(ctx *gin.Context) {
	svc.recordHandler(ctx, routes.CreateRecord)
}
func (svc *ServiceClient) UpdateRecord(ctx *gin.Context) {
	svc.recordHandler(ctx, routes.UpdateRecord)
}
func (svc *ServiceClient) RemoveRecord(ctx *gin.Context) {
	svc.recordHandler(ctx, routes.RemoveRecord)
}
