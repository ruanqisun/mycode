package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0014IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0014I
}

type IL8R0014ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0014O
}

type IL8R0014RequestForm struct {
	Form []IL8R0014IDataForm
}

type IL8R0014ResponseForm struct {
	Form []IL8R0014ODataForm
}

type IL8R0014I struct {
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	AcctBookStusCd                 string  `json:"AcctBookStusCd"`
	AcctBookSeqNo                  int  `json:"AcctBookSeqNo"`
	SubjLength                     int  `json:"SubjLength"`
	OnshetPngdgacctSubjNo          string  `json:"OnshetPngdgacctSubjNo"`
	OffshetPngdgacctSubjNo         string  `json:"OffshetPngdgacctSubjNo"`
	UnDistrMarginSubjNo            string  `json:"UnDistrMarginSubjNo"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	EngineVaritNm                  string  `json:"EngineVaritNm"`
	InterLiqdSubjNo                string  `json:"InterLiqdSubjNo"`
	EfftDt                         string  `json:"EfftDt"`
	InvalidDt                      string  `json:"InvalidDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0014O struct {
	Records                      []IL8R0014ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0014ORecords struct {
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	AcctBookStusCd                 string  `json:"AcctBookStusCd"`
	AcctBookSeqNo                  int  `json:"AcctBookSeqNo"`
	SubjLength                     int  `json:"SubjLength"`
	OnshetPngdgacctSubjNo          string  `json:"OnshetPngdgacctSubjNo"`
	OffshetPngdgacctSubjNo         string  `json:"OffshetPngdgacctSubjNo"`
	UnDistrMarginSubjNo            string  `json:"UnDistrMarginSubjNo"`
	EngineVaritCd                  string  `json:"EngineVaritCd"`
	EngineVaritNm                  string  `json:"EngineVaritNm"`
	InterLiqdSubjNo                string  `json:"InterLiqdSubjNo"`
	EfftDt                         string  `json:"EfftDt"`
	InvalidDt                      string  `json:"InvalidDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0014RequestForm) PackRequest(il8r0014I IL8R0014I) (responseBody []byte, err error) {

	requestForm := IL8R0014RequestForm{
		Form: []IL8R0014IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0014I",
				},

				FormData: il8r0014I,
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
func (o *IL8R0014RequestForm) UnPackRequest(request []byte) (IL8R0014I, error) {
	il8r0014I := IL8R0014I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0014I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0014I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0014ResponseForm) PackResponse(il8r0014O IL8R0014O) (responseBody []byte, err error) {
	responseForm := IL8R0014ResponseForm{
		Form: []IL8R0014ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0014O",
				},
				FormData: il8r0014O,
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
func (o *IL8R0014ResponseForm) UnPackResponse(request []byte) (IL8R0014O, error) {

	il8r0014O := IL8R0014O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0014O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0014O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0014I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
