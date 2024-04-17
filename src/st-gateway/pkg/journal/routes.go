package journal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/menezmethod/st-server/src/st-gateway/configs"
	"github.com/menezmethod/st-server/src/st-gateway/pkg/auth"
	authPb "github.com/menezmethod/st-server/src/st-gateway/pkg/auth/pb"
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

func GetUserFromContext(ctx *gin.Context) (*authPb.User, error) {
	userID, exists := ctx.Get("userID")
	if !exists {
		return nil, fmt.Errorf("user ID not found in context")
	}

	userIDUint, ok := userID.(uint64)
	if !ok {
		return nil, fmt.Errorf("invalid user ID type: expected uint64, got %T", userID)
	}

	return &authPb.User{
		Id: userIDUint,
	}, nil
}

func (svc *ServiceClient) ListJournals(ctx *gin.Context) {
	routes.ListJournals(ctx, svc.JournalClient)
}

func (svc *ServiceClient) GetJournal(ctx *gin.Context) {
	routes.FineOneJournal(ctx, svc.JournalClient)
}

func (svc *ServiceClient) CreateJournal(ctx *gin.Context) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	routes.CreateJournal(ctx, svc.JournalClient, user)
}

func (svc *ServiceClient) UpdateJournal(ctx *gin.Context) {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	routes.UpdateJournal(ctx, svc.JournalClient, user)
}

func (svc *ServiceClient) RemoveJournal(ctx *gin.Context) {
	routes.RemoveJournal(ctx, svc.JournalClient)
}
