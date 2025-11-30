import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MysteriumClockComponent } from './mysterium-clock.component';

describe('MysteriumClockComponent', () => {
  let component: MysteriumClockComponent;
  let fixture: ComponentFixture<MysteriumClockComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MysteriumClockComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MysteriumClockComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
