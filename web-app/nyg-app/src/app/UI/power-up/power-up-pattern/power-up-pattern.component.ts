import { Component, Input, Output, EventEmitter, ViewEncapsulation, OnInit, Renderer2, ViewChild, ElementRef, QueryList, ViewChildren, Inject, PLATFORM_ID, AfterViewInit } from '@angular/core';
import { ControlContainer, FormGroupDirective, ReactiveFormsModule } from '@angular/forms';
import { RemoveClassDirective } from './remove-class.directive';

@Component({
  selector: 'app-power-up-pattern',
  imports: [ReactiveFormsModule, RemoveClassDirective],
  templateUrl: './power-up-pattern.component.html',
  styleUrls: ['./touchup/animations.scss', './touchup/power-up.scss', './touchup/color.scss'],
  viewProviders: [{ provide: ControlContainer, useExisting: FormGroupDirective }],
  encapsulation: ViewEncapsulation.Emulated
})
export class PowerUpPatternComponent implements OnInit, AfterViewInit {

  @Input() OCSformControl: string = ""
  @Input() OCSlvl: string = "1"
  @Input() OCSpower: string = ""
  @Output() OCSEvent = new EventEmitter<void>()
  @Input() OCSdisableHover: boolean = false
  @ViewChild('_animate_',) _C!: ElementRef

  constructor(@Inject(PLATFORM_ID) private id: Object, private renderer: Renderer2) { }

  ngOnInit(): void { }

  ngAfterViewInit(): void {
    if (this.OCSdisableHover) {
      this.renderer.removeClass(this._C.nativeElement, '_animate')
    }
  }

  Event() {
    this.OCSEvent.emit()
  }
}
