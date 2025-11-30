import { ElementRef } from '@angular/core';
import { ScrollBehvaiorDirective } from './scroll-behvaior.directive';

describe('ScrollBehvaiorDirective', () => {
  let dummy: HTMLElement
  let _ref: ElementRef<HTMLElement>
  it('should create an instance', () => {
    dummy = document.createElement('div')
    _ref = new ElementRef<HTMLElement>(dummy)
    const directive = new ScrollBehvaiorDirective(_ref);
    expect(directive).toBeTruthy();
  });
});
