import { AfterViewInit, Directive, ElementRef, Inject, Input, OnChanges, OnInit, PLATFORM_ID, Renderer2, RendererStyleFlags2, SimpleChanges } from '@angular/core';
import { OnChanceDirective } from '../UI/bar/on-chance.directive';
import { debounceTime, Subject } from 'rxjs';
import { isPlatformBrowser } from '@angular/common';

@Directive({
  selector: '[appOnMode]'
})
export class OnModeDirective implements OnInit {

  @Input({ required: true }) toColor: string = ''
  @Input({ required: true }) fromColor: string = ''
  @Input({ required: true }) bg: string = ''
  @Input({ required: false }) isSingleColor: boolean = false
  @Input() singleColor: string = ''
  colorChangeSubject = new Subject<void>()
  constructor(@Inject(PLATFORM_ID) private pd: Object, private _ref: ElementRef, private _rend: Renderer2) {
  }
  ngOnInit(): void {
    this.Color();
  }

  Color() {
    this._rend.setStyle(
      this._ref.nativeElement,
      '--from',
      this.fromColor, RendererStyleFlags2.DashCase
    );
    this._rend.setStyle(
      this._ref.nativeElement,
      '--to',
      this.toColor, RendererStyleFlags2.DashCase
    );

    this._rend.setStyle(this._ref.nativeElement, '--bg', this.bg, RendererStyleFlags2.DashCase)
    if (this.isSingleColor)
      this._rend.setStyle(this._ref.nativeElement, '--col_', this.singleColor, RendererStyleFlags2.DashCase)

  }
}


