package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRKH5IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH5I
}

type DAILRKH5ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH5O
}

type DAILRKH5RequestForm struct {
	Form []DAILRKH5IDataForm
}

type DAILRKH5ResponseForm struct {
	Form []DAILRKH5ODataForm
}

type DAILRKH5I struct {
	ProcesId                     string  `json:"ProcesId"`
	DelvrDt                      string  `json:"DelvrDt"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	MobileNo                     string  `json:"MobileNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	DelvrBefCustMgrEmpnbr        string  `json:"DelvrBefCustMgrEmpnbr"`
	DelvrBefOrgNo                string  `json:"DelvrBefOrgNo"`
	DelvrAftCustMgrEmpnbr        string  `json:"DelvrAftCustMgrEmpnbr"`
	DelvrAftOrgNo                string  `json:"DelvrAftOrgNo"`
	DelvrPrvlgTypCd              string  `json:"DelvrPrvlgTypCd"`
	DelvrPersonRelCd             string  `json:"DelvrPersonRelCd"`
	DelvrComnt                   string  `json:"DelvrComnt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRKH5O struct {
	Records                      []DAILRKH5ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRKH5ORecords struct {
	ProcesId                     string  `json:"ProcesId"`
	DelvrDt                      string  `json:"DelvrDt"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	MobileNo                     string  `json:"MobileNo"`
	LoanProdtNo                  string  `json:"LoanProdtNo"`
	LoanProdtVersNo              string  `json:"LoanProdtVersNo"`
	DelvrBefCustMgrEmpnbr        string  `json:"DelvrBefCustMgrEmpnbr"`
	DelvrBefOrgNo                string  `json:"DelvrBefOrgNo"`
	DelvrAftCustMgrEmpnbr        string  `json:"DelvrAftCustMgrEmpnbr"`
	DelvrAftOrgNo                string  `json:"DelvrAftOrgNo"`
	DelvrPrvlgTypCd              string  `json:"DelvrPrvlgTypCd"`
	DelvrPersonRelCd             string  `json:"DelvrPersonRelCd"`
	DelvrComnt                   string  `json:"DelvrComnt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRKH5RequestForm) PackRequest(dailrkh5I DAILRKH5I) (responseBody []byte, err error) {

	requestForm := DAILRKH5RequestForm{
		Form: []DAILRKH5IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH5I",
				},

				FormData: dailrkh5I,
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
func (o *DAILRKH5RequestForm) UnPackRequest(request []byte) (DAILRKH5I, error) {
	dailrkh5I := DAILRKH5I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh5I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh5I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRKH5ResponseForm) PackResponse(dailrkh5O DAILRKH5O) (responseBody []byte, err error) {
	responseForm := DAILRKH5ResponseForm{
		Form: []DAILRKH5ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH5O",
				},
				FormData: dailrkh5O,
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
func (o *DAILRKH5ResponseForm) UnPackResponse(request []byte) (DAILRKH5O, error) {

	dailrkh5O := DAILRKH5O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh5O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh5O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRKH5I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
