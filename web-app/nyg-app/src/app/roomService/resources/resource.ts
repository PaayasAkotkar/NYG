import { filter, Observable } from "rxjs";
import { RoomFields } from "../vouchers/vouchers";
import { SessionService } from "../../storage/session/session.service";
import { LocalService } from "../../storage/local/local.service";
import { NgZone } from "@angular/core";
import { CookieService } from "../../storage/cookie/cookie.service";


interface signedUp {
    leaveRoom: boolean
    createRoom: boolean // if true
    joinRoom: boolean // if true
    friend: boolean
    set: boolean // if true than set nickname
    setToss: boolean // if the toss session to keep on or off
    starter: boolean // whether they want to make use of fullest
    reverse: boolean // whether they want to add a1-a1 or a1-b1
    privateRoom: boolean // whether they want to create a private room
    changeSettings: boolean
    roomCreated: boolean
    nexusPower: boolean
    tagPower: boolean
    rewindPower: boolean
    freezePower: boolean
    covertPower: boolean
    betPower: boolean
    drawPower: boolean
    to: string       // then join to this
    roomName: string   // then create this room
    id: string

    category: string
    book: string
    field: string

    gameTime: number
    decisionTime: number

    code: string
    // nickname: string      // to display the name to all connected players in room

    roomCapacity: number
    clash: boolean
}

// note: only when the player is in the lobby
interface LogOut {
    roomName: string
    id: string
    loggout: boolean
}

interface Guess {
    redScoreCount: number
    blueScoreCount: number
    set: number
    round: number
    chances: number
    time: number
    onFire: number

    roomName: string
    guessToken: string
    challengeToken: string
    HeadTails: string
    dictionaryToken: string
    betOn: string

    /**end of room data*/
    start: boolean // start the game
    challengeDiscussion: boolean
    dictionarySession: boolean
    session: boolean
    TossSession: boolean
    challengeSet: boolean
    power: boolean
    freeze: boolean
    nexus: boolean
    rewind: boolean
    draw: boolean
    tag: boolean
    covert: boolean
    unfreeze: boolean
    bet: boolean
    drawAccept: boolean
    drawSession: boolean
    betSession: boolean
    timeUp: boolean
    discussionTimeUp: boolean
    clash: boolean
    finalBoss: boolean
    lastDance: boolean
    spectate: boolean
    myView: ISpectate
}
export interface ISpectate {
    scrollPos: number // current scroll pos
    onClick: string   // current clicked event
    word: string      // current word written on the guest
    scrollHappend: boolean
}
interface Profile {
    name: string
    set: boolean
    id: string
}

export class RoomResource extends RoomFields {
    constructor(private _cookie: CookieService, private _session: SessionService, private _local: LocalService, private run: NgZone) {
        super(_session, _local)
    }

    connect() {
        if (this.Conn && this.Conn.readyState === WebSocket.CONNECTING) {
            return
        }
        // saves the id in the local storage of the web so that we dont have to generate new id each time
        // note: this is valid till the one tab 
        // meaning that one client can play with 3 different ids for its 3 open tab
        var ___id = this._session.getItem("session")
        if (!___id) {
            ___id = this.id
        }
        var _id = ___id
        this.getId = _id

        const URL = `ws://localhost:3000/ws/${_id}/${this.roomName}`
        if (!this.Conn || this.Conn.readyState === WebSocket.CLOSED) {
            this.Conn = new WebSocket(URL)
        }
        this.Conn.onopen = (e) => { }
        this.Conn.onmessage = (token) => {
            this.run.run(() => {
                this.receiveTokens.next(token.data)
            })
        }
        this.Conn.onerror = (e) => {
            console.error("ERROR___SCHEMA_____ WEBSOCKET 🍋", e)
        }
        this.Conn.onclose = () => {
        }
    }

    Online(): boolean {
        return this.Conn.OPEN == this.Conn.readyState
    }

    Leave(roomName: string) {
        const token: LogOut = { roomName: roomName.toLowerCase(), id: this.getId, loggout: true }
        const sendToken = JSON.stringify(token)
        this.Conn.send(sendToken)
    }

