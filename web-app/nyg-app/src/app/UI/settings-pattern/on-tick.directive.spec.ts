import { ElementRef, Renderer2 } from '@angular/core';
import { OnTickDirective } from './on-tick.directive';
import { TestBed } from '@angular/core/testing';

describe('OnTickDirective', () => {
    let dummyRenderer: Renderer2

    beforeEach(() => {

        dummyRenderer = {
            setProperty: jasmine.createSpy('setProperty'),
            addClass: jasmine.createSpy('addClass'),
            removeClass: jasmine.createSpy('removeClass')
        } as unknown as Renderer2
        TestBed.configureTestingModule(
            {
                declarations: [OnTickDirective],
            }
        )
    })

    it('should create an instance', () => {
        let dummyDiv = document.createElement('div')
        let dummy: ElementRef = new ElementRef(dummyDiv)
        const pid = 'browser'
        const directive = new OnTickDirective(dummyRenderer, dummy);
        expect(directive).toBeTruthy();
    });
});
