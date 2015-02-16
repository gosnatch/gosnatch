package tvdb

import (
	"testing"
	"encoding/xml"
)

const BANSHEE_TEST = `
<?xml version="1.0" encoding="UTF-8" ?>
<Data>
<Series>
<seriesid>259765</seriesid>
<language>en</language>
<SeriesName>Banshee</SeriesName>
<banner>graphical/259765-g2.jpg</banner>
<Overview>Lucas Hood is an ex-con and master thief who assumes the identity of the sheriff of Banshee, Pa., where he continues his criminal activities, even as he’s hunted by the shadowy gangsters he betrayed years earlier.</Overview>
<FirstAired>2013-01-11</FirstAired>
<IMDB_ID>tt2017109</IMDB_ID>
<zap2it_id>EP01564099</zap2it_id>
<id>259765</id>
</Series>
</Data>
`

func TestBanshee(t *testing.T) {
	var data GetSeriesData
	err := xml.Unmarshal([]byte(BANSHEE_TEST), &data)
	if err != nil {
		t.Error(err)
	}

	if data.Series[0].SeriesName != "Banshee" {
		t.Error("SeriesName != Banshee")
	}
}

const TOUCH_TEST = `
<?xml version="1.0" encoding="UTF-8" ?>
<Data>
<Series>
<seriesid>83105</seriesid>
<language>en</language>
<SeriesName>Touch</SeriesName>
<banner>graphical/83105-g.jpg</banner>
<Overview>Touch follows the lives of three people: Kazuya and Tatsuya Uesugi (identical twin brothers) and Minami Asakura. They've lived next to each other since they were babies, and their parents even built a playhouse in the yard in order to keep them from damaging the houses. As they grew into their teens, they suddenly noticed that one of them was a girl, and Kazuya and Minami became closer. Kazuya, the younger twin, was an excellent student (as was Minami).</Overview>
<FirstAired>1985-03-24</FirstAired>
<id>83105</id>
</Series>
<Series>
<seriesid>80229</seriesid>
<language>en</language>
<SeriesName>Touch Me, I'm Karen Taylor</SeriesName>
<banner>graphical/80229-g.jpg</banner>
<Overview>A brand new comedy series co-writen by The BAFTA-winning comedian Karen Taylor who brings several unforgettable characters to life in this brilliant new series as seen through Karen Taylors eyes </Overview>
<FirstAired>2006-03-28</FirstAired>
<id>80229</id>
</Series>
<Series>
<seriesid>70409</seriesid>
<language>en</language>
<SeriesName>The Evil Touch</SeriesName>
<FirstAired>1973-09-01</FirstAired>
<id>70409</id>
</Series>
<Series>
<seriesid>76379</seriesid>
<language>en</language>
<SeriesName>The Gentle Touch</SeriesName>
<banner>graphical/76379-g.jpg</banner>
<Overview>In 1980s Britain, with disaffected punks and skinheads taking to the streets, and the repercussions of racism triggering explosive riots, it is not the best time to be a member of Her Majesty's Constabulary - particularly if you are a woman. Set on location in colourful Soho and Covent Garden, The Gentle Touch tells the story of tough cop Maggie Forbes. But this is no conventional cops and robbers series - this is real life drama. At times shocking, at times moving, and always utterly gripping</Overview>
<FirstAired>1980-04-01</FirstAired>
<id>76379</id>
</Series>
<Series>
<seriesid>248935</seriesid>
<language>en</language>
<SeriesName>Touch (2012)</SeriesName>
<banner>graphical/248935-g2.jpg</banner>
<Overview>"Touch" is a preternatural drama where science and spirituality meet in which we are all interconnected. The show follows a group of unrelated characters. One of these is Martin Bohm (Kiefer Sutherland), a widower and a single father who is haunted by an inability to connect to his mute and severely autistic 10-year-old son, Jake. Martin has tried everything he could do in order to reach his son, but at no success. To spend his time, Jake has cast-off cell phones, disassembling them and manipulating the parts. This allows him to see the world in a different way entirely. Martin is visited by social worker Clea Hopkins. She insists on doing an evaluation of the living situation.

Clea sees Martin as a man whose life has become dominated by a child he can no longer control.</Overview>
<FirstAired>2012-01-25</FirstAired>
<IMDB_ID>tt1821681</IMDB_ID>
<zap2it_id>SH01419237</zap2it_id>
<id>248935</id>
</Series>
<Series>
<seriesid>260750</seriesid>
<language>en</language>
<SeriesName>A Touch of Cloth</SeriesName>
<banner>graphical/260750-g4.jpg</banner>
<Overview>Hannah plays DI Jack Cloth, who is called in to investigate an apparent series of serial killings alongside his new partner, DC Anne Oldman, described as a "plucky, no-nonsense sidekick". Playing with the cliches and conventions of British police dramas, subplots include Cloth dealing with visions of his dead wife and the bisexual DC Oldman coming to grips with her feelings for both her female fiancee and Cloth.</Overview>
<FirstAired>2012-08-26</FirstAired>
<IMDB_ID>tt2240991</IMDB_ID>
<id>260750</id>
</Series>
<Series>
<seriesid>207201</seriesid>
<language>en</language>
<SeriesName>A Touch Away</SeriesName>
<banner>graphical/207201-g.jpg</banner>
<Overview>Impossible love story between Zorik - new immigrant from Russia and Roha'Le - a daughter of Hassidic family.</Overview>
<FirstAired>2007-01-23</FirstAired>
<IMDB_ID>tt0896576</IMDB_ID>
<id>207201</id>
</Series>
<Series>
<seriesid>267651</seriesid>
<language>en</language>
<SeriesName>A Touch Away</SeriesName>
<id>267651</id>
</Series>
<Series>
<seriesid>76247</seriesid>
<language>en</language>
<SeriesName>A Touch of Frost</SeriesName>
<banner>graphical/3804-g.jpg</banner>
<Overview>Detective Inspector Jack Frost is a disorganised DI for the Denton Police Force and will do anything to see that justice is done, even if he has to break the rules.</Overview>
<FirstAired>1992-12-01</FirstAired>
<IMDB_ID>tt0108967</IMDB_ID>
<zap2it_id>EP00080032</zap2it_id>
<id>76247</id>
</Series>
</Data>
`
