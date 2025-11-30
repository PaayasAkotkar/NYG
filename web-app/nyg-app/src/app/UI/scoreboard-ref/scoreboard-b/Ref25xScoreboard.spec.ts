import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Ref25xScoreboardComponent } from './Ref25xScoreboard.component';

describe(' Ref25xScoreboardComponent ', () => {
  let component: Ref25xScoreboardComponent;
  let fixture: ComponentFixture<Ref25xScoreboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Ref25xScoreboardComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Ref25xScoreboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
