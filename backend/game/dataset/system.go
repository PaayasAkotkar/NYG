package dataset

import (
	"fmt"
	"maps"
	"sort"
	"strings"
)

// SearchSystem_  this is the parent where key=[]string
type SearchSystem_ struct {

	// domains, main-categories, sub-categories, fields, items, sub-items
	// note: items and list-of-items can be called as mapped items
	set map[string]map[string]map[string]map[string]map[string]map[string][]string
	// example:
	// domain: sports
	// main-category: cricket
	// sub-category: international
	// field: icc
	// item: india
	// sub-items: kohli, sharma, bumrah
	defaults map[string]bool
}

// Constructor Default-Constructor
func (ss SearchSystem_) Constructor() SearchSystem_ {
	Domains := []string{"sports", "entertainment"}
	MainCategories := []string{"cricket", "movies"} // for sports its cricket
	SubCategories := []string{"international", "national", "domestic"}
	Fields := []string{"icc"}
	// for cricket international players
	MapItems := map[string][]string{
		"india": {"rohit sharma", "virat kohli", "jasprit bumrah", "ravindra jadeja", "mohammad shami",
			"mohammed siraj", "kannur lokesh rahul",
			"shubman gill", "hardik pandya", "suryakumar yadav", "rishabh pant", "kuldeep yadav", "axar patel", "yashasvi jaiswal",
			"rinku singh", "tilak verma", "ruturaj gaikwad", "shardul thakur", "shivam dube", "ravi bishnoi", "jitesh sharma",
			"washington sundar", "mukesh kumar", "sanju samson", "arshdeep singh", "kona srikar bharat", "prasidh kirshna", "avesh khan",
			"rajat patidar", "sarfaraz khan", "dhruv jurel"},

		"austrillia": {"sean abbot", "xavier bartlett", "scott boland", "alex carey",
			"pat cummins", "nathan ellis", "cameron green", "aaron hardie",
			"josh hazlewood", "travis head", "josh inglis",
			"usman khawaja", "marnus labuschagne", "nathan lyon", "mitchell marsh",
			"glenn maxwell", "lance morris", "todd murphy", "jhye richardson",
			"matt short", "steve smith", "mitchell starc", "adam zampa"},

		"england": {
			"Gus Atkinson", "Harry Brook", "Jos Buttler", "Joe Root", "Jamie Smith", "Ben Stokes", "Mark Wood",
			"Rehan Ahmed", "Jofra Archer", "Jonny Bairstow", "Shoaib Bashir", "Brydon Carse", "Zak Crawley",
			"Sam Curran", "Ben Duckett", "Will Jacks", "Jack Leach", "Liam Livingstone", "Ollie Pope", "Matthew Potts",
			"Adil Rashid", "Phil Salt", "Olly Stone", "Josh Tongue", "Reece Topley",
			"Chris Woakes", "Jacob Bethell", "Josh Hull", "John Turner"},

		// to do: make sure the input for the names like new zeland south africa
		// to be converted into newzeland southafrica
		// note: player having three names or four to be converted into two names making it first name and last name only
		"newzeland": {
			"Tom Blundell", "Michael Bracewell", "Mark Chapman", "Josh Clarkson", "Jacob Duffy",
			"Matt Henry", "Kyle Jamieson", "Tom Latham", "Daryl Mitchell", "Henry Nicholls", "Will O’Rourke",
			"Ajaz Patel", "Glenn Phillips", "Rachin Ravindra", "Mitchell Santner",
			"Ben Sears", "Nathan Smith", "Ish Sodhi", "Tim Southee", "Will Young",
		},

		"southafrica": {
			"Temba Bavuma", "David Bedingham", "Nandre Burger", "Gerald Coetzee",
			"Tony de Zorzi", "Reeza Hendricks", "Marco Jansen", "Keshav Maharaj",
			"Kwena Maphaka", "Aiden Markram", "Wiaan Mulder", "Senuran Muthusamy", "Lungi Ngidi",
			"Kagiso Rabada", "Ryan Rickelton", "Tristan Stubbs", "Kyle Verreynne",
			"Lizaad Williams", "David Miller", "Rassie van der Dussen",
		},
		"westindies": {
			"Alick Athanaze", "Kraigg Brathwaite", "Keacy Carty", "Tagenarine Chanderpaul", "Joshua Da Silva", "Jason Holder",
			"Shai Hope", "Akeal Hosein", "Alzarri Joseph", "Brandon King", "Kyle Mayers",
			"Gudakesh Motie", "Nicholas Pooran", "Rovman Powell", "Kemar Roach", "Jayden Seales", "Romario Shepherd",
		},
		// note this not the acutal 2025 squad tho
		"srilanka": {
			"Charith Asalanka", "Pathum Nissanka", "Avishka Fernando",
			"Kusal Mendis", "Kamindu Mendis", "Janith Liyanage", "Nishan Madushka",
			"Nuwanidu Fernando", "Wanindu Hasaranga", "Maheesh Theekshana", "Dunith Wellalage", "Jeffrey Vandersay", "Asitha Fernando", "Lahiru Kumara", "Mohamed Shiraz", "Eshan Malinga",
		},
	}
	d := map[string]bool{}
	d["2025-international-cricket-team"] = true
	set := map[string]map[string]map[string]map[string]map[string]map[string][]string{}
	set["2025-international-cricket-team"] = map[string]map[string]map[string]map[string]map[string][]string{
		Domains[0]: {
			MainCategories[0]: {
				SubCategories[0]: {
					Fields[0]: MapItems,
				},
			},
		},
	}

	return SearchSystem_{set: set, defaults: d}
}

