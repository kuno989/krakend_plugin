package gin

import (
	"context"
	"fmt"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
	"github.com/kuno989/cert_plugin"
	"github.com/kuno989/cert_plugin/engine/pkg"
	"github.com/kuno989/cert_plugin/engine/schema"
	"net/http"
)

const HeaderAuthorization = "x-api-key"

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
		userToken := ctx.GetHeader(HeaderAuthorization)
		if userToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, schema.ResponseMessage{
				Code: http.StatusUnauthorized,
				Error:"Forbidden",
				Message: "Forbidden",
			})
			return
		}

		key, err := mongo.ApiKeyCheck(ctx, userToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, schema.ResponseMessage{
				Code: http.StatusUnauthorized,
				Error:"Forbidden",
				Message: "Forbidden",
			})
			return
		}
		fmt.Printf("허용된 api key %s\n",key.ApiKey)
		ctx.Next()
	}
}