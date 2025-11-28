package update

import (
	"fmt"
	"maps"
	"nyg/dataset"
)

type SportsParcel struct {
	Event       string
	Category    string
	Field       string
	Book        string
	MappedItems map[string][]string
}

type CricketParcel struct {
	Pack []SportsParcel
}

// keys
const (
	CricketWorldCup2007Key = "ICC WORLD T20 2007"
	CricketWorldCup2009Key = "ICC WORLD T20 2009"

	CricketChampionsWorldCup2009Key = "ICC CHAMPIONS TROPHY 2009"
	CricketChampionsWorldCup2013Key = "ICC CHAMPIONS TROPHY 2013"

	CricketWorldCup2011Key = "ICC WORLD CUP 2011"
	CricketWorldCup2012Key = "ICC WORLD CUP 2012"
	CricketWorldCup2015Key = "ICC WORLD CUP 2015"
	CricketWorldCup2016Key = "ICC WORLD CUP 2016"

	IPL2008Key = "IPL 2008"

	CricketKey = "CRICKET"
	SportKey   = "Sports"
)

// international

func CricketWorldCup2007() dataset.SearchSystem_ {
	twoT07WC := map[string][]string{}

	twoT07WC["austrillia"] = []string{
		"Adam Gilchrist",
		"Michael Clarke",
		"Brad Haddin",
		"Matthew Hayden",
		"Brad Hodge",
		"Michael Hussey",
		"Ricky Ponting",
		"Brad Hogg",
		"Andrew Symonds",
		"Shane Watson",
		"Nathan Bracken",
		"Stuart Clark",
		"Ben Hilfenhaus",
		"Mitchell Johnson",
		"Brett Lee",
		"Shaun Tait",
	}

	twoT07WC["bangladesh"] = []string{
		"Mohammad Ashraful",
		"Aftab Ahmed",
		"Zunaed Siddique",
		"Mushfiqur Rahim",
		"Nadif Chowdhury",
		"Nazimuddin",
		"Tamim Iqbal",
		"Alok Kapali",
		"Farhad Reza",
		"Mahmudullah",
		"Shakib Al Hasan",
		"Ziaur Rahman",
		"Mashrafe Mortaza",
		"Abdur Razzak",
		"Syed Rasel",
	}

	twoT07WC["england"] = []string{
		"Kevin Pietersen",
		"Matt Prior",
		"Owais Shah",
		"Vikram Solanki",
		"Luke Wright",
		"Paul Collingwood",
		"Andrew Flintoff",
		"Darren Maddy",
		"Dimitri Mascarenhas",
		"Chris Schofield",
		"Ravi Bopara",
		"James Anderson",
		"Stuart Broad",
		"James Kirtley",
		"Jeremy Snape",
		"Chris Tremlett",
		"Ryan Sidebottom",
	}

	twoT07WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Yuvraj Singh",
		"Gautam Gambhir",
		"Dinesh Karthik",
		"Virender Sehwag",
		"Rohit Sharma",
		"Robin Uthappa",
		"Piyush Chawla",
		"Joginder Sharma",
		"Yusuf Pathan",
		"Irfan Pathan",
		"Ajit Agarkar",
		"Harbhajan Singh",
		"RP Singh",
		"Sreesanth",
	}

	twoT07WC["kenya"] = []string{
		"Tanmay Mishra",
		"Alex Obanda",
		"David Obuya",
		"Morris Ouma",
		"Steve Tikolo",
		"Collins Obuya",
		"Thomas Odoyo",
		"Tony Suji",
		"Rajesh Bhudia",
		"Jimmy Kamande",
		"Nehemiah Odhiambo",
		"Peter Ongondo",
		"Lameck Onyango",
		"Elijah Otieno",
		"Hiren Varaiya",
	}

	twoT07WC["new zeland"] = []string{
		"Peter Fulton",
		"Gareth Hopkins",
		"Brendon McCullum",
		"Craig McMillan",
		"Ross Taylor",
		"Lou Vincent",
		"Daniel Vettori",
		"Jacob Oram",
		"Scott Styris",
		"Shane Bond",
		"Mark Gillespie",
		"Chris Martin",
		"Nathan McCullum",
		"Jeetan Patel",
		"Bradley Scott",
	}

	twoT07WC["pakistan"] = []string{
		"Fawad Alam",
		"Imran Nazir",
		"Kamran Akmal",
		"Misbah ul Haq",
		"Salman Butt",
		"Younis Khan",
		"Shoaib Malik",
		"Mohammad Hafeez",
		"Shahid Afridi",
		"Yasir Arafat",
		"Abdur Rehman",
		"Iftikhar Anjum",
		"Mohammad Asif",
		"Sohail Tanvir",
		"Umar Gul",
		"Shoaib Akhtar",
	}

	twoT07WC["scotland"] = []string{
		"Ryan Watson",
		"Gavin Hamilton",
		"Gregor Maiden",
		"Neil McCallum",
		"Navdeep Poonia",
		"Qasim Sheikh",
		"Colin Smith",
		"Fraser Watts",
		"John Blain",
		"Dougie Brown",
		"Gordon Drummond",
		"Majid Haq",
		"Dewald Nel",
		"Craig Wright",
		"Ross Lyons",
	}

	twoT07WC["south africa"] = []string{
		"Graeme Smith",
		"Gulam Bodi",
		"Mark Boucher",
		"AB de Villiers",
		"Herschelle Gibbs",
		"Loots Bosman",
		"Jean Paul Duminy",
		"Justin Kemp",
		"Albie Morkel",
		"Vernon Philander",
		"Shaun Pollock",
		"Johan van der Wath",
		"Morne Morkel",
		"Andre Nel",
		"Makhaya Ntini",
		"Thandi Tshabalala",
	}

	twoT07WC["sri lanka"] = []string{
		"Mahela Jayawardene",
		"Kumar Sangakkara",
		"Chamara Silva",
		"Upul Tharanga",
		"Tillakaratne Dilshan",
		"Hasantha Fernando",
		"Sanath Jayasuriya",
		"Kaushal Lokuarachchi",
		"Farveez Maharoof",
		"Jehan Mubarak",
		"Dilruwan Perera",
		"Gayan Wijekoon",
		"Dilhara Fernando",
		"Lasith Malinga",
		"Chaminda Vaas",
		"Muthiah Muralidaran",
	}

	twoT07WC["west indies"] = []string{
		"Ramnaresh Sarwan",
		"Shivnarine Chanderpaul",
		"Narsingh Deonarine",
		"Runako Morton",
		"Denesh Ramdin",
		"Marlon Samuels",
		"Devon Smith",
		"Dwayne Bravo",
		"Chris Gayle",
		"Daren Sammy",
		"Dwayne Smith",
		"Pedro Collins",
		"Fidel Edwards",
		"Daren Powell",
		"Ravi Rampaul",
	}
	twoT07WC["zimbabwe"] = []string{
		"Chamu Chibhabha",
		"Timycen Maruma",
		"Hamilton Masakadza",
		"Stuart Matsikenyeri",
		"Vusi Sibanda",
		"Tatenda Taibu",
		"Brendan Taylor",
		"Elton Chigumbura",
		"Keith Dabengwa",
		"Sean Williams",
		"Prosper Utseya",
		"Gary Brent",
		"Chris Mpofu",
		"Tawanda Mupariwa",
		"Johnson Marumisa",
	}

	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2007Key, SportKey, CricketKey, twoT07WC, true, false, false)
	return ICCTEWC
}
func CricketWorldCup2009() dataset.SearchSystem_ {
	twoT09WC := map[string][]string{}

	twoT09WC["austrillia"] = []string{
		"Ricky Ponting",
		"Michael Clarke",
		"Brad Haddin",
		"Michael Hussey",
		"David Warner",
		"Cameron White",
		"James Hopes",
		"David Hussey",
		"Shane Watson",
		"Andrew Symonds",
		"Nathan Bracken",
		"Nathan Hauritz",
		"Ben Hilfenhaus",
		"Mitchell Johnson",
		"Brett Lee",
		"Peter Siddle",
	}

	twoT09WC["bangladesh"] = []string{
		"Mohammad Ashraful",
		"Zunaed Siddique",
		"Mohammad Mithun",
		"Mushfiqur Rahim",
		"Raqibul Hasan",
		"Shamsur Rahman",
		"Tamim Iqbal",
		"Mahmudullah",
		"Naeem Islam",
		"Shakib Al Hasan",
		"Abdur Razzak",
		"Mashrafe Mortaza",
		"Rubel Hossain",
		"Shahadat Hossain",
		"Syed Rasel",
	}

	twoT09WC["england"] = []string{
		"James Foster",
		"Rob Key",
		"Eoin Morgan",
		"Kevin Pietersen",
		"Owais Shah",
		"Luke Wright",
		"Paul Collingwood",
		"Ravi Bopara",
		"Dimitri Mascarenhas",
		"Graham Napier",
		"Andrew Flintoff",
		"James Anderson",
		"Stuart Broad",
		"Adil Rashid",
		"Ryan Sidebottom",
		"Graeme Swann",
	}

	twoT09WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Yuvraj Singh",
		"Gautam Gambhir",
		"Dinesh Karthik",
		"Suresh Raina",
		"Rohit Sharma",
		"Virender Sehwag",
		"Ravindra Jadeja",
		"Yusuf Pathan",
		"Irfan Pathan",
		"Harbhajan Singh",
		"Zaheer Khan",
		"Praveen Kumar",
		"Pragyan Ojha",
		"Ishant Sharma",
		"RP Singh",
	}

	twoT09WC["ireland"] = []string{
		"William Porterfield",
		"Niall O'Brien",
		"Gary Wilson",
		"Andre Botha",
		"Kyle McCallan",
		"John Mooney",
		"Kevin O'Brien",
		"Paul Stirling",
		"Andrew White",
		"Peter Connell",
		"Alex Cusack",
		"Trent Johnston",
		"Boyd Rankin",
		"Regan West",
		"Jeremy Bray",
	}
	twoT09WC["netherland"] = []string{
		"Jeroen Smits",
		"Tom de Grooth",
		"Tim Gruijters",
		"Alexei Kervezee",
		"Darron Reekers",
		"Eric Szwarczynski",
		"Bas Zuiderent",
		"Peter Borren",
		"Mudassar Bukhari",
		"Pieter Seelaar",
		"Ryan ten Doeschate",
		"Daan van Bunge",
		"Maurits Jonkman",
		"Dirk Nannes",
		"Edgar Schiferli",
		"Ruud Nijman",
	}
	twoT09WC["new zeland"] = []string{
		"Neil Broom",
		"Brendon Diamanti",
		"Martin Guptill",
		"Brendon McCullum",
		"Peter McGlashan",
		"Aaron Redmond",
		"Ross Taylor",
		"Daniel Vettori",
		"James Franklin",
		"Jacob Oram",
		"Scott Styris",
		"Jesse Ryder",
		"Ian Butler",
		"Nathan McCullum",
		"Kyle Mills",
		"Iain O'Brien",
	}

	twoT09WC["pakistan"] = []string{
		"Younis Khan",
		"Ahmed Shehzad",
		"Fawad Alam",
		"Kamran Akmal",
		"Misbah ul Haq",
		"Salman Butt",
		"Shahzaib Hasan",
		"Abdul Razzaq",
		"Shahid Afridi",
		"Shoaib Malik",
		"Yasir Arafat",
		"Iftikhar Anjum",
		"Mohammad Amir",
		"Saeed Ajmal",
		"Sohail Tanvir",
		"Umar Gul",
		"Shoaib Akhtar",
	}

	twoT09WC["scotland"] = []string{
		"Gavin Hamilton",
		"Richie Berrington",
		"Kyle Coetzer",
		"Calum MacLeod",
		"Neil McCallum",
		"Navdeep Poonia",
		"Colin Smith",
		"Ryan Watson",
		"Fraser Watts",
		"Gordon Drummond",
		"Majid Haq",
		"Dewald Nel",
		"Glenn Rogers",
		"Jan Stander",
		"Craig Wright",
		"John Blain",
	}

	twoT09WC["south africa"] = []string{
		"Graeme Smith",
		"Mark Boucher",
		"AB de Villiers",
		"Herschelle Gibbs",
		"Johan Botha",
		"Jean Paul Duminy",
		"Jacques Kallis",
		"Albie Morkel",
		"Justin Ontong",
		"Roelof van der Merwe",
		"Yusuf Abdulla",
		"Morne Morkel",
		"Wayne Parnell",
		"Robin Peterson",
		"Dale Steyn",
	}

	twoT09WC["sri lanka"] = []string{
		"Kumar Sangakkara",
		"Mahela Jayawardene",
		"Chamara Silva",
		"Tillakaratne Dilshan",
		"Sanath Jayasuriya",
		"Farveez Maharoof",
		"Angelo Mathews",
		"Jehan Mubarak",
		"Isuru Udana",
		"Muthiah Muralidaran",
		"Nuwan Kulasekara",
		"Lasith Malinga",
		"Ajantha Mendis",
		"Thilan Thushara",
		"Indika de Saram",
	}

	twoT09WC["west indies"] = []string{
		"Denesh Ramdin",
		"Shivnarine Chanderpaul",
		"Andre Fletcher",
		"Xavier Marshall",
		"Ramnaresh Sarwan",
		"Lendl Simmons",
		"Chris Gayle",
		"Dwayne Bravo",
		"Kieron Pollard",
		"Daren Sammy",
		"Lionel Baker",
		"Sulieman Benn",
		"Fidel Edwards",
		"Jerome Taylor",
		"David Bernard",
	}

	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2009Key, SportKey, CricketKey, twoT09WC, true, false, false)
	return ICCTEWC
}
func CricketChampionsTrophy2009() dataset.SearchSystem_ {
	twoT09CT := map[string][]string{}

	twoT09CT["austrillia"] = []string{
		"Ricky Ponting",
		"Callum Ferguson",
		"Michael Hussey",
		"Tim Paine",
		"Adam Voges",
		"Cameron White",
		"Michael Clarke",
		"Brad Haddin",
		"James Hopes",
		"David Hussey",
		"Shane Watson",
		"Doug Bollinger",
		"Nathan Hauritz",
		"Ben Hilfenhaus",
		"Mitchell Johnson",
		"Brett Lee",
		"Peter Siddle",
		"Nathan Bracken",
	}

	twoT09CT["england"] = []string{
		"Andrew Strauss",
		"Steven Davies",
		"Joe Denly",
		"Eoin Morgan",
		"Owais Shah",
		"Luke Wright",
		"Matt Prior",
		"Ravi Bopara",
		"Tim Bresnan",
		"Paul Collingwood",
		"Andrew Flintoff",
		"James Anderson",
		"Stuart Broad",
		"Graham Onions",
		"Adil Rashid",
		"Ryan Sidebottom",
		"Graeme Swann",
	}

	twoT09CT["india"] = []string{
		"Mahendra Singh Dhoni",
		"Rahul Dravid",
		"Gautam Gambhir",
		"Dinesh Karthik",
		"Virat Kohli",
		"Suresh Raina",
		"Sachin Tendulkar",
		"Yuvraj Singh",
		"Abhishek Nayar",
		"Yusuf Pathan",
		"Harbhajan Singh",
		"Praveen Kumar",
		"Amit Mishra",
		"Ashish Nehra",
		"Ishant Sharma",
		"RP Singh",
	}

	twoT09CT["new zeland"] = []string{
		"Neil Broom",
		"Brendon Diamanti",
		"Martin Guptill",
		"Gareth Hopkins",
		"Brendon McCullum",
		"Aaron Redmond",
		"Ross Taylor",
		"Daniel Vettori",
		"Grant Elliott",
		"James Franklin",
		"Jacob Oram",
		"Jesse Ryder",
		"Shane Bond",
		"Ian Butler",
		"Kyle Mills",
		"Iain O'Brien",
		"Jeetan Patel",
		"Daryl Tuffey",
	}

	twoT09CT["pakistan"] = []string{
		"Younis Khan",
		"Fawad Alam",
		"Imran Nazir",
		"Kamran Akmal",
		"Misbah ul Haq",
		"Mohammad Yousuf",
		"Umar Akmal",
		"Shahid Afridi",
		"Shoaib Malik",
		"Iftikhar Anjum",
		"Mohammad Amir",
		"Mohammad Asif",
		"Naved ul Hasan",
		"Saeed Ajmal",
		"Umar Gul",
	}

	twoT09CT["south africa"] = []string{
		"Graeme Smith",
		"Hashim Amla",
		"Mark Boucher",
		"AB de Villiers",
		"Herschelle Gibbs",
		"Johan Botha",
		"Jean Paul Duminy",
		"Jacques Kallis",
		"Albie Morkel",
		"Roelof van der Merwe",
		"Makhaya Ntini",
		"Wayne Parnell",
		"Robin Peterson",
		"Dale Steyn",
		"Lonwabo Tsotsobe",
	}

	twoT09CT["sri lanka"] = []string{
		"Kumar Sangakkara",
		"Mahela Jayawardene",
		"Thilina Kandamby",
		"Chamara Kapugedera",
		"Thilan Samaraweera",
		"Upul Tharanga",
		"Tillakaratne Dilshan",
		"Sanath Jayasuriya",
		"Angelo Mathews",
		"Nuwan Kulasekara",
		"Lasith Malinga",
		"Ajantha Mendis",
		"Muthiah Muralidaran",
		"Dhammika Prasad",
		"Thilan Thushara",
	}

	twoT09CT["west indies"] = []string{
		"Floyd Reifer",
		"Royston Crandon",
		"Travis Dowlin",
		"Andre Fletcher",
		"Kieran Powell",
		"Dale Richards",
		"Devon Smith",
		"Chadwick Walton",
		"Daren Sammy",
		"Tino Best",
		"Nikita Miller",
		"Kemar Roach",
		"Gavin Tonge",
		"Daren Powell",
		"David Bernard",
		"Kevin McClean",
	}

	ICCTECT := dataset.SearchSystem_{}
	ICCTECT.Constructor() // important
	ICCTECT.Include(CricketChampionsWorldCup2009Key, SportKey, CricketKey, twoT09CT, true, false, false)
	return ICCTECT
}
func IPL2008() dataset.SearchSystem_ {
	ipl2008 := map[string][]string{}
	ipl2008["chennai super kings"] = []string{
		"Mahendra Singh Dhoni",
		"Parthiv Patel",
		"Abhinav Mukund",
		"Srikkanth Anirudha",
		"Subramaniam Badrinath",
		"Suresh Raina",
		"Michael Hussey",
		"Arun Karthik",
		"Ravichandran Ashwin",
		"Albie Morkel",
		"Jacob Oram",
		"Joginder Sharma",
		"Shadab Jakati",
		"Sudeep Tyagi",
		"Lakshmipathy Balaji",
		"Manpreet Gony",
		"Muttiah Muralitharan",
	}
	ipl2008["mumbai indians"] = []string{
		"Sachin Tendulkar",
		"Yogesh Takawale",
		"Ajinkya Rahane",
		"Ankeet Chavan",
		"Siddharth Chitnis",
		"Saurabh Tiwary",
		"Manish Pandey",
		"Robin Uthappa",
		"Abhishek Nayar",
		"Dwayne Bravo",
		"Dwayne Smith",
		"Harbhajan Singh",
		"Swapnil Singh",
		"Lasith Malinga",
		"Dhawal Kulkarni",
	}
	ipl2008["kolkata knight riders"] = []string{
		"Brendon McCullum",
		"Wriddhiman Saha",
		"Debabrata Das",
		"Aakash Chopra",
		"Cheteshwar Pujara",
		"Brad Hodge",
		"Ricky Ponting",
		"Sourav Ganguly",
		"Laxmi Ratan Shukla",
		"Chris Gayle",
		"Ajit Agarkar",
		"David Hussey",
		"Iqbal Abdulla",
		"Ajantha Mendis",
		"Ishant Sharma",
		"Ashok Dinda",
		"Murali Kartik",
	}
	ipl2008["delhi daredevils"] = []string{
		"Virender Sehwag",
		"Dinesh Karthik",
		"AB de Villiers",
		"Shikhar Dhawan",
		"Gautam Gambhir",
		"Manoj Tiwary",
		"Mithun Manhas",
		"Yogesh Nagar",
		"Tillakaratne Dilshan",
		"Daniel Vettori",
		"Rajat Bhatia",
		"Yo Mahesh",
		"Pradeep Sangwan",
		"Amit Mishra",
	}
	ipl2008["kings XI punjab"] = []string{
		"Kumar Sangakkara",
		"Nitin Saini",
		"Sunny Sohal",
		"Tanmay Srivastava",
		"Luke Pomersbach",
		"Mahela Jayawardene",
		"Shaun Marsh",
		"Yuvraj Singh",
		"Rishi Dhawan",
		"Irfan Pathan",
		"Brett Lee",
		"Piyush Chawla",
		"Ramesh Powar",
		"Shanthakumaran Sreesanth",
	}
	ipl2008["royal challengers banglore"] = []string{
		"Rahul Dravid",
		"Shreevats Goswami",
		"Ross Taylor",
		"Bharat Chipli",
		"Virat Kohli",
		"Cameron White",
		"Jacques Kallis",
		"Praveen Kumar",
		"KP Appanna",
		"Zaheer Khan",
		"Vinay Kumar",
		"Dale Steyn",
	}
	ipl2008["rajasthan royals"] = []string{
		"Mahesh Rawat",
		"Graeme Smith",
		"Mohammad Kaif",
		"Ravindra Jadeja",
		"Gajendra Singh",
		"Yusuf Pathan",
		"Shane Watson",
		"Dimitri Mascarenhas",
		"Dinesh Salunkhe",
		"Pankaj Singh",
		"Siddharth Trivedi",
		"Munaf Patel",
		"Morne Morkel",
	}
	ipl2008["deccan charges"] = []string{
		"VVS Laxman",
		"Herschelle Gibbs",
		"Adam Gilchrist",
		"Dwaraka Ravi Teja",
		"Rohit Sharma",
		"Chamara Silva",
		"Venugopal Rao",
		"Arjun Yadav",
		"Shahid Afridi",
		"Scott Styris",
		"Andrew Symonds",
		"Pragyan Ojha",
		"RP Singh",
		"Chaminda Vaas",
		"Doddapaneni Kalyankrishna",
		"PM Sarvesh Kumar",
		"Paidikalva Vijaykumar",
	}
	c := dataset.SearchSystem_{}
	c.Constructor() // important
	c.Include(IPL2008Key, SportKey, CricketKey, ipl2008, false, true, false)
	return c
}

