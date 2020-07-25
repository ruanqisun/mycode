package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0020IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0020I
}

type IL8R0020ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0020O
}

type IL8R0020RequestForm struct {
	Form []IL8R0020IDataForm
}

type IL8R0020ResponseForm struct {
	Form []IL8R0020ODataForm
}

type IL8R0020I struct {
	JrnlzId                        string  `json:"JrnlzId"`
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxTm                           string  `json:"TxTm"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	CurCd                          string  `json:"CurCd"`
	DebitCrdtFlgCd                 string  `json:"DebitCrdtFlgCd"`
	TxAmt                          float64  `json:"TxAmt"`
	SubjNo                         string  `json:"SubjNo"`
	ReqePtySoftProdtCd             string  `json:"ReqePtySoftProdtCd"`
	KepacctSuccFlg                 string  `json:"KepacctSuccFlg"`
	AbstCd                         string  `json:"AbstCd"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SumId                          string  `json:"SumId"`
	WrtffFlg                       string  `json:"WrtffFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0020O struct {
	Records                      []IL8R0020ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0020ORecords struct {
	JrnlzId                        string  `json:"JrnlzId"`
	AcctnDt                        string  `json:"AcctnDt"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxTm                           string  `json:"TxTm"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	CurCd                          string  `json:"CurCd"`
	DebitCrdtFlgCd                 string  `json:"DebitCrdtFlgCd"`
	TxAmt                          float64  `json:"TxAmt"`
	SubjNo                         string  `json:"SubjNo"`
	ReqePtySoftProdtCd             string  `json:"ReqePtySoftProdtCd"`
	KepacctSuccFlg                 string  `json:"KepacctSuccFlg"`
	AbstCd                         string  `json:"AbstCd"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SumId                          string  `json:"SumId"`
	WrtffFlg                       string  `json:"WrtffFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0020RequestForm) PackRequest(il8r0020I IL8R0020I) (responseBody []byte, err error) {

	requestForm := IL8R0020RequestForm{
		Form: []IL8R0020IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0020I",
				},

				FormData: il8r0020I,
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
func (o *IL8R0020RequestForm) UnPackRequest(request []byte) (IL8R0020I, error) {
	il8r0020I := IL8R0020I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0020I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0020I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0020ResponseForm) PackResponse(il8r0020O IL8R0020O) (responseBody []byte, err error) {
	responseForm := IL8R0020ResponseForm{
		Form: []IL8R0020ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0020O",
				},
				FormData: il8r0020O,
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
func (o *IL8R0020ResponseForm) UnPackResponse(request []byte) (IL8R0020O, error) {

	il8r0020O := IL8R0020O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0020O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0020O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0020I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
