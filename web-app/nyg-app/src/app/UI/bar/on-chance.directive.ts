import { Directive, ElementRef, Input, OnInit, Renderer2 } from '@angular/core';

@Directive({
  selector: '[appOnChance]'
})
export class OnChanceDirective implements OnInit {
  to: string = '#45556B'
  from: string = '#f5f4de'
  Chance = 10
  @Input() NYGTeamName: string = ''
  @Input({ required: true }) OCScurrentChance = 0
  constructor(private _ref: ElementRef, private renderer: Renderer2) { }
  ngOnInit(): void {
    switch (this.NYGTeamName) {
      case 'RED':
        this.renderer.setStyle(this._ref.nativeElement, '--bg', 'blue')
        break
      case 'BLUE':
        this.renderer.setStyle(this._ref.nativeElement, '--bg', 'black')

    }
    this.renderer.setStyle(this._ref.nativeElement, '--bg', 'black')

    const children = this._ref.nativeElement.children

    for (let i = 0; i < children.length && i < this.OCScurrentChance; i++) {
      const child = children[i];

      // Set CSS variables
      this.renderer.setStyle(child, '--from', this.from);
      this.renderer.setStyle(child, '--to', this.to);
    }
  }
}
