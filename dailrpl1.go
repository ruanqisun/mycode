package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRPL1IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL1I
}

type DAILRPL1ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRPL1O
}

type DAILRPL1RequestForm struct {
	Form []DAILRPL1IDataForm
}

type DAILRPL1ResponseForm struct {
	Form []DAILRPL1ODataForm
}

type DAILRPL1I struct {
	JobNo                        string  `json:"JobNo"`
	JobNm                        string  `json:"JobNm"`
	JobBgnExecDt                 string  `json:"JobBgnExecDt"`
	JobBgnExecTm                 string  `json:"JobBgnExecTm"`
	JobEndExecDt                 string  `json:"JobEndExecDt"`
	JobEndExecTm                 string  `json:"JobEndExecTm"`
	CurtyJobStusCd               string  `json:"CurtyJobStusCd"`
	JobCanRepetExecFlg           string  `json:"JobCanRepetExecFlg"`
	PgmCdnmAssmblg               string  `json:"PgmCdnmAssmblg"`
	BizDt                        string  `json:"BizDt"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRPL1O struct {
	Records                      []DAILRPL1ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRPL1ORecords struct {
	JobNo                        string  `json:"JobNo"`
	JobNm                        string  `json:"JobNm"`
	JobBgnExecDt                 string  `json:"JobBgnExecDt"`
	JobBgnExecTm                 string  `json:"JobBgnExecTm"`
	JobEndExecDt                 string  `json:"JobEndExecDt"`
	JobEndExecTm                 string  `json:"JobEndExecTm"`
	CurtyJobStusCd               string  `json:"CurtyJobStusCd"`
	JobCanRepetExecFlg           string  `json:"JobCanRepetExecFlg"`
	PgmCdnmAssmblg               string  `json:"PgmCdnmAssmblg"`
	BizDt                        string  `json:"BizDt"`
}

// @Desc Build request message
func (o *DAILRPL1RequestForm) PackRequest(dailrpl1I DAILRPL1I) (responseBody []byte, err error) {

	requestForm := DAILRPL1RequestForm{
		Form: []DAILRPL1IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL1I",
				},

				FormData: dailrpl1I,
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
func (o *DAILRPL1RequestForm) UnPackRequest(request []byte) (DAILRPL1I, error) {
	dailrpl1I := DAILRPL1I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl1I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl1I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRPL1ResponseForm) PackResponse(dailrpl1O DAILRPL1O) (responseBody []byte, err error) {
	responseForm := DAILRPL1ResponseForm{
		Form: []DAILRPL1ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRPL1O",
				},
				FormData: dailrpl1O,
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
func (o *DAILRPL1ResponseForm) UnPackResponse(request []byte) (DAILRPL1O, error) {

	dailrpl1O := DAILRPL1O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrpl1O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrpl1O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRPL1I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
