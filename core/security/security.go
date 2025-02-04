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

package security ;import (_b "bytes";_cc "crypto/aes";_f "crypto/cipher";_g "crypto/md5";_fb "crypto/rand";_ef "crypto/rc4";_d "crypto/sha256";_fc "crypto/sha512";_ccc "encoding/binary";_fcf "errors";_bd "fmt";_a "github.com/unidoc/unipdf/v3/common";_eb "hash";
_c "io";_efd "math";);func _dc (_ae _f .Block )_f .BlockMode {return (*ecbEncrypter )(_ec (_ae ))};func (_ace stdHandlerR6 )alg8 (_bee *StdEncryptDict ,_efef []byte ,_dga []byte )error {if _gbc :=_dcd ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_efef );
_gbc !=nil {return _gbc ;};var _fcg [16]byte ;if _ ,_ede :=_c .ReadFull (_fb .Reader ,_fcg [:]);_ede !=nil {return _ede ;};_gdd :=_fcg [0:8];_ddf :=_fcg [8:16];_dgg :=make ([]byte ,len (_dga )+len (_gdd ));_ddg :=copy (_dgg ,_dga );copy (_dgg [_ddg :],_gdd );
_gfa ,_gbfe :=_ace .alg2b (_bee .R ,_dgg ,_dga ,nil );if _gbfe !=nil {return _gbfe ;};U :=make ([]byte ,len (_gfa )+len (_gdd )+len (_ddf ));_ddg =copy (U ,_gfa [:32]);_ddg +=copy (U [_ddg :],_gdd );copy (U [_ddg :],_ddf );_bee .U =U ;_ddg =len (_dga );
copy (_dgg [_ddg :],_ddf );_gfa ,_gbfe =_ace .alg2b (_bee .R ,_dgg ,_dga ,nil );if _gbfe !=nil {return _gbfe ;};_gee ,_gbfe :=_cb (_gfa [:32]);if _gbfe !=nil {return _gbfe ;};_ccb :=make ([]byte ,_cc .BlockSize );_aaa :=_f .NewCBCEncrypter (_gee ,_ccb );
UE :=make ([]byte ,32);_aaa .CryptBlocks (UE ,_efef [:32]);_bee .UE =UE ;return nil ;};

// Authenticate implements StdHandler interface.
func (_gaf stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _gaf .alg2a (d ,pass );};var _ StdHandler =stdHandlerR6 {};type ecb struct{_bb _f .Block ;_ee int ;};type ecbDecrypter ecb ;

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};func (_fff stdHandlerR4 )alg6 (_gd *StdEncryptDict ,_decb []byte )([]byte ,error ){var (_fbef []byte ;_efe error ;);_ebg :=_fff .alg2 (_gd ,_decb );if _gd .R ==2{_fbef ,_efe =_fff .alg4 (_ebg ,_decb );
}else if _gd .R >=3{_fbef ,_efe =_fff .alg5 (_ebg ,_decb );}else {return nil ,_fcf .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _efe !=nil {return nil ,_efe ;};_a .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_fbef ),string (_gd .U ));
_ac :=_fbef ;_be :=_gd .U ;if _gd .R >=3{if len (_ac )> 16{_ac =_ac [0:16];};if len (_be )> 16{_be =_be [0:16];};};if !_b .Equal (_ac ,_be ){return nil ,nil ;};return _ebg ,nil ;};var _ StdHandler =stdHandlerR4 {};type ecbEncrypter ecb ;func _dcd (_gfe ,_dcb string ,_adb int ,_cff []byte )error {if len (_cff )< _adb {return errInvalidField {Func :_gfe ,Field :_dcb ,Exp :_adb ,Got :len (_cff )};
};return nil ;};func (_dff stdHandlerR6 )alg11 (_cgb *StdEncryptDict ,_cgdeb []byte )([]byte ,error ){if _bfg :=_dcd ("\u0061\u006c\u00671\u0031","\u0055",48,_cgb .U );_bfg !=nil {return nil ,_bfg ;};_ecdf :=make ([]byte ,len (_cgdeb )+8);_egaf :=copy (_ecdf ,_cgdeb );
_egaf +=copy (_ecdf [_egaf :],_cgb .U [32:40]);_debe ,_fgg :=_dff .alg2b (_cgb .R ,_ecdf ,_cgdeb ,nil );if _fgg !=nil {return nil ,_fgg ;};_debe =_debe [:32];if !_b .Equal (_debe ,_cgb .U [:32]){return nil ,nil ;};return _debe ,nil ;};type stdHandlerR4 struct{Length int ;
ID0 string ;};

