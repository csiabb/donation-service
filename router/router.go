/*
Copyright ArxanChain Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package router

import (
	"fmt"
	"net/http"

	"github.com/csiabb/donation-service/common/log"
	srvctx "github.com/csiabb/donation-service/context"
	"github.com/csiabb/donation-service/controllers/org"
	"github.com/csiabb/donation-service/controllers/pub"
	"github.com/csiabb/donation-service/controllers/version"
	"github.com/csiabb/donation-service/middleware"

	"github.com/gin-gonic/gin"
)

// api version
const (
	APIVersion = "v1"
)

var (
	logger = log.MustGetLogger("router")

	// main url prefix
	apiPrefix = fmt.Sprintf("api/%s", APIVersion)

	// checkAPI
	checkVersionURL = "version"
)

const (
	urlPubFundsReceive    = "pub/funds/receive"
	urlPubFundsQuery      = "pub/funds/query"
	urlPubFundsDetail     = "pub/funds/detail"
	urlPubSuppliesReceive = "pub/supplies/receive"
	urlPubSuppliesQuery   = "pub/supplies/query"
	urlPubSuppliesDetail  = "pub/supplies/detail"
	urlPubList            = "pub/list"

	// org
	urlOrgsQuery = "orgs/query"
)

// Router service router
type Router struct {
	context        *srvctx.Context
	versionHandler *version.RestHandler
	pubHandler     *pub.RestHandler
	orgHandler     *org.RestHandler
}

// InitRouter init router
func (r *Router) InitRouter(ctx *srvctx.Context) error {
	if nil == ctx {
		return fmt.Errorf("param is nil")
	}

	r.context = ctx

	// Init version handler
	var err error
	r.versionHandler, err = version.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create version rest http handler instance, %+v", err)
		return err
	}

	r.pubHandler, err = pub.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create pub rest http handler instance, %+v", err)
		return err
	}

	r.orgHandler, err = org.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create organization rest http handler instance, %+v", err)
		return err
	}

	return nil
}

// SetupRouter add routes for rest api server
func (r *Router) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Delims("{{", "}}")
	router.Use(Cors())

	// service version
	router.GET(checkVersionURL, r.versionHandler.Version)

	// v1 group api
	apiPrefix := router.Group(apiPrefix)
	{
		// log reponse and request
		apiPrefix.Use(middleware.RequestResponseLogger())

		// publicity
		apiPrefix.POST(urlPubFundsReceive, r.pubHandler.ReceiveFunds)
		apiPrefix.GET(urlPubFundsQuery, r.pubHandler.QueryFunds)
		apiPrefix.GET(urlPubFundsDetail, r.pubHandler.QueryFundsDetail)
		apiPrefix.POST(urlPubSuppliesReceive, r.pubHandler.ReceiveSupplies)
		apiPrefix.GET(urlPubSuppliesQuery, r.pubHandler.QuerySupplies)
		apiPrefix.GET(urlPubSuppliesDetail, r.pubHandler.QuerySuppliesDetail)
		apiPrefix.GET(urlPubList, r.pubHandler.PubUserList)

		// org
		apiPrefix.GET(urlOrgsQuery, r.orgHandler.QueryOrganizations)
	}
	return router
}

// Cors ...
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()

	}
}
