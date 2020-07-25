package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRFK2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRFK2I
}

type DAILRFK2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRFK2O
}

type DAILRFK2RequestForm struct {
	Form []DAILRFK2IDataForm
}

type DAILRFK2ResponseForm struct {
	Form []DAILRFK2ODataForm
}

type DAILRFK2I struct {
	MakelnAplySn                 string  `json:"MakelnAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	NodeNm                       string  `json:"NodeNm"`
	TxSn                         string  `json:"TxSn"`
	ApprvNodeTypCd               string  `json:"ApprvNodeTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	MakelnApprvOperResultCd      string  `json:"MakelnApprvOperResultCd"`
	AfltInfo                     string  `json:"AfltInfo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRFK2O struct {
	Records                      []DAILRFK2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRFK2ORecords struct {
	MakelnAplySn                 string  `json:"MakelnAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo                       string  `json:"CustNo"`
	NodeNm                       string  `json:"NodeNm"`
	TxSn                         string  `json:"TxSn"`
	ApprvNodeTypCd               string  `json:"ApprvNodeTypCd"`
	OperPersonEmpnbr             string  `json:"OperPersonEmpnbr"`
	OperOrgNo                    string  `json:"OperOrgNo"`
	OperTm                       string  `json:"OperTm"`
	MakelnApprvOperResultCd      string  `json:"MakelnApprvOperResultCd"`
	AfltInfo                     string  `json:"AfltInfo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRFK2RequestForm) PackRequest(dailrfk2I DAILRFK2I) (responseBody []byte, err error) {

	requestForm := DAILRFK2RequestForm{
		Form: []DAILRFK2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRFK2I",
				},

				FormData: dailrfk2I,
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
func (o *DAILRFK2RequestForm) UnPackRequest(request []byte) (DAILRFK2I, error) {
	dailrfk2I := DAILRFK2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrfk2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrfk2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRFK2ResponseForm) PackResponse(dailrfk2O DAILRFK2O) (responseBody []byte, err error) {
	responseForm := DAILRFK2ResponseForm{
		Form: []DAILRFK2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRFK2O",
				},
				FormData: dailrfk2O,
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
func (o *DAILRFK2ResponseForm) UnPackResponse(request []byte) (DAILRFK2O, error) {

	dailrfk2O := DAILRFK2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrfk2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrfk2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRFK2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
