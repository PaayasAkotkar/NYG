import { Component, Input, OnChanges, OnInit, signal, SimpleChanges, WritableSignal } from '@angular/core';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ColorPalletes as theme, HAND } from '../../helpers/hand';
import { OnModeDirective } from '../../directives/on-mode.directive';
@Component({
  selector: 'app-list',
  imports: [OnModeDirective],
  templateUrl: './list.component.html',
  styleUrls: ['./touchup/list.scss', './touchup/color.scss', './touchup/animations.scss'],

})
export class ListComponent implements OnChanges, OnInit {
  @Input({ required: false }) NYGlist: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) NYGproceed: WritableSignal<boolean> = signal(false)
  @Input() OCSHeader: string = "LIST"
  @Input() NYGMode: string = ''

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['NYGproceed']) {
      if (changes['NYGproceed'].currentValue == true)
        this.Fill()
    }
  }

  fromColor: string = ''
  toColor: string = ''
  bg: string = ''
  singleColor: string = ''

  ngOnInit(): void {
    this.Fill()
    this.update()
  }
  
  update() {
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
  List: WritableSignal<OCSFieldFormat[]> = signal([])
  Fill() {
    this.List = this.NYGlist
  }
}
