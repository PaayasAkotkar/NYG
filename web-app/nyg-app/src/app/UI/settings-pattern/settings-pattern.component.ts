import { Component, Input, EventEmitter, OnChanges, Output, SimpleChanges, OnInit, Renderer2, HostListener, ViewEncapsulation } from '@angular/core';
import { ControlContainer, FormBuilder, FormControl, FormGroup, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { OnTickDirective } from './on-tick.directive';

@Component({
  selector: 'app-settings-pattern',
  imports: [ReactiveFormsModule,],
  templateUrl: './settings-pattern.component.html',
  styleUrls: ['./touchup/settings-pattern.scss', './touchup/color.scss', './touchup/animations.scss'],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
  encapsulation: ViewEncapsulation.Emulated
})
export class SettingsPatternComponent implements OnInit, OnChanges {
  @Input() NYGheader: string = "NYG"
  Link: "ON" | "NO" = "NO"
  track: boolean = false
  @Input() NYGsetToolTip: boolean = false
  @Input() NYGtoolTip: string = ""
  @Input() NYGformControl!: string
  @Output() NYGEVENT = new EventEmitter<any>()
  @Output() NYGclick = new EventEmitter<{ click: boolean, formControl: string }>()
  ngOnChanges(changes: SimpleChanges): void { }

  ngOnInit(): void { }

  click() {
    this.track = !this.track
    this.NYGclick.emit({ click: this.track, formControl: this.NYGformControl })
    if (this.track) {
      this.Link = "ON"
    } else {
      this.Link = "NO"
    }
    this.NYGEVENT.emit();

  }

  test() {
    this.track = !this.track
    if (this.track) {
      this.Link = "ON"
    } else {
      this.Link = "NO"
    }
    this.NYGEVENT.emit();
  }

}
