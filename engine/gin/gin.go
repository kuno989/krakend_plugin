package gin

import (
	"context"
	"fmt"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
	"github.com/kuno989/cert_plugin"
	"github.com/kuno989/cert_plugin/engine/pkg"
	"net/http"
)

func Register(cfg *config.ServiceConfig, logger logging.Logger, engine *gin.Engine){
	certCfg := cert_plugin.ParseConfig(cfg.ExtraConfig, logger)
	if cfg == nil {
		return
	}
	ctx := context.Background()
	mongo, err := pkg.NewMongo(ctx, *certCfg)
	if err != nil {
		fmt.Errorf("mongo connection")
	}

	engine.Use(middleware(mongo, logger))
}

func middleware(mongo *pkg.Mongo, logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key, err := mongo.ApiKeyCheck(ctx, "123")
		if err != nil {
			logger.Error(fmt.Sprintf("%s 허용되지않은 라이센스 키"), key)
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		fmt.Println("허용됨",key)
		ctx.Next()
	}
}