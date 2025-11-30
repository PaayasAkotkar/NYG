import { Component, effect, HostListener, Input, OnInit, signal, ViewEncapsulation, WritableSignal } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomService } from '../../roomService/room.service';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { PowerUpPatternComponent } from './power-up-pattern/power-up-pattern.component';
import { PowerService } from './power.service';

@Component({
  selector: 'app-power-up',
  imports: [PowerUpPatternComponent, ReactiveFormsModule],
  templateUrl: './power-up.component.html',
  styleUrls: ['./touchup/power.scss'],
  encapsulation: ViewEncapsulation.Emulated
})

export class PowerUpComponent implements OnInit {
  constructor(private room: RoomService, private ps: PowerService) { }

  ngOnInit(): void { }

  usePowerUpForm = new FormGroup(
    {
      FreezeToken: new FormControl(false),
      NexusToken: new FormControl(false),
      RewindToken: new FormControl(false),
      DrawToken: new FormControl(false),
      TagToken: new FormControl(false),
      CovertToken: new FormControl(false),
      BetToken: new FormControl(false)
    }
  )

  useDrawToken: WritableSignal<boolean | undefined | null | undefined | null> = signal(this.usePowerUpForm.get('DrawToken')?.value)
  useTagToken: WritableSignal<boolean | undefined | null | undefined | null> = signal(this.usePowerUpForm.get('TagToken')?.value)
  useFreezeToken: WritableSignal<boolean | undefined | null | undefined | null> = signal(this.usePowerUpForm.get('FreezeToken')?.value)
  useNexusToken: WritableSignal<boolean | undefined | null> = signal(this.usePowerUpForm.get('NexusToken')?.value)
  useRewindToken: WritableSignal<boolean | undefined | null> = signal(this.usePowerUpForm.get('RevindToken')?.value)
  useCovertToken: WritableSignal<boolean | undefined | null> = signal(this.usePowerUpForm.get('CovertToken')?.value)
  useBetToken: WritableSignal<boolean | undefined | null> = signal(this.usePowerUpForm.get('BetToken')?.value)

  @Input({ required: true }) NYGcanUsePower: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGclash: WritableSignal<boolean>
  @Input() OCSfreezePowerLevel = "1"
  @Input() OCSnexusPowerLevel = "1"
  @Input() OCSdrawPowerLevel = "MAX"
  @Input() OCStagPowerLevel = "MAX"
  @Input() OCSbetPowerLevel = "MAX"
  @Input() OCSrewindPowerLevel = "MAX"
  @Input() OCScovertPowerLevel = "MAX"

  // for room settings that has included these powers in the playing
  @Input({ required: true }) NYGincludeR: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeD: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeT: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeN: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeF: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeC: WritableSignal<boolean | undefined | null | undefined>
  @Input({ required: true }) NYGincludeB: WritableSignal<boolean | undefined | null | undefined>

  // for capturing the players signals for using the power 
  @Input({ required: true }) NYGnexus: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGrewind: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGfreeze: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGdraw: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGtag: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGbet: WritableSignal<boolean | undefined | null>
  @Input({ required: true }) NYGcovert: WritableSignal<boolean | undefined | null>

  @Input({ required: true }) NYGblock: WritableSignal<boolean | undefined | null>

  // the false has been given here make sure to chage to true
  // @HostListener('document:keydown.alt.1', ['$event'])
  // useFreeze(event: KeyboardEvent) {
  //   event.preventDefault()
  //   this.useFreeze = false;
  //   console.log("freeze activated")
  // }

  // @HostListener('document:keydown.alt.2', ['$event'])
  // useNexus(event: KeyboardEvent) {
  //   event.preventDefault()
  //   this.useNexus = false;
  //   console.log("add activated")
  // }

  // @HostListener('document:keydown.alt.3', ['$event'])
  // useRewind(event: KeyboardEvent) {
  //   event.preventDefault()
  //   this.useRewind = false;
  //   console.log("rewind activated")
  // }

  // @HostListener('document:keydown.alt.4', ['$event'])
  // useTag(event: KeyboardEvent) {
  //   event.preventDefault()
  //   this.useTag = false;
  //   console.log("tag activated")
  // }
  // @HostListener('document:keydown.alt.5', ['$event'])
  // useDraw(event: KeyboardEvent) {
  //   event.preventDefault()
  //   this.useDraw = false;
  //   console.log("draw activated")
  // }

  @Input({ required: true }) NYGroomname: string = ""

  closeSetup() {
    this.ps.closeSetup()
  }
  FUse() {
    this.useFreezeToken.set(this.usePowerUpForm.get('FreezeToken')?.value)
    if (this.useFreezeToken() || this.NYGfreeze()) {
      this.room.UseFreezePower(this.NYGroomname, this.NYGclash())
      this.closeSetup()
    }
  }

  NUse() {
    this.useNexusToken.set(this.usePowerUpForm.get('NexusToken')?.value)
    if (this.useNexusToken() || this.NYGnexus()) {
      this.room.UseNexusPower(this.NYGroomname, this.NYGclash())
      this.closeSetup()
    }
  }

  RUse() {
    this.useRewindToken.set(this.usePowerUpForm.get('RewindToken')?.value)
    if (this.useRewindToken() || this.NYGrewind()) {
      this.room.UseRewindPower(this.NYGroomname, this.NYGclash())
      this.closeSetup()
    }
  }

  DUse() {
    this.useDrawToken.set(this.usePowerUpForm.get('DrawToken')?.value)
    if (this.useDrawToken() || this.NYGdraw()) {
      this.room.UseDrawPower(this.NYGroomname)
      this.closeSetup()
    }
  }

  TUse() {
    this.useTagToken.set(this.usePowerUpForm.get('TagToken')?.value)
    if (this.useTagToken() || this.NYGtag()) {
      this.room.UseTagPower(this.NYGroomname)
      this.closeSetup()
    }
  }

  CUse() {
    this.useCovertToken.set(this.usePowerUpForm.get('CovertToken')?.value)
    if (this.useCovertToken() || this.NYGcovert()) {
      this.room.UseCovertPower(this.NYGroomname, this.NYGclash())
      this.closeSetup()
    }
  }

  BUse() {
    this.useBetToken.set(this.usePowerUpForm.get('BetToken')?.value)
    if (this.useBetToken() || this.NYGbet()) {
      this.room.UseBetPower(this.NYGroomname, this.NYGclash())
      this.closeSetup()
    }
  }

}
