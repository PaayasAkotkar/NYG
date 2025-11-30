import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LobbyBoxComponent } from './lobby-box.component';

describe('LobbyBoxComponent', () => {
  let component: LobbyBoxComponent;
  let fixture: ComponentFixture<LobbyBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LobbyBoxComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LobbyBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
