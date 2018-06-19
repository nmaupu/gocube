package cli

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/jawher/mow.cli"
	"github.com/nmaupu/gocube/config"
	"github.com/nmaupu/gocube/data"
	"github.com/nmaupu/gocube/image"
	"github.com/signintech/gopdf"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	//595.28, 841.89 = A4
	PdfWidth                 = 595.28
	PdfHeight                = 841.89
	PdfMarginTop             = 15.0
	PdfMarginLeft            = 20.0
	PdfMarginRight           = 20.0
	PdfTitleTextHeight       = 20.0
	PdfTitlePaddingBottom    = 5.0
	PdfSubTitleTextHeight    = 20.0
	PdfSubTitlePaddingBottom = 5.0
	CellTitleTextHeight      = 20.0
	CellTitlePaddingBottom   = 5.0
	CellSetupTextHeight      = 15.0
	CellSetupPaddingBottom   = 10.0
	CellAlgTextHeight        = 15.0
	CellMarginTop            = 5.0
	CellMarginRight          = 5.0
	CellPaddingTop           = 5.0
	CellPaddingBottom        = 10.0
	CellPaddingLeft          = 5.0
	ImgWidthPtFull           = 110.0
	ImgWidthPtTop            = 80.0
	ImgWidthPt3d             = 80.0
	ImgPaddingRight          = 10.0

	FontFileName     = "ttf/rockwell.ttf"
	FontName         = "rockwell"
	FontTitleSize    = 18
	FontSubTitleSize = 16
	FontAlgTitleSize = 14
	FontSetupAlgSize = 12
	FontAlgSize      = 10
)

func exportPDF(cmd *cli.Cmd) {
	cmd.Spec = "[-f] -o"
	file := cmd.StringOpt("f file", "config.yaml", "Config file name (do not provide extension)")
	output := cmd.StringOpt("o output", "out.pdf", "Output PDF file name")

	cmd.Action = func() {
		var err error
		conf := config.Configure(*file)

		// Creating temp dir
		tmpDir, err := ioutil.TempDir("", "gocube")
		if err != nil {
			log.Fatal(err)
		}

		defer os.RemoveAll(tmpDir) // clean up

		log.Printf("Temporary dir = %s", tmpDir)

		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: PdfWidth, H: PdfHeight}})
		err = pdf.AddTTFFont(FontName, FontFileName)
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = pdf.SetFont(FontName, "", FontAlgSize)
		if err != nil {
			log.Print(err.Error())
			return
		}

		// Generating cube's images in temp dir
		var ctx *gg.Context
		nbImages := 0
		for kDraw, draw := range conf.Draw {
			preAlg := data.NewAlg(draw.PreAlg)
			postAlg := data.NewAlg(draw.PostAlg)
			colors := data.GetColors(draw.Colors...)

			// Create a new page for each draw section
			createNewPdfPage(&pdf, conf.Pdf.Title, draw.Title)
			nbImages = 0

			var mx, my float64
			for kSpec, v := range draw.Spec {
				c := data.NewCubeColors(conf.Cube.Size, float64(*cubieSize), colors)
				setupAlg := data.NewAlg(v.Algs[0]).Reverse()
				c.Execute(preAlg)
				c.Execute(setupAlg)
				c.Execute(postAlg)

				var imgWidthPt float64
				switch draw.View {
				case "top":
					ctx = c.DrawTopView("white")
					imgWidthPt = ImgWidthPtTop
				case "full":
					ctx = c.Draw()
					imgWidthPt = ImgWidthPtFull
				case "3d":
					ctx = c.Draw3d(600)
					imgWidthPt = ImgWidthPt3d
					ctx, _ = image.TrimImageWhite(ctx)
				default:
					panic(fmt.Sprintf("Incorrect view, %s", draw.View))
				}

				tmpFile := filepath.Join(tmpDir, fmt.Sprintf("file%d-%d", kDraw, kSpec))
				ctx.SavePNG(tmpFile)

				if nbImages%10 == 0 && nbImages > 0 {
					createNewPdfPage(&pdf, conf.Pdf.Title, draw.Title)
					nbImages = 0
				}

				// Building PDF cell
				imgSizeRatio := float64(ctx.Width()) / float64(ctx.Height())
				imgHeightPt := imgWidthPt / imgSizeRatio
				cellHeight := (imgHeightPt +
					CellPaddingTop +
					CellTitleTextHeight + CellTitlePaddingBottom +
					CellSetupTextHeight + CellSetupPaddingBottom +
					CellPaddingBottom)
				if nbImages%2 == 0 {
					mx = PdfMarginLeft
					my = PdfMarginTop + PdfTitleTextHeight + PdfTitlePaddingBottom + PdfSubTitleTextHeight + PdfSubTitlePaddingBottom + float64(nbImages/2)*(cellHeight+CellMarginTop)
				} else {
					mx = PdfWidth/2 + CellMarginRight/2
				}
				nbImages++

				printPdfCell(
					&pdf,
					v.Name,
					tmpFile,
					imgWidthPt,
					imgHeightPt,
					setupAlg.String(),
					v.Algs,
					mx,
					my)
			}
		}

		pdf.WritePdf(*output)
	}
}

