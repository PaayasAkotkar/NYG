import { isPlatformBrowser } from '@angular/common';
import { AfterViewInit, Directive, ElementRef, Inject, input, Input, OnChanges, OnInit, PLATFORM_ID, Renderer2, SimpleChanges, ViewChildren } from '@angular/core';

@Directive({
  selector: '[appOnUpgradeColor]'
})
export class OnUpgradeColorDirective implements AfterViewInit, OnChanges {
  // fromColor -> toColor
  // toColor is the base where fromColor covers it
  @Input() fromColor = '#C2BD00'
  @Input() toColor = '#E5E3E1'

  @Input({ required: true }) currentLevel = 3

  constructor(private renderer: Renderer2, private ele: ElementRef,
    @Inject(PLATFORM_ID) private platformId: Object) { }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes['currentLevel']) {
      if (isPlatformBrowser(this.platformId)) {

        // important
        this.renderer.setProperty(this.ele.nativeElement, '--fromPer', `${this.currentLevel}%`)
        this.renderer.setProperty(this.ele.nativeElement, '--toPer', `${this.currentLevel}%`)

      }
    }
  }

  ngAfterViewInit(): void {
  }

}
