import { FormBuilder, FormControl, FormGroup, Validators, } from "@angular/forms"
import { OCSFieldFormat } from "../../helpers/patterns"
import { BehaviorSubject, single } from "rxjs"
import { signal, WritableSignal } from "@angular/core"
import { GameMiscelleanous, IGameOver, IProfile_, JoinedProfile, NamesAsPerTeams, NamesAsPerTeamsRes, Pairing, RecBetItems, RoomSettings, RoomSettingsParcel, Rounds, SheetUpdate, SpectateInfo } from "../resources/lobbyResource"

export interface FetchPattern {
    Pack: string[]
}

export class LobbyFields {
    getBook: WritableSignal<OCSFieldFormat[]> = signal([])

    alertMsg: string = "what up you all fuck off"
    TagMsgs: WritableSignal<boolean> = signal(false)

    getRoomMode: string = ""
    getRoomBook: string = ""
    getRoomCategory: string = ""
    getRoomGameTime: number = 10
    getRoomDecisionTime: number = 20
    getDictionaries: WritableSignal<OCSFieldFormat[]> = signal([])

    getLists: WritableSignal<OCSFieldFormat[]> = signal([])
    getSetupDictionaries: WritableSignal<OCSFieldFormat[]> = signal([])

    getRoomSettings: RoomSettingsParcel = { decisionTime: 0, gameTime: 0, book: "", powers: [] }

    StartClock: WritableSignal<boolean> = signal(true)

    PorceedIntro: WritableSignal<boolean> = signal(false)
    DictionarySetupProceed: WritableSignal<boolean> = signal(false)
    DictionaryProceed: WritableSignal<boolean> = signal(false)
    ChallengeProceed: WritableSignal<boolean> = signal(false)
    EntranceProceed: WritableSignal<boolean> = signal(false)
    isChallengedRemove: WritableSignal<boolean> = signal(false)
    isDictionrayRemove: WritableSignal<boolean> = signal(false)
    getPlayerAgainst: GameMiscelleanous = { matchup: {} }

    /***
     ***** powers controls ******
     * 
     */
    // for freeze
    unFreezeControl: WritableSignal<boolean> = signal(false)
    unfreezeAll: WritableSignal<boolean> = signal(true)
    // for covert
    underTest: WritableSignal<boolean> = signal(false)
    // for draw
    drawOffer: WritableSignal<boolean> = signal(false)
    isDDecline: WritableSignal<boolean> = signal(false)
    drawMeetingDone: WritableSignal<boolean> = signal(false)

    onlineNicknames: WritableSignal<string[]> = signal([])
    refOinlineNickNames: string[] = []
    // for rewind
    clockBack: WritableSignal<boolean> = signal(false)
    // for tag 
    lock: WritableSignal<boolean> = signal(false)
    // for nexus
    nexusWord: WritableSignal<string> = signal("")
    tossAlert: WritableSignal<boolean> = signal(false)
    tossMsg: string = ""
    // for bet
    betItem: WritableSignal<OCSFieldFormat[]> = signal([])
    betPicker: WritableSignal<boolean> = signal(false)
    betting: WritableSignal<boolean> = signal(false)
    betCups: WritableSignal<RecBetItems> = signal({ firstCup: "", secondCup: "", thirdCup: "" })

    getDictionaryEvent: string = ""
    DictionarySet: WritableSignal<boolean> = signal(false)

    MyRoom = "NYG"
    JoinCode = "NYG"
    Recruit = ""
    RoomSpec = ""
    _MyId = ""
    // _guessIn = ""
    _chat = ""
    ChatToken = ""
    _GetChallenge: WritableSignal<string> = signal("NYG")
    _Erase = ""

    Hosts = ['']
    HostIds = ['']
    Rooms = ['']
    // isEntertainment: WritableSignal<boolean> = signal(false)
    // keys to get the players details
    $$$$$$$$$$$BLUE: string = 'BLUE'
    $$$$$$$$$$$RED: string = 'RED'
    isPlaying: "YES" | "NO" = "NO"

    getNickNames: string[] = []

    entertainmentCollection: OCSFieldFormat[] = []
    sportsCollection: OCSFieldFormat[] = []

    updateCheatSheet: Record<string, Record<number, string>> = {}

    MyTeam: string = "RED"
    MyNickName: string = "NYG"
    // vote: number = 1
    roundsKeys: string[] = ["Round1", "Round2", "Round3", "Round4", "Round5", "Round6"]
    flip: WritableSignal<OCSFieldFormat[]> = signal(
        [
            { OCStemplate: 'TAIL', OCSformControl: false, OCSToolTip: "" },
            { OCStemplate: 'HEAD', OCSformControl: false, OCSToolTip: "" },
        ])

