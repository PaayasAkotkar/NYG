import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RoomNameComponent } from './room-name.component';

describe('RoomNameComponent', () => {
  let component: RoomNameComponent;
  let fixture: ComponentFixture<RoomNameComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RoomNameComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RoomNameComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
