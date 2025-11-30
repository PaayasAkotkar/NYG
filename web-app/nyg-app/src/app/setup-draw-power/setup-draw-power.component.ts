import { Component, Input, OnDestroy, OnInit, signal, WritableSignal } from '@angular/core';
import { DrawPowerService } from '../UI/power-up/power-draw/power-draw.service';

@Component({
  selector: 'app-setup-draw-power',
  imports: [],
  templateUrl: './setup-draw-power.component.html',
})
export class SetupDrawPowerComponent implements OnInit, OnDestroy {
  @Input({ required: true }) NYGdrawOffer: WritableSignal<boolean> = signal(false)
  @Input({ required: true }) NYGdmyRoom: string = ""
  @Input({ required: true }) NYGmyTeam: string = ""

  ngOnInit(): void {
    if (this.NYGdrawOffer()) {
      this.ds.Setup(this.NYGdmyRoom, this.NYGmyTeam)
    }
  }
  ngOnDestroy(): void {
    this.ds.closeSetup()
  }
  constructor(private ds: DrawPowerService) {

  }
}
