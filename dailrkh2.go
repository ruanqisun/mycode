package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRKH2IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH2I
}

type DAILRKH2ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRKH2O
}

type DAILRKH2RequestForm struct {
	Form []DAILRKH2IDataForm
}

type DAILRKH2ResponseForm struct {
	Form []DAILRKH2ODataForm
}

type DAILRKH2I struct {
	OrgNo                        string  `json:"OrgNo"`
	SeqNo                        int     `json:"SeqNo"`
	BlklistNo                    string  `json:"BlklistNo"`
	BizSn                        string  `json:"BizSn"`
	SorcDescr                    string  `json:"SorcDescr"`
	BlklistTypCd                 string  `json:"BlklistTypCd"`
	CustNmlstNewaddTypCd         string  `json:"CustNmlstNewaddTypCd"`
	RecmndEmpnbr                 string  `json:"RecmndEmpnbr"`
	RecmndNm	                 string  `json:"RecmndNm"`
	MobileNo                     string  `json:"MobileNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	BlklistCustClsfCd            string  `json:"BlklistCustClsfCd"`
	BckltRsnCd                   string  `json:"BckltRsnCd"`
	BckltRsnComnt                string  `json:"BckltRsnComnt"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	DeregisRsn                   string  `json:"DeregisRsn"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int     `json:"PageNo"`
	PageRecCount                 int     `json:"PageRecCount"`
}

type DAILRKH2O struct {
	Records                      []DAILRKH2ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRKH2ORecords struct {
	OrgNo                        string  `json:"OrgNo"`
	SeqNo                        int     `json:"SeqNo"`
	BlklistNo                    string  `json:"BlklistNo"`
	BizSn                        string  `json:"BizSn"`
	SorcDescr                    string  `json:"SorcDescr"`
	BlklistTypCd                 string  `json:"BlklistTypCd"`
	CustNmlstNewaddTypCd         string  `json:"CustNmlstNewaddTypCd"`
	RecmndEmpnbr                 string  `json:"RecmndEmpnbr"`
	RecmndNm	                 string  `json:"RecmndNm"`
	MobileNo                     string  `json:"MobileNo"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	BlklistCustClsfCd            string  `json:"BlklistCustClsfCd"`
	BckltRsnCd                   string  `json:"BckltRsnCd"`
	BckltRsnComnt                string  `json:"BckltRsnComnt"`
	KeprcdStusCd                 string  `json:"KeprcdStusCd"`
	DeregisRsn                   string  `json:"DeregisRsn"`
	CrtDt                        string  `json:"CrtDt"`
	CrtEmpnbr                    string  `json:"CrtEmpnbr"`
	CrtOrgNo                     string  `json:"CrtOrgNo"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRKH2RequestForm) PackRequest(dailrkh2I DAILRKH2I) (responseBody []byte, err error) {

	requestForm := DAILRKH2RequestForm{
		Form: []DAILRKH2IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH2I",
				},

				FormData: dailrkh2I,
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
func (o *DAILRKH2RequestForm) UnPackRequest(request []byte) (DAILRKH2I, error) {
	dailrkh2I := DAILRKH2I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh2I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh2I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRKH2ResponseForm) PackResponse(dailrkh2O DAILRKH2O) (responseBody []byte, err error) {
	responseForm := DAILRKH2ResponseForm{
		Form: []DAILRKH2ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRKH2O",
				},
				FormData: dailrkh2O,
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
func (o *DAILRKH2ResponseForm) UnPackResponse(request []byte) (DAILRKH2O, error) {

	dailrkh2O := DAILRKH2O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrkh2O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrkh2O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRKH2I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