func CricketWorldCup2011() dataset.SearchSystem_ {
	twoT11WC := map[string][]string{}

	twoT11WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Virender Sehwag",
		"Gautam Gambhir",
		"Virat Kohli",
		"Suresh Raina",
		"Sachin Tendulkar",
		"Yuvraj Singh",
		"Ravichandran Ashwin",
		"Piyush Chawla",
		"Yusuf Pathan",
		"Harbhajan Singh",
		"Zaheer Khan",
		"Ashish Nehra",
		"Munaf Patel",
		"Sreesanth",
		"Praveen Kumar",
	}
	twoT11WC["austrillia"] = []string{
		"Ricky Ponting",
		"Michael Clarke",
		"Callum Ferguson",
		"Brad Haddin",
		"Michael Hussey",
		"Tim Paine",
		"Steven Smith",
		"Cameron White",
		"John Hastings",
		"David Hussey",
		"Shane Watson",
		"Mitchell Johnson",
		"Jason Krejza",
		"Brett Lee",
		"Shaun Tait",
		"Doug Bollinger",
		"Nathan Hauritz",
	}
	twoT11WC["england"] = []string{
		"Andrew Strauss",
		"Ian Bell",
		"Eoin Morgan",
		"Matt Prior",
		"Jonathan Trott",
		"Luke Wright",
		"Kevin Pietersen",
		"Ravi Bopara",
		"Tim Bresnan",
		"Paul Collingwood",
		"James Tredwell",
		"Michael Yardy",
		"James Anderson",
		"Jade Dernbach",
		"Adil Rashid",
		"Graeme Swann",
		"Chris Tremlett",
		"Stuart Broad",
		"Ajmal Shahzad",
	}
	twoT11WC["pakistan"] = []string{
		"Misbah ul Haq",
		"Ahmed Shehzad",
		"Asad Shafiq",
		"Kamran Akmal",
		"Umar Akmal",
		"Younis Khan",
		"Shahid Afridi",
		"Abdul Razzaq",
		"Mohammad Hafeez",
		"Abdur Rehman",
		"Junaid Khan",
		"Saeed Ajmal",
		"Shoaib Akhtar",
		"Umar Gul",
		"Wahab Riaz",
		"Sohail Tanvir",
	}
	twoT11WC["srilanka"] = []string{
		"Kumar Sangakkara",
		"Mahela Jayawardene",
		"Chamara Kapugedera",
		"Thilan Samaraweera",
		"Chamara Silva",
		"Upul Tharanga",
		"Tillakaratne Dilshan",
		"Thisara Perera",
		"Angelo Mathews",
		"Nuwan Kulasekara",
		"Lasith Malinga",
		"Ajantha Mendis",
		"Muthiah Muralidaran",
		"Suraj Randiv",
	}
	twoT11WC["south africa"] = []string{

		"Graeme Smith",
		"Hashim Amla",
		"AB de Villiers",
		"Faf du Plessis",
		"Colin Ingram",
		"Morne van Wyk",
		"Johan Botha",
		"Jean Paul Duminy",
		"Jacques Kallis",
		"Imran Tahir",
		"Morne Morkel",
		"Wayne Parnell",
		"Robin Peterson",
		"Dale Steyn",
		"Lonwabo Tsotsobe",
	}
	twoT11WC["bangladesh"] = []string{
		"Tamim Iqbal",
		"Imrul Kayes",
		"Zunaed Siddique",
		"Mohammad Ashraful",
		"Mushfiqur Rahim",
		"Raqibul Hasan",
		"Shahriar Nafees",
		"Shakib Al Hasan",
		"Mahmudullah",
		"Naeem Islam",
		"Abdur Razzak",
		"Nazmul Hossain",
		"Rubel Hossain",
		"Shafiul Islam",
		"Sohrawordi Shuvo",
	}
	twoT11WC["new zeland"] = []string{

		"Martin Guptill",
		"Jamie How",
		"Brendon McCullum",
		"Ross Taylor",
		"Kane Williamson",
		"Daniel Vettori",
		"James Franklin",
		"Jacob Oram",
		"Jesse Ryder",
		"Scott Styris",
		"Luke Woodcock",
		"Nathan McCullum",
		"Andy McKay",
		"Tim Southee",
		"Daryl Tuffey",
		"Hamish Bennett",
		"Kyle Mills",
	}
	twoT11WC["west indies"] = []string{

		"Darren Bravo",
		"Shivnarine Chanderpaul",
		"Kirk Edwards",
		"Ramnaresh Sarwan",
		"Devon Smith",
		"Devon Thomas",
		"Adrian Barath",
		"Carlton Baugh",
		"Daren Sammy",
		"Chris Gayle",
		"Kieron Pollard",
		"Andre Russell",
		"Dwayne Bravo",
		"Sulieman Benn",
		"Devendra Bishoo",
		"Nikita Miller",
		"Ravi Rampaul",
		"Kemar Roach",
	}
	twoT11WC["ireland"] = []string{
		"William Porterfield",
		"Ed Joyce",
		"Niall O'Brien",
		"Gary Wilson",
		"Andre Botha",
		"George Dockrell",
		"Nigel Jones",
		"John Mooney",
		"Kevin O'Brien",
		"Paul Stirling",
		"Andrew White",
		"Alex Cusack",
		"Trent Johnston",
		"Boyd Rankin",
		"Albert van der Merwe",
	}
	twoT11WC["netherland"] = []string{
		"Wesley Barresi",
		"Atse Buurman",
		"Tom Cooper",
		"Tom de Grooth",
		"Alexei Kervezee",
		"Eric Szwarczynski",
		"Bas Zuiderent",
		"Peter Borren",
		"Bernard Loots",
		"Mudassar Bukhari",
		"Ryan ten Doeschate",
		"Adeel Raja",
		"Bradley Kruger",
		"Berend Westdijk",
	}
	twoT11WC["zimbabwe"] = []string{

		"Regis Chakabva",
		"Charles Coventry",
		"Terry Duffin",
		"Craig Ervine",
		"Vusi Sibanda",
		"Tatenda Taibu",
		"Brendan Taylor",
		"Elton Chigumbura",
		"Greg Lamb",
		"Sean Ervine",
		"Sean Williams",
		"Graeme Cremer",
		"Shingi Masakadza",
		"Chris Mpofu",
		"Tinashe Panyangara",
		"Ray Price",
		"Prosper Utseya",
		"Ed Rainsford",
	}
	twoT11WC["cananda"] = []string{
		"Ashish Bagai",
		"Rizwan Cheema",
		"Tyson Gordon",
		"Ruvindu Gunasekera",
		"Nitish Kumar",
		"Hiral Patel",
		"Zubin Surkari",
		"Karl Whatham",
		"Harvir Baidwan",
		"John Davison",
		"Jimmy Hansra",
		"Balaji Rao",
		"Parth Desai",
		"Khurram Chohan",
		"Henry Osinde",
	}
	twoT11WC["kenya"] = []string{

		"Tanmay Mishra",
		"Alex Obanda",
		"David Obuya",
		"Morris Ouma",
		"Rakep Patel",
		"Seren Waters",
		"Collins Obuya",
		"Thomas Odoyo",
		"Steve Tikolo",
		"Jimmy Kamande",
		"James Ngoche",
		"Shem Ngoche",
		"Nehemiah Odhiambo",
		"Peter Ongondo",
		"Elijah Otieno",
	}
	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2011Key, SportKey, CricketKey, twoT11WC, true, false, false)
	return ICCTEWC
}
func CricketWorldCup2012() dataset.SearchSystem_ {
	twoT12WC := map[string][]string{}

	twoT12WC["afghanistan"] = []string{
		"Nawroz Mangal",
		"Asghar Afghan",
		"Javed Ahmadi",
		"Karim Sadiq",
		"Mohammad Shahzad",
		"Najibullah Zadran",
		"Shafiqullah",
		"Gulbadin Naib",
		"Mohammad Nabi",
		"Samiullah Shinwari",
		"Dawlat Zadran",
		"Hamid Hassan",
		"Izatullah Dawlatzai",
		"Mohammad Nasim Baras",
		"Shapoor Zadran",
	}

	twoT12WC["austrillia"] = []string{
		"George Bailey",
		"Michael Hussey",
		"Matthew Wade",
		"David Warner",
		"Cameron White",
		"Shane Watson",
		"Dan Christian",
		"Brad Hogg",
		"David Hussey",
		"Glenn Maxwell",
		"Pat Cummins",
		"Xavier Doherty",
		"Ben Hilfenhaus",
		"Clint McKay",
		"Mitchell Starc",
	}

	twoT12WC["bangladesh"] = []string{
		"Mushfiqur Rahim",
		"Jahurul Islam",
		"Zunaed Siddique",
		"Mohammad Ashraful",
		"Tamim Iqbal",
		"Mahmudullah",
		"Farhad Reza",
		"Nasir Hossain",
		"Shakib Al Hasan",
		"Ziaur Rahman",
		"Abdur Razzak",
		"Abul Hasan",
		"Mashrafe Mortaza",
		"Shafiul Islam",
	}

	twoT12WC["england"] = []string{
		"Jonny Bairstow",
		"Jos Buttler",
		"Alex Hales",
		"Craig Kieswetter",
		"Michael Lumb",
		"Eoin Morgan",
		"Luke Wright",
		"Ravi Bopara",
		"Tim Bresnan",
		"Samit Patel",
		"Stuart Broad",
		"Danny Briggs",
		"Jade Dernbach",
		"Steven Finn",
		"Graeme Swann",
	}

	twoT12WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Gautam Gambhir",
		"Virat Kohli",
		"Suresh Raina",
		"Virender Sehwag",
		"Rohit Sharma",
		"Manoj Tiwary",
		"Yuvraj Singh",
		"Ravichandran Ashwin",
		"Piyush Chawla",
		"Irfan Pathan",
		"Lakshmipathy Balaji",
		"Ashok Dinda",
		"Harbhajan Singh",
		"Zaheer Khan",
	}

	twoT12WC["ireland"] = []string{
		"William Porterfield",
		"Ed Joyce",
		"Niall O'Brien",
		"Gary Wilson",
		"George Dockrell",
		"Nigel Jones",
		"Kevin O'Brien",
		"Paul Stirling",
		"Andrew White",
		"Alex Cusack",
		"Trent Johnston",
		"Tim Murtagh",
		"Boyd Rankin",
		"Max Sorensen",
		"Stuart Thompson",
	}

	twoT12WC["new zeland"] = []string{
		"Ross Taylor",
		"Martin Guptill",
		"Brendon McCullum",
		"Rob Nicol",
		"BJ Watling",
		"Kane Williamson",
		"James Franklin",
		"Jacob Oram",
		"Daniel Vettori",
		"Doug Bracewell",
		"Ronnie Hira",
		"Nathan McCullum",
		"Kyle Mills",
		"Adam Milne",
		"Tim Southee",
	}

	twoT12WC["pakistan"] = []string{
		"Asad Shafiq",
		"Imran Nazir",
		"Kamran Akmal",
		"Nasir Jamshed",
		"Umar Akmal",
		"Mohammad Hafeez",
		"Abdul Razzaq",
		"Shahid Afridi",
		"Shoaib Malik",
		"Yasir Arafat",
		"Mohammad Sami",
		"Raza Hasan",
		"Saeed Ajmal",
		"Sohail Tanvir",
		"Umar Gul",
	}

	twoT12WC["south africa"] = []string{
		"AB de Villiers",
		"Hashim Amla",
		"Faf du Plessis",
		"Richard Levi",
		"Farhaan Behardien",
		"Johan Botha",
		"Jean Paul Duminy",
		"Jacques Kallis",
		"Albie Morkel",
		"Justin Ontong",
		"Morne Morkel",
		"Wayne Parnell",
		"Robin Peterson",
		"Dale Steyn",
		"Lonwabo Tsotsobe",
	}

	twoT12WC["sri lanka"] = []string{
		"Mahela Jayawardene",
		"Dinesh Chandimal",
		"Dilshan Munaweera",
		"Kumar Sangakkara",
		"Lahiru Thirimanne",
		"Angelo Mathews",
		"Akila Dananjaya",
		"Tillakaratne Dilshan",
		"Jeevan Mendis",
		"Thisara Perera",
		"Shaminda Eranga",
		"Rangana Herath",
		"Nuwan Kulasekara",
		"Lasith Malinga",
		"Ajantha Mendis",
	}

	twoT12WC["netherland"] = []string{
		"Wesley Barresi",
		"Atse Buurman",
		"Tom Cooper",
		"Tom de Grooth",
		"Alexei Kervezee",
		"Eric Szwarczynski",
		"Bas Zuiderent",
		"Peter Borren",
		"Bernard Loots",
		"Mudassar Bukhari",
		"Ryan ten Doeschate",
		"Adeel Raja",
		"Bradley Kruger",
		"Berend Westdijk",
	}

	twoT12WC["west indies"] = []string{
		"Darren Bravo",
		"Johnson Charles",
		"Denesh Ramdin",
		"Marlon Samuels",
		"Lendl Simmons",
		"Daren Sammy",
		"Dwayne Bravo",
		"Chris Gayle",
		"Sunil Narine",
		"Kieron Pollard",
		"Andre Russell",
		"Dwayne Smith",
		"Samuel Badree",
		"Fidel Edwards",
		"Ravi Rampaul",
	}

	twoT12WC["zimbabwe"] = []string{
		"Brendan Taylor",
		"Craig Ervine",
		"Hamilton Masakadza",
		"Stuart Matsikenyeri",
		"Forster Mutizwa",
		"Vusi Sibanda",
		"Malcolm Waller",
		"Elton Chigumbura",
		"Graeme Cremer",
		"Kyle Jarvis",
		"Chris Mpofu",
		"Ray Price",
		"Prosper Utseya",
		"Brian Vitori",
		"Richard Muzhange",
	}

	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2012Key, SportKey, CricketKey, twoT12WC, true, false, false)
	return ICCTEWC
}
func CricketChampionsTrophy2013() dataset.SearchSystem_ {
	twoT13CT := map[string][]string{}

	twoT13CT["austrillia"] = []string{
		"Michael Clarke",
		"George Bailey",
		"Phillip Hughes",
		"Adam Voges",
		"Matthew Wade",
		"David Warner",
		"Nathan Coulter Nile",
		"James Faulkner",
		"Mitchell Marsh",
		"Glenn Maxwell",
		"Shane Watson",
		"Xavier Doherty",
		"Mitchell Johnson",
		"Clint McKay",
		"Mitchell Starc",
	}

	twoT13CT["england"] = []string{
		"Alastair Cook",
		"Jonny Bairstow",
		"Ian Bell",
		"Jos Buttler",
		"Eoin Morgan",
		"Joe Root",
		"Jonathan Trott",
		"Ravi Bopara",
		"Tim Bresnan",
		"James Tredwell",
		"Chris Woakes",
		"James Anderson",
		"Stuart Broad",
		"Steven Finn",
		"Graeme Swann",
	}

	twoT13CT["india"] = []string{
		"Mahendra Singh Dhoni",
		"Shikhar Dhawan",
		"Dinesh Karthik",
		"Virat Kohli",
		"Suresh Raina",
		"Rohit Sharma",
		"Murali Vijay",
		"Ravichandran Ashwin",
		"Ravindra Jadeja",
		"Irfan Pathan",
		"Bhuvneshwar Kumar",
		"Amit Mishra",
		"Ishant Sharma",
		"Vinay Kumar",
		"Umesh Yadav",
	}

	twoT13CT["new zeland"] = []string{
		"Brendon McCullum",
		"Martin Guptill",
		"Colin Munro",
		"Luke Ronchi",
		"Ross Taylor",
		"Kane Williamson",
		"Corey Anderson",
		"James Franklin",
		"Daniel Vettori",
		"Grant Elliott",
		"Andrew Ellis",
		"Doug Bracewell",
		"Ian Butler",
		"Mitchell McClenaghan",
		"Nathan McCullum",
		"Kyle Mills",
		"Tim Southee",
		"Trent Boult",
	}

	twoT13CT["pakistan"] = []string{
		"Misbah ul Haq",
		"Asad Shafiq",
		"Imran Farhat",
		"Kamran Akmal",
		"Nasir Jamshed",
		"Umar Amin",
		"Mohammad Hafeez",
		"Shoaib Malik",
		"Abdur Rehman",
		"Asad Ali",
		"Ehsan Adil",
		"Junaid Khan",
		"Mohammad Irfan",
		"Saeed Ajmal",
		"Wahab Riaz",
	}

	twoT13CT["south africa"] = []string{
		"AB de Villiers",
		"Hashim Amla",
		"Faf du Plessis",
		"Colin Ingram",
		"David Miller",
		"Alviro Petersen",
		"Graeme Smith",
		"Farhaan Behardien",
		"Jean Paul Duminy",
		"Ryan McLaren",
		"Chris Morris",
		"Rory Kleinveldt",
		"Robin Peterson",
		"Aaron Phangiso",
		"Dale Steyn",
		"Lonwabo Tsotsobe",
		"Morne Morkel",
	}

	twoT13CT["sri lanka"] = []string{
		"Dinesh Chandimal",
		"Mahela Jayawardene",
		"Kusal Perera",
		"Kumar Sangakkara",
		"Lahiru Thirimanne",
		"Angelo Mathews",
		"Tillakaratne Dilshan",
		"Jeevan Mendis",
		"Thisara Perera",
		"Shaminda Eranga",
		"Rangana Herath",
		"Nuwan Kulasekara",
		"Lasith Malinga",
		"Sachithra Senanayake",
		"Chanaka Welegedara",
		"Dilhara Lokuhettige",
	}

	twoT13CT["west indies"] = []string{
		"Denesh Ramdin",
		"Darren Bravo",
		"Johnson Charles",
		"Marlon Samuels",
		"Ramnaresh Sarwan",
		"Devon Smith",
		"Dwayne Bravo",
		"Chris Gayle",
		"Jason Holder",
		"Sunil Narine",
		"Kieron Pollard",
		"Daren Sammy",
		"Tino Best",
		"Ravi Rampaul",
		"Kemar Roach",
	}

	ICCTECT := dataset.SearchSystem_{}
	ICCTECT.Constructor() // important
	ICCTECT.Include(CricketChampionsWorldCup2013Key, SportKey, CricketKey, twoT13CT, true, false, false)
	return ICCTECT
}
func CricketWorldCup2015() dataset.SearchSystem_ {
	twoT15WC := map[string][]string{}

	twoT15WC["afghanistan"] = []string{
		"Afsar Zazai",
		"Asghar Afghan",
		"Javed Ahmadi",
		"Najibullah Zadran",
		"Nasir Jamal",
		"Nawroz Mangal",
		"Shafiqullah",
		"Usman Ghani",
		"Mohammad Nabi",
		"Gulbadin Naib",
		"Samiullah Shinwari",
		"Aftab Alam",
		"Dawlat Zadran",
		"Hamid Hassan",
		"Shapoor Zadran",
		"Mirwais Ashraf",
	}

	twoT15WC["austrillia"] = []string{
		"Michael Clarke",
		"George Bailey",
		"Aaron Finch",
		"Brad Haddin",
		"Steven Smith",
		"David Warner",
		"James Faulkner",
		"Mitchell Marsh",
		"Glenn Maxwell",
		"Shane Watson",
		"Pat Cummins",
		"Xavier Doherty",
		"Josh Hazlewood",
		"Mitchell Johnson",
		"Mitchell Starc",
	}

	twoT15WC["bangladesh"] = []string{
		"Anamul Haque",
		"Imrul Kayes",
		"Mushfiqur Rahim",
		"Tamim Iqbal",
		"Mahmudullah",
		"Mominul Haque",
		"Nasir Hossain",
		"Sabbir Rahman",
		"Shakib Al Hasan",
		"Soumya Sarkar",
		"Taijul Islam",
		"Mashrafe Mortaza",
		"Arafat Sunny",
		"Rubel Hossain",
		"Taskin Ahmed",
		"Al Amin Hossain",
	}

	twoT15WC["england"] = []string{
		"Eoin Morgan",
		"Jos Buttler",
		"Gary Ballance",
		"Ian Bell",
		"Alex Hales",
		"Joe Root",
		"James Taylor",
		"Moeen Ali",
		"Ravi Bopara",
		"James Tredwell",
		"Chris Woakes",
		"James Anderson",
		"Stuart Broad",
		"Steven Finn",
		"Chris Jordan",
	}

	twoT15WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Shikhar Dhawan",
		"Virat Kohli",
		"Ajinkya Rahane",
		"Suresh Raina",
		"Ambati Rayudu",
		"Rohit Sharma",
		"Ravichandran Ashwin",
		"Stuart Binny",
		"Ravindra Jadeja",
		"Axar Patel",
		"Bhuvneshwar Kumar",
		"Mohammed Shami",
		"Mohit Sharma",
		"Umesh Yadav",
		"Ishant Sharma",
	}

	twoT15WC["ireland"] = []string{
		"William Porterfield",
		"Andy Balbirnie",
		"Ed Joyce",
		"Niall O'Brien",
		"Gary Wilson",
		"George Dockrell",
		"Andy McBrine",
		"John Mooney",
		"Kevin O'Brien",
		"Paul Stirling",
		"Peter Chase",
		"Alex Cusack",
		"Max Sorensen",
		"Stuart Thompson",
		"Craig Young",
		"Tim Murtagh",
	}

	twoT15WC["new zeland"] = []string{
		"Brendon McCullum",
		"Martin Guptill",
		"Tom Latham",
		"Luke Ronchi",
		"Ross Taylor",
		"Kane Williamson",
		"Corey Anderson",
		"Grant Elliott",
		"Daniel Vettori",
		"Trent Boult",
		"Matt Henry",
		"Mitchell McClenaghan",
		"Nathan McCullum",
		"Kyle Mills",
		"Tim Southee",
		"Adam Milne",
	}

	twoT15WC["pakistan"] = []string{
		"Misbah ul Haq",
		"Ahmed Shehzad",
		"Haris Sohail",
		"Nasir Jamshed",
		"Sarfaraz Ahmed",
		"Sohaib Maqsood",
		"Umar Akmal",
		"Younis Khan",
		"Shahid Afridi",
		"Mohammad Hafeez",
		"Ehsan Adil",
		"Mohammad Irfan",
		"Rahat Ali",
		"Sohail Khan",
		"Wahab Riaz",
		"Yasir Shah",
		"Junaid Khan",
	}

	twoT15WC["scotland"] = []string{
		"Preston Mommsen",
		"Kyle Coetzer",
		"Richie Berrington",
		"Freddie Coleman",
		"Matthew Cross",
		"Hamish Gardiner",
		"Matt Machan",
		"Calum MacLeod",
		"Michael Leask",
		"Rob Taylor",
		"Josh Davey",
		"Alasdair Evans",
		"Majid Haq",
		"Safyaan Sharif",
		"Iain Wardlaw",
	}

	twoT15WC["south africa"] = []string{
		"AB de Villiers",
		"Hashim Amla",
		"Quinton de Kock",
		"Faf du Plessis",
		"David Miller",
		"Rilee Rossouw",
		"Farhaan Behardien",
		"Jean Paul Duminy",
		"Vernon Philander",
		"Kyle Abbott",
		"Imran Tahir",
		"Morne Morkel",
		"Wayne Parnell",
		"Aaron Phangiso",
		"Dale Steyn",
	}

	twoT15WC["sri lanka"] = []string{
		"Lahiru Thirimanne",
		"Mahela Jayawardene",
		"Kusal Perera",
		"Kumar Sangakkara",
		"Upul Tharanga",
		"Dinesh Chandimal",
		"Dimuth Karunaratne",
		"Angelo Mathews",
		"Tillakaratne Dilshan",
		"Thisara Perera",
		"Seekkuge Prasanna",
		"Jeevan Mendis",
		"Dushmantha Chameera",
		"Tharindu Kaushal",
		"Nuwan Kulasekara",
		"Suranga Lakmal",
		"Lasith Malinga",
		"Sachithra Senanayake",
		"Rangana Herath",
		"Dhammika Prasad",
	}

	twoT15WC["united arab emirates"] = []string{
		"Amjad Ali",
		"Andri Berenger",
		"Khurram Khan",
		"Swapnil Patil",
		"Rohan Mustafa",
		"Saqlain Haider",
		"Shaiman Anwar",
		"Amjad Javed",
		"Krishna Chandran",
		"Mohammad Tauqir",
		"Fahad Alhashmi",
		"Manjula Guruge",
		"Kamran Shazad",
		"Mohammad Naveed",
		"Nasir Aziz",
	}

	twoT15WC["west indies"] = []string{
		"Marlon Samuels",
		"Johnson Charles",
		"Denesh Ramdin",
		"Lendl Simmons",
		"Darren Bravo",
		"Jason Holder",
		"Jonathan Carter",
		"Chris Gayle",
		"Andre Russell",
		"Daren Sammy",
		"Dwayne Smith",
		"Sunil Narine",
		"Sulieman Benn",
		"Sheldon Cottrell",
		"Nikita Miller",
		"Kemar Roach",
		"Jerome Taylor",
	}
	twoT15WC["zimbabwe"] = []string{
		"Regis Chakabva",
		"Chamu Chibhabha",
		"Craig Ervine",
		"Hamilton Masakadza",
		"Stuart Matsikenyeri",
		"Solomon Mire",
		"Brendan Taylor",
		"Elton Chigumbura",
		"Tafadzwa Kamungozi",
		"Sikandar Raza",
		"Sean Williams",
		"Tendai Chatara",
		"Tawanda Mupariwa",
		"Tinashe Panyangara",
		"Prosper Utseya",
	}

	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2015Key, SportKey, CricketKey, twoT15WC, true, false, false)
	return ICCTEWC
}
func CricketWorldCup2016() dataset.SearchSystem_ {
	twoT16WC := map[string][]string{}

	twoT16WC["afghanistan"] = []string{
		"Asghar Afghan",
		"Karim Sadiq",
		"Mohammad Shahzad",
		"Najibullah Zadran",
		"Noor Ali Zadran",
		"Shafiqullah",
		"Usman Ghani",
		"Amir Hamza",
		"Gulbadin Naib",
		"Mohammad Nabi",
		"Rashid Khan",
		"Samiullah Shinwari",
		"Dawlat Zadran",
		"Hamid Hassan",
		"Shapoor Zadran",
	}

	twoT16WC["austrillia"] = []string{
		"Steven Smith",
		"David Warner",
		"Aaron Finch",
		"Usman Khawaja",
		"Peter Nevill",
		"Nathan Coulter Nile",
		"James Faulkner",
		"John Hastings",
		"Mitchell Marsh",
		"Glenn Maxwell",
		"Shane Watson",
		"Ashton Agar",
		"Josh Hazlewood",
		"Andrew Tye",
		"Adam Zampa",
	}

	twoT16WC["bangladesh"] = []string{
		"Mohammad Mithun",
		"Mushfiqur Rahim",
		"Nurul Hasan",
		"Tamim Iqbal",
		"Shakib Al Hasan",
		"Abu Hider",
		"Mahmudullah",
		"Nasir Hossain",
		"Sabbir Rahman",
		"Shuvagata Hom",
		"Soumya Sarkar",
		"Mashrafe Mortaza",
		"Al Amin Hossain",
		"Mustafizur Rahman",
		"Saqlain Sajib",
		"Arafat Sunny",
		"Taskin Ahmed",
	}

	twoT16WC["england"] = []string{
		"Eoin Morgan",
		"Sam Billings",
		"Jos Buttler",
		"Alex Hales",
		"Joe Root",
		"Jason Roy",
		"James Vince",
		"Moeen Ali",
		"Liam Dawson",
		"Ben Stokes",
		"David Willey",
		"Chris Jordan",
		"Liam Plunkett",
		"Adil Rashid",
		"Reece Topley",
		"Steven Finn",
	}
	twoT16WC["hong kong"] = []string{
		"Waqas Khan",
		"Anshy Rath",
		"Jamie Atkinson",
		"Babar Hayat",
		"Ryan Campbell",
		"Christopher Carter",
		"Kinchit Shah",
		"Waqas Barkat",
		"Tanwir Afzal",
		"Aizaz Khan",
		"Mark Chapman",
		"Nizakat Khan",
		"Haseeb Amjad",
		"Adil Mehmood",
		"Nadeem Ahmed",
	}
	twoT16WC["india"] = []string{
		"Mahendra Singh Dhoni",
		"Shikhar Dhawan",
		"Virat Kohli",
		"Manish Pandey",
		"Ajinkya Rahane",
		"Suresh Raina",
		"Rohit Sharma",
		"Yuvraj Singh",
		"Ravichandran Ashwin",
		"Ravindra Jadeja",
		"Hardik Pandya",
		"Jasprit Bumrah",
		"Harbhajan Singh",
		"Mohammed Shami",
		"Pawan Negi",
		"Ashish Nehra",
	}

	twoT16WC["ireland"] = []string{
		"William Porterfield",
		"Andy Balbirnie",
		"Niall O'Brien",
		"Stuart Poynter",
		"Andrew Poynter",
		"Gary Wilson",
		"George Dockrell",
		"Andy McBrine",
		"Kevin O'Brien",
		"Paul Stirling",
		"Tim Murtagh",
		"Boyd Rankin",
		"Max Sorensen",
		"Stuart Thompson",
		"Craig Young",
	}

	twoT16WC["netherland"] = []string{
		"Wesley Barresi",
		"Tom Cooper",
		"Ben Cooper",
		"Stephan Myburgh",
		"Max O'Dowd",
		"Peter Borren",
		"Mudassar Bukhari",
		"Michael Rippon",
		"Pieter Seelaar",
		"Roelof van der Merwe",
		"Ahsan Malik",
		"Vivian Kingma",
		"Logan van Beek",
		"Timm van der Gugten",
		"Paul van Meekeren",
	}
	twoT16WC["new zeland"] = []string{
		"Kane Williamson",
		"Martin Guptill",
		"Colin Munro",
		"Henry Nicholls",
		"Luke Ronchi",
		"Ross Taylor",
		"Corey Anderson",
		"Grant Elliott",
		"Mitchell Santner",
		"Trent Boult",
		"Mitchell McClenaghan",
		"Nathan McCullum",
		"Adam Milne",
		"Ish Sodhi",
		"Tim Southee",
	}
	twoT16WC["oman"] = []string{
		"Sultan Ahmed",
		"Aamir Kaleem",
		"Amir Ali",
		"Jatinder Singh",
		"Khawar Ali",
		"Mehran Khan",
		"Vaibhav Wategaonkar",
		"Zeeshan Maqsood",
		"Zeeshan Siddiqui",
		"Ajay Lalcheta",
		"Munis Ansari",
		"Bilal Khan",
		"Rajeshkumar Ranpura",
		"Sufyan Mehmood",
		"Adnan Ilyas",
	}
	twoT16WC["pakistan"] = []string{
		"Ahmed Shehzad",
		"Khalid Latif",
		"Sarfaraz Ahmed",
		"Sharjeel Khan",
		"Umar Akmal",
		"Babar Azam",
		"Iftikhar Ahmed",
		"Khurram Manzoor",
		"Shahid Afridi",
		"Imad Wasim",
		"Mohammad Hafeez",
		"Mohammad Nawaz",
		"Shoaib Malik",
		"Anwar Ali",
		"Mohammad Amir",
		"Mohammad Irfan",
		"Mohammad Sami",
		"Wahab Riaz",
		"Rumman Raees",
	}

	twoT16WC["scotland"] = []string{
		"Preston Mommsen",
		"Kyle Coetzer",
		"Richie Berrington",
		"Matthew Cross",
		"Matt Machan",
		"Calum MacLeod",
		"George Munsey",
		"Con de Lange",
		"Michael Leask",
		"Rob Taylor",
		"Josh Davey",
		"Alasdair Evans",
		"Gavin Main",
		"Safyaan Sharif",
		"Mark Watt",
	}

	twoT16WC["south africa"] = []string{
		"Faf du Plessis",
		"Hashim Amla",
		"Quinton de Kock",
		"AB de Villiers",
		"David Miller",
		"Rilee Rossouw",
		"Farhaan Behardien",
		"Jean Paul Duminy",
		"Chris Morris",
		"David Wiese",
		"Kyle Abbott",
		"Imran Tahir",
		"Aaron Phangiso",
		"Kagiso Rabada",
		"Dale Steyn",
	}

	twoT16WC["sri lanka"] = []string{
		"Dinesh Chandimal",
		"Chamara Kapugedera",
		"Lahiru Thirimanne",
		"Niroshan Dickwella",
		"Angelo Mathews",
		"Tillakaratne Dilshan",
		"Shehan Jayasuriya",
		"Thisara Perera",
		"Dasun Shanaka",
		"Milinda Siriwardana",
		"Dushmantha Chameera",
		"Rangana Herath",
		"Nuwan Kulasekara",
		"Suranga Lakmal",
		"Sachithra Senanayake",
		"Jeffrey Vandersay",
		"Lasith Malinga",
	}

	twoT16WC["west indies"] = []string{
		"Johnson Charles",
		"Evin Lewis",
		"Denesh Ramdin",
		"Marlon Samuels",
		"Lendl Simmons",
		"Darren Bravo",
		"Andre Fletcher",
		"Daren Sammy",
		"Carlos Brathwaite",
		"Dwayne Bravo",
		"Chris Gayle",
		"Jason Holder",
		"Andre Russell",
		"Sunil Narine",
		"Kieron Pollard",
		"Samuel Badree",
		"Sulieman Benn",
		"Ashley Nurse",
		"Jerome Taylor",
	}
	twoT16WC["zimbabwe"] = []string{
		"Hamilton Masakadza",
		"Chamu Chibhabha",
		"Peter Moor",
		"Richmond Mutumbami",
		"Vusi Sibanda",
		"Malcolm Waller",
		"Elton Chigumbura",
		"Sikandar Raza",
		"Sean Williams",
		"Luke Jongwe",
		"Neville Madziva",
		"Tendai Chatara",
		"Tendai Chisoro",
		"Wellington Masakadza",
		"Tawanda Mupariwa",
		"Tinashe Panyangara",
		"Donald Tiripano",
		"Graeme Cremer",
	}

	ICCTEWC := dataset.SearchSystem_{}
	ICCTEWC.Constructor() // important
	ICCTEWC.Include(CricketWorldCup2016Key, SportKey, CricketKey, twoT16WC, true, false, false)
	return ICCTEWC
}

