
import { RoomService } from "../../roomService/room.service";
import { FetchPattern, LobbyFields, } from "../vouchers/lobbyVouchers";
import { asyncScheduler, debounceTime, distinct, distinctUntilChanged, filter, interval, map, observeOn, repeat, Subject, takeUntil, timer } from "rxjs";
import { NgZone } from "@angular/core";
import { OCSFieldFormat } from "../../helpers/patterns";
import { ISpectate } from "../../roomService/resources/resource";
export interface ClashPowers {
    covert: boolean
    nexus: boolean
    freeze: boolean
    rewind: boolean
    draw: boolean
    tag: boolean
    bet: boolean
}

export interface ClashPowerReset {
    isBetting: boolean
    betting: boolean
    unfreeze: boolean
    underTest: boolean
}
interface ClashMatchResult {
    roomname: string
    pairing: Record<string, Pairing>
    round: number
}

interface ClashMatchToken {
    lock: boolean
    block: boolean
}
export interface ClashOn {
    roomname: string
    gameBegin: boolean
    chances: number
    mode: string
    pairing: Record<string, Pairing> // id-teamname
    ids: string[]
    teamName: string
    list: string[]
    round: number
    finalBoss: boolean
    lastDance: boolean
    gameTime: number
    discussionTime: number
}
export interface Pairing {
    teamname: string
    chances: number
}

export interface RoomTime {
    gameTime: number
    discussionTime: number
}
export interface NamesAsPerTeams {
    [teamName: string]: string
}
export interface RoomSettingsParcel {
    book: string
    gameTime: number
    decisionTime: number
    powers: string[]
}
export interface NamesAsPerTeamsRes {
    matchup: Record<string, string>
}
export interface RecBetItems {
    firstCup: string
    secondCup: string
    thirdCup: string
}
interface ClashProfile {
    NickNames: string[]
    id: string
}
interface ClashGameReward {
    onFire: number
    currentChance: number
}

export interface NestedMap {
    [key: string]: Rounds | undefined
}
export interface Rounds {
    round1: number,
    round2: number,
    round3: number,
    round4: number,
    round5: number,
    round6: number,
}
export interface Message {
    gameBegin: boolean
    startGame: boolean
    clash: boolean
    session: boolean
}

export interface SheetUpdate {
    name: string[]
    sheet: NestedMap
}
export interface PlayBet {
    doBet: boolean
    bets: string[]
    challenge: boolean
}
export interface ClashBet {
    isBetting: boolean
    cups: string[]
}

export interface ClasMatchUpdate {
    finalBoss: boolean
    lastDance: boolean
    imBoss: boolean
}
export interface MatchUp {
    img: string
    name: string
}
export interface GameMiscelleanous {
    matchup: Record<string, MatchUp>
}

export interface FreezePowerPackage {
    unfreeze: boolean
    canUnfreeze: boolean
    freezeUse: boolean
    freezeTime: number
}

export interface Powers {
    key: Record<string, boolean>

    canUsePower: boolean
}
export interface BSession {
    toss: boolean
    dictionary: boolean
    challenge: boolean
    game: boolean
    clash: boolean
}
export interface LocifyGame {
    redScore: number
    blueScore: number
    round: number
    set: number
    powers: Record<string, boolean>
    session: boolean
    mode: string
    roomname: string
    teamname: string
    stats: Record<string, Record<number, string>>
    nicknames: string[]
}
export interface RoomSettings {
    book: string
    mode: string
    category: string
    gameTime: number
    code: string
    roomname: string
    decisionTime: number
    powers: Record<string, boolean>
    cap: number
    isEntertainment: boolean
    friend: boolean
    private: boolean
    reverse: boolean
    starter: boolean
    setupToss: boolean
    client: boolean
    session: boolean
}
export interface Spectate {
    toss: boolean
    dictionary: boolean
    bet: boolean
    challenge: boolean
    tossCoin: boolean
    game: boolean
    nexusWord: boolean
    powers: boolean
    bets: []
    setDictionary: string
    setChallenge: string
    setBet: string
    oppoSetChallenge: string
    tossBody: string
}

export interface SpectateInfo {
    teamName: string
    playerName: string
}
export interface Frame {
    frames: string[]
}
export interface Credit {
    coins: string
    spurs: string
    boardTheme: string
    guessTheme: string
    id: string
    imageURL: string
    credits: IProfile_
    nexus: number
    freeze: number


}
export interface IProfile_ {
    gamesPlayed: number,
    tier: string,
    rating: number,
    points: number,
}
export interface JoinedProfile {
    name: string,
    profile: IProfile_
}
export interface IGameOver {
    ratingIncrementedBy: number
    ratingDecrementedBy: number
    newRating: number
    alert: boolean
}

export interface IMessage {
    message: string
    alert: boolean
}
export interface IRemovePower {
    power: Record<string, boolean>
}
/********   README    ***********
 * the framework is case-insensitive
 * because the framework is made using look-after rxjs technique
 * suppose token
 * var test = /Set:(?=[\s\S\w])/
 * var t = "ChallengeSet: 1"
 * var s = "Set: 1"
 * note: if you test s and t for Set:.... reg exp for both it will be tru
 * but if you test ChallengeSet.... reg exp it will give true and false
 * so make sure that anyting before colon must not be the same token
 * in-case if you won't able to get the desired result you must find that key 
 * ********* */

// @TODO: if the rxjs not able to pass make sure to check on the new browser or clear the cache

