package activityLog

import (
	"encoding/json"
	"latihan/infrastructure/configuration"
	"net/http"
	"time"
)

const (
	EndpointInsertLogActivity = "/api/v1/log"
	ServiceConsumer           = "MyGram"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 18 * time.Second,
	}
}

type LoggerService struct {
	Payload interface{} `json:"payload"`
	Url     string      `json:"url"`
	Method  string      `json:"method"`
	Path    string      `json:"path"`
}

type Service struct {
	conf *configuration.AppConfig
}

func NewService(conf *configuration.AppConfig) *Service {
	return &Service{
		conf: conf,
	}
}

// func (svc *Service) Call(request *http.Request, data interface{}) []byte {
// 	url := svc.conf.Hades.URL + EndpointInsertLogActivity
// 	var body = LoggerService{
// 		Payload: data,
// 		Url:     request.Host,
// 		Method:  request.Method,
// 		Path:    request.URL.Path,
// 	}
// 	jsonRequest, _ := json.Marshal(body)
// 	payload := bytes.NewBuffer(jsonRequest)

// 	newRequest, err := http.NewRequest(http.MethodPost, url, payload)
// 	if err != nil {
// 		return Error(err)
// 	}

// 	newRequest.Header.Add(utils.MyGramHeaderUserID, request.Header.Get(utils.MyGramHeaderUserID))
// 	newRequest.Header.Add(utils.MyGramHeaderIPAddress, request.Header.Get(utils.MyGramHeaderIPAddress))
// 	newRequest.Header.Add(utils.MyGramHeaderUserAgent, request.Header.Get(utils.MyGramHeaderUserAgent))
// 	newRequest.Header.Add(utils.MyGramHeaderTimeStamp, times.TimeStampString())

// 	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
// 		for key, val := range via[0].Header {
// 			req.Header[key] = val
// 		}
// 		return err
// 	}

// 	resp, err := utils.Request(newRequest)
// 	if err != nil {
// 		return Error(err)
// 	}
// 	defer resp.Body.Close()

// 	respByte, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Error(err)
// 	}

// 	var jsonb map[string]interface{}
// 	err = json.Unmarshal(respByte, &jsonb)
// 	if err != nil {
// 		return Error(err)
// 	}

// 	return respByte // optional needed

// }

func Error(err error) []byte {
	var x = make(map[string]interface{})
	x = map[string]interface{}{"rc": "01", "msg": err.Error()}

	data, _ := json.Marshal(x)
	return data
}
