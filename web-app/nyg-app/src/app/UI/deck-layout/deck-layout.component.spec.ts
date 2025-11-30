import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeckLayoutComponent } from './deck-layout.component';
import { ApolloTestingModule } from 'apollo-angular/testing';
import { Apollo } from 'apollo-angular';

describe('DeckLayoutComponent', () => {
  let component: DeckLayoutComponent;
  let fixture: ComponentFixture<DeckLayoutComponent>;
  const mockApollo = {
    watchQuery: jasmine.createSpy('watchQuery').and.returnValue({
      valueChanges: { subscribe: () => { } }
    }),
    mutate: jasmine.createSpy('mutate').and.returnValue(Promise.resolve({}))
  };

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [, DeckLayoutComponent],
      providers: [{ provide: Apollo, useValue: mockApollo }]
    })
      .compileComponents();

    fixture = TestBed.createComponent(DeckLayoutComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
