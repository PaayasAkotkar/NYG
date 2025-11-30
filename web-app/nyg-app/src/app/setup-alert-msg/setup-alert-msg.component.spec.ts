import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SetupAlertMsgComponent } from './setup-alert-msg.component';

describe('SetupAlertMsgComponent', () => {
  let component: SetupAlertMsgComponent;
  let fixture: ComponentFixture<SetupAlertMsgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SetupAlertMsgComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SetupAlertMsgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
