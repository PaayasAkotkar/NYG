import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MakeRoomComponent } from './make-room.component';
import { ApolloTestingController, ApolloTestingModule } from 'apollo-angular/testing';
describe('MakeRoomComponent', () => {
  let component: MakeRoomComponent;
  let fixture: ComponentFixture<MakeRoomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [MakeRoomComponent, ApolloTestingModule],

    })
      .compileComponents();

    fixture = TestBed.createComponent(MakeRoomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
