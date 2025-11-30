import { AfterViewInit, Component, effect, ElementRef, EventEmitter, Input, OnChanges, OnInit, Output, Renderer2, signal, SimpleChanges, ViewChild, ViewEncapsulation, WritableSignal } from '@angular/core';
import { ControlContainer, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ColorPalletes as theme } from '../../helpers/hand';
import { OnModeDirective } from '../../directives/on-mode.directive';

import { RoomService } from '../../roomService/room.service';
import { ISpectate } from '../../roomService/resources/resource';
export interface ItypeScroll {
  scrollPos: number,
  isScroll: boolean
}
@Component({
  selector: 'app-field',
  imports: [OnModeDirective, ReactiveFormsModule],
  templateUrl: './field.component.html',
  styleUrls: ['./touchup/book.scss', './touchup/color.scss',
    './touchup/animations.scss'
  ],
  encapsulation: ViewEncapsulation.Emulated,
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
})
export class FieldComponent implements OnInit {
  @Input({ required: false }) NYGcollection: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: false }) NYGheader: string = "ENTERTAINMENT"
  @Output() NYGEVENT = new EventEmitter<string>()
  @Output() NYGposition = new EventEmitter<number>();
  @Input({ required: false }) NYGMode: string = 'locify'

  fromColor: string = ''
  toColor: string = ''
  bg: string = ''
  singleColor: string = ''

  ScrollPos: number = 0

  onScroll(token: number) {
    this.NYGposition.emit(token)
  }

  ngOnInit(): void {
    switch (this.NYGMode) {

      case 'mysterium':
        this.bg = theme.$nyg_brown_palletes.ad5_brown
        this.fromColor = theme.$nyg_green_palletes.ad18_green
        this.toColor = theme.$nyg_red_palletes.ad16_red
        this.singleColor = theme.$nyg_white_palletes.ad8_white
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
      case 'CREATE':
        this.bg = theme.$nyg_white_palletes.ad8_white
        this.singleColor = theme.$nyg_black_palletes.ad7_black
        this.toColor = theme.$nyg_red_palletes.ad11_red
        this.fromColor = theme.$nyg_blue_palletes.ad11_blue
        break
      default:
        this.bg = theme.$nyg_white_palletes.ad8_white
        this.toColor = theme.$nyg_red_palletes.ad11_red
        this.fromColor = theme.$nyg_blue_palletes.ad11_blue
        this.singleColor = theme.$nyg_white_palletes.ad8_white

        break
    }
  }

  OnClickValue(value: string) {
    this.NYGEVENT.emit(value)
  }

  constructor(private room: RoomService, private render: Renderer2) {

  }
}