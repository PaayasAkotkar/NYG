import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OcsIndoxMsgComponent } from './ocs-indox-msg.component';

describe('OcsIndoxMsgComponent', () => {
  let component: OcsIndoxMsgComponent;
  let fixture: ComponentFixture<OcsIndoxMsgComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OcsIndoxMsgComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OcsIndoxMsgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
