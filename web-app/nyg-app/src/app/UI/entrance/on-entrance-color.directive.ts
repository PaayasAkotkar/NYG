import { Directive, ElementRef, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appOnEntranceColor]'
})
export class OnEntranceColorDirective implements OnInit {

  isRedTeam: Boolean = false
  redColor: string = '#FF0012'
  blueColor: string = '#0016A4'
  constructor(private el: ElementRef<HTMLDivElement>, private Renderer: Renderer2) { }
  ngOnInit(): void {
    this.Fill()
  }
  Fill() {
    if (this.isRedTeam) {
      this.Renderer.setStyle(this.el.nativeElement, 'background', this.redColor)
    } else {
      this.Renderer.setStyle(this.el.nativeElement, 'background', this.blueColor)
    }
  }
}
