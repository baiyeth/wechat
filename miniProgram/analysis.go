package miniProgram

import (
	"github.com/silenceper/wechat/v2/miniprogram/analysis"
)

// Analysis 微信小程序相关API
type Analysis struct {
	*MiniProgram
}

func NewAnalysis(mp *MiniProgram) *Analysis {
	return &Analysis{mp}
}

func (a *Analysis) GetAnalysisDailyRetain(beginDate, endDate string) (result analysis.ResAnalysisRetain, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisDailyRetain(beginDate, endDate)
}

func (a *Analysis) GetAnalysisDailySummary(beginDate, endDate string) (result analysis.ResAnalysisDailySummary, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisDailySummary(beginDate, endDate)
}

func (a *Analysis) GetAnalysisDailyVisitTrend(beginDate, endDate string) (result analysis.ResAnalysisVisitTrend, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisDailyVisitTrend(beginDate, endDate)
}

func (a *Analysis) GetAnalysisMonthlyRetain(beginDate, endDate string) (result analysis.ResAnalysisRetain, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisMonthlyRetain(beginDate, endDate)
}

func (a *Analysis) GetAnalysisMonthlyVisitTrend(beginDate, endDate string) (result analysis.ResAnalysisVisitTrend, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisMonthlyVisitTrend(beginDate, endDate)
}

func (a *Analysis) GetAnalysisUserPortrait(beginDate, endDate string) (result analysis.ResAnalysisUserPortrait, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisUserPortrait(beginDate, endDate)
}

func (a *Analysis) GetAnalysisVisitDistribution(beginDate, endDate string) (result analysis.ResAnalysisVisitDistribution, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisVisitDistribution(beginDate, endDate)
}

func (a *Analysis) GetAnalysisVisitPage(beginDate, endDate string) (result analysis.ResAnalysisVisitPage, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisVisitPage(beginDate, endDate)
}

func (a *Analysis) GetAnalysisWeeklyRetain(beginDate, endDate string) (result analysis.ResAnalysisRetain, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisWeeklyRetain(beginDate, endDate)
}

func (a *Analysis) GetAnalysisWeeklyVisitTrend(beginDate, endDate string) (result analysis.ResAnalysisVisitTrend, err error) {
	return a.GetMp().GetAnalysis().GetAnalysisWeeklyVisitTrend(beginDate, endDate)
}