// See returns the default value
func (ss SearchSystem_) See() map[string]map[string]map[string]map[string]map[string]map[string][]string {
	_copy := make(map[string]map[string]map[string]map[string]map[string]map[string][]string)
	maps.Copy(_copy, ss.set)
	return _copy
}

// ADomains returns all the domains in the dataset
func (ss SearchSystem_) ADomains() []string {
	doms := make([]string, 0, len(ss.set))
	for r := range ss.set {
		doms = append(doms, r)
	}
	return doms
}

// AMainCategories returns all the main categories in the dataset
func (ss SearchSystem_) AMainCategories(_sort bool) []string {
	MC := []string{}
	for r := range ss.set {
		for r2 := range ss.set[r] {
			for r3 := range ss.set[r][r2] {
				MC = append(MC, r3)
			}
		}
	}
	if _sort {
		sort.Strings(MC)
	}
	return MC
}

// ASubCategories returns all the sub categories in the dataset
func (ss SearchSystem_) ASubCategories(_sort bool) []string {
	SC := []string{}
	set := ss.set
	for r := range set {
		for r2 := range set[r] {
			for r3 := range set[r][r2] {
				for r4 := range set[r][r2][r3] {
					SC = append(SC, r4)
				}
			}
		}
	}
	if _sort {
		sort.Strings(SC)
	}
	return SC
}

// AFileds returns all the fields in the dataset
func (ss SearchSystem_) AFileds(_sort bool) []string {
	F := []string{}
	set := ss.set
	for r := range set {
		for r2 := range set[r] {
			for r3 := range set[r][r2] {
				for r4 := range set[r][r2][r3] {
					for r5 := range set[r][r2][r3][r4] {
						F = append(F, r5)
					}
				}

			}
		}
	}
	if _sort {
		sort.Strings(F)
	}
	return F
}

// AItems returns all the items in the dataset
func (ss SearchSystem_) AItems(_sort bool) []string {
	I := []string{}
	set := ss.set
	for r := range set {
		for r2 := range set[r] {
			for r3 := range set[r][r2] {
				for r4 := range set[r][r2][r3] {
					for r5 := range set[r][r2][r3][r4] {
						for r6 := range set[r][r2][r3][r4][r5] {
							fmt.Println(r6)
						}
					}
				}
			}
		}
	}
	if _sort {
		sort.Strings(I)
	}
	return I
}