// avoiding implemnting the interface for now
export class LobbyResource extends LobbyFields {
    Unsub = new Subject<void>()
    Clean() {
        this.Unsub.next()
        this.Unsub.complete()
    }
    constructor(private r: RoomService, private run: NgZone) {
        super()
        var tokens = this.r.Rooms().pipe(takeUntil(this.Unsub),
            filter((token): token is string => token != undefined)
        )
        var locifyMissceleanous = this.r.Rooms().pipe(takeUntil(this.Unsub))
        locifyMissceleanous.subscribe({
            next: (token) => {
                var NC = /NYGcredits:(?=[\s\S\w])/
                var getCredits = NC.test(token)
                var NFPattern = /NYGFrames:(?=[\s\S\w])/
                var NSIPattern = /NYGspectateInfo:(?=[\s\S\w])/
                var spectateinfo = NSIPattern.test(token)
                var frames = NFPattern.test(token)
                var NMP = /NYGmessagePort:(?=[\s\S\w])/
                var getMessage = NMP.test(token)
                var NPJ = /NYGplayerJoin:(?=[\s\S\w])/
                var getPlayerJoin = NPJ.test(token)
                var NMN = /NYGmyNickname:(?=[\s\S\w])/
                var getMyNickname = NMN.test(token)
                var RP = /NYGremovePower:(?=[\s\S\w])/
                var removePower = RP.test(token)
                const count = this.Search(token, ":")
                switch (true) {
                    case removePower:
                        var p = 'NYGremovePower:'
                        var _p: IRemovePower = { power: {} }
                        var _p1: IRemovePower = this.afterColonJson(token, p, _p)
                        var z = Object.fromEntries(
                            Object.entries(_p1.power).map(([key, value]) => [
                                key.toLowerCase(),
                                value
                            ]),

                        ) as Record<string, boolean>
                        var k = Object.keys(z)

                        for (let i of k) {
                            switch (i) {
                                case "covert":
                                    this.includeCovertPower.set(_p1.power["covert"])
                                    break
                                case "draw":
                                    this.includeDrawPower.set(_p1.power["draw"])
                                    break
                                case "tag":
                                    this.includeTagPower.set(_p1.power["tag"])
                                    break
                                case "freeze":
                                    this.includeFreezePower.set(_p1.power["freeze"])
                                    break
                                case "rewind":
                                    this.includeRewindPower.set(_p1.power["rewind"])
                                    break
                                case "nexus":
                                    this.includeNexusPower.set(_p1.power["nexus"])
                                    break
                                case "bet":
                                    this.includeBetPower.set(_p1.power["bet"])
                                    break
                            }
                        }
                        console.log('powers: ', _p1)
                        console.log('converted keys: ', k)

                        break
                    case getMyNickname:
                        var n = this.afterColon(token, count, false, false)
                        this.MyNickName = n
                        break

                    case getPlayerJoin:
                        var p = 'NYGplayerJoin: '
                        var x: JoinedProfile = { name: "", profile: { gamesPlayed: 0, points: 0, tier: "", rating: 0 } }
                        var x_: JoinedProfile = this.afterColonJson(token, p, x)
                        this.joinedPlayer.push(x_)
                        break

                    case getMessage:
                        var p = 'NYGmessagePort: '
                        var g: IMessage = { message: "", alert: false }
                        var g_: IMessage = this.afterColonJson(token, p, g)
                        this.alertMsg = g_.message
                        this.messagePort.set(g_.alert)
                        timer(1500).subscribe({
                            next: () => {
                                this.messagePort.set(false)
                            }
                        })
                        break

                    case getCredits:
                        var p = 'NYGcredits: '
                        var gc: Credit = { guessTheme: "", boardTheme: "", id: "", coins: "", spurs: "", credits: { gamesPlayed: 0, points: 0, tier: "", rating: 0 }, imageURL: "", freeze: 0, nexus: 0 }
                        var gc_: Credit = this.afterColonJson(token, p, gc)
                        this.boardTheme = gc_.boardTheme
                        this.guessTheme = gc_.guessTheme
                        this.rating = gc_.credits.rating
                        break

                    case spectateinfo:
                        var p = 'NYGspectateInfo:'
                        var si: SpectateInfo = { teamName: "", playerName: "" }
                        var si_: SpectateInfo = this.afterColonJson(token, p, si)
                        this.spectateInfo = si_
                        break

                    case frames:
                        var p = 'NYGFrames:'
                        var j: Frame = { frames: [] }
                        var _f: Frame = this.afterColonJson(token, p, j)
                        var j_: OCSFieldFormat[] = []
                        for (let i of _f.frames) {
                            j_.push({ OCStemplate: i, OCSformControl: false, OCSToolTip: "" })
                        }

                        const uniqueMap = new Map<string, OCSFieldFormat>();

                        for (const element of j_) {
                            uniqueMap.set(element.OCStemplate, element);
                        }

                        const unique: OCSFieldFormat[] = Array.from(uniqueMap.values());
                        this.frames.set(unique)
                        console.log('freames: ', frames)
                        break
                }
            }
        })
        var locifyRoom = this.r.Rooms().pipe(takeUntil(this.Unsub))
        locifyRoom.subscribe({
            next: (token) => {

                var LRPattern = /LocifyRoom:(?=[\s\S\w])/
                var roomDets = LRPattern.test(token)
                var NGT = /NYGgameTime:(?=[\s\S\w])/
                var getGameTime = NGT.test(token)
                const count = this.Search(token, ":")
                switch (true) {
                    case getGameTime:
                        this.getRoomGameTime = this.afterColon(token, count, true)
                        break
                    case roomDets:
                        var p = 'LocifyRoom:'
                        var lr: RoomSettings = {
                            category: "", mode: "",
                            book: "", gameTime: -1, decisionTime: -1,
                            powers: {}, cap: 0, isEntertainment: false,
                            roomname: "", friend: false, private: false
                            , code: "", starter: false, setupToss: false, reverse: false,
                            client: false, session: false
                        }

                        var _lr: RoomSettings = this.afterColonJson(token, p, lr)
                        this.NYGroomSettings = _lr
                        this.getRoomCategory = _lr.category
                        this.getRoomMode = _lr.mode
                        this.getRoomDecisionTime = _lr.decisionTime
                        break
                }
            }
        })
        var locify = this.r.Rooms().pipe(takeUntil(this.Unsub))

        locify.subscribe({
            next: (token) => {
                var LGPattern = /LocifyGame:(?=[\s\S\w])/
                var _game = LGPattern.test(token)
                var NSPattern = /NYGSessionUpdate:(?=[\s\S\w])/
                var sessionUpdate = NSPattern.test(token)
                switch (true) {
                    case sessionUpdate:
                        var p = 'NYGSessionUpdate:'
                        var _ns: BSession = {
                            toss: false,
                            challenge: false, dictionary: false,
                            clash: false, game: false
                        }
                        var ns: BSession = this.afterColonJson(token, p, _ns)
                        this._TossSession.set(ns.toss)
                        this._ChallengeDiscussion.set(ns.challenge)
                        this._DictionaryDiscussion.set(ns.dictionary)
                        this._StartGame.set(ns.game)
                        break
                    case _game:
                        var p = 'LocifyGame:'
                        var lg: LocifyGame = {
                            redScore: 0, blueScore: 0,
                            round: 0, set: 0, powers: {},
                            session: false, mode: "", roomname: "",
                            teamname: "", nicknames: [], stats: {}
                        }
                        var _lg: LocifyGame = this.afterColonJson(token, p, lg)
                        this.Session.set(_lg.session)
                        this.___currentRound = _lg.round
                        this.___currentSet = _lg.set
                        this.MyRoom = _lg.roomname
                        this.MyTeam = _lg.teamname
                        this.RedScore = _lg.redScore
                        this.BlueScore = _lg.blueScore
                        this.updateCheatSheet = _lg.stats
                        this.getNickNames = _lg.nicknames
                        this.NYGMode.set(_lg.mode)
                        if (_lg.powers != null) {
                            var z = Object.fromEntries(
                                Object.entries(_lg.powers).map(([key, value]) => [
                                    key.toLowerCase(),
                                    value
                                ]),

                            ) as Record<string, boolean>

                            var keys = Object.keys(z)
                            for (let i of keys) {
                                switch (i) {
                                    case "covert":
                                        this.includeCovertPower.set(z["covert"])
                                        break
                                    case "draw":
                                        this.includeDrawPower.set(z["draw"])
                                        break
                                    case "tag":
                                        this.includeTagPower.set(z["tag"])
                                        break
                                    case "freeze":
                                        this.includeFreezePower.set(z["freeze"])
                                        break
                                    case "rewind":
                                        this.includeRewindPower.set(z["rewind"])
                                        break
                                    case "nexus":
                                        this.includeNexusPower.set(z["nexus"])
                                        break
                                    case "bet":
                                        this.includeBetPower.set(z["bet"])
                                        break
                                }
                            }
                        }
                        break
                }
            }
        })
        var powerTokens = this.r.Rooms().pipe(takeUntil(this.Unsub),
            filter((token): token is string => token != undefined))
        powerTokens.subscribe({
            next: (token) => {
                var _freeze = /NYGFreeze:(?=[\s\S\w])/
                var _freezeuse = _freeze.test(token)
                var power = /NYGPowers:(?=[\s\S\w])/
                var usedPower = power.test(token)
                switch (true) {
                    case usedPower:
                        var s: Powers = { key: {}, canUsePower: false }
                        var p = 'NYGPowers:'
                        var o: Powers = this.afterColonJson(token, p, s)
                        this.canUsePower.set(o.canUsePower)
                        for (let i of Object.keys(o.key)) {
                            switch (i) {
                                case "covert":
                                    this.useCovertPower.set(o.key["covert"])
                                    break
                                case "draw":
                                    this.useDrawPower.set(o.key["draw"])
                                    break
                                case "tag":
                                    this.useTagPower.set(o.key["tag"])
                                    break
                                case "freeze":
                                    this.useFreezePower.set(o.key["freeze"])
                                    break
                                case "rewind":
                                    this.useRewindPower.set(o.key["rewind"])
                                    break
                                case "nexus":
                                    this.useNexusPower.set(o.key["nexus"])
                                    break
                            }
                        }
                        break

                    case _freezeuse:
                        var p = 'NYGFreeze:'
                        var fu: FreezePowerPackage = { unfreeze: false, freezeUse: false, canUnfreeze: false, freezeTime: 0 }
                        var openFu: FreezePowerPackage = this.afterColonJson(token, p, fu)
                        this.useFreezePower.set(openFu.freezeUse)
                        this.unFreezeControl.set(openFu.canUnfreeze)
                        this.unfreezeAll.set(openFu.unfreeze)
                        this.freezeAccess = openFu.freezeTime
                        console.log('freeze: ', openFu)
                        break
                }
            }
        })
        var clashTokens = this.r.Rooms().pipe(takeUntil(this.Unsub),
            filter((token): token is string => token != undefined)
        )
        clashTokens.subscribe({
            next: (token) => {
                var NMPattern = /NYGMatchup:(?=[\s\S\w])/
                var nygmatch = NMPattern.test(token)
                var EMPattern = /Eliminated:(?=[\s\S\w])/
                var isEliminated = EMPattern.test(token)
                var CMUAPattern = /ClashMatchUpdate:(?=[\s\S\w])/
                var matchUpdate = CMUAPattern.test(token)
                var CPRPattern = /ClashPowerRest:(?=[\s\S\w])/
                var clashPowerReset = CPRPattern.test(token)
                var CIBPattern = /ClashUnderClashBet:(?=[\s\S\w])/
                var clashStruct = CIBPattern.test(token)
                var CBPattern = /ClashBet:(?=[\s\S\w])/
                var clashBet = CBPattern.test(token)
                var CRPattern = /ClashGameReward:(?=[\s\S\w])/
                var clashReward = CRPattern.test(token)
                var CMRPattern = /ClashGameResult:(?=[\s\S\w])/
                var getClashMatchResult = CMRPattern.test(token)
                var CPPattern = /ClashProfile:(?=[\s\S\w])/
                var clashProfile = CPPattern.test(token)
                var COPattern = /ClashOn:(?=[\s\S\w])/
                var clashOn = COPattern.test(token)
                var CPPattern = /ClashPowers:(?=[\s\S\w])/
                var clashPower = CPPattern.test(token)
                var CMUPattern = /ClashMatchUp:(?=[\s\S\w])/
                var getMatchups = CMUPattern.test(token)
                var CPattern = /Clash:(?=[\s\S\w])/
                var isClash = CPattern.test(token)
                var CWRPattern = /ClashResultWait:(?=[\s\S\w])/
                var clashWaitResult = CWRPattern.test(token)
                var NGO = /NYGGameOver:(?=[\s\S\w])/
                var getGameOver = NGO.test(token)
                const count = this.Search(token, ":")

                switch (true) {
                    case getGameOver:
                        var p = 'NYGGameOver:'
                        var _o: IGameOver = { ratingDecrementedBy: 0, ratingIncrementedBy: 0, newRating: 0, alert: false }
                        var jj: IGameOver = this.afterColonJson(token, p, _o)
                        this.gameOver = jj
                        this.gameOverAlert.set(this.gameOver.alert)
                        break

                    case nygmatch:
                        var p = "NYGMatchup:"
                        var o_: GameMiscelleanous = { matchup: {} }
                        var nm: GameMiscelleanous = this.afterColonJson(token, p, o_)
                        this.getPlayerAgainst = nm
                        break

                    case isEliminated:
                        var p = "Eliminated:"
                        var x_: Message = {
                            gameBegin: false, session: false,
                            startGame: false, clash: false
                        }
                        var j_: Message = this.afterColonJson(token, p, x_)
                        this.Session.set(j_.session)
                        this.isClash.set(j_.clash)
                        this._StartGame.set(j_.startGame)
                        break

                    case matchUpdate:
                        var p = 'ClashMatchUpdate:'
                        var mcw: ClasMatchUpdate = { finalBoss: false, lastDance: false, imBoss: false }
                        var mc: ClasMatchUpdate = this.afterColonJson(token, p, mcw)
                        this.FinalBoss.set(mc.finalBoss)
                        this.LastDance.set(mc.lastDance)
                        this.IMBoss.set(mc.imBoss)
                        break

                    case clashPowerReset:
                        var p = 'ClashPowerReset:'
                        var cpr: ClashPowerReset = { isBetting: false, underTest: false, unfreeze: true, betting: false }
                        var crp2: ClashPowerReset = this.afterColonJson(token, p, cpr)
                        this.betting.set(crp2.betting)
                        this.betPicker.set(crp2.isBetting)
                        this.unfreezeAll.set(crp2.unfreeze)
                        this.underTest.set(crp2.underTest)
                        break

                    case clashStruct:
                        var z: ClashBet = { isBetting: false, cups: [] }
                        var p = 'ClashUnderClashBet:'
                        var be: ClashBet = this.afterColonJson(token, p, z)
                        this.betting.set(true)
                        var _formate: OCSFieldFormat[] = []
                        for (let i of be.cups) {
                            _formate.push({ OCSformControl: false, OCStemplate: i, OCSToolTip: "" })
                        }
                        this.betItem.set(_formate)
                        var zq: RecBetItems = { firstCup: be.cups[0], secondCup: be.cups[1], thirdCup: be.cups[2] }
                        this.betCups.set(zq)
                        this._ChallengeDiscussion.set(false) // explicitly set to false
                        break

                    case clashBet:
                        var _i: PlayBet = { challenge: false, doBet: false, bets: [] }
                        var p = 'ClashBet:'
                        var bp: PlayBet = this.afterColonJson(token, p, _i)
                        this.betPicker.set(bp.doBet)
                        this._ChallengeDiscussion.set(bp.challenge)
                        var _formate: OCSFieldFormat[] = []
                        for (let i of bp.bets) {
                            _formate.push({ OCSformControl: false, OCStemplate: i, OCSToolTip: "" })
                        }
                        this.betItem.set(_formate)
                        break

                    case clashReward:
                        var p = 'ClashGameReward:'
                        var __f: ClashGameReward = {
                            onFire: 0, currentChance: 0
                        }
                        var getCrw: ClashGameReward = this.afterColonJson(token, p, __f)
                        this.OnFire = getCrw.onFire
                        this.ChancesLeft.set(getCrw.currentChance)
                        break

                    case getClashMatchResult:
                        var p = 'ClashGameResult:'
                        var _f: ClashMatchResult = {
                            roomname: "", pairing: {},
                            round: 0
                        }
                        var getCr: ClashMatchResult = this.afterColonJson(token,
                            p, _f)
                        this.ClashProfile = getCr.pairing
                        this.___currentRound = getCr.round
                        break

                    case clashWaitResult:
                        var cwr = this.afterColon<false, false>(token, count, false, false)
                        this.ClashMessage.set(true)
                        this.alertMsg = cwr
                        break

                    case isClash:
                        var __s = this.afterColon<false, true>(token, count, false, true)
                        this.isClash.set(__s)
                        break

                    case getMatchups:
                        var _j: ClashMatchToken = { block: false, lock: false }
                        var _parse = 'ClashMatchUp:'
                        var __open: ClashMatchToken = this.afterColonJson(token, _parse, _j)

                        this.lock.set(__open.lock)
                        this.block.set(__open.block)
                        break

                    case clashProfile:
                        let j: ClashProfile = { NickNames: [], id: "" }
                        var _parse = 'ClashProfile:'
                        var _open: ClashProfile = this.afterColonJson(token, _parse, j)
                        this.getNickNames = _open.NickNames
                        this.getNickNames = this.getNickNames.map(e => e.substring(0, 8))
                        break

                    case clashPower:
                        var i: ClashPowers = {
                            covert: false, tag: false, draw: false,
                            nexus: false, bet: false, freeze: false, rewind: false
                        }
                        var parse = "ClashPowers:"
                        var open: ClashPowers = this.afterColonJson(token, parse, i)
                        var ma: Record<string, boolean> = {}
                        for (let key of Object.keys(open)) {
                            ma[key] = open[key as keyof ClashPowers]
                        }
                        this.includeBetPower.set(ma['bet'])
                        this.includeDrawPower.set(ma['draw'])
                        this.includeTagPower.set(ma['tag'])
                        this.includeNexusPower.set(ma['nexus'])
                        this.includeRewindPower.set(ma['rewind'])
                        this.includeFreezePower.set(ma['freeze'])
                        this.includeCovertPower.set(ma['covert'])
                        break

                    case clashOn:
                        var parse = "ClashOn:"
                        var o: ClashOn = {
                            roomname: "", chances: 10, mode: "", gameBegin: false, pairing: {}, ids: [],
                            teamName: "", list: [], round: 0, finalBoss: false, lastDance: false,
                            gameTime: 0, discussionTime: 0
                        }
                        var ___open: ClashOn = this.afterColonJson(token, parse, o)

                        this.isClash.set(true)

                        this.___currentRound = ___open.round
                        this.NYGMode.set(___open.mode)
                        this.ChancesLeft.set(___open.chances)
                        this.Session.set(___open.gameBegin)
                        this.MyRoom = ___open.roomname
                        this.getRoomGameTime = ___open.gameTime
                        this.getRoomDecisionTime = ___open.discussionTime
                        for (let nikcnames of Object.keys(___open.pairing)) {
                            this.getNickNames.push(nikcnames)
                        }
                        this.MyTeam = ___open.teamName.toUpperCase()
                        this.ClashProfile = ___open.pairing

                        break
                }
            }
        })
        tokens.subscribe({
            next: (token: string) => {
                run.run(() => {
                    const ROCpattern = /RoomCode:(?=[\s\S\w])/
                    const Rpattern = /RoomCap:(?=[\d])/
                    const Hpattern = /HostID:(?=[\s\S\w])/
                    const Gpattern = /Searching: (?=[\s\S\w])/
                    const Sppattern = /Room Spec:(?=[\s\S\w])/
                    const Apattern = /ActiveConns:(?=[\s\S\d])/
                    const Shpattern = /GameBegin:(?=[\s\S\w])/
                    const UBRpattern = /UpdateBoardR:(?=[\s\S\w])/
                    const UBBpattern = /UpdateBoardB:(?=[\s\S\w])/
                    const ROpattern = /RoundOver:(?=[\s\S\w])/
                    const LockPpattern = /Lock:(?=[\s\S\w])/
                    const GDpattern = /GameDone:(?=[\s\S\w])/
                    const SGpattern = /StartGame:(?=[\s\S\w])/
                    const GIpattern = /GuessIn:(?=[\s\S\w])/
                    const Bpattern = /Block:(?=[\s\S\w])/
                    const Ppattern = /PlayingSession:(?=[\s\S\w])/
                    const Cpattern = /ChatToken:(?=[\s\S\w])/
                    const Wpattern = /Waiting:(?=[\s\S\w])/
                    const TSOpattern = /TossSession:(?=[\s\S\w])/
                    const CDpattern = /ChallengeDiscussion:(?=[\s\S\w])/
                    const CSTpattern = /ChallengeGuess:(?=[\s\S\w])/
                    const JDpattern = /JoinDone:(?=[\s\S\w])/
                    const TCpattern = /TossCoin:(?=[\s\s\w])/
                    const STpattern = /Set:(?=[\s\S\w])/
                    const RDpattern = /Round:(?=[\s\S\w])/
                    const TNpattern = /Team:(?=[\s\S\w])/
                    const GNpattern = /NicknameLists:(?=[\s\S\w])/
                    const UCSpattern = /UpdateCheatSheet[\s\S\w]/
                    const TOpattern = /Toss:(?=[\s\S\w])/
                    const DDpattern = /DictionaryDiscussion:(?=[\s\S\w])/
                    const PApattern = /PowerActivated:(?=[\s\S\w])/

                    const NPVpattern = /NexusWord:(?=[\s\S\w])/
                    const MRpattern = /MyRoom: (?=[\s\S\w])/
                    const FRpattern = /Friend:(?=[\s\S\w])/
                    const INpattern = /IsEnterTainment:(?=[\s\S\w])/
                    const NNpattern = /MyNickName:(?=[\s\S\w])/

                    const RRpattern = /ClockRestart:(?=[\s\S\w])/
                    const DOpattern = /DrawOffer:(?=[\s\S\w])/
                    const UTpattern = /UnderTest:(?=[\s\S\w])/

                    // powers used
                    const DPUpattern = /Draw:(?=[\s\S\w])/
                    const NPUpattern = /Nexus:(?=[\s\S\w])/
                    const TPUpattern = /Tag:(?=[\s\S\w])/
                    const FPUpattern = /Freeze:(?=[\s\S\w])/
                    const RPUpattern = /Rewind:(?=[\s\S\w])/
                    const BPUpattern = /Bet:(?=[\s\S\w])/
                    const CPUpattern = /Covert:(?=[\s\S\w])/

                    const ALpattern = /AlertTeam:(?=[\s\S\w])/

                    const alertTeam = ALpattern.test(token)

                    const useNexus = NPUpattern.test(token)
                    const useBetPower = BPUpattern.test(token)
                    const useCovertPower = CPUpattern.test(token)
                    const useDrawPower = DPUpattern.test(token)
                    const useRewindPower = RPUpattern.test(token)
                    const useTagPower = TPUpattern.test(token)
                    const useFreezePower = FPUpattern.test(token)

                    const underTest = UTpattern.test(token)
                    const drawOffer = DOpattern.test(token)
                    const rewindRestart = RRpattern.test(token)

                    const myNickName = NNpattern.test(token)
                    const isEntertainment = INpattern.test(token)
                    const isFriend = FRpattern.test(token)
                    const myRoom = MRpattern.test(token)

                    const getFriendCode = ROCpattern.test(token)

                    const getNexus = NPVpattern.test(token)

                    // const FPIpattern = /PowerFreeze:(?=[\s\S\w])/
                    // const NPIpattern = /PowerNexus:(?=[\s\sW])/
                    // const RPIpattern = /PowerRewind:(?=[\s\S\w])/
                    // const DPIpattern = /PowerDraw:(?=[\s\S\w])/
                    // const TPIpattern = /PowerTag:(?=[\s\S\w])/
                    // const CPIpattern = /PowerCovert:(?=[\s\S\w])/
                    // const BPIpattern = /PowerBet:(?=[\s\S\w])/

                    // const includeDrawPower = DPIpattern.test(token)
                    // const includeNexusPower = NPIpattern.test(token)
                    // const includeFreezePower = FPIpattern.test(token)
                    // const includeRewindPower = RPIpattern.test(token)
                    // const includeTagPower = TPIpattern.test(token)
                    // const includeCovertPower = CPIpattern.test(token)
                    // const includeBetPower = BPIpattern.test(token)

                    const isPowerActivated = PApattern.test(token)
                    const dictionaryDiscussion = DDpattern.test(token)
                    const setToss = TOpattern.test(token)
                    const getCheatSheetUpdate = UCSpattern.test(token)

                    const getNN = GNpattern.test(token)
                    const myTeam = TNpattern.test(token)
                    const tossCoin = TCpattern.test(token)
                    const currentSet = STpattern.test(token)
                    const currentRound = RDpattern.test(token)
                    const isJoinDone = JDpattern.test(token)
                    const getChallenge = CSTpattern.test(token)
                    const challengeDiscussion = CDpattern.test(token)
                    const gameDone = GDpattern.test(token)
                    const tossSession = TSOpattern.test(token)
                    const waiting = Wpattern.test(token)
                    const chatting = Cpattern.test(token)
                    const playingSession = Ppattern.test(token)
                    const isBlock = Bpattern.test(token)
                    const guessIn = GIpattern.test(token)
                    const startGame = SGpattern.test(token)
                    const lockPlayers = LockPpattern.test(token)
                    const updateRoomCapacity = Rpattern.test(token)
                    const viewHostIDs = Hpattern.test(token)
                    const viewRecruitment = Gpattern.test(token)
                    const viewRoomSpecification = Sppattern.test(token)
                    const viewActiveConnections = Apattern.test(token)
                    const isPlaying = Shpattern.test(token)
                    const isRoundOver = ROpattern.test(token)

                    const RDTpattern = /RoomDecisionTime:(?=[\s\S\w])/
                    const RGTpattern = /RoomGameTime:(?=[\s\s\w])/
                    const RCApattern = /RoomCategory:(?=[\s\S\w])/
                    const RBpattern = /RoomBook:(?=[\s\sW])/
                    const RMpattern = /RoomMode:(?=[\s\S\w])/

                    const getRoomMode = RMpattern.test(token)
                    const getRoomBook = RBpattern.test(token)
                    const getRoomCategory = RCApattern.test(token)
                    const getRoomGameTime = RGTpattern.test(token)
                    const getRoomDecisionTime = RDTpattern.test(token)

                    const updateRedScore = UBRpattern.test(token)
                    const updateBlueScore = UBBpattern.test(token)

                    const PRpattern = /Press:(?=[\s\S\w])/
                    const UFpattern = /Unfreeze:(?=[\s\S\w])/

                    const unfreeze = UFpattern.test(token)
                    const press = PRpattern.test(token)

                    const DOfpattern = /DrawOfferDeclined:(?=[\s\S\w])/
                    const isOfferAccepted = DOfpattern.test(token)

                    const TMpattern = /TagMsg:(?=[\s\S\w])/
                    const tagMsg = TMpattern.test(token)

                    const DMSpattern = /DrawMsg:(?=[\s\Sw])/
                    const drawMsg = DMSpattern.test(token)

                    const DMpattern = /DrawMeetingDone:(?=[\s\S\w])/
                    const isDrawMeetingDone = DMpattern.test(token)
                    const BIpattern = /BetItem:(?=[\s\S\w])/
                    const betItems = BIpattern.test(token)
                    const BEpattern = /Betting:(?=[\s\S\w])/
                    const isBetting = BEpattern.test(token)
                    const DSpattern = /DictionarySet:(?=[\s\S\w])/
                    const isDictionarySet = DSpattern.test(token)
                    const BPpattern = /BetPicker:(?=[\s\S\w])/
                    const betPicker = BPpattern.test(token)

                    const DEpattern = /DictionaryEvent:(?=[\s\S\W])/
                    const getDictionaryEvent = DEpattern.test(token)

                    const BCpattern = /BetCups:(?=[\s\S\w])/
                    const betCups = BCpattern.test(token)

                    const DUpattern = /DURL:(?=[\s\S\w])/
                    const IUpattern = /CURL:(?=[\s\S\w])/

                    const getDURL = DUpattern.test(token)
                    const getCURL = IUpattern.test(token)

                    const FBpattern = /FetchedBook:(?=[\s\S\w])/
                    const getBooks = FBpattern.test(token)

                    const RDipattern = /DictionaryRemove:(?=[\s\S\w])/
                    const removeDictionary = RDipattern.test(token)

                    const RCipattern = /ChallengeRemove:(?=[\s\S\w])/
                    const removeChallenge = RCipattern.test(token)

                    const TDMepattern = /TrashDMessage:(?=[\s\S\w])/
                    const trashDMessage = TDMepattern.test(token)
                    const TCMepattern = /TrashCMessage:(?=[\s\S\w])/
                    const trashCMessage = TCMepattern.test(token)
                    const DLpattern = /DURLT:(?=[\s\S\w])/
                    const getDictionaryToSetup = DLpattern.test(token)

                    const IHpattern = /Heads:(?=[\s\S\w])/
                    const isHead = IHpattern.test(token)

                    const ONNpattern = /OnlineNickName:(?=[\s\S\w])/
                    const onlineNickName = ONNpattern.test(token)

                    const TApattern = /TeamAgainstx:(?=[\s\S\w])/
                    const getTeamAgainst = TApattern.test(token)
                    const RSpattern = /ParcelRoomSettings:(?=[\s\S\w])/
                    const getRoomSetting = RSpattern.test(token)

                    const GTpattern = /RoomTime:(?=[\s\S\w])/
                    const _getTime = GTpattern.test(token)


                    const TMSpattern = /TossMsg:(?=[\s\S\w])/
                    const getTossMsg = TMSpattern.test(token)

                    const count = this.Search(token, ":")

                    switch (true) {

                        case getTossMsg:
                            this.tossMsg = this.afterColon<false, false>(token, count, false, false)
                            this.tossAlert.set(true)

                            break

                        // case getRoomSetting:
                        //     var parse = "ParcelRoomSettings:"
                        //     var g: RoomSettingsParcel = { powers: [], book: "", gameTime: 0, decisionTime: 0 }
                        //     var _temp: RoomSettingsParcel = this.afterColonJson(token, parse, g)
                        //     this.getRoomSettings = _temp
                        //     for (let i of _temp.powers) {
                        //         switch (true) {
                        //             case i == "DRAW":
                        //                 this.RincludeDrawPower.set(true)
                        //                 break
                        //             case i == "COVERT":
                        //                 this.RincludeCovertPower.set(true)
                        //                 break
                        //             case i == "TAG":
                        //                 this.RincludeTagPower.set(true)
                        //                 break
                        //             case i == "FREEZE":
                        //                 this.RincludeFreezePower.set(true)
                        //                 break
                        //             case i == "REWIND":
                        //                 this.RincludeRewindPower.set(true)
                        //                 break
                        //             case i == "BET":
                        //                 this.RincludeBetPower.set(true)
                        //                 break
                        //             case i == "NEXUS":
                        //                 this.RincludeNexusPower.set(true)
                        //                 break
                        //         }
                        //     }
                        //     this.RgameTime = this.getRoomSettings.gameTime
                        //     this.RdecisionTime = this.getRoomSettings.decisionTime
                        //     break

                        case _getTime:
                            var parse = "RoomTime:"
                            var __x: RoomTime = { gameTime: 0, discussionTime: 0 }
                            var _j: RoomTime = this.afterColonJson(token, parse, __x)
                            this.getRoomGameTime = _j.gameTime
                            this.getRoomDecisionTime = _j.discussionTime
                            break

                        case onlineNickName:
                            var _token = this.afterColon(token, count, false, false)
                            this.refOinlineNickNames.push(_token)
                            this.onlineNicknames.set(this.refOinlineNickNames)
                            break

                        case isHead:
                            this.isHead.set(this.afterColon<false, true>(token, count, false, true))
                            break

                        case getDictionaryToSetup:
                            var parse = "DURLT:"
                            var f: FetchPattern = { Pack: [] }
                            var _getDictionary: FetchPattern = this.afterColonJson(token, parse, f)
                            var store: OCSFieldFormat[] = []
                            for (let dict of _getDictionary.Pack) {
                                store.push({ OCStemplate: dict, OCSformControl: false, OCSToolTip: "" })
                            }
                            this.getSetupDictionaries.set(store)

                            if (this.getSetupDictionaries().length > 0)
                                this.DictionarySetupProceed.set(true)
                            break

                        case trashCMessage:
                            this.isChallengedRemove.set(true)
                            this.alertMsg = this.afterColon<false, false>(token, count, false, false)
                            this.isDictionrayRemove.set(false)
                            break

                        case trashDMessage:
                            this.isDictionrayRemove.set(true)
                            this.alertMsg = this.afterColon<false, false>(token, count, false, false)
                            break

                        case removeDictionary:
                            var remove = this.afterColon<false, false>(token, count, false, false)
                            this.getDictionaries.set(this.getDictionaries().filter((e) => e.OCStemplate != remove))

                            this.getSetupDictionaries.set(this.getSetupDictionaries().filter(e => e.OCStemplate != remove))
                            break

                        case removeChallenge:
                            var remove = this.afterColon<false, false>(token, count, false, false).trim()
                            this.getLists.set(this.getLists().filter((e) => e.OCStemplate != remove))
                            break

                        case getBooks:
                            var parse = 'FetchedBook:'
                            var tr: FetchPattern = { Pack: [] }
                            const fetched: FetchPattern = this.afterColonJson(token, parse, tr)
                            var store: OCSFieldFormat[] = []
                            for (let book of fetched.Pack) {
                                store.push({ OCStemplate: book, OCSformControl: false, OCSToolTip: "" })
                            }
                            this.getBook.set(store)
                            break

                        case getCURL:
                            var parse = "CURL:"
                            var f: FetchPattern = { Pack: [] }
                            var _getDictionary: FetchPattern = this.afterColonJson(token, parse, f)
                            var store: OCSFieldFormat[] = []
                            for (let dict of _getDictionary.Pack) {
                                store.push({ OCStemplate: dict, OCSformControl: false, OCSToolTip: "" })
                            }
                            this.getLists.set(store)

                            if (this.getLists().length > 0)
                                this.ChallengeProceed.set(true)
                            break

                        case getDURL:
                            var parse = "DURL:"
                            var f: FetchPattern = { Pack: [] }
                            var _getDictionary: FetchPattern = this.afterColonJson(token, parse, f)

                            var store: OCSFieldFormat[] = []
                            for (let dict of _getDictionary.Pack) {
                                store.push({ OCStemplate: dict, OCSformControl: false, OCSToolTip: "" })
                            }
                            this.getDictionaries.set(store)

                            if (this.getDictionaries().length > 0)
                                this.DictionaryProceed.set(true)
                            break

                        case drawMsg:
                            this.alertMsg = this.afterColon<false, false>(token, count, false, false)
                            this.drawOffer.set(true)
                            break

                        case isDrawMeetingDone:
                            this.drawMeetingDone.set(this.afterColon(token, count, false, true))
                            break

                        case betCups:
                            var parse = "BetCups: "
                            var ix: RecBetItems = { firstCup: "", secondCup: "", thirdCup: "" }
                            var temp = this.afterColonJson(token, parse, ix) as RecBetItems

                            this.betCups.set(temp)
                            break

                        case isBetting:
                            this.betting.set(this.afterColon(token, count, false, true))
                            break

                        case isDictionarySet:
                            this.DictionarySet
                                .set(this.afterColon(token, count, false, true))
                            break

                        case getDictionaryEvent:
                            this.getDictionaryEvent = this.afterColon<false, false>(token, count, false, false)
                            break

                        case getRoomMode:
                            this.getRoomMode = this.afterColon<false, false>(token, count, false, false)
                            break

                        case getRoomBook:
                            this.getRoomBook = this.afterColon<false, false>(token, count, false, false)
                            break

                        case getRoomCategory:
                            this.getRoomCategory = this.afterColon<false, false>(token, count, false, false)
                            break

                        case betPicker:
                            this.betPicker.set(this.afterColon(token, count, false, true))
                            break

                        case betItems:
                            var parse = "BetItem: "
                            var ix: RecBetItems = { firstCup: "", secondCup: "", thirdCup: "" }
                            var temp = this.afterColonJson(token, parse, ix) as RecBetItems
                            var conv: OCSFieldFormat[] = [
                                { OCSformControl: false, OCStemplate: temp.firstCup, OCSToolTip: "" },
                                { OCSformControl: false, OCStemplate: temp.secondCup, OCSToolTip: "" },
                                { OCSformControl: false, OCStemplate: temp.thirdCup, OCSToolTip: "" },
                            ]
                            this.betItem.set(conv)
                            if (this.betItem().length > 0)
                                this.BetProceed.set(true)
                            break

                        case tagMsg:
                            this.TagMsgs
                                .set(this.afterColon(token, count, false, true))
                            break

                        case alertTeam:
                            this.alertMsg = this.afterColon<false, false>(token, count, false, false)
                            break

                        case isOfferAccepted:
                            this.isDDecline
                                .set(this.afterColon(token, count, false, true))
                            break

                        case unfreeze:
                            this.unfreezeAll
                                .set(this.afterColon(token, count, false, true))
                            break

                        // case includeDrawPower:
                        //     this.includeDrawPower
                        //         .set(this.afterColon(token, count, false, true))
                        //     break

                        // case includeFreezePower:
                        //     this.includeFreezePower

                        //         .set(this.afterColon(token, count, false, true))
                        //     break

                        // case includeTagPower:
                        //     this.includeTagPower
                        //         .set(this.afterColon(token, count, false, true))

                        //     break

                        // case includeRewindPower:
                        //     this.includeRewindPower
                        //         .set(this.afterColon(token, count, false, true))

                        //     break

                        // case includeNexusPower:
                        //     this.includeNexusPower
                        //         .set(this.afterColon(token, count, false, true))
                        //     break

                        // case includeCovertPower:
                        //     this.includeCovertPower
                        //         .set(this.afterColon(token, count, false, true))
                        //     break

                        // case includeBetPower:
                        //     this.includeBetPower
                        //         .set(this.afterColon(token, count, false, true))
                        //     break

                        case useNexus:
                            this.useNexusPower.set(this.afterColon(token, count, false, true))
                            break

                        case useBetPower:
                            this.useBetPower.set(this.afterColon(token, count, false, true))
                            break

                        case useRewindPower:
                            this.useRewindPower.set(this.afterColon(token, count, false, true))
                            break

                        case useDrawPower:
                            this.useDrawPower.set(this.afterColon(token, count, false, true))
                            break

                        case useTagPower:
                            this.useTagPower.set(this.afterColon(token, count, false, true))
                            break

                        case useCovertPower:
                            this.useCovertPower.set(this.afterColon(token, count, false, true))
                            break

                        case useFreezePower:
                            this.useFreezePower.set(this.afterColon(token, count, false, true))
                            break


                        case underTest:
                            this.underTest.set(this.afterColon(token, count, false, true))
                            break

                        case drawOffer:
                            this.drawOffer.set(this.afterColon(token, count, false, true))
                            break

                        case press:
                            this.unFreezeControl.set(this.afterColon(token, count, false, true))
                            break

                        case rewindRestart:
                            this.clockBack.set(this.afterColon(token, count, false, true))
                            break

                        case myNickName:
                            this.MyNickName = this.afterColon(token, count, false, false)
                            break

                        // case isEntertainment:
                        //     this.isEntertainment.set(this.afterColon(token, count, false, true))
                        //     break

                        case isFriend:
                            this.Friend.set(this.afterColon(token, count, false, true))
                            break

                        case getFriendCode:
                            this.JoinCode = this.afterColon<false, false>(token, count, false, false)
                            break

                        case myRoom:
                            this.MyRoom = this.afterColon<false, false>(token, count, false, false)
                            this.MyRoom = this.MyRoom.toUpperCase()
                            break

                        case updateRedScore:
                            this.RedScore += this.afterColon(token, count, true, false)
                            break

                        case updateBlueScore:
                            this.BlueScore += this.afterColon(token, count, true, false)
                            break

                        case getNexus:
                            this.nexusWord.set(this.afterColon<false, false>(token, count, false,))
                            break

                        case isPowerActivated:
                            this.canUsePower.set(this.afterColon(token, count, false, true))
                            break

                        case dictionaryDiscussion:
                            this._DictionaryDiscussion.set(this.afterColon(token, count, false, true))
                            break

                        case tossCoin:
                            timer(1450).subscribe({
                                next: () => {
                                    this.TossCoin.set(this.afterColon(token, count, false, true))
                                }
                            })
                            break

                        case getNN:
                            this.getNickNames.push(this.afterColon<false, false>(token, count, false, false))
                            break

                        case myTeam:
                            this.MyTeam = this.afterColon<false, false>(token, count, false, false)
                            break

                        case currentSet:
                            this.___currentSet = this.afterColon(token, count, true, false)
                            break

                        case currentRound:
                            this.___currentRound = this.afterColon(token, count, true, false)
                            break

                        case updateRoomCapacity:
                            this.RoomCap = this.afterColon(token, count, true, false)
                            break

                        case viewHostIDs:
                            this.HostIds.push(this.afterColon<false, false>(token, count, false, false))
                            break

                        case viewRecruitment:
                            this.Recruit = this.afterColon<false, false>(token, count, false, false)
                            break

                        case viewRoomSpecification:
                            this.RoomSpec = this.afterColon<false, false>(token, count, false, false)
                            break

                        case viewActiveConnections:
                            this.ActiveConn.set(this.afterColon(token, count, true, false))
                            break

                        case isPlaying:
                            this.Session.set(this.afterColon(token, count, false, true))
                            // erase
                            // if (this.Session == false) {
                            //     this.HostIds = this.HostIds.slice(this.HostIds.length - 1, 0)
                            //     this.Hosts = this.Hosts.slice(this.Hosts.length - 1, 0)
                            //     this.TeamA = this.TeamA.slice(this.TeamA.length - 1, 0)
                            //     this.TeamB = this.TeamB.slice(this.TeamB.length - 1, 0)
                            //     this.lock = true
                            //     this.HomeScore = 0
                            //     this.AwayScore = 0
                            // }
                            break

                        case isJoinDone:
                            this.JoinDone.set(this.afterColon(token, count, false, true))
                            break

                        case isRoundOver:
                            this.RoundOver.set(this.afterColon(token, count, false, true))
                            break

                        case lockPlayers:
                            this.lock.set(this.afterColon(token, count, false, true))
                            if (this.lock()) {
                                this.isPlaying = "YES"
                            } else {
                                this.isPlaying = "NO"
                            }
                            break

                        case startGame:
                            // delay so that we can play animation
                            // setTimeout(() => {
                            this._StartGame.set(this.afterColon(token, count, false, true))
                            // }, 1000); // 1 sec =1000ms
                            break

                        case getChallenge:
                            const ___token = this.afterColon<false, false>(token, count, false, false)
                            if (___token)
                                this._GetChallenge.set(___token)
                            if (this._GetChallenge().length > 0)
                                this.PorceedIntro.set(true)
                            break

                        // case guessIn:
                        //     this._guessIn = this.afterColon(token, count, false)
                        //     break

                        case isBlock:
                            this.block.set(this.afterColon(token, count, false, true))
                            break

                        // case playingSession:
                        //     this.StartGameSession.set(this.afterColon(token, count, false, true))

                        //     break

                        case chatting:
                            var h = this.afterColon<false, false>(token, count, false, false)
                            if (h) {
                                this.ChatToken = h
                            }
                            break

                        case tossSession:
                            this._TossSession.set(this.afterColon(token, count, false, true))
                            break

                        case challengeDiscussion:
                            this._ChallengeDiscussion.set(this.afterColon(token, count, false, true))
                            break

                        case waiting:
                            this._Waiting.set(this.afterColon(token, count, false, true))
                            break

                        case gameDone:
                            this._GameDone.set(this.afterColon(token, count, false, true))
                            break

                        default:
                    }
                })
            },
            error: (errr) => {
                console.error("error in fethc: ", errr)
            }
        })
    }

