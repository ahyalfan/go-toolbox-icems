package restapi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/ahyalfan/go-toolbox-icems/gateway/model"
	"github.com/ahyalfan/go-toolbox-icems/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContentType string

var (
	ContentTypeJson     ContentType = "application/json"
	ContentTypeForm     ContentType = "application/x-www-form-urlencoded"
	ContentTypeMulipart ContentType = "multipart/form-data"
)

type Rest[T model.Response] struct {
	ApiKey string
	URL    string
	Log    *logrus.Logger
}

func (r *Rest[T]) GetData(header map[string]string) (T, error) {
	var something T
	agent := fiber.Get(r.URL)

	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "get data failed",
		}).Error(errs)
		return something, errs[0]
	}
	if statusCode != fiber.StatusOK {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
		return something, fmt.Errorf("failed get data with statusCode %d", statusCode)
	}
	err := json.Unmarshal(body, &something)
	if err != nil {
		return something, err
	}
	return something, nil
}

func (r *Rest[T]) SendDataJson(header map[string]string, req model.Request) (T, error) {
	agent := fiber.Post(r.URL)
	agent = agent.Set("Content-Type", string(ContentTypeJson))
	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	// set header
	for k, v := range header {
		agent = agent.Set(k, v)
	}

	reqJson, _ := json.Marshal(req)
	agent = agent.Body(reqJson)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendDataForm(header map[string]string, req model.Request) (T, error) {
	agent := fiber.Post(r.URL)
	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	// set header
	for k, v := range header {
		agent = agent.Set(k, v)
	}

	args := fiber.AcquireArgs()
	defer fiber.ReleaseArgs(args)
	results := utils.StructToMapString(req, "form")
	for k, v := range results {
		args.Set(k, v)
	}
	agent = agent.Form(args)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendDataMultiForm(header map[string]string, req model.RequestMultiForm) (T, error) {
	agent := fiber.Post(r.URL)
	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	// set header
	for k, v := range header {
		agent = agent.Set(k, v)
	}
	multiPartFile := req.GetFile()

	file, err := multiPartFile.Open()
	if err != nil {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
		}).Error("open file failed")
		var zero T
		return zero, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
		}).Error("read file failed")
		var zero T
		return zero, err
	}
	results := utils.StructToMapString(req, "form")
	ff1 := &fiber.FormFile{
		Fieldname: "file",
		Name:      req.GetFileName(),
		Content:   fileBytes,
	}

	args := fiber.AcquireArgs()
	defer fiber.ReleaseArgs(args)

	for k, v := range results {
		args.Set(k, v)
	}
	agent = agent.
		FileData(ff1).
		MultipartForm(args)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendAgent(agent *fiber.Agent) (T, error) {
	var response T
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "send data failed",
		}).Error(errs)
		return response, errs[0]
	}
	if statusCode > 300 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
		return response, fmt.Errorf("failed get data with statusCode %d", statusCode)
	}

	err := json.Unmarshal(body, &response)
	if err != nil {
		r.Log.WithError(err).Error("failed to unmarshal response")
		return response, err
	}
	return response, nil
}