    Clash(coverPower: boolean, nexusPower: boolean, rewindPower: boolean, freezePower: boolean, betPower: boolean) {
        var name = this._cookie.getCookie('PROFILE')

        var siup: signedUp = {
            joinRoom: false, to: "", leaveRoom: false, createRoom: true,
            roomName: "", roomCapacity: 0, id: this.getId,
            category: "", book: "", field: "", roomCreated: false,
            friend: false, code: "", set: true, changeSettings: false,
            setToss: false, starter: false, reverse: false, privateRoom: false,
            drawPower: false, nexusPower: nexusPower, tagPower: false, rewindPower: rewindPower, freezePower: freezePower
            , covertPower: coverPower, betPower: betPower, gameTime: -1, decisionTime: -1,
            clash: true,
        }
        this.signUPDetails = siup
        const token = JSON.stringify(siup)
        this.Conn.send(token)
    }
    joinRoom(to: string,
        friend: boolean, code: string,): void {

        var siup: signedUp = {
            joinRoom: true, to: to, leaveRoom: false, createRoom: false,
            roomName: "", roomCapacity: this.signUPDetails.roomCapacity,
            category: "", id: this.getId,
            setToss: false, reverse: false, roomCreated: false,
            field: "", privateRoom: false, drawPower: false, nexusPower: false, tagPower: false, rewindPower: false, freezePower: false,
            book: "", starter: false, changeSettings: false,
            friend: friend, code: "", set: true,
            covertPower: false, betPower: false, decisionTime: -1, gameTime: -1,
            clash: false
        }
        if (friend) {
            var siup: signedUp = {
                joinRoom: true, to: to, leaveRoom: false, createRoom: false,
                roomName: "", roomCapacity: this.signUPDetails.roomCapacity,
                category: "", id: this.getId, roomCreated: false,
                book: "", privateRoom: false, reverse: false, drawPower: false, nexusPower: false, tagPower: false, rewindPower: false, freezePower: false,
                setToss: false, starter: false, changeSettings: false,
                field: "", friend: friend, code: code, set: true,
                covertPower: false, betPower: false, gameTime: -1, decisionTime: -1,
                clash: false
            }
        }
        this.signUPDetails = siup
        const token = JSON.stringify(siup)
        this.Conn.send(token)
    }

    /**
     *  @RETURNS rooms name, live connection
    */
    Rooms(): Observable<string> {
        return this.receiveTokens.asObservable()
    }

    CreateRoom(roomName: string, limit: number, friend: boolean, setToss: boolean, starter: boolean, reverse: boolean, privateRoom: boolean, category: string, field: string, book: string,
        drawPower: boolean, nexusPower: boolean, tagPower: boolean, rewindPower: boolean, freezePower: boolean, coverPower: boolean, betPower: boolean,
        gameTime: number, decisionTime: number
    ): void {
        var siup: signedUp = {
            joinRoom: false, to: "", leaveRoom: false, createRoom: true,
            roomName: roomName, roomCapacity: limit, id: this.getId,
            category: category.toLowerCase(), book: book.toLowerCase(), field: field, roomCreated: false,
            friend: friend, code: "", set: true, changeSettings: false,
            setToss: setToss, starter: starter, reverse: reverse, privateRoom: privateRoom,
            drawPower: drawPower, nexusPower: nexusPower, tagPower: tagPower, rewindPower: rewindPower, freezePower: freezePower
            , covertPower: coverPower, betPower: betPower, gameTime: gameTime, decisionTime: decisionTime,
            clash: false,
        }
        this.signUPDetails = siup
        const token = JSON.stringify(siup)
        this.Conn.send(token)
    }

    ChangeRoomSettings(starter: boolean = false, reverse: boolean = false, toss: boolean = false, book: string, drawPower: boolean, nexusPower: boolean, tagPower: boolean,
        rewindPower: boolean, freezePower: boolean, coverPower: boolean, betPower: boolean, gameTime: number, decisionTime: number) {
        var siup: signedUp = {
            joinRoom: false, to: "", leaveRoom: false, createRoom: true,
            roomName: this.signUPDetails.roomName, roomCapacity: -1, id: this.getId,
            field: "", roomCreated: true,
            friend: false, code: "", set: false, changeSettings: true,
            setToss: toss, starter: starter, reverse: reverse, privateRoom: false, book: book, category: "",
            drawPower: drawPower, nexusPower: nexusPower, tagPower: tagPower, rewindPower: rewindPower, freezePower: freezePower
            , covertPower: coverPower, betPower: betPower, gameTime: gameTime, decisionTime: decisionTime, clash: false

        }
        const token = JSON.stringify(siup)
        this.Conn.send(token)
    }

    // hold
    // LeaveRoom() {
    //     const siup: signedUp = {
    //         joinRoom: false, to: "", leaveRoom: true, createRoom: false,
    //         roomName: this.roomName, // important else we may not able to find which room to eliminate
    //         roomCapacity: 0, category: "", subCategory: "", global: "", event: ""
    //     }
    //     this.signUPDetails = siup
    //     const token = JSON.stringify(siup)
    //     this.Conn.send(token)
    // }

