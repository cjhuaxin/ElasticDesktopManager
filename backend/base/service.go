package base

import (
	"github.com/cjhuaxin/ElasticDesktopManager/backend/models"
)

func BuildSucess(data interface{}) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: "0",
		ErrMsg:  "",
		Data:    data,
	}
}

func BuildFailed(errCode, errMsg string) *models.BaseResponse {
	return &models.BaseResponse{
		ErrCode: errCode,
		ErrMsg:  errMsg,
		Data:    nil,
	}
}
