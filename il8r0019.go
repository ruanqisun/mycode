package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0019IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0019I
}

type IL8R0019ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0019O
}

type IL8R0019RequestForm struct {
	Form []IL8R0019IDataForm
}

type IL8R0019ResponseForm struct {
	Form []IL8R0019ODataForm
}

type IL8R0019I struct {
	SubjNo                         string  `json:"SubjNo"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SubjEngNm                      string  `json:"SubjEngNm"`
	SubjNm                         string  `json:"SubjNm"`
	SubjCategCd                    string  `json:"SubjCategCd"`
	KepacctManrCd                  string  `json:"KepacctManrCd"`
	AtmtcGnrtGlFlg                 string  `json:"AtmtcGnrtGlFlg"`
	DtlSubjFlg                     string  `json:"DtlSubjFlg"`
	SubjBalZeroFlgCd               string  `json:"SubjBalZeroFlgCd"`
	BalDrctCd                      string  `json:"BalDrctCd"`
	SubjStusCd                     string  `json:"SubjStusCd"`
	SubjOpenAcctDt                 string  `json:"SubjOpenAcctDt"`
	EfftDt                         string  `json:"EfftDt"`
	SubjHeavyOpenDt                string  `json:"SubjHeavyOpenDt"`
	DeregisDt                      string  `json:"DeregisDt"`
	InvalidDt                      string  `json:"InvalidDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0019O struct {
	Records                      []IL8R0019ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0019ORecords struct {
	SubjNo                         string  `json:"SubjNo"`
	AcctBookCategCd                string  `json:"AcctBookCategCd"`
	SubjEngNm                      string  `json:"SubjEngNm"`
	SubjNm                         string  `json:"SubjNm"`
	SubjCategCd                    string  `json:"SubjCategCd"`
	KepacctManrCd                  string  `json:"KepacctManrCd"`
	AtmtcGnrtGlFlg                 string  `json:"AtmtcGnrtGlFlg"`
	DtlSubjFlg                     string  `json:"DtlSubjFlg"`
	SubjBalZeroFlgCd               string  `json:"SubjBalZeroFlgCd"`
	BalDrctCd                      string  `json:"BalDrctCd"`
	SubjStusCd                     string  `json:"SubjStusCd"`
	SubjOpenAcctDt                 string  `json:"SubjOpenAcctDt"`
	EfftDt                         string  `json:"EfftDt"`
	SubjHeavyOpenDt                string  `json:"SubjHeavyOpenDt"`
	DeregisDt                      string  `json:"DeregisDt"`
	InvalidDt                      string  `json:"InvalidDt"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0019RequestForm) PackRequest(il8r0019I IL8R0019I) (responseBody []byte, err error) {

	requestForm := IL8R0019RequestForm{
		Form: []IL8R0019IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0019I",
				},

				FormData: il8r0019I,
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
func (o *IL8R0019RequestForm) UnPackRequest(request []byte) (IL8R0019I, error) {
	il8r0019I := IL8R0019I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0019I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0019I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0019ResponseForm) PackResponse(il8r0019O IL8R0019O) (responseBody []byte, err error) {
	responseForm := IL8R0019ResponseForm{
		Form: []IL8R0019ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0019O",
				},
				FormData: il8r0019O,
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
func (o *IL8R0019ResponseForm) UnPackResponse(request []byte) (IL8R0019O, error) {

	il8r0019O := IL8R0019O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0019O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0019O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0019I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
