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

package pdfutil ;import (_bb "github.com/unidoc/unipdf/v3/common";_cd "github.com/unidoc/unipdf/v3/contentstream";_c "github.com/unidoc/unipdf/v3/contentstream/draw";_e "github.com/unidoc/unipdf/v3/core";_g "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_g .PdfPage )error {_cdc ,_cdb :=page .GetMediaBox ();if _cdb !=nil {return _cdb ;};_ba ,_cdb :=page .GetRotate ();if _cdb !=nil {_bb .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_cdb .Error ());
};_bg :=_ba %360!=0&&_ba %90==0;_cdc .Normalize ();_cf ,_gg ,_ec ,_be :=_cdc .Llx ,_cdc .Lly ,_cdc .Width (),_cdc .Height ();_bab :=_cf !=0||_gg !=0;if !_bg &&!_bab {return nil ;};_a :=func (_ggg ,_d ,_eg float64 )_c .BoundingBox {return _c .Path {Points :[]_c .Point {_c .NewPoint (0,0).Rotate (_eg ),_c .NewPoint (_ggg ,0).Rotate (_eg ),_c .NewPoint (0,_d ).Rotate (_eg ),_c .NewPoint (_ggg ,_d ).Rotate (_eg )}}.GetBoundingBox ();
};_ga :=_cd .NewContentCreator ();var _af float64 ;if _bg {_af =-float64 (_ba );_f :=_a (_ec ,_be ,_af );_ga .Translate ((_f .Width -_ec )/2+_ec /2,(_f .Height -_be )/2+_be /2);_ga .RotateDeg (_af );_ga .Translate (-_ec /2,-_be /2);_ec ,_be =_f .Width ,_f .Height ;
};if _bab {_ga .Translate (-_cf ,-_gg );};_ea :=_ga .Operations ();_db ,_cdb :=_e .MakeStream (_ea .Bytes (),_e .NewFlateEncoder ());if _cdb !=nil {return _cdb ;};_eag :=_e .MakeArray (_db );_eag .Append (page .GetContentStreamObjs ()...);*_cdc =_g .PdfRectangle {Urx :_ec ,Ury :_be };
if _afd :=page .CropBox ;_afd !=nil {_afd .Normalize ();_bf ,_fc ,_fg ,_gaf :=_afd .Llx -_cf ,_afd .Lly -_gg ,_afd .Width (),_afd .Height ();if _bg {_fcg :=_a (_fg ,_gaf ,_af );_fg ,_gaf =_fcg .Width ,_fcg .Height ;};*_afd =_g .PdfRectangle {Llx :_bf ,Lly :_fc ,Urx :_bf +_fg ,Ury :_fc +_gaf };
};_bb .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_af ,_ea ,_cdc );page .Contents =_eag ;page .Rotate =nil ;
return nil ;};