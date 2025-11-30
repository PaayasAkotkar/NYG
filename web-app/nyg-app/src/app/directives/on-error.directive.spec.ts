import { ElementRef, PLATFORM_ID, Renderer2 } from '@angular/core';
import { OnErrorDirective } from './on-error.directive';
import { TestBed } from '@angular/core/testing';

describe('OnErrorDirective', () => {
  let dummyRenderer: Renderer2

  beforeEach(() => {

    dummyRenderer = {
      setProperty: jasmine.createSpy('setProperty'),
      addClass: jasmine.createSpy('addClass'),
      removeClass: jasmine.createSpy('removeClass')
    } as unknown as Renderer2
    TestBed.configureTestingModule(
      {
        imports: [OnErrorDirective],
        providers: [{ provide: PLATFORM_ID, useValue: 'bowser' }],
      }
    )
  })

  it('should create an instance', () => {
    let dummyDiv = document.createElement('div')
    let dummy: ElementRef = new ElementRef(dummyDiv)
    const pid = 'browser'
    const directive = new OnErrorDirective(pid, dummy, dummyRenderer);
    expect(directive).toBeTruthy();
  });
});
