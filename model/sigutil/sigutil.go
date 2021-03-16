//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package sigutil ;import (_d "bytes";_g "crypto";_cg "crypto/x509";_bf "encoding/asn1";_f "encoding/pem";_df "errors";_e "fmt";_aa "github.com/unidoc/timestamp";_ff "github.com/unidoc/unipdf/v3/common";_cbe "golang.org/x/crypto/ocsp";_b "io";_ce "io/ioutil";
_cb "net/http";_a "time";);

// IsCA returns true if the provided certificate appears to be a CA certificate.
func (_db *CertClient )IsCA (cert *_cg .Certificate )bool {return cert .IsCA &&_d .Equal (cert .RawIssuer ,cert .RawSubject );};

// NewTimestampClient returns a new timestamp client.
func NewTimestampClient ()*TimestampClient {return &TimestampClient {HTTPClient :_ffe ()}};

// NewTimestampRequest returns a new timestamp request based
// on the specified options.
func NewTimestampRequest (body _b .Reader ,opts *_aa .RequestOptions )(*_aa .Request ,error ){if opts ==nil {opts =&_aa .RequestOptions {};};if opts .Hash ==0{opts .Hash =_g .SHA256 ;};if !opts .Hash .Available (){return nil ,_cg .ErrUnsupportedAlgorithm ;
};_bd :=opts .Hash .New ();if _ ,_eaf :=_b .Copy (_bd ,body );_eaf !=nil {return nil ,_eaf ;};return &_aa .Request {HashAlgorithm :opts .Hash ,HashedMessage :_bd .Sum (nil ),Certificates :opts .Certificates ,TSAPolicyOID :opts .TSAPolicyOID ,Nonce :opts .Nonce },nil ;
};

// Get retrieves the certificate at the specified URL.
func (_bb *CertClient )Get (url string )(*_cg .Certificate ,error ){if _bb .HTTPClient ==nil {_bb .HTTPClient =_ffe ();};_dfa ,_ee :=_bb .HTTPClient .Get (url );if _ee !=nil {return nil ,_ee ;};defer _dfa .Body .Close ();_fe ,_ee :=_ce .ReadAll (_dfa .Body );
if _ee !=nil {return nil ,_ee ;};if _ceb ,_ :=_f .Decode (_fe );_ceb !=nil {_fe =_ceb .Bytes ;};_aaf ,_ee :=_cg .ParseCertificate (_fe );if _ee !=nil {return nil ,_ee ;};return _aaf ,nil ;};

// GetEncodedToken executes the timestamp request and returns the DER encoded
// timestamp token bytes.
func (_gg *TimestampClient )GetEncodedToken (serverURL string ,req *_aa .Request )([]byte ,error ){if serverURL ==""{return nil ,_e .Errorf ("\u006d\u0075\u0073\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061m\u0070\u0020\u0073\u0065\u0072\u0076\u0065r\u0020\u0055\u0052\u004c");
};if req ==nil {return nil ,_e .Errorf ("\u0074\u0069\u006de\u0073\u0074\u0061\u006dp\u0020\u0072\u0065\u0071\u0075\u0065\u0073t\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006e\u0069\u006c");};_fff ,_feg :=req .Marshal ();if _feg !=nil {return nil ,_feg ;
};_eae :=_gg .HTTPClient ;if _eae ==nil {_eae =_ffe ();};_gc ,_feg :=_eae .Post (serverURL ,"a\u0070\u0070\u006c\u0069\u0063\u0061t\u0069\u006f\u006e\u002f\u0074\u0069\u006d\u0065\u0073t\u0061\u006d\u0070-\u0071u\u0065\u0072\u0079",_d .NewBuffer (_fff ));
if _feg !=nil {return nil ,_feg ;};defer _gc .Body .Close ();_gdg ,_feg :=_ce .ReadAll (_gc .Body );if _feg !=nil {return nil ,_feg ;};if _gc .StatusCode !=_cb .StatusOK {return nil ,_e .Errorf ("\u0075\u006e\u0065x\u0070\u0065\u0063\u0074e\u0064\u0020\u0048\u0054\u0054\u0050\u0020s\u0074\u0061\u0074\u0075\u0073\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0064",_gc .StatusCode );
};var _ed struct{Version _bf .RawValue ;Content _bf .RawValue ;};if _ ,_feg =_bf .Unmarshal (_gdg ,&_ed );_feg !=nil {return nil ,_feg ;};return _ed .Content .FullBytes ,nil ;};

// OCSPClient represents a OCSP (Online Certificate Status Protocol) client.
// It is used to request revocation data from OCSP servers.
type OCSPClient struct{

// HTTPClient is the HTTP client used to make OCSP requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_cb .Client ;

// Hash is the hash function  used when constructing the OCSP
// requests. If zero, SHA-1 will be used.
Hash _g .Hash ;};

// NewOCSPClient returns a new OCSP client.
func NewOCSPClient ()*OCSPClient {return &OCSPClient {HTTPClient :_ffe (),Hash :_g .SHA1 }};

