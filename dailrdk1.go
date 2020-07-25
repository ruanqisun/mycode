package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK1I
}

type DAILRDK1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK1O
}

type DAILRDK1RequestForm struct {
	Form []DAILRDK1IDataForm
}

type DAILRDK1ResponseForm struct {
	Form []DAILRDK1ODataForm
}

type DAILRDK1I struct {
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	AcctiAcctNo                    string  `json:"AcctiAcctNo"`
	ContrtStusCd                   string  `json:"ContrtStusCd"`
	CustNo                         string  `json:"CustNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	LoanProdtVersNo                string  `json:"LoanProdtVersNo"`
	MakelnOrgNo                    string  `json:"MakelnOrgNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	RevneCmpdCd                    string  `json:"RevneCmpdCd"`
	MakelnManrCd                   string  `json:"MakelnManrCd"`
	RelaAcctDtrmnManrCd            string  `json:"RelaAcctDtrmnManrCd"`
	RepayDayDtrmnManrCd            string  `json:"RepayDayDtrmnManrCd"`
	LoanDeadl	               	   int     `json:"LoanDeadl"`
	LoanDeadlCycCd                 string  `json:"LoanDeadlCycCd"`
	PmitAdvRepayTms                int     `json:"PmitAdvRepayTms"`
	AdvRepayLimitBgnDt             string  `json:"AdvRepayLimitBgnDt"`
	AdvRepayLimitDays              int     `json:"AdvRepayLimitDays"`
	LimitTrmInsidPmitPartlRepayFlg string  `json:"LimitTrmInsidPmitPartlRepayFlg"`
	LimitTrmInsidPmitPayOffFlg     string  `json:"LimitTrmInsidPmitPayOffFlg"`
	OpenAcctDt                     string  `json:"OpenAcctDt"`
	FsttmForsprtDt                 string  `json:"FsttmForsprtDt"`
	BegintDt                       string  `json:"BegintDt"`
	OrgnlMatrDt                    string  `json:"OrgnlMatrDt"`
	CurCd                          string  `json:"CurCd"`
	LoanAmt                        float64 `json:"LoanAmt"`
	EmbFlg                         string  `json:"EmbFlg"`
	DecdEmbDt                      string  `json:"DecdEmbDt"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	CtrtNo                         string  `json:"CtrtNo"`
	LoanGuarManrCd                 string  `json:"LoanGuarManrCd"`
	OthConsmTypCd                  string  `json:"OthConsmTypCd"`
	RepayManrCd                    string  `json:"RepayManrCd"`
	IntStlDayDtrmnManrCd           string  `json:"IntStlDayDtrmnManrCd"`
	AdvRepayTms                    int     `json:"AdvRepayTms"`
	TranNormlLoanDt                string  `json:"TranNormlLoanDt"`
	BldInstltRepayFlg              string  `json:"BldInstltRepayFlg"`
	IntStlPrtyTypCd                string  `json:"IntStlPrtyTypCd"`
	LoanIntrtAdjCycCd              string  `json:"LoanIntrtAdjCycCd"`
	LoanIntrtAdjCycQty             int     `json:"LoanIntrtAdjCycQty"`
	StpDeductFlg                   string  `json:"StpDeductFlg"`
	StpDeductRsnTypCd              string  `json:"StpDeductRsnTypCd"`
	StpDeductCnfrmDt               string  `json:"StpDeductCnfrmDt"`
	LoanIntrtAdjManrCd             string  `json:"LoanIntrtAdjManrCd"`
	ExpdayPayOffManrCd             string  `json:"ExpdayPayOffManrCd"`
	GraceTrmIntacrFlg              string  `json:"GraceTrmIntacrFlg"`
	EvrpridMaxGraceTrmDays         int     `json:"EvrpridMaxGraceTrmDays"`
	ContrtPrdGraceTrmTotlDays      int     `json:"ContrtPrdGraceTrmTotlDays"`
	AdvRepayColtfeFlg              string  `json:"AdvRepayColtfeFlg"`
	ColtfeManrCd                   string  `json:"ColtfeManrCd"`
	SglColtfeAmt                   float64 `json:"SglColtfeAmt"`
	ColtfeAmtCrdnlnbrCd            string  `json:"ColtfeAmtCrdnlnbrCd"`
	ColtfeRatio                    float64 `json:"ColtfeRatio"`
	AcrdgRatioColtfeCeilAmt        float64 `json:"AcrdgRatioColtfeCeilAmt"`
	AcrdgRatioColtfeFloorAmt       float64 `json:"AcrdgRatioColtfeFloorAmt"`
	CurrExecTmprd                  int     `json:"CurrExecTmprd"`
	RepayPlanAdjFlg                string  `json:"RepayPlanAdjFlg"`
	RestFlg                        string  `json:"RestFlg"`
	RestDt                         string  `json:"RestDt"`
	TranDvalDt                     string  `json:"TranDvalDt"`
	TranDvalFlg                    string  `json:"TranDvalFlg"`
	ExtsnTms                       int     `json:"ExtsnTms"`
	WaitExtsnFlg                   string  `json:"WaitExtsnFlg"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	MatrDt                         string  `json:"MatrDt"`
	PayOffDt                       string  `json:"PayOffDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int  `json:"PageNo"`
	PageRecCount                   int  `json:"PageRecCount"`
}

type DAILRDK1O struct {
	Records                        []DAILRDK1ORecords
	PageNo                         int  `json:"PageNo"`
	PageRecCount                   int  `json:"PageRecCount"`
}

type DAILRDK1ORecords struct {
	LoanDubilNo                    string  `json:"LoanDubilNo"`
	AcctiAcctNo                    string  `json:"AcctiAcctNo"`
	ContrtStusCd                   string  `json:"ContrtStusCd"`
	CustNo                         string  `json:"CustNo"`
	LoanProdtNo                    string  `json:"LoanProdtNo"`
	LoanProdtVersNo                string  `json:"LoanProdtVersNo"`
	MakelnOrgNo                    string  `json:"MakelnOrgNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	RevneCmpdCd                    string  `json:"RevneCmpdCd"`
	MakelnManrCd                   string  `json:"MakelnManrCd"`
	RelaAcctDtrmnManrCd            string  `json:"RelaAcctDtrmnManrCd"`
	RepayDayDtrmnManrCd            string  `json:"RepayDayDtrmnManrCd"`
	LoanDeadl	               	   int     `json:"LoanDeadl"`
	LoanDeadlCycCd                 string  `json:"LoanDeadlCycCd"`
	PmitAdvRepayTms                int     `json:"PmitAdvRepayTms"`
	AdvRepayLimitBgnDt             string  `json:"AdvRepayLimitBgnDt"`
	AdvRepayLimitDays              int     `json:"AdvRepayLimitDays"`
	LimitTrmInsidPmitPartlRepayFlg string  `json:"LimitTrmInsidPmitPartlRepayFlg"`
	LimitTrmInsidPmitPayOffFlg     string  `json:"LimitTrmInsidPmitPayOffFlg"`
	OpenAcctDt                     string  `json:"OpenAcctDt"`
	FsttmForsprtDt                 string  `json:"FsttmForsprtDt"`
	BegintDt                       string  `json:"BegintDt"`
	OrgnlMatrDt                    string  `json:"OrgnlMatrDt"`
	CurCd                          string  `json:"CurCd"`
	LoanAmt                        float64 `json:"LoanAmt"`
	EmbFlg                         string  `json:"EmbFlg"`
	DecdEmbDt                      string  `json:"DecdEmbDt"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	CtrtNo                         string  `json:"CtrtNo"`
	LoanGuarManrCd                 string  `json:"LoanGuarManrCd"`
	OthConsmTypCd                  string  `json:"OthConsmTypCd"`
	RepayManrCd                    string  `json:"RepayManrCd"`
	IntStlDayDtrmnManrCd           string  `json:"IntStlDayDtrmnManrCd"`
	AdvRepayTms                    int     `json:"AdvRepayTms"`
	TranNormlLoanDt                string  `json:"TranNormlLoanDt"`
	BldInstltRepayFlg              string  `json:"BldInstltRepayFlg"`
	IntStlPrtyTypCd                string  `json:"IntStlPrtyTypCd"`
	LoanIntrtAdjCycCd              string  `json:"LoanIntrtAdjCycCd"`
	LoanIntrtAdjCycQty             int     `json:"LoanIntrtAdjCycQty"`
	StpDeductFlg                   string  `json:"StpDeductFlg"`
	StpDeductRsnTypCd              string  `json:"StpDeductRsnTypCd"`
	StpDeductCnfrmDt               string  `json:"StpDeductCnfrmDt"`
	LoanIntrtAdjManrCd             string  `json:"LoanIntrtAdjManrCd"`
	ExpdayPayOffManrCd             string  `json:"ExpdayPayOffManrCd"`
	GraceTrmIntacrFlg              string  `json:"GraceTrmIntacrFlg"`
	EvrpridMaxGraceTrmDays         int     `json:"EvrpridMaxGraceTrmDays"`
	ContrtPrdGraceTrmTotlDays      int     `json:"ContrtPrdGraceTrmTotlDays"`
	AdvRepayColtfeFlg              string  `json:"AdvRepayColtfeFlg"`
	ColtfeManrCd                   string  `json:"ColtfeManrCd"`
	SglColtfeAmt                   float64 `json:"SglColtfeAmt"`
	ColtfeAmtCrdnlnbrCd            string  `json:"ColtfeAmtCrdnlnbrCd"`
	ColtfeRatio                    float64 `json:"ColtfeRatio"`
	AcrdgRatioColtfeCeilAmt        float64 `json:"AcrdgRatioColtfeCeilAmt"`
	AcrdgRatioColtfeFloorAmt       float64 `json:"AcrdgRatioColtfeFloorAmt"`
	CurrExecTmprd                  int     `json:"CurrExecTmprd"`
	RepayPlanAdjFlg                string  `json:"RepayPlanAdjFlg"`
	RestFlg                        string  `json:"RestFlg"`
	RestDt                         string  `json:"RestDt"`
	TranDvalDt                     string  `json:"TranDvalDt"`
	TranDvalFlg                    string  `json:"TranDvalFlg"`
	ExtsnTms                       int     `json:"ExtsnTms"`
	WaitExtsnFlg                   string  `json:"WaitExtsnFlg"`
	CurrLoanRiskClsfCd             string  `json:"CurrLoanRiskClsfCd"`
	MatrDt                         string  `json:"MatrDt"`
	PayOffDt                       string  `json:"PayOffDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK1RequestForm) PackRequest(dailrdk1I DAILRDK1I) (responseBody []byte, err error) {

	requestForm := DAILRDK1RequestForm{
		Form: []DAILRDK1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK1I",
				},
				FormData: dailrdk1I,
			},
		},
	}

	responseBody, err = json.Marshal(requestForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.REQPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing request message
func (o *DAILRDK1RequestForm) UnPackRequest(request []byte) (DAILRDK1I, error) {
	dailrdk1I := DAILRDK1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK1ResponseForm) PackResponse(dailrdk1O DAILRDK1O) (responseBody []byte, err error) {
	responseForm := DAILRDK1ResponseForm{
		Form: []DAILRDK1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK1O",
				},
				FormData: dailrdk1O,
			},
		},
	}

	responseBody, err = json.Marshal(responseForm)
	if err != nil {
		return nil, errors.Wrap(err, 0, constant.RSPPACKERR)
	}

	return responseBody, nil
}

// @Desc Parsing response message
func (o *DAILRDK1ResponseForm) UnPackResponse(request []byte) (DAILRDK1O, error) {

	dailrdk1O := DAILRDK1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
