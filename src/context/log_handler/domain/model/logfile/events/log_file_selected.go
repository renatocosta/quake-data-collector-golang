package events

const LogFileSelectedEvent = "LogFileSelected"

type LogFileSelected struct {
	Content []string
	//MetaData logfile.Metadata
}