    GuessToken(token: string, redscore: number,
        bluescore: number, to: string = "",
        round: number, set: number,
        chances: number, clash: boolean, time: number, onFire: number, finalBoss: boolean, lastDance: boolean) {

        var tokenGuess: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: token, roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: redscore, blueScoreCount: bluescore,
            dictionarySession: false, dictionaryToken: "", draw: false, tag: false,
            start: true, challengeDiscussion: false, session: true,
            TossSession: false, challengeSet: false, round: round, set: set, power: false, freeze: false, nexus: false, rewind: false,
            covert: false, unfreeze: false, bet: false,
            drawAccept: false, drawSession: false, betSession: false, betOn: "", clash: clash, chances: chances,
            time: time, onFire: onFire, finalBoss: finalBoss, lastDance: lastDance,
            spectate: false, myView: {
                onClick: "", word: "", scrollPos:
                    -1, scrollHappend: false

            }
        }
        const sendToken = JSON.stringify(tokenGuess)
        this.Conn.send(sendToken)
    }

    SetDictionary(token: string, to: string, clash: boolean,
        chances: number, finalBoss: boolean,
        lastDance: boolean) {

        var tokenDictionary: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: -1, blueScoreCount: -1,
            dictionarySession: true, dictionaryToken: token,
            draw: false, tag: false,
            start: false, challengeDiscussion: false, session: true,
            TossSession: false, challengeSet: false, round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, chances: chances,
            time: -1, onFire: -1, finalBoss: finalBoss,
            lastDance: lastDance, spectate: false,
            myView: {
                onClick: "", word: "",
                scrollPos: -1, scrollHappend: false,
            },
        }
        const sendToken = JSON.stringify(tokenDictionary)
        this.Conn.send(sendToken)
    }

    SetProfile(token: string) {
        var _token: Profile = {
            name: token, set: true, id: this.getId
        }
        const sendToken = JSON.stringify(_token)
        this.Conn.send(sendToken)
    }

    challengeToken(challenge: string, to: string, clash: boolean, chances: number, finalBoss: boolean, lastDance: boolean, round: number) {
        var tokenChallenge: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), HeadTails: "", challengeToken: challenge,
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "", draw: false, tag: false,
            session: true, start: false, challengeDiscussion: true,
            TossSession: false,
            challengeSet: true, round: round, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, chances: chances,
            time: -1, onFire: -1, finalBoss: finalBoss, lastDance: lastDance, spectate: false,
            myView: {
                onClick: "", word: "", scrollPos: -1, scrollHappend: false

            },

        }
        const sendToken = JSON.stringify(tokenChallenge)
        this.Conn.send(sendToken)

    }

    TossSession(HeadTails: string, friend: boolean, to: string, clash: boolean, chances: number, finalBoss: boolean, lastDance: boolean) {
        var TossDiscussion: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: HeadTails,
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            session: true, challengeDiscussion: false, start: false,
            TossSession: true, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false, bet: false,
            covert: false, unfreeze: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, chances: chances,
            time: -1, onFire: -1, finalBoss: finalBoss,
            lastDance: lastDance, spectate: false,
            myView: {
                onClick: "", word: "",
                scrollPos: -1, scrollHappend: false
            },
        }
        const sendToken = JSON.stringify(TossDiscussion)
        this.Conn.send(sendToken)

    }

    // todo set the room name param
    UseFreezePower(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: true, nexus: false, rewind: false
            , draw: false, tag: false, drawSession: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false,
            betSession: false, betOn: "", clash: clash,
            time: -1, onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: {
                onClick: "", word: "", scrollPos: -1, scrollHappend: false
            },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    UseNexusPower(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: true, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash,
            time: -1, onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    Unfreeze(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: true,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false, unfreeze: true,
            covert: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash,
            time: -1, onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    UseRewindPower(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: false, rewind: true
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "",
            clash: clash, time: -1, onFire: -1,
            finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }
    UseDrawPower(to: string) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: false, rewind: false
            , draw: true, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "",
            clash: false, time: -1, onFire: -1,
            finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }
    UseTagPower(to: string,) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: false, rewind: false
            , draw: false, tag: true,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "",
            clash: false, time: -1, onFire: -1,
            finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    UseCovertPower(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: true, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "",
            clash: clash, time: -1, onFire: -1,
            finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }
    UseBetPower(to: string, clash: boolean) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: true, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: true, drawAccept: false, drawSession: false,
            betSession: false, betOn: "",
            clash: clash, time: -1, onFire: -1,
            finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    SendDrawOffer(accept: boolean, to: string) {
        var tokenPower: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true,
            challengeDiscussion: false, start: false,
            TossSession: false,
            challengeSet: false,
            round: -1, set: -1,
            power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false,
            drawAccept: accept, drawSession: true,
            betSession: false, betOn: "", clash: false, time: -1,
            onFire: -1, finalBoss: false, lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }

        const sendToken = JSON.stringify(tokenPower)
        this.Conn.send(sendToken)
    }

    BetOn(to: string, bet: string, clash: boolean) {
        var tokenBet: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(),
            challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1,
            power: false, freeze: false,
            nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false,
            bet: false, drawAccept: false,
            drawSession: false,
            betSession: true,
            betOn: bet, clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
    }

    SendBetGuess(to: string, guess: string, clash: boolean,) {
        var tokenBet: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: true,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: true, betOn: guess, clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
    }

    GameTimeUp(roomName: string, clash: boolean, chances: number) {
        var tokenBet: Guess = {
            discussionTimeUp: false, timeUp: true,
            guessToken: "", roomName: roomName.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: true,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
    }

    DecisionTimeUp(roomName: string, clash: boolean, chances: number) {
        var tokenBet: Guess = {
            discussionTimeUp: true, timeUp: false,
            guessToken: "", roomName: roomName.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
    }

    FreezeTimeOut(roomName: string, clash: boolean) {
        const token: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: roomName.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: true,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: true, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(token)
        this.Conn.send(sendToken)
    }

    Spectate(roomname: string, isClash: boolean, isFinalBoss: boolean,
        myView: ISpectate) {

        const token: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: roomname.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: true, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: isClash, time: -1,
            onFire: -1, finalBoss: isFinalBoss,
            lastDance: false, spectate: true, myView: myView,
        }

        const sendToken = JSON.stringify(token)
        this.Conn.send(sendToken)
    }
    TextUpdate(roomname: string, isClash: boolean, isFinalBoss: boolean, myView: ISpectate,) {

        const token: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: roomname.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: true,
            TossSession: false, challengeSet: false,
            round: -1, set: -1, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: true, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: isClash, time: -1,
            onFire: -1, finalBoss: isFinalBoss,
            lastDance: false, spectate: true, myView: myView,
        }

        const sendToken = JSON.stringify(token)
        this.Conn.send(sendToken)
    }

    ChallengeTimeOut(roomname: string, clash: boolean, challengeToken: string, round: number, set: number) {
        var tokenBet: Guess = {
            discussionTimeUp: true, timeUp: false,
            guessToken: "", roomName: roomname.toLowerCase(), challengeToken: challengeToken, HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: true, start: false,
            TossSession: false, challengeSet: false,
            round: round, set: set, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
        console.log('timeout: ', sendToken)

    }
    // for toss coin timeout use the coin toss and send it
    TossTimeOut(roomname: string, clash: boolean, tossToken: string, round: number, set: number) {
        var tokenBet: Guess = {
            discussionTimeUp: true, timeUp: false,
            guessToken: "", roomName: roomname.toLowerCase(), challengeToken: "", HeadTails: tossToken,
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: true, challengeSet: false,
            round: round, set: set, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
        console.log('timeout: ', sendToken)

    }
    DictionaryTimeOut(roomname: string, clash: boolean, dictionaryToken: string, round: number, set: number) {
        var tokenBet: Guess = {
            discussionTimeUp: true, timeUp: false,
            guessToken: "", roomName: roomname.toLowerCase(), challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: true, dictionaryToken: dictionaryToken,
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: round, set: set, power: false, freeze: false, nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false, bet: false, drawAccept: false, drawSession: false,
            betSession: false, betOn: "", clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
        console.log('timeout: ', sendToken)
    }
    BetOnTimeOut(to: string, bet: string, clash: boolean, round: number, set: number) {
        var tokenBet: Guess = {
            discussionTimeUp: false, timeUp: false,
            guessToken: "", roomName: to.toLowerCase(),
            challengeToken: "", HeadTails: "",
            redScoreCount: 0, blueScoreCount: 0,
            dictionarySession: false, dictionaryToken: "",
            chances: -1,
            session: true, challengeDiscussion: false, start: false,
            TossSession: false, challengeSet: false,
            round: round, set: set,
            power: false, freeze: false,
            nexus: false, rewind: false
            , draw: false, tag: false,
            covert: false, unfreeze: false,
            bet: false, drawAccept: false,
            drawSession: false,
            betSession: true,
            betOn: bet, clash: clash, time: -1,
            onFire: -1, finalBoss: false,
            lastDance: false, spectate: false, myView: { onClick: "", word: "", scrollPos: -1, scrollHappend: false },
        }
        const sendToken = JSON.stringify(tokenBet)
        this.Conn.send(sendToken)
        console.log('timeout: ', sendToken)

    }
    LeaveRoom(room: string) {
        const token: LogOut = {
            roomName: room, id: this.id, loggout: true
        }
        const sendToken = JSON.stringify(token)
        this.Conn.send(sendToken)
    }
}


