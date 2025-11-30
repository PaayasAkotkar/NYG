import { Component, Input, OnChanges, OnInit, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { SettingsPatternComponent } from "../settings-pattern/settings-pattern.component";
import { OcsInputComponent } from "../ocs-input/ocs-input.component";
import { OcsSelectComponent } from "../ocs-select/ocs-select.component";
import { PowerSetttingsPatternComponent } from "../power-setttings-pattern/power-setttings-pattern.component";
import { OcsBtn2Component } from "../../setup/ocs-btn-2/ocs-btn-2.component";
import { OcsCounterComponent } from "../ocs-counter/ocs-counter.component";
import { OcsHeaderComponent } from "../../setup/ocs-header/ocs-header.component";
import { HAND } from '../../helpers/hand';
import { BehaviorSubject, filter, Subscription, timer } from 'rxjs';
import { OCSFieldFormat } from '../../helpers/patterns';
import { FieldComponent } from '../../field/field/field.component';
// max limit for naming:
// 8
// error system:
// if the player sets the name over max limit=> than we can remove letters - overhead
// no special chars are allowed so if they use _ @ than remove that and alert them that we have done this you can go back and create new name too
@Component({
  selector: 'app-create-room',
  imports: [FieldComponent, ReactiveFormsModule, SettingsPatternComponent, OcsInputComponent, OcsSelectComponent, PowerSetttingsPatternComponent, OcsBtn2Component, OcsCounterComponent, OcsHeaderComponent],
  templateUrl: './create-room.component.html',
  styleUrls: ['./touchup/create-room.scss', './touchup/container.scss',
    './touchup/stepper.scss', './touchup/color.scss',
    './touchup/animation.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated
})
export class CreateRoomComponent implements OnInit, OnChanges {

  constructor(private FB: FormBuilder, private room: RoomService, private router: Router) { }

  @Input({ required: true }) NYGentertainmentCollection: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) NYGproceed: boolean = false
  @Input({ required: true }) NYGsportsCollection: WritableSignal<OCSFieldFormat[]> = signal([])

  createRoomForm = new FormBuilder().group({
    CreateToken: ['',],
    RoomLimitToken: [0,],
    // CreateNickNameToken: ['',],
    FriendZoneToken: [false,], // if the owner of the room wish to play with his friend
    SetupTossToken: [false,], // if the owner of the room wish to set the toss system 
    PrivateRoomToken: [false,], // if the owner of the room wish to create a private room
    StarterToken: [false,], // affects on the power
    ReverseToken: [false], // if the owner wants to set the a1-a1 or a1-b1 pattern
    CategoryToken: [' ',],
    BookToken: [''],
    NexusToken: [false],
    TagToken: [false],
    RewindToken: [false],
    DrawToken: [false],
    FreezeToken: [false],
    CovertToken: [false],
    BetToken: [false],
  })

  // gtime: [5s,10s]
  GTime: number = 10 // game time
  // dtime: [5s,2m]
  DTime: number = 2 // dicussion time minute
  DTimes: number = 0 // dicussion time seconds

  DiscussionTime = `${this.DTime} : ${this.DTimes}s`
  GameTime = `${this.GTime}s`
  Gcount____ = signal<number>(this.GTime)
  Dcount____ = signal<number>(this.DTimes)
  isEntertainment: boolean = false
  Friend: boolean = false

  getcreateRoomToken = String(this.createRoomForm.get('CreateToken')?.value)
  getRoomLimitToken = parseInt(String(this.createRoomForm.get('RoomLimitToken')?.value))
  getCategoryToken = String(this.createRoomForm.get('CategoryToken')?.value)
  getSetupTossToken = Boolean(this.createRoomForm.get('SetupTossToken')?.value)
  getPrivateRoomToken = Boolean(this.createRoomForm.get('PrivateRoomToken')?.value)
  getStarterToken = Boolean(this.createRoomForm.get('StarterToken')?.value)
  getReverseToken = Boolean(this.createRoomForm.get('ReverseToken')?.value)
  getCFreezeToken = Boolean(this.createRoomForm.get('FreezeToken')?.value)
  getCTagToken = Boolean(this.createRoomForm.get('TagToken')?.value)
  getCRewindToken = Boolean(this.createRoomForm.get('RewindToken')?.value)
  getCDrawToken = Boolean(this.createRoomForm.get('DrawToken')?.value)
  getCNeuxsToken = Boolean(this.createRoomForm.get('NexusToken')?.value)
  getCCovertToken = Boolean(this.createRoomForm.get('CovertToken')?.value)
  getCBetToken = Boolean(this.createRoomForm.get('BetToken')?.value)
  getFriendZoneToken = Boolean(this.createRoomForm.get('FriendZoneToken')?.value)
  getBookToken = String(this.createRoomForm.get('BookToken')?.value)
  // getCNickName: string = String(this.createRoomForm.get('CreateNickNameToken')?.value)

  ENavigation: 'CONTROLS' | 'SET' | 'BOOK' | 'POWERS' = 'SET'
  Entertainment: boolean = false
  navigationValidation: BehaviorSubject<string> = new BehaviorSubject('')
  OCSmode: OCSFieldFormat[] = [
    { OCStemplate: 'ONE on ONE', OCSformControl: 2, OCSToolTip: "" },
    { OCStemplate: 'TWO on TWO', OCSformControl: 4, OCSToolTip: "" },
  ]

  OCScategory: OCSFieldFormat[] = [
    { OCStemplate: 'SPORTS', OCSformControl: 'SPORTS', OCSToolTip: "" },
    { OCStemplate: 'ENTERTAINMENT', OCSformControl: 'ENTERTAINMENT', OCSToolTip: "" },
  ]

  OCSfields: OCSFieldFormat[] = [
    { OCStemplate: 'INTERNATIONAL', OCSformControl: 'INTERNATIONAL', OCSToolTip: "" },
    { OCStemplate: 'NATIONAL', OCSformControl: 'NATIONAL', OCSToolTip: "" },
    { OCStemplate: 'DOMESTIC', OCSformControl: 'DOMESTIC', OCSToolTip: "" },
  ]

  selectedBook = new HAND.Queue()
  selectedPowerKeys = new HAND.Queue()

  powerKeys: string[] = []
  collectionKeys: string[] = []

  powers: OCSFieldFormat[] = [
    { OCStemplate: 'NEXUS', OCSformControl: 'NexusToken', OCSToolTip: "" },
    { OCStemplate: 'TAG', OCSformControl: 'TagToken', OCSToolTip: "" },
    { OCStemplate: 'REWIND', OCSformControl: 'RewindToken', OCSToolTip: "" },
    { OCStemplate: 'DRAW', OCSformControl: 'DrawToken', OCSToolTip: "" },
    { OCStemplate: 'FREEZE', OCSformControl: 'FreezeToken', OCSToolTip: "" },
    { OCStemplate: 'COVERT', OCSformControl: 'CovertToken', OCSToolTip: "" },
    { OCStemplate: 'BET', OCSformControl: 'BetToken', OCSToolTip: "" },
  ]

  oneVonePowers: OCSFieldFormat[] = [
    { OCStemplate: 'NEXUS', OCSformControl: 'NexusToken', OCSToolTip: "" },
    { OCStemplate: 'REWIND', OCSformControl: 'RewindToken', OCSToolTip: "" },
    { OCStemplate: 'FREEZE', OCSformControl: 'FreezeToken', OCSToolTip: "" },
    { OCStemplate: 'COVERT', OCSformControl: 'CovertToken', OCSToolTip: "" },
    { OCStemplate: 'BET', OCSformControl: 'BetToken', OCSToolTip: "" },
  ]

  twoVtwoPowers: OCSFieldFormat[] = [
    { OCStemplate: 'NEXUS', OCSformControl: 'NexusToken', OCSToolTip: "" },
    { OCStemplate: 'TAG', OCSformControl: 'TagToken', OCSToolTip: "" },
    { OCStemplate: 'REWIND', OCSformControl: 'RewindToken', OCSToolTip: "" },
    { OCStemplate: 'DRAW', OCSformControl: 'DrawToken', OCSToolTip: "" },
    { OCStemplate: 'FREEZE', OCSformControl: 'FreezeToken', OCSToolTip: "" },
    { OCStemplate: 'COVERT', OCSformControl: 'CovertToken', OCSToolTip: "" },
    { OCStemplate: 'BET', OCSformControl: 'BetToken', OCSToolTip: "" },
  ]
  powerToolTip: string = 'abilities to gain advantage'
  bookToolTip: string = 'style of game'
  roomLimitToolTip: string = "match type"
  categoryToolTip: string = "genre of the game"
  nameToolTip: string = "name of the room"
  nicknameToolTip: string = "name to display assoicated players"
  tokens: OCSFieldFormat[] = [
    { OCStemplate: 'INVITE', OCSformControl: 'FriendZoneToken', OCSToolTip: "teamup with friend" },
    { OCStemplate: 'TOSS', OCSformControl: 'SetupTossToken', OCSToolTip: "after each round toss for the dictionary" },
    { OCStemplate: 'PRIVATE', OCSformControl: 'PrivateRoomToken', OCSToolTip: "accessible via code" },
    { OCStemplate: 'STARTER', OCSformControl: 'StarterToken', OCSToolTip: "winner initiates the dictionary and challenge" },
    { OCStemplate: 'REVERSE', OCSformControl: 'ReverseToken', OCSToolTip: "one by one initiates" },
  ]

  // signal to start the bevahior
  signalPower: boolean = false
  NYGbookerror: string = ""
  isNYGbookError: boolean = false
  Subscription: Subscription | undefined

  isNYGnameError: boolean = false
  isNYGnicknameError: boolean = false
  NDEFAULT = "the name is not valid remove"
  SDEFAULT = "the filed must not left empty"
  NYGnameError: string = ""
  NYGnicknamenameError: string = ""
  NYGcategoryError: string
  NYGtypeError: string
  isNYGcategoryError: boolean = false
  isNYGtypeError: boolean = false
  isNYGpowerError: boolean = false
  NYGpowerError: string
  proccedNavigation: boolean = false
  PDEFAULT = "MUST SELECT 4 POWERs"
  MatchType: string = ""

  NYGgameTimeToolTip = "kickoff time"
  NYGdecisionTimeToolTip = "strategic discussion time"

  powerTriggered: BehaviorSubject<string[]> = new BehaviorSubject<string[]>([])
  click(token: { click: boolean, formControl: string }) {
    this.createRoomForm.get(token.formControl)?.setValue(token.click, { emitEvent: false })
  }

  ngOnInit(): void {
    this.NYGpowerError = this.PDEFAULT
    this.NYGtypeError = this.SDEFAULT
    this.NYGcategoryError = this.SDEFAULT
    this.MatchType = '2'
    this.powers = this.oneVonePowers
    this.createRoomForm.get('CategoryToken')?.setValue('SPORTS', { emitEvent: false })
    this.createRoomForm.get('RoomLimitToken')?.setValue(2, { emitEvent: false })
    this.powerTriggered.subscribe({
      next: (token) => {
        if (this.signalPower) {
          if (token.length < 4) {
            var cal = 4 - this.powerKeys.length
            this.isNYGpowerError = true
            this.NYGpowerError = `MUST SELECT ${cal.toString()}  MORE POWERS`
          } else {
            this.signalPower = false
            this.isNYGpowerError = false
          }
        }
      }
    })

    this.NYGnameError = "the name is not valid  "
    const restriction = /[^a-zA-Z]$/;
    this.createRoomForm.get('CreateToken')?.valueChanges.pipe(filter(token => token != undefined)).subscribe({
      next: (token) => {
        if (restriction.test(token)) {
          for (let i = 0; i < token.length; i++) {
            if (restriction.test(token[i])) {
              if (!this.NYGnameError.includes(token[i])) {
                this.isNYGnameError = true
                this.NYGnameError += token[i]
              }
            }
          }
          token = token.replace(restriction, "")
          setTimeout(() => {
            this.isNYGnameError = false
            this.NYGnameError = this.NDEFAULT
          }, 2500)
          this.createRoomForm.get('CreateToken')?.setValue(token, { emitEvent: false })
        }
        else if (token.length > 8) {
          this.isNYGnameError = true
          this.NYGnameError = "REACHED MAX LENGTH 8"
          this.createRoomForm.get('CreateToken')?.setValue(token.substring(0, 8))
          setTimeout(() => {
            this.isNYGnameError = false
          }, 3500)
        }

      }
    })

    this.FillControls()
    this.Subscription = this.createRoomForm.get('CategoryToken')?.valueChanges.subscribe({
      next: (token) => {
        if (token == "ENTERTAINMENT") {
          this.Entertainment = true
        } else if (token == "SPORTS") {
          this.Entertainment = false
        }
      }
    })
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGproceed']) {
      if (changes['NYGproceed'].currentValue == true)
        this.FillControls()
    }
  }

  GAdd() {
    if (this.Gcount____() < 10) {
      this.Gcount____.update(v => v + 1)
      this.GTime = this.Gcount____()
    }
    this.GameTime = `${this.GTime}s`
  }

  GSub() {
    if (this.Gcount____() > 3) {
      this.Gcount____.update(v => v - 1)
      this.GTime = this.Gcount____()
    }
    this.GameTime = `${this.GTime}s`
  }

  DAdd() {
    if (this.Dcount____() < 60) {
      this.Dcount____.update(v => v + 1)
      this.DTimes = this.Dcount____()

      this.DiscussionTime = `${this.DTime} : ${this.DTimes}s`
    }
  }

  DSub() {
    if (this.Dcount____() > 0) {
      this.Dcount____.update(v => v - 1)
      this.DTimes = this.Dcount____()
      this.DiscussionTime = `${this.DTime} : ${this.DTimes}s`
    }
  }

  FillControls() {
    this.NYGsportsCollection().forEach(item => {
      this.createRoomForm.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
    })
    this.NYGentertainmentCollection().forEach(item => {
      this.createRoomForm.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
    })
  }

  toBook() {
    this.ENavigation = 'BOOK'
  }

  toSet() {
    this.ENavigation = 'SET'
  }
  toControls() {
    this.ENavigation = 'CONTROLS'
  }
  toPowers() {
    this.ENavigation = 'POWERS'
  }
  GetMatchType(token: string) {
    this.MatchType = token
    if (this.MatchType == '2') {
      this.powers = this.oneVonePowers
    } else {
      this.powers = this.twoVtwoPowers
    }
  }

  Update() {
    const formValue: any = this.createRoomForm.value;
    const value = Object.keys(formValue).filter(key => {

      var isBooks = false

      for (let i of this.NYGsportsCollection()) {
        if (i.OCStemplate == key) {
          isBooks = true
        }
      }

      for (let i of this.NYGentertainmentCollection()) {
        if (i.OCStemplate == key) {
          isBooks = true
        }
      }

      var res = ''
      if (isBooks) {
        if (formValue[key] === true) {
          res = key
        }
      }
      return res
    }).join(' ')
    this.selectedBook.enqueue(value)
  }

  // PowerToken get's the name of the newly clicked formvalue and manipulate the click system 
  PowerToken(token: string) {
    this.powerKeys.push(token)
    const len = this.powerKeys.length

    if (len > 4) {
      const newly = this.powerKeys[len - 1]

      var temp
      if (newly) {
        temp = this.powerKeys[this.powerKeys.indexOf(newly) - 1]
      }
      const old = temp

      // this condition prevents the newly and old value to get false  
      if (old && this.createRoomForm.get(newly)?.value == true) {
        this.powerKeys = this.powerKeys.filter(e => e != old)
        this.powerKeys.forEach(e => {
          if (this.createRoomForm.get(e)?.value == true) {
            this.createRoomForm.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }

    // fetch only active power
    for (let power of this.powerKeys) {
      var active = this.createRoomForm.get(power)?.value
      if (!active) {
        this.powerKeys = this.powerKeys.filter(p => p != power)
      }
    }
    this.powerTriggered.next(this.powerKeys)
  }

  CollectionToken(token: string) {
    this.collectionKeys.push(token)

    const len = this.collectionKeys.length
    if (len > 1) {
      const newly = this.collectionKeys[len - 1]

      var temp
      if (newly) {
        temp = this.collectionKeys[this.collectionKeys.indexOf(newly) - 1]

      }
      const old = temp

      // this condition prevents the newly and old value to get false  
      if (old && this.createRoomForm.get(newly)?.value == true) {
        this.collectionKeys = this.collectionKeys.filter(e => e != old)
        this.collectionKeys.forEach(e => {
          if (this.createRoomForm.get(e)?.value == true) {
            this.createRoomForm.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }

    }
  }

  RoomCreate(e: Event) {
    e.preventDefault()

    // this.collectionKeys.forEach((_, i) => {
    //   if (this.collectionKeys[i] == this.collectionKeys[i - 1]) {

    //   }
    // })

    this.getFriendZoneToken = Boolean(this.createRoomForm.get('FriendZoneToken')?.value)
    this.getSetupTossToken = Boolean(this.createRoomForm.get('SetupTossToken')?.value)
    this.getPrivateRoomToken = Boolean(this.createRoomForm.get('PrivateRoomToken')?.value)
    this.getStarterToken = Boolean(this.createRoomForm.get('StarterToken')?.value)
    this.getReverseToken = Boolean(this.createRoomForm.get('ReverseToken')?.value)
    this.getBookToken = this.collectionKeys[0]
    // this.getCNickName = String(this.createRoomForm.get('CreateNickNameToken')?.value)
    this.getCategoryToken = String(this.createRoomForm.get('CategoryToken')?.value)
    this.getcreateRoomToken = String(this.createRoomForm.get('CreateToken')?.value)
    this.getRoomLimitToken = parseInt(String(this.createRoomForm.get('RoomLimitToken')?.value))
    this.getCFreezeToken = Boolean(this.createRoomForm.get('FreezeToken')?.value)
    this.getCTagToken = Boolean(this.createRoomForm.get('TagToken')?.value)
    this.getCRewindToken = Boolean(this.createRoomForm.get('RewindToken')?.value)
    this.getCDrawToken = Boolean(this.createRoomForm.get('DrawToken')?.value)
    this.getCNeuxsToken = Boolean(this.createRoomForm.get('NexusToken')?.value)
    this.getCCovertToken = Boolean(this.createRoomForm.get('CovertToken')?.value)
    this.getCBetToken = Boolean(this.createRoomForm.get('BetToken')?.value)
    const getGameTime = this.GTime
    const getDecisionTime = this.DTimes
    this.NYGnameError = 'name is empty'
    this.NYGnicknamenameError = 'nickname is empty'

    // if (this.getCNickName != "" && this.getCNickName != undefined) {
    //   this.isNYGnicknameError = false
    // }
    if (this.getcreateRoomToken != "" && this.getcreateRoomToken != undefined) {
      this.isNYGnameError = false
    }
    if (this.getcreateRoomToken == "" || this.getcreateRoomToken == undefined) {
      this.isNYGnicknameError = true
      this.isNYGnameError = true
      this.ENavigation = 'SET'
    }
    // if (this.getCNickName == "" || this.getCNickName == undefined) {
    //   this.isNYGnicknameError = true
    //   this.ENavigation = 'SET'
    // }
    if (this.getcreateRoomToken == "" || this.getcreateRoomToken == "") {
      this.isNYGnameError = true
      this.ENavigation = 'SET'
    }
    if (this.powerKeys.length < 4) {
      this.signalPower = true
      var cal = 4 - this.powerKeys.length
      this.isNYGpowerError = true
      this.NYGpowerError = `MUST SELECT ${cal.toString()}  MORE POWERS`
    } else {
      this.signalPower = false
      this.isNYGpowerError = false
    }
    if (this.collectionKeys[0] == undefined) {
      this.NYGbookerror = `SELECT AT LEAST ONE BOOK`
      this.isNYGbookError = true
    }

    if (!this.isNYGnameError && !this.isNYGnicknameError) {
      if (this.collectionKeys[0] == undefined) {
        this.NYGbookerror = `SELECT AT LEAST ONE BOOK`
        this.isNYGbookError = true
        this.ENavigation = 'BOOK'
      } else {
        this.isNYGbookError = false
      }
      if (this.powerKeys.length < 4) {
        var cal = 4 - this.powerKeys.length
        this.isNYGpowerError = true
        this.NYGpowerError = `MUST SELECT ${cal.toString()}  MORE POWERS`
        this.ENavigation = 'POWERS'
      } else {
        this.isNYGpowerError = false
      }
    }

    if (!this.isNYGnameError && !this.isNYGnicknameError && !this.isNYGpowerError && !this.isNYGbookError) {
      if (this.getFriendZoneToken) {
        this.room.CreateRoom(this.getcreateRoomToken,
          this.getRoomLimitToken,
          true, this.getSetupTossToken, this.getStarterToken,
          this.getReverseToken,
          this.getPrivateRoomToken, this.getCategoryToken, "",
          this.getBookToken,
          this.getCDrawToken, this.getCNeuxsToken,
          this.getCTagToken, this.getCRewindToken,
          this.getCFreezeToken,
          this.getCCovertToken, this.getCBetToken, getGameTime, getDecisionTime
        )
        this.router.navigate(['/Game-Hall'], { skipLocationChange: true })

      } else {
        this.room.CreateRoom(this.getcreateRoomToken,
          this.getRoomLimitToken,
          false, this.getSetupTossToken, this.getStarterToken,
          this.getReverseToken,
          this.getPrivateRoomToken, this.getCategoryToken,
          "", this.getBookToken,
          this.getCDrawToken, this.getCNeuxsToken,
          this.getCTagToken, this.getCRewindToken,
          this.getCFreezeToken,
          this.getCCovertToken, this.getCBetToken, getGameTime, getDecisionTime
        )
        this.Friend = false
      }
      this.router.navigate(['/Game-Hall'], { skipLocationChange: false })

    }

  }

}