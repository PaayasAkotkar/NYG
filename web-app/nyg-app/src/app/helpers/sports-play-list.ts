
interface Sports {
    [country: string]: string[],
}

export class SportsData {
    protected Cricket: Sports = {
        ["india"]: ["rohit sharma", "virat kohli", "jasprit bumrah", "ravindra jadeja", "mohammad shami",
            "mohammed siraj", "kannur lokesh rahul",
            "shubman gill", "hardik pandya", "suryakumar yadav", "rishabh pant", "kuldeep yadav", "axar patel", "yashasvi jaiswal",
            "rinku singh", "tilak verma", "ruturaj gaikwad", "shardul thakur", "shivam dube", "ravi bishnoi", "jitesh sharma",
            "washington sundar", "mukesh kumar", "sanju samson", "arshdeep singh", "kona srikar bharat", "prasidh kirshna", "avesh khan",
            "rajat patidar", "sarfaraz khan", "dhruv jurel"],

        ["australia"]: ["sean abbot", "xavier bartlett", "scott boland", "alex carey",
            "pat cummins", "nathan ellis", "cameron green", "aaron hardie",
            "josh hazlewood", "travis head", "josh inglis",
            "usman khawaja", "marnus labuschagne", "nathan lyon", "mitchell marsh",
            "glenn maxwell", "lance morris", "todd murphy", "jhye richardson",
            "matt short", "steve smith", "mitchell starc", "adam zampa"]
    }

    constructor(team_name: string) {
        this.currentTeam = team_name
    }

    /**fields*/

    // you can find details for other teams too
    currentTeam: string = " "

    /**contains detail of the given team */
    myTeam = new class {

        constructor(private sd: SportsData) { }

        /**methods*/

        /**
         * 
         * @returns if the team is the top 10 
         */
        bTop10(): boolean {
            var has = false
            // avoiding complications by lower-caseing the team name
            const team = this.sd.currentTeam.trim().toLowerCase()
            if (this.sd.Cricket[team]) {
                has = true
            }
            return has
        }

        /**
         * 
         * @returns total number of players
         */
        playerCount(): number {
            var totalp = 0
            // avoiding complications by lower-caseing and whitespacing
            const team = this.sd.currentTeam.trim().toLowerCase()
            if (this.bTop10()) {
                totalp = this.sd.Cricket[team].length
            }
            return totalp
        }
        /**
         * 
         * @returns active players name 
         */
        activePlayers(): string {
            // avoiding complications by lower-caseing and whitespacing
            var team = this.sd.currentTeam.trim().toLowerCase()
            var players = this.sd.Cricket[team].toLocaleString()
            return players
        }

        /**
         * 
         * @param player_name to check for in the current roster
         * @returns true if present
         */
        bactivePlayer(player_name: string): boolean {
            // checks if the given player is not empty
            try {
                player_name = player_name.trim()
                if (player_name == " ") {
                    throw "empty name "
                }
            } catch (e) {
                throw "empty name "
            }
            // variables
            var bactive: boolean = false
            var f_lname: string[] = [" "], f_l: string | undefined = " ", fullname: string | undefined = " ", aPlayerList: string[] = [" "]

            // avoiding complications by lower-caseing and whitespacing
            const fname: string = player_name.trim().toLowerCase()
            const team = this.sd.currentTeam.trim().toLowerCase()

            aPlayerList = this.sd.Cricket[team]

            // stores spliited first and last names in the f_lname
            for (let i = 0; i < aPlayerList.length; ++i) {
                var f = this.split_up(aPlayerList[i], true, false)
                var l = this.split_up(aPlayerList[i], false, true)
                f_lname.push(f.toString())
                f_lname.push(l.toString())
            }
            // checks for the value if provided as the first name or last name of the player 
            f_l = f_lname.find((v) => { return v == fname })
            // checks for the value if provided as the full name of the player
            fullname = aPlayerList.find((v) => {
                return v == fname
            })
            // check if the condition matches for provided value
            switch (fname) {
                // found in either f or l name
                case f_l: { bactive = true } break
                // found in full name
                case fullname: { bactive = true } break
                // not found
                default: { bactive = false }
            }
            return bactive
        }

        /**
         * @NOTE it only returns based on the first caught whitespace
         * @param from first and last name
         * @param f if true returns first separated
         * @param l if true returns second separated
         * @param if both false returns an array 
         * @returns string split in two half portion
         */
        private split_up(from: string, f: boolean, l: boolean = false) {
            // to catch if the string doesnt have whitespace
            try {
                var i = from.search(' ')
                if (from[i] != ' ') {
                    throw 'already splited'
                }
            } catch (e) {
                throw 'already splited'
            }
            var col: string[] = []
            var first: string, last: string, whitespace: number = 0
            for (let i = 0; i < from.length; i++) {
                whitespace = from.search(' ')
            }
            first = from.slice(0, whitespace).trim() // trime: removes the whitespace left
            last = from.slice(whitespace).trim() // trim: removes the whitespace left
            col.push(first, last)
            if (f == true && l == false) { return first }
            if (l == true && f == false) { return last }
            else {
                return [...col]
            }
        }

    }(this)
    /**methods */
    /**
            * 
            * @returns total number of team
            */
    teamCount() {
        var totalt = Object.keys(this.Cricket).length
        return totalt
    }

    /**
     * 
     * @returns worlds best team name
     */
    worldsEliteTeam() {
        var eliteTeam = Object.keys(this.Cricket).toLocaleString()
        return eliteTeam
    }
}