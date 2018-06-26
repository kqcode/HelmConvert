package convertopt

type ConvertOptions struct {
	//chart files to convert
	InputFiles []string
	//cloud provider
	Provider string
	//the .tgz file to convert
	TgzFile []string
}
