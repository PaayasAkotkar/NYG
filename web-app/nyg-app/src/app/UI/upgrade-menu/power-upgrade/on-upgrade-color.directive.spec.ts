import { ElementRef, PLATFORM_ID, Renderer2 } from '@angular/core';
import { OnUpgradeColorDirective } from './on-upgrade-color.directive';
import { TestBed } from '@angular/core/testing';

describe('OnUpgradeColorDirective', () => {
    let dummyRenderer: Renderer2
    beforeEach(() => {
        dummyRenderer = {
            addClass: jasmine.createSpy('addClass')
        } as unknown as Renderer2
        TestBed.configureTestingModule(
            {
                declarations: [OnUpgradeColorDirective],
                providers: [{ provide: PLATFORM_ID, use_value: 'browser' }],
            }
        )
    })
    it('should create an instance', () => {
        let dummyDiv = document.createElement('div')
        let dummy: ElementRef = new ElementRef(dummyDiv)
        const pid = 'browser'
        const directive = new OnUpgradeColorDirective(dummyRenderer, dummy, pid);
        expect(directive).toBeTruthy();
    });
});
