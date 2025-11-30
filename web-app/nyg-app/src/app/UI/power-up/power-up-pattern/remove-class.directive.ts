import { Directive, ElementRef, input, Input, OnChanges, OnInit, Renderer2, SimpleChanges } from '@angular/core';

@Directive({
  selector: '[RemoveClass]'
})
// best suit for animation
export class RemoveClassDirective implements OnInit, OnChanges {
  // to trigger remove 
  @Input({ required: false }) doIt: boolean = false
  // to remove class via doIt command
  @Input({ required: false }) toRemove: string = ""
  // to replace the class via replaceIt command
  @Input() Replace: string = ""
  // to replace the class with the new class via replaceIt command
  @Input() ReplaceWith: string = ""
  // to trigger replace
  @Input() replaceIt: boolean = false
  // to stop the current process
  // note you implicilty have to create the timer func to make the delay false
  // @Input() delay: boolean = false

  constructor(private _ref: ElementRef, private _re: Renderer2) { }

  ngOnInit(): void { }

  _doIt(toRemove: string) {
    this._re.removeClass(this._ref.nativeElement, toRemove)
  }
  _notDoIt() {
    this._re.addClass(this._ref.nativeElement, this.toRemove)
  }
  _replaceWith(replace: string, replaceWith: string) {
    this._re.removeClass(this._ref.nativeElement, replaceWith)
    this._re.addClass(this._ref.nativeElement, replace)
  }
  _replace(replace: string, replaceWith: string) {
    this._re.removeClass(this._ref.nativeElement, replace)
    this._re.addClass(this._ref.nativeElement, replaceWith)
  }
  ngOnChanges(changes: SimpleChanges): void {

    if (changes['doIt']) {
      if (this.doIt) {
        this._doIt(this.toRemove)
      } else {
        this._notDoIt()
      }
    } else if (changes['replaceIt']) {
      if (this.replaceIt) {
        this._replaceWith(this.Replace, this.ReplaceWith)
      } else {
        this._replace(this.Replace, this.ReplaceWith)
      }
    } else {
      // if (this.delay) {
      //   this._doIt(this.Replace)
      //   this._doIt(this.ReplaceWith)
      //   this._doIt(this.toRemove)
      // }
    }
  }
}
