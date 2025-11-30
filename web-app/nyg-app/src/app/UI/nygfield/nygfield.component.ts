import { EventEmitter, Component, Input, Output, output, ViewEncapsulation, OnInit, OnChanges, SimpleChanges, ViewChild, ElementRef, WritableSignal, signal, Renderer2, effect } from '@angular/core';
import { ControlContainer, FormArray, FormBuilder, FormControl, FormGroup, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { OCSFieldFormat } from '../../helpers/patterns';
import { ColorPalletes as theme } from '../../helpers/hand';
import { RoomService } from '../../roomService/room.service';
import { OnModeDirective } from "../../directives/on-mode.directive";
import { ScrollViewDirective } from "../../directives/scroll-view.directive";

/***
 * difference between the V1 and V2 is of the content placement
 */
@Component({
  selector: 'ref-field',
  imports: [ReactiveFormsModule, OnModeDirective, ScrollViewDirective],
  templateUrl: './nygfield.component.html',
  styleUrls: ['./touchup/book.scss', './touchup/color.scss', './touchup/animations.scss'],
  encapsulation: ViewEncapsulation.Emulated,
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],

})
export class NYGfieldV2Component implements OnInit {
  @Input({ required: true }) NYGcollection: WritableSignal<OCSFieldFormat[]> = signal([])
  @Input({ required: true }) NYGheader: string = ""
  @Input({ required: true }) NYGsetTopPosition: WritableSignal<number> = signal(0)

  @Input({ required: true }) NYGMode: WritableSignal<string> = signal("")

  @Output() NYGEVENT = new EventEmitter<string>()
  @Output() NYGposition = new EventEmitter<number>();

  fromColor: string = ''
  toColor: string = ''

  bg: string = ''
  singleColor: string = ''

  ngOnInit(): void {
    switch (this.NYGMode()) {

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

      default:
        this.bg = theme.$nyg_white_palletes.ad8_white
        this.toColor = theme.$nyg_red_palletes.ad11_red
        this.fromColor = theme.$nyg_blue_palletes.ad11_blue
        this.singleColor = theme.$nyg_white_palletes.ad8_white
        break
    }
  }

  constructor(private room: RoomService, private renderer: Renderer2) { }
}
