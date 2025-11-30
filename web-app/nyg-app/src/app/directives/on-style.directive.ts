import { AfterViewInit, ChangeDetectorRef, Directive, ElementRef, Input, input, NgZone, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appOnStyle]'
})
export class OnStyleDirective implements AfterViewInit, OnInit {

  @Input() bgColor: string = ''
  @Input() isLinearBg: boolean = false
  @Input() toColor: string = ''
  @Input() fromColor: string = ''
  constructor(private zone: NgZone, private cdr: ChangeDetectorRef, private el: ElementRef, private renderer: Renderer2) { }
  ngAfterViewInit(): void {

    if (!this.isLinearBg) {
      this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', this.bgColor
      )
      this.cdr.detectChanges()

    } else {
      this.zone.run(() => {
        const q = `linear-gradient(to right,${this.fromColor},${this.toColor})`
        this.renderer.setStyle(this.el.nativeElement, 'backgroundImage', q)
        this.cdr.detectChanges()
      })

    }
  }
  ngOnInit(): void {

  }
}
