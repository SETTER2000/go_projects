package main

import (
	"fmt"
	"github.com/SETTER2000/lengconv"
	"github.com/jessevdk/go-flags" // Импорт пакета go-flags с присвоением псевдонима flags
	"os"
	"strconv"
)

var opts struct {
	MeterToCentimeter      bool `short:"q" long:"mc" description:"Конвертировать метры в сантиметры."`
	MeterToFoot            bool `short:"z" long:"mf" description:"Конвертировать метры в футы."`
	MeterToMillimeter      bool `short:"х" long:"mmm" description:"Конвертировать метры в миллиметры."`
	CentimeterToMeter      bool `short:"w" long:"cm" description:"Конвертировать сантиметры в метры."`
	CentimeterToFoot       bool `short:"u" long:"cf" description:"Конвертировать сантиметры в футы."`
	CentimeterToMillimeter bool `short:"v" long:"cmm" description:"Конвертировать сантиметры в миллиметры."`
	FootToCentimeter       bool `short:"g" long:"fc" description:"Конвертировать футы в сантиметры."`
	FootToMeter            bool `short:"j" long:"fm" description:"Конвертировать футы в метры."`
	FootToMillimeter       bool `short:"b" long:"fmm" description:"Конвертировать футы в миллиметры."`
	MillimeterToMeter      bool `short:"j" long:"fm" description:"Конвертировать миллиметры в метры."`
	MillimeterToFoot       bool `short:"j" long:"fm" description:"Конвертировать миллиметры в футы."`
	MillimeterToCentimeter bool `short:"j" long:"fm" description:"Конвертировать миллиметры в сантиметры."`
}

func main() {
	flags.Parse(&opts)
	st := "%s = %.2f "
	for _, arg := range os.Args[2:] {
		d, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := lengconv.Meter(d)
		f := lengconv.Foot(d)
		c := lengconv.Centimeter(d)
		mm := lengconv.Millimeter(d)

		if opts.MeterToCentimeter {
			fmt.Printf(st+"см\n", m, lengconv.MToCm(m))
		}

		if opts.MeterToFoot {
			fmt.Printf(st+"ft\n", m, lengconv.MToFt(m))
		}

		if opts.MeterToMillimeter {
			fmt.Printf(st+"mm\n", m, lengconv.MToMm(m))
		}

		if opts.CentimeterToMeter {
			fmt.Printf(st+"м\n", c, lengconv.CmToM(c))
		}

		if opts.CentimeterToFoot {
			fmt.Printf(st+"ft\n", c, lengconv.CmToFt(c))
		}

		if opts.CentimeterToMillimeter {
			fmt.Printf(st+"mm\n", c, lengconv.CmToMm(c))
		}

		if opts.FootToCentimeter {
			fmt.Printf(st+"cm\n", f, lengconv.FtToCm(f))
		}

		if opts.FootToMeter {
			fmt.Printf(st+"m\n", f, lengconv.FtToM(f))
		}

		if opts.FootToMillimeter {
			fmt.Printf(st+"mm\n", f, lengconv.FtToMm(f))
		}

		if opts.MillimeterToMeter {
			fmt.Printf(st+"m\n", mm, lengconv.MmToM(mm))
		}
		if opts.MillimeterToFoot {
			fmt.Printf(st+"ft\n", mm, lengconv.MmToM(mm))
		}
		if opts.MillimeterToCentimeter {
			fmt.Printf(st+"cm\n", mm, lengconv.MmToM(mm))
		}
	}
}