// ASubItems returns all the sub items in the dataset
func (ss SearchSystem_) ASubItems(_sort bool) []string {
	SI := []string{}
	set := ss.set
	for r := range set {
		for r2 := range set[r] {
			for r3 := range set[r][r2] {
				for r4 := range set[r][r2][r3] {
					for r5 := range set[r][r2][r3][r4] {
						for r6 := range set[r][r2][r3][r4][r5] {
							SI = append(SI, set[r][r2][r3][r4][r5][r6]...)
						}

					}
				}
			}
		}
	}
	if _sort {
		sort.Strings(SI)
	}
	return SI
}

// HasField returns true if the value found in the dataset
// TIP you can even pass the forename, surname, or any other character that is separated via whitespace(for example: apple pie)
// NOTE it searches through the list all fields
func (ss SearchSystem_) HasField(item string) bool {
	item = strings.ToLower(item)
	items := ss.AFileds(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// HasItem returns true if the value found in the dataset
// NOTE it searches through the list all items
func (ss SearchSystem_) HasItem(item string) bool {
	item = strings.ToLower(item)
	items := ss.AItems(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// HasSubCategory returns true if the value found in the dataset
// NOTE it searches through the list all sub category
func (ss SearchSystem_) HasSubCategory(item string) bool {
	item = strings.ToLower(item)
	items := ss.ASubCategories(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// HasMainCategory returns true if the value found in the dataset
// NOTE it searches through the list all main category
func (ss SearchSystem_) HasMainCategory(item string) bool {
	item = strings.ToLower(item)
	items := ss.AMainCategories(false)
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// HasDomain returns true if the value found in the dataset
// NOTE it searches through the list all domains
func (ss SearchSystem_) HasDomain(item string) bool {
	item = strings.ToLower(item)
	items := ss.ADomains()
	found := false
	EqualizeString_(items)
	_name := []string{} // either by first or last

	// separating the first and last name
	for r := range items {
		_name = append(_name, strings.Split(items[r], " ")...)
	}
	for r := range _name {
		switch true {
		case _name[r] == item:
			{
				found = true
				break
			}
		case Includes(items, item):
			{
				found = true
			}
		}
	}
	return found
}

// PosItem returns the position or index of the items from whole collection
// if not found returns -1
func (ss SearchSystem_) PosItem(item string, _in []string) int {
	items := _in
	found := 0
	for !Includes(items, item) {
		found += 1
		if found > len(items) {
			found = -1
			break
		}
	}
	return found
}

func (ss SearchSystem_) mapEtoLower(m map[string][]string) map[string][]string {
	for r, v := range m {
		for _, r2 := range v {
			m[r] = []string{strings.ToLower(r2)}
		}
	}
	return m
}

// Include  returns adds the new key and respective value in it
// NOTE if you have not assigned the constructor than it wont mapped
// example: create:=SearchSystem_{}.Constructor()
func (ss SearchSystem_) Include(key string, domain string, newMainC string, newF string, mapItems map[string][]string, national bool, domestic bool, international bool) {
	for r := range ss.set {
		if key == r {
			panic(key + "cannot override the key")
		}
	}
	// all the elements to lower cases
	newMainC = strings.ToLower(newMainC)
	newF = strings.ToLower(newF)
	mapItems = ss.mapEtoLower(mapItems)

	// to check domain name, mc, fields,
	switch true {
	case ss.HasDomain(domain):
		panic(domain + "cannot override default domain name")
	case ss.HasMainCategory(newMainC):
		panic(newMainC + "cannot override default main cateogry ")
	case ss.HasField(newF):
		panic(newF + "cannot override default field")

	}
	switch true {
	case international:
		ss.set[key] = map[string]map[string]map[string]map[string]map[string][]string{
			domain: {
				newMainC: {
					"international": {
						newF: mapItems,
					},
				},
			},
		}
	case national:
		ss.set[key] = map[string]map[string]map[string]map[string]map[string][]string{
			domain: {
				newMainC: {
					"national": {
						newF: mapItems,
					},
				},
			},
		}
	case domestic:
		ss.set[key] = map[string]map[string]map[string]map[string]map[string][]string{
			domain: {
				newMainC: {
					"domestic": {
						newF: mapItems,
					},
				},
			},
		}
	}
}

// RemoveDomain returns removes the key and associated data with it
func (ss SearchSystem_) RemoveDomain(domain string) {
	s := SearchSystem_{}.Constructor()
	for r := range s.defaults {
		fmt.Println(r, domain)
		if r == domain {
			panic("cannot remove default domain")
		}
	}
	if !ss.HasDomain(domain) {
		panic("there's no domin named " + domain + " exists")
	}
	delete(ss.set, domain)
}

// ReplaceItemInDomain returns replace the item in the domain in the item category
func (ss SearchSystem_) ReplaceItemInDomain(domain string, search string, replace string) {
	s := ss.set
	for r := range s[domain] {
		for r2 := range s[domain][r] {
			for r3 := range s[domain][r][r2] {
				for r4 := range s[domain][r][r2][r3] {
					for r5 := range s[domain][r][r2][r3][r4] {
						for _, r7 := range s[domain][r][r2][r3][r4][r5] {
							if search == r7 {
								Replace(s[domain][r][r2][r3][r4][r5], search, replace)
							}
						}
					}
				}
			}
		}
	}
	ss.set[domain] = s[domain]
}

// DomainsCount returns total number of keys
func (ss SearchSystem_) DomainsCount() int {
	count := len(ss.ADomains())
	return count
}

// MainCategoriesCount returns total number of categories
func (ss SearchSystem_) MainCategoriesCount() int {
	count := len(ss.AMainCategories(false))
	return count
}

// FieldsCount returns total number of items
func (ss SearchSystem_) FieldsCount() int {
	count := len(ss.AFileds(false))
	return count
}

// ItemCount returns total number of items
func (ss SearchSystem_) ItemCount() int {
	count := len(ss.AItems(false))
	return count
}

// SubItemCount returns total number of items
func (ss SearchSystem_) SubItemCount() int {
	count := len(ss.ASubItems(false))
	return count
}

// Shift returns shifts the value in the key or sorts the collection if true
func (ss SearchSystem_) Shift(domain string, inItem string, appenditem string, _sort bool) {
	s := ss.set
	if ss.defaults[domain] {
		panic("cannot shift value in default domain")
	}
	for r := range s[domain] {
		for r2 := range s[domain][r] {
			for r3 := range s[domain][r][r2] {
				for r4 := range s[domain][r][r2][r3] {
					for r5 := range s[domain][r][r2][r3][r4] {
						fmt.Println(r5, inItem)
						if r5 == inItem {
							s[domain][r][r2][r3][r4][r5] = append(s[domain][r][r2][r3][r4][r5], appenditem)
						}
					}
				}
			}
		}
	}
	ss.set[domain] = s[domain]
}

// Match returns true if the element that matches in the given key
func (ss SearchSystem_) Match(domain string, inItem string, search string) bool {
	// equalize the string
	s := ss.set
	search = strings.ToLower(search)
	found := false

	// spearate the names in first and last
	_names := []string{}
	for r := range s[domain] {
		for r2 := range s[domain][r] {
			for r3 := range s[domain][r][r2] {
				for r4 := range s[domain][r][r2][r3] {
					for r5 := range s[domain][r][r2][r3][r4] {
						if r5 == inItem {
							fmt.Println(true)
							_names = append(_names, s[domain][r][r2][r3][r4][r5]...)
						}
					}
				}
			}
		}
	}
	_names = ParseWords(_names)

	for _, r := range _names {
		if r == search {
			found = true
		}
	}
	return found
}

// ParseWords returns separates the words in two category namely first word and last word
func ParseWords(item []string) []string {
	_name := []string{}
	for r := range item {
		_name = append(_name, strings.Split(item[r], " ")...)
	}
	return _name
}
