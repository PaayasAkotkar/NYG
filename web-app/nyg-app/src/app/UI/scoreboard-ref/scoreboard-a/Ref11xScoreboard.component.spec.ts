import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Ref11xScoreboardComponent } from './Ref11xScoreboard.component';

describe('Ref11xScoreboardComponent', () => {
  let component: Ref11xScoreboardComponent;
  let fixture: ComponentFixture<Ref11xScoreboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Ref11xScoreboardComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Ref11xScoreboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
