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
	"github.com/csiabb/donation-service/controllers/acc"
	"github.com/csiabb/donation-service/controllers/bc"
	"github.com/csiabb/donation-service/controllers/image"
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
	// acc
	urlAccLoginWXApp = "acc/login/wxapp"

	// block chain
	urlBCCallBack = "bc/cb"

	// pub
	urlPubFunds          = "pub/funds"
	urlPubFundsDetail    = "pub/funds/detail"
	urlPubSupplies       = "pub/supplies"
	urlPubSuppliesDetail = "pub/supplies/detail"
	urlPubList           = "pub/list"

	// org
	urlOrgCharities       = "org/charities"
	urlOrgCharitiesDetail = "org/charities/detail"

	// image
	urlImageUpload = "image/upload"
	urlImageShare  = "image/share"
)

// Router service router
type Router struct {
	context        *srvctx.Context
	versionHandler *version.RestHandler
	pubHandler     *pub.RestHandler
	orgHandler     *org.RestHandler
	accHandler     *acc.RestHandler
	imageHandler   *image.RestHandler
	bcHandler      *bc.RestHandler
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

	r.accHandler, err = acc.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create account rest http handler instance, %+v", err)
		return err
	}

	r.imageHandler, err = image.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create image rest http handler instance, %+v", err)
		return err
	}

	r.bcHandler, err = bc.NewRestHandler(r.context)
	if err != nil {
		logger.Errorf("Failed to create block chain rest http handler instance, %+v", err)
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

		// account
		apiPrefix.POST(urlAccLoginWXApp, r.accHandler.LoginWXApp)

		// block chain
		apiPrefix.POST(urlBCCallBack, r.bcHandler.BlockChainCallBack)

		// publicity
		apiPrefix.POST(urlPubFunds, r.pubHandler.ReceiveFunds)
		apiPrefix.GET(urlPubFunds, r.pubHandler.QueryFunds)
		apiPrefix.GET(urlPubFundsDetail, r.pubHandler.QueryFundsDetail)
		apiPrefix.POST(urlPubSupplies, r.pubHandler.ReceiveSupplies)
		apiPrefix.GET(urlPubSupplies, r.pubHandler.QuerySupplies)
		apiPrefix.GET(urlPubSuppliesDetail, r.pubHandler.QuerySuppliesDetail)
		apiPrefix.GET(urlPubList, r.pubHandler.PubUserList)

		// org
		apiPrefix.GET(urlOrgCharities, r.orgHandler.QueryOrgCharities)
		apiPrefix.GET(urlOrgCharitiesDetail, r.orgHandler.QueryOrgCharitiesDetail)

		// image
		apiPrefix.POST(urlImageUpload, r.imageHandler.Upload)
		apiPrefix.GET(urlImageShare, r.imageHandler.Share)
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
