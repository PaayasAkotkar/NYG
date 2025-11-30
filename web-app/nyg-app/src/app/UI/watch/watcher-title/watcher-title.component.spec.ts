import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WatcherTitleComponent } from './watcher-title.component';

describe('WatcherTitleComponent', () => {
  let component: WatcherTitleComponent;
  let fixture: ComponentFixture<WatcherTitleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WatcherTitleComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WatcherTitleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
