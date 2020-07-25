package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0031IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0031I
}

type IL8R0031ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0031O
}

type IL8R0031RequestForm struct {
	Form []IL8R0031IDataForm
}

type IL8R0031ResponseForm struct {
	Form []IL8R0031ODataForm
}

type IL8R0031I struct {
	BizSn	                       string  `json:"BizSn"`
	TxAcctnDt                      string  `json:"TxAcctnDt"`
	TxTm                           string  `json:"TxTm"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	TxTelrNo                       string  `json:"TxTelrNo"`
	AuthTelrNo                     string  `json:"AuthTelrNo"`
	TxLunchChnlCd                  string  `json:"TxLunchChnlCd"`
	AccessChnlCd                   string  `json:"AccessChnlCd"`
	BizSysSoftProdtCd              string  `json:"BizSysSoftProdtCd"`
	DebitCrdtFlg                   string  `json:"DebitCrdtFlg"`
	CardNoOrAcctNo                 string  `json:"CardNoOrAcctNo"`
	CntptyAcctNo                   string  `json:"CntptyAcctNo"`
	AcctiAcctNo                    string  `json:"AcctiAcctNo"`
	TxCurCd                        string  `json:"TxCurCd"`
	Amount                         float64  `json:"Amount"`
	OrginlTxAcctnDt                string  `json:"OrginlTxAcctnDt"`
	OrginlBizFolnNo                string  `json:"OrginlBizFolnNo"`
	OrginlTxTelrNo                 string  `json:"OrginlTxTelrNo"`
	TxStusCd                       string  `json:"TxStusCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0031O struct {
	Records                      []IL8R0031ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0031ORecords struct {
	BizSn	                       string  `json:"BizSn"`
	TxAcctnDt                      string  `json:"TxAcctnDt"`
	TxTm                           string  `json:"TxTm"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	TxTelrNo                       string  `json:"TxTelrNo"`
	AuthTelrNo                     string  `json:"AuthTelrNo"`
	TxLunchChnlCd                  string  `json:"TxLunchChnlCd"`
	AccessChnlCd                   string  `json:"AccessChnlCd"`
	BizSysSoftProdtCd              string  `json:"BizSysSoftProdtCd"`
	DebitCrdtFlg                   string  `json:"DebitCrdtFlg"`
	CardNoOrAcctNo                 string  `json:"CardNoOrAcctNo"`
	CntptyAcctNo                   string  `json:"CntptyAcctNo"`
	AcctiAcctNo                    string  `json:"AcctiAcctNo"`
	TxCurCd                        string  `json:"TxCurCd"`
	Amount                         float64  `json:"Amount"`
	OrginlTxAcctnDt                string  `json:"OrginlTxAcctnDt"`
	OrginlBizFolnNo                string  `json:"OrginlBizFolnNo"`
	OrginlTxTelrNo                 string  `json:"OrginlTxTelrNo"`
	TxStusCd                       string  `json:"TxStusCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0031RequestForm) PackRequest(il8r0031I IL8R0031I) (responseBody []byte, err error) {

	requestForm := IL8R0031RequestForm{
		Form: []IL8R0031IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0031I",
				},

				FormData: il8r0031I,
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
func (o *IL8R0031RequestForm) UnPackRequest(request []byte) (IL8R0031I, error) {
	il8r0031I := IL8R0031I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0031I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0031I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0031ResponseForm) PackResponse(il8r0031O IL8R0031O) (responseBody []byte, err error) {
	responseForm := IL8R0031ResponseForm{
		Form: []IL8R0031ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0031O",
				},
				FormData: il8r0031O,
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
func (o *IL8R0031ResponseForm) UnPackResponse(request []byte) (IL8R0031O, error) {

	il8r0031O := IL8R0031O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0031O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0031O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0031I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
