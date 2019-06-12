package scanner

import (
	"strconv"
	"strings"

	"github.com/nicksellen/audiotags"
	"github.com/pkg/errors"
)

type tags struct {
	map_  map[string]string
	props *audiotags.AudioProperties
}

func readTags(path string) (*tags, error) {
	map_, props, err := audiotags.Read(path)
	if err != nil {
		return nil, errors.Wrap(err, "audiotags module")
	}
	return &tags{
		map_:  map_,
		props: props,
	}, nil
}

func (t *tags) firstTag(keys ...string) string {
	for _, key := range keys {
		if val, ok := t.map_[key]; ok {
			return val
		}
	}
	return ""
}

func intSep(in, sep string) int {
	if in == "" {
		return 0
	}
	start := strings.SplitN(in, sep, 2)[0]
	out, err := strconv.Atoi(start)
	if err != nil {
		return 0
	}
	return out
}

func (t *tags) Title() string       { return t.firstTag("title") }
func (t *tags) Artist() string      { return t.firstTag("artist") }
func (t *tags) Album() string       { return t.firstTag("album") }
func (t *tags) AlbumArtist() string { return t.firstTag("albumartist", "album artist") }
func (t *tags) Year() int           { return intSep(t.firstTag("date", "year"), "-") } // eg. 2019-6-11
func (t *tags) TrackNumber() int    { return intSep(t.firstTag("tracknumber"), "/") }  // eg. 5/12
func (t *tags) DiscNumber() int     { return intSep(t.firstTag("discnumber"), "/") }   // eg. 1/2
func (t *tags) Length() int         { return t.props.Length }
func (t *tags) Bitrate() int        { return t.props.Bitrate }