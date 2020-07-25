package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRGG1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRGG1I
}

type DAILRGG1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRGG1O
}

type DAILRGG1RequestForm struct {
	Form []DAILRGG1IDataForm
}

type DAILRGG1ResponseForm struct {
	Form []DAILRGG1ODataForm
}

type DAILRGG1I struct {
	TmplNo                         string  `json:"TmplNo"`
	TmplNm                         string  `json:"TmplNm"`
	TmplContent                    string  `json:"TmplContent"`
	TmplComnt                      string  `json:"TmplComnt"`
	TmplTypCd                      string  `json:"TmplTypCd"`
	PmitSndFlg                     string  `json:"PmitSndFlg"`
	EfftDt                         string  `json:"EfftDt"`
	AdvSndDays                     int     `json:"AdvSndDays"`
	SndBgnTm                       string  `json:"SndBgnTm"`
	SndEndTm                       string  `json:"SndEndTm"`
	RiskClsfValidFlg               string  `json:"RiskClsfValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type DAILRGG1O struct {
	Records                      []DAILRGG1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRGG1ORecords struct {
	TmplNo                         string  `json:"TmplNo"`
	TmplNm                         string  `json:"TmplNm"`
	TmplContent                    string  `json:"TmplContent"`
	TmplComnt                      string  `json:"TmplComnt"`
	TmplTypCd                      string  `json:"TmplTypCd"`
	PmitSndFlg                     string  `json:"PmitSndFlg"`
	EfftDt                         string  `json:"EfftDt"`
	AdvSndDays                     int     `json:"AdvSndDays"`
	SndBgnTm                       string  `json:"SndBgnTm"`
	SndEndTm                       string  `json:"SndEndTm"`
	RiskClsfValidFlg               string  `json:"RiskClsfValidFlg"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRGG1RequestForm) PackRequest(dailrgg1I DAILRGG1I) (responseBody []byte, err error) {

	requestForm := DAILRGG1RequestForm{
		Form: []DAILRGG1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRGG1I",
				},

				FormData: dailrgg1I,
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
func (o *DAILRGG1RequestForm) UnPackRequest(request []byte) (DAILRGG1I, error) {
	dailrgg1I := DAILRGG1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrgg1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrgg1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRGG1ResponseForm) PackResponse(dailrgg1O DAILRGG1O) (responseBody []byte, err error) {
	responseForm := DAILRGG1ResponseForm{
		Form: []DAILRGG1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRGG1O",
				},
				FormData: dailrgg1O,
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
func (o *DAILRGG1ResponseForm) UnPackResponse(request []byte) (DAILRGG1O, error) {

	dailrgg1O := DAILRGG1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrgg1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrgg1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRGG1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
