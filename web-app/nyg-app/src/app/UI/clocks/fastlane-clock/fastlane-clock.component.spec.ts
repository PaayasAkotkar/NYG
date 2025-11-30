import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FastlaneClockComponent } from './fastlane-clock.component';

describe('FastlaneClockComponent', () => {
  let component: FastlaneClockComponent;
  let fixture: ComponentFixture<FastlaneClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FastlaneClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FastlaneClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