    toBool(v: string): boolean {
        var test = false
        switch (v) {
            case "yes":
                test = true
                break
            case "no":
                test = false
        }
        return test
    }

    removeDuplicate(token: string[]): string[] {
        token = token.filter((e, i) => {
            if (e != token[i - 1]) {
                return e
            } else {
                return token[i]
            }
        })
        return token
    }

    /**
     * 
     * @param str to find whitespace index
     * @returns index of whitesapce
     */
    firstIndexofWhitespace(str: string): number {
        var ws = /\s/
        var count = 0
        for (let i = 0; i < str.length; i++) {
            if (ws.test(str[i])) {
                count = i
                break
            }
        }
        return count
    }

    _ToString(str: string): string {
        var temp = str.toString()
        for (let i = 0; i <= str.length; i++) {
            if (str[i] == ',') {
                str = str.replace(',', '')
            }
        }
        return str
    }

    // afterColonJson removes the parse value from the prefix and returns in given interface
    afterColonJson(token: string | undefined, parse: string, _interface: any) {
        if (token) {
            token = token.substring(parse.length)
            _interface = JSON.parse(token)
        }
        return _interface
    }
    /**
     * 
     * @param token  to convert into number, boolean, or string
     * @param inx  to slice index from
     * @param num  if true returns in number else string
     * @param bool if true returns in boolean
     * @returns conversion of the rxjs token
     */
    afterColon<T extends boolean, U extends boolean>(
        token: string,
        inx: number,
        num: T,
        bool?: U
    ): T extends true
        ? number
        : U extends true
        ? boolean
        : string {
        var count: any
        if (token) {
            // change 1 to 2 if you get any whitespace
            token = token.slice(inx + 1)
            switch (true) {
                // conv to num
                case num && !bool: {
                    count = parseInt(token)
                } break
                // conv to string
                case !num && !bool: {
                    count = token.trim()
                } break
                // conv to boolean
                default:
                    const pattern = /true/
                    count = pattern.test(token)
            }
        }
        return count
    }

