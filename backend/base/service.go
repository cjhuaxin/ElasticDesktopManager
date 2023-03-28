package base

import (
	"github.com/99designs/keyring"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"os"
)

type Service struct {
	Ctx     *models.EdmContext
	Paths   *Paths
	Log     *zap.SugaredLogger
	Keyring keyring.Keyring
}

type Paths struct {
	HomeDir string
	ConfDir string
	DbDir   string
	LogDir  string
	TmpDir  string
}

func (*Service) BuildSucess(data interface{}) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: "0",
		ErrMsg:  "",
		Data:    data,
	}
}

func (*Service) BuildFailed(errCode, errMsg string) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    nil,
	}
}

func (*Service) InitEsClient(urls string, username, password string) (*elasticsearch.Client, error) {
	urlSlice, err := util.NormalizeUrls(urls)
	if err != nil {
		return nil, err
	}
	return elasticsearch.NewClient(elasticsearch.Config{
		// 设置ES服务地址，支持多个地址
		Addresses: urlSlice,
		Username:  username,
		Password:  password,
		//		CACert:    cert,
		Logger: &elastictransport.CurlLogger{
			Output: os.Stdout,
		},
	},
	)
}
