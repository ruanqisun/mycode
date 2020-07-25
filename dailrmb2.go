package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRMB2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB2I
}

type DAILRMB2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRMB2O
}

type DAILRMB2RequestForm struct {
	Form []DAILRMB2IDataForm
}

type DAILRMB2ResponseForm struct {
	Form []DAILRMB2ODataForm
}

type DAILRMB2I struct {
	TmplNo                       string  `json:"TmplNo"`
	EfftDt                       string  `json:"EfftDt"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	SuitHavegProdtFlg            string  `json:"SuitHavegProdtFlg"`
	TmplNm                       string  `json:"TmplNm"`
	TmplContent                  string  `json:"TmplContent"`
	TmplComnt                    string  `json:"TmplComnt"`
	TextTmplTypCd                string  `json:"TextTmplTypCd"`
	PmitSndFlg                   string  `json:"PmitSndFlg"`
	AdvSndDays                   int     `json:"AdvSndDays"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	SndBgnTm                     string  `json:"SndBgnTm"`
	SndEndTm                     string  `json:"SndEndTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRMB2O struct {
	Records                      []DAILRMB2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRMB2ORecords struct {
	TmplNo                       string  `json:"TmplNo"`
	EfftDt                       string  `json:"EfftDt"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	SuitHavegProdtFlg            string  `json:"SuitHavegProdtFlg"`
	TmplNm                       string  `json:"TmplNm"`
	TmplContent                  string  `json:"TmplContent"`
	TmplComnt                    string  `json:"TmplComnt"`
	TextTmplTypCd                string  `json:"TextTmplTypCd"`
	PmitSndFlg                   string  `json:"PmitSndFlg"`
	AdvSndDays                   int     `json:"AdvSndDays"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	SndBgnTm                     string  `json:"SndBgnTm"`
	SndEndTm                     string  `json:"SndEndTm"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRMB2RequestForm) PackRequest(dailrmb2I DAILRMB2I) (responseBody []byte, err error) {

	requestForm := DAILRMB2RequestForm{
		Form: []DAILRMB2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB2I",
				},

				FormData: dailrmb2I,
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
func (o *DAILRMB2RequestForm) UnPackRequest(request []byte) (DAILRMB2I, error) {
	dailrmb2I := DAILRMB2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRMB2ResponseForm) PackResponse(dailrmb2O DAILRMB2O) (responseBody []byte, err error) {
	responseForm := DAILRMB2ResponseForm{
		Form: []DAILRMB2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRMB2O",
				},
				FormData: dailrmb2O,
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
func (o *DAILRMB2ResponseForm) UnPackResponse(request []byte) (DAILRMB2O, error) {

	dailrmb2O := DAILRMB2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrmb2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrmb2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRMB2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
