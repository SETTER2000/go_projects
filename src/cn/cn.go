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
	MeterToInch            bool `short:"n" long:"mi" description:"метры в дюймы"`
	CentimeterToMeter      bool `short:"w" long:"cm" description:"сантиметры в метры"`
	CentimeterToFoot       bool `short:"u" long:"cf" description:"сантиметры в футы"`
	CentimeterToMillimeter bool `short:"v" long:"cmm" description:"сантиметры в миллиметры"`
	CentimeterToInch       bool `short:"c" long:"ci" description:"сантиметры в дюймы"`
	FootToCentimeter       bool `short:"g" long:"fc" description:"футы в сантиметры"`
	FootToMeter            bool `short:"j" long:"fm" description:"футы в метры"`
	FootToMillimeter       bool `short:"b" long:"fmm" description:"футы в миллиметры"`
	FootToInch             bool `short:"l" long:"fi" description:"футы в дюймы"`
	MillimeterToMeter      bool `short:"t" long:"mmtom" description:"миллиметры в метры"`
	MillimeterToFoot       bool `short:"p" long:"mmf" description:"миллиметры в футы"`
	MillimeterToCentimeter bool `short:"i" long:"mmc" description:"миллиметры в сантиметры"`
	MillimeterToInch       bool `short:"k" long:"mmi" description:"миллиметры в дюймы"`
	InchToMillimeter       bool `short:"m" long:"imm" description:"дюймы в миллиметры"`
	InchToCentimeter       bool `short:"s" long:"ic" description:"дюймы в сантиметры"`
	InchToMeter            bool `short:"d" long:"im" description:"дюймы в метры"`
	InchToFoot             bool `short:"f" long:"if" description:"дюймы в футы"`
}

func main() {
	flags.Parse(&opts)
	st := "%-4s = %.2f "
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
		i := lengconv.Inch(d)

		if opts.All {
			fmt.Printf(st+"ft\n"+st+"m\n"+st+"mm\n"+st+"in\n"+st+"cm\n"+st+"m\n"+st+"mm\n"+st+"in\n"+st+"ft\n"+st+"cm\n"+st+"mm\n"+st+"in\n"+st+"cm\n"+st+"ft\n"+st+"m\n"+st+"in\n"+st+"mm\n"+st+"cm\n"+st+"m\n"+st+"ft\n",
				c, lengconv.CmToFt(c),
				c, lengconv.CmToM(c),
				c, lengconv.CmToMm(c),
				c, lengconv.CmToIn(c),
				f, lengconv.FtToCm(f),
				f, lengconv.FtToM(f),
				f, lengconv.FtToMm(f),
				f, lengconv.FtToIn(f),
				m, lengconv.MToFt(m),
				m, lengconv.MToCm(m),
				m, lengconv.MToMm(m),
				m, lengconv.MToIn(m),
				mm, lengconv.MmToCm(mm),
				mm, lengconv.MmToFt(mm),
				mm, lengconv.MmToM(mm),
				mm, lengconv.MmToIn(mm),
				i, lengconv.InToMm(i),
				i, lengconv.InToCm(i),
				i, lengconv.InToM(i),
				i, lengconv.InToF(i),
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

		if opts.MeterToInch {
			fmt.Printf(st+"in\n", m, lengconv.MToIn(m))
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

		if opts.CentimeterToInch {
			fmt.Printf(st+"in\n", c, lengconv.CmToIn(c))
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

		if opts.FootToInch {
			fmt.Printf(st+"in\n", f, lengconv.FtToIn(f))
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

		if opts.MillimeterToInch {
			fmt.Printf(st+"in\n", mm, lengconv.MmToIn(mm))
		}

		if opts.InchToMillimeter {
			fmt.Printf(st+"mm\n", i, lengconv.InToMm(i))
		}

		if opts.InchToCentimeter {
			fmt.Printf(st+"cm\n", i, lengconv.InToCm(i))
		}

		if opts.InchToMeter {
			fmt.Printf(st+"m\n", i, lengconv.InToM(i))
		}

		if opts.InchToFoot {
			fmt.Printf(st+"ft\n", i, lengconv.InToF(i))
		}

	}
}
