import { Directive, ElementRef, Input, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[MyTeamColor]'
})
export class OnTeamColorDirective implements OnInit {
  redColor = '#B60120'
  blueColor = '#0068FF'

  constructor(private ele: ElementRef, private renderer: Renderer2) { }
  @Input('MyTeamColor') myTeam: string = ""

  ngOnInit(): void {
    console.log('team color: ', this.myTeam)
    switch (true) {
      case this.myTeam == "RED":
        this.renderer.setStyle(this.ele.nativeElement, 'backgroundColor', this.redColor)
        break
      case this.myTeam == "BLUE":
        this.renderer.setStyle(this.ele.nativeElement, 'backgroundColor', this.blueColor)
        break
    }
  }
}