// Allowed checks if a set of permissions can be granted.
func (_gbf Permissions )Allowed (p2 Permissions )bool {return _gbf &p2 ==p2 };const (PermOwner =Permissions (_efd .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);PermAnnotate =Permissions (1<<5);
PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e"););func (_db errInvalidField )Error ()string {return _bd .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_db .Func ,_db .Field ,_db .Exp ,_db .Got );
};func (_add stdHandlerR4 )alg3 (R int ,_fbd ,_abe []byte )([]byte ,error ){var _bac []byte ;if len (_abe )> 0{_bac =_add .alg3Key (R ,_abe );}else {_bac =_add .alg3Key (R ,_fbd );};_ggd ,_deb :=_ef .NewCipher (_bac );if _deb !=nil {return nil ,_fcf .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_aac :=_add .paddedPass (_fbd );_eaf :=make ([]byte ,len (_aac ));_ggd .XORKeyStream (_eaf ,_aac );if R >=3{_adbf :=make ([]byte ,len (_bac ));for _bdg :=0;_bdg < 19;_bdg ++{for _ff :=0;_ff < len (_bac );_ff ++{_adbf [_ff ]=_bac [_ff ]^byte (_bdg +1);
};_bg ,_bab :=_ef .NewCipher (_adbf );if _bab !=nil {return nil ,_fcf .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_bg .XORKeyStream (_eaf ,_eaf );};};return _eaf ,nil ;};func (_ced stdHandlerR4 )alg5 (_ed []byte ,_fbda []byte )([]byte ,error ){_dfa :=_g .New ();
_dfa .Write ([]byte (_ce ));_dfa .Write ([]byte (_ced .ID0 ));_aec :=_dfa .Sum (nil );_a .Log .Trace ("\u0061\u006c\u0067\u0035");_a .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_ed );_a .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_ced .ID0 );
if len (_aec )!=16{return nil ,_fcf .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");};_efg ,_ebd :=_ef .NewCipher (_ed );if _ebd !=nil {return nil ,_fcf .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_gca :=make ([]byte ,16);_efg .XORKeyStream (_gca ,_aec );_bba :=make ([]byte ,len (_ed ));for _edc :=0;_edc < 19;_edc ++{for _egb :=0;_egb < len (_ed );_egb ++{_bba [_egb ]=_ed [_egb ]^byte (_edc +1);};_efg ,_ebd =_ef .NewCipher (_bba );if _ebd !=nil {return nil ,_fcf .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_efg .XORKeyStream (_gca ,_gca );_a .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_edc ,_bba );_a .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_edc ,_gca );
};_cd :=make ([]byte ,32);for _fe :=0;_fe < 16;_fe ++{_cd [_fe ]=_gca [_fe ];};_ ,_ebd =_fb .Read (_cd [16:32]);if _ebd !=nil {return nil ,_fcf .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _cd ,nil ;};func (_dg stdHandlerR4 )alg4 (_ded []byte ,_ge []byte )([]byte ,error ){_bc ,_da :=_ef .NewCipher (_ded );if _da !=nil {return nil ,_fcf .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_dfd :=[]byte (_ce );
_fbe :=make ([]byte ,len (_dfd ));_bc .XORKeyStream (_fbe ,_dfd );return _fbe ,nil ;};func (_ad *ecbDecrypter )BlockSize ()int {return _ad ._ee };type stdHandlerR6 struct{};func (_aa stdHandlerR4 )alg3Key (R int ,_ggf []byte )[]byte {_ab :=_g .New ();_ccf :=_aa .paddedPass (_ggf );
_ab .Write (_ccf );if R >=3{for _gga :=0;_gga < 50;_gga ++{_ba :=_ab .Sum (nil );_ab =_g .New ();_ab .Write (_ba );};};_fd :=_ab .Sum (nil );if R ==2{_fd =_fd [0:5];}else {_fd =_fd [0:_aa .Length /8];};return _fd ;};func (stdHandlerR4 )paddedPass (_gg []byte )[]byte {_fcd :=make ([]byte ,32);
_fceb :=copy (_fcd ,_gg );for ;_fceb < 32;_fceb ++{_fcd [_fceb ]=_ce [_fceb -len (_gg )];};return _fcd ;};func _ec (_cf _f .Block )*ecb {return &ecb {_bb :_cf ,_ee :_cf .BlockSize ()}};func (_cgd stdHandlerR6 )alg2a (_cad *StdEncryptDict ,_ebb []byte )([]byte ,Permissions ,error ){if _fcea :=_dcd ("\u0061\u006c\u00672\u0061","\u004f",48,_cad .O );
_fcea !=nil {return nil ,0,_fcea ;};if _ggfe :=_dcd ("\u0061\u006c\u00672\u0061","\u0055",48,_cad .U );_ggfe !=nil {return nil ,0,_ggfe ;};if len (_ebb )> 127{_ebb =_ebb [:127];};_ebgb ,_gab :=_cgd .alg12 (_cad ,_ebb );if _gab !=nil {return nil ,0,_gab ;
};var (_dae []byte ;_bdgd []byte ;_eec []byte ;);var _gfg Permissions ;if len (_ebgb )!=0{_gfg =PermOwner ;_cfb :=make ([]byte ,len (_ebb )+8+48);_bbg :=copy (_cfb ,_ebb );_bbg +=copy (_cfb [_bbg :],_cad .O [40:48]);copy (_cfb [_bbg :],_cad .U [0:48]);
_dae =_cfb ;_bdgd =_cad .OE ;_eec =_cad .U [0:48];}else {_ebgb ,_gab =_cgd .alg11 (_cad ,_ebb );if _gab ==nil &&len (_ebgb )==0{_ebgb ,_gab =_cgd .alg11 (_cad ,[]byte (""));};if _gab !=nil {return nil ,0,_gab ;}else if len (_ebgb )==0{return nil ,0,nil ;
};_gfg =_cad .P ;_faa :=make ([]byte ,len (_ebb )+8);_adc :=copy (_faa ,_ebb );copy (_faa [_adc :],_cad .U [40:48]);_dae =_faa ;_bdgd =_cad .UE ;_eec =nil ;};if _abc :=_dcd ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_bdgd );_abc !=nil {return nil ,0,_abc ;
};_bdgd =_bdgd [:32];_bcf ,_gab :=_cgd .alg2b (_cad .R ,_dae ,_ebb ,_eec );if _gab !=nil {return nil ,0,_gab ;};_aecg ,_gab :=_cc .NewCipher (_bcf [:32]);if _gab !=nil {return nil ,0,_gab ;};_efb :=make ([]byte ,_cc .BlockSize );_bacg :=_f .NewCBCDecrypter (_aecg ,_efb );
_gcae :=make ([]byte ,32);_bacg .CryptBlocks (_gcae ,_bdgd );if _cad .R ==5{return _gcae ,_gfg ,nil ;};_gab =_cgd .alg13 (_cad ,_gcae );if _gab !=nil {return nil ,0,_gab ;};return _gcae ,_gfg ,nil ;};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;func (_egd stdHandlerR4 )alg2 (_df *StdEncryptDict ,_eeb []byte )[]byte {_a .Log .Trace ("\u0061\u006c\u0067\u0032");_bbc :=_egd .paddedPass (_eeb );_dec :=_g .New ();_dec .Write (_bbc );_dec .Write (_df .O );var _cg [4]byte ;_ccc .LittleEndian .PutUint32 (_cg [:],uint32 (_df .P ));
_dec .Write (_cg [:]);_a .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_cg );_dec .Write ([]byte (_egd .ID0 ));_a .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_df .R ,_df .EncryptMetadata );
if (_df .R >=4)&&!_df .EncryptMetadata {_dec .Write ([]byte {0xff,0xff,0xff,0xff});};_gc :=_dec .Sum (nil );if _df .R >=3{_dec =_g .New ();for _ea :=0;_ea < 50;_ea ++{_dec .Reset ();_dec .Write (_gc [0:_egd .Length /8]);_gc =_dec .Sum (nil );};};if _df .R >=3{return _gc [0:_egd .Length /8];
};return _gc [0:5];};func (_gfb stdHandlerR6 )alg13 (_dfe *StdEncryptDict ,_egdd []byte )error {if _begc :=_dcd ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_egdd );_begc !=nil {return _begc ;};if _fca :=_dcd ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_dfe .Perms );
_fca !=nil {return _fca ;};_abb :=make ([]byte ,16);copy (_abb ,_dfe .Perms [:16]);_dggd ,_acea :=_cc .NewCipher (_egdd [:32]);if _acea !=nil {return _acea ;};_bgd :=_ecd (_dggd );_bgd .CryptBlocks (_abb ,_abb );if !_b .Equal (_abb [9:12],[]byte ("\u0061\u0064\u0062")){return _fcf .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_aff :=Permissions (_ccc .LittleEndian .Uint32 (_abb [0:4]));if _aff !=_dfe .P {return _fcf .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _acc bool ;if _abb [8]=='T'{_acc =true ;}else if _abb [8]=='F'{_acc =false ;}else {return _fcf .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _acc !=_dfe .EncryptMetadata {return _fcf .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};func (_bf *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_bf ._ee !=0{_a .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_a .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_bf ._bb .Decrypt (dst ,src [:_bf ._ee ]);src =src [_bf ._ee :];dst =dst [_bf ._ee :];};};func (_gf *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_gf ._ee !=0{_a .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_a .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_gf ._bb .Encrypt (dst ,src [:_gf ._ee ]);src =src [_gf ._ee :];dst =dst [_gf ._ee :];};};func (_dcc *ecbEncrypter )BlockSize ()int {return _dcc ._ee };

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_bfc stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_cbg :=make ([]byte ,32);if _ ,_cgbe :=_c .ReadFull (_fb .Reader ,_cbg );_cgbe !=nil {return nil ,_cgbe ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;d .Perms =nil ;
if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _fbfc :=_bfc .alg8 (d ,_cbg ,upass );_fbfc !=nil {return nil ,_fbfc ;};if _cde :=_bfc .alg9 (d ,_cbg ,opass );_cde !=nil {return nil ,_cde ;};if d .R ==5{return _cbg ,nil ;
};if _daee :=_bfc .alg10 (d ,_cbg );_daee !=nil {return nil ,_daee ;};return _cbg ,nil ;};func _cb (_bbb []byte )(_f .Block ,error ){_fbb ,_af :=_cc .NewCipher (_bbb );if _af !=nil {_a .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_af );
return nil ,_af ;};return _fbb ,nil ;};func (_aegg stdHandlerR6 )alg10 (_cgde *StdEncryptDict ,_dcfb []byte )error {if _fed :=_dcd ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_dcfb );_fed !=nil {return _fed ;};_gde :=uint64 (uint32 (_cgde .P ))|(_efd .MaxUint32 <<32);
Perms :=make ([]byte ,16);_ccc .LittleEndian .PutUint64 (Perms [:8],_gde );if _cgde .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");if _ ,_ecgf :=_c .ReadFull (_fb .Reader ,Perms [12:16]);_ecgf !=nil {return _ecgf ;
};_faag ,_bda :=_cb (_dcfb [:32]);if _bda !=nil {return _bda ;};_fec :=_dc (_faag );_fec .CryptBlocks (Perms ,Perms );_cgde .Perms =Perms [:16];return nil ;};

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_cccg stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_ffg :=_cccg .alg3 (d .R ,upass ,opass );if _ffg !=nil {_a .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ffg );
return nil ,_ffg ;};d .O =O ;_a .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_gcf :=_cccg .alg2 (d ,upass );U ,_ffg :=_cccg .alg5 (_gcf ,upass );if _ffg !=nil {_a .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ffg );
return nil ,_ffg ;};d .U =U ;_a .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _gcf ,nil ;};

