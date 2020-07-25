package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRKH4IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH4I
}

type DAILRKH4ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH4O
}

type DAILRKH4RequestForm struct {
	Form []DAILRKH4IDataForm
}

type DAILRKH4ResponseForm struct {
	Form []DAILRKH4ODataForm
}

type DAILRKH4I struct {
	CustMgrEmpnbr                string  `json:"CustMgrEmpnbr"`
	CorgCustMgrEmpnbr            string  `json:"CorgCustMgrEmpnbr"`
	CorgCustMgrName              string  `json:"CorgCustMgrName"`
	CustMgrName                  string  `json:"CustMgrName"`
	LprOrgNo                     string  `json:"LprOrgNo"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CustNo                       string  `json:"CustNo"`
	ExcutvPrvlgTypCd             string  `json:"ExcutvPrvlgTypCd"`
	BlngtoOrgNo                  string  `json:"BlngtoOrgNo"`
	RelBlngTypCd                 string  `json:"RelBlngTypCd"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CustCotactTelNo              string  `json:"CustCotactTelNo"`
	CrtDt                        string  `json:"CrtDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRKH4O struct {
	Records                      []DAILRKH4ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRKH4ORecords struct {
	CustMgrEmpnbr                string  `json:"CustMgrEmpnbr"`
	CorgCustMgrEmpnbr            string  `json:"CorgCustMgrEmpnbr"`
	CorgCustMgrName              string  `json:"CorgCustMgrName"`
	CustMgrName                  string  `json:"CustMgrName"`
	LprOrgNo                     string  `json:"LprOrgNo"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	CustNo                       string  `json:"CustNo"`
	ExcutvPrvlgTypCd             string  `json:"ExcutvPrvlgTypCd"`
	BlngtoOrgNo                  string  `json:"BlngtoOrgNo"`
	RelBlngTypCd                 string  `json:"RelBlngTypCd"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	CustCotactTelNo              string  `json:"CustCotactTelNo"`
	CrtDt                        string  `json:"CrtDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRKH4RequestForm) PackRequest(dailrkh4I DAILRKH4I) (responseBody []byte, err error) {

	requestForm := DAILRKH4RequestForm{
		Form: []DAILRKH4IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH4I",
				},

				FormData: dailrkh4I,
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
func (o *DAILRKH4RequestForm) UnPackRequest(request []byte) (DAILRKH4I, error) {
	dailrkh4I := DAILRKH4I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh4I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh4I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRKH4ResponseForm) PackResponse(dailrkh4O DAILRKH4O) (responseBody []byte, err error) {
	responseForm := DAILRKH4ResponseForm{
		Form: []DAILRKH4ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH4O",
				},
				FormData: dailrkh4O,
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
func (o *DAILRKH4ResponseForm) UnPackResponse(request []byte) (DAILRKH4O, error) {

	dailrkh4O := DAILRKH4O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh4O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh4O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRKH4I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
