import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRoomComponent } from './create-room.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('CreateRoomComponent', () => {
  let component: CreateRoomComponent;
  let fixture: ComponentFixture<CreateRoomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReactiveFormsModule, CreateRoomComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(CreateRoomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
