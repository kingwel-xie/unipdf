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

package transform ;import (_b "fmt";_f "github.com/unidoc/unipdf/v3/common";_c "math";);const _cfe =1e-6;func (_gc *Matrix )Concat (b Matrix ){*_gc =Matrix {b [0]*_gc [0]+b [1]*_gc [3],b [0]*_gc [1]+b [1]*_gc [4],0,b [3]*_gc [0]+b [4]*_gc [3],b [3]*_gc [1]+b [4]*_gc [4],0,b [6]*_gc [0]+b [7]*_gc [3]+_gc [6],b [6]*_gc [1]+b [7]*_gc [4]+_gc [7],1};
_gc .clampRange ();};const _bbc =1e9;func (_ffe *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_gf :=NewMatrix (a ,b ,c ,d ,tx ,ty );_ffe .transformByMatrix (_gf );};func (_deg Point )Displace (delta Point )Point {return Point {_deg .X +delta .X ,_deg .Y +delta .Y }};
func (_cc Matrix )Scale (xScale ,yScale float64 )Matrix {return _cc .Mult (ScaleMatrix (xScale ,yScale ))};func (_gb Matrix )Transform (x ,y float64 )(float64 ,float64 ){_ff :=x *_gb [0]+y *_gb [3]+_gb [6];_ba :=x *_gb [1]+y *_gb [4]+_gb [7];return _ff ,_ba ;
};func (_fcf Matrix )Translation ()(float64 ,float64 ){return _fcf [6],_fcf [7]};func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );};func (_cbc Point )String ()string {return _b .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_cbc .X ,_cbc .Y );
};func RotationMatrix (angle float64 )Matrix {_fc :=_c .Cos (angle );_g :=_c .Sin (angle );return NewMatrix (_fc ,_g ,-_g ,_fc ,0,0);};func (_gbg Matrix )Singular ()bool {return _c .Abs (_gbg [0]*_gbg [4]-_gbg [1]*_gbg [3])< _cea };func (_bg Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_bg .X +t *b .X ,Y :(1-t )*_bg .Y +t *b .Y };
};func (_dg Matrix )String ()string {_ce ,_ac ,_cf ,_cfa ,_fg ,_cgd :=_dg [0],_dg [1],_dg [3],_dg [4],_dg [6],_dg [7];return _b .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_ce ,_ac ,_cf ,_cfa ,_fg ,_cgd );
};func (_de Matrix )Translate (tx ,ty float64 )Matrix {return _de .Mult (TranslationMatrix (tx ,ty ))};func (_ae *Matrix )Clone ()Matrix {return NewMatrix (_ae [0],_ae [1],_ae [3],_ae [4],_ae [6],_ae [7])};func (_gca *Point )Set (x ,y float64 ){_gca .X ,_gca .Y =x ,y };
type Point struct{X float64 ;Y float64 ;};const _cea =1e-10;func (_eg Matrix )Unrealistic ()bool {_aa ,_bdb ,_df ,_ca :=_c .Abs (_eg [0]),_c .Abs (_eg [1]),_c .Abs (_eg [3]),_c .Abs (_eg [4]);_cbb :=_aa > _cfe &&_ca > _cfe ;_bea :=_bdb > _cfe &&_df > _cfe ;
return !(_cbb ||_bea );};func (_d Matrix )Identity ()bool {return _d [0]==1&&_d [1]==0&&_d [2]==0&&_d [3]==0&&_d [4]==1&&_d [5]==0&&_d [6]==0&&_d [7]==0&&_d [8]==1;};func (_ed *Point )transformByMatrix (_cgc Matrix ){_ed .X ,_ed .Y =_cgc .Transform (_ed .X ,_ed .Y )};
func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};func (_ad Matrix )Rotate (theta float64 )Matrix {return _ad .Mult (RotationMatrix (theta ))};func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};
func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};func (_ebd Matrix )Mult (b Matrix )Matrix {_ebd .Concat (b );return _ebd };func (_be *Matrix )clampRange (){for _gbc ,_bd :=range _be {if _bd > _bbc {_f .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_bd ,_bbc );
_be [_gbc ]=_bbc ;}else if _bd < -_bbc {_f .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_bd ,-_bbc );_be [_gbc ]=-_bbc ;};};};func (_cb Matrix )ScalingFactorX ()float64 {return _c .Hypot (_cb [0],_cb [1])};
func (_eag Point )Rotate (theta float64 )Point {_dd :=_c .Hypot (_eag .X ,_eag .Y );_fe :=_c .Atan2 (_eag .Y ,_eag .X );_aba ,_gg :=_c .Sincos (_fe +theta /180.0*_c .Pi );return Point {_dd *_gg ,_dd *_aba };};func (_ccb Matrix )Angle ()float64 {_ab :=_c .Atan2 (-_ccb [1],_ccb [0]);
if _ab < 0.0{_ab +=2*_c .Pi ;};return _ab /_c .Pi *180.0;};func (_ee Matrix )Inverse ()(Matrix ,bool ){_ada ,_db :=_ee [0],_ee [1];_bb ,_gcb :=_ee [3],_ee [4];_cd ,_add :=_ee [6],_ee [7];_eef :=_ada *_gcb -_db *_bb ;if _c .Abs (_eef )< _ga {return Matrix {},false ;
};_fgg ,_bc :=_gcb /_eef ,-_db /_eef ;_ef ,_bf :=-_bb /_eef ,_ada /_eef ;_fa :=-(_fgg *_cd +_ef *_add );_ge :=-(_bc *_cd +_bf *_add );return NewMatrix (_fgg ,_bc ,_ef ,_bf ,_fa ,_ge ),true ;};func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_cg :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};
_cg .clampRange ();return _cg ;};func (_acb Point )Distance (b Point )float64 {return _c .Hypot (_acb .X -b .X ,_acb .Y -b .Y )};type Matrix [9]float64 ;func (_eb Matrix )Round (precision float64 )Matrix {for _a :=range _eb {_eb [_a ]=_c .Round (_eb [_a ]/precision )*precision ;
};return _eb ;};func (_fd Matrix )ScalingFactorY ()float64 {return _c .Hypot (_fd [3],_fd [4])};const _ga =1.0e-6;func (_ea *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_ea [0],_ea [1]=a ,b ;_ea [3],_ea [4]=c ,d ;_ea [6],_ea [7]=tx ,ty ;_ea .clampRange ();
};func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};func (_ebg *Matrix )Shear (x ,y float64 ){_ebg .Concat (ShearMatrix (x ,y ))};