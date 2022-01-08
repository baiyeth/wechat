package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram/analysis"
)

func GetAnalysisDailyRetain(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisRetain, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisDailyRetain(beginDate, endDate)
}

func GetAnalysisDailySummary(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisDailySummary, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisDailySummary(beginDate, endDate)
}

func GetAnalysisDailyVisitTrend(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisVisitTrend, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisDailyVisitTrend(beginDate, endDate)
}

func GetAnalysisMonthlyRetain(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisRetain, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisMonthlyRetain(beginDate, endDate)
}

func GetAnalysisMonthlyVisitTrend(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisVisitTrend, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisMonthlyVisitTrend(beginDate, endDate)
}

func GetAnalysisUserPortrait(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisUserPortrait, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisUserPortrait(beginDate, endDate)
}

func GetAnalysisVisitDistribution(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisVisitDistribution, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisVisitDistribution(beginDate, endDate)
}

func GetAnalysisVisitPage(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisVisitPage, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisVisitPage(beginDate, endDate)
}

func GetAnalysisWeeklyRetain(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisRetain, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisWeeklyRetain(beginDate, endDate)
}

func GetAnalysisWeeklyVisitTrend(beginDate, endDate string, appid, appSecret string) (result analysis.ResAnalysisVisitTrend, err error) {
	miniprogram := GetMiniProgram(appid, appSecret)
	return miniprogram.GetAnalysis().GetAnalysisWeeklyVisitTrend(beginDate, endDate)
}
