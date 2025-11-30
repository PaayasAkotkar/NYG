import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RefTagPowerComponent } from './ref-tag-power.component';

describe('RefTagPowerComponent', () => {
  let component: RefTagPowerComponent;
  let fixture: ComponentFixture<RefTagPowerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RefTagPowerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RefTagPowerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