func createNewPdfPage(pdf *gopdf.GoPdf, pageTitle, pageSubTitle string) {
	var my float64
	var rect gopdf.Rect

	pdf.AddPage()

	// Writing title
	my = PdfMarginTop
	pdf.SetX(0)
	pdf.SetY(my)
	pdf.SetFont(FontName, "", FontTitleSize)
	rect = gopdf.Rect{W: PdfWidth, H: PdfTitleTextHeight}
	pdf.CellWithOption(&rect, pageTitle, gopdf.CellOption{Align: gopdf.Middle | gopdf.Center})

	// Writing Subtitle
	my += PdfTitleTextHeight + PdfTitlePaddingBottom
	pdf.SetX(0)
	pdf.SetY(my)
	pdf.SetFont(FontName, "", FontSubTitleSize)
	rect = gopdf.Rect{W: PdfWidth, H: PdfSubTitleTextHeight}
	pdf.CellWithOption(&rect, pageSubTitle, gopdf.CellOption{Align: gopdf.Middle | gopdf.Center})
}

// Beware: pdf are using pt whereas images' are using px!
func printPdfCell(pdf *gopdf.GoPdf, title string, imgFileName string, imgWidthPt, imgHeightPt float64, setupAlg string, algs []string, x, y float64) {
	var mx, my float64
	var rect gopdf.Rect

	cellWidth := (PdfWidth-CellMarginRight)/2 - PdfMarginLeft
	cellHeight := (imgHeightPt +
		CellPaddingTop +
		CellTitleTextHeight + CellTitlePaddingBottom +
		CellSetupTextHeight + CellSetupPaddingBottom +
		CellPaddingBottom)

	pdf.SetLineWidth(1)
	pdf.RectFromUpperLeft(
		x, y,
		cellWidth,
		cellHeight)

	// Print title
	mx = x + CellPaddingLeft
	my = y + CellPaddingTop
	pdf.SetX(mx)
	pdf.SetY(my)
	pdf.SetFont(FontName, "", FontAlgTitleSize)
	rect = gopdf.Rect{W: cellWidth - CellPaddingLeft*2, H: CellTitleTextHeight}
	pdf.CellWithOption(&rect, title, gopdf.CellOption{Border: gopdf.Bottom, Align: gopdf.Middle | gopdf.Center})

	// Print setup alg
	//mx = x + CellPaddingLeft
	my += CellTitleTextHeight + CellTitlePaddingBottom
	pdf.SetX(mx)
	pdf.SetY(my)
	pdf.SetFont(FontName, "", FontSetupAlgSize)
	rect = gopdf.Rect{W: cellWidth - CellPaddingLeft*2, H: CellSetupTextHeight}
	pdf.CellWithOption(&rect, fmt.Sprintf("Setup: %s", setupAlg), gopdf.CellOption{Align: gopdf.Left | gopdf.Middle})

	// Insert image
	//mx = x + CellPaddingLeft
	my += CellSetupTextHeight + CellSetupPaddingBottom
	rect = gopdf.Rect{W: imgWidthPt, H: imgHeightPt}
	pdf.Image(imgFileName, mx, my, &rect)

	// Print all algs
	pdf.SetFont(FontName, "", FontAlgSize)
	mx = x + CellPaddingLeft + imgWidthPt + ImgPaddingRight
	//my = y + CellPaddingTop + CellTitleTextHeight + CellSetupTextHeight

	for k, alg := range algs {
		pdf.SetX(mx)
		pdf.SetY(my + float64(k*CellAlgTextHeight))
		pdf.Cell(nil, alg)
	}
}
