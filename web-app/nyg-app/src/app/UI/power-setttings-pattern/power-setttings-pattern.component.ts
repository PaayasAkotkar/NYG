import { Component, Input, EventEmitter, Output, ViewEncapsulation, ViewChild, ElementRef, OnInit } from '@angular/core';
import { ControlContainer, FormControl, FormGroup, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';


@Component({
  selector: 'app-power-setttings-pattern',
  imports: [ReactiveFormsModule],
  templateUrl: './power-setttings-pattern.component.html',
  styleUrls: ['./touchup/power-setttings-pattern.scss', './touchup/color.scss'],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
  encapsulation: ViewEncapsulation.Emulated,
})
export class PowerSetttingsPatternComponent implements OnInit {
  @Input() NYGheader: string = "NEXUS"
  Link: "ON" | "NO" = "NO"
  track: boolean = false
  @Input() NYGformControl: string = "NEXUS"
  @Output() NYGEVENT = new EventEmitter<string>()
  test() {
    this.NYGEVENT.emit(this.NYGformControl) // send the fromControlName
  }
  ngOnInit(): void {
  }
}
