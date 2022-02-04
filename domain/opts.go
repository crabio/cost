package domain

// CLI programm options
type Opts struct {
	// Path to scheme config file
	FilePath string `short:"f" long:"file" description:"A path to scheme config file" required:"true"`
}
