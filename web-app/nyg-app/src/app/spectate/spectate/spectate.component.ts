import { Component, EventEmitter, effect, Input, Output, signal, WritableSignal, OnChanges, SimpleChanges, input, OnInit, NgZone, OnDestroy } from '@angular/core';
import { OCSFieldFormat } from '../../helpers/patterns';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { WatchComponent } from "../../UI/watch/watch.component";
import { NYGfieldV2Component } from "../../UI/nygfield/nygfield.component";
import { Tags2Component } from "../../UI/tags2/tags2.component";
import { RoomService } from '../../roomService/room.service';
import { ISpectate } from '../../roomService/resources/resource';
import { HAND } from '../../helpers/hand';
import { single } from 'rxjs';
import { GameMiscelleanous, SpectateInfo } from '../../lobby/resources/lobbyResource';

@Component({
  selector: 'app-spectate',
  imports: [ReactiveFormsModule, WatchComponent, NYGfieldV2Component, Tags2Component],
  templateUrl: './spectate.component.html',
  styleUrl: './spectate.component.scss'
})
export class SpectateComponent implements OnInit, OnDestroy {
  NYGteamRedText: WritableSignal<string> = signal("")
  NYGteamBlueText: WritableSignal<string> = signal("")

  NYGteamRedName: WritableSignal<string> = signal("technical issue")
  NYGteamBlueName: WritableSignal<string> = signal("technical issue")

  @Input({ required: true }) NYGredTeamScore: string | number = 0
  @Input({ required: true }) NYGblueTeamScore: string | number = 0

  NYGsetTopPosition: WritableSignal<number> = signal(0)
  NYGonClick: WritableSignal<string> = signal("")

  @Input({ required: true }) NYGgame: WritableSignal<boolean>
  @Input({ required: true }) NYGheader: string
  @Input() NYGmode: WritableSignal<string> = signal('')
  @Input({ required: true }) NYGproceed: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) OCScollections: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) NYGspecInfo: SpectateInfo
  @Input({}) OCSteamPlayers: GameMiscelleanous = { matchup: {} }
  ngOnDestroy(): void {

  }

  FillControls() {
    this.OCScollections().forEach(item => {
      this._ref.addControl<any>(item.OCStemplate, this.FB.control(item.OCSformControl))
    })

  }

  ngOnInit(): void {

    this.NYGteamRedName.set(this.OCSteamPlayers.matchup["RED"].name)
    this.NYGteamBlueName.set(this.OCSteamPlayers.matchup["BLUE"].name)

    this._ref.valueChanges.subscribe({
      next: (selectedKey) => {
        if (selectedKey) {
          Object.keys(this._ref.controls).forEach(otherKey => {
            // only uncheck the previous selected value of book
            if (otherKey !== selectedKey) {
              this._ref.get(otherKey)?.setValue(false, { emitEvent: false });
            }

          });
        }
        // this.Update()
      }
    })

    var Spectate = this._room.Rooms()
    Spectate.subscribe({
      next: (token) => {
        var NVPattern = /NYGView:(?=[\s\S\w])/
        var NSPattern = /NYGSpectate:(?=[\s\S\w])/
        var NBSPattern = /NYGblueSpectate:(?=[\s\S\w])/
        var NRSPattern = /NYGredSpectate:(?=[\s\S\w])/

        var blueView = NBSPattern.test(token)
        var redView = NRSPattern.test(token)
        var spectate = NSPattern.test(token)
        var view = NVPattern.test(token)
        const count = this.Search(token, ":")

        switch (true) {
          case blueView:
            var bt = this.afterColon(token, count, false, false)
            this.NYGteamBlueText.set(bt)
            break
          case redView:
            var bt = this.afterColon(token, count, false, false)
            this.NYGteamRedText.set(bt)
            break

          case spectate:

            break;
          case view:
            var p = 'NYGView:'
            var _view: ISpectate = {
              word: "",
              scrollPos: -1, onClick: "",
              scrollHappend: false,
            }

            var _p: ISpectate = this.afterColonJson(token, p, _view)
            if (_p.scrollHappend) {
              this.NYGsetTopPosition.set(_p.scrollPos)
              this.NYGonClick.set(_p.onClick)
              this._ref.get(this.NYGonClick())?.setValue(true, { emitEvent: false })
              this.Update() // just for safety in case the on init doesnt triggered

            } else {
              this.NYGonClick.set(_p.onClick)
              this._ref.get(this.NYGonClick())?.setValue(true, { emitEvent: false })
              this.Update()
            }
            break
        }
      }
    })
  }

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

  Update() {
    const formValue: any = this._ref.value;
    Object.keys(formValue).filter(key => {
      this._ref.get(key)?.valueChanges.subscribe({
        next: (token) => {
          if (token) {

            Object.keys(this._ref.controls).forEach(otherKey => {
              // only uncheck the previous selected value of book
              if (otherKey !== key) {
                this._ref.get(otherKey)?.setValue(false, { emitEvent: false });
              }

            });
          }
        }
      })
    })

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

  afterColonJson(token: string | undefined, parse: string, _interface: any) {
    if (token) {
      token = token.substring(parse.length)
      _interface = JSON.parse(token)
    }
    return _interface
  }

  constructor(private run: NgZone, private FB: FormBuilder, private _room: RoomService) {
    effect(() => {
      if (this.NYGproceed()) {
        this.FillControls()
      }
      if (this.NYGonClick() != "") {
        this._ref.get(this.NYGonClick())?.setValue(true, { emitEvent: false })
      }
    })

  }

  _ref: FormGroup = new FormGroup({})
}
