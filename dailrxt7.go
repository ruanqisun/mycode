package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRXT7IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT7I
}

type DAILRXT7ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRXT7O
}

type DAILRXT7RequestForm struct {
	Form []DAILRXT7IDataForm
}

type DAILRXT7ResponseForm struct {
	Form []DAILRXT7ODataForm
}

type DAILRXT7I struct {
	TaskExampId                  string  `json:"TaskExampId"`
	TaskNm                       string  `json:"TaskNm"`
	BizId                        string  `json:"BizId"`
	ProcesExampId                string  `json:"ProcesExampId"`
	DistrPersonEmpnbr            string  `json:"DistrPersonEmpnbr"`
	CanExctrRoleNo               string  `json:"CanExctrRoleNo"`
	CrtTm                        string  `json:"CrtTm"`
	OperTm                       string  `json:"OperTm"`
	ProcesTmplId                 string  `json:"ProcesTmplId"`
	CorgTaskStusCd               string  `json:"CorgTaskStusCd"`
	CorgTaskExampId              string  `json:"CorgTaskExampId"`
	CustNo                       string  `json:"CustNo"`
	TaskTypCd                    string  `json:"TaskTypCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LprOrgNo                     string  `json:"LprOrgNo"`
	BizOrgNo                     string  `json:"BizOrgNo"`
	SticDt                       string  `json:"SticDt"`
	SticTm                       string  `json:"SticTm"`
	TaskTmplId                   string  `json:"TaskTmplId"`
	MenuFuncId                   string  `json:"MenuFuncId"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	CrtDt                        string  `json:"CrtDt"`
	OperDt                       string  `json:"OperDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRXT7O struct {
	Records                      []DAILRXT7ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRXT7ORecords struct {
	TaskExampId                  string  `json:"TaskExampId"`
	TaskNm                       string  `json:"TaskNm"`
	BizId                        string  `json:"BizId"`
	ProcesExampId                string  `json:"ProcesExampId"`
	DistrPersonEmpnbr            string  `json:"DistrPersonEmpnbr"`
	CanExctrRoleNo               string  `json:"CanExctrRoleNo"`
	CrtTm                        string  `json:"CrtTm"`
	OperTm                       string  `json:"OperTm"`
	ProcesTmplId                 string  `json:"ProcesTmplId"`
	CorgTaskStusCd               string  `json:"CorgTaskStusCd"`
	CorgTaskExampId              string  `json:"CorgTaskExampId"`
	CustNo                       string  `json:"CustNo"`
	TaskTypCd                    string  `json:"TaskTypCd"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LprOrgNo                     string  `json:"LprOrgNo"`
	BizOrgNo                     string  `json:"BizOrgNo"`
	SticDt                       string  `json:"SticDt"`
	SticTm                       string  `json:"SticTm"`
	TaskTmplId                   string  `json:"TaskTmplId"`
	MenuFuncId                   string  `json:"MenuFuncId"`
	TaskDlwthStusCd              string  `json:"TaskDlwthStusCd"`
	CrtDt                        string  `json:"CrtDt"`
	OperDt                       string  `json:"OperDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRXT7RequestForm) PackRequest(dailrxt7I DAILRXT7I) (responseBody []byte, err error) {

	requestForm := DAILRXT7RequestForm{
		Form: []DAILRXT7IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT7I",
				},

				FormData: dailrxt7I,
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
func (o *DAILRXT7RequestForm) UnPackRequest(request []byte) (DAILRXT7I, error) {
	dailrxt7I := DAILRXT7I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt7I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt7I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRXT7ResponseForm) PackResponse(dailrxt7O DAILRXT7O) (responseBody []byte, err error) {
	responseForm := DAILRXT7ResponseForm{
		Form: []DAILRXT7ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRXT7O",
				},
				FormData: dailrxt7O,
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
func (o *DAILRXT7ResponseForm) UnPackResponse(request []byte) (DAILRXT7O, error) {

	dailrxt7O := DAILRXT7O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrxt7O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrxt7O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRXT7I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
