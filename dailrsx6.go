package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX6IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX6I
}

type DAILRSX6ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX6O
}

type DAILRSX6RequestForm struct {
	Form []DAILRSX6IDataForm
}

type DAILRSX6ResponseForm struct {
	Form []DAILRSX6ODataForm
}

type DAILRSX6I struct {
	MakelnAplySn                 string  `json:"MakelnAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo    		             string  `json:"CustNo"`
	ConterCustNo                 string  `json:"ConterCustNo"`
	ConterNm                     string  `json:"ConterNm"`
	ConterTelNo                  string  `json:"ConterTelNo"`
	KinRelCd                     string  `json:"KinRelCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRSX6O struct {
	Records                      []DAILRSX6ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX6ORecords struct {
	MakelnAplySn                 string  `json:"MakelnAplySn"`
	SeqNo                        int     `json:"SeqNo"`
	CustNo    		             string  `json:"CustNo"`
	ConterCustNo                 string  `json:"ConterCustNo"`
	ConterNm                     string  `json:"ConterNm"`
	ConterTelNo                  string  `json:"ConterTelNo"`
	KinRelCd                     string  `json:"KinRelCd"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX6RequestForm) PackRequest(dailrsx6I DAILRSX6I) (responseBody []byte, err error) {

	requestForm := DAILRSX6RequestForm{
		Form: []DAILRSX6IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX6I",
				},

				FormData: dailrsx6I,
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
func (o *DAILRSX6RequestForm) UnPackRequest(request []byte) (DAILRSX6I, error) {
	dailrsx6I := DAILRSX6I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx6I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx6I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX6ResponseForm) PackResponse(dailrsx6O DAILRSX6O) (responseBody []byte, err error) {
	responseForm := DAILRSX6ResponseForm{
		Form: []DAILRSX6ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX6O",
				},
				FormData: dailrsx6O,
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
func (o *DAILRSX6ResponseForm) UnPackResponse(request []byte) (DAILRSX6O, error) {

	dailrsx6O := DAILRSX6O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx6O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx6O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX6I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
