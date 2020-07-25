package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRHT2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRHT2I
}

type DAILRHT2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRHT2O
}

type DAILRHT2RequestForm struct {
	Form []DAILRHT2IDataForm
}

type DAILRHT2ResponseForm struct {
	Form []DAILRHT2ODataForm
}

type DAILRHT2I struct {
	CtrtNo                       string  `json:"CtrtNo"`
	SeqNo                        int  `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	TxSn                         string  `json:"TxSn"`
	NodeNm                       string  `json:"NodeNm"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperBefCtrtStusCd            string  `json:"OperBefCtrtStusCd"`
	OperAftCtrtStusCd            string  `json:"OperAftCtrtStusCd"`
	AfltInfo                     string  `json:"AfltInfo"`
	OperEmpnbr                   string  `json:"OperEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	OperResultCd                 string  `json:"OperResultCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRHT2O struct {
	Records                      []DAILRHT2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRHT2ORecords struct {
	CtrtNo                       string  `json:"CtrtNo"`
	SeqNo                        int  `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	TxSn                         string  `json:"TxSn"`
	NodeNm                       string  `json:"NodeNm"`
	OperTypCd                    string  `json:"OperTypCd"`
	OperBefCtrtStusCd            string  `json:"OperBefCtrtStusCd"`
	OperAftCtrtStusCd            string  `json:"OperAftCtrtStusCd"`
	AfltInfo                     string  `json:"AfltInfo"`
	OperEmpnbr                   string  `json:"OperEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	OperResultCd                 string  `json:"OperResultCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRHT2RequestForm) PackRequest(dailrht2I DAILRHT2I) (responseBody []byte, err error) {

	requestForm := DAILRHT2RequestForm{
		Form: []DAILRHT2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRHT2I",
				},

				FormData: dailrht2I,
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
func (o *DAILRHT2RequestForm) UnPackRequest(request []byte) (DAILRHT2I, error) {
	dailrht2I := DAILRHT2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrht2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrht2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRHT2ResponseForm) PackResponse(dailrht2O DAILRHT2O) (responseBody []byte, err error) {
	responseForm := DAILRHT2ResponseForm{
		Form: []DAILRHT2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRHT2O",
				},
				FormData: dailrht2O,
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
func (o *DAILRHT2ResponseForm) UnPackResponse(request []byte) (DAILRHT2O, error) {

	dailrht2O := DAILRHT2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrht2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrht2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRHT2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