    // basically for nickname->teamname
    ClashProfile: Record<string, Pairing> = {}

    ___currentRound: number = 1
    ___currentSet: number = 1

    RoomCap: number = 0
    ActiveConn: WritableSignal<number> = signal(0)
    RedScore: number = 0
    BlueScore: number = 0
    roomCapacity: number = 0
    BetProceed: WritableSignal<boolean> = signal(false)
    isHead: WritableSignal<boolean> = signal(false)
    block: WritableSignal<boolean> = signal(false)
    Session: WritableSignal<boolean> = signal(false)
    RoundOver: WritableSignal<boolean> = signal(false)

    Friend: WritableSignal<boolean> = signal(false)
    //StartGameSession: WritableSignal<boolean> = signal(false)
    JoinDone: WritableSignal<boolean> = signal(false)
    TossCoin: WritableSignal<boolean> = signal(false)
    canUsePower: WritableSignal<boolean> = signal(false)

    useNexusPower: WritableSignal<boolean> = signal(false)
    useRewindPower: WritableSignal<boolean> = signal(false)
    useFreezePower: WritableSignal<boolean> = signal(false)
    useDrawPower: WritableSignal<boolean> = signal(false)
    useTagPower: WritableSignal<boolean> = signal(false)
    useBetPower: WritableSignal<boolean> = signal(false)
    useCovertPower: WritableSignal<boolean> = signal(false)

    includeDrawPower: WritableSignal<boolean> = signal(false)
    includeFreezePower: WritableSignal<boolean> = signal(false)
    includeTagPower: WritableSignal<boolean> = signal(false)
    includeRewindPower: WritableSignal<boolean> = signal(false)
    includeNexusPower: WritableSignal<boolean> = signal(false)
    includeCovertPower: WritableSignal<boolean> = signal(false)
    includeBetPower: WritableSignal<boolean> = signal(false)
    NYGroomSettings: RoomSettings = {
        category: "", mode: "",
        book: "", gameTime: -1, decisionTime: 0,
        powers: {}, cap: 0, isEntertainment: false, roomname: '', private: false, friend: false
        , code: "", starter: false, setupToss: false, reverse: false, client: false, session: false
    }
    frames: WritableSignal<OCSFieldFormat[]> = signal([])
    // these tokens are same used in golang too // make sure to remove this line
    // signals
    _StartGame: WritableSignal<boolean> = signal(false)
    _Waiting: WritableSignal<boolean> = signal(false)
    _GameDone: WritableSignal<boolean> = signal(false)

    // on going session 
    _TossSession: WritableSignal<boolean> = signal(true)
    rWaiting: WritableSignal<boolean> = signal(false) // for reverse waiting
    _ChallengeDiscussion: WritableSignal<boolean> = signal(false)
    _DictionaryDiscussion: WritableSignal<boolean> = signal(false)
    // set name
    setNameForm = new FormGroup(
        { NameToken: new FormControl(' ') }
    )
    ChancesLeft: WritableSignal<number> = signal(0)
    ClashMessage: WritableSignal<boolean> = signal(false)
    OnFire: number = 0
    isClash: WritableSignal<boolean> = signal(false)
    FinalBoss: WritableSignal<boolean> = signal(false)
    LastDance: WritableSignal<boolean> = signal(false)
    IMBoss: WritableSignal<boolean> = signal(false)
    refToss: WritableSignal<OCSFieldFormat[]> = signal([
        { OCSformControl: false, OCStemplate: "HEADS", OCSToolTip: "" },
        { OCSformControl: false, OCStemplate: "TAILS", OCSToolTip: "" },
    ])
    spectateInfo: SpectateInfo = { teamName: "", playerName: "" }
    nullColl: WritableSignal<OCSFieldFormat[]> = signal([])
    getTopPosition: WritableSignal<number> = signal(0)
    currClick: WritableSignal<string> = signal("")
    defaultProceed: WritableSignal<boolean> = signal(true)
    guessTheme: string = "text25x"
    boardTheme: string = 'text25x'
    NYGMode: WritableSignal<string> = signal('mysterium')
    rating: number = 0
    gameOver: IGameOver = { alert: false, ratingDecrementedBy: 0, ratingIncrementedBy: 0, newRating: 0 }
    gameOverAlert: WritableSignal<boolean> = signal(false)
    messagePort: WritableSignal<boolean> = signal(false)
    joinedPlayer: JoinedProfile[] = []
    spectateTheme: WritableSignal<string> = signal('')
    nillRef: OCSFieldFormat[] = []
    freezeAccess: number = 7
}
export interface RoundKeys {
    r1: string | undefined
    r2: string | undefined
    r3: string | undefined
    r4: string | undefined
    r5: string | undefined
    r6: string | undefined
}