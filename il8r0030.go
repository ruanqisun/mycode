package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0030IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0030I
}

type IL8R0030ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0030O
}

type IL8R0030RequestForm struct {
	Form []IL8R0030IDataForm
}

type IL8R0030ResponseForm struct {
	Form []IL8R0030ODataForm
}

type IL8R0030I struct {
	MakelnAplySn                   string  `json:"MakelnAplySn"`
	SeqNo                          int  `json:"SeqNo"`
	CustNo                         string  `json:"CustNo"`
	DcmkApprvResultCd              string  `json:"DcmkApprvResultCd"`
	DcmkApprvDescr                 string  `json:"DcmkApprvDescr"`
	HitExcludRuleComnt             string  `json:"HitExcludRuleComnt"`
	BlklistTypCd                   string  `json:"BlklistTypCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr               string  `json:"FinlModifrEmpnbr"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0030O struct {
	Records                      []IL8R0030ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0030ORecords struct {
	MakelnAplySn                   string  `json:"MakelnAplySn"`
	SeqNo                          int  `json:"SeqNo"`
	CustNo                         string  `json:"CustNo"`
	DcmkApprvResultCd              string  `json:"DcmkApprvResultCd"`
	DcmkApprvDescr                 string  `json:"DcmkApprvDescr"`
	HitExcludRuleComnt             string  `json:"HitExcludRuleComnt"`
	BlklistTypCd                   string  `json:"BlklistTypCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModifrEmpnbr               string  `json:"FinlModifrEmpnbr"`
}

// @Desc Build request message
func (o *IL8R0030RequestForm) PackRequest(il8r0030I IL8R0030I) (responseBody []byte, err error) {

	requestForm := IL8R0030RequestForm{
		Form: []IL8R0030IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0030I",
				},

				FormData: il8r0030I,
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
func (o *IL8R0030RequestForm) UnPackRequest(request []byte) (IL8R0030I, error) {
	il8r0030I := IL8R0030I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0030I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0030I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0030ResponseForm) PackResponse(il8r0030O IL8R0030O) (responseBody []byte, err error) {
	responseForm := IL8R0030ResponseForm{
		Form: []IL8R0030ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0030O",
				},
				FormData: il8r0030O,
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
func (o *IL8R0030ResponseForm) UnPackResponse(request []byte) (IL8R0030O, error) {

	il8r0030O := IL8R0030O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0030O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0030O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0030I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
