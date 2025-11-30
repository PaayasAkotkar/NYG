import { ChangeDetectorRef, Component, effect, Input, NgZone, OnDestroy, OnInit, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { PlaygroundComponent, } from '../UI/playground/playground.component';
import { GameMiscelleanous, NamesAsPerTeamsRes, RecBetItems } from '../lobby/resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { FreezePowerService } from '../UI/power-up/power-freeze/power-freeze.service';
import { Subject, Subscription, takeUntil, takeWhile, timer, of, concat, tap, switchMap, concatMap, take } from 'rxjs';
import { NexusPowerService } from '../UI/power-up/power-nexus/power-nexus.service';
import { ViewGuessSheetService } from '../view-guess-sheet/view-guess-sheet.service';
import { BetGuessPowerService } from '../UI/power-up/power-bet/bet-picker/power-bet-picker.service';
import { TimerService } from '../timer/timer.service';
import { PowerRewindService } from '../UI/power-up/power-rewind/power-rewind.service';
import { PowerCovertService } from '../UI/power-up/power-covert/power-covert.service';
import { SpotLightService } from '../UI/spot-light/spot-light.service';
import { EntranceService } from '../UI/entrance/entrance.service';
import { OnLoadService } from './on-load.service';
import { FreezeCountService } from '../UI/power-up/power-freeze/freeze-count/freeze-count.service';

@Component({
  selector: 'app-game',
  imports: [PlaygroundComponent],
  templateUrl: './game.component.html',
  encapsulation: ViewEncapsulation.Emulated
})
export class GameComponent implements OnInit, OnDestroy {
  @Input({ required: true }) NYGClash: WritableSignal<boolean>
  @Input({ required: true }) NYGChances: WritableSignal<number>
  @Input({ required: true }) NYGblueScore: number = 0
  @Input({ required: true }) NYGredScore: number = 0
  @Input({ required: true }) NYGpowerActivated: WritableSignal<boolean> = signal(false)

  @Input({ required: false }) NYGmyroom: string
  @Input({ required: false }) NYGset: number
  @Input({ required: false }) NYGround: number
  @Input({ required: false }) NYGrewindRestart: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGmyTeam: string

  @Input({ required: false }) NYGfriend: WritableSignal<boolean> = signal(false)
  @Input({ required: false }) NYGsheetValue: WritableSignal<string> = signal("NYG")
  @Input({ required: false }) NYGroundOver: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) OCSguessTheme: string = "text11x"
  @Input({ required: true }) OCSboardTheme: string = "text11x"


  // power controls
  @Input({ required: true }) NYGcanUnfreeze: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGnexusValue: WritableSignal<string> = signal("NYG")
  @Input({ required: true }) OCSgameTime: number = 10

  // power used
  @Input({ required: true }) NYGuseNexusPower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGuseFreezePower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGuseCovertPower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGuseRewindPower: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGunFreeze: WritableSignal<boolean> = signal(true)
  @Input({ required: true }) NYGunderTast: WritableSignal<boolean> = signal(false)

  // for betting
  @Input({ required: true }) NYGbetting: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGbetCups: WritableSignal<RecBetItems> = signal({ firstCup: "", secondCup: "", thirdCup: "" })

  @Input({ required: true }) NYGStartIntro: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGonFire: number = 0
  @Input({ required: true }) NYGfinalBoss: WritableSignal<boolean>
  @Input({ required: true }) NYGlastDance: WritableSignal<boolean>
  @Input({ required: true }) OCSplayersAgainst: GameMiscelleanous
  @Input({ required: true }) EntranceProceed: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGfreezeAccess: number = 0

  sessionDone: WritableSignal<boolean> = signal(false)

  StoreTime: number

  // RewindSub: Subscription = new Subscription()
  // UnreezeSub: Subscription = new Subscription()
  // Subscription: Subscription = new Subscription()

  Stop: Subject<void> = new Subject<void>()
  StopSub: Subject<void> = new Subject<void>()

  // to track the game time for rewind restart
  // note to use this as initial time
  TrackGameTime: WritableSignal<number> = signal(0)

  StopTime(stop: boolean) {
    this.sessionDone.set(stop)
    if (this.sessionDone()) {
      this.MarkComplete()
    }
  }

  firstInterval: number = 1900
  secondInterval: number = 1745
  thirdInterval: number = this.firstInterval + this.secondInterval - 25
  defaultGameTime: number = 0
  OCSstartClock: WritableSignal<boolean> = signal(false)
  OCStimeout: boolean = false

  ngOnInit(): void {
    this.show.Setup(this.OCSplayersAgainst, this.NYGnexusValue, this.firstInterval, this.secondInterval)
    this.defaultGameTime = this.OCSgameTime

    const isUsedNexusPower = this.NYGuseNexusPower()
    const isBetting = this.NYGbetting()
    const isUnderTest = this.NYGunderTast()

    timer(this.thirdInterval).subscribe(() => {
      switch (true) {
        case isUsedNexusPower:
          this.nx.Setup(this.NYGnexusValue)
          break

        case isUsedNexusPower && isUnderTest:
          this.cov.Setup()
          this.nx.Setup(this.NYGnexusValue)
          break

        case isBetting && isUsedNexusPower:
          this.bsg.Setup(this.NYGmyroom, this.NYGbetCups, this.NYGClash)
          this.nx.Setup(this.NYGnexusValue)
          break

        case isBetting:
          this.bsg.Setup(this.NYGmyroom, this.NYGbetCups, this.NYGClash)
          break

        case isBetting && isUnderTest:
          this.bsg.Setup(this.NYGmyroom, this.NYGbetCups, this.NYGClash)
          this.cov.Setup()
          break

        case isUnderTest:
          this.cov.Setup()
          break

        case !this.NYGunFreeze():
          this.fs.Setup(this.NYGmyroom, this.NYGfreezeAccess, this.NYGcanUnfreeze,
            this.NYGunFreeze,
            this.NYGunFreeze, this.NYGClash)
          break

        default:
          this.nx.closeSetup()
          this.rew.closeSetup()
          this.fs.closeSetup()
          this.nx.closeSetup()
          this.bsg.closeSetup()
          this.cov.closeSetup()
      }
    })
  }

  constructor(private fcs: FreezeCountService, private show: OnLoadService, private ent: EntranceService, private gt: SpotLightService, private cov: PowerCovertService, private rew: PowerRewindService, private bsg: BetGuessPowerService, private vgs: ViewGuessSheetService, private nx: NexusPowerService, private room: RoomService, private fs: FreezePowerService, private cdr: ChangeDetectorRef, private Zone: NgZone, private OCStimer: TimerService) {
    this.StoreTime = this.OCSgameTime
    effect(() => {
      const isUnfreeze = this.NYGunFreeze()
      const isRewindRestart = this.NYGrewindRestart()
      // this.thirdInterval =
      //   this.firstInterval + this.secondInterval - 25

      if (isRewindRestart) {
        this.rew.Setup()
        this.fs.closeSetup()
        this.OCStimer.RewindTime(this.TrackGameTime(), this.StoreTime, 1000, 3000)
        this.OCStimer.ReStart().pipe(
          takeUntil(this.StopSub),
        ).subscribe({
          next: (token: number) => {
            this.OCSstartClock.set(true)
            this.OCSgameTime = token
            this.cdr.detectChanges()

            if (this.OCSgameTime == 0) {
              this.OCStimer.SignalStop(true)
              this.OCStimeout = true
              this.TimeUp()
            }

          }
        })
      }

      if (this.NYGStartIntro()) {
        timer(this.thirdInterval).pipe(
          take(1),
          concatMap(() => {
            return timer(1).pipe(
              takeUntil(this.Stop),
              tap(() => {
                this.vgs.Setup(this.NYGsheetValue)
                if (isUnfreeze && !isRewindRestart) {
                  this.fs.closeSetup()
                  this.fcs.closeSetup()
                  const speed = 1200
                  this.OCStimer.StartFast(this.OCSgameTime, this.OCSgameTime, speed)
                  this.OCStimer.Start().pipe(
                    takeWhile(c => c >= 0),
                    takeUntil(this.StopSub),

                  ).subscribe({
                    next: (token: number) => {
                      this.OCSgameTime = token
                      this.TrackGameTime.set(token) // save the game time
                      this.cdr.detectChanges()
                    }
                  })
                }
              }))
          })
        ).subscribe()
      }
    })
  }

  TimeUp() {
    this.room.GameTimeUp(this.NYGmyroom, this.NYGClash(), this.NYGChances())
  }


  MarkComplete() {
    this.Stop.next()
    this.Stop.complete()
    this.StopSub.next()
    this.StopSub.complete()
  }

  CloseCdk() {
    this.vgs.closeSetup()
    this.nx.closeSetup()
    this.fs.closeSetup()
    this.bsg.closeSetup()
    this.cov.closeSetup()
  }

  ngOnDestroy(): void {
    this.MarkComplete()
    this.CloseCdk()
  }

}