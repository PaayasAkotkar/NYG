import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SettingsPatternComponent } from './settings-pattern.component';

describe('SettingsPatternComponent', () => {
  let component: SettingsPatternComponent;
  let fixture: ComponentFixture<SettingsPatternComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SettingsPatternComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SettingsPatternComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
