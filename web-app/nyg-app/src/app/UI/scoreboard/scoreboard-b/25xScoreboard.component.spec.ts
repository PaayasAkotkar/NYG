import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Scoreboard25xComponent } from './25xScoreboard.component';

describe('Scoreboard25xComponent', () => {
  let component: Scoreboard25xComponent;
  let fixture: ComponentFixture<Scoreboard25xComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Scoreboard25xComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Scoreboard25xComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
