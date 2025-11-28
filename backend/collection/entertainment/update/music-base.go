// Package update guesses can be made via lead singer, song
package update

import (
	"fmt"
	"maps"
	"nyg/dataset"
)

type MusicParcel struct {
	Pack []EntertainmentParcel
}

// keys
const (
	MarathiSongs2014 = "MARATHI SONGS 2014"
	MarathiSongs2015 = "MARATHI SONGS 2015"
	MarathiSongs2016 = "MARATHI SONGS 2016"
	MarathiSongs2009 = "MARATHI SONGS 2014"

	HindiSongs2011 = "HINDI SONGS 2011"
	HindiSongs2012 = "HINDI SONGS 2011"

	EnglishSongs2010 = "ENGLISH SONGS 2010"
	EnglishSongs2012 = "ENGLISH SONGS 2012"
	EnglishSongs2015 = "ENGLISH SONGS 2015"

	MusicKey = "MUSIC"
)

func AlbumsMarathi2009() dataset.SearchSystem_ {
	m2009 := map[string][]string{}
	m2009["hridayamandhale gaane"] = []string{
		"ardhya raati",
		"bela shende",
		"door nabhanchya",
		"hridayamandhale gaane",
		"jantar mantar",
		"nilya manache",
		"runjhu runjhun",
		"tujhya vina",
	}
	m2009["mee shivajiraje bhosale boltoy"] = []string{
		"masuli wani tujhi ga jwani",
		"ajit parab",
		"bela shende",
		"maharajachi kirti befam (powada)",
		"nandesh umap",
		"o'raaje",
		"sukhwinder singh",
		"maharashtra geet",
		"umesh kamarkar",
		"neha rajpal",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(MarathiSongs2009, EntertainmentKey, MusicKey, m2009, false, true, false)
	return M
}
func AlbumsMarathi2014() dataset.SearchSystem_ {
	m2014 := map[string][]string{}
	m2014["lokmanya-ek yugpurush"] = []string{
		"gajananaa gajananaa",
		"shankar mahadevan",
		"guru thakur",
		"powada",
		"nandesh umap",
		"geeta nirupan",
		"narayan parshuram",
		"ajit sameer",
		"power of unity",
		"punashcha hari om",
		"spirit of lokmanya",
	}
	m2014["lai bhari"] = []string{
		"mauli mauli",
		"ajay gogavale",
		"new nava tarana",
		"kunal ganjawala",
		"jeev bhulala",
		"sonu nigam",
		"shreya ghoshal",
		"ala holicha sar lai bhari",
		"swapnil bhandodkar",
		"yogita godbole",
		"ye na saajna",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(MarathiSongs2014, EntertainmentKey, MusicKey, m2014, false, true, false)
	return M
}
func AlbumsMarathi2015() dataset.SearchSystem_ {
	m2015 := map[string][]string{}
	m2015["katyar kaljat ghusli"] = []string{
		"sur niragas ho",
		"anandi joshi",
		"shankar mahadevan",
		"dl ki tapish",
		"rahul deshpande",
		"tejonidhi lohagol",
		"ghei chhand",
		"man mandira",
		"man mandira tejane",
		"ghei chhand makarand",
		"lagi karejwa katar",
		"din gale",
		"sur se saji",
		"jitendra abhisheki",
		"muralidhar shyam",
		"bhola bhandari",
		"arjit singh",
		"yaar lllahi qawwali",
		"arshad muhammad",
		"divya kumar",
		"tarana",
		"suraj piya ki",
		"mahesh kale",
		"aruni kirani",
		"katyar kaljat ghusli-theme song",
		"shankar-ehsaan-loy",
	}
	m2015["nilkantha master"] = []string{
		"adhir man jhale",
		"shreya ghoshal",
		"paratun ye na",
		"javel ali",
		"kaunse des chala",
		"ajay gogavale",
		"vande mataram",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(MarathiSongs2015, EntertainmentKey, MusicKey, m2015, false, true, false)
	return M
}

func AlbumsMarathi2016() dataset.SearchSystem_ {
	m2016 := map[string][]string{}
	m2016["jau dya na balasheb"] = []string{
		"dolby walya",
		"nagesh morwekar",
		"earl edgar",
		"mona darling",
		"suman sridhar",
		"shreya ghoshal",
		"kunal ganjawala",
		"sonu nigam",
		"vaat disu de",
		"yogita godbole",
		"bring it on",
		"gondhal",
	}
	m2016["sairat"] = []string{
		"yad lagla",
		"ajay gogavale",
		"aatach baya ka baavaria",
		"sairat jhala ji",
		"chinmayi",
		"zingaat",
		"atul gogavale",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(MarathiSongs2016, EntertainmentKey, MusicKey, m2016, false, true, false)
	return M
}

func AlbumsHindi2011() dataset.SearchSystem_ {
	h2011 := map[string][]string{}
	h2011["rockstart"] = []string{
		"phir se ud chala",
		"mohit chauhan",
		"jo bhi main",
		"katiya karun",
		"harshdeep kaur",
		"kun faya kun",
		"a.r. rahman",
		"javed ali",
		"seher emin",
		"karthik",
		"haawa haawa",
		"aur ho",
		"alma ferovic",
		"tango for taj",
		"tum ko",
		"kavita krishnamurthy",
		"the dichotomy of fame",
		"nadan parinde",
		"tum ho",
		"suzanne d'mello",
		"sadda haq",
		"meeting place",
		"ranbir kapoor",
		"jaagran(rockstart)",
	}
	h2011["ra-one"] = []string{
		"chammak challo",
		"vishal-shekhar",
		"akon",
		"hamsika iyer",
		"vishal dadlani",
		"niranjan iyengar",
		"dildaara(stand by me)",
		"shafqat amanat ali",
		"kumaar",
		"criminal",
		"shruti pathak",
		"bhare naina",
		"nandini srikar",
		"panchhi jalonvi",
		"right by your side",
		"sidd coutto",
		"anvita dutt guptan",
		"raftaarein",
		"shekhar ravjani",
		"jiya more ghabrayee (the chase)",
		"sukhwinder singh",
		"anubhav sinha",
		"comes the light (theme)",
		"i'm on (theme)",
		"son of the end (theme)",
		"chammak challo (international version)",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(HindiSongs2011, EntertainmentKey, MusicKey, h2011, false, true, false)
	return M
}

func AlbumsHindi2012() dataset.SearchSystem_ {
	h2012 := map[string][]string{}
	h2012["vicky donor"] = []string{
		"rokda",
		"akshay verma",
		"aditi singh sharma",
		"kho jaane de",
		"clinton cerejo",
		"juhi chaturvedi",
		"rum whisky",
		"kusum verma",
		"pani da rang - male voclas",
		"aayushmann khurrana",
		"rochak kohli",
		"mar jayian - romantic",
		"vishal dadlani",
		"sunidhi chauhan",
		"bann chakraborty",
		"swanand kirkire",
		"chaddha",
		"mika singh",
		"vijay maurya",
		"pani da rang -female vocals",
		"sukanya purayastha",
		"mar jayian -sad",
	}
	h2012["khiladi786"] = []string{
		"lonely",
		"himesh reshammiya",
		"yo yo honey singh",
		"balma",
		"shreya ghoshal",
		"long drive",
		"mika singh",
		"sari sari raat",
		"hookah bar",
		"vineet singh",
		"aaman trikha",
		"khiladi title track",
		"yashraj kapil",
		"alamgir khan",
		"rajdeep chatterjee",
		"tu hoor pari",
		"javel ali",
		"chandrakala singh",
		"harshdeep kaur",
		"kiran kamath",
		"dj a sen",
		"dj amann nagpal",
		"teenu arora",
	}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(HindiSongs2012, EntertainmentKey, MusicKey, h2012, false, true, false)
	return M
}

func AlbumsEnglish2010() dataset.SearchSystem_ {
	e2010 := map[string][]string{}
	e2010["teenage dream"] = []string{"katy perry"}
	e2010["mylo xyloto"] = []string{"coldplay"}
	e2010["25"] = []string{"adele"}
	e2010["x"] = []string{"ed sheeran"}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(EnglishSongs2010, EntertainmentKey, MusicKey, e2010, true, false, false)
	return M
}
func AlbumsEnglish2012() dataset.SearchSystem_ {
	e2012 := map[string][]string{}
	e2012["blunderbuss"] = []string{"jack white"}
	e2012["wrecking ball"] = []string{"bruce springsteen"}
	e2012["mdna"] = []string{"madonna"}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(EnglishSongs2012, EntertainmentKey, MusicKey, e2012, true, false, false)
	return M
}
func AlbumsEnglish2015() dataset.SearchSystem_ {
	e2015 := map[string][]string{}
	e2015["the pinkprint"] = []string{"nicki minaj"}
	e2015["honeymoon"] = []string{"lana del rey"}
	e2015["bouquet (ep)"] = []string{"the chainsmokers"}
	M := dataset.SearchSystem_{}
	M.Constructor() // important
	M.Include(EnglishSongs2015, EntertainmentKey, MusicKey, e2015, true, false, false)
	return M
}

func BuildMusic() MusicParcel {
	m2009, m2014, m2015, m2016 := AlbumsMarathi2009(), AlbumsMarathi2014(), AlbumsMarathi2015(), AlbumsMarathi2016()
	h2011, h2012 := AlbumsHindi2011(), AlbumsHindi2012()
	e2010, e2012, e2015 := AlbumsEnglish2010(), AlbumsEnglish2012(), AlbumsEnglish2015()
	me2009, mc2009, mf2009, mb2009, mi2009 := m2009.Manager(MarathiSongs2009)
	me2014, mc2014, mf2014, mb2014, mi2014 := m2014.Manager(MarathiSongs2014)
	me2015, mc2015, mf2015, mb2015, mi2015 := m2015.Manager(MarathiSongs2015)
	me2016, mc2016, mf2016, mb2016, mi2016 := m2016.Manager(MarathiSongs2016)
	he2011, hc2011, hf2011, hb2011, hi2011 := h2011.Manager(HindiSongs2011)
	he2012, hc2012, hf2012, hb2012, hi2012 := h2012.Manager(HindiSongs2012)
	ee2010, ec2010, ef2010, eb2010, ei2010 := e2010.Manager(EnglishSongs2010)
	ee2012, ec2012, ef2012, eb2012, ei2012 := e2012.Manager(EnglishSongs2012)
	ee2015, ec2015, ef2015, eb2015, ei2015 := e2015.Manager(EnglishSongs2015)

	parcel1 := EntertainmentParcel{Event: me2009, Category: mc2009, Field: mf2009, Book: mb2009, MappedItems: mi2009}
	parcel2 := EntertainmentParcel{Event: me2014, Category: mc2014, Field: mf2014, Book: mb2014, MappedItems: mi2014}
	parcel3 := EntertainmentParcel{Event: me2015, Category: mc2015, Field: mf2015, Book: mb2015, MappedItems: mi2015}
	parcel4 := EntertainmentParcel{Event: me2016, Category: mc2016, Field: mf2016, Book: mb2016, MappedItems: mi2016}
	parcel5 := EntertainmentParcel{Event: he2011, Category: hc2011, Field: hf2011, Book: hb2011, MappedItems: hi2011}
	parcel6 := EntertainmentParcel{Event: he2012, Category: hc2012, Field: hf2012, Book: hb2012, MappedItems: hi2012}
	parcel7 := EntertainmentParcel{Event: ee2010, Category: ec2010, Field: ef2010, Book: eb2010, MappedItems: ei2010}
	parcel8 := EntertainmentParcel{Event: ee2012, Category: ec2012, Field: ef2012, Book: eb2012, MappedItems: ei2012}
	parcel9 := EntertainmentParcel{Event: ee2015, Category: ec2015, Field: ef2015, Book: eb2015, MappedItems: ei2015}

	ready := []EntertainmentParcel{parcel1, parcel2, parcel3, parcel4, parcel5, parcel6, parcel7, parcel8, parcel9}
	pack := MusicParcel{Pack: ready}
	return pack
}

func MusicEvents() []string {
	events := []string{}
	openBuild := BuildMusic()
	for _, event := range openBuild.Pack {
		events = append(events, event.Event)
	}
	return events
}

func MusicCategory() string {
	category := []string{}
	token := BuildMusic().Pack
	for _, cat := range token {
		category = append(category, cat.Category)
	}
	category = dataset.EraseDuplicate(category)
	cat := category[0]
	fmt.Println(category)
	return cat
}

func MusicBook() string {
	fi := ""
	token := BuildMusic().Pack
	for _, cat := range token {
		fi = cat.Book
	}
	return fi
}

func MusicLists() []string {
	ma := []string{}
	token := BuildMusic().Pack
	for _, t := range token {
		for items := range t.MappedItems {
			ma = append(ma, items)
		}
	}
	return ma
}

func MusicSheet() map[string][]string {
	sheet := map[string][]string{}
	token := BuildMusic().Pack
	for _, t := range token {
		maps.Copy(sheet, t.MappedItems)
	}
	return sheet
}
