package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDH6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH6I
}

type DAILRDH6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDH6O
}

type DAILRDH6RequestForm struct {
	Form []DAILRDH6IDataForm
}

type DAILRDH6ResponseForm struct {
	Form []DAILRDH6ODataForm
}

type DAILRDH6I struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	SeqNo                        int     `json:"SeqNo"`
	DvalBizAplyNo                string  `json:"DvalBizAplyNo"`
	DvalDlwthStusCd              string  `json:"DvalDlwthStusCd"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustNm                       string  `json:"CustNm"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	GrantDt                      string  `json:"GrantDt"`
	MatrDt                       string  `json:"MatrDt"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	TranDvalComnt                string  `json:"TranDvalComnt"`
	DvalFlg                      string  `json:"DvalFlg"`
	DuePrin                      float64 `json:"DuePrin"`
	DuePrinBgnDt                 string  `json:"DuePrinBgnDt"`
	OvdueIntr                    float64 `json:"OvdueIntr"`
	IntrOvdueBgnDt               string  `json:"IntrOvdueBgnDt"`
	OvduePnltint                 float64 `json:"OvduePnltint"`
	PnltintBgnDt                 string  `json:"PnltintBgnDt"`
	DvalAplyDt                   string  `json:"DvalAplyDt"`
	AplyEmpnbr                   string  `json:"AplyEmpnbr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRDH6O struct {
	Records                      []DAILRDH6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDH6ORecords struct {
	LoanAcctNo                   string  `json:"LoanAcctNo"`
	SeqNo                        int     `json:"SeqNo"`
	DvalBizAplyNo                string  `json:"DvalBizAplyNo"`
	DvalDlwthStusCd              string  `json:"DvalDlwthStusCd"`
	CtrtNo                       string  `json:"CtrtNo"`
	CustNo                       string  `json:"CustNo"`
	CustNm                       string  `json:"CustNm"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	LoanProdtNm                  string  `json:"LoanProdtNm"`
	LoanGuarManrCd               string  `json:"LoanGuarManrCd"`
	GrantDt                      string  `json:"GrantDt"`
	MatrDt                       string  `json:"MatrDt"`
	RepayManrCd                  string  `json:"RepayManrCd"`
	LoanRiskClsfCd               string  `json:"LoanRiskClsfCd"`
	TranDvalComnt                string  `json:"TranDvalComnt"`
	DvalFlg                      string  `json:"DvalFlg"`
	DuePrin                      float64 `json:"DuePrin"`
	DuePrinBgnDt                 string  `json:"DuePrinBgnDt"`
	OvdueIntr                    float64 `json:"OvdueIntr"`
	IntrOvdueBgnDt               string  `json:"IntrOvdueBgnDt"`
	OvduePnltint                 float64 `json:"OvduePnltint"`
	PnltintBgnDt                 string  `json:"PnltintBgnDt"`
	DvalAplyDt                   string  `json:"DvalAplyDt"`
	AplyEmpnbr                   string  `json:"AplyEmpnbr"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDH6RequestForm) PackRequest(dailrdh6I DAILRDH6I) (responseBody []byte, err error) {

	requestForm := DAILRDH6RequestForm{
		Form: []DAILRDH6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH6I",
				},

				FormData: dailrdh6I,
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
func (o *DAILRDH6RequestForm) UnPackRequest(request []byte) (DAILRDH6I, error) {
	dailrdh6I := DAILRDH6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDH6ResponseForm) PackResponse(dailrdh6O DAILRDH6O) (responseBody []byte, err error) {
	responseForm := DAILRDH6ResponseForm{
		Form: []DAILRDH6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDH6O",
				},
				FormData: dailrdh6O,
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
func (o *DAILRDH6ResponseForm) UnPackResponse(request []byte) (DAILRDH6O, error) {

	dailrdh6O := DAILRDH6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdh6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdh6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDH6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
