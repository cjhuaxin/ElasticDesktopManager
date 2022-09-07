package base

import (
	"github.com/99designs/keyring"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
	"github.com/cjhuaxin/ElasticDesktopManager/backend/util"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

type BaseService struct {
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

func (*BaseService) BuildSucess(data interface{}) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: "0",
		ErrMsg:  "",
		Data:    data,
	}
}

func (*BaseService) BuildFailed(errCode, errMsg string) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    nil,
	}
}

func (*BaseService) InitEsClient(urls string, username, password string) (*elastic.Client, error) {
	options := make([]elastic.ClientOptionFunc, 0)
	urlSlice, err := util.NormalizeUrls(urls)
	if err != nil {
		return nil, err
	}
	options = append(options, elastic.SetURL(urlSlice...))
	if username != "" {
		options = append(options, elastic.SetBasicAuth(username, password))
	}
	options = append(options, elastic.SetSniff(false))

	return elastic.NewClient(options...)
}
