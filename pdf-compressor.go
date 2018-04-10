package main

import (
	"flag"
	"log"
	"os/exec"
	"fmt"
	"os"
	"strings"
	"path"
	"runtime"
)

func main() {

	var gsExecName string
	if runtime.GOOS == "windows" {
		gsExecName = "gswin32c"
	} else {
		gsExecName = "gs"
	}

	_, err := exec.LookPath(gsExecName)
	if err != nil {
		log.Fatal("ghostscript is not installed")
	}

	// TODO: Add usage examples to help message

	pathPtr := flag.String("i", "", "Path to the input file. REQUIRED")
	// TODO: Add output file flag
	sizePtr := flag.String("size", "medium", "Size of output file. Valid choices are small, medium, or large")
	noResize := flag.Bool("no-resize", false, "Whether to resize output or not. For presentation decks this should always be omitted.")
	aspectRatioX := flag.Float64("aspect-ratio-x", 16, "X aspect ratio. Defaults to 16:9.")
	aspectRatioY := flag.Float64("aspect-ratio-y", 9, "Y aspect ratio. Defaults to 16:9.")
	compatibilityLevelPtr := flag.String("compatibility-level", "1.7", "PDF Compatibility Level. Reduce if sharing broadly. 1.3 - 1.7 are supported.")
	lossinessPtr := flag.String("lossiness", "lossless", "Use Lossy or Lossless image compression. Lossy images are much smaller but tend to have compression artifacts.")
	flag.Parse()

	aspectRatio := *aspectRatioX / *aspectRatioY

	ValidatePath(*pathPtr)
	ValidateFileExistsAndIsPdf(*pathPtr)
	ValidateSizeString(*sizePtr)
	ValidateCompatibilityLevel(*compatibilityLevelPtr)
	ValidateLossinessString(*lossinessPtr)

	var args []string

	args = append(args, "-o", OutputFileName(*pathPtr, *sizePtr))
	args = append(args, "-sDEVICE=pdfwrite")
	args = append(args, "-dDOINTERPOLATE")
	args = append(args, "-dCompatibilityLevel=" + *compatibilityLevelPtr)
	args = append(args, "-dCompressPages=true")
	args = append(args, "-dCompressFonts=true")

	if !*noResize {
		width := fmt.Sprintf("%d", WidthInPoints(*sizePtr, aspectRatio))
		height := fmt.Sprintf("%d", HeightInPoints(*sizePtr, aspectRatio))

		args = append(args, "-dDEVICEWIDTHPOINTS=" + width)
		args = append(args, "-dDEVICEHEIGHTPOINTS=" + height)
		//args = append(args, "-dFIXEDMEDIA")
		args = append(args, "-dFitPage")
	}

	var dpi string

	if *lossinessPtr == "lossless" {
		dpi = "72"
		args = append(args, "-dAutoFilterColorImages=false")
		args = append(args, "-dAutoFilterGrayImages=false")
		args = append(args, "-dAutoFilterMonoImages=false")
		args = append(args, "-dColorImageFilter=/FlateEncode")
		args = append(args, "-dGrayImageFilter=/FlateEncode")
		args = append(args, "-dMonoImageFilter=/FlateEncode")
	} else {
		dpi = DPI(*sizePtr)
		// TODO: Figure out how to set options on JPEG filter in ghostscript. See https://www.ghostscript.com/doc/current/Devices.htm
		//args = append(args, "-r" + dpi)
		//args = append(args, "-dJPEGQ=100")
	}

	args = append(args, "-dDownsampleColorImages=true")
	args = append(args, "-dDownsampleGrayImages=true")
	args = append(args, "-dDownsampleMonoImages=true")
	args = append(args, "-dColorImageResolution=" + dpi)
	args = append(args, "-dGrayImageResolution=" + dpi)
	args = append(args, "-dMonoImageResolution=" + dpi)
	args = append(args, "-dDetectDuplicateImages=true")
	args = append(args, "-dColorImageDownsampleThreshold=1.0")
	args = append(args, "-dGrayImageDownsampleThreshold=1.0")
	args = append(args, "-dMonoImageDownsampleThreshold=1.0")
	args = append(args, "-dColorImageDownsampleType=/Bicubic")
	args = append(args, "-dGrayImageDownsampleType=/Bicubic")
	args = append(args, "-dMonoImageDownsampleType=/Bicubic")
	args = append(args, "-f", *pathPtr)

	cmd := exec.Command(gsExecName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ValidatePath(pathStr string) {
	if pathStr == "" {
		fmt.Println("Please specify an input PDF file.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateFileExistsAndIsPdf(pathStr string) {
	if _, err := os.Stat(pathStr); os.IsNotExist(err) {
		fmt.Println("Input file does not exist.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if strings.ToLower(path.Ext(pathStr)) != ".pdf" {
		fmt.Println("Input file extension is not PDF.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateSizeString(sizeStr string) {
	validSizes := []string{"small", "medium", "large"}
	if !Contains(validSizes, strings.ToLower(sizeStr)) {
		fmt.Println("Invalid size argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateCompatibilityLevel(levelStr string) {
	validLevels := []string{"1.3", "1.4", "1.5", "1.6", "1.7"}
	if !Contains(validLevels, strings.ToLower(levelStr)) {
		fmt.Println("Invalid compatibility level argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func ValidateLossinessString(lossinessStr string) {
	validConfigurations := []string{"lossy", "lossless"}
	if !Contains(validConfigurations, strings.ToLower(lossinessStr)) {
		fmt.Println("Invalid lossiness argument.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func WidthInPoints(sizeStr string, aspectRatio float64) int {
	if aspectRatio > 1 {
		return int(ShortLength(sizeStr) * aspectRatio)
	} else {
		return int(ShortLength(sizeStr))
	}
}

func HeightInPoints(sizeStr string, aspectRatio float64) int {
	if aspectRatio > 1 {
		return int(ShortLength(sizeStr))
	} else {
		return int(ShortLength(sizeStr) * (1 / aspectRatio))
	}
}

// Returns size of the short side of the paper based on size string argument.
func ShortLength(sizeStr string) float64 {
	switch strings.ToLower(sizeStr) {
	case "small":
		return 9 * 72
	case "large":
		return 16 * 72
	default:
		return 12 * 72
	}
}

// Returns string representation of DPI of images based on size string argument.
func DPI(sizeStr string) string {
	switch strings.ToLower(sizeStr) {
	case "small":
		return "150"
	case "large":
		return "300"
	default:
		return "225"
	}
}

func OutputFileName(pathStr string, sizeStr string) string {
	cutIndex := strings.LastIndex(strings.ToLower(pathStr), ".pdf")
	outputFileName := pathStr[:cutIndex] + "-" + sizeStr + ".pdf"
	fmt.Println(outputFileName)
	return outputFileName
}

// TODO: Pull out into utility file
func Contains(vs []string, t string) bool {
	for _, v := range vs {
		if v == t {
			return true
		}
	}
	return false
}