    Search(token: string, _search: string): number {
        var count = -1
        for (let i = 0; i < token.length; i++) {
            if (token[i] == _search) {
                count = i
            }
        }
        return count
    }

    // // note: if you were to use the experimental your map value must string|string[]
    // GetTeam(token: string[], teamRed: boolean): Map<string, string[]> {
    //     // remove any empty space
    //     token = token
    //         .map(item => {
    //             item.trim();
    //             return item.toLowerCase()
    //         })
    //         .filter(item => item.length > 0)
    //     const matchB = /team\sblue/
    //     const matchR = /team\sred/
    //     var m = new Map<string, string[]>([])

    //     var team = []
    //     var TeamName = [] // either red or blue as per the boolean
    //     var playerName: any
    //     switch (true) {
    //         case teamRed:
    //             for (let i = 0; i < token.length; i++) {
    //                 if (matchR.test(token[i])) {
    //                     team.push(token[i])
    //                 }
    //             }
    //             if (team.length > 1) {
    //                 for (let i = 0; i < team.length; i++) {
    //                     TeamName.push([team[i]])
    //                 }

    //                 playerName = TeamName.join(' ').split(' ').filter((e) => { return e != 'team' && e != 'red' })
    //                 for (let i = 0; i < playerName.length; i++) {

