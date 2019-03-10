package gotafseer

type Chapter struct {
	Name  string
	Index int
}

type Tafseer struct {
	Id        int
	Name      string
	Language  string
	Author    string
	Book_Name string
}

type Verse struct {
	Ayah_Number int
	Sura_Index  int
	Sura_Name   string
	Text        string
}

type VerseTafseer struct {
	Ayah_Number  int
	Ayah_Url     string
	Tafseer_Id   int
	Tafseer_Name string
	Text         string
}
