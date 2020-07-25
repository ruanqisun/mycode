package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT6I
}

type DAILRXT6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT6O
}

type DAILRXT6RequestForm struct {
	Form []DAILRXT6IDataForm
}

type DAILRXT6ResponseForm struct {
	Form []DAILRXT6ODataForm
}

type DAILRXT6I struct {
	Empnbr                       string  `json:"Empnbr"`
	EmplyName                    string  `json:"EmplyName"`
	UserOrgNo                    string  `json:"UserOrgNo"`
	MtguarStusCd                 string  `json:"MtguarStusCd"`
	LgnVouchCd                   string  `json:"LgnVouchCd"`
	RecntMtguarDt                string  `json:"RecntMtguarDt"`
	RecntMtguarTm                string  `json:"RecntMtguarTm"`
	RecntLeaveOfficeDt           string  `json:"RecntLeaveOfficeDt"`
	RecntLeaveOfficeTm           string  `json:"RecntLeaveOfficeTm"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT6O struct {
	Records                      []DAILRXT6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT6ORecords struct {
	Empnbr                       string  `json:"Empnbr"`
	EmplyName                    string  `json:"EmplyName"`
	UserOrgNo                    string  `json:"UserOrgNo"`
	MtguarStusCd                 string  `json:"MtguarStusCd"`
	LgnVouchCd                   string  `json:"LgnVouchCd"`
	RecntMtguarDt                string  `json:"RecntMtguarDt"`
	RecntMtguarTm                string  `json:"RecntMtguarTm"`
	RecntLeaveOfficeDt           string  `json:"RecntLeaveOfficeDt"`
	RecntLeaveOfficeTm           string  `json:"RecntLeaveOfficeTm"`
}

// @Desc Build request message
func (o *DAILRXT6RequestForm) PackRequest(dailrxt6I DAILRXT6I) (responseBody []byte, err error) {

	requestForm := DAILRXT6RequestForm{
		Form: []DAILRXT6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT6I",
				},

				FormData: dailrxt6I,
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
func (o *DAILRXT6RequestForm) UnPackRequest(request []byte) (DAILRXT6I, error) {
	dailrxt6I := DAILRXT6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT6ResponseForm) PackResponse(dailrxt6O DAILRXT6O) (responseBody []byte, err error) {
	responseForm := DAILRXT6ResponseForm{
		Form: []DAILRXT6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT6O",
				},
				FormData: dailrxt6O,
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
func (o *DAILRXT6ResponseForm) UnPackResponse(request []byte) (DAILRXT6O, error) {

	dailrxt6O := DAILRXT6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
