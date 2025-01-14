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

package uuid ;import (_c "crypto/rand";_a "encoding/hex";_d "io";);var _b =_c .Reader ;func _ee (_cb []byte ,_df UUID ){_a .Encode (_cb ,_df [:4]);_cb [8]='-';_a .Encode (_cb [9:13],_df [4:6]);_cb [13]='-';_a .Encode (_cb [14:18],_df [6:8]);_cb [18]='-';
_a .Encode (_cb [19:23],_df [8:10]);_cb [23]='-';_a .Encode (_cb [24:],_df [10:]);};func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_bf :=_d .ReadFull (_b ,uuid [:]);if _bf !=nil {return _eba ,_bf ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;
return uuid ,nil ;};var _eba UUID ;func (_ag UUID )String ()string {var _be [36]byte ;_ee (_be [:],_ag );return string (_be [:])};var Nil =_eba ;type UUID [16]byte ;func MustUUID ()UUID {uuid ,_eb :=NewUUID ();if _eb !=nil {panic (_eb );};return uuid ;
};