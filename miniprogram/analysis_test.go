package wechat_test

import (
	"context"
	"fmt"
	"testing"

	wechat "github.com/baiyeth/wechat/miniprogram"
)

var (
	an = wechat.NewAnalysis(context.Background())
)

func TestGetAnalysisDailyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisDailyRetain(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisDailySummary(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisDailySummary(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisDailyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisDailyVisitTrend(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisMonthlyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisMonthlyRetain(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisMonthlyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisMonthlyVisitTrend(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisUserPortrait(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisUserPortrait(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisVisitDistribution(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisVisitDistribution(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisVisitPage(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisVisitPage(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisWeeklyRetain(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisWeeklyRetain(beginDate, endDate)
	fmt.Print(data, err)
}

func TestGetAnalysisWeeklyVisitTrend(t *testing.T) {
	t.Parallel()
	beginDate := "2020-01-01"
	endDate := "2021-01-01"
	data, err := an.GetAnalysisWeeklyVisitTrend(beginDate, endDate)
	fmt.Print(data, err)
}
