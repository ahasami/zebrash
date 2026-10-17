# Zebrash

> ## ℹ️ About this fork
> This is a fork of [ingridhq/zebrash](https://github.com/ingridhq/zebrash).
> The original library is no longer maintained, so I forked it to fix some bugs that we ran into
> when rendering real-world carrier labels. I will mainly maintain the features that I need 
> myself, but pull requests are welcome.

Library for rendering ZPL (Zebra Programming Language) files as raster images

- Partially based on https://github.com/BinaryKits/BinaryKits.Zpl
- Uses slightly modified implementations of:
	- PDF417, Aztec, Code 2 of 5 and EAN13 barcodes from https://github.com/boombuler/barcode/
	- Code 128, Code 39, QR Code and DataMatrix from https://github.com/makiuchi-d/gozxing

## Description

This library emulates subset of ZPL engine and allows you to view most of the ZPL labels that are used by carriers such as Fedex, UPS or DHL as PNGs without the need to possess physical Zebra-compatible printer.
Think of https://labelary.com/viewer.html except it is completely free for commercial use, has no API limits and can easily be self-hosted or plugged into existing Go application so you don't need to send labels with real customers information to some 3rd-party servers

Example of the output:

![UPS label](testdata/zebrash-default/ups_grayscale.png)


## Usage:

```go

	file, err := os.ReadFile("./testdata/label.zpl")
	if err != nil {
		t.Fatal(err)
	}

	parser := zebrash.NewParser()

	res, err := parser.Parse(file)
	if err != nil {
		t.Fatal(err)
	}

	var buff bytes.Buffer

	drawer := zebrash.NewDrawer()

	err = drawer.DrawLabelAsPng(res[0], &buff, drawers.DrawerOptions{
		LabelWidthMm:         101.6,
		LabelHeightMm:        203.2,
		Dpmm:                 8,
		EnableInvertedLabels: true,
		GrayscaleOutput:      true,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile("./testdata/label.png", buff.Bytes(), 0744)
	if err != nil {
		t.Fatal(err)
	}

```

### Command line

Render a ZPL file as a PNG next to it (`label.zpl` -> `label.png`):

```sh
go run ./cmd/zpl2png testdata/label.zpl
```

Options: `-w` label width in mm (default `101.6`), `-h` label height in mm (default `203.2`), `-dpmm` dots per mm (default `8`, use `12` for 300 dpi).
