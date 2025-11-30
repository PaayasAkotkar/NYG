import { ElementRef, Renderer2 } from '@angular/core';
import { ScrollViewDirective } from './scroll-view.directive';

describe('ScrollViewDirective', () => {
  let renderer: Renderer2

  beforeEach(() => {
    renderer = {
      addClass: jasmine.createSpy('addClass')
    } as unknown as Renderer2
  })
  it('should create an instance', () => {
    let divE = document.createElement('div')
    let dummy = new ElementRef(divE)
    const directive = new ScrollViewDirective(dummy, renderer);
    expect(directive).toBeTruthy();

  });
});
