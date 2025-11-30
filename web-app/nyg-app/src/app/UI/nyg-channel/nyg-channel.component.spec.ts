import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NygChannelComponent } from './nyg-channel.component';

describe('NygChannelComponent', () => {
  let component: NygChannelComponent;
  let fixture: ComponentFixture<NygChannelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NygChannelComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NygChannelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