    //                     m.set("TEAM RED", playerName)
    //                 }

    //             }
    //             // experimental
    //             // else {
    //             //     var ___teamBlue = team.join(' ').split(' ')
    //             //     var ____name = ___teamBlue.slice(0, 1).join('').toUpperCase()
    //             //     m.set("TEAM BLUE", ____name)
    //             // }

    //             break
    //         case !teamRed:
    //             for (let i = 0; i < token.length; i++) {
    //                 if (matchB.test(token[i])) {
    //                     team.push(token[i])
    //                 }
    //             }
    //             if (team.length > 1) {

    //                 for (let i = 0; i < team.length; i++) {
    //                     TeamName.push([team[i]])
    //                 }

    //                 playerName = TeamName.join(' ').split(' ').filter((e) => { return e != 'team' && e != 'blue' })
    //                 for (let i = 0; i < playerName.length; i++) {

    //                     m.set("TEAM BLUE", playerName)
    //                 }

    //             }

    //             // keep it here in-case the team.length<1
    //             // else {
    //             //     var ___teamBlue = team.join(' ').split(' ')
    //             //     var ____name = ___teamBlue.slice(0, 1).join('').toUpperCase()
    //             //     m.set("TEAM RED", ____name)
    //             // }

    //             break
    //         default:
    //             console.error("Not found")
    //     }
    //     return m
    // }
    /**
     * 
     * @param token any token that has that has 2 white space [example: XX X1 X2]
     * @param pattern that matches the first letter [example: XX X1 X2, pattern:/XX[\s\S\w]/]
     * @param key to save the sheets
     * @param remove1 to remove value after first whitespace
     * @param remove2 to remove value after seconds whitespace
     * @returns collection of main data
     * @example input: [fromXserver INDIA PM: MODI] output: map[INDIA] has {PM: MODI}
     * @note if the token and pattern doesnt matches it would send the same token back
     */
    UpdateSheets(_token: string, pattern: RegExp, key: string, remove1: string, remove2: string): Map<string, Map<string, string>> {
        if (pattern.test(_token)) {
            _token = _token.split(' ').filter(e => e != remove1 && e != remove2).join(' ').toString()
        }
        var result = _token.split(':')

        var updateSheetBoard = new Map<string, Map<string, string>>()
        var getUpdate = new Map<string, string>()
        getUpdate.set(result[0], result[1])
        updateSheetBoard.set(key, getUpdate)
        return updateSheetBoard
    }

