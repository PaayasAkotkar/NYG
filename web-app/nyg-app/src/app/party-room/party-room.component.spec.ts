import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PartyRoomComponent } from './party-room.component';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { Apollo } from 'apollo-angular';
import { ApolloTestingModule } from 'apollo-angular/testing';

describe('PartyRoomComponent', () => {
  let component: PartyRoomComponent;
  let fixture: ComponentFixture<PartyRoomComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PartyRoomComponent, ApolloTestingModule],
      providers: [provideHttpClient, Apollo, provideHttpClientTesting]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PartyRoomComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
