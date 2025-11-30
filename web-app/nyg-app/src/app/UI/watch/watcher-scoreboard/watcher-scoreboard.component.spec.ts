import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WatcherScoreboardComponent } from './watcher-scoreboard.component';

describe('WatcherScoreboardComponent', () => {
  let component: WatcherScoreboardComponent;
  let fixture: ComponentFixture<WatcherScoreboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WatcherScoreboardComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WatcherScoreboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
