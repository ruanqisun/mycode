package models

import (
	"git.forms.io/legobank/legoapp/constant"
	"git.forms.io/legobank/legoapp/errors"
	"git.forms.io/universe/common/json"
)

type DAILRSX3IDataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX3I
}

type DAILRSX3ODataForm struct {
	FormHead CommonFormHead
	FormData DAILRSX3O
}

type DAILRSX3RequestForm struct {
	Form []DAILRSX3IDataForm
}

type DAILRSX3ResponseForm struct {
	Form []DAILRSX3ODataForm
}

type DAILRSX3I struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	TelNo                        string  `json:"TelNo"`
	MobileNo                     string  `json:"MobileNo"`
	GenderCd                     string  `json:"GenderCd"`
	Email                        string  `json:"Email"`
	StateAndRgnCd                string  `json:"StateAndRgnCd"`
	HousdRgstScAdcmCd            string  `json:"HousdRgstScAdcmCd"`
	HousdRgstDtlAddr             string  `json:"HousdRgstDtlAddr"`
	HousdRgstAddr                string  `json:"HousdRgstAddr"`
	CmunicAddrScAdcmCd           string  `json:"CmunicAddrScAdcmCd"`
	CotactDtlAddr                string  `json:"CotactDtlAddr"`
	CotactAddr                   string  `json:"CotactAddr"`
	CotactAddrPostCd             string  `json:"CotactAddrPostCd"`
	RsdnceScAdcmCd               string  `json:"RsdnceScAdcmCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	RsdnceAddr                   string  `json:"RsdnceAddr"`
	RsdncePostCd                 string  `json:"RsdncePostCd"`
	EdctbkgCd                    string  `json:"EdctbkgCd"`
	CareerTypCd                  string  `json:"CareerTypCd"`
	IndusTypCd                   string  `json:"IndusTypCd"`
	WorkUnitNm                   string  `json:"WorkUnitNm"`
	WorkUnitAddr                 string  `json:"WorkUnitAddr"`
	WorkUnitCharcCd              string  `json:"WorkUnitCharcCd"`
	DutyCd                       string  `json:"DutyCd"`
	DutyNm                       string  `json:"DutyNm"`
	UnitTelNo                    string  `json:"UnitTelNo"`
	MarrgSituationCd             string  `json:"MarrgSituationCd"`
	HavNotChildFlg               string  `json:"HavNotChildFlg"`
	BirthDt                      string  `json:"BirthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX3O struct {
	Records                      []DAILRSX3ORecords
	PageNo                       int  `json:"PageNo"`
	PageRecCount                 int  `json:"PageRecCount"`
}

type DAILRSX3ORecords struct {
	CrdtAplySn                   string  `json:"CrdtAplySn"`
	CustNo                       string  `json:"CustNo"`
	CustName                     string  `json:"CustName"`
	IndvCrtfTypCd                string  `json:"IndvCrtfTypCd"`
	IndvCrtfNo                   string  `json:"IndvCrtfNo"`
	TelNo                        string  `json:"TelNo"`
	MobileNo                     string  `json:"MobileNo"`
	GenderCd                     string  `json:"GenderCd"`
	Email                        string  `json:"Email"`
	StateAndRgnCd                string  `json:"StateAndRgnCd"`
	HousdRgstScAdcmCd            string  `json:"HousdRgstScAdcmCd"`
	HousdRgstDtlAddr             string  `json:"HousdRgstDtlAddr"`
	HousdRgstAddr                string  `json:"HousdRgstAddr"`
	CmunicAddrScAdcmCd           string  `json:"CmunicAddrScAdcmCd"`
	CotactDtlAddr                string  `json:"CotactDtlAddr"`
	CotactAddr                   string  `json:"CotactAddr"`
	CotactAddrPostCd             string  `json:"CotactAddrPostCd"`
	RsdnceScAdcmCd               string  `json:"RsdnceScAdcmCd"`
	RsdnceDtlAddr                string  `json:"RsdnceDtlAddr"`
	RsdnceAddr                   string  `json:"RsdnceAddr"`
	RsdncePostCd                 string  `json:"RsdncePostCd"`
	EdctbkgCd                    string  `json:"EdctbkgCd"`
	CareerTypCd                  string  `json:"CareerTypCd"`
	IndusTypCd                   string  `json:"IndusTypCd"`
	WorkUnitNm                   string  `json:"WorkUnitNm"`
	WorkUnitAddr                 string  `json:"WorkUnitAddr"`
	WorkUnitCharcCd              string  `json:"WorkUnitCharcCd"`
	DutyCd                       string  `json:"DutyCd"`
	DutyNm                       string  `json:"DutyNm"`
	UnitTelNo                    string  `json:"UnitTelNo"`
	MarrgSituationCd             string  `json:"MarrgSituationCd"`
	HavNotChildFlg               string  `json:"HavNotChildFlg"`
	BirthDt                      string  `json:"BirthDt"`
	FinlModfyDt                  string  `json:"FinlModfyDt"`
	FinlModfyTm                  string  `json:"FinlModfyTm"`
	FinlModfyOrgNo               string  `json:"FinlModfyOrgNo"`
	FinlModfyTelrNo              string  `json:"FinlModfyTelrNo"`
}

// @Desc Build request message
func (o *DAILRSX3RequestForm) PackRequest(dailrsx3I DAILRSX3I) (responseBody []byte, err error) {

	requestForm := DAILRSX3RequestForm{
		Form: []DAILRSX3IDataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX3I",
				},

				FormData: dailrsx3I,
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
func (o *DAILRSX3RequestForm) UnPackRequest(request []byte) (DAILRSX3I, error) {
	dailrsx3I := DAILRSX3I{}
	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx3I, errors.Wrap(err, 0, constant.REQUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx3I, errors.New("UnPackRequest failed.", constant.REQUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

// @Desc Build response message
func (o *DAILRSX3ResponseForm) PackResponse(dailrsx3O DAILRSX3O) (responseBody []byte, err error) {
	responseForm := DAILRSX3ResponseForm{
		Form: []DAILRSX3ODataForm{
			{
				FormHead: CommonFormHead{
					FormId: "DAILRSX3O",
				},
				FormData: dailrsx3O,
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
func (o *DAILRSX3ResponseForm) UnPackResponse(request []byte) (DAILRSX3O, error) {

	dailrsx3O := DAILRSX3O{}

	if err := json.Unmarshal(request, o); nil != err {
		return dailrsx3O, errors.Wrap(err, 0, constant.RSPUNPACKERR)
	}

	if len(o.Form) < 1 {
		return dailrsx3O, errors.New("UnPackResponse failed.", constant.RSPUNPACKERR)
	}

	return o.Form[0].FormData, nil
}

func (w *DAILRSX3I) Validate() error {
	validate := validator.New()
	return validate.Struct(w)
}
