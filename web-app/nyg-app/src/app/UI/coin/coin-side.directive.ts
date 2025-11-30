import { Directive, effect, ElementRef, Input, OnInit, Renderer2, signal, WritableSignal } from '@angular/core';

@Directive({
  selector: '[appCoinSide]'
})
export class CoinSideDirective implements OnInit {

  @Input({ required: true }) isHead: WritableSignal<boolean> = signal(false)
  headFace: string = 'head-face'
  tailFace: string = 'tail-face'

  constructor(private ele: ElementRef, private render: Renderer2) {

    effect(() => {
      if (this.isHead()) {
        this.render.setStyle(this.ele.nativeElement, 'zIndex', '2');
      } else {
        this.render.setStyle(this.ele.nativeElement, 'zIndex', '1');

      }
    })
  }


  ngOnInit(): void {

  }
}
