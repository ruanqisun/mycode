package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0018IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0018I
}

type IL8R0018ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0018O
}

type IL8R0018RequestForm struct {
	Form []IL8R0018IDataForm
}

type IL8R0018ResponseForm struct {
	Form []IL8R0018ODataForm
}

type IL8R0018I struct {
	AcctnDt                        string  `json:"AcctnDt"`
	SeqNo                          int  `json:"SeqNo"`
	TxTm                           string  `json:"TxTm"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	BizEventCd                     string  `json:"BizEventCd"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	FundFlgCd                      string  `json:"FundFlgCd"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	AccessChnlCd                   string  `json:"AccessChnlCd"`
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	AmtDrctCd                      string  `json:"AmtDrctCd"`
	ProdtNo                        string  `json:"ProdtNo"`
	EventUdfInfo                   string  `json:"EventUdfInfo"`
	CurCd                          string  `json:"CurCd"`
	Amount                         float64  `json:"Amount"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	WastStusCd                     string  `json:"WastStusCd"`
	AbstCd                         string  `json:"AbstCd"`
	WrtffFlg                       string  `json:"WrtffFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	TelrNo                         string  `json:"TelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0018O struct {
	Records                      []IL8R0018ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0018ORecords struct {
	AcctnDt                        string  `json:"AcctnDt"`
	SeqNo                          int  `json:"SeqNo"`
	TxTm                           string  `json:"TxTm"`
	BizFolnNo                      string  `json:"BizFolnNo"`
	BizEventCd                     string  `json:"BizEventCd"`
	SysFolnNo                      string  `json:"SysFolnNo"`
	FundFlgCd                      string  `json:"FundFlgCd"`
	TxOrgNo                        string  `json:"TxOrgNo"`
	AccessChnlCd                   string  `json:"AccessChnlCd"`
	SoftProdtCd                    string  `json:"SoftProdtCd"`
	ActngOrgNo                     string  `json:"ActngOrgNo"`
	AmtDrctCd                      string  `json:"AmtDrctCd"`
	ProdtNo                        string  `json:"ProdtNo"`
	EventUdfInfo                   string  `json:"EventUdfInfo"`
	CurCd                          string  `json:"CurCd"`
	Amount                         float64  `json:"Amount"`
	MansbjTypCd                    string  `json:"MansbjTypCd"`
	WastStusCd                     string  `json:"WastStusCd"`
	AbstCd                         string  `json:"AbstCd"`
	WrtffFlg                       string  `json:"WrtffFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	TelrNo                         string  `json:"TelrNo"`
}

// @Desc Build request message
func (o *IL8R0018RequestForm) PackRequest(il8r0018I IL8R0018I) (responseBody []byte, err error) {

	requestForm := IL8R0018RequestForm{
		Form: []IL8R0018IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0018I",
				},

				FormData: il8r0018I,
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
func (o *IL8R0018RequestForm) UnPackRequest(request []byte) (IL8R0018I, error) {
	il8r0018I := IL8R0018I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0018I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0018I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0018ResponseForm) PackResponse(il8r0018O IL8R0018O) (responseBody []byte, err error) {
	responseForm := IL8R0018ResponseForm{
		Form: []IL8R0018ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0018O",
				},
				FormData: il8r0018O,
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
func (o *IL8R0018ResponseForm) UnPackResponse(request []byte) (IL8R0018O, error) {

	il8r0018O := IL8R0018O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0018O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0018O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0018I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
