package models

// DumpMovie struct is the structure of a movie from a JSON file dumped
//  by tMDB daily with an added processed field to tell if the row in mySQL
//  database.
type DumpMovie struct {
	adult      bool
	id         int32
	title      string
	popularity float64
	video      bool
	processed  bool
}
