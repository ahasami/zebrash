// Command zpl2png renders a ZPL file as a PNG next to it (label.zpl -> label.png).
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ingridhq/zebrash"
	"github.com/ingridhq/zebrash/drawers"
)

func main() {
	widthMm := flag.Float64("w", 101.6, "label width in mm")
	heightMm := flag.Float64("h", 203.2, "label height in mm")
	dpmm := flag.Int("dpmm", 8, "dots per mm (8 = 203 dpi, 12 = 300 dpi)")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: zpl2png [-w mm] [-h mm] [-dpmm n] file.zpl...")
		os.Exit(2)
	}

	for _, src := range flag.Args() {
		if err := render(src, *widthMm, *heightMm, *dpmm); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", src, err)
			os.Exit(1)
		}
	}
}

func render(src string, widthMm, heightMm float64, dpmm int) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	labels, err := zebrash.NewParser().Parse(data)
	if err != nil {
		return err
	}

	base := strings.TrimSuffix(src, ".zpl")
	drawer := zebrash.NewDrawer()

	for i, label := range labels {
		var buff bytes.Buffer

		err := drawer.DrawLabelAsPng(label, &buff, drawers.DrawerOptions{
			LabelWidthMm:  widthMm,
			LabelHeightMm: heightMm,
			Dpmm:          dpmm,
		})
		if err != nil {
			return err
		}

		dst := base + ".png"
		if i > 0 {
			dst = fmt.Sprintf("%s_%d.png", base, i+1)
		}

		if err := os.WriteFile(dst, buff.Bytes(), 0644); err != nil {
			return err
		}
		fmt.Println("wrote", dst)
	}

	return nil
}
