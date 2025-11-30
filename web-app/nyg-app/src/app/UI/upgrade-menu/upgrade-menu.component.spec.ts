import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpgradeMenuComponent } from './upgrade-menu.component';

describe('UpgradeMenuComponent', () => {
  let component: UpgradeMenuComponent;
  let fixture: ComponentFixture<UpgradeMenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UpgradeMenuComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(UpgradeMenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should distribute the powers in min where front row contains powers to upgrade and back no upgrades', () => {
    component.OCSlist = [{
      power: "nexus",
      spur: 10,
      counter: 0,
      level: 0,
      allow: true,
      display: "10S"
    },
    {
      power: "freeze",
      spur: 10,
      counter: 0,
      level: 0,
      allow: true,
      display: "10S"
    },
    {
      power: "draw",
      spur: 10,
      counter: 0,
      level: 100,
      allow: false,
      display: "S"
    },
    {
      power: "tag",
      spur: 10,
      counter: 0,
      level: 0,
      allow: false,
      display: "S"
    },
    {
      power: "bet",
      spur: 10,
      counter: 0,
      level: 0,
      allow: false,
      display: "S"
    },
    {
      power: "rewind",
      spur: 10,
      counter: 0,
      level: 0,
      allow: false,
      display: "S"
    },
    {
      power: "covert",
      spur: 10,
      counter: 0,
      level: 0,
      allow: false,
      display: "S"
    },
    ]
    component.Arrange()
    fixture.detectChanges()
    expect(component.frontRowPowers.length).toBe(4)
    expect(component.frontRowPowers.length).toBe(3)

  })
});