// TimestampClient represents a RFC 3161 timestamp client.
// It is used to obtain signed tokens from timestamp authority servers.
type TimestampClient struct{

// HTTPClient is the HTTP client used to make timestamp requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_cb .Client ;};func _ffe ()*_cb .Client {return &_cb .Client {Timeout :5*_a .Second }};

// MakeRequest makes a CRL request to the specified server and returns the
// response. If a server URL is not provided, it is extracted from the certificate.
func (_cd *CRLClient )MakeRequest (serverURL string ,cert *_cg .Certificate )([]byte ,error ){if _cd .HTTPClient ==nil {_cd .HTTPClient =_ffe ();};if serverURL ==""{if len (cert .CRLDistributionPoints )==0{return nil ,_df .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063\u0061t\u0065\u0020\u0064o\u0065\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0070ec\u0069\u0066\u0079 \u0061\u006ey\u0020\u0043\u0052\u004c\u0020\u0073e\u0072\u0076e\u0072\u0073");
};serverURL =cert .CRLDistributionPoints [0];};_de ,_gda :=_cd .HTTPClient .Get (serverURL );if _gda !=nil {return nil ,_gda ;};defer _de .Body .Close ();_fb ,_gda :=_ce .ReadAll (_de .Body );if _gda !=nil {return nil ,_gda ;};if _ceg ,_ :=_f .Decode (_fb );
_ceg !=nil {_fb =_ceg .Bytes ;};return _fb ,nil ;};

// CertClient represents a X.509 certificate client. Its primary purpose
// is to download certificates.
type CertClient struct{

// HTTPClient is the HTTP client used to make certificate requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_cb .Client ;};

// NewCertClient returns a new certificate client.
func NewCertClient ()*CertClient {return &CertClient {HTTPClient :_ffe ()}};

// MakeRequest makes a OCSP request to the specified server and returns
// the parsed and raw responses. If a server URL is not provided, it is
// extracted from the certificate.
func (_ge *OCSPClient )MakeRequest (serverURL string ,cert ,issuer *_cg .Certificate )(*_cbe .Response ,[]byte ,error ){if _ge .HTTPClient ==nil {_ge .HTTPClient =_ffe ();};if serverURL ==""{if len (cert .OCSPServer )==0{return nil ,nil ,_df .New ("\u0063e\u0072\u0074i\u0066\u0069\u0063a\u0074\u0065\u0020\u0064\u006f\u0065\u0073 \u006e\u006f\u0074\u0020\u0073\u0070e\u0063\u0069\u0066\u0079\u0020\u0061\u006e\u0079\u0020\u004f\u0043S\u0050\u0020\u0073\u0065\u0072\u0076\u0065\u0072\u0073");
};serverURL =cert .OCSPServer [0];};_ac ,_ecg :=_cbe .CreateRequest (cert ,issuer ,&_cbe .RequestOptions {Hash :_ge .Hash });if _ecg !=nil {return nil ,nil ,_ecg ;};_eb ,_ecg :=_ge .HTTPClient .Post (serverURL ,"\u0061p\u0070\u006c\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u006fc\u0073\u0070\u002d\u0072\u0065\u0071\u0075\u0065\u0073\u0074",_d .NewReader (_ac ));
if _ecg !=nil {return nil ,nil ,_ecg ;};defer _eb .Body .Close ();_cdg ,_ecg :=_ce .ReadAll (_eb .Body );if _ecg !=nil {return nil ,nil ,_ecg ;};if _ca ,_ :=_f .Decode (_cdg );_ca !=nil {_cdg =_ca .Bytes ;};_dec ,_ecg :=_cbe .ParseResponseForCert (_cdg ,cert ,issuer );
if _ecg !=nil {return nil ,nil ,_ecg ;};return _dec ,_cdg ,nil ;};

// CRLClient represents a CRL (Certificate revocation list) client.
// It is used to request revocation data from CRL servers.
type CRLClient struct{

// HTTPClient is the HTTP client used to make CRL requests.
// By default, an HTTP client with a 5 second timeout per request is used.
HTTPClient *_cb .Client ;};

// NewCRLClient returns a new CRL client.
func NewCRLClient ()*CRLClient {return &CRLClient {HTTPClient :_ffe ()}};

// GetIssuer retrieves the issuer of the provided certificate.
func (_ffg *CertClient )GetIssuer (cert *_cg .Certificate )(*_cg .Certificate ,error ){for _ ,_bfa :=range cert .IssuingCertificateURL {_be ,_dg :=_ffg .Get (_bfa );if _dg !=nil {_ff .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074 \u0064\u006f\u0077\u006e\u006c\u006f\u0061\u0064\u0020\u0069\u0073\u0073\u0075e\u0072\u0020\u0066\u006f\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066ic\u0061\u0074\u0065\u0020\u0025\u0076\u003a\u0020\u0025\u0076",cert .Subject .CommonName ,_dg );
continue ;};return _be ,nil ;};return nil ,_e .Errorf ("\u0069\u0073\u0073\u0075e\u0072\u0020\u0063\u0065\u0072\u0074\u0069\u0066\u0069\u0063a\u0074e\u0020\u006e\u006f\u0074\u0020\u0066\u006fu\u006e\u0064");};