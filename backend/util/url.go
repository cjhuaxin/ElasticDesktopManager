package util

import (
	"fmt"
	"strings"

	"github.com/cjhuaxin/ElasticDesktopManager/backend/resource"
	"github.com/goware/urlx"
)

func NormalizeUrls(urls string) ([]string, error) {
	urlSlice := strings.Split(urls, ",")
	endpoints := make([]string, 0)
	for _, addr := range urlSlice {
		parsed, err := urlx.Parse(addr)
		if err != nil {
			return nil, err
		}
		host, port, err := urlx.SplitHostPort(parsed)
		if err != nil {
			return nil, err
		}
		if port == "" {
			port = resource.EsDefaultPort
		}
		endpoints = append(endpoints, fmt.Sprintf("%s://%s:%s", parsed.Scheme, host, port))
	}

	return endpoints, nil
}