import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Scoreboard11xComponent } from './11xScoreboard.component';

describe('Scoreboard11xComponent', () => {
  let component: Scoreboard11xComponent;
  let fixture: ComponentFixture<Scoreboard11xComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Scoreboard11xComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(Scoreboard11xComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
