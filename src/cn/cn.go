package main

import (
	"fmt"
	"github.com/SETTER2000/lengconv"
	"github.com/jessevdk/go-flags" // Импорт пакета go-flags с присвоением псевдонима flags
	"os"
	"strconv"
)

var opts struct {
	All                    bool `short:"a" long:"all" description:"все доступные варианты"`
	MeterToCentimeter      bool `short:"q" long:"mc" description:"метры в сантиметры"`
	MeterToFoot            bool `short:"z" long:"mf" description:"метры в футы"`
	MeterToMillimeter      bool `short:"x" long:"mtomm" description:"метры в миллиметры"`
	CentimeterToMeter      bool `short:"w" long:"cm" description:"сантиметры в метры"`
	CentimeterToFoot       bool `short:"u" long:"cf" description:"сантиметры в футы"`
	CentimeterToMillimeter bool `short:"v" long:"cmm" description:"сантиметры в миллиметры"`
	FootToCentimeter       bool `short:"g" long:"fc" description:"футы в сантиметры"`
	FootToMeter            bool `short:"j" long:"fm" description:"футы в метры"`
	FootToMillimeter       bool `short:"b" long:"fmm" description:"футы в миллиметры"`
	MillimeterToMeter      bool `short:"t" long:"mmtom" description:"миллиметры в метры"`
	MillimeterToFoot       bool `short:"p" long:"mmf" description:"миллиметры в футы"`
	MillimeterToCentimeter bool `short:"i" long:"mmc" description:"миллиметры в сантиметры"`
}

func main() {
	flags.Parse(&opts)
	st := "%-6s = %.2f "
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

		if opts.All {
			fmt.Printf(st+"ft\n"+st+"m\n"+st+"mm\n"+st+"cm\n"+st+"m\n"+st+"mm\n"+st+"ft\n"+st+"cm\n"+st+"mm\n"+st+"cm\n"+st+"ft\n"+st+"m\n",
				c, lengconv.CmToFt(c),
				c, lengconv.CmToM(c),
				c, lengconv.CmToMm(c),
				f, lengconv.FtToCm(f),
				f, lengconv.FtToM(f),
				f, lengconv.FtToMm(f),
				m, lengconv.MToFt(m),
				m, lengconv.MToCm(m),
				m, lengconv.MToMm(m),
				mm, lengconv.MmToCm(mm),
				mm, lengconv.MmToFt(mm),
				mm, lengconv.MmToM(mm),
			)
		}

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
