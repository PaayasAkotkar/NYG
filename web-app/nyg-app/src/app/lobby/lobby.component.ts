import { ChangeDetectorRef, Component, effect, ElementRef, HostListener, Input, NgZone, OnChanges, OnDestroy, OnInit, Query, signal, SimpleChanges, ViewChild, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from './resources/lobbyResource';
import { RoomService } from '../roomService/room.service';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { TossComponent } from "../UI/toss/toss.component";
import { SetChallengeComponent } from "../UI/set-challenge/set-challenge.component";
import { GameComponent } from '../game/game.component';
import { SetupBoardComponent } from "../setup-board/setup-board.component";
import { SetDictionaryComponent } from "../UI/set-dictionary/set-dictionary.component";
import { CoinComponent } from "../UI/coin/coin.component";
import { SetupDrawPowerComponent } from "../setup-draw-power/setup-draw-power.component";
import { SetupAlertMsgComponent } from "../setup-alert-msg/setup-alert-msg.component";
import { MatGridListModule } from '@angular/material/grid-list';
import { PowerBetComponent } from "../UI/power-up/power-bet/power-bet.component";
import { OnEnter } from '../animations/enter-leave';
import { ClashLobbyComponent } from '../UI/clash-lobby/clash-lobby.component';
import { PartyRoomComponent } from '../party-room/party-room.component';

import { SpectateComponent } from "../spectate/spectate/spectate.component";
import { GameOverComponent } from '../game-over/game-over.component';
import { timer } from 'rxjs';
import { PowerTagService } from '../UI/power-up/power-tag/power-tag.service';

@Component({
  animations: [OnEnter],
  selector: 'app-lobby',
  imports: [PartyRoomComponent, ClashLobbyComponent, MatGridListModule, GameComponent, RouterModule, ReactiveFormsModule, TossComponent, SetChallengeComponent, SetupBoardComponent, CoinComponent, SetupDrawPowerComponent, SetupAlertMsgComponent, SetDictionaryComponent, PowerBetComponent, SpectateComponent, GameOverComponent],
  templateUrl: './lobby.component.html',
  encapsulation: ViewEncapsulation.Emulated,
  styles: `
  
  .setup{
    display:flex;
    flex-direction: column-reverse;
    width: 100%;
    height: 100%;
    justify-content: center;
    gap: 10px;
    align-items: center;
  };
  `
})

// todo: when the new connection is joined and a client is unable to see live connection
// after some seconds suggest them to refresh the page 
export class LobbyComponent extends LobbyResource implements OnInit, OnDestroy {
  timeOut: boolean = false

  constructor(private tg: PowerTagService, private __run: NgZone, private cdr: ChangeDetectorRef, private room: RoomService) {
    super(room, __run)
    effect(() => {
      const tag = this.useTagPower()
      if (tag) {
        this.tg.Setup()
      } else {
        this.tg.closeSetup()
      }
    })
  }

  ngOnInit(): void { }

  ngOnDestroy(): void {
  }
  reset(e: boolean) {
    this.timeOut = e
  }
  _timeout(e: boolean) {
    if (!this.block() && this.lock() && !this._Waiting()) {
      this.timeOut = e
    } else {
      this.timeOut = false
    }

    console.log('timeout!!!⏲️ beach')
  }
}
