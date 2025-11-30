import { AfterViewInit, Component, ElementRef, Input, NgZone, OnChanges, OnInit, Renderer2, signal, SimpleChanges, ViewChild, ViewEncapsulation, WritableSignal } from '@angular/core';
import { RoomService } from '../../roomService/room.service';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { SettingsPatternComponent } from "../settings-pattern/settings-pattern.component";
import { PowerSetttingsPatternComponent } from "../power-setttings-pattern/power-setttings-pattern.component";
import { OcsHeaderComponent } from "../../setup/ocs-header/ocs-header.component";
import { OcsCounterComponent } from "../ocs-counter/ocs-counter.component";
import { OcsBtn2Component } from "../../setup/ocs-btn-2/ocs-btn-2.component";
import { OCSFieldFormat } from '../../helpers/patterns';
import { FieldComponent } from '../../field/field/field.component';
import { RoomSettings } from '../../lobby/resources/lobbyResource';
import { BehaviorSubject } from 'rxjs';
import { RoomSettingsService } from './room-settings.service';

@Component({
  selector: 'app-room-settings',
  imports: [ReactiveFormsModule, SettingsPatternComponent, PowerSetttingsPatternComponent, FieldComponent, OcsHeaderComponent, OcsCounterComponent, OcsBtn2Component],
  templateUrl: './room-settings.component.html',
  styleUrls: ['./touchup/room-settings.scss', './touchup/color.scss', './touchup/powerS.scss'],
  encapsulation: ViewEncapsulation.Emulated
})
export class RoomSettingsComponent implements AfterViewInit, OnInit, OnChanges {
  @Input({ required: false }) Entertainment: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) nygFormControl: OCSFieldFormat[] = [
    { OCStemplate: 'TOSS', OCSformControl: 'SetupTossToken', OCSToolTip: "after each round toss for the dictionary" },
    { OCStemplate: 'STARTER', OCSformControl: 'StarterToken', OCSToolTip: "winner initiates the dictionary and challenge" },
    { OCStemplate: 'REVERSE', OCSformControl: 'ReverseToken', OCSToolTip: "one by one initiates" }
  ]

  @Input({ required: true }) NYGcollection: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) DTimes: number = 0 // dicussion time seconds
  @Input({ required: true }) GTime: number = 3 // game time
  @Input({ required: true }) setToss: boolean = false
  @Input({ required: true }) setReverse: boolean = false
  @Input({ required: true }) setStarter: boolean = false
  @Input({ required: true }) NYGroomSettings: RoomSettings

  @Input({ required: true }) Powers: Record<string, boolean> = {}
  @Input({ required: true }) NYGclient: boolean = false // if it is not the room owner
  @Input({ required: true }) Maxlength: number = 4

  @ViewChild('items') NYGrestrictedElem!: ElementRef<HTMLDivElement>

  powerTriggered: BehaviorSubject<string[]> = new BehaviorSubject<string[]>([])
  changeRoomSettings: FormGroup = new FormBuilder().group({
    SetupTossToken: [false,], // if the owner of the room wish to set the toss system 
    StarterToken: [false,], // affects on the power
    ReverseToken: [false], // if the owner wants to set the a1-a1 or a1-b1 pattern
    BookToken: [''],
    NexusToken: [false],
    TagToken: [false],
    RewindToken: [false],
    DrawToken: [false],
    FreezeToken: [false],
    CovertToken: [false],
    BetToken: [false],
  })
  click(token: { click: boolean, formControl: string }) {
    this.changeRoomSettings.get(token.formControl)?.setValue(token.click, { emitEvent: false })
  }
  getCRFreezeToken = Boolean(this.changeRoomSettings.get('FreezeToken')?.value)
  getCRTagToken = Boolean(this.changeRoomSettings.get('TagToken')?.value)
  getCRRewindToken = Boolean(this.changeRoomSettings.get('RewindToken')?.value)
  getCRDrawToken = Boolean(this.changeRoomSettings.get('DrawToken')?.value)
  getCRNeuxsToken = Boolean(this.changeRoomSettings.get('NexusToken')?.value)
  getCRCovertToken = Boolean(this.changeRoomSettings.get('CovertToken')?.value)
  getCRBetToken = Boolean(this.changeRoomSettings.get('BetToken')?.value)

  getCRSetupTossToken: WritableSignal<boolean>
  getCRReverseToken: WritableSignal<boolean>
  getCRBookToken: string
  getCRStarterToken: WritableSignal<boolean>

  ENavigation: 'CONTROLS' | 'SET' | 'BOOK' | 'POWERS' = 'CONTROLS'
  signalPower: boolean = false
  NYGbookerror: string = ""
  isNYGbookError: boolean = false
  bookToolTip: string = 'style of game'

  powers: OCSFieldFormat[] = [
    { OCStemplate: 'NEXUS', OCSformControl: 'NexusToken', OCSToolTip: "" },
    { OCStemplate: 'TAG', OCSformControl: 'TagToken', OCSToolTip: "" },
    { OCStemplate: 'REWIND', OCSformControl: 'RewindToken', OCSToolTip: "" },
    { OCStemplate: 'DRAW', OCSformControl: 'DrawToken', OCSToolTip: "" },
    { OCStemplate: 'FREEZE', OCSformControl: 'FreezeToken', OCSToolTip: "" },
    { OCStemplate: 'COVERT', OCSformControl: 'CovertToken', OCSToolTip: "" },
    { OCStemplate: 'BET', OCSformControl: 'BetToken', OCSToolTip: "" },
  ]
  collectionKeys: string[] = []

  // dtime: [5s,2m]
  // constant
  readonly DTime: number = 2 // dicussion time minute

  DiscussionTime = `${this.DTime} : ${this.DTimes}s`
  GameTime = `${this.GTime}s`

  Gcount____: WritableSignal<number> = signal<number>(1)
  Dcount____: WritableSignal<number> = signal<number>(1)
  Header = ""
  isNYGpowerError: boolean = false
  NYGgameTimeToolTip = "kickoff time"
  NYGdecisionTimeToolTip = "strategic discussion time"
  NYGpowerError: string
  proccedNavigation: boolean = false

  powerKeys: string[] = []

  constructor(private s: RoomSettingsService, private renderer: Renderer2, private room: RoomService, private FB: FormBuilder) {
  }

  ngAfterViewInit(): void {
    this.NYGclient = this.NYGroomSettings.client
    if (this.NYGclient) {
      if (this.NYGrestrictedElem) {
        this.renderer.setStyle(this.NYGrestrictedElem.nativeElement, 'pointer-events', 'none')
      }
    }
  }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGproceed']) {
      if (changes['NYGproceed'].currentValue == true) {
        this.UpdateTokens()
        this.FillControls()
        this.NYGclient = this.NYGroomSettings.client

      }
    }
  }

  ngOnInit(): void {
    if (this.NYGroomSettings.cap == 2) {
      this.powers = this.powers.filter((e) => e.OCStemplate != 'TAG' && e.OCStemplate != 'DRAW')
    }

    this.FillControls()
    this.UpdateTokens()
    this.NYGclient = this.NYGroomSettings.client

    if (this.Entertainment() == true) {
      this.Header = "Entertainment"
    } else {
      this.Header = "Sports"
    }
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

  }

  UpdateTokens() {

    this.GameTime = this.NYGroomSettings.gameTime.toString()
    this.Gcount____.set(this.NYGroomSettings.gameTime)
    this.Dcount____.set(this.NYGroomSettings.decisionTime)
    this.DTimes = this.Dcount____()
    this.DiscussionTime = `${this.DTime} : ${this.DTimes}s`

    this.setReverse = this.NYGroomSettings.reverse
    this.setStarter = this.NYGroomSettings.starter
    this.setToss = this.NYGroomSettings.setupToss

    this.changeRoomSettings.get("StarterToken")?.setValue(this.setStarter, { emitEvent: false })
    this.changeRoomSettings.get("SetupTossToken")?.setValue(this.setToss, { emitEvent: false })
    this.changeRoomSettings.get("ReverseToken")?.setValue(this.setReverse, { emitEvent: false })

    var pow = Object.fromEntries(
      Object.entries(this.Powers).map(([key, val]) => [
        key.toLowerCase(),
        val
      ])
    ) as Record<string, boolean>

    for (let p of Object.keys(pow)) {
      switch (p) {
        case "nexus":
          this.changeRoomSettings.get('NexusToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "tag":
          this.changeRoomSettings.get('TagToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "rewind":
          this.changeRoomSettings.get('RewindToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "draw":
          this.changeRoomSettings.get('DrawToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "freeze":
          this.changeRoomSettings.get('FreezeToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "covert":
          this.changeRoomSettings.get('CovertToken')?.setValue(pow[p], { emitEvent: false })
          break
        case "bet":
          this.changeRoomSettings.get('BetToken')?.setValue(pow[p], { emitEvent: false })
          break
      }
      if (pow[p]) {
        this.powerKeys.push(p)
      }
    }

    this.changeRoomSettings.get(this.NYGroomSettings.book.toUpperCase())?.setValue(true, { emitEvent: false })
    this.collectionKeys.push(this.NYGroomSettings.book)

    console.log('room book: ', this.NYGroomSettings.book)
  }

  FillControls() {
    this.NYGcollection().forEach(item => {
      this.changeRoomSettings.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
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
      if (old && this.changeRoomSettings.get(newly)?.value == true) {
        this.powerKeys = this.powerKeys.filter(e => e != old)
        this.powerKeys.forEach(e => {
          if (this.changeRoomSettings.get(e)?.value == true) {
            this.changeRoomSettings.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }

    // fetch only active power
    for (let power of this.powerKeys) {
      var active = this.changeRoomSettings.get(power)?.value
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
      if (old && this.changeRoomSettings.get(newly)?.value == true) {
        this.collectionKeys = this.collectionKeys.filter(e => e != old)
        this.collectionKeys.forEach(e => {
          if (this.changeRoomSettings.get(e)?.value == true) {
            this.changeRoomSettings.get(old)?.setValue(false, { emitEvent: false })
          }
        })
      }
    }
  }

  SaveSettings(ev: Event) {

    ev.preventDefault()

    this.getCRSetupTossToken = signal(this.changeRoomSettings.get('SetupTossToken')?.value)
    this.getCRStarterToken = signal(this.changeRoomSettings.get('StarterToken')?.value)
    this.getCRReverseToken = signal(this.changeRoomSettings.get('ReverseToken')?.value)
    this.getCRBookToken = this.collectionKeys[0]
    this.getCRFreezeToken = Boolean(this.changeRoomSettings.get('FreezeToken')?.value)
    this.getCRTagToken = Boolean(this.changeRoomSettings.get('TagToken')?.value)
    this.getCRRewindToken = Boolean(this.changeRoomSettings.get('RewindToken')?.value)
    this.getCRDrawToken = Boolean(this.changeRoomSettings.get('DrawToken')?.value)
    this.getCRNeuxsToken = Boolean(this.changeRoomSettings.get('NexusToken')?.value)
    this.getCRCovertToken = Boolean(this.changeRoomSettings.get('CovertToken')?.value)
    this.getCRBetToken = Boolean(this.changeRoomSettings.get('BetToken')?.value)
    const getGameTime = this.GTime
    const getDecisionTime = this.DTimes

    // if selected than it must save 4 powers
    if (this.powerKeys.length < 4 && this.powerKeys.length !== 0) {
      this.signalPower = true
      var cal = 4 - this.powerKeys.length
      this.isNYGpowerError = true
      this.NYGpowerError = `MUST SELECT ${cal.toString()}  MORE POWERS`
      this.ENavigation = 'POWERS'
    } else {
      this.signalPower = false
      this.isNYGpowerError = false
    }

    if (!this.isNYGpowerError) {
      this.room.ChangeRoomSettings(this.getCRStarterToken(), this.getCRReverseToken(), this.getCRSetupTossToken(), this.getCRBookToken,
        this.getCRDrawToken, this.getCRNeuxsToken, this.getCRTagToken, this.getCRRewindToken,
        this.getCRFreezeToken, this.getCRCovertToken, this.getCRBetToken, getGameTime, getDecisionTime
      )
      this.s.closeSetup()
    }
  }

}
