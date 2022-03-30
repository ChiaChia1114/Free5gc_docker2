/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package management

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/nrf/internal/logger"
	"github.com/free5gc/nrf/internal/sbi/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

// DeregisterNFInstance - Deregisters a given NF Instance
func HTTPDeregisterNFInstance(c *gin.Context) {
	// parse nfInstanceId

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["nfInstanceID"] = c.Params.ByName("nfInstanceID")

	httpResponse := producer.HandleNFDeregisterRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}

// GetNFInstance - Read the profile of a given NF Instance
func HTTPGetNFInstance(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["nfInstanceID"] = c.Params.ByName("nfInstanceID")

	httpResponse := producer.HandleGetNFInstanceRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}

// RegisterNFInstance - Register a new NF Instance
func HTTPRegisterNFInstance(c *gin.Context) {
	var nfprofile models.NfProfile

	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.ManagementLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	// step 2: convert requestBody to openapi models
	err = openapi.Deserialize(&nfprofile, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.ManagementLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	// step 3: encapsulate the request by http_wrapper package
	req := httpwrapper.NewRequest(c.Request, nfprofile)

	// step 4: call producer
	httpResponse := producer.HandleNFRegisterRequest(req)

	for key, val := range httpResponse.Header {
		c.Header(key, val[0])
	}

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}

// UpdateNFInstance - Update NF Instance profile
func HTTPUpdateNFInstance(c *gin.Context) {
	// step 1: retrieve http request body
	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.ManagementLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["nfInstanceID"] = c.Params.ByName("nfInstanceID")
	req.Body = requestBody

	httpResponse := producer.HandleUpdateNFInstanceRequest(req)

	responseBody, err := openapi.Serialize(httpResponse.Body, "application/json")
	if err != nil {
		logger.ManagementLog.Warnln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(httpResponse.Status, "application/json", responseBody)
	}
}
