package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL6I
}

type DAILRPL6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL6O
}

type DAILRPL6RequestForm struct {
	Form []DAILRPL6IDataForm
}

type DAILRPL6ResponseForm struct {
	Form []DAILRPL6ODataForm
}

type DAILRPL6I struct {
	ProcesId                        string  `json:"ProcesId"`
	LoanAcctNo                      string  `json:"LoanAcctNo"`
	AdjDt                           string  `json:"AdjDt"`
	AdjManrCd                       string  `json:"AdjManrCd"`
	AdjPersonOrgNo                  string  `json:"AdjPersonOrgNo"`
	AdjPersonEmpnbr                 string  `json:"AdjPersonEmpnbr"`
	UserNm                          string  `json:"UserNm"`
	ApprvEmpnbr                     string  `json:"ApprvEmpnbr"`
	ApprvNm                         string  `json:"ApprvNm"`
	TxSn                            string  `json:"TxSn"`
	TxTm                            string  `json:"TxTm"`
	AdjBefLoanRiskClsfCd            string  `json:"AdjBefLoanRiskClsfCd"`
	AdjAftLoanRiskClsfCd            string  `json:"AdjAftLoanRiskClsfCd"`
	InterAdjBefLoanRiskClsfCd		string  `json:"InterAdjBefLoanRiskClsfCd"`
	InterAdjAftLoanRiskClsfCd       string  `json:"InterAdjAftLoanRiskClsfCd"`
	AdjAftPrin                      float64 `json:"AdjAftPrin"`
	AdjAftIntr                      float64 `json:"AdjAftIntr"`
	AdjAftCmpdAmt                   float64 `json:"AdjAftCmpdAmt"`
	AdjAftPondg                     float64 `json:"AdjAftPondg"`
	AdjAftPnltint                   float64 `json:"AdjAftPnltint"`
	AdjAftCmpndint                  float64 `json:"AdjAftCmpndint"`
	AdjBefCustOvdueDays             int     `json:"AdjBefCustOvdueDays"`
	AdjAftCustOvdueDays             int     `json:"AdjAftCustOvdueDays"`
	FinlModfyDt                     string  `json:"FinlModfyDt"`
	FinlModfyTm                     string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                  string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                 string  `json:"FinlModfyTelrNo"`
	PageNo                          int     `json:"PageNo"`
	PageRecCount                    int     `json:"PageRecCount"`
}

type DAILRPL6O struct {
	Records                      []DAILRPL6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL6ORecords struct {
	ProcesId                        string  `json:"ProcesId"`
	LoanAcctNo                      string  `json:"LoanAcctNo"`
	AdjDt                           string  `json:"AdjDt"`
	AdjManrCd                       string  `json:"AdjManrCd"`
	AdjPersonOrgNo                  string  `json:"AdjPersonOrgNo"`
	AdjPersonEmpnbr                 string  `json:"AdjPersonEmpnbr"`
	UserNm                          string  `json:"UserNm"`
	ApprvEmpnbr                     string  `json:"ApprvEmpnbr"`
	ApprvNm                         string  `json:"ApprvNm"`
	TxSn                            string  `json:"TxSn"`
	TxTm                            string  `json:"TxTm"`
	AdjBefLoanRiskClsfCd            string  `json:"AdjBefLoanRiskClsfCd"`
	AdjAftLoanRiskClsfCd            string  `json:"AdjAftLoanRiskClsfCd"`
	InterAdjBefLoanRiskClsfCd		string  `json:"InterAdjBefLoanRiskClsfCd"`
	InterAdjAftLoanRiskClsfCd       string  `json:"InterAdjAftLoanRiskClsfCd"`
	AdjAftPrin                      float64 `json:"AdjAftPrin"`
	AdjAftIntr                      float64 `json:"AdjAftIntr"`
	AdjAftCmpdAmt                   float64 `json:"AdjAftCmpdAmt"`
	AdjAftPondg                     float64 `json:"AdjAftPondg"`
	AdjAftPnltint                   float64 `json:"AdjAftPnltint"`
	AdjAftCmpndint                  float64 `json:"AdjAftCmpndint"`
	AdjBefCustOvdueDays             int     `json:"AdjBefCustOvdueDays"`
	AdjAftCustOvdueDays             int     `json:"AdjAftCustOvdueDays"`
	FinlModfyDt                     string  `json:"FinlModfyDt"`
	FinlModfyTm                     string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                  string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                 string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRPL6RequestForm) PackRequest(dailrpl6I DAILRPL6I) (responseBody []byte, err error) {

	requestForm := DAILRPL6RequestForm{
		Form: []DAILRPL6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL6I",
				},

				FormData: dailrpl6I,
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
func (o *DAILRPL6RequestForm) UnPackRequest(request []byte) (DAILRPL6I, error) {
	dailrpl6I := DAILRPL6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL6ResponseForm) PackResponse(dailrpl6O DAILRPL6O) (responseBody []byte, err error) {
	responseForm := DAILRPL6ResponseForm{
		Form: []DAILRPL6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL6O",
				},
				FormData: dailrpl6O,
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
func (o *DAILRPL6ResponseForm) UnPackResponse(request []byte) (DAILRPL6O, error) {

	dailrpl6O := DAILRPL6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
