package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type IL8R0011IDataForm struct {
	FormHead CommonFormHead
	FormData IL8R0011I
}

type IL8R0011ODataForm struct {
	FormHead CommonFormHead
	FormData IL8R0011O
}

type IL8R0011RequestForm struct {
	Form []IL8R0011IDataForm
}

type IL8R0011ResponseForm struct {
	Form []IL8R0011ODataForm
}

type IL8R0011I struct {
	CtrtNo                         string  `json:"CtrtNo"`
	CustNo                         string  `json:"CustNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	CtrtStusCd                     string  `json:"CtrtStusCd"`
	CtrtStartDt                    string  `json:"CtrtStartDt"`
	CtrtMatrDt                     string  `json:"CtrtMatrDt"`
	ApprvDt                        string  `json:"ApprvDt"`
	SignOrgNo                      string  `json:"SignOrgNo"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	ComslNo                        string  `json:"ComslNo"`
	TmplNo                         string  `json:"TmplNo"`
	CtrtImgNo                      string  `json:"CtrtImgNo"`
	BrwmnyCtrtTypCd                string  `json:"BrwmnyCtrtTypCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
	PageNo                         int     `json:"PageNo"`
	PageRecCount                   int     `json:"PageRecCount"`
}

type IL8R0011O struct {
	Records                      []IL8R0011ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type IL8R0011ORecords struct {
	CtrtNo                         string  `json:"CtrtNo"`
	CustNo                         string  `json:"CustNo"`
	IndvCrtfTypCd                  string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                     string  `json:"IndvCrtfNo"`
	CtrtStusCd                     string  `json:"CtrtStusCd"`
	CtrtStartDt                    string  `json:"CtrtStartDt"`
	CtrtMatrDt                     string  `json:"CtrtMatrDt"`
	ApprvDt                        string  `json:"ApprvDt"`
	SignOrgNo                      string  `json:"SignOrgNo"`
	OprOrgNo                       string  `json:"OprOrgNo"`
	OprorEmpnbr                    string  `json:"OprorEmpnbr"`
	ComslNo                        string  `json:"ComslNo"`
	TmplNo                         string  `json:"TmplNo"`
	CtrtImgNo                      string  `json:"CtrtImgNo"`
	BrwmnyCtrtTypCd                string  `json:"BrwmnyCtrtTypCd"`
	FinlModfyDt                    string  `json:"FinlModfyDt"`
	FinlModfyTm                    string  `json:"FinlModfyTm"`
	FinlModfyOrgNo                 string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo                string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *IL8R0011RequestForm) PackRequest(il8r0011I IL8R0011I) (responseBody []byte, err error) {

	requestForm := IL8R0011RequestForm{
		Form: []IL8R0011IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0011I",
				},

				FormData: il8r0011I,
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
func (o *IL8R0011RequestForm) UnPackRequest(request []byte) (IL8R0011I, error) {
	il8r0011I := IL8R0011I{}
	if err := json.Unmarshal(request, o); nil != err {
		return il8r0011I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0011I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *IL8R0011ResponseForm) PackResponse(il8r0011O IL8R0011O) (responseBody []byte, err error) {
	responseForm := IL8R0011ResponseForm{
		Form: []IL8R0011ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "IL8R0011O",
				},
				FormData: il8r0011O,
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
func (o *IL8R0011ResponseForm) UnPackResponse(request []byte) (IL8R0011O, error) {

	il8r0011O := IL8R0011O{}

	if err := json.Unmarshal(request, o); nil != err {
		return il8r0011O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return il8r0011O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *IL8R0011I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