    /**
     * 
     * @param token any token that has that has 2 white space [example: XX X1 X2]
     * @param pattern that matches the first letter [example: XX X1 X2, pattern:/XX[\s\S\w]/]
     * @param key to save the sheets
     * @param remove1 to remove value after first whitespace
     * @param remove2 to remove value after seconds whitespace
     * @returns collection of main data
     * @example input: [fromXserver BLUE XXX: BLUE] output: map[XXX] has {XXX: BLUE}
     * @note if the token and pattern doesnt matches it would send the same token back
     */
    GetpNickATeam(_token: string, pattern: RegExp, remove1: string): Map<string, Map<string, string>> {
        if (pattern.test(_token)) {
            _token = _token.split(' ').filter(e => e != remove1).join(' ').toString()
        }

        var result = _token.split(':')

        var updateSheetBoard = new Map<string, Map<string, string>>()
        var getUpdate = new Map<string, string>()
        getUpdate.set(result[0].split(' ')[0].trim(), result[1].trim())
        updateSheetBoard.set(result[0].split(' ')[1].trim(), getUpdate)
        return updateSheetBoard
    }

    GetpTeam(_token: string, pattern: RegExp, remove1: string): Map<string, Map<string, string>> {
        if (pattern.test(_token)) {
            _token = _token.split(' ').filter(e => e != remove1).join(' ').toString()
        }

        var result = _token.split(':')

        var updateSheetBoard = new Map<string, Map<string, string>>()
        var getUpdate = new Map<string, string>()

        getUpdate.set(result[0].split(' ')[1].trim(), result[0].split(' ')[0])

        updateSheetBoard.set(result[0].split(' ')[1].trim(), getUpdate)
        return updateSheetBoard
    }
    UpdateFrameSheet(token: string, pattern: RegExp, remove: string): Map<string, string> {
        var m2 = new Map<string, string>()
        if (pattern.test(token)) {
            var temp = token.split(' ').filter(e => e != remove).join(' ').split(' ')
            m2.set(temp[0], temp[2])
        }
        return m2
    }
    UpdateCheatSheet(token: string, pattern: RegExp, key: string, remove: string): Map<string, string> {
        var m2 = new Map<string, string>()
        if (pattern.test(token)) {
            var temp = token.split(' ').filter(e => e != remove).join(' ').split(' ')
            m2.set(key, temp[2])
        }
        return m2
    }
}