// Authenticate implements StdHandler interface.
func (_aca stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_a .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_ga ,_ccg :=_aca .alg7 (d ,pass );if _ccg !=nil {return nil ,0,_ccg ;};if _ga !=nil {_a .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _ga ,PermOwner ,nil ;
};_a .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_ga ,_ccg =_aca .alg6 (d ,pass );if _ccg !=nil {return nil ,0,_ccg ;};
if _ga !=nil {_a .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _ga ,d .P ,nil ;};return nil ,0,nil ;};func (_bbf stdHandlerR6 )alg2b (R int ,_fgd ,_cdg ,_beg []byte )([]byte ,error ){if R ==5{return _gae (_fgd );
};return _aef (_fgd ,_cdg ,_beg );};func _ecd (_ecb _f .Block )_f .BlockMode {return (*ecbDecrypter )(_ec (_ecb ))};func _gae (_eda []byte )([]byte ,error ){_ecbe :=_d .New ();_ecbe .Write (_eda );return _ecbe .Sum (nil ),nil ;};func (_fba stdHandlerR6 )alg12 (_beeb *StdEncryptDict ,_ccd []byte )([]byte ,error ){if _cfc :=_dcd ("\u0061\u006c\u00671\u0032","\u0055",48,_beeb .U );
_cfc !=nil {return nil ,_cfc ;};if _bga :=_dcd ("\u0061\u006c\u00671\u0032","\u004f",48,_beeb .O );_bga !=nil {return nil ,_bga ;};_gaa :=make ([]byte ,len (_ccd )+8+48);_dcff :=copy (_gaa ,_ccd );_dcff +=copy (_gaa [_dcff :],_beeb .O [32:40]);_dcff +=copy (_gaa [_dcff :],_beeb .U [0:48]);
_eee ,_ege :=_fba .alg2b (_beeb .R ,_gaa ,_ccd ,_beeb .U [0:48]);if _ege !=nil {return nil ,_ege ;};_eee =_eee [:32];if !_b .Equal (_eee ,_beeb .O [:32]){return nil ,nil ;};return _eee ,nil ;};func _aba (_fg []byte ,_ecg int ){_ega :=_ecg ;for _ega < len (_fg ){copy (_fg [_ega :],_fg [:_ega ]);
_ega *=2;};};const _ce ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";func (_cge stdHandlerR4 )alg7 (_gce *StdEncryptDict ,_bcd []byte )([]byte ,error ){_bfa :=_cge .alg3Key (_gce .R ,_bcd );
_dfdf :=make ([]byte ,len (_gce .O ));if _gce .R ==2{_fbf ,_fbc :=_ef .NewCipher (_bfa );if _fbc !=nil {return nil ,_fcf .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_fbf .XORKeyStream (_dfdf ,_gce .O );}else if _gce .R >=3{_fdf :=append ([]byte {},_gce .O ...);
for _dcg :=0;_dcg < 20;_dcg ++{_ag :=append ([]byte {},_bfa ...);for _dcf :=0;_dcf < len (_bfa );_dcf ++{_ag [_dcf ]^=byte (19-_dcg );};_ead ,_bag :=_ef .NewCipher (_ag );if _bag !=nil {return nil ,_fcf .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_ead .XORKeyStream (_dfdf ,_fdf );_fdf =append ([]byte {},_dfdf ...);};}else {return nil ,_fcf .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};_ca ,_fde :=_cge .alg6 (_gce ,_dfdf );if _fde !=nil {return nil ,nil ;};return _ca ,nil ;};type errInvalidField struct{Func string ;
Field string ;Exp int ;Got int ;};func _aef (_cce ,_bfb ,_gba []byte )([]byte ,error ){var (_dece ,_dd ,_dce _eb .Hash ;);_dece =_d .New ();_deda :=make ([]byte ,64);_cccga :=_dece ;_cccga .Write (_cce );K :=_cccga .Sum (_deda [:0]);_fad :=make ([]byte ,64*(127+64+48));
_cfe :=func (_aacc int )([]byte ,error ){_fea :=len (_bfb )+len (K )+len (_gba );_fbdb :=_fad [:_fea ];_afg :=copy (_fbdb ,_bfb );_afg +=copy (_fbdb [_afg :],K [:]);_afg +=copy (_fbdb [_afg :],_gba );if _afg !=_fea {_a .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_fcf .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_fad [:_fea *64];_aba (K1 ,_fea );_cef ,_abg :=_cb (K [0:16]);if _abg !=nil {return nil ,_abg ;};_dge :=_f .NewCBCEncrypter (_cef ,K [16:32]);_dge .CryptBlocks (K1 ,K1 );
E :=K1 ;_gef :=0;for _abeb :=0;_abeb < 16;_abeb ++{_gef +=int (E [_abeb ]%3);};var _aga _eb .Hash ;switch _gef %3{case 0:_aga =_dece ;case 1:if _dd ==nil {_dd =_fc .New384 ();};_aga =_dd ;case 2:if _dce ==nil {_dce =_fc .New ();};_aga =_dce ;};_aga .Reset ();
_aga .Write (E );K =_aga .Sum (_deda [:0]);return E ,nil ;};for _bce :=0;;{E ,_feg :=_cfe (_bce );if _feg !=nil {return nil ,_feg ;};_eab :=E [len (E )-1];_bce ++;if _bce >=64&&_eab <=uint8 (_bce -32){break ;};};return K [:32],nil ;};func (_daa stdHandlerR6 )alg9 (_ffge *StdEncryptDict ,_geeg []byte ,_fdff []byte )error {if _acf :=_dcd ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_geeg );
_acf !=nil {return _acf ;};if _cee :=_dcd ("\u0061\u006c\u0067\u0039","\u0055",48,_ffge .U );_cee !=nil {return _cee ;};var _addg [16]byte ;if _ ,_eeg :=_c .ReadFull (_fb .Reader ,_addg [:]);_eeg !=nil {return _eeg ;};_gcd :=_addg [0:8];_aea :=_addg [8:16];
_cgg :=_ffge .U [:48];_aeac :=make ([]byte ,len (_fdff )+len (_gcd )+len (_cgg ));_bfd :=copy (_aeac ,_fdff );_bfd +=copy (_aeac [_bfd :],_gcd );_bfd +=copy (_aeac [_bfd :],_cgg );_abebb ,_aeg :=_daa .alg2b (_ffge .R ,_aeac ,_fdff ,_cgg );if _aeg !=nil {return _aeg ;
};O :=make ([]byte ,len (_abebb )+len (_gcd )+len (_aea ));_bfd =copy (O ,_abebb [:32]);_bfd +=copy (O [_bfd :],_gcd );_bfd +=copy (O [_bfd :],_aea );_ffge .O =O ;_bfd =len (_fdff );_bfd +=copy (_aeac [_bfd :],_aea );_abebb ,_aeg =_daa .alg2b (_ffge .R ,_aeac ,_fdff ,_cgg );
if _aeg !=nil {return _aeg ;};_aaccf ,_aeg :=_cb (_abebb [:32]);if _aeg !=nil {return _aeg ;};_fgc :=make ([]byte ,_cc .BlockSize );_cag :=_f .NewCBCEncrypter (_aaccf ,_fgc );OE :=make ([]byte ,32);_cag .CryptBlocks (OE ,_geeg [:32]);_ffge .OE =OE ;return nil ;
};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_fce *StdEncryptDict ,_de ,_eg []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_efc *StdEncryptDict ,_gb []byte )([]byte ,Permissions ,error );};