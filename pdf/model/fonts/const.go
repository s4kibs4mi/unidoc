/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */
/*
 * The embedded character metrics specified in this file are distributed under the terms listed in
 * ./afms/MustRead.html.
 */

package fonts

type FontWeight int

const (
	FontWeightMedium FontWeight = iota
	FontWeightBold
	FontWeightRoman
)

type FontMetrics struct {
	FontName     string
	FontFamily   string
	FirstChar    int
	LastChar     int
	Descent      float64
	FontBbox     [4]float64
	FontWeight   FontWeight
	CapHeight    float64
	Flags        uint
	XHeight      float64
	ItalicAngle  float64
	Ascent       float64
	Widths       []float64
	MissingWidth uint
	Leading      uint
	Vscale       float64
	Hscale       float64
	FontMatrix   [6]float64
}

var (
	Standard14Fonts = map[string]Font{
		"Courier":               NewFontCourier(),
		"Courier-Bold":          NewFontCourierBold(),
		"Courier-BoldOblique":   NewFontCourierBoldOblique(),
		"Courier-Oblique":       NewFontCourierOblique(),
		"Helvetica":             NewFontHelvetica(),
		"Helvetica-Bold":        NewFontHelveticaBold(),
		"Helvetica-BoldOblique": NewFontHelveticaBoldOblique(),
		"Helvetica-Oblique":     NewFontHelveticaOblique(),
		"Times-Roman":           NewFontTimesRoman(),
		"Times-Bold":            NewFontTimesBold(),
		"Times-BoldItalic":      NewFontTimesBoldItalic(),
		"Times-Italic":          NewFontTimesItalic(),
		"Symbol":                NewFontSymbol(),
		"ZapfDingbats":          NewFontZapfDingbats(),
	}

	Standard14FontMetrics = map[string]FontMetrics{
		"Courier":               fmCourier,
		"Courier-Bold":          fmCourierBold,
		"Courier-BoldOblique":   fmCourierBoldOblique,
		"Courier-Oblique":       fmCourierOblique,
		"Helvetica":             fmHelvetica,
		"Helvetica-Bold":        fmHelveticaBold,
		"Helvetica-BoldOblique": fmHelveticaBoldOblique,
		"Helvetica-Oblique":     fmHelveticaOblique,
		"Times-Roman":           fmTimesRoman,
		"Times-Bold":            fmTimesBold,
		"Times-BoldItalic":      fmTimesBoldItalic,
		"Times-Italic":          fmTimesItalic,
		"Symbol":                fmSymbol,
		"ZapfDingbats":          fmZapfDingbats,
	}

	fmCourier = FontMetrics{
		FontName:     "Courier",
		FontFamily:   "Courier",
		FirstChar:    0,
		LastChar:     0,
		Descent:      -194.0,
		FontBbox:     [4]float64{-6.0, -249.0, 639.0, 803.0},
		FontWeight:   FontWeightMedium,
		CapHeight:    572.0,
		Flags:        64,
		XHeight:      434.0,
		ItalicAngle:  0.0,
		Ascent:       627.0,
		Widths:       []float64{},
		MissingWidth: 600,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmCourierBold = FontMetrics{
		FontName:     "Courier-Bold",
		FontFamily:   "Courier",
		FirstChar:    0,
		LastChar:     0,
		Descent:      -194.0,
		FontBbox:     [4]float64{-88.0, -249.0, 697.0, 811.0},
		FontWeight:   FontWeightBold,
		CapHeight:    572.0,
		Flags:        64,
		XHeight:      434.0,
		ItalicAngle:  0.0,
		Ascent:       627.0,
		Widths:       []float64{},
		MissingWidth: 600,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmCourierBoldOblique = FontMetrics{
		FontName:     "Courier-BoldOblique",
		FontFamily:   "Courier",
		FirstChar:    0,
		LastChar:     0,
		Descent:      -194.0,
		FontBbox:     [4]float64{-49.0, -249.0, 758.0, 811.0},
		FontWeight:   FontWeightBold,
		CapHeight:    572.0,
		Flags:        64,
		XHeight:      434.0,
		ItalicAngle:  -11.0,
		Ascent:       627.0,
		Widths:       []float64{},
		MissingWidth: 600,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmCourierOblique = FontMetrics{
		FontName:     "Courier-Oblique",
		FontFamily:   "Courier",
		FirstChar:    0,
		LastChar:     0,
		Descent:      -194.0,
		FontBbox:     [4]float64{-49.0, -249.0, 749, 803},
		FontWeight:   FontWeightMedium,
		CapHeight:    572.0,
		Flags:        64,
		XHeight:      434.0,
		ItalicAngle:  -11.0,
		Ascent:       627.0,
		Widths:       []float64{},
		MissingWidth: 600,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmHelvetica = FontMetrics{
		FontName:    "Helvetica",
		FontFamily:  "Helvetica",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -207.0,
		FontBbox:    [4]float64{-166.0, -225.0, 1000.0, 931.0},
		FontWeight:  FontWeightMedium,
		CapHeight:   718.0,
		Flags:       0,
		XHeight:     523.0,
		ItalicAngle: 0.0,
		Ascent:      718.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			278, 278, 355, 556, 556, 889, 667, 222, 333, 333, 389, 584, 278, 333, 278, 278,
			556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 278, 278, 584, 584, 584, 556,
			1015, 667, 667, 722, 722, 667, 611, 778, 722, 278, 500, 667, 556, 833, 722, 778,
			667, 778, 722, 667, 611, 722, 667, 944, 667, 667, 611, 278, 278, 278, 469, 556,
			222, 556, 556, 500, 556, 556, 278, 556, 556, 222, 222, 500, 222, 833, 556, 556,
			556, 556, 333, 500, 278, 556, 500, 722, 500, 500, 500, 334, 260, 334, 584, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 333, 556, 556, 167, 556, 556, 556, 556, 191, 333, 556, 333, 333,
			500, 500, 0, 556, 556, 556, 278, 0, 537, 350, 222, 333, 333, 556, 1000, 1000, 0,
			611, 0, 333, 333, 333, 333, 333, 333, 333, 333, 0, 333, 333, 0, 333, 333, 333,
			1000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1000, 0, 370, 0, 0, 0, 0,
			556, 778, 1000, 365, 0, 0, 0, 0, 0, 889, 0, 0, 0, 278, 0, 0, 222, 611, 944, 611,
			0, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmHelveticaBold = FontMetrics{
		FontName:    "Helvetica-Bold",
		FontFamily:  "Helvetica",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -207.0,
		FontBbox:    [4]float64{-170.0, -228.0, 1003.0, 962.0},
		FontWeight:  FontWeightBold,
		CapHeight:   718.0,
		Flags:       0,
		XHeight:     532.0,
		ItalicAngle: 0.0,
		Ascent:      718.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			278, 333, 474, 556, 556, 889, 722, 278, 333, 333, 389, 584, 278, 333,
			278, 278, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 333, 333,
			584, 584, 584, 611, 975, 722, 722, 722, 722, 667, 611, 778, 722, 278,
			556, 722, 611, 833, 722, 778, 667, 778, 722, 667, 611, 722, 667, 944,
			667, 667, 611, 333, 278, 333, 584, 556, 278, 556, 611, 556, 611, 556,
			333, 611, 611, 278, 278, 556, 278, 889, 611, 611, 611, 611, 389, 556,
			333, 611, 556, 778, 556, 556, 500, 389, 280, 389, 584, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 333, 556, 556, 167, 556, 556, 556, 556, 238, 500,
			556, 333, 333, 611, 611, 0, 556, 556, 556, 278, 0, 556, 350, 278, 500,
			500, 556, 1000, 1000, 0, 611, 0, 333, 333, 333, 333, 333, 333, 333,
			333, 0, 333, 333, 0, 333, 333, 333, 1000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1000, 0, 370, 0, 0, 0, 0, 611, 778, 1000, 365, 0, 0,
			0, 0, 0, 889, 0, 0, 0, 278, 0, 0, 278, 611, 944, 611, 0, 0, 0, 0,
		},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmHelveticaBoldOblique = FontMetrics{
		FontName:    "Helvetica-BoldOblique",
		FontFamily:  "Helvetica",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -207.0,
		FontBbox:    [4]float64{-175.0, -228.0, 1114.0, 962.0},
		FontWeight:  FontWeightBold,
		CapHeight:   718.0,
		Flags:       0,
		XHeight:     532.0,
		ItalicAngle: -12.0,
		Ascent:      718.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			278, 333, 474, 556, 556, 889, 722, 278, 333, 333, 389, 584, 278, 333,
			278, 278, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 333, 333,
			584, 584, 584, 611, 975, 722, 722, 722, 722, 667, 611, 778, 722, 278,
			556, 722, 611, 833, 722, 778, 667, 778, 722, 667, 611, 722, 667, 944,
			667, 667, 611, 333, 278, 333, 584, 556, 278, 556, 611, 556, 611, 556,
			333, 611, 611, 278, 278, 556, 278, 889, 611, 611, 611, 611, 389, 556,
			333, 611, 556, 778, 556, 556, 500, 389, 280, 389, 584, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 333, 556, 556, 167, 556, 556, 556, 556, 238, 500,
			556, 333, 333, 611, 611, 0, 556, 556, 556, 278, 0, 556, 350, 278, 500,
			500, 556, 1000, 1000, 0, 611, 0, 333, 333, 333, 333, 333, 333, 333, 333,
			0, 333, 333, 0, 333, 333, 333, 1000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 1000, 0, 370, 0, 0, 0, 0, 611, 778, 1000, 365, 0, 0, 0, 0,
			0, 889, 0, 0, 0, 278, 0, 0, 278, 611, 944, 611, 0, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmHelveticaOblique = FontMetrics{
		FontName:    "Helvetica-Oblique",
		FontFamily:  "Helvetica",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -207.0,
		FontBbox:    [4]float64{-171.0, -225.0, 1116.0, 931.0},
		FontWeight:  FontWeightMedium,
		CapHeight:   718.0,
		Flags:       0,
		XHeight:     523.0,
		ItalicAngle: -12.0,
		Ascent:      718.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			278, 278, 355, 556, 556, 889, 667, 222, 333, 333, 389, 584, 278, 333, 278,
			278, 556, 556, 556, 556, 556, 556, 556, 556, 556, 556, 278, 278, 584, 584,
			584, 556, 1015, 667, 667, 722, 722, 667, 611, 778, 722, 278, 500, 667, 556,
			833, 722, 778, 667, 778, 722, 667, 611, 722, 667, 944, 667, 667, 611, 278,
			278, 278, 469, 556, 222, 556, 556, 500, 556, 556, 278, 556, 556, 222, 222,
			500, 222, 833, 556, 556, 556, 556, 333, 500, 278, 556, 500, 722, 500, 500,
			500, 334, 260, 334, 584, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 333, 556, 556, 167, 556,
			556, 556, 556, 191, 333, 556, 333, 333, 500, 500, 0, 556, 556, 556, 278, 0,
			537, 350, 222, 333, 333, 556, 1000, 1000, 0, 611, 0, 333, 333, 333, 333, 333,
			333, 333, 333, 0, 333, 333, 0, 333, 333, 333, 1000, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 1000, 0, 370, 0, 0, 0, 0, 556, 778, 1000, 365, 0, 0,
			0, 0, 0, 889, 0, 0, 0, 278, 0, 0, 222, 611, 944, 611, 0, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmTimesRoman = FontMetrics{
		FontName:    "Times-Roman",
		FontFamily:  "Times",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -217.0,
		FontBbox:    [4]float64{-168.0, -218.0, 1000.0, 898.0},
		FontWeight:  FontWeightRoman,
		CapHeight:   662.0,
		Flags:       0,
		XHeight:     450.0,
		ItalicAngle: 0.0,
		Ascent:      683.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			250, 333, 408, 500, 500, 833, 778, 333, 333, 333, 500, 564, 250, 333, 250, 278,
			500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 278, 278, 564, 564, 564, 444,
			921, 722, 667, 667, 722, 611, 556, 722, 722, 333, 389, 722, 611, 889, 722, 722,
			556, 722, 667, 556, 611, 722, 722, 944, 722, 722, 611, 333, 278, 333, 469, 500,
			333, 444, 500, 444, 500, 444, 333, 500, 500, 278, 278, 500, 278, 778, 500, 500,
			500, 500, 333, 389, 278, 500, 500, 722, 500, 500, 444, 480, 200, 480, 541, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 333, 500, 500, 167, 500, 500, 500, 500, 180, 444, 500, 333, 333,
			556, 556, 0, 500, 500, 500, 250, 0, 453, 350, 333, 444, 444, 500, 1000, 1000, 444,
			0, 333, 333, 333, 333, 333, 333, 333, 333, 0, 333, 333, 0, 333, 333, 333, 1000,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 889, 0, 276, 0, 0, 0, 0, 611, 722,
			889, 310, 0, 0, 0, 0, 0, 667, 0, 0, 0, 278, 0, 0, 278, 500, 722, 500, 0, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmTimesBold = FontMetrics{
		FontName:    "Times-Bold",
		FontFamily:  "Times",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -217.0,
		FontBbox:    [4]float64{-168.0, -218.0, 1000.0, 935.0},
		FontWeight:  FontWeightBold,
		CapHeight:   676.0,
		Flags:       0,
		XHeight:     461.0,
		ItalicAngle: 0.0,
		Ascent:      683.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			250, 333, 555, 500, 500, 1000, 833, 333, 333, 333, 500, 570, 250, 333, 250,
			278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 333, 333, 570, 570,
			570, 500, 930, 722, 667, 722, 722, 667, 611, 778, 778, 389, 500, 778, 667,
			944, 722, 778, 611, 778, 722, 556, 667, 722, 722, 1000, 722, 722, 667, 333,
			278, 333, 581, 500, 333, 500, 556, 444, 556, 444, 333, 500, 556, 278, 333,
			556, 278, 833, 556, 500, 556, 556, 444, 389, 333, 556, 500, 722, 500, 500,
			444, 394, 220, 394, 520, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 333, 500, 500, 167, 500,
			500, 500, 500, 278, 500, 500, 333, 333, 556, 556, 0, 500, 500, 500, 250, 0,
			540, 350, 333, 500, 500, 500, 1000, 1000, 0, 500, 0, 333, 333, 333, 333,
			333, 333, 333, 333, 0, 333, 333, 0, 333, 333, 333, 1000, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1000, 0, 300, 0, 0, 0, 0, 667, 778, 1000, 330, 0,
			0, 0, 0, 0, 722, 0, 0, 0, 278, 0, 0, 278, 500, 722, 556, 0, 0, 0, 0,
		},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmTimesBoldItalic = FontMetrics{
		FontName:    "Times-BoldItalic",
		FontFamily:  "Times",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -217.0,
		FontBbox:    [4]float64{-200.0, -218.0, 996.0, 921.0},
		FontWeight:  FontWeightBold,
		CapHeight:   669.0,
		Flags:       0,
		XHeight:     462.0,
		ItalicAngle: -15.0,
		Ascent:      683.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			250, 389, 555, 500, 500, 833, 778, 333, 333, 333, 500, 570, 250, 333,
			250, 278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 333, 333,
			570, 570, 570, 500, 832, 667, 667, 667, 722, 667, 667, 722, 778, 389,
			500, 667, 611, 889, 722, 722, 611, 722, 667, 556, 611, 722, 667, 889,
			667, 611, 611, 333, 278, 333, 570, 500, 333, 500, 500, 444, 500, 444,
			333, 500, 556, 278, 278, 500, 278, 778, 556, 500, 500, 500, 389, 389,
			278, 556, 444, 667, 500, 444, 389, 348, 220, 348, 570, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 389, 500, 500, 167, 500, 500, 500, 500, 278, 500,
			500, 333, 333, 556, 556, 0, 500, 500, 500, 250, 0, 500, 350, 333, 500,
			500, 500, 1000, 1000, 0, 500, 0, 333, 333, 333, 333, 333, 333, 333,
			333, 0, 333, 333, 0, 333, 333, 333, 1000, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 944, 0, 266, 0, 0, 0, 0, 611, 722, 944, 300, 0,
			0, 0, 0, 0, 722, 0, 0, 0, 278, 0, 0, 278, 500, 722, 500, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmTimesItalic = FontMetrics{
		FontName:    "Times-Italic",
		FontFamily:  "Times",
		FirstChar:   0,
		LastChar:    255,
		Descent:     -217.0,
		FontBbox:    [4]float64{-169.0, -217.0, 1010.0, 883.0},
		FontWeight:  FontWeightMedium,
		CapHeight:   653.0,
		Flags:       0,
		XHeight:     441.0,
		ItalicAngle: -15.5,
		Ascent:      683.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			250, 333, 420, 500, 500, 833, 778, 333, 333, 333, 500, 675, 250, 333, 250,
			278, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 333, 333, 675, 675,
			675, 500, 920, 611, 611, 667, 722, 611, 611, 722, 722, 333, 444, 667, 556,
			833, 667, 722, 611, 722, 611, 500, 556, 722, 611, 833, 611, 556, 556, 389,
			278, 389, 422, 500, 333, 500, 500, 444, 500, 444, 278, 500, 500, 278, 278,
			444, 278, 722, 500, 500, 500, 500, 389, 389, 278, 500, 444, 667, 444, 444,
			389, 400, 275, 400, 541, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 389, 500, 500, 167, 500,
			500, 500, 500, 214, 556, 500, 333, 333, 500, 500, 0, 500, 500, 500, 250, 0,
			523, 350, 333, 556, 556, 500, 889, 1000, 0, 500, 0, 333, 333, 333, 333, 333,
			333, 333, 333, 0, 333, 333, 0, 333, 333, 333, 889, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 889, 0, 276, 0, 0, 0, 0, 556, 722, 944, 310, 0, 0,
			0, 0, 0, 667, 0, 0, 0, 278, 0, 0, 278, 500, 667, 500, 0, 0, 0, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmSymbol = FontMetrics{
		FontName:    "Symbol",
		FontFamily:  "Symbol",
		FirstChar:   0,
		LastChar:    255,
		Descent:     0.0,
		FontBbox:    [4]float64{-180.0, -293.0, 1090.0, 1010.0},
		FontWeight:  FontWeightMedium,
		CapHeight:   676.0,
		Flags:       0,
		XHeight:     0.0,
		ItalicAngle: 0.0,
		Ascent:      0.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			250, 333, 713, 500, 549, 833, 778, 439, 333, 333, 500, 549, 250, 549, 250, 278, 500,
			500, 500, 500, 500, 500, 500, 500, 500, 500, 278, 278, 549, 549, 549, 444, 549, 722,
			667, 722, 612, 611, 763, 603, 722, 333, 631, 722, 686, 889, 722, 722, 768, 741, 556,
			592, 611, 690, 439, 768, 645, 795, 611, 333, 863, 333, 658, 500, 500, 631, 549, 549,
			494, 439, 521, 411, 603, 329, 603, 549, 549, 576, 521, 549, 549, 521, 549, 603, 439,
			576, 713, 686, 493, 686, 494, 480, 200, 480, 549, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 750, 620, 247, 549,
			167, 713, 500, 753, 753, 753, 753, 1042, 987, 603, 987, 603, 400, 549, 411, 549, 549,
			713, 494, 460, 549, 549, 549, 549, 1000, 603, 1000, 658, 823, 686, 795, 987, 768, 768,
			823, 768, 768, 713, 713, 713, 713, 713, 713, 713, 768, 713, 790, 790, 890, 823, 549,
			250, 713, 603, 603, 1042, 987, 603, 987, 603, 494, 329, 790, 790, 786, 713, 384,
			384, 384, 384, 384, 384, 494, 494, 494, 494, 329, 274, 0, 686, 686, 686, 384, 384,
			384, 384, 384, 384, 494, 494, 494, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}

	fmZapfDingbats = FontMetrics{
		FontName:    "ZapfDingbats",
		FontFamily:  "ITC",
		FirstChar:   0,
		LastChar:    255,
		Descent:     0.0,
		FontBbox:    [4]float64{-1.0, -143.0, 981.0, 820.0},
		FontWeight:  FontWeightMedium,
		CapHeight:   718.0,
		Flags:       0,
		XHeight:     0.0,
		ItalicAngle: 0.0,
		Ascent:      0.0,
		Widths: []float64{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			278, 974, 961, 974, 980, 719, 789, 790, 791, 690, 960, 939, 549, 855, 911, 933,
			911, 945, 974, 755, 846, 762, 761, 571, 677, 763, 760, 759, 754, 494, 552, 537,
			577, 692, 786, 788, 788, 790, 793, 794, 816, 823, 789, 841, 823, 833, 816, 831,
			923, 744, 723, 749, 790, 792, 695, 776, 768, 792, 759, 707, 708, 682, 701, 826,
			815, 789, 789, 707, 687, 696, 689, 786, 787, 713, 791, 785, 791, 873, 761, 762,
			762, 759, 759, 892, 892, 788, 784, 438, 138, 277, 415, 392, 392, 668, 668, 0,
			390, 390, 317, 317, 276, 276, 509, 509, 410, 410, 234, 234, 334, 334, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 732, 544, 544, 910, 667, 760, 760,
			776, 595, 694, 626, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788,
			788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788,
			788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 788, 894, 838, 1016, 458,
			748, 924, 748, 918, 927, 928, 928, 834, 873, 828, 924, 924, 917, 930, 931, 463,
			883, 836, 836, 867, 867, 696, 696, 874, 0, 874, 760, 946, 771, 865, 771, 888, 967,
			888, 831, 873, 927, 970, 918, 0},
		MissingWidth: 0,
		Leading:      0,
		Vscale:       0.0,
		Hscale:       0.0,
		FontMatrix:   [6]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0},
	}
)
