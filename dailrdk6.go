package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRDK6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK6I
}

type DAILRDK6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRDK6O
}

type DAILRDK6RequestForm struct {
	Form []DAILRDK6IDataForm
}

type DAILRDK6ResponseForm struct {
	Form []DAILRDK6ODataForm
}

type DAILRDK6I struct {
	TxSn                         string  `json:"TxSn"`
	QtaChgSn                     string  `json:"QtaChgSn"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	AcctiAcctNo                  string  `json:"AcctiAcctNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	RepayAcctNo                  string  `json:"RepayAcctNo"`
	RepayTotlAmt                 float64 `json:"RepayTotlAmt"`
	RepayPrin                    float64 `json:"RepayPrin"`
	RepayIntr                    float64 `json:"RepayIntr"`
	RepayPnltint                 float64 `json:"RepayPnltint"`
	RepayCmpndint                float64 `json:"RepayCmpndint"`
	RepayCmpdAmt                 float64 `json:"RepayCmpdAmt"`
	RepbyTotlAmt                 float64 `json:"RepbyTotlAmt"`
	RepayPondg                   float64 `json:"RepayPondg"`
	RepayOthFee                  float64 `json:"RepayOthFee"`
	RepayCurCd                   string  `json:"RepayCurCd"`
	RepayStusCd                  string  `json:"RepayStusCd"`
	FailRsn                      string  `json:"FailRsn"`
	RepayFuncCd                  string  `json:"RepayFuncCd"`
	RepayAplyDt                  string  `json:"RepayAplyDt"`
	ActlRepayDt                  string  `json:"ActlRepayDt"`
	CrtTm                        string  `json:"CrtTm"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK6O struct {
	Records                      []DAILRDK6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRDK6ORecords struct {
	TxSn                         string  `json:"TxSn"`
	QtaChgSn                     string  `json:"QtaChgSn"`
	LoanDubilNo                  string  `json:"LoanDubilNo"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	AcctiAcctNo                  string  `json:"AcctiAcctNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	AcctnDt                      string  `json:"AcctnDt"`
	RepayAcctNo                  string  `json:"RepayAcctNo"`
	RepayTotlAmt                 float64 `json:"RepayTotlAmt"`
	RepayPrin                    float64 `json:"RepayPrin"`
	RepayIntr                    float64 `json:"RepayIntr"`
	RepayPnltint                 float64 `json:"RepayPnltint"`
	RepayCmpndint                float64 `json:"RepayCmpndint"`
	RepayCmpdAmt                 float64 `json:"RepayCmpdAmt"`
	RepbyTotlAmt                 float64 `json:"RepbyTotlAmt"`
	RepayPondg                   float64 `json:"RepayPondg"`
	RepayOthFee                  float64 `json:"RepayOthFee"`
	RepayCurCd                   string  `json:"RepayCurCd"`
	RepayStusCd                  string  `json:"RepayStusCd"`
	FailRsn                      string  `json:"FailRsn"`
	RepayFuncCd                  string  `json:"RepayFuncCd"`
	RepayAplyDt                  string  `json:"RepayAplyDt"`
	ActlRepayDt                  string  `json:"ActlRepayDt"`
	CrtTm                        string  `json:"CrtTm"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRDK6RequestForm) PackRequest(dailrdk6I DAILRDK6I) (responseBody []byte, err error) {

	requestForm := DAILRDK6RequestForm{
		Form: []DAILRDK6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK6I",
				},
				FormData: dailrdk6I,
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
func (o *DAILRDK6RequestForm) UnPackRequest(request []byte) (DAILRDK6I, error) {
	dailrdk6I := DAILRDK6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRDK6ResponseForm) PackResponse(dailrdk6O DAILRDK6O) (responseBody []byte, err error) {
	responseForm := DAILRDK6ResponseForm{
		Form: []DAILRDK6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRDK6O",
				},
				FormData: dailrdk6O,
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
func (o *DAILRDK6ResponseForm) UnPackResponse(request []byte) (DAILRDK6O, error) {

	dailrdk6O := DAILRDK6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrdk6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrdk6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRDK6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
