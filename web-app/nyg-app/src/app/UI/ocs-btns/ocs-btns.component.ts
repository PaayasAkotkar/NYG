import { Component, Input, Output, EventEmitter } from '@angular/core';
import { OnClickColorDirective } from '../upgrade-menu/power-upgrade/on-click-color.directive';
import { ColorPalletes as theme } from '../../helpers/hand';
import { OnModeDirective } from '../../directives/on-mode.directive';

@Component({
  selector: 'app-ocs-btns',
  imports: [OnModeDirective,],
  templateUrl: './ocs-btns.component.html',
  styleUrls: ['./touchup/ocs-btn.scss', './touchup/color.scss',
    './touchup/animations.scss'
  ],
})
export class OcsBtnsComponent {
  @Input({ required: true }) OCSheader: string = "CLICK"
  @Input({ required: false }) OCSteam: string = "BLUE"
  @Output() OCSclick = new EventEmitter<any>()
  @Input({ required: true }) OCSType: 'submit' | 'button' = 'submit'
  Click() {
    this.OCSclick.emit(this.OCSheader)
  }
  fromColor: string = ''
  toColor: string = ''
  bg: string = ''
  singleColor: string = ''
  @Input({ required: false }) NYGMode: string = ''

  ngOnInit(): void {
    switch (this.NYGMode) {

      case 'mysterium':
        this.bg = theme.$nyg_brown_palletes.ad5_brown
        this.fromColor = theme.$nyg_green_palletes.ad18_green
        this.toColor = theme.$nyg_red_palletes.ad16_red
        this.singleColor = theme.$nyg_black_palletes.ad8_black
        break

      case 'node':
        this.bg = theme.$nyg_white_palletes.ad9_white
        this.fromColor = theme.$nyg_yellow_palletes.ad9_yellow
        this.toColor = theme.$nyg_white_palletes.ad8_white
        this.singleColor = theme.$nyg_white_palletes.ad8_white
        break

      case 'pshycic':
        this.bg = theme.$nyg_voilet_palletes.ad8_voilet
        this.fromColor = theme.$nyg_white_palletes.ad8_white
        this.toColor = theme.$nyg_yellow_palletes.ad4_yellow
        this.singleColor = theme.$nyg_black_palletes.umber_black
        break

      case 'fastlane':
        this.bg = theme.$nyg_blue_palletes.ad17_blue
        this.fromColor = theme.$nyg_orange_palletes.ad8_orange
        this.toColor = theme.$nyg_yellow_palletes.ad8_yellow
        this.singleColor = theme.$nyg_black_palletes.umber_black
        break
      case 'gym':
        this.bg = theme.$nyg_black_palletes.ad7_black
        this.fromColor = theme.$nyg_brown_palletes.ad4_brown
        this.toColor = theme.$nyg_yellow_palletes.ad7_yellow
        this.singleColor = theme.$nyg_white_palletes.ad8_white
        break
      case 'locify':
        this.bg = theme.$nyg_green_palletes.ad20_green
        this.toColor = theme.$nyg_red_palletes.ad17_red
        this.fromColor = theme.$nyg_yellow_palletes.ad10_yellow
        this.singleColor = theme.$nyg_white_palletes.ad8_white
        break

      default:
        this.bg = theme.$nyg_white_palletes.ad8_white
        this.singleColor = theme.$nyg_white_palletes.ad8_white
        break
    }
  }
}
