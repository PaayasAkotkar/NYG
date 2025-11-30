import { Directive, ElementRef, Input, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appOnTeamColor]'
})
export class OnTeamColorDirective implements OnInit {
  @Input({ required: true }) Team: string = ""
  blueColor: string = '#3725A6'
  redColor: string = '#B10000'
  blackColor: string = '#FFFF00'
  greyColor: string = '#AFAFAF'
  @Input() doBackground: boolean = false
  constructor(private el: ElementRef, private renderer: Renderer2) { }
  ngOnInit(): void {
    this.Color()
  }
  Color() {
    if (!this.doBackground) {
      if (this.Team == "RED") {
        this.renderer.setStyle(this.el.nativeElement, 'color', this.redColor)
      } else if (this.Team == "BLUE") {
        this.renderer.setStyle(this.el.nativeElement, 'color', this.blueColor)
      }
    } else {
      if (this.Team == "RED") {
        this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', this.redColor)
      } else if (this.Team == "BLUE") {
        this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', this.blueColor)
      } else if (this.Team == "GREY") {
        this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', this.greyColor)
      } else {
        this.renderer.setStyle(this.el.nativeElement, 'backgroundColor', this.blackColor)
      }

    }
  }
}
