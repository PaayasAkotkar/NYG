import { Component, EventEmitter, Output, Input, OnChanges, OnInit, signal, SimpleChanges, ViewEncapsulation, WritableSignal } from '@angular/core';
import { LobbyResource } from '../../lobby/resources/lobbyResource';
import { RoomService } from '../../roomService/room.service';
import { FlipCoin, X } from './animation/animation';
import { CoinSideDirective } from './coin-side.directive';

@Component({
  selector: 'app-coin',
  imports: [CoinSideDirective],
  templateUrl: './coin.component.html',
  styleUrls: ['./touchup/coin.scss', './touchup/animation.scss'],
  animations: [FlipCoin, X],
  encapsulation: ViewEncapsulation.Emulated
})
export class CoinComponent implements OnInit, OnChanges {
  constructor(private room: RoomService) { }

  @Input({ required: true }) NYGtimeout: boolean
  @Input({ required: true }) NYGfriend: WritableSignal<boolean>
  @Input({ required: true }) NYGmyRoom: string
  @Input({ required: true }) NYGisHead: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGclash: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGchance: WritableSignal<number> = signal(0)
  @Input({ required: true }) NYGfinalBoss: boolean = false
  @Input({ required: true }) NYGlastDance: boolean = false
  @Input({ required: true }) NYGround: number = 0
  @Input({ required: true }) NYGset: number = 0
  @Output() NYGstop: EventEmitter<boolean> = new EventEmitter<boolean>() // this basically sets to default value

  isHead: WritableSignal<boolean> = signal(false)
  isTail: WritableSignal<boolean> = signal(false)

  coinState: string = ""
  isOpen: boolean = false
  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGtimeout']) {
      if (changes['NYGtimeout'].currentValue == true) {
        this.Bot()
        this.NYGstop.emit(false)
      }
    }
  }

  Bot() {
    this.room.TossTimeOut(this.NYGmyRoom, this.NYGclash(), "", this.NYGround, this.NYGset)
  }
  ngOnInit(): void {
    const randomInt = Math.floor(Math.random() * 11)
    if (randomInt % 2 == 0) {

      this.isHead.set(true)
      this.isTail.set(false)
    } else {
      this.isTail.set(true)
      this.isHead.set(false)
    }
  }

  Result() {
    if (this.NYGisHead()) {
      this.isHead.set(true)
      this.isTail.set(false)
    } else {
      this.isTail.set(true)
      this.isHead.set(false)
    }
  }

  Toss() {
    this.isOpen = !this.isOpen
    this.isHead != this.isHead
    // if (!this.NYGfriend && !this.NYGclash()) {
    //   this.room.TossSession("", false, this.NYGmyRoom, false, 0, false, false)
    // } else if (this.NYGclash()) {
    //   this.room.TossSession("", true, this.NYGmyRoom, true, this.NYGchance(), this.NYGfinalBoss, this.NYGlastDance)
    // } else {
    //   this.room.TossSession("", true, this.NYGmyRoom, false, 0, false, false)
    // }

    this.room.TossSession("", this.NYGfriend(), this.NYGmyRoom, this.NYGclash(), this.NYGchance(), this.NYGfinalBoss, this.NYGlastDance)

  }

}
