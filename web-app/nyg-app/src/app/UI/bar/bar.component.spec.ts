import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BarComponent } from './bar.component';
import { DebugElement } from '@angular/core';
import { OnChanceDirective } from './on-chance.directive';

describe('BarComponent', () => {
  let component: BarComponent;
  let fixture: ComponentFixture<BarComponent>;
  let hostEl: DebugElement

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BarComponent, OnChanceDirective]
    })
      .compileComponents();

    fixture = TestBed.createComponent(BarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    hostEl = fixture.nativeElement.querySelector('div')
  });
  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should set the bg color', () => {
    component.NYGTeamName = 'blue'
    component.OCScurrentChance = 0
    fixture.detectChanges()
    const bg = getComputedStyle(hostEl.nativeElement,).getPropertyValue('-bg').trim()
    expect(bg).toBe('blue')
  })
});
