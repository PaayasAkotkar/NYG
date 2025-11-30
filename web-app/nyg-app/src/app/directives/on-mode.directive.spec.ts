import { TestBed } from '@angular/core/testing';
import { OnModeDirective } from './on-mode.directive';
import { PLATFORM_ID } from '@angular/core';
class _test {
  from: string = ''
  to: string = ''
  isSignal: boolean = false
  single: string = ''
}
describe('OnModeDirective', () => {
  it('should create an instance', () => {
    TestBed.configureTestingModule({
      imports: [OnModeDirective, _test],
      providers: [{ PLATFORM_ID, useValue: 'browser' }]
    }).compileComponents()
  });
});
