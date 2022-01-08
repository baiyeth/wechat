package wechat_test

import (
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

func TestGetAnalysisDailyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisDailyRetain(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisDailySummary(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisDailySummary(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisDailyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisDailyVisitTrend(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisMonthlyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisMonthlyRetain(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisMonthlyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisMonthlyVisitTrend(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisUserPortrait(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisUserPortrait(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisVisitDistribution(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisVisitDistribution(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisVisitPage(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisVisitPage(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisWeeklyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisWeeklyRetain(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}

func TestGetAnalysisWeeklyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	appid := ""
	appSecret := ""
	data, err := wechat.GetAnalysisWeeklyVisitTrend(beginDate, endDate, appid, appSecret)
	fmt.Print(data, err)
}