// end of international

// BuildCricket creates a go to ready parcel to the server
func BuildCricket() CricketParcel {
	ICC2007, ICC2009, ICCc2009, ICC2011, ICC2012, ICCc2013, ICC2015, ICC2016 := CricketWorldCup2007(), CricketWorldCup2009(),
		CricketChampionsTrophy2009(), CricketWorldCup2011(), CricketWorldCup2012(), CricketChampionsTrophy2013(), CricketWorldCup2015(), CricketWorldCup2016()

	events7, category7, field7, book7, items7 := ICC2007.Manager(CricketWorldCup2007Key)
	events9, category9, field9, book9, items9 := ICC2009.Manager(CricketWorldCup2009Key)
	events09, category09, field09, book09, items09 := ICCc2009.Manager(CricketChampionsWorldCup2009Key)
	events11, category11, field11, book11, items11 := ICC2011.Manager(CricketWorldCup2011Key)
	events12, category12, field12, book12, items12 := ICC2012.Manager(CricketWorldCup2012Key)
	events13, category13, field13, book13, items13 := ICCc2013.Manager(CricketChampionsWorldCup2013Key)
	events15, category15, field15, book15, items15 := ICC2015.Manager(CricketWorldCup2015Key)
	events16, category16, field16, book16, items16 := ICC2016.Manager(CricketWorldCup2016Key)

	parcel7 := SportsParcel{Event: events7, Category: category7, Field: field7,
		Book: book7, MappedItems: items7}
	parcel9 := SportsParcel{Event: events9, Category: category9, Field: field9,
		Book: book9, MappedItems: items9}
	parcel09 := SportsParcel{Event: events09, Category: category09, Field: field09,
		Book: book09, MappedItems: items09}

	parcel11 := SportsParcel{Event: events11, Category: category11, Field: field11,
		Book: book11, MappedItems: items11}
	parcel12 := SportsParcel{Event: events12, Category: category12, Field: field12,
		Book: book12, MappedItems: items12}
	parcel13 := SportsParcel{Event: events13, Category: category13, Field: field13,
		Book: book13, MappedItems: items13}
	parcel15 := SportsParcel{Event: events15, Category: category15, Field: field15,
		Book: book15, MappedItems: items15}
	parcel16 := SportsParcel{Event: events16, Category: category16, Field: field16,
		Book: book16, MappedItems: items16}
	ready := []SportsParcel{parcel7, parcel9, parcel09, parcel11, parcel12, parcel13, parcel15, parcel16}
	pack := CricketParcel{Pack: ready}
	return pack
}

func CricketEvents() []string {
	events := []string{}
	openBuild := BuildCricket()
	for _, event := range openBuild.Pack {
		events = append(events, event.Event)
	}
	return events
}

func CricketCategory() string {
	category := []string{}
	token := BuildCricket().Pack
	for _, cat := range token {
		category = append(category, cat.Category)
	}
	category = dataset.EraseDuplicate(category)
	cat := category[0]
	fmt.Println(category)
	return cat
}

func CricketBook() string {
	fi := ""
	token := BuildCricket().Pack
	for _, cat := range token {
		fi = cat.Book
	}
	return fi
}

func CricketLists() []string {
	ma := []string{}
	token := BuildCricket().Pack
	for _, t := range token {
		for items := range t.MappedItems {
			ma = append(ma, items)
		}
	}
	return ma
}

func CricketSheet() map[string][]string {
	sheet := map[string][]string{}
	token := BuildCricket().Pack
	for _, t := range token {
		maps.Copy(sheet, t.MappedItems)
	}
	return sheet
}
