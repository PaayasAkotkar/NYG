import { Component, effect, Input, OnInit, signal, WritableSignal } from '@angular/core';
import { OcsImgComponent } from "../../profile/ocs-img/ocs-img.component";
import { OcsHeaderV3Component } from "../ocs-header-v3/ocs-header-v3.component";
import { GameMiscelleanous, NamesAsPerTeamsRes } from '../../lobby/resources/lobbyResource';

@Component({
  selector: 'app-entrance',
  imports: [OcsImgComponent, OcsHeaderV3Component],
  templateUrl: './entrance.component.html',
  styleUrls: ['./touchup/entrance.scss', './touchup/color.scss', './touchup/animations.scss']
})
export class EntranceComponent implements OnInit {
  @Input({ required: false }) OCSteamPlayers: GameMiscelleanous

  RedPlayer: WritableSignal<string> = signal('')
  BluePlayer: WritableSignal<string> = signal('')

  @Input({ required: false }) OCSproceed: WritableSignal<boolean> = signal(false)
  @Input() OCSPicURL: WritableSignal<string> = signal('')
  OSCredImg: WritableSignal<string> = signal('')
  OSCblueImg: WritableSignal<string> = signal('')

  ngOnInit(): void {

    this.RedPlayer.set(this.OCSteamPlayers.matchup["RED"].name)
    this.BluePlayer.set(this.OCSteamPlayers.matchup["BLUE"].name)
    this.OSCredImg.set(this.OCSteamPlayers.matchup['RED'].img)
    this.OSCblueImg.set(this.OCSteamPlayers.matchup['BLUE'].img)

  }

  constructor() {
    effect(() => {
      if (this.OCSproceed()) {

        this.RedPlayer.set(this.OCSteamPlayers.matchup["RED"].name)
        this.BluePlayer.set(this.OCSteamPlayers.matchup["BLUE"].name)
        this.OSCredImg.set(this.OCSteamPlayers.matchup['RED'].img)
        this.OSCblueImg.set(this.OCSteamPlayers.matchup['BLUE'].img)
      }
    })
  }
}
