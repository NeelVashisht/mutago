package v1

import (
	"errors"
	"io/ioutil"
)

var genres = []string{
	"Blues", "Classic Rock", "Country", "Dance",
	"Disco", "Funk", "Grunge", "Hip-Hop",
	"Jazz", "Metal", "New Age", "Oldies",
	"Other", "Pop", "R&B", "Rap",
	"Reggae", "Rock", "Techno", "Industrial",
	"Alternative", "Ska", "Death Metal", "Pranks",
	"Soundtrack", "Euro-Techno", "Ambient", "Trip-Hop",
	"Vocal", "Jazz+Funk", "Fusion", "Trance",
	"Classical", "Instrumental", "Acid", "House",
	"Game", "Sound Clip", "Gospel", "Noise",
	"AlternRock", "Bass", "Soul", "Punk",
	"Space", "Meditative", "Instrumental Pop", "Instrumental Rock",
	"Ethnic", "Gothic", "Darkwave", "Techno-Industrial",
	"Electronic", "Pop-Folk", "Eurodance", "Dream",
	"Southern Rock", "Comedy", "Cult", "Gangsta",
	"Top 40", "Christian Rap", "Pop/Funk", "Jungle",
	"Native American", "Cabaret", "New Wave", "Psychadelic",
	"Rave", "Showtunes", "Trailer", "Lo-Fi",
	"Tribal", "Acid Punk", "Acid Jazz", "Polka",
	"Retro", "Musical", "Rock & Roll", "Hard Rock",
}

func ParseV1(reader io.ReadSeeker) (map[string]string, error) {

	reader.Seek(-128, os.SEEK_END)
	/*
	   v1 tags are in the last 128 bytes
	*/
	content := make([]byte, 128)

	_, err := io.ReadFull(reader, content)

	if err != nil || string(content[:3]) != "TAG" {
		return nil, errors.New("Error : ID3v1 tag not present")
	}

	tags := make(map[string]string)
	tags["TIT2"] = string(content[3:33])   //TITLE
	tags["TPE1"] = string(content[33:63])  //ARTIST
	tags["TALB"] = string(content[63:93])  //ALBUM
	tags["TYER"] = string(content[93:97])  //YEAR
	tags["COMM"] = string(content[97:127]) //COMMENTS
	tags["TCON"] = genres[content[127]]    //GENRE

	return tags, nil
}